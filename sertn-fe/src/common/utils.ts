import { PublicClient, getContract } from "viem";
import { ModelDBAddress, backendUrl } from "./constants";
import {
  InferenceResponse,
  Log,
  LogLevel,
  NewTaskArgs,
  TaskChallengedArgs,
  TaskChallengedSuccessfullyArgs,
  TaskChallengedUnsuccessfullyArgs,
  TaskRespondedWithProofArgs,
  TaskResponseArgs,
} from "./types";
import { ModelDBAbi } from "../ABI/ModelDB.abi";

export const handleEvents = (
  logs: Log[],
  logger: {
    log: (
      level: LogLevel,
      message: string,
      index: number,
      blockNumber: string
    ) => void;
  }
) => {
  logs.map((log: Log) => {
    const blockNumber = log.blockNumber.toString();
    if (log.eventName === "NewTaskCreated") {
      const args = log.args as NewTaskArgs;
      logger.log(
        "INFO",
        "New Task Created. " + " ZK Inputs: " + args.task.inputs,
        args.taskIndex,
        blockNumber
      );
    }
    if (log.eventName === "TaskResponded") {
      const args = log.args as TaskResponseArgs;
      logger.log(
        "INFO",
        "New Task Response. " + "Output: " + args.taskResponse.output,
        args.taskResponse.referenceTaskIndex,
        blockNumber
      );
    }
    if (log.eventName === "TaskChallenged") {
      const args = log.args as TaskChallengedArgs;
      logger.log("WARN", "Task Challenged", args.taskIndex, blockNumber);
    }
    if (log.eventName === "TaskChallengedUnsuccessfully") {
      const args = log.args as TaskChallengedUnsuccessfullyArgs;
      logger.log(
        "SUCCESS",
        "Task Proven. Prover: " + args.prover,
        args.taskIndex,
        blockNumber
      );
    }
    if (log.eventName === "TaskChallengedSuccessfully") {
      const args = log.args as TaskChallengedSuccessfullyArgs;
      logger.log(
        "ERROR",
        "Task Challenge Confirmed. Challenger: " + args.challenger,
        args.taskIndex,
        blockNumber
      );
    }
    if (log.eventName === "TaskRespondedWithProof") {
      const args = log.args as TaskRespondedWithProofArgs;
      logger.log(
        "SUCCESS",
        "Task Responded With Proof. Confirmed Output: " +
          args.output.toString() +
          ". Proven by: " +
          args.prover +
          ".",
        args.taskIndex,
        blockNumber
      );
    }
  });
};

export const fetchInference = async (
  inputs: string,
  provenOnResponse: boolean,
  modelVerifier = "0xb63A6908A6cd558799E17B92E10E843d254248F2"
) => {
  const response = await fetch(
    backendUrl +
      `?inputs=${inputs}&provenOnResponse=${provenOnResponse}&modelVerifier=${modelVerifier}`
  );
  const text = await response.text();
  return JSON.parse(text) as InferenceResponse;
};

export const getModels = async (client: PublicClient) => {
  const modelContract = getContract({
    address: ModelDBAddress,
    abi: ModelDBAbi,
    client,
  });

  const latest = await modelContract.read.currentModelId();
  const models = [];

  for (let i = 0n; i < latest; i++) {
    let address = await modelContract.read.modelAddresses([i]);
    let model = await modelContract.read.modelVerifiers([address]);
    const [name, description, _] = model;
    models.push({ name, description, address });
  }

  return models;
};
