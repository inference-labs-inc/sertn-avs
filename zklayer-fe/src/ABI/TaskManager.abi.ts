export const zkLayer = {
  abi: [
    {
      type: "constructor",
      inputs: [
        {
          name: "_registryCoordinator",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
        {
          name: "_taskResponseWindowBlock",
          type: "uint32",
          internalType: "uint32",
        },
      ],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "TASK_CHALLENGE_WINDOW_BLOCK",
      inputs: [],
      outputs: [{ name: "", type: "uint32", internalType: "uint32" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "TASK_RESPONSE_WINDOW_BLOCK",
      inputs: [],
      outputs: [{ name: "", type: "uint32", internalType: "uint32" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "aggregator",
      inputs: [],
      outputs: [{ name: "", type: "address", internalType: "address" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "allTaskHashes",
      inputs: [{ name: "", type: "uint32", internalType: "uint32" }],
      outputs: [{ name: "", type: "bytes32", internalType: "bytes32" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "allTaskResponses",
      inputs: [{ name: "", type: "uint32", internalType: "uint32" }],
      outputs: [{ name: "", type: "bytes32", internalType: "bytes32" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "blsApkRegistry",
      inputs: [],
      outputs: [
        { name: "", type: "address", internalType: "contract IBLSApkRegistry" },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "challengeInstances",
      inputs: [
        { name: "id", type: "uint32", internalType: "uint32" },
        { name: "index", type: "uint256", internalType: "uint256" },
      ],
      outputs: [{ name: "", type: "uint256", internalType: "uint256" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "checkSignatures",
      inputs: [
        { name: "msgHash", type: "bytes32", internalType: "bytes32" },
        { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
        {
          name: "referenceBlockNumber",
          type: "uint32",
          internalType: "uint32",
        },
        {
          name: "params",
          type: "tuple",
          internalType:
            "struct IBLSSignatureChecker.NonSignerStakesAndSignature",
          components: [
            {
              name: "nonSignerQuorumBitmapIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "nonSignerPubkeys",
              type: "tuple[]",
              internalType: "struct BN254.G1Point[]",
              components: [
                { name: "X", type: "uint256", internalType: "uint256" },
                { name: "Y", type: "uint256", internalType: "uint256" },
              ],
            },
            {
              name: "quorumApks",
              type: "tuple[]",
              internalType: "struct BN254.G1Point[]",
              components: [
                { name: "X", type: "uint256", internalType: "uint256" },
                { name: "Y", type: "uint256", internalType: "uint256" },
              ],
            },
            {
              name: "apkG2",
              type: "tuple",
              internalType: "struct BN254.G2Point",
              components: [
                { name: "X", type: "uint256[2]", internalType: "uint256[2]" },
                { name: "Y", type: "uint256[2]", internalType: "uint256[2]" },
              ],
            },
            {
              name: "sigma",
              type: "tuple",
              internalType: "struct BN254.G1Point",
              components: [
                { name: "X", type: "uint256", internalType: "uint256" },
                { name: "Y", type: "uint256", internalType: "uint256" },
              ],
            },
            {
              name: "quorumApkIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "totalStakeIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "nonSignerStakeIndices",
              type: "uint32[][]",
              internalType: "uint32[][]",
            },
          ],
        },
      ],
      outputs: [
        {
          name: "",
          type: "tuple",
          internalType: "struct IBLSSignatureChecker.QuorumStakeTotals",
          components: [
            {
              name: "signedStakeForQuorum",
              type: "uint96[]",
              internalType: "uint96[]",
            },
            {
              name: "totalStakeForQuorum",
              type: "uint96[]",
              internalType: "uint96[]",
            },
          ],
        },
        { name: "", type: "bytes32", internalType: "bytes32" },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "confirmChallenge",
      inputs: [
        {
          name: "task",
          type: "tuple",
          internalType: "struct IOmronTaskManager.Task",
          components: [
            { name: "inputs", type: "uint256[5]", internalType: "uint256[5]" },
            {
              name: "taskCreatedBlock",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
            {
              name: "quorumThresholdPercentage",
              type: "uint32",
              internalType: "uint32",
            },
          ],
        },
        {
          name: "taskResponse",
          type: "tuple",
          internalType: "struct IOmronTaskManager.TaskResponse",
          components: [
            {
              name: "referenceTaskIndex",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "output", type: "uint256", internalType: "uint256" },
          ],
        },
        {
          name: "taskResponseMetadata",
          type: "tuple",
          internalType: "struct IOmronTaskManager.TaskResponseMetadata",
          components: [
            {
              name: "taskResponsedBlock",
              type: "uint32",
              internalType: "uint32",
            },
            {
              name: "hashOfNonSigners",
              type: "bytes32",
              internalType: "bytes32",
            },
          ],
        },
        {
          name: "pubkeysOfNonSigningOperators",
          type: "tuple[]",
          internalType: "struct BN254.G1Point[]",
          components: [
            { name: "X", type: "uint256", internalType: "uint256" },
            { name: "Y", type: "uint256", internalType: "uint256" },
          ],
        },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "createNewTask",
      inputs: [
        { name: "inputs", type: "uint256[5]", internalType: "uint256[5]" },
        {
          name: "quorumThresholdPercentage",
          type: "uint32",
          internalType: "uint32",
        },
        { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "delegation",
      inputs: [],
      outputs: [
        {
          name: "",
          type: "address",
          internalType: "contract IDelegationManager",
        },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "generator",
      inputs: [],
      outputs: [{ name: "", type: "address", internalType: "address" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "getBatchOperatorFromId",
      inputs: [
        {
          name: "registryCoordinator",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
        { name: "operatorIds", type: "bytes32[]", internalType: "bytes32[]" },
      ],
      outputs: [
        { name: "operators", type: "address[]", internalType: "address[]" },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "getBatchOperatorId",
      inputs: [
        {
          name: "registryCoordinator",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
        { name: "operators", type: "address[]", internalType: "address[]" },
      ],
      outputs: [
        { name: "operatorIds", type: "bytes32[]", internalType: "bytes32[]" },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "getCheckSignaturesIndices",
      inputs: [
        {
          name: "registryCoordinator",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
        {
          name: "referenceBlockNumber",
          type: "uint32",
          internalType: "uint32",
        },
        { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
        {
          name: "nonSignerOperatorIds",
          type: "bytes32[]",
          internalType: "bytes32[]",
        },
      ],
      outputs: [
        {
          name: "",
          type: "tuple",
          internalType: "struct OperatorStateRetriever.CheckSignaturesIndices",
          components: [
            {
              name: "nonSignerQuorumBitmapIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "quorumApkIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "totalStakeIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "nonSignerStakeIndices",
              type: "uint32[][]",
              internalType: "uint32[][]",
            },
          ],
        },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "getOperatorState",
      inputs: [
        {
          name: "registryCoordinator",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
        { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
        { name: "blockNumber", type: "uint32", internalType: "uint32" },
      ],
      outputs: [
        {
          name: "",
          type: "tuple[][]",
          internalType: "struct OperatorStateRetriever.Operator[][]",
          components: [
            { name: "operator", type: "address", internalType: "address" },
            { name: "operatorId", type: "bytes32", internalType: "bytes32" },
            { name: "stake", type: "uint96", internalType: "uint96" },
          ],
        },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "getOperatorState",
      inputs: [
        {
          name: "registryCoordinator",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
        { name: "operatorId", type: "bytes32", internalType: "bytes32" },
        { name: "blockNumber", type: "uint32", internalType: "uint32" },
      ],
      outputs: [
        { name: "", type: "uint256", internalType: "uint256" },
        {
          name: "",
          type: "tuple[][]",
          internalType: "struct OperatorStateRetriever.Operator[][]",
          components: [
            { name: "operator", type: "address", internalType: "address" },
            { name: "operatorId", type: "bytes32", internalType: "bytes32" },
            { name: "stake", type: "uint96", internalType: "uint96" },
          ],
        },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "getQuorumBitmapsAtBlockNumber",
      inputs: [
        {
          name: "registryCoordinator",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
        { name: "operatorIds", type: "bytes32[]", internalType: "bytes32[]" },
        { name: "blockNumber", type: "uint32", internalType: "uint32" },
      ],
      outputs: [{ name: "", type: "uint256[]", internalType: "uint256[]" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "getTaskResponseWindowBlock",
      inputs: [],
      outputs: [{ name: "", type: "uint32", internalType: "uint32" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "initialize",
      inputs: [
        {
          name: "_pauserRegistry",
          type: "address",
          internalType: "contract IPauserRegistry",
        },
        { name: "initialOwner", type: "address", internalType: "address" },
        { name: "_aggregator", type: "address", internalType: "address" },
        { name: "_generator", type: "address", internalType: "address" },
        { name: "_inferenceDB", type: "address", internalType: "address" },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "latestTaskNum",
      inputs: [],
      outputs: [{ name: "", type: "uint32", internalType: "uint32" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "owner",
      inputs: [],
      outputs: [{ name: "", type: "address", internalType: "address" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "pause",
      inputs: [
        { name: "newPausedStatus", type: "uint256", internalType: "uint256" },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "pauseAll",
      inputs: [],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "paused",
      inputs: [{ name: "index", type: "uint8", internalType: "uint8" }],
      outputs: [{ name: "", type: "bool", internalType: "bool" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "paused",
      inputs: [],
      outputs: [{ name: "", type: "uint256", internalType: "uint256" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "pauserRegistry",
      inputs: [],
      outputs: [
        { name: "", type: "address", internalType: "contract IPauserRegistry" },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "proveResultAccurate",
      inputs: [
        { name: "taskId", type: "uint32", internalType: "uint32" },
        { name: "instances", type: "uint256[]", internalType: "uint256[]" },
        { name: "proof", type: "bytes", internalType: "bytes" },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "raiseChallenge",
      inputs: [
        {
          name: "task",
          type: "tuple",
          internalType: "struct IOmronTaskManager.Task",
          components: [
            { name: "inputs", type: "uint256[5]", internalType: "uint256[5]" },
            {
              name: "taskCreatedBlock",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
            {
              name: "quorumThresholdPercentage",
              type: "uint32",
              internalType: "uint32",
            },
          ],
        },
        {
          name: "taskResponse",
          type: "tuple",
          internalType: "struct IOmronTaskManager.TaskResponse",
          components: [
            {
              name: "referenceTaskIndex",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "output", type: "uint256", internalType: "uint256" },
          ],
        },
        {
          name: "taskResponseMetadata",
          type: "tuple",
          internalType: "struct IOmronTaskManager.TaskResponseMetadata",
          components: [
            {
              name: "taskResponsedBlock",
              type: "uint32",
              internalType: "uint32",
            },
            {
              name: "hashOfNonSigners",
              type: "bytes32",
              internalType: "bytes32",
            },
          ],
        },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "registryCoordinator",
      inputs: [],
      outputs: [
        {
          name: "",
          type: "address",
          internalType: "contract IRegistryCoordinator",
        },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "renounceOwnership",
      inputs: [],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "respondToTask",
      inputs: [
        {
          name: "task",
          type: "tuple",
          internalType: "struct IOmronTaskManager.Task",
          components: [
            { name: "inputs", type: "uint256[5]", internalType: "uint256[5]" },
            {
              name: "taskCreatedBlock",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
            {
              name: "quorumThresholdPercentage",
              type: "uint32",
              internalType: "uint32",
            },
          ],
        },
        {
          name: "taskResponse",
          type: "tuple",
          internalType: "struct IOmronTaskManager.TaskResponse",
          components: [
            {
              name: "referenceTaskIndex",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "output", type: "uint256", internalType: "uint256" },
          ],
        },
        {
          name: "nonSignerStakesAndSignature",
          type: "tuple",
          internalType:
            "struct IBLSSignatureChecker.NonSignerStakesAndSignature",
          components: [
            {
              name: "nonSignerQuorumBitmapIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "nonSignerPubkeys",
              type: "tuple[]",
              internalType: "struct BN254.G1Point[]",
              components: [
                { name: "X", type: "uint256", internalType: "uint256" },
                { name: "Y", type: "uint256", internalType: "uint256" },
              ],
            },
            {
              name: "quorumApks",
              type: "tuple[]",
              internalType: "struct BN254.G1Point[]",
              components: [
                { name: "X", type: "uint256", internalType: "uint256" },
                { name: "Y", type: "uint256", internalType: "uint256" },
              ],
            },
            {
              name: "apkG2",
              type: "tuple",
              internalType: "struct BN254.G2Point",
              components: [
                { name: "X", type: "uint256[2]", internalType: "uint256[2]" },
                { name: "Y", type: "uint256[2]", internalType: "uint256[2]" },
              ],
            },
            {
              name: "sigma",
              type: "tuple",
              internalType: "struct BN254.G1Point",
              components: [
                { name: "X", type: "uint256", internalType: "uint256" },
                { name: "Y", type: "uint256", internalType: "uint256" },
              ],
            },
            {
              name: "quorumApkIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "totalStakeIndices",
              type: "uint32[]",
              internalType: "uint32[]",
            },
            {
              name: "nonSignerStakeIndices",
              type: "uint32[][]",
              internalType: "uint32[][]",
            },
          ],
        },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "setPauserRegistry",
      inputs: [
        {
          name: "newPauserRegistry",
          type: "address",
          internalType: "contract IPauserRegistry",
        },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "setStaleStakesForbidden",
      inputs: [{ name: "value", type: "bool", internalType: "bool" }],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "stakeRegistry",
      inputs: [],
      outputs: [
        { name: "", type: "address", internalType: "contract IStakeRegistry" },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "staleStakesForbidden",
      inputs: [],
      outputs: [{ name: "", type: "bool", internalType: "bool" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "taskNumber",
      inputs: [],
      outputs: [{ name: "", type: "uint32", internalType: "uint32" }],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "transferOwnership",
      inputs: [{ name: "newOwner", type: "address", internalType: "address" }],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "function",
      name: "trySignatureAndApkVerification",
      inputs: [
        { name: "msgHash", type: "bytes32", internalType: "bytes32" },
        {
          name: "apk",
          type: "tuple",
          internalType: "struct BN254.G1Point",
          components: [
            { name: "X", type: "uint256", internalType: "uint256" },
            { name: "Y", type: "uint256", internalType: "uint256" },
          ],
        },
        {
          name: "apkG2",
          type: "tuple",
          internalType: "struct BN254.G2Point",
          components: [
            { name: "X", type: "uint256[2]", internalType: "uint256[2]" },
            { name: "Y", type: "uint256[2]", internalType: "uint256[2]" },
          ],
        },
        {
          name: "sigma",
          type: "tuple",
          internalType: "struct BN254.G1Point",
          components: [
            { name: "X", type: "uint256", internalType: "uint256" },
            { name: "Y", type: "uint256", internalType: "uint256" },
          ],
        },
      ],
      outputs: [
        { name: "pairingSuccessful", type: "bool", internalType: "bool" },
        { name: "siganatureIsValid", type: "bool", internalType: "bool" },
      ],
      stateMutability: "view",
    },
    {
      type: "function",
      name: "unpause",
      inputs: [
        { name: "newPausedStatus", type: "uint256", internalType: "uint256" },
      ],
      outputs: [],
      stateMutability: "nonpayable",
    },
    {
      type: "event",
      name: "Initialized",
      inputs: [
        {
          name: "version",
          type: "uint8",
          indexed: false,
          internalType: "uint8",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "NewTaskCreated",
      inputs: [
        {
          name: "taskIndex",
          type: "uint32",
          indexed: true,
          internalType: "uint32",
        },
        {
          name: "task",
          type: "tuple",
          indexed: false,
          internalType: "struct IOmronTaskManager.Task",
          components: [
            { name: "inputs", type: "uint256[5]", internalType: "uint256[5]" },
            {
              name: "taskCreatedBlock",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "quorumNumbers", type: "bytes", internalType: "bytes" },
            {
              name: "quorumThresholdPercentage",
              type: "uint32",
              internalType: "uint32",
            },
          ],
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "OwnershipTransferred",
      inputs: [
        {
          name: "previousOwner",
          type: "address",
          indexed: true,
          internalType: "address",
        },
        {
          name: "newOwner",
          type: "address",
          indexed: true,
          internalType: "address",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "Paused",
      inputs: [
        {
          name: "account",
          type: "address",
          indexed: true,
          internalType: "address",
        },
        {
          name: "newPausedStatus",
          type: "uint256",
          indexed: false,
          internalType: "uint256",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "PauserRegistrySet",
      inputs: [
        {
          name: "pauserRegistry",
          type: "address",
          indexed: false,
          internalType: "contract IPauserRegistry",
        },
        {
          name: "newPauserRegistry",
          type: "address",
          indexed: false,
          internalType: "contract IPauserRegistry",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "StaleStakesForbiddenUpdate",
      inputs: [
        { name: "value", type: "bool", indexed: false, internalType: "bool" },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "TaskChallenged",
      inputs: [
        {
          name: "taskIndex",
          type: "uint32",
          indexed: true,
          internalType: "uint32",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "TaskChallengedSuccessfully",
      inputs: [
        {
          name: "taskIndex",
          type: "uint32",
          indexed: true,
          internalType: "uint32",
        },
        {
          name: "challenger",
          type: "address",
          indexed: true,
          internalType: "address",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "TaskChallengedUnsuccessfully",
      inputs: [
        {
          name: "taskIndex",
          type: "uint32",
          indexed: true,
          internalType: "uint32",
        },
        {
          name: "prover",
          type: "address",
          indexed: true,
          internalType: "address",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "TaskCompleted",
      inputs: [
        {
          name: "taskIndex",
          type: "uint32",
          indexed: true,
          internalType: "uint32",
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "TaskResponded",
      inputs: [
        {
          name: "taskResponse",
          type: "tuple",
          indexed: false,
          internalType: "struct IOmronTaskManager.TaskResponse",
          components: [
            {
              name: "referenceTaskIndex",
              type: "uint32",
              internalType: "uint32",
            },
            { name: "output", type: "uint256", internalType: "uint256" },
          ],
        },
        {
          name: "taskResponseMetadata",
          type: "tuple",
          indexed: false,
          internalType: "struct IOmronTaskManager.TaskResponseMetadata",
          components: [
            {
              name: "taskResponsedBlock",
              type: "uint32",
              internalType: "uint32",
            },
            {
              name: "hashOfNonSigners",
              type: "bytes32",
              internalType: "bytes32",
            },
          ],
        },
      ],
      anonymous: false,
    },
    {
      type: "event",
      name: "Unpaused",
      inputs: [
        {
          name: "account",
          type: "address",
          indexed: true,
          internalType: "address",
        },
        {
          name: "newPausedStatus",
          type: "uint256",
          indexed: false,
          internalType: "uint256",
        },
      ],
      anonymous: false,
    },
  ],
};
