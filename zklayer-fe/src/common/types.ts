export interface Log {
  eventName:
    | "NewTaskCreated"
    | "TaskResponded"
    | "TaskCompleted"
    | "TaskChallengedSuccessfully"
    | "TaskChallengedUnsuccessfully"
    | "TaskChallenged";
  args:
    | NewTaskArgs
    | TaskResponseArgs
    | TaskChallengedArgs
    | TaskChallengedSuccessfullyArgs
    | TaskChallengedUnsuccessfullyArgs;
  address: string;
  topics: string[];
  data: string;
  blockHash: string;
  blockNumber: string;
  transactionHash: string;
  transactionIndex: number;
  logIndex: number;
  transactionLogIndex: string;
  removed: boolean;
}

export interface NewTaskArgs {
  taskIndex: number;
  task: Task;
}

export interface TaskChallengedArgs {
  taskIndex: number;
}

export interface TaskChallengedSuccessfullyArgs {
  challenger: string;
  taskIndex: number;
}

export interface TaskChallengedUnsuccessfullyArgs {
  taskIndex: number;
  prover: string;
}

export interface Task {
  inputs: string[];
  taskCreatedBlock: number;
  quorumNumbers: string;
  quorumThresholdPercentage: number;
}

export interface TaskResponseArgs {
  taskResponse: TaskResponse;
  taskResponseMetadata: TaskResponseMetadata;
}

export interface TaskResponse {
  referenceTaskIndex: number;
  output: string;
}

export interface TaskResponseMetadata {
  taskResponsedBlock: number;
  hashOfNonSigners: string;
}

export type InferenceResponse = {
  inputs: string;
  taskIndex: number;
};
export const ProofTypes = ["Challenge", "Pre Prove"] as const;

export type InferenceForm = {
  elements: {
    type: {
      value: string;
    };
    input: {
      value: string;
    };
  };
};

export const LogStyles = {
  ERROR: "text-rose-400 font-normal",
  INFO: "text-teal-400 font-normal",
  WARN: "text-amber-400 font-normal",
  SUCCESS: "text-emerald-400 font-normal",
} as const;

export type LogLevel = "ERROR" | "WARN" | "INFO" | "SUCCESS";

export type LogMessage = {
  taskIndex: number;
  level: LogLevel;
  message: string;
  blockNumber: string;
};
