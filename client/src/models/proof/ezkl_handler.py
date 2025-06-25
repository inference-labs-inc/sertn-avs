import json
import os
import subprocess
import traceback
from enum import Enum
from typing import Optional, Iterable

import ezkl

from common.console import console, styles
from common.constants import MODELS_FOLDER, PROOFS_FOLDER

LOCAL_EZKL_PATH = os.path.join(os.path.expanduser("~"), ".ezkl", "ezkl")


class EZKLInputType(Enum):
    F16 = ezkl.PyInputType.F16
    F32 = ezkl.PyInputType.F32
    F64 = ezkl.PyInputType.F64
    Int = ezkl.PyInputType.Int
    Bool = ezkl.PyInputType.Bool
    TDim = ezkl.PyInputType.TDim


class EZKLHandler:
    """
    Handler for the EZKL proof system.
    This class provides methods for generating and verifying proofs using EZKL.
    """

    def __init__(self, model_id: str, task_id: str, inputs: Optional[list[float]]):
        self.input_data = inputs

        model_path = os.path.join(MODELS_FOLDER, f"{model_id}")
        self.compiled_model_path = os.path.join(model_path, "model.compiled")
        self.pk_path = os.path.join(model_path, "pk.key")  # TODO: ...
        self.vk_path = os.path.join(model_path, "vk.key")
        self.settings_path = os.path.join(model_path, "settings.json")
        self.settings = json.load(open(self.settings_path, "r", encoding="utf-8"))

        self.proof_dir = os.path.join(PROOFS_FOLDER, task_id)
        self.input_path = os.path.join(self.proof_dir, "inputs.json")
        self.witness_path = os.path.join(self.proof_dir, "witness.json")
        self.proof_filepath = os.path.join(self.proof_dir, "proof.json")

    def gen_input_file(self):
        console.print("Generating input file", style=styles.op_info)
        if isinstance(self.input_data, list):
            input_data = self.input_data
        else:
            input_data = self.input_data.to_array()
        data = {"input_data": [input_data]}
        os.makedirs(os.path.dirname(self.input_path), exist_ok=True)
        with open(self.input_path, "w", encoding="utf-8") as f:
            json.dump(data, f)
        console.print(f"Generated input.json with data: {data}", style=styles.op_info)

    def gen_proof(self) -> tuple[str, str]:
        try:
            console.print("Starting proof generation...", style=styles.op_info)

            self.generate_witness()
            console.print("Generating proof", style=styles.debug)

            result = subprocess.run(
                [
                    LOCAL_EZKL_PATH,
                    "prove",
                    "--witness",
                    self.witness_path,
                    "--compiled-circuit",
                    self.compiled_model_path,
                    "--pk-path",
                    self.pk_path,
                    "--proof-path",
                    self.proof_filepath,
                ],
                check=True,
                capture_output=True,
                text=True,
            )

            console.print(
                f"Proof generated: {self.proof_filepath}, result: {result.stdout}",
                style=styles.op_info,
            )

            with open(self.proof_filepath, "r", encoding="utf-8") as f:
                proof = json.load(f)

            return json.dumps(proof), json.dumps(proof["instances"])

        except Exception as e:
            console.print(
                f"An error occurred during proof generation: {e}", style=styles.error
            )
            traceback.print_exc()
            raise

    def verify_proof(
        self,
        validator_inputs: Iterable[float],
        proof: str | dict,
    ) -> bool:
        if not proof:
            return False

        if isinstance(proof, str):
            proof_json = json.loads(proof)
        else:
            proof_json = proof

        input_instances = self.translate_inputs_to_instances(validator_inputs)
        proof_json["instances"] = [
            input_instances[:] + proof_json["instances"][0][len(input_instances) :]
        ]
        proof_json["transcript_type"] = "EVM"

        with open(self.proof_filepath, "w", encoding="utf-8") as f:
            json.dump(proof_json, f)

        try:
            result = subprocess.run(
                [
                    LOCAL_EZKL_PATH,
                    "verify",
                    "--settings-path",
                    self.settings_path,
                    "--proof-path",
                    self.proof_filepath,
                    "--vk-path",
                    self.vk_path,
                ],
                check=True,
                capture_output=True,
                text=True,
                timeout=600,
            )
            return "verified: true" in result.stdout
        except subprocess.TimeoutExpired:
            console.print(
                "Verification process timed out after 60 seconds", style=styles.warning
            )
            return False
        except subprocess.CalledProcessError:
            return False

    def generate_witness(self, return_content: bool = False) -> list | dict:
        console.print("Generating witness...", style=styles.debug)

        result = subprocess.run(
            [
                LOCAL_EZKL_PATH,
                "gen-witness",
                "--data",
                self.input_path,
                "--compiled-circuit",
                self.compiled_model_path,
                "--output",
                self.witness_path,
                "--vk-path",
                self.vk_path,
            ],
            check=True,
            capture_output=True,
            text=True,
        )

        console.print(f"Gen witness result: {result.stdout}", style=styles.debug)

        if return_content:
            with open(self.witness_path, "r", encoding="utf-8") as f:
                return json.load(f)
        return result.stdout

    def translate_inputs_to_instances(self, validator_inputs) -> list[int]:
        scale_map = self.settings.get("model_input_scales", [])
        type_map = self.settings.get("input_types", [])
        return [
            ezkl.float_to_felt(x, scale_map[i], EZKLInputType[type_map[i]].value)
            for i, arr in enumerate(validator_inputs)
            for x in arr
        ]
