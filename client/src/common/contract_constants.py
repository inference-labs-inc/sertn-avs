from enum import IntEnum


class TaskStructMap(IntEnum):
    """
    Task model representation, declared in `contracts/interfaces/ISertnTaskManager.sol`
    Basically it's just a fields order in the `ISertnTaskManager.Task` struct.
    """

    START_BLOCK = 0
    MODEL_ID = 1
    INPUTS = 2
    PROOF_HASH = 3
    USER = 4
    NONCE = 5
    OPERATOR = 6
    STATE = 7
    OUTPUT = 8
    FEE = 9


class TaskStateMap(IntEnum):
    """
    Python representation for task states, declared in `contracts/interfaces/ISertnTaskManager.sol`
    `ISertnTaskManager.TaskState`
    """

    CREATED = 0
    ASSIGNED = 1
    COMPLETED = 2
    CHALLENGED = 3
    REJECTED = 4
    RESOLVED = 5

    def __str__(self):
        return self.name

    @classmethod
    def from_int(cls, value: int):
        """Convert integer from blockchain to TaskState"""
        try:
            return cls(value)
        except ValueError:
            raise ValueError(f"Invalid task state value: {value}")


class VerificationStrategiesMap(IntEnum):
    """
    Python representation of `VerificationStrategy` enum.
    """

    NONE = 0
    ONCHAIN = 1
    OFFCHAIN = 2
