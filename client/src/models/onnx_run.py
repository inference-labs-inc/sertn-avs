import os

import numpy as np
import onnxruntime as ort

from common.constants import MODELS_FOLDER


def run_onnx(model_id: str, input_data: list[float]) -> list[float]:
    """
    Run an ONNX model with the given input data.
    Args:
        model_path (str): Path to the ONNX model file.
        input_data (list[float]): Input data for the model.
    Returns:
        np.ndarray: The output of the model as a NumPy array.
    """
    # Path to the ONNX model
    model_path = os.path.join(MODELS_FOLDER, f"model_{model_id}", "network.onnx")
    if not os.path.exists(model_path):
        raise FileNotFoundError(f"Model file not found: {model_path}")

    # Load the ONNX model
    session = ort.InferenceSession(model_path)

    # Get input and output names
    input_name = session.get_inputs()[0].name
    output_names = [output.name for output in session.get_outputs()]
    options = ort.RunOptions()
    options.log_severity_level = 3

    # Run the model
    outputs = session.run(
        output_names, {input_name: [[[el] for el in input_data]]}, options
    )

    flattened = []
    for out in outputs:
        flattened.extend(out.flatten())

    return np.array(flattened).tolist()
