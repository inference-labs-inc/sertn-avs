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
      name: "raiseChallenger",
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
  bytecode: {
    object:
      "0x6101206040523480156200001257600080fd5b5060405162006232380380620062328339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615f3b620002f7600039600081816102e20152818161063101526133e70152600081816105fa0152612ab60152600081816104b3015281816116790152612c980152600081816104ed01528181612e6e01526130300152600081816105140152818161139d015281816127a0015281816129190152612b530152615f3b6000f3fe608060405234801561001057600080fd5b50600436106102485760003560e01c80635c975abb1161013b5780638b00ce7c116100b8578063df5cf7231161007c578063df5cf723146105f5578063f2fde38b1461061c578063f5c9899d1461062f578063f63c5bab14610655578063fabc1cbc1461065d57600080fd5b80638b00ce7c146105935780638da5cb5b146105a3578063b98d0908146105b4578063bdeea6cc146105c1578063cefdc1d4146105d457600080fd5b80636efb4636116100ff5780636efb463614610536578063715018a61461055757806372d18e8d1461055f5780637afa1eed1461056d578063886f11951461058057600080fd5b80635c975abb146104a65780635df45946146104ae57806367d6be44146104d557806368304835146104e85780636d14a9871461050f57600080fd5b806331b36bd9116101c95780634f739f741161018d5780634f739f741461041857806358a7cd2614610438578063595c6a671461044b5780635ac86ab7146104535780635c1556621461048657600080fd5b806331b36bd9146103925780633563b0d1146103b2578063416c7e5e146103d2578063480bab6b146103e55780634d2b57fe146103f857600080fd5b8063171f1d5b11610210578063171f1d5b146102ae5780631ad43189146102dd578063245a7bfc146103195780632cb223d5146103445780632d89f6fc1461037257600080fd5b80630627721e1461024d57806310d67a2f14610262578063136439dd146102755780631459457a14610288578063162d8a3f1461029b575b600080fd5b61026061025b366004614783565b610670565b005b6102606102703660046147ff565b6107b9565b61026061028336600461481c565b610875565b610260610296366004614835565b6109b4565b6102606102a93660046148d1565b610b13565b6102c16102bc366004614a95565b610c5f565b6040805192151583529015156020830152015b60405180910390f35b6103047f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102d4565b60cc5461032c906001600160a01b031681565b6040516001600160a01b0390911681526020016102d4565b610364610352366004614ae6565b60cb6020526000908152604090205481565b6040519081526020016102d4565b610364610380366004614ae6565b60ca6020526000908152604090205481565b6103a56103a0366004614b26565b610de9565b6040516102d49190614c14565b6103c56103c0366004614c27565b610f05565b6040516102d49190614d79565b6102606103e0366004614d9a565b61139b565b6102606103f3366004614e26565b611510565b61040b610406366004614f07565b6117fc565b6040516102d49190614f56565b61042b610426366004614fe7565b611911565b6040516102d491906150b1565b61026061044636600461516c565b612037565b6102606120db565b6104766104613660046151fd565b606654600160ff9092169190911b9081161490565b60405190151581526020016102d4565b61049961049436600461521a565b6121a2565b6040516102d4919061527d565b606654610364565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b6103646104e33660046152b5565b61236a565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b6105496105443660046154ee565b6123ed565b6040516102d49291906155ae565b6102606132e5565b60c95463ffffffff16610304565b60cd5461032c906001600160a01b031681565b60655461032c906001600160a01b031681565b60c9546103049063ffffffff1681565b6033546001600160a01b031661032c565b6097546104769060ff1681565b6102606105cf3660046155f7565b6132f9565b6105e76105e236600461566b565b6135b1565b6040516102d49291906156a2565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b61026061062a3660046147ff565b613743565b7f0000000000000000000000000000000000000000000000000000000000000000610304565b610304606481565b61026061066b36600461481c565b6137b9565b60cd546001600160a01b0316331461068757600080fd5b61068f6145f8565b6040805160a081810190925290869060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f8501839004830281018301909152838152908490849081908401838280828437600092019190915250505050604080830191909152516107199082906020016156c3565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f54977349061077c9084906156c3565b60405180910390a260c9546107989063ffffffff16600161578f565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561080c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083091906157b7565b6001600160a01b0316336001600160a01b0316146108695760405162461bcd60e51b8152600401610860906157d4565b60405180910390fd5b61087281613915565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156108bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e1919061581e565b6108fd5760405162461bcd60e51b81526004016108609061583b565b606654818116146109765760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610860565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156109d45750600054600160ff909116105b806109ee5750303b1580156109ee575060005460ff166001145b610a515760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610860565b6000805460ff191660011790558015610a74576000805461ff0019166101001790555b610a7f866000613a0c565b610a8885613af6565b60cc80546001600160a01b038087166001600160a01b03199283161790925560cd805486841690831617905560ce8054928516929091169190911790558015610b0b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6000610b226020840184614ae6565b63ffffffff8116600090815260cb6020526040902054909150610b4457600080fd5b8282604051602001610b579291906158a1565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb90935291205414610b8d57600080fd5b6064610b9c6020840184614ae6565b610ba6919061578f565b63ffffffff164363ffffffff161115610bbe57600080fd5b60ce54604051631150ef3360e21b81526001600160a01b0390911690634543bccc90610bf690849088906020890135906004016158d7565b600060405180830381600087803b158015610c1057600080fd5b505af1158015610c24573d6000803e3d6000fd5b505060405163ffffffff841692507f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a9150600090a250505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610ca757610ca76158fc565b60200201518951600160200201518a60200151600060028110610ccc57610ccc6158fc565b60200201518b60200151600160028110610ce857610ce86158fc565b602090810291909101518c518d830151604051610d459a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d689190615912565b9050610ddb610d81610d7a8884613b48565b8690613bdf565b610d89613c73565b610dd1610dc285610dbc604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613b48565b610dcb8c613d33565b90613bdf565b886201d4c0613dc3565b909890975095505050505050565b606081516001600160401b03811115610e0457610e04614930565b604051908082528060200260200182016040528015610e2d578160200160208202803683370190505b50905060005b8251811015610efe57836001600160a01b03166313542a4e848381518110610e5d57610e5d6158fc565b60200260200101516040518263ffffffff1660e01b8152600401610e9091906001600160a01b0391909116815260200190565b602060405180830381865afa158015610ead573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed19190615934565b828281518110610ee357610ee36158fc565b6020908102919091010152610ef78161594d565b9050610e33565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6b91906157b7565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd191906157b7565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611013573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103791906157b7565b9050600086516001600160401b0381111561105457611054614930565b60405190808252806020026020018201604052801561108757816020015b60608152602001906001900390816110725790505b50905060005b875181101561138f5760008882815181106110aa576110aa6158fc565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa15801561110b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111339190810190615968565b905080516001600160401b0381111561114e5761114e614930565b60405190808252806020026020018201604052801561119957816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161116c5790505b508484815181106111ac576111ac6158fc565b602002602001018190525060005b8151811015611379576040518060600160405280876001600160a01b03166347b314e88585815181106111ef576111ef6158fc565b60200260200101516040518263ffffffff1660e01b815260040161121591815260200190565b602060405180830381865afa158015611232573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125691906157b7565b6001600160a01b03168152602001838381518110611276576112766158fc565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106112a4576112a46158fc565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611300573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132491906159f8565b6001600160601b0316815250858581518110611342576113426158fc565b6020026020010151828151811061135b5761135b6158fc565b602002602001018190525080806113719061594d565b9150506111ba565b50505080806113879061594d565b91505061108d565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141d91906157b7565b6001600160a01b0316336001600160a01b0316146114c95760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610860565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b600061151f6020850185614ae6565b9050600082516001600160401b0381111561153c5761153c614930565b604051908082528060200260200182016040528015611565578160200160208202803683370190505b50905060005b83518110156115d7576115a8848281518110611589576115896158fc565b6020026020010151805160009081526020918201519091526040902090565b8282815181106115ba576115ba6158fc565b6020908102919091010152806115cf8161594d565b91505061156b565b5060006115ea60c0880160a08901614ae6565b826040516020016115fc929190615a21565b6040516020818303038152906040528051906020012090508460200135811461162457600080fd5b600084516001600160401b0381111561163f5761163f614930565b604051908082528060200260200182016040528015611668578160200160208202803683370190505b50905060005b855181101561175b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106116b8576116b86158fc565b60200260200101516040518263ffffffff1660e01b81526004016116de91815260200190565b602060405180830381865afa1580156116fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171f91906157b7565b828281518110611731576117316158fc565b6001600160a01b0390921660209283029190910190910152806117538161594d565b91505061166e565b5060ce5460405163959b127960e01b815263ffffffff861660048201526001600160a01b039091169063959b127990602401600060405180830381600087803b1580156117a757600080fd5b505af11580156117bb573d6000803e3d6000fd5b505060405133925063ffffffff871691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050565b606081516001600160401b0381111561181757611817614930565b604051908082528060200260200182016040528015611840578160200160208202803683370190505b50905060005b8251811015610efe57836001600160a01b031663296bb064848381518110611870576118706158fc565b60200260200101516040518263ffffffff1660e01b815260040161189691815260200190565b602060405180830381865afa1580156118b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118d791906157b7565b8282815181106118e9576118e96158fc565b6001600160a01b039092166020928302919091019091015261190a8161594d565b9050611846565b61193c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561197c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119a091906157b7565b90506119cd6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906119fd908b9089908990600401615a9f565b600060405180830381865afa158015611a1a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a429190810190615ac8565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611a74908b908b908b90600401615b7f565b600060405180830381865afa158015611a91573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ab99190810190615ac8565b6040820152856001600160401b03811115611ad657611ad6614930565b604051908082528060200260200182016040528015611b0957816020015b6060815260200190600190039081611af45790505b50606082015260005b60ff8116871115611f48576000856001600160401b03811115611b3757611b37614930565b604051908082528060200260200182016040528015611b60578160200160208202803683370190505b5083606001518360ff1681518110611b7a57611b7a6158fc565b602002602001018190525060005b86811015611e485760008c6001600160a01b03166304ec63518a8a85818110611bb357611bb36158fc565b905060200201358e88600001518681518110611bd157611bd16158fc565b60200260200101516040518463ffffffff1660e01b8152600401611c0e9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611c2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4f9190615b9f565b90506001600160c01b038116611cf35760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610860565b8a8a8560ff16818110611d0857611d086158fc565b6001600160c01b03841692013560f81c9190911c600190811614159050611e3557856001600160a01b031663dd9846b98a8a85818110611d4a57611d4a6158fc565b905060200201358d8d8860ff16818110611d6657611d666158fc565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611dbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de09190615bc8565b85606001518560ff1681518110611df957611df96158fc565b60200260200101518481518110611e1257611e126158fc565b63ffffffff9092166020928302919091019091015282611e318161594d565b9350505b5080611e408161594d565b915050611b88565b506000816001600160401b03811115611e6357611e63614930565b604051908082528060200260200182016040528015611e8c578160200160208202803683370190505b50905060005b82811015611f0d5784606001518460ff1681518110611eb357611eb36158fc565b60200260200101518181518110611ecc57611ecc6158fc565b6020026020010151828281518110611ee657611ee66158fc565b63ffffffff9092166020928302919091019091015280611f058161594d565b915050611e92565b508084606001518460ff1681518110611f2857611f286158fc565b602002602001018190525050508080611f4090615be5565b915050611b12565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fad91906157b7565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611fe0908b908b908e90600401615c05565b600060405180830381865afa158015611ffd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120259190810190615ac8565b60208301525098975050505050505050565b60ce54604051632c53e69360e11b81526001600160a01b03909116906358a7cd269061206f9088908890889088908890600401615c2f565b600060405180830381600087803b15801561208957600080fd5b505af115801561209d573d6000803e3d6000fd5b505060405133925063ffffffff881691507ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015612123573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612147919061581e565b6121635760405162461bcd60e51b81526004016108609061583b565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016121d4929190615c6e565b600060405180830381865afa1580156121f1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122199190810190615ac8565b9050600084516001600160401b0381111561223657612236614930565b60405190808252806020026020018201604052801561225f578160200160208202803683370190505b50905060005b855181101561236057866001600160a01b03166304ec635187838151811061228f5761228f6158fc565b6020026020010151878685815181106122aa576122aa6158fc565b60200260200101516040518463ffffffff1660e01b81526004016122e79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612304573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123289190615b9f565b6001600160c01b0316828281518110612343576123436158fc565b6020908102919091010152806123588161594d565b915050612265565b5095945050505050565b60ce546040516319f5af9160e21b815263ffffffff84166004820152602481018390526000916001600160a01b0316906367d6be4490604401602060405180830381865afa1580156123c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e49190615934565b90505b92915050565b60408051808201909152606080825260208201526000846124645760405162461bcd60e51b81526020600482015260376024820152600080516020615ee683398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610860565b6040830151518514801561247c575060a08301515185145b801561248c575060c08301515185145b801561249c575060e08301515185145b6125065760405162461bcd60e51b81526020600482015260416024820152600080516020615ee683398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610860565b8251516020840151511461257e5760405162461bcd60e51b815260206004820152604460248201819052600080516020615ee6833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610860565b4363ffffffff168463ffffffff16106125ed5760405162461bcd60e51b815260206004820152603c6024820152600080516020615ee683398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610860565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561262e5761262e614930565b604051908082528060200260200182016040528015612657578160200160208202803683370190505b506020820152866001600160401b0381111561267557612675614930565b60405190808252806020026020018201604052801561269e578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156126d2576126d2614930565b6040519080825280602002602001820160405280156126fb578160200160208202803683370190505b5081526020860151516001600160401b0381111561271b5761271b614930565b604051908082528060200260200182016040528015612744578160200160208202803683370190505b50816020018190525060006128168a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156127ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128119190615c8d565b613fe7565b905060005b876020015151811015612a925761284188602001518281518110611589576115896158fc565b83602001518281518110612857576128576158fc565b60209081029190910101528015612917576020830151612878600183615caa565b81518110612888576128886158fc565b602002602001015160001c836020015182815181106128a9576128a96158fc565b602002602001015160001c11612917576040805162461bcd60e51b8152602060048201526024810191909152600080516020615ee683398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610860565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061295c5761295c6158fc565b60200260200101518b8b60000151858151811061297b5761297b6158fc565b60200260200101516040518463ffffffff1660e01b81526004016129b89392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156129d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129f99190615b9f565b6001600160c01b031683600001518281518110612a1857612a186158fc565b602002602001018181525050612a7e610d7a612a528486600001518581518110612a4457612a446158fc565b602002602001015116614071565b8a602001518481518110612a6857612a686158fc565b602002602001015161409c90919063ffffffff16565b945080612a8a8161594d565b91505061281b565b5050612a9d83614180565b60975490935060ff16600081612ab4576000612b36565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b369190615934565b905060005b8a8110156131b4578215612c96578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612b9257612b926158fc565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612bd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf69190615934565b612c009190615cc1565b11612c965760405162461bcd60e51b81526020600482015260666024820152600080516020615ee683398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610860565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612cd757612cd76158fc565b9050013560f81c60f81b60f81c8c8c60a001518581518110612cfb57612cfb6158fc565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612d57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d7b9190615cd9565b6001600160401b031916612d9e8a604001518381518110611589576115896158fc565b67ffffffffffffffff191614612e3a5760405162461bcd60e51b81526020600482015260616024820152600080516020615ee683398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610860565b612e6a89604001518281518110612e5357612e536158fc565b602002602001015187613bdf90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612ead57612ead6158fc565b9050013560f81c60f81b60f81c8c8c60c001518581518110612ed157612ed16158fc565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612f2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f5191906159f8565b85602001518281518110612f6757612f676158fc565b6001600160601b03909216602092830291909101820152850151805182908110612f9357612f936158fc565b602002602001015185600001518281518110612fb157612fb16158fc565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561319f5761302986600001518281518110612ffb57612ffb6158fc565b60200260200101518f8f86818110613015576130156158fc565b600192013560f81c9290921c811614919050565b1561318d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061306f5761306f6158fc565b9050013560f81c60f81b60f81c8e89602001518581518110613093576130936158fc565b60200260200101518f60e0015188815181106130b1576130b16158fc565b602002602001015187815181106130ca576130ca6158fc565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561312e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061315291906159f8565b8751805185908110613166576131666158fc565b6020026020010181815161317a9190615d04565b6001600160601b03169052506001909101905b806131978161594d565b915050612fd5565b505080806131ac9061594d565b915050612b3b565b5050506000806131ce8c868a606001518b60800151610c5f565b915091508161323f5760405162461bcd60e51b81526020600482015260436024820152600080516020615ee683398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610860565b806132a05760405162461bcd60e51b81526020600482015260396024820152600080516020615ee683398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610860565b505060008782602001516040516020016132bb929190615a21565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6132ed61421b565b6132f76000613af6565b565b60cc546001600160a01b0316331461331057600080fd5b600061332260c0850160a08601614ae6565b905036600061333460c0870187615d2c565b9092509050600061334c610100880160e08901614ae6565b905060ca600061335f6020890189614ae6565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161338b9190615d72565b60405160208183030381529060405280519060200120146133ab57600080fd5b600060cb816133bd60208a018a614ae6565b63ffffffff1663ffffffff16815260200190815260200160002054146133e257600080fd5b61340c7f00000000000000000000000000000000000000000000000000000000000000008561578f565b63ffffffff164363ffffffff16111561342457600080fd5b6000866040516020016134379190615e1b565b60405160208183030381529060405280519060200120905060008061345f8387878a8c6123ed565b9150915060005b858110156134f3578460ff1683602001518281518110613488576134886158fc565b602002602001015161349a9190615e29565b6001600160601b03166064846000015183815181106134bb576134bb6158fc565b60200260200101516001600160601b03166134d69190615e58565b10156134e157600080fd5b806134eb8161594d565b915050613466565b5060408051808201825263ffffffff43168152602080820184905291519091613520918c91849101615e77565b6040516020818303038152906040528051906020012060cb60008c600001602081019061354d9190614ae6565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a8260405161359c929190615e77565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106135ec576135ec6158fc565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906136289088908690600401615c6e565b600060405180830381865afa158015613645573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261366d9190810190615ac8565b60008151811061367f5761367f6158fc565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156136eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061370f9190615b9f565b6001600160c01b03169050600061372582614275565b9050816137338a838a610f05565b9550955050505050935093915050565b61374b61421b565b6001600160a01b0381166137b05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610860565b61087281613af6565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561380c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061383091906157b7565b6001600160a01b0316336001600160a01b0316146138605760405162461bcd60e51b8152600401610860906157d4565b6066541981196066541916146138de5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610860565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016109a9565b6001600160a01b0381166139a35760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610860565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b0316158015613a2d57506001600160a01b03821615155b613aaf5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610860565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613af282613915565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613b64614626565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613b9757613b99565bfe5b5080613bd75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610860565b505092915050565b6040805180820190915260008082526020820152613bfb614644565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613b97575080613bd75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610860565b613c7b614662565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613d63600080516020615ec683398151915286615912565b90505b613d6f81614341565b9093509150600080516020615ec6833981519152828309831415613da9576040805180820190915290815260208101919091529392505050565b600080516020615ec6833981519152600182089050613d66565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613df5614687565b60005b6002811015613fba576000613e0e826006615e58565b9050848260028110613e2257613e226158fc565b60200201515183613e34836000615cc1565b600c8110613e4457613e446158fc565b6020020152848260028110613e5b57613e5b6158fc565b60200201516020015183826001613e729190615cc1565b600c8110613e8257613e826158fc565b6020020152838260028110613e9957613e996158fc565b6020020151515183613eac836002615cc1565b600c8110613ebc57613ebc6158fc565b6020020152838260028110613ed357613ed36158fc565b6020020151516001602002015183613eec836003615cc1565b600c8110613efc57613efc6158fc565b6020020152838260028110613f1357613f136158fc565b602002015160200151600060028110613f2e57613f2e6158fc565b602002015183613f3f836004615cc1565b600c8110613f4f57613f4f6158fc565b6020020152838260028110613f6657613f666158fc565b602002015160200151600160028110613f8157613f816158fc565b602002015183613f92836005615cc1565b600c8110613fa257613fa26158fc565b60200201525080613fb28161594d565b915050613df8565b50613fc36146a6565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613ff3846143c3565b9050808360ff166001901b116123e45760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610860565b6000805b82156123e757614086600184615caa565b909216918061409481615ea3565b915050614075565b60408051808201909152600080825260208201526102008261ffff16106140f85760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610860565b8161ffff166001141561410c5750816123e7565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061417557600161ffff871660ff83161c81161415614158576141558484613bdf565b93505b6141628384613bdf565b92506201fffe600192831b169101614128565b509195945050505050565b604080518082019091526000808252602082015281511580156141a557506020820151155b156141c3575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615ec683398151915284602001516141f69190615912565b61420e90600080516020615ec6833981519152615caa565b905292915050565b919050565b6033546001600160a01b031633146132f75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610860565b606060008061428384614071565b61ffff166001600160401b0381111561429e5761429e614930565b6040519080825280601f01601f1916602001820160405280156142c8576020820181803683370190505b5090506000805b8251821080156142e0575061010081105b15614337576001811b935085841615614327578060f81b838381518110614309576143096158fc565b60200101906001600160f81b031916908160001a9053508160010191505b6143308161594d565b90506142cf565b5090949350505050565b60008080600080516020615ec68339815191526003600080516020615ec683398151915286600080516020615ec68339815191528889090908905060006143b7827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615ec6833981519152614550565b91959194509092505050565b60006101008251111561444c5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610860565b815161445a57506000919050565b60008083600081518110614470576144706158fc565b0160200151600160f89190911c81901b92505b84518110156145475784818151811061449e5761449e6158fc565b0160200151600160f89190911c1b91508282116145335760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610860565b918117916145408161594d565b9050614483565b50909392505050565b60008061455b6146a6565b6145636146c4565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613b975750826145ed5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610860565b505195945050505050565b604051806080016040528061460b6146e2565b81526000602082018190526060604083018190529091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280614675614700565b8152602001614682614700565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461087257600080fd5b80356142168161471e565b60008083601f84011261474d57600080fd5b5081356001600160401b0381111561476457600080fd5b60208301915083602082850101111561477c57600080fd5b9250929050565b60008060008060e0858703121561479957600080fd5b60a08501868111156147aa57600080fd5b859450356147b78161471e565b925060c08501356001600160401b038111156147d257600080fd5b6147de8782880161473b565b95989497509550505050565b6001600160a01b038116811461087257600080fd5b60006020828403121561481157600080fd5b81356123e4816147ea565b60006020828403121561482e57600080fd5b5035919050565b600080600080600060a0868803121561484d57600080fd5b8535614858816147ea565b94506020860135614868816147ea565b93506040860135614878816147ea565b92506060860135614888816147ea565b91506080860135614898816147ea565b809150509295509295909350565b600061010082840312156148b957600080fd5b50919050565b6000604082840312156148b957600080fd5b600080600060a084860312156148e657600080fd5b83356001600160401b038111156148fc57600080fd5b614908868287016148a6565b93505061491885602086016148bf565b915061492785606086016148bf565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561496857614968614930565b60405290565b60405161010081016001600160401b038111828210171561496857614968614930565b604051601f8201601f191681016001600160401b03811182821017156149b9576149b9614930565b604052919050565b6000604082840312156149d357600080fd5b6149db614946565b9050813581526020820135602082015292915050565b600082601f830112614a0257600080fd5b604051604081018181106001600160401b0382111715614a2457614a24614930565b8060405250806040840185811115614a3b57600080fd5b845b81811015614175578035835260209283019201614a3d565b600060808284031215614a6757600080fd5b614a6f614946565b9050614a7b83836149f1565b8152614a8a83604084016149f1565b602082015292915050565b6000806000806101208587031215614aac57600080fd5b84359350614abd86602087016149c1565b9250614acc8660608701614a55565b9150614adb8660e087016149c1565b905092959194509250565b600060208284031215614af857600080fd5b81356123e48161471e565b60006001600160401b03821115614b1c57614b1c614930565b5060051b60200190565b60008060408385031215614b3957600080fd5b8235614b44816147ea565b91506020838101356001600160401b03811115614b6057600080fd5b8401601f81018613614b7157600080fd5b8035614b84614b7f82614b03565b614991565b81815260059190911b82018301908381019088831115614ba357600080fd5b928401925b82841015614bca578335614bbb816147ea565b82529284019290840190614ba8565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614c0957815187529582019590820190600101614bed565b509495945050505050565b6020815260006123e46020830184614bd9565b600080600060608486031215614c3c57600080fd5b8335614c47816147ea565b92506020848101356001600160401b0380821115614c6457600080fd5b818701915087601f830112614c7857600080fd5b813581811115614c8a57614c8a614930565b614c9c601f8201601f19168501614991565b91508082528884828501011115614cb257600080fd5b808484018584013760008482840101525080945050505061492760408501614730565b600081518084526020808501808196508360051b810191508286016000805b86811015614d6b578385038a52825180518087529087019087870190845b81811015614d5657835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614d12565b50509a87019a95505091850191600101614cf4565b509298975050505050505050565b6020815260006123e46020830184614cd5565b801515811461087257600080fd5b600060208284031215614dac57600080fd5b81356123e481614d8c565b600082601f830112614dc857600080fd5b81356020614dd8614b7f83614b03565b82815260069290921b84018101918181019086841115614df757600080fd5b8286015b84811015614e1b57614e0d88826149c1565b835291830191604001614dfb565b509695505050505050565b60008060008060c08587031215614e3c57600080fd5b84356001600160401b0380821115614e5357600080fd5b614e5f888389016148a6565b9550614e6e88602089016148bf565b9450614e7d88606089016148bf565b935060a0870135915080821115614e9357600080fd5b50614ea087828801614db7565b91505092959194509250565b600082601f830112614ebd57600080fd5b81356020614ecd614b7f83614b03565b82815260059290921b84018101918181019086841115614eec57600080fd5b8286015b84811015614e1b5780358352918301918301614ef0565b60008060408385031215614f1a57600080fd5b8235614f25816147ea565b915060208301356001600160401b03811115614f4057600080fd5b614f4c85828601614eac565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614f975783516001600160a01b031683529284019291840191600101614f72565b50909695505050505050565b60008083601f840112614fb557600080fd5b5081356001600160401b03811115614fcc57600080fd5b6020830191508360208260051b850101111561477c57600080fd5b6000806000806000806080878903121561500057600080fd5b863561500b816147ea565b9550602087013561501b8161471e565b945060408701356001600160401b038082111561503757600080fd5b6150438a838b0161473b565b9096509450606089013591508082111561505c57600080fd5b5061506989828a01614fa3565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614c0957815163ffffffff168752958201959082019060010161508f565b6000602080835283516080828501526150cd60a085018261507b565b905081850151601f19808684030160408701526150ea838361507b565b92506040870151915080868403016060870152615107838361507b565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561515e578487830301845261514c82875161507b565b95880195938801939150600101615132565b509998505050505050505050565b60008060008060006060868803121561518457600080fd5b853561518f8161471e565b945060208601356001600160401b03808211156151ab57600080fd5b6151b789838a01614fa3565b909650945060408801359150808211156151d057600080fd5b506151dd8882890161473b565b969995985093965092949392505050565b60ff8116811461087257600080fd5b60006020828403121561520f57600080fd5b81356123e4816151ee565b60008060006060848603121561522f57600080fd5b833561523a816147ea565b925060208401356001600160401b0381111561525557600080fd5b61526186828701614eac565b92505060408401356152728161471e565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614f9757835183529284019291840191600101615299565b600080604083850312156152c857600080fd5b82356152d38161471e565b946020939093013593505050565b600082601f8301126152f257600080fd5b81356020615302614b7f83614b03565b82815260059290921b8401810191818101908684111561532157600080fd5b8286015b84811015614e1b5780356153388161471e565b8352918301918301615325565b600082601f83011261535657600080fd5b81356020615366614b7f83614b03565b82815260059290921b8401810191818101908684111561538557600080fd5b8286015b84811015614e1b5780356001600160401b038111156153a85760008081fd5b6153b68986838b01016152e1565b845250918301918301615389565b600061018082840312156153d757600080fd5b6153df61496e565b905081356001600160401b03808211156153f857600080fd5b615404858386016152e1565b8352602084013591508082111561541a57600080fd5b61542685838601614db7565b6020840152604084013591508082111561543f57600080fd5b61544b85838601614db7565b604084015261545d8560608601614a55565b606084015261546f8560e086016149c1565b608084015261012084013591508082111561548957600080fd5b615495858386016152e1565b60a08401526101408401359150808211156154af57600080fd5b6154bb858386016152e1565b60c08401526101608401359150808211156154d557600080fd5b506154e284828501615345565b60e08301525092915050565b60008060008060006080868803121561550657600080fd5b8535945060208601356001600160401b038082111561552457600080fd5b61553089838a0161473b565b9096509450604088013591506155458261471e565b9092506060870135908082111561555b57600080fd5b50615568888289016153c4565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614c095781516001600160601b031687529582019590820190600101615589565b60408152600083516040808401526155c96080840182615575565b90506020850151603f198483030160608501526155e68282615575565b925050508260208301529392505050565b60008060006080848603121561560c57600080fd5b83356001600160401b038082111561562357600080fd5b61562f878388016148a6565b945061563e87602088016148bf565b9350606086013591508082111561565457600080fd5b50615661868287016153c4565b9150509250925092565b60008060006060848603121561568057600080fd5b833561568b816147ea565b92506020840135915060408401356152728161471e565b8281526040602082015260006156bb6040830184614cd5565b949350505050565b6020808252825160009190828483015b60058210156156f25782518152918301916001919091019083016156d3565b50505063ffffffff818501511660c084015260408401516101008060e086015281518061012087015260005b8181101561573b578381018501518782016101400152840161571e565b8181111561574e57600061014083890101525b50606087015163ffffffff8116878401529350601f01601f1916949094016101400195945050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168185168083038211156157ae576157ae615779565b01949350505050565b6000602082840312156157c957600080fd5b81516123e4816147ea565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561583057600080fd5b81516123e481614d8c565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b803561588e8161471e565b63ffffffff168252602090810135910152565b608081016158af8285615883565b82356158ba8161471e565b63ffffffff16604083015260209290920135606090910152919050565b63ffffffff8416815260e0810160a084602084013760c0919091019190915292915050565b634e487b7160e01b600052603260045260246000fd5b60008261592f57634e487b7160e01b600052601260045260246000fd5b500690565b60006020828403121561594657600080fd5b5051919050565b600060001982141561596157615961615779565b5060010190565b6000602080838503121561597b57600080fd5b82516001600160401b0381111561599157600080fd5b8301601f810185136159a257600080fd5b80516159b0614b7f82614b03565b81815260059190911b820183019083810190878311156159cf57600080fd5b928401925b828410156159ed578351825292840192908401906159d4565b979650505050505050565b600060208284031215615a0a57600080fd5b81516001600160601b03811681146123e457600080fd5b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615a5c57815185529382019390820190600101615a40565b5092979650505050505050565b81835260006001600160fb1b03831115615a8257600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff84168152604060208201526000615abf604083018486615a69565b95945050505050565b60006020808385031215615adb57600080fd5b82516001600160401b03811115615af157600080fd5b8301601f81018513615b0257600080fd5b8051615b10614b7f82614b03565b81815260059190911b82018301908381019087831115615b2f57600080fd5b928401925b828410156159ed578351615b478161471e565b82529284019290840190615b34565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615abf604083018486615b56565b600060208284031215615bb157600080fd5b81516001600160c01b03811681146123e457600080fd5b600060208284031215615bda57600080fd5b81516123e48161471e565b600060ff821660ff811415615bfc57615bfc615779565b60010192915050565b604081526000615c19604083018587615b56565b905063ffffffff83166020830152949350505050565b63ffffffff86168152606060208201526000615c4f606083018688615a69565b8281036040840152615c62818587615b56565b98975050505050505050565b63ffffffff831681526040602082015260006156bb6040830184614bd9565b600060208284031215615c9f57600080fd5b81516123e4816151ee565b600082821015615cbc57615cbc615779565b500390565b60008219821115615cd457615cd4615779565b500190565b600060208284031215615ceb57600080fd5b815167ffffffffffffffff19811681146123e457600080fd5b60006001600160601b0383811690831681811015615d2457615d24615779565b039392505050565b6000808335601e19843603018112615d4357600080fd5b8301803591506001600160401b03821115615d5d57600080fd5b60200191503681900382131561477c57600080fd5b6020815260a0826020830137600060c08201600080825260a0850135615d978161471e565b63ffffffff1690915260c08401359036859003601e19018212615db8578081fd5b9084019081356001600160401b03811115615dd1578182fd5b803603861315615ddf578182fd5b61010091508160e0860152615dfc61012086018260208601615b56565b925050615e0b60e08601614730565b63ffffffff811685830152614337565b604081016123e78284615883565b60006001600160601b0380831681851681830481118215151615615e4f57615e4f615779565b02949350505050565b6000816000190483118215151615615e7257615e72615779565b500290565b60808101615e858285615883565b63ffffffff8351166040830152602083015160608301529392505050565b600061ffff80831681811415615ebb57615ebb615779565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122030fe8bbfb10863b393afb4d1dbda7dcd468ef978225ba381ca7287ef3ad22d7164736f6c634300080c0033",
    sourceMap:
      "846:12230:80:-:0;;;2367:222;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;2497:20;1701::38;-1:-1:-1;;;;;1679:42:38;;;-1:-1:-1;;;;;1679:42:38;;;;;1747:20;-1:-1:-1;;;;;1747:34:38;;:36;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;1731:52:38;;;-1:-1:-1;;;;;1731:52:38;;;;;1810:20;-1:-1:-1;;;;;1810:35:38;;:37;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;1793:54:38;;;-1:-1:-1;;;;;1793:54:38;;;;;1870:13;;-1:-1:-1;;;;;1870:24:38;;:26;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;1857:39:38;;;-1:-1:-1;1915:20:38;:27;;-1:-1:-1;;1915:27:38;1938:4;1915:27;;;2529:53:80::1;;1915:27:38::0;2529:53:80;-1:-1:-1;846:12230:80;;14:153:82;-1:-1:-1;;;;;111:31:82;;101:42;;91:70;;157:1;154;147:12;91:70;14:153;:::o;172:468::-;280:6;288;341:2;329:9;320:7;316:23;312:32;309:52;;;357:1;354;347:12;309:52;389:9;383:16;408:53;455:5;408:53;:::i;:::-;530:2;515:18;;509:25;480:5;;-1:-1:-1;578:10:82;565:24;;553:37;;543:65;;604:1;601;594:12;543:65;627:7;617:17;;;172:468;;;;;:::o;645:297::-;739:6;792:2;780:9;771:7;767:23;763:32;760:52;;;808:1;805;798:12;760:52;840:9;834:16;859:53;906:5;859:53;:::i;:::-;931:5;645:297;-1:-1:-1;;;645:297:82:o;1250:300::-;846:12230:80;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;",
    linkReferences: {},
  },
  deployedBytecode: {
    object:
      "0x608060405234801561001057600080fd5b50600436106102485760003560e01c80635c975abb1161013b5780638b00ce7c116100b8578063df5cf7231161007c578063df5cf723146105f5578063f2fde38b1461061c578063f5c9899d1461062f578063f63c5bab14610655578063fabc1cbc1461065d57600080fd5b80638b00ce7c146105935780638da5cb5b146105a3578063b98d0908146105b4578063bdeea6cc146105c1578063cefdc1d4146105d457600080fd5b80636efb4636116100ff5780636efb463614610536578063715018a61461055757806372d18e8d1461055f5780637afa1eed1461056d578063886f11951461058057600080fd5b80635c975abb146104a65780635df45946146104ae57806367d6be44146104d557806368304835146104e85780636d14a9871461050f57600080fd5b806331b36bd9116101c95780634f739f741161018d5780634f739f741461041857806358a7cd2614610438578063595c6a671461044b5780635ac86ab7146104535780635c1556621461048657600080fd5b806331b36bd9146103925780633563b0d1146103b2578063416c7e5e146103d2578063480bab6b146103e55780634d2b57fe146103f857600080fd5b8063171f1d5b11610210578063171f1d5b146102ae5780631ad43189146102dd578063245a7bfc146103195780632cb223d5146103445780632d89f6fc1461037257600080fd5b80630627721e1461024d57806310d67a2f14610262578063136439dd146102755780631459457a14610288578063162d8a3f1461029b575b600080fd5b61026061025b366004614783565b610670565b005b6102606102703660046147ff565b6107b9565b61026061028336600461481c565b610875565b610260610296366004614835565b6109b4565b6102606102a93660046148d1565b610b13565b6102c16102bc366004614a95565b610c5f565b6040805192151583529015156020830152015b60405180910390f35b6103047f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102d4565b60cc5461032c906001600160a01b031681565b6040516001600160a01b0390911681526020016102d4565b610364610352366004614ae6565b60cb6020526000908152604090205481565b6040519081526020016102d4565b610364610380366004614ae6565b60ca6020526000908152604090205481565b6103a56103a0366004614b26565b610de9565b6040516102d49190614c14565b6103c56103c0366004614c27565b610f05565b6040516102d49190614d79565b6102606103e0366004614d9a565b61139b565b6102606103f3366004614e26565b611510565b61040b610406366004614f07565b6117fc565b6040516102d49190614f56565b61042b610426366004614fe7565b611911565b6040516102d491906150b1565b61026061044636600461516c565b612037565b6102606120db565b6104766104613660046151fd565b606654600160ff9092169190911b9081161490565b60405190151581526020016102d4565b61049961049436600461521a565b6121a2565b6040516102d4919061527d565b606654610364565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b6103646104e33660046152b5565b61236a565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b6105496105443660046154ee565b6123ed565b6040516102d49291906155ae565b6102606132e5565b60c95463ffffffff16610304565b60cd5461032c906001600160a01b031681565b60655461032c906001600160a01b031681565b60c9546103049063ffffffff1681565b6033546001600160a01b031661032c565b6097546104769060ff1681565b6102606105cf3660046155f7565b6132f9565b6105e76105e236600461566b565b6135b1565b6040516102d49291906156a2565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b61026061062a3660046147ff565b613743565b7f0000000000000000000000000000000000000000000000000000000000000000610304565b610304606481565b61026061066b36600461481c565b6137b9565b60cd546001600160a01b0316331461068757600080fd5b61068f6145f8565b6040805160a081810190925290869060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f8501839004830281018301909152838152908490849081908401838280828437600092019190915250505050604080830191909152516107199082906020016156c3565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f54977349061077c9084906156c3565b60405180910390a260c9546107989063ffffffff16600161578f565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561080c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083091906157b7565b6001600160a01b0316336001600160a01b0316146108695760405162461bcd60e51b8152600401610860906157d4565b60405180910390fd5b61087281613915565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156108bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e1919061581e565b6108fd5760405162461bcd60e51b81526004016108609061583b565b606654818116146109765760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610860565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156109d45750600054600160ff909116105b806109ee5750303b1580156109ee575060005460ff166001145b610a515760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610860565b6000805460ff191660011790558015610a74576000805461ff0019166101001790555b610a7f866000613a0c565b610a8885613af6565b60cc80546001600160a01b038087166001600160a01b03199283161790925560cd805486841690831617905560ce8054928516929091169190911790558015610b0b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6000610b226020840184614ae6565b63ffffffff8116600090815260cb6020526040902054909150610b4457600080fd5b8282604051602001610b579291906158a1565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb90935291205414610b8d57600080fd5b6064610b9c6020840184614ae6565b610ba6919061578f565b63ffffffff164363ffffffff161115610bbe57600080fd5b60ce54604051631150ef3360e21b81526001600160a01b0390911690634543bccc90610bf690849088906020890135906004016158d7565b600060405180830381600087803b158015610c1057600080fd5b505af1158015610c24573d6000803e3d6000fd5b505060405163ffffffff841692507f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a9150600090a250505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610ca757610ca76158fc565b60200201518951600160200201518a60200151600060028110610ccc57610ccc6158fc565b60200201518b60200151600160028110610ce857610ce86158fc565b602090810291909101518c518d830151604051610d459a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d689190615912565b9050610ddb610d81610d7a8884613b48565b8690613bdf565b610d89613c73565b610dd1610dc285610dbc604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613b48565b610dcb8c613d33565b90613bdf565b886201d4c0613dc3565b909890975095505050505050565b606081516001600160401b03811115610e0457610e04614930565b604051908082528060200260200182016040528015610e2d578160200160208202803683370190505b50905060005b8251811015610efe57836001600160a01b03166313542a4e848381518110610e5d57610e5d6158fc565b60200260200101516040518263ffffffff1660e01b8152600401610e9091906001600160a01b0391909116815260200190565b602060405180830381865afa158015610ead573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed19190615934565b828281518110610ee357610ee36158fc565b6020908102919091010152610ef78161594d565b9050610e33565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6b91906157b7565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd191906157b7565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611013573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103791906157b7565b9050600086516001600160401b0381111561105457611054614930565b60405190808252806020026020018201604052801561108757816020015b60608152602001906001900390816110725790505b50905060005b875181101561138f5760008882815181106110aa576110aa6158fc565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa15801561110b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111339190810190615968565b905080516001600160401b0381111561114e5761114e614930565b60405190808252806020026020018201604052801561119957816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161116c5790505b508484815181106111ac576111ac6158fc565b602002602001018190525060005b8151811015611379576040518060600160405280876001600160a01b03166347b314e88585815181106111ef576111ef6158fc565b60200260200101516040518263ffffffff1660e01b815260040161121591815260200190565b602060405180830381865afa158015611232573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125691906157b7565b6001600160a01b03168152602001838381518110611276576112766158fc565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106112a4576112a46158fc565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611300573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132491906159f8565b6001600160601b0316815250858581518110611342576113426158fc565b6020026020010151828151811061135b5761135b6158fc565b602002602001018190525080806113719061594d565b9150506111ba565b50505080806113879061594d565b91505061108d565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141d91906157b7565b6001600160a01b0316336001600160a01b0316146114c95760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610860565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b600061151f6020850185614ae6565b9050600082516001600160401b0381111561153c5761153c614930565b604051908082528060200260200182016040528015611565578160200160208202803683370190505b50905060005b83518110156115d7576115a8848281518110611589576115896158fc565b6020026020010151805160009081526020918201519091526040902090565b8282815181106115ba576115ba6158fc565b6020908102919091010152806115cf8161594d565b91505061156b565b5060006115ea60c0880160a08901614ae6565b826040516020016115fc929190615a21565b6040516020818303038152906040528051906020012090508460200135811461162457600080fd5b600084516001600160401b0381111561163f5761163f614930565b604051908082528060200260200182016040528015611668578160200160208202803683370190505b50905060005b855181101561175b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106116b8576116b86158fc565b60200260200101516040518263ffffffff1660e01b81526004016116de91815260200190565b602060405180830381865afa1580156116fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171f91906157b7565b828281518110611731576117316158fc565b6001600160a01b0390921660209283029190910190910152806117538161594d565b91505061166e565b5060ce5460405163959b127960e01b815263ffffffff861660048201526001600160a01b039091169063959b127990602401600060405180830381600087803b1580156117a757600080fd5b505af11580156117bb573d6000803e3d6000fd5b505060405133925063ffffffff871691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050565b606081516001600160401b0381111561181757611817614930565b604051908082528060200260200182016040528015611840578160200160208202803683370190505b50905060005b8251811015610efe57836001600160a01b031663296bb064848381518110611870576118706158fc565b60200260200101516040518263ffffffff1660e01b815260040161189691815260200190565b602060405180830381865afa1580156118b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118d791906157b7565b8282815181106118e9576118e96158fc565b6001600160a01b039092166020928302919091019091015261190a8161594d565b9050611846565b61193c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561197c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119a091906157b7565b90506119cd6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906119fd908b9089908990600401615a9f565b600060405180830381865afa158015611a1a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a429190810190615ac8565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611a74908b908b908b90600401615b7f565b600060405180830381865afa158015611a91573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ab99190810190615ac8565b6040820152856001600160401b03811115611ad657611ad6614930565b604051908082528060200260200182016040528015611b0957816020015b6060815260200190600190039081611af45790505b50606082015260005b60ff8116871115611f48576000856001600160401b03811115611b3757611b37614930565b604051908082528060200260200182016040528015611b60578160200160208202803683370190505b5083606001518360ff1681518110611b7a57611b7a6158fc565b602002602001018190525060005b86811015611e485760008c6001600160a01b03166304ec63518a8a85818110611bb357611bb36158fc565b905060200201358e88600001518681518110611bd157611bd16158fc565b60200260200101516040518463ffffffff1660e01b8152600401611c0e9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611c2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4f9190615b9f565b90506001600160c01b038116611cf35760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610860565b8a8a8560ff16818110611d0857611d086158fc565b6001600160c01b03841692013560f81c9190911c600190811614159050611e3557856001600160a01b031663dd9846b98a8a85818110611d4a57611d4a6158fc565b905060200201358d8d8860ff16818110611d6657611d666158fc565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611dbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de09190615bc8565b85606001518560ff1681518110611df957611df96158fc565b60200260200101518481518110611e1257611e126158fc565b63ffffffff9092166020928302919091019091015282611e318161594d565b9350505b5080611e408161594d565b915050611b88565b506000816001600160401b03811115611e6357611e63614930565b604051908082528060200260200182016040528015611e8c578160200160208202803683370190505b50905060005b82811015611f0d5784606001518460ff1681518110611eb357611eb36158fc565b60200260200101518181518110611ecc57611ecc6158fc565b6020026020010151828281518110611ee657611ee66158fc565b63ffffffff9092166020928302919091019091015280611f058161594d565b915050611e92565b508084606001518460ff1681518110611f2857611f286158fc565b602002602001018190525050508080611f4090615be5565b915050611b12565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fad91906157b7565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611fe0908b908b908e90600401615c05565b600060405180830381865afa158015611ffd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120259190810190615ac8565b60208301525098975050505050505050565b60ce54604051632c53e69360e11b81526001600160a01b03909116906358a7cd269061206f9088908890889088908890600401615c2f565b600060405180830381600087803b15801561208957600080fd5b505af115801561209d573d6000803e3d6000fd5b505060405133925063ffffffff881691507ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015612123573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612147919061581e565b6121635760405162461bcd60e51b81526004016108609061583b565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016121d4929190615c6e565b600060405180830381865afa1580156121f1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122199190810190615ac8565b9050600084516001600160401b0381111561223657612236614930565b60405190808252806020026020018201604052801561225f578160200160208202803683370190505b50905060005b855181101561236057866001600160a01b03166304ec635187838151811061228f5761228f6158fc565b6020026020010151878685815181106122aa576122aa6158fc565b60200260200101516040518463ffffffff1660e01b81526004016122e79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612304573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123289190615b9f565b6001600160c01b0316828281518110612343576123436158fc565b6020908102919091010152806123588161594d565b915050612265565b5095945050505050565b60ce546040516319f5af9160e21b815263ffffffff84166004820152602481018390526000916001600160a01b0316906367d6be4490604401602060405180830381865afa1580156123c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e49190615934565b90505b92915050565b60408051808201909152606080825260208201526000846124645760405162461bcd60e51b81526020600482015260376024820152600080516020615ee683398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610860565b6040830151518514801561247c575060a08301515185145b801561248c575060c08301515185145b801561249c575060e08301515185145b6125065760405162461bcd60e51b81526020600482015260416024820152600080516020615ee683398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610860565b8251516020840151511461257e5760405162461bcd60e51b815260206004820152604460248201819052600080516020615ee6833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610860565b4363ffffffff168463ffffffff16106125ed5760405162461bcd60e51b815260206004820152603c6024820152600080516020615ee683398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610860565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561262e5761262e614930565b604051908082528060200260200182016040528015612657578160200160208202803683370190505b506020820152866001600160401b0381111561267557612675614930565b60405190808252806020026020018201604052801561269e578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156126d2576126d2614930565b6040519080825280602002602001820160405280156126fb578160200160208202803683370190505b5081526020860151516001600160401b0381111561271b5761271b614930565b604051908082528060200260200182016040528015612744578160200160208202803683370190505b50816020018190525060006128168a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156127ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128119190615c8d565b613fe7565b905060005b876020015151811015612a925761284188602001518281518110611589576115896158fc565b83602001518281518110612857576128576158fc565b60209081029190910101528015612917576020830151612878600183615caa565b81518110612888576128886158fc565b602002602001015160001c836020015182815181106128a9576128a96158fc565b602002602001015160001c11612917576040805162461bcd60e51b8152602060048201526024810191909152600080516020615ee683398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610860565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061295c5761295c6158fc565b60200260200101518b8b60000151858151811061297b5761297b6158fc565b60200260200101516040518463ffffffff1660e01b81526004016129b89392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156129d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129f99190615b9f565b6001600160c01b031683600001518281518110612a1857612a186158fc565b602002602001018181525050612a7e610d7a612a528486600001518581518110612a4457612a446158fc565b602002602001015116614071565b8a602001518481518110612a6857612a686158fc565b602002602001015161409c90919063ffffffff16565b945080612a8a8161594d565b91505061281b565b5050612a9d83614180565b60975490935060ff16600081612ab4576000612b36565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b369190615934565b905060005b8a8110156131b4578215612c96578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612b9257612b926158fc565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612bd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf69190615934565b612c009190615cc1565b11612c965760405162461bcd60e51b81526020600482015260666024820152600080516020615ee683398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610860565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612cd757612cd76158fc565b9050013560f81c60f81b60f81c8c8c60a001518581518110612cfb57612cfb6158fc565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612d57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d7b9190615cd9565b6001600160401b031916612d9e8a604001518381518110611589576115896158fc565b67ffffffffffffffff191614612e3a5760405162461bcd60e51b81526020600482015260616024820152600080516020615ee683398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610860565b612e6a89604001518281518110612e5357612e536158fc565b602002602001015187613bdf90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612ead57612ead6158fc565b9050013560f81c60f81b60f81c8c8c60c001518581518110612ed157612ed16158fc565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612f2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f5191906159f8565b85602001518281518110612f6757612f676158fc565b6001600160601b03909216602092830291909101820152850151805182908110612f9357612f936158fc565b602002602001015185600001518281518110612fb157612fb16158fc565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561319f5761302986600001518281518110612ffb57612ffb6158fc565b60200260200101518f8f86818110613015576130156158fc565b600192013560f81c9290921c811614919050565b1561318d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061306f5761306f6158fc565b9050013560f81c60f81b60f81c8e89602001518581518110613093576130936158fc565b60200260200101518f60e0015188815181106130b1576130b16158fc565b602002602001015187815181106130ca576130ca6158fc565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561312e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061315291906159f8565b8751805185908110613166576131666158fc565b6020026020010181815161317a9190615d04565b6001600160601b03169052506001909101905b806131978161594d565b915050612fd5565b505080806131ac9061594d565b915050612b3b565b5050506000806131ce8c868a606001518b60800151610c5f565b915091508161323f5760405162461bcd60e51b81526020600482015260436024820152600080516020615ee683398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610860565b806132a05760405162461bcd60e51b81526020600482015260396024820152600080516020615ee683398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610860565b505060008782602001516040516020016132bb929190615a21565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6132ed61421b565b6132f76000613af6565b565b60cc546001600160a01b0316331461331057600080fd5b600061332260c0850160a08601614ae6565b905036600061333460c0870187615d2c565b9092509050600061334c610100880160e08901614ae6565b905060ca600061335f6020890189614ae6565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161338b9190615d72565b60405160208183030381529060405280519060200120146133ab57600080fd5b600060cb816133bd60208a018a614ae6565b63ffffffff1663ffffffff16815260200190815260200160002054146133e257600080fd5b61340c7f00000000000000000000000000000000000000000000000000000000000000008561578f565b63ffffffff164363ffffffff16111561342457600080fd5b6000866040516020016134379190615e1b565b60405160208183030381529060405280519060200120905060008061345f8387878a8c6123ed565b9150915060005b858110156134f3578460ff1683602001518281518110613488576134886158fc565b602002602001015161349a9190615e29565b6001600160601b03166064846000015183815181106134bb576134bb6158fc565b60200260200101516001600160601b03166134d69190615e58565b10156134e157600080fd5b806134eb8161594d565b915050613466565b5060408051808201825263ffffffff43168152602080820184905291519091613520918c91849101615e77565b6040516020818303038152906040528051906020012060cb60008c600001602081019061354d9190614ae6565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a8260405161359c929190615e77565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106135ec576135ec6158fc565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906136289088908690600401615c6e565b600060405180830381865afa158015613645573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261366d9190810190615ac8565b60008151811061367f5761367f6158fc565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156136eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061370f9190615b9f565b6001600160c01b03169050600061372582614275565b9050816137338a838a610f05565b9550955050505050935093915050565b61374b61421b565b6001600160a01b0381166137b05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610860565b61087281613af6565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561380c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061383091906157b7565b6001600160a01b0316336001600160a01b0316146138605760405162461bcd60e51b8152600401610860906157d4565b6066541981196066541916146138de5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610860565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016109a9565b6001600160a01b0381166139a35760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610860565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b0316158015613a2d57506001600160a01b03821615155b613aaf5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610860565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613af282613915565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613b64614626565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613b9757613b99565bfe5b5080613bd75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610860565b505092915050565b6040805180820190915260008082526020820152613bfb614644565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613b97575080613bd75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610860565b613c7b614662565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613d63600080516020615ec683398151915286615912565b90505b613d6f81614341565b9093509150600080516020615ec6833981519152828309831415613da9576040805180820190915290815260208101919091529392505050565b600080516020615ec6833981519152600182089050613d66565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613df5614687565b60005b6002811015613fba576000613e0e826006615e58565b9050848260028110613e2257613e226158fc565b60200201515183613e34836000615cc1565b600c8110613e4457613e446158fc565b6020020152848260028110613e5b57613e5b6158fc565b60200201516020015183826001613e729190615cc1565b600c8110613e8257613e826158fc565b6020020152838260028110613e9957613e996158fc565b6020020151515183613eac836002615cc1565b600c8110613ebc57613ebc6158fc565b6020020152838260028110613ed357613ed36158fc565b6020020151516001602002015183613eec836003615cc1565b600c8110613efc57613efc6158fc565b6020020152838260028110613f1357613f136158fc565b602002015160200151600060028110613f2e57613f2e6158fc565b602002015183613f3f836004615cc1565b600c8110613f4f57613f4f6158fc565b6020020152838260028110613f6657613f666158fc565b602002015160200151600160028110613f8157613f816158fc565b602002015183613f92836005615cc1565b600c8110613fa257613fa26158fc565b60200201525080613fb28161594d565b915050613df8565b50613fc36146a6565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613ff3846143c3565b9050808360ff166001901b116123e45760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610860565b6000805b82156123e757614086600184615caa565b909216918061409481615ea3565b915050614075565b60408051808201909152600080825260208201526102008261ffff16106140f85760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610860565b8161ffff166001141561410c5750816123e7565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061417557600161ffff871660ff83161c81161415614158576141558484613bdf565b93505b6141628384613bdf565b92506201fffe600192831b169101614128565b509195945050505050565b604080518082019091526000808252602082015281511580156141a557506020820151155b156141c3575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615ec683398151915284602001516141f69190615912565b61420e90600080516020615ec6833981519152615caa565b905292915050565b919050565b6033546001600160a01b031633146132f75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610860565b606060008061428384614071565b61ffff166001600160401b0381111561429e5761429e614930565b6040519080825280601f01601f1916602001820160405280156142c8576020820181803683370190505b5090506000805b8251821080156142e0575061010081105b15614337576001811b935085841615614327578060f81b838381518110614309576143096158fc565b60200101906001600160f81b031916908160001a9053508160010191505b6143308161594d565b90506142cf565b5090949350505050565b60008080600080516020615ec68339815191526003600080516020615ec683398151915286600080516020615ec68339815191528889090908905060006143b7827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615ec6833981519152614550565b91959194509092505050565b60006101008251111561444c5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610860565b815161445a57506000919050565b60008083600081518110614470576144706158fc565b0160200151600160f89190911c81901b92505b84518110156145475784818151811061449e5761449e6158fc565b0160200151600160f89190911c1b91508282116145335760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610860565b918117916145408161594d565b9050614483565b50909392505050565b60008061455b6146a6565b6145636146c4565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613b975750826145ed5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610860565b505195945050505050565b604051806080016040528061460b6146e2565b81526000602082018190526060604083018190529091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280614675614700565b8152602001614682614700565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461087257600080fd5b80356142168161471e565b60008083601f84011261474d57600080fd5b5081356001600160401b0381111561476457600080fd5b60208301915083602082850101111561477c57600080fd5b9250929050565b60008060008060e0858703121561479957600080fd5b60a08501868111156147aa57600080fd5b859450356147b78161471e565b925060c08501356001600160401b038111156147d257600080fd5b6147de8782880161473b565b95989497509550505050565b6001600160a01b038116811461087257600080fd5b60006020828403121561481157600080fd5b81356123e4816147ea565b60006020828403121561482e57600080fd5b5035919050565b600080600080600060a0868803121561484d57600080fd5b8535614858816147ea565b94506020860135614868816147ea565b93506040860135614878816147ea565b92506060860135614888816147ea565b91506080860135614898816147ea565b809150509295509295909350565b600061010082840312156148b957600080fd5b50919050565b6000604082840312156148b957600080fd5b600080600060a084860312156148e657600080fd5b83356001600160401b038111156148fc57600080fd5b614908868287016148a6565b93505061491885602086016148bf565b915061492785606086016148bf565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561496857614968614930565b60405290565b60405161010081016001600160401b038111828210171561496857614968614930565b604051601f8201601f191681016001600160401b03811182821017156149b9576149b9614930565b604052919050565b6000604082840312156149d357600080fd5b6149db614946565b9050813581526020820135602082015292915050565b600082601f830112614a0257600080fd5b604051604081018181106001600160401b0382111715614a2457614a24614930565b8060405250806040840185811115614a3b57600080fd5b845b81811015614175578035835260209283019201614a3d565b600060808284031215614a6757600080fd5b614a6f614946565b9050614a7b83836149f1565b8152614a8a83604084016149f1565b602082015292915050565b6000806000806101208587031215614aac57600080fd5b84359350614abd86602087016149c1565b9250614acc8660608701614a55565b9150614adb8660e087016149c1565b905092959194509250565b600060208284031215614af857600080fd5b81356123e48161471e565b60006001600160401b03821115614b1c57614b1c614930565b5060051b60200190565b60008060408385031215614b3957600080fd5b8235614b44816147ea565b91506020838101356001600160401b03811115614b6057600080fd5b8401601f81018613614b7157600080fd5b8035614b84614b7f82614b03565b614991565b81815260059190911b82018301908381019088831115614ba357600080fd5b928401925b82841015614bca578335614bbb816147ea565b82529284019290840190614ba8565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614c0957815187529582019590820190600101614bed565b509495945050505050565b6020815260006123e46020830184614bd9565b600080600060608486031215614c3c57600080fd5b8335614c47816147ea565b92506020848101356001600160401b0380821115614c6457600080fd5b818701915087601f830112614c7857600080fd5b813581811115614c8a57614c8a614930565b614c9c601f8201601f19168501614991565b91508082528884828501011115614cb257600080fd5b808484018584013760008482840101525080945050505061492760408501614730565b600081518084526020808501808196508360051b810191508286016000805b86811015614d6b578385038a52825180518087529087019087870190845b81811015614d5657835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614d12565b50509a87019a95505091850191600101614cf4565b509298975050505050505050565b6020815260006123e46020830184614cd5565b801515811461087257600080fd5b600060208284031215614dac57600080fd5b81356123e481614d8c565b600082601f830112614dc857600080fd5b81356020614dd8614b7f83614b03565b82815260069290921b84018101918181019086841115614df757600080fd5b8286015b84811015614e1b57614e0d88826149c1565b835291830191604001614dfb565b509695505050505050565b60008060008060c08587031215614e3c57600080fd5b84356001600160401b0380821115614e5357600080fd5b614e5f888389016148a6565b9550614e6e88602089016148bf565b9450614e7d88606089016148bf565b935060a0870135915080821115614e9357600080fd5b50614ea087828801614db7565b91505092959194509250565b600082601f830112614ebd57600080fd5b81356020614ecd614b7f83614b03565b82815260059290921b84018101918181019086841115614eec57600080fd5b8286015b84811015614e1b5780358352918301918301614ef0565b60008060408385031215614f1a57600080fd5b8235614f25816147ea565b915060208301356001600160401b03811115614f4057600080fd5b614f4c85828601614eac565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614f975783516001600160a01b031683529284019291840191600101614f72565b50909695505050505050565b60008083601f840112614fb557600080fd5b5081356001600160401b03811115614fcc57600080fd5b6020830191508360208260051b850101111561477c57600080fd5b6000806000806000806080878903121561500057600080fd5b863561500b816147ea565b9550602087013561501b8161471e565b945060408701356001600160401b038082111561503757600080fd5b6150438a838b0161473b565b9096509450606089013591508082111561505c57600080fd5b5061506989828a01614fa3565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614c0957815163ffffffff168752958201959082019060010161508f565b6000602080835283516080828501526150cd60a085018261507b565b905081850151601f19808684030160408701526150ea838361507b565b92506040870151915080868403016060870152615107838361507b565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561515e578487830301845261514c82875161507b565b95880195938801939150600101615132565b509998505050505050505050565b60008060008060006060868803121561518457600080fd5b853561518f8161471e565b945060208601356001600160401b03808211156151ab57600080fd5b6151b789838a01614fa3565b909650945060408801359150808211156151d057600080fd5b506151dd8882890161473b565b969995985093965092949392505050565b60ff8116811461087257600080fd5b60006020828403121561520f57600080fd5b81356123e4816151ee565b60008060006060848603121561522f57600080fd5b833561523a816147ea565b925060208401356001600160401b0381111561525557600080fd5b61526186828701614eac565b92505060408401356152728161471e565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614f9757835183529284019291840191600101615299565b600080604083850312156152c857600080fd5b82356152d38161471e565b946020939093013593505050565b600082601f8301126152f257600080fd5b81356020615302614b7f83614b03565b82815260059290921b8401810191818101908684111561532157600080fd5b8286015b84811015614e1b5780356153388161471e565b8352918301918301615325565b600082601f83011261535657600080fd5b81356020615366614b7f83614b03565b82815260059290921b8401810191818101908684111561538557600080fd5b8286015b84811015614e1b5780356001600160401b038111156153a85760008081fd5b6153b68986838b01016152e1565b845250918301918301615389565b600061018082840312156153d757600080fd5b6153df61496e565b905081356001600160401b03808211156153f857600080fd5b615404858386016152e1565b8352602084013591508082111561541a57600080fd5b61542685838601614db7565b6020840152604084013591508082111561543f57600080fd5b61544b85838601614db7565b604084015261545d8560608601614a55565b606084015261546f8560e086016149c1565b608084015261012084013591508082111561548957600080fd5b615495858386016152e1565b60a08401526101408401359150808211156154af57600080fd5b6154bb858386016152e1565b60c08401526101608401359150808211156154d557600080fd5b506154e284828501615345565b60e08301525092915050565b60008060008060006080868803121561550657600080fd5b8535945060208601356001600160401b038082111561552457600080fd5b61553089838a0161473b565b9096509450604088013591506155458261471e565b9092506060870135908082111561555b57600080fd5b50615568888289016153c4565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614c095781516001600160601b031687529582019590820190600101615589565b60408152600083516040808401526155c96080840182615575565b90506020850151603f198483030160608501526155e68282615575565b925050508260208301529392505050565b60008060006080848603121561560c57600080fd5b83356001600160401b038082111561562357600080fd5b61562f878388016148a6565b945061563e87602088016148bf565b9350606086013591508082111561565457600080fd5b50615661868287016153c4565b9150509250925092565b60008060006060848603121561568057600080fd5b833561568b816147ea565b92506020840135915060408401356152728161471e565b8281526040602082015260006156bb6040830184614cd5565b949350505050565b6020808252825160009190828483015b60058210156156f25782518152918301916001919091019083016156d3565b50505063ffffffff818501511660c084015260408401516101008060e086015281518061012087015260005b8181101561573b578381018501518782016101400152840161571e565b8181111561574e57600061014083890101525b50606087015163ffffffff8116878401529350601f01601f1916949094016101400195945050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168185168083038211156157ae576157ae615779565b01949350505050565b6000602082840312156157c957600080fd5b81516123e4816147ea565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561583057600080fd5b81516123e481614d8c565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b803561588e8161471e565b63ffffffff168252602090810135910152565b608081016158af8285615883565b82356158ba8161471e565b63ffffffff16604083015260209290920135606090910152919050565b63ffffffff8416815260e0810160a084602084013760c0919091019190915292915050565b634e487b7160e01b600052603260045260246000fd5b60008261592f57634e487b7160e01b600052601260045260246000fd5b500690565b60006020828403121561594657600080fd5b5051919050565b600060001982141561596157615961615779565b5060010190565b6000602080838503121561597b57600080fd5b82516001600160401b0381111561599157600080fd5b8301601f810185136159a257600080fd5b80516159b0614b7f82614b03565b81815260059190911b820183019083810190878311156159cf57600080fd5b928401925b828410156159ed578351825292840192908401906159d4565b979650505050505050565b600060208284031215615a0a57600080fd5b81516001600160601b03811681146123e457600080fd5b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615a5c57815185529382019390820190600101615a40565b5092979650505050505050565b81835260006001600160fb1b03831115615a8257600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff84168152604060208201526000615abf604083018486615a69565b95945050505050565b60006020808385031215615adb57600080fd5b82516001600160401b03811115615af157600080fd5b8301601f81018513615b0257600080fd5b8051615b10614b7f82614b03565b81815260059190911b82018301908381019087831115615b2f57600080fd5b928401925b828410156159ed578351615b478161471e565b82529284019290840190615b34565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615abf604083018486615b56565b600060208284031215615bb157600080fd5b81516001600160c01b03811681146123e457600080fd5b600060208284031215615bda57600080fd5b81516123e48161471e565b600060ff821660ff811415615bfc57615bfc615779565b60010192915050565b604081526000615c19604083018587615b56565b905063ffffffff83166020830152949350505050565b63ffffffff86168152606060208201526000615c4f606083018688615a69565b8281036040840152615c62818587615b56565b98975050505050505050565b63ffffffff831681526040602082015260006156bb6040830184614bd9565b600060208284031215615c9f57600080fd5b81516123e4816151ee565b600082821015615cbc57615cbc615779565b500390565b60008219821115615cd457615cd4615779565b500190565b600060208284031215615ceb57600080fd5b815167ffffffffffffffff19811681146123e457600080fd5b60006001600160601b0383811690831681811015615d2457615d24615779565b039392505050565b6000808335601e19843603018112615d4357600080fd5b8301803591506001600160401b03821115615d5d57600080fd5b60200191503681900382131561477c57600080fd5b6020815260a0826020830137600060c08201600080825260a0850135615d978161471e565b63ffffffff1690915260c08401359036859003601e19018212615db8578081fd5b9084019081356001600160401b03811115615dd1578182fd5b803603861315615ddf578182fd5b61010091508160e0860152615dfc61012086018260208601615b56565b925050615e0b60e08601614730565b63ffffffff811685830152614337565b604081016123e78284615883565b60006001600160601b0380831681851681830481118215151615615e4f57615e4f615779565b02949350505050565b6000816000190483118215151615615e7257615e72615779565b500290565b60808101615e858285615883565b63ffffffff8351166040830152602083015160608301529392505050565b600061ffff80831681811415615ebb57615ebb615779565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122030fe8bbfb10863b393afb4d1dbda7dcd468ef978225ba381ca7287ef3ad22d7164736f6c634300080c0033",
    sourceMap:
      "846:12230:80:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;3288:692;;;;;;:::i;:::-;;:::i;:::-;;5826:138:34;;;;;;:::i;:::-;;:::i;3832:392::-;;;;;;:::i;:::-;;:::i;2772:425:80:-;;;;;;:::i;:::-;;:::i;6576:854::-;;;;;;:::i;:::-;;:::i;13991::38:-;;;;;;:::i;:::-;;:::i;:::-;;;;6762:14:82;;6755:22;6737:41;;6821:14;;6814:22;6809:2;6794:18;;6787:50;6710:18;13991:854:38;;;;;;;;1171:50:80;;;;;;;;7121:10:82;7109:23;;;7091:42;;7079:2;7064:18;1171:50:80;6947:192:82;1854:25:80;;;;;-1:-1:-1;;;;;1854:25:80;;;;;;-1:-1:-1;;;;;7308:32:82;;;7290:51;;7278:2;7263:18;1854:25:80;7144:203:82;1797:50:80;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;7748:25:82;;;7736:2;7721:18;1797:50:80;7602:177:82;1654:47:80;;;;;;:::i;:::-;;;;;;;;;;;;;;10650:380:39;;;;;;:::i;:::-;;:::i;:::-;;;;;;;:::i;3037:1255::-;;;;;;:::i;:::-;;:::i;:::-;;;;;;;:::i;2175:168:38:-;;;;;;:::i;:::-;;:::i;7714:5235:80:-;;;;;;:::i;:::-;;:::i;11441:390:39:-;;;;;;:::i;:::-;;:::i;:::-;;;;;;;:::i;5476:3709::-;;;;;;:::i;:::-;;:::i;:::-;;;;;;;:::i;7436:272:80:-;;;;;;:::i;:::-;;:::i;4299:136:34:-;;;:::i;5606:149::-;;;;;;:::i;:::-;5724:7;;5695:1;:10;;;;;;;;5724:14;;;5723:24;;5606:149;;;;21262:14:82;;21255:22;21237:41;;21225:2;21210:18;5606:149:34;21097:187:82;9602:654:39;;;;;;:::i;:::-;;:::i;:::-;;;;;;;:::i;5418:87:34:-;5491:7;;5418:87;;1125:47:38;;;;;2595:171:80;;;;;;:::i;:::-;;:::i;1074:45:38:-;;;;;1011:57;;;;;4469:9040;;;;;;:::i;:::-;;:::i;:::-;;;;;;;;:::i;2071:101:0:-;;;:::i;6480:90:80:-;6550:13;;;;6480:90;;1885:24;;;;;-1:-1:-1;;;;;1885:24:80;;;1825:37:34;;;;;-1:-1:-1;;;;;1825:37:34;;;1397:27:80;;;;;;;;;1441:85:0;1513:6;;-1:-1:-1;;;;;1513:6:0;1441:85;;1363:32:38;;;;;;;;;4041:2433:80;;;;;;:::i;:::-;;:::i;1757:712:39:-;;;;;;:::i;:::-;;:::i;:::-;;;;;;;;:::i;1178:46:38:-;;;;;2321:198:0;;;;;;:::i;:::-;;:::i;12955:119:80:-;13041:26;12955:119;;1227:56;;1280:3;1227:56;;4911:437:34;;;;;;:::i;:::-;;:::i;3288:692:80:-;2333:9;;-1:-1:-1;;;;;2333:9:80;2319:10;:23;2311:32;;;;;;3506:19:::1;;:::i;:::-;3535:23;::::0;;;;;::::1;::::0;;;;3552:6;;3535:23:::1;::::0;;;3552:6;;3535:23;3552:6;3535:23;::::1;;::::0;::::1;::::0;;;;-1:-1:-1;;;3535:23:80;;3568:47:::1;3602:12;3568:47:::0;::::1;:24;::::0;;::::1;:47:::0;;;;3625:61;;::::1;:33;::::0;::::1;:61:::0;3696:37:::1;::::0;;3535:23:::1;3696:37:::0;::::1;::::0;;::::1;::::0;::::1;::::0;;;;;;;;;;;3720:13;;;;;;3696:37;::::1;3720:13:::0;;;;3696:37;::::1;;::::0;::::1;::::0;;;;-1:-1:-1;;;;3696:21:80::1;::::0;;::::1;:37:::0;;;;3857:19;::::1;::::0;3696:7;;3857:19:::1;;;:::i;:::-;;::::0;;-1:-1:-1;;3857:19:80;;::::1;::::0;;;;;;3847:30;;3857:19:::1;3847:30:::0;;::::1;::::0;3830:13:::1;::::0;;::::1;::::0;;::::1;3816:28;::::0;;;:13:::1;:28:::0;;;;;;;:61;3907:13;::::1;::::0;3892:38:::1;::::0;::::1;::::0;3922:7;;3892:38:::1;:::i;:::-;;;;;;;;3956:13;::::0;:17:::1;::::0;:13:::1;;::::0;:17:::1;:::i;:::-;3940:13;:33:::0;;-1:-1:-1;;3940:33:80::1;;::::0;;;::::1;::::0;;;::::1;::::0;;-1:-1:-1;;;;;3288:692:80:o;5826:138:34:-;2285:14;;;;;;;;;-1:-1:-1;;;;;2285:14:34;-1:-1:-1;;;;;2285:23:34;;:25;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;2271:39:34;:10;-1:-1:-1;;;;;2271:39:34;;2263:94;;;;-1:-1:-1;;;2263:94:34;;;;;;;:::i;:::-;;;;;;;;;5920:37:::1;5939:17;5920:18;:37::i;:::-;5826:138:::0;:::o;3832:392::-;2125:14;;:35;;-1:-1:-1;;;2125:35:34;;2149:10;2125:35;;;7290:51:82;-1:-1:-1;;;;;2125:14:34;;;;:23;;7263:18:82;;2125:35:34;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;2117:88;;;;-1:-1:-1;;;2117:88:34;;;;;;;:::i;:::-;4064:7:::1;::::0;4034:25;;::::1;4033:38;4025:107;;;::::0;-1:-1:-1;;;4025:107:34;;34749:2:82;4025:107:34::1;::::0;::::1;34731:21:82::0;34788:2;34768:18;;;34761:30;34827:34;34807:18;;;34800:62;34898:26;34878:18;;;34871:54;34942:19;;4025:107:34::1;34547:420:82::0;4025:107:34::1;4142:7;:25:::0;;;4182:35:::1;::::0;7748:25:82;;;4189:10:34::1;::::0;4182:35:::1;::::0;7736:2:82;7721:18;4182:35:34::1;;;;;;;;3832:392:::0;:::o;2772:425:80:-;3111:19:1;3134:13;;;;;;3133:14;;3179:34;;;;-1:-1:-1;3197:12:1;;3212:1;3197:12;;;;:16;3179:34;3178:108;;;-1:-1:-1;3258:4:1;1476:19:2;:23;;;3219:66:1;;-1:-1:-1;3268:12:1;;;;;:17;3219:66;3157:201;;;;-1:-1:-1;;;3157:201:1;;35174:2:82;3157:201:1;;;35156:21:82;35213:2;35193:18;;;35186:30;35252:34;35232:18;;;35225:62;-1:-1:-1;;;35303:18:82;;;35296:44;35357:19;;3157:201:1;34972:410:82;3157:201:1;3368:12;:16;;-1:-1:-1;;3368:16:1;3383:1;3368:16;;;3394:65;;;;3428:13;:20;;-1:-1:-1;;3428:20:1;;;;;3394:65;2985:47:80::1;3003:15;2000:1:34;2985:17:80;:47::i;:::-;3042:32;3061:12;3042:18;:32::i;:::-;3084:10;:24:::0;;-1:-1:-1;;;;;3084:24:80;;::::1;-1:-1:-1::0;;;;;;3084:24:80;;::::1;;::::0;;;3118:9:::1;:22:::0;;;;::::1;::::0;;::::1;;::::0;;3150:11:::1;:40:::0;;;;::::1;::::0;;;::::1;::::0;;;::::1;::::0;;3479:99:1;;;;3529:5;3513:21;;-1:-1:-1;;3513:21:1;;;3553:14;;-1:-1:-1;35539:36:82;;3553:14:1;;35527:2:82;35512:18;3553:14:1;;;;;;;3479:99;3101:483;2772:425:80;;;;;:::o;6576:854::-;6758:25;6786:31;;;;:12;:31;:::i;:::-;6866:36;;;6914:1;6866:36;;;:16;:36;;;;;;6758:59;;-1:-1:-1;6858:59:80;;;;;;7025:12;7039:20;7014:46;;;;;;;;;:::i;:::-;;;;-1:-1:-1;;7014:46:80;;;;;;;;;7004:57;;7014:46;7004:57;;;;6948:36;;;;;;;:16;:36;;;;;;:113;6927:144;;;;;;1280:3;7143:39;;;;:20;:39;:::i;:::-;:89;;;;:::i;:::-;7103:129;;7110:12;7103:129;;;;7082:160;;;;;;7253:11;;:121;;-1:-1:-1;;;7253:121:80;;-1:-1:-1;;;;;7253:11:80;;;;:21;;:121;;7288:18;;7320:4;;7345:19;;;;;7253:121;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;7389:34:80;;;;;;-1:-1:-1;7389:34:80;;-1:-1:-1;7389:34:80;;;6748:682;6576:854;;;:::o;13991::38:-;14188:22;14212;14321:13;2037:77:52;14372:7:38;14381:3;:5;;;14388:3;:5;;;14395;:7;;;14403:1;14395:10;;;;;;;:::i;:::-;;;;;14407:7;;14415:1;14407:10;;;;14419:5;:7;;;14427:1;14419:10;;;;;;;:::i;:::-;;;;;14431:5;:7;;;14439:1;14431:10;;;;;;;:::i;:::-;;;;;;;;;;14443:7;;14452;;;;14355:105;;;;;;;;;;;37381:19:82;;;37425:2;37416:12;;37409:28;;;;37462:2;37453:12;;37446:28;;;;37499:2;37490:12;;37483:28;;;;37536:3;37527:13;;37520:29;;;;37574:3;37565:13;;37558:29;37612:3;37603:13;;37596:29;37650:3;37641:13;;37634:29;37688:3;37679:13;;37672:29;37726:3;37717:13;;37028:708;14355:105:38;;;;;;;;;;;;;14345:116;;;;;;14337:125;;:144;;;;:::i;:::-;14321:160;-1:-1:-1;14564:274:38;14599:33;14610:21;:3;14321:160;14610:14;:21::i;:::-;14599:5;;:10;:33::i;:::-;14650:22;:20;:22::i;:::-;14690:67;14719:37;14750:5;14719:19;-1:-1:-1;;;;;;;;;;;;;;;;;2392:13:52;;;;;;;;2400:1;2392:13;;2403:1;2392:13;;;;;2313:99;14719:19:38;:30;;:37::i;:::-;14690:23;14705:7;14690:14;:23::i;:::-;:28;;:67::i;:::-;14775:5;998:6;14564:17;:274::i;:::-;14523:315;;;;-1:-1:-1;13991:854:38;-1:-1:-1;;;;;;13991:854:38:o;10650:380:39:-;10793:28;10861:9;:16;-1:-1:-1;;;;;10847:31:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;10847:31:39;;10833:45;;10893:9;10888:136;10912:9;:16;10908:1;:20;10888:136;;;10966:19;-1:-1:-1;;;;;10966:33:39;;11000:9;11010:1;11000:12;;;;;;;;:::i;:::-;;;;;;;10966:47;;;;;;;;;;;;;;-1:-1:-1;;;;;7308:32:82;;;;7290:51;;7278:2;7263:18;;7144:203;10966:47:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;10949:11;10961:1;10949:14;;;;;;;;:::i;:::-;;;;;;;;;;:64;10930:3;;;:::i;:::-;;;10888:136;;;;10650:380;;;;:::o;3037:1255::-;3205:19;3236:28;3267:19;-1:-1:-1;;;;;3267:33:39;;:35;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;3236:66;;3312:28;3343:19;-1:-1:-1;;;;;3343:33:39;;:35;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;3312:66;;3388:30;3421:19;-1:-1:-1;;;;;3421:34:39;;:36;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;3388:69;;3468:29;3517:13;:20;-1:-1:-1;;;;;3500:38:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;3468:70;;3553:9;3548:699;3572:13;:20;3568:1;:24;3548:699;;;3613:18;3640:13;3654:1;3640:16;;;;;;;;:::i;:::-;;;;;3702:69;;-1:-1:-1;;;3702:69:39;;3640:16;;;;;3702:69;;;39476:36:82;;;39560:10;39548:23;;39528:18;;;39521:51;3640:16:39;-1:-1:-1;;;;;;;;3702:42:39;;;;;39449:18:82;;3702:69:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;3702:69:39;;;;;;;;;;;;:::i;:::-;3671:100;;3815:11;:18;-1:-1:-1;;;;;3800:34:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;;;;;;;;;;;;;;;;;;;;;;3800:34:39;;-1:-1:-1;;3800:34:39;;;;;;;;;;;;3785:9;3795:1;3785:12;;;;;;;;:::i;:::-;;;;;;:49;;;;3853:9;3848:389;3872:11;:18;3868:1;:22;3848:389;;;3933:289;;;;;;;;3974:14;-1:-1:-1;;;;;3974:40:39;;4015:11;4027:1;4015:14;;;;;;;;:::i;:::-;;;;;;;3974:56;;;;;;;;;;;;;7748:25:82;;7736:2;7721:18;;7602:177;3974:56:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;3933:289:39;;;;;4072:11;4084:1;4072:14;;;;;;;;:::i;:::-;;;;;;;3933:289;;;;4116:13;-1:-1:-1;;;;;4116:35:39;;4160:11;4172:1;4160:14;;;;;;;;:::i;:::-;;;;;;;;;;;4116:87;;-1:-1:-1;;;;;;4116:87:39;;;;;;;;;;40665:25:82;;;;40738:4;40726:17;;40706:18;;;40699:45;4116:87:39;40780:23:82;;40760:18;;;40753:51;40638:18;;4116:87:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;3933:289:39;;;;3915:9;3925:1;3915:12;;;;;;;;:::i;:::-;;;;;;;3928:1;3915:15;;;;;;;;:::i;:::-;;;;;;:307;;;;3892:3;;;;;:::i;:::-;;;;3848:389;;;;3599:648;;3594:3;;;;;:::i;:::-;;;;3548:699;;;-1:-1:-1;4276:9:39;3037:1255;-1:-1:-1;;;;;;;3037:1255:39:o;2175:168:38:-;1466:19;-1:-1:-1;;;;;1466:25:38;;:27;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;1452:41:38;:10;-1:-1:-1;;;;;1452:41:38;;1444:146;;;;-1:-1:-1;;;1444:146:38;;41318:2:82;1444:146:38;;;41300:21:82;41357:2;41337:18;;;41330:30;41396:34;41376:18;;;41369:62;41467:34;41447:18;;;41440:62;41539:30;41518:19;;;41511:59;41587:19;;1444:146:38;41116:496:82;1444:146:38;2260:20:::1;:28:::0;;-1:-1:-1;;2260:28:38::1;::::0;::::1;;::::0;;::::1;::::0;;;2303:33:::1;::::0;21237:41:82;;;2303:33:38::1;::::0;21225:2:82;21210:18;2303:33:38::1;;;;;;;2175:168:::0;:::o;7714:5235:80:-;7958:25;7986:31;;;;:12;:31;:::i;:::-;7958:59;;8149:53;8232:28;:35;-1:-1:-1;;;;;8205:72:80;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;8205:72:80;;8149:128;;8293:6;8288:205;8309:28;:35;8305:1;:39;8288:205;;;8437:45;:28;8466:1;8437:31;;;;;;;;:::i;:::-;;;;;;;10534:9:52;;10473:16;10524:20;;;10580:4;10576:13;;;10570:20;10557:34;;;10629:4;10616:18;;;10404:246;8437:45:80;8365:36;8419:1;8365:69;;;;;;;;:::i;:::-;;;;;;;;;;:117;8346:3;;;;:::i;:::-;;;;8288:205;;;-1:-1:-1;9176:27:80;9263:21;;;;;;;;:::i;:::-;9302:36;9229:123;;;;;;;;;:::i;:::-;;;;;;;;;;;;;9206:156;;;;;;9176:186;;9403:20;:37;;;9380:19;:60;9372:69;;;;;;9508:46;9584:28;:35;-1:-1:-1;;;;;9557:72:80;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;9557:72:80;;9508:121;;9644:6;9639:254;9660:28;:35;9656:1;:39;9639:254;;;9791:14;-1:-1:-1;;;;;9751:90:80;;9842:36;9879:1;9842:39;;;;;;;;:::i;:::-;;;;;;;9751:131;;;;;;;;;;;;;7748:25:82;;7736:2;7721:18;;7602:177;9751:131:80;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;9716:29;9746:1;9716:32;;;;;;;;:::i;:::-;-1:-1:-1;;;;;9716:166:80;;;:32;;;;;;;;;;;:166;9697:3;;;;:::i;:::-;;;;9639:254;;;-1:-1:-1;9903:11:80;;:48;;-1:-1:-1;;;9903:48:80;;7121:10:82;7109:23;;9903:48:80;;;7091:42:82;-1:-1:-1;;;;;9903:11:80;;;;:28;;7064:18:82;;9903:48:80;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;12884:58:80;;12931:10;;-1:-1:-1;12884:58:80;;;;-1:-1:-1;12884:58:80;;;;;7948:5001;;;;7714:5235;;;;:::o;11441:390:39:-;11590:26;11654:11;:18;-1:-1:-1;;;;;11640:33:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;11640:33:39;;11628:45;;11688:9;11683:142;11707:11;:18;11703:1;:22;11683:142;;;11761:19;-1:-1:-1;;;;;11761:37:39;;11799:11;11811:1;11799:14;;;;;;;;:::i;:::-;;;;;;;11761:53;;;;;;;;;;;;;7748:25:82;;7736:2;7721:18;;7602:177;11761:53:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;11746:9;11756:1;11746:12;;;;;;;;:::i;:::-;-1:-1:-1;;;;;11746:68:39;;;:12;;;;;;;;;;;:68;11727:3;;;:::i;:::-;;;11683:142;;5476:3709;5716:29;-1:-1:-1;;;;;;;;;;;;;;;;;;;;;;;;;;;;;5716:29:39;5757:28;5788:19;-1:-1:-1;;;;;5788:33:39;;:35;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;5757:66;;5833:52;-1:-1:-1;;;;;;;;;;;;;;;;;;;;;;;;;;;;;5833:52:39;6065:99;;-1:-1:-1;;;6065:99:39;;-1:-1:-1;;;;;6065:55:39;;;;;:99;;6121:20;;6143;;;;6065:99;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;6065:99:39;;;;;;;;;;;;:::i;:::-;6011:153;;6322:84;;-1:-1:-1;;;6322:84:39;;-1:-1:-1;;;;;6322:47:39;;;;;:84;;6370:20;;6392:13;;;;6322:84;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;6322:84:39;;;;;;;;;;;;:::i;:::-;6279:40;;;:127;6487:13;-1:-1:-1;;;;;6472:36:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;6425:44:39;;;:83;6523:23;6518:2307;6552:40;;;;-1:-1:-1;6518:2307:39;;;6629:30;6883:20;-1:-1:-1;;;;;6870:41:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;6870:41:39;;6804:22;:44;;;6849:17;6804:63;;;;;;;;;;:::i;:::-;;;;;;:107;;;;6931:6;6926:1405;6943:31;;;6926:1405;;;7091:29;7144:19;-1:-1:-1;;;;;7144:55:39;;7225:20;;7246:1;7225:23;;;;;;;:::i;:::-;;;;;;;7275:20;7322:22;:51;;;7374:1;7322:54;;;;;;;;:::i;:::-;;;;;;;7144:254;;;;;;;;;;;;;;;;44768:25:82;;;44812:10;44858:15;;;44853:2;44838:18;;44831:43;44910:15;44905:2;44890:18;;44883:43;44756:2;44741:18;;44569:363;7144:254:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;7091:307;-1:-1:-1;;;;;;7441:26:39;;7433:131;;;;-1:-1:-1;;;7433:131:39;;45434:2:82;7433:131:39;;;45416:21:82;45473:2;45453:18;;;45446:30;45512:34;45492:18;;;45485:62;45583:34;45563:18;;;45556:62;45655:30;45634:19;;;45627:59;45703:19;;7433:131:39;45232:496:82;7433:131:39;7750:13;;7764:17;7750:32;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;7719:64:39;;7750:32;;;;;7719:64;;;;7787:1;7718:70;;;:75;7714:603;;-1:-1:-1;7714:603:39;;8024:13;-1:-1:-1;;;;;8024:46:39;;8096:20;;8117:1;8096:23;;;;;;;:::i;:::-;;;;;;;8151:13;;8165:17;8151:32;;;;;;;;;:::i;:::-;8024:228;;-1:-1:-1;;;;;;8024:228:39;;;;;;;;;;40665:25:82;;;;8151:32:39;;;;;;;40706:18:82;;;40699:45;-1:-1:-1;8024:228:39;40780:23:82;;40760:18;;;40753:51;40638:18;;8024:228:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;7934:22;:44;;;7979:17;7934:63;;;;;;;;;;:::i;:::-;;;;;;;7998:22;7934:87;;;;;;;;:::i;:::-;:318;;;;:87;;;;;;;;;;;:318;8274:24;;;;:::i;:::-;;;;7714:603;-1:-1:-1;6976:3:39;;;;:::i;:::-;;;;6926:1405;;;;8421:46;8483:22;-1:-1:-1;;;;;8470:36:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;8470:36:39;;8421:85;;8525:6;8520:185;8541:22;8537:1;:26;8520:185;;;8624:22;:44;;;8669:17;8624:63;;;;;;;;;;:::i;:::-;;;;;;;8688:1;8624:66;;;;;;;;:::i;:::-;;;;;;;8588:30;8619:1;8588:33;;;;;;;;:::i;:::-;:102;;;;:33;;;;;;;;;;;:102;8565:3;;;;:::i;:::-;;;;8520:185;;;;8784:30;8718:22;:44;;;8763:17;8718:63;;;;;;;;;;:::i;:::-;;;;;;:96;;;;6615:2210;;6594:19;;;;;:::i;:::-;;;;6518:2307;;;;8835:30;8868:19;-1:-1:-1;;;;;8868:34:39;;:36;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;9060:78;;-1:-1:-1;;;9060:78:39;;8835:69;;-1:-1:-1;;;;;;9060:41:39;;;;;:78;;9102:13;;;;9117:20;;9060:78;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;9060:78:39;;;;;;;;;;;;:::i;:::-;9018:39;;;:120;-1:-1:-1;9018:39:39;5476:3709;-1:-1:-1;;;;;;;;5476:3709:39:o;7436:272:80:-;7581:11;;:57;;-1:-1:-1;;;7581:57:80;;-1:-1:-1;;;;;7581:11:80;;;;:31;;:57;;7613:6;;7621:9;;;;7632:5;;;;7581:57;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;7653:48:80;;7690:10;;-1:-1:-1;7653:48:80;;;;-1:-1:-1;7653:48:80;;;;;7436:272;;;;;:::o;4299:136:34:-;2125:14;;:35;;-1:-1:-1;;;2125:35:34;;2149:10;2125:35;;;7290:51:82;-1:-1:-1;;;;;2125:14:34;;;;:23;;7263:18:82;;2125:35:34;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;2117:88;;;;-1:-1:-1;;;2117:88:34;;;;;;;:::i;:::-;-1:-1:-1;;4349:7:34::1;:27:::0;;;4391:37:::1;::::0;7748:25:82;;;4398:10:34::1;::::0;4391:37:::1;::::0;7736:2:82;7721:18;4391:37:34::1;;;;;;;4299:136::o:0;9602:654:39:-;9786:16;9814:35;9852:19;-1:-1:-1;;;;;9852:55:39;;9908:11;9921;9852:81;;;;;;;;;;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;9852:81:39;;;;;;;;;;;;:::i;:::-;9814:119;;9943:30;9990:11;:18;-1:-1:-1;;;;;9976:33:39;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;9976:33:39;;9943:66;;10024:9;10019:201;10043:11;:18;10039:1;:22;10019:201;;;10101:19;-1:-1:-1;;;;;10101:55:39;;10157:11;10169:1;10157:14;;;;;;;;:::i;:::-;;;;;;;10173:11;10186:19;10206:1;10186:22;;;;;;;;:::i;:::-;;;;;;;10101:108;;;;;;;;;;;;;;;;44768:25:82;;;44812:10;44858:15;;;44853:2;44838:18;;44831:43;44910:15;44905:2;44890:18;;44883:43;44756:2;44741:18;;44569:363;10101:108:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;10082:127:39;:13;10096:1;10082:16;;;;;;;;:::i;:::-;;;;;;;;;;:127;10063:3;;;;:::i;:::-;;;;10019:201;;;-1:-1:-1;10236:13:39;9602:654;-1:-1:-1;;;;;9602:654:39:o;2595:171:80:-;2718:11;;:41;;-1:-1:-1;;;2718:41:80;;47623:10:82;47611:23;;2718:41:80;;;47593:42:82;47651:18;;;47644:34;;;2692:7:80;;-1:-1:-1;;;;;2718:11:80;;:30;;47566:18:82;;2718:41:80;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;2711:48;;2595:171;;;;;:::o;4469:9040:38:-;-1:-1:-1;;;;;;;;;;;;;;;;;4751:7:38;4791:25;4783:93;;;;-1:-1:-1;;;4783:93:38;;48080:2:82;4783:93:38;;;48062:21:82;48119:2;48099:18;;;48092:30;-1:-1:-1;;;;;;;;;;;48138:18:82;;;48131:62;48229:25;48209:18;;;48202:53;48272:19;;4783:93:38;47878:419:82;4783:93:38;4933:17;;;;:24;4909:48;;4908:122;;;;-1:-1:-1;4999:23:38;;;;:30;4975:54;;4908:122;:195;;;;-1:-1:-1;5071:24:38;;;;:31;5047:55;;4908:195;:272;;;;-1:-1:-1;5144:28:38;;;;:35;5120:59;;4908:272;4887:384;;;;-1:-1:-1;;;4887:384:38;;48504:2:82;4887:384:38;;;48486:21:82;48543:2;48523:18;;;48516:30;-1:-1:-1;;;;;;;;;;;48562:18:82;;;48555:62;48653:34;48633:18;;;48626:62;-1:-1:-1;;;48704:19:82;;;48697:32;48746:19;;4887:384:38;48302:469:82;4887:384:38;5337:35;;:42;5303:23;;;;:30;:76;5282:191;;;;-1:-1:-1;;;5282:191:38;;48978:2:82;5282:191:38;;;48960:21:82;49017:2;48997:18;;;48990:30;;;-1:-1:-1;;;;;;;;;;;49036:18:82;;;49029:62;49127:34;49107:18;;;49100:62;-1:-1:-1;;;49178:19:82;;;49171:35;49223:19;;5282:191:38;48776:472:82;5282:191:38;5522:12;5492:43;;:20;:43;;;5484:116;;;;-1:-1:-1;;;5484:116:38;;49455:2:82;5484:116:38;;;49437:21:82;49494:2;49474:18;;;49467:30;-1:-1:-1;;;;;;;;;;;49513:18:82;;;49506:62;49604:30;49584:18;;;49577:58;49652:19;;5484:116:38;49253:424:82;5484:116:38;6117:19;;;;;;;;6090:24;6117:19;;;;;;;;;;;-1:-1:-1;;;;;;;;;;;;;;;;6117:19:38;6473:13;-1:-1:-1;;;;;6460:34:38;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;6460:34:38;-1:-1:-1;6426:31:38;;;:68;6552:13;-1:-1:-1;;;;;6539:34:38;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;6539:34:38;-1:-1:-1;6504:69:38;;-1:-1:-1;;;;;;;;;;;;;;;;;6666:6:38;:23;;;:30;-1:-1:-1;;;;;6652:45:38;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;6652:45:38;-1:-1:-1;6625:72:38;;6747:23;;;;:30;-1:-1:-1;;;;;6733:45:38;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;6733:45:38;;6707:10;:23;;:71;;;;6957:27;6987:87;7025:13;;6987:87;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;7040:33:38;;;-1:-1:-1;;;7040:33:38;;;;-1:-1:-1;;;;;7040:19:38;:31;;-1:-1:-1;7040:31:38;;-1:-1:-1;7040:33:38;;;;;;;;;;;;;;:31;:33;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;6987:37;:87::i;:::-;6957:117;;7094:9;7089:1638;7113:6;:23;;;:30;7109:1;:34;7089:1638;;;7405:40;:6;:23;;;7429:1;7405:26;;;;;;;;:::i;:40::-;7376:10;:23;;;7400:1;7376:26;;;;;;;;:::i;:::-;;;;;;;;;;:69;7467:6;;7463:277;;7576:23;;;;7600:5;7604:1;7600;:5;:::i;:::-;7576:30;;;;;;;;:::i;:::-;;;;;;;7568:39;;7538:10;:23;;;7562:1;7538:26;;;;;;;;:::i;:::-;;;;;;;7530:35;;:77;7497:224;;;;;-1:-1:-1;;;7497:224:38;;50266:2:82;7497:224:38;;;50248:21:82;50285:18;;;50278:30;;;;-1:-1:-1;;;;;;;;;;;50324:18:82;;;50317:62;50415:34;50395:18;;;50388:62;50467:19;;7497:224:38;50064:428:82;7497:224:38;7901:19;-1:-1:-1;;;;;7901:55:38;;7995:10;:23;;;8019:1;7995:26;;;;;;;;:::i;:::-;;;;;;;8060:20;8113:6;:35;;;8149:1;8113:38;;;;;;;;:::i;:::-;;;;;;;7901:273;;;;;;;;;;;;;;;;44768:25:82;;;44812:10;44858:15;;;44853:2;44838:18;;44831:43;44910:15;44905:2;44890:18;;44883:43;44756:2;44741:18;;44569:363;7901:273:38;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;7850:324:38;:10;:24;;;7875:1;7850:27;;;;;;;;:::i;:::-;;;;;;:324;;;;;8465:247;8495:199;8592:75;8647:19;8617:10;:24;;;8642:1;8617:27;;;;;;;;:::i;:::-;;;;;;;:49;8592:24;:75::i;:::-;8495:6;:23;;;8519:1;8495:26;;;;;;;;:::i;:::-;;;;;;;:67;;:199;;;;:::i;8465:247::-;8459:253;-1:-1:-1;7145:3:38;;;;:::i;:::-;;;;7089:1638;;;;6789:1948;9010:12;:3;:10;:12::i;:::-;9354:20;;9004:18;;-1:-1:-1;9354:20:38;;9325:26;9354:20;9420:65;;9484:1;9420:65;;;9444:10;-1:-1:-1;;;;;9444:35:38;;:37;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;9388:97;;9505:9;9500:3138;9520:24;;;9500:3138;;;9728:21;9724:368;;;9901:20;9806:115;;9877:21;9806:19;-1:-1:-1;;;;;9806:43:38;;9856:13;;9870:1;9856:16;;;;;;;:::i;:::-;9806:68;;;;;;-1:-1:-1;;;;;;9806:68:38;;;9856:16;;;;;9806:68;;;35539:36:82;-1:-1:-1;35512:18:82;;9806:68:38;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;:92;;;;:::i;:::-;:115;9773:300;;;;-1:-1:-1;;;9773:300:38;;51021:2:82;9773:300:38;;;51003:21:82;51060:3;51040:18;;;51033:31;-1:-1:-1;;;;;;;;;;;51080:18:82;;;51073:62;51171:34;51151:18;;;51144:62;51243:34;51222:19;;;51215:63;-1:-1:-1;;;51294:19:82;;;51287:37;51341:19;;9773:300:38;50819:547:82;9773:300:38;10361:14;-1:-1:-1;;;;;10361:46:38;;10458:13;;10472:1;10458:16;;;;;;;:::i;:::-;;;;;;;;;10452:23;;10518:20;10575:6;:23;;;10599:1;10575:26;;;;;;;;:::i;:::-;;;;;;;;;;;10361:267;;-1:-1:-1;;;;;;10361:267:38;;;;;;;51596:4:82;51584:17;;;10361:267:38;;;51566:36:82;10361:267:38;51667:15:82;;;51647:18;;;51640:43;51719:15;51699:18;;;51692:43;51539:18;;10361:267:38;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;10289:339:38;;10297:34;:6;:17;;;10315:1;10297:20;;;;;;;;:::i;:34::-;-1:-1:-1;;10289:339:38;;10260:507;;;;-1:-1:-1;;;10260:507:38;;52247:2:82;10260:507:38;;;52229:21:82;52286:2;52266:18;;;52259:30;-1:-1:-1;;;;;;;;;;;52305:18:82;;;52298:62;52396:34;52376:18;;;52369:62;52468:34;52447:19;;;52440:63;-1:-1:-1;;;52519:19:82;;;52512:32;52561:19;;10260:507:38;52045:541:82;10260:507:38;10791:30;10800:6;:17;;;10818:1;10800:20;;;;;;;;:::i;:::-;;;;;;;10791:3;:8;;:30;;;;:::i;:::-;10785:36;;10996:13;-1:-1:-1;;;;;10996:49:38;;11092:13;;11106:1;11092:16;;;;;;;:::i;:::-;;;;;;;;;11086:23;;11148:20;11201:6;:24;;;11226:1;11201:27;;;;;;;;:::i;:::-;;;;;;;;;;;10996:255;;-1:-1:-1;;;;;;10996:255:38;;;;;;;51596:4:82;51584:17;;;10996:255:38;;;51566:36:82;10996:255:38;51667:15:82;;;51647:18;;;51640:43;51719:15;51699:18;;;51692:43;51539:18;;10996:255:38;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;10938:11;:31;;;10970:1;10938:34;;;;;;;;:::i;:::-;-1:-1:-1;;;;;10938:313:38;;;:34;;;;;;;;;;:313;11307:31;;;:34;;11339:1;;11307:34;;;;;;:::i;:::-;;;;;;;11269:11;:32;;;11302:1;11269:35;;;;;;;;:::i;:::-;;;;;;:72;-1:-1:-1;;;;;11269:72:38;;;-1:-1:-1;;;;;11269:72:38;;;;;11428:31;11727:9;11722:902;11746:6;:23;;;:30;11742:1;:34;11722:902;;;11918:71;11936:10;:24;;;11961:1;11936:27;;;;;;;;:::i;:::-;;;;;;;11971:13;;11985:1;11971:16;;;;;;;:::i;:::-;7404:1:53;11971:16:38;;;;;7387:13:53;;;;7386:19;;7380:26;;;-1:-1:-1;7292:121:53;11918:71:38;11914:692;;;12084:13;-1:-1:-1;;;;;12084:43:38;;12182:13;;12196:1;12182:16;;;;;;;:::i;:::-;;;;;;;;;12176:23;;12246:20;12312:10;:23;;;12336:1;12312:26;;;;;;;;:::i;:::-;;;;;;;12379:6;:28;;;12408:1;12379:31;;;;;;;;:::i;:::-;;;;;;;12411:23;12379:56;;;;;;;;:::i;:::-;;;;;;;;;;;12084:382;;-1:-1:-1;;;;;;12084:382:38;;;;;;;52845:4:82;52833:17;;;12084:382:38;;;52815:36:82;12084:382:38;52916:15:82;;;52896:18;;;52889:43;52948:18;;;52941:34;;;;53011:15;52991:18;;;52984:43;52787:19;;12084:382:38;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;12017:32;;:35;;12050:1;;12017:35;;;;;;:::i;:::-;;;;;;:449;;;;;;;:::i;:::-;-1:-1:-1;;;;;12017:449:38;;;-1:-1:-1;12532:25:38;;;;;11914:692;11778:3;;;;:::i;:::-;;;;11722:902;;;;9551:3087;9546:3;;;;;:::i;:::-;;;;9500:3138;;;;9311:3337;;12708:22;12732:21;12757:153;12805:7;12831:3;12853:6;:12;;;12884:6;:12;;;12757:30;:153::i;:::-;12707:203;;;;12932:17;12924:97;;;;-1:-1:-1;;;12924:97:38;;53482:2:82;12924:97:38;;;53464:21:82;53521:2;53501:18;;;53494:30;-1:-1:-1;;;;;;;;;;;53540:18:82;;;53533:62;53631:34;53611:18;;;53604:62;-1:-1:-1;;;53682:19:82;;;53675:34;53726:19;;12924:97:38;53280:471:82;12924:97:38;13043:16;13035:86;;;;-1:-1:-1;;;13035:86:38;;53958:2:82;13035:86:38;;;53940:21:82;53997:2;53977:18;;;53970:30;-1:-1:-1;;;;;;;;;;;54016:18:82;;;54009:62;54107:27;54087:18;;;54080:55;54152:19;;13035:86:38;53756:421:82;13035:86:38;12657:475;;13206:27;13263:20;13285:10;:23;;;13246:63;;;;;;;;;:::i;:::-;;;;;;;-1:-1:-1;;13246:63:38;;;;;;13236:74;;13246:63;13236:74;;;;13469:11;;13236:74;;-1:-1:-1;4469:9040:38;;-1:-1:-1;;;;;;;;;4469:9040:38:o;2071:101:0:-;1334:13;:11;:13::i;:::-;2135:30:::1;2162:1;2135:18;:30::i;:::-;2071:101::o:0;4041:2433:80:-;2024:10;;-1:-1:-1;;;;;2024:10:80;2010;:24;2002:33;;;;;;4248:23:::1;4274:21;::::0;;;::::1;::::0;::::1;;:::i;:::-;4248:47:::0;-1:-1:-1;4305:28:80::1;;4336:18;;::::0;::::1;:4:::0;:18:::1;:::i;:::-;4305:49:::0;;-1:-1:-1;4305:49:80;-1:-1:-1;4364:32:80::1;4399:30;::::0;;;::::1;::::0;::::1;;:::i;:::-;4364:65:::0;-1:-1:-1;4607:13:80::1;:46;4621:31;;::::0;::::1;:12:::0;:31:::1;:::i;:::-;4607:46;;;;;;;;;;;;;;;;4581:4;4570:16;;;;;;;;:::i;:::-;;;;;;;;;;;;;4560:27;;;;;;:93;4539:124;;;::::0;::::1;;4786:1;4725:16;4786:1:::0;4742:31:::1;;::::0;::::1;:12:::0;:31:::1;:::i;:::-;4725:49;;;;;;;;;;;;;;;;:63;4704:94;;;::::0;::::1;;4869:45;4888:26;4869:16:::0;:45:::1;:::i;:::-;4829:85;;4836:12;4829:85;;;;4808:116;;;::::0;::::1;;5055:15;5094:12;5083:24;;;;;;;;:::i;:::-;;;;;;;;;;;;;5073:35;;;;;;5055:53;;5168:42;5224:24:::0;5261:164:::1;5294:7;5319:13;;5350:16;5384:27;5261:15;:164::i;:::-;5154:271;;;;5526:6;5521:507;5538:24:::0;;::::1;5521:507;;;5977:25;5904:99;;:17;:37;;;5942:1;5904:40;;;;;;;;:::i;:::-;;;;;;;:99;;;;:::i;:::-;-1:-1:-1::0;;;;;5794:209:80::1;1340:3;5794:17;:38;;;5833:1;5794:41;;;;;;;;:::i;:::-;;;;;;;-1:-1:-1::0;;;;;5794:86:80::1;;;;;:::i;:::-;:209;;5769:248;;;::::0;::::1;;5564:3:::0;::::1;::::0;::::1;:::i;:::-;;;;5521:507;;;-1:-1:-1::0;6089:94:80::1;::::0;;;;::::1;::::0;;::::1;6130:12;6089:94;::::0;;::::1;::::0;;::::1;::::0;;;6320:46;;6089:94;;6320:46:::1;::::0;6331:12;;6089:94;;6320:46:::1;;:::i;:::-;;;;;;;;;;;;;6297:79;;;;;;6245:16;:49;6262:12;:31;;;;;;;;;;:::i;:::-;6245:49;;;;;;;;;;;;;;;:131;;;;6418:49;6432:12;6446:20;6418:49;;;;;;;:::i;:::-;;;;;;;;4238:2236;;;;;;;;4041:2433:::0;;;:::o;1757:712:39:-;1991:16;;;2005:1;1991:16;;;;;;;;;1920:7;;1929:19;;1920:7;;1991:16;;;;;;;;;;;-1:-1:-1;1991:16:39;1960:47;;2034:10;2017:11;2029:1;2017:14;;;;;;;;:::i;:::-;;;;;;;;;;:27;2070:81;;-1:-1:-1;;;2070:81:39;;2054:13;;-1:-1:-1;;;;;2070:55:39;;;;;:81;;2126:11;;2139;;2070:81;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;;2070:81:39;;;;;;;;;;;;:::i;:::-;2152:1;2070:84;;;;;;;;:::i;:::-;;;;;;;;;;;2192:87;;-1:-1:-1;;;2192:87:39;;;;;57327:25:82;;;2054:100:39;57388:23:82;;;57368:18;;;57361:51;2054:100:39;;;57428:18:82;;;57421:34;;;2054:100:39;-1:-1:-1;2169:20:39;;-1:-1:-1;;;;;2192:55:39;;;;;57300:18:82;;2192:87:39;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;2169:110:39;;;2290:26;2319:44;2350:12;2319:30;:44::i;:::-;2290:73;;2382:12;2396:65;2413:19;2434:13;2449:11;2396:16;:65::i;:::-;2374:88;;;;;;;;1757:712;;;;;;:::o;2321:198:0:-;1334:13;:11;:13::i;:::-;-1:-1:-1;;;;;2409:22:0;::::1;2401:73;;;::::0;-1:-1:-1;;;2401:73:0;;57668:2:82;2401:73:0::1;::::0;::::1;57650:21:82::0;57707:2;57687:18;;;57680:30;57746:34;57726:18;;;57719:62;-1:-1:-1;;;57797:18:82;;;57790:36;57843:19;;2401:73:0::1;57466:402:82::0;2401:73:0::1;2484:28;2503:8;2484:18;:28::i;4911:437:34:-:0;2285:14;;;;;;;;;-1:-1:-1;;;;;2285:14:34;-1:-1:-1;;;;;2285:23:34;;:25;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;-1:-1:-1;;;;;2271:39:34;:10;-1:-1:-1;;;;;2271:39:34;;2263:94;;;;-1:-1:-1;;;2263:94:34;;;;;;;:::i;:::-;5164:7:::1;;5163:8;5141:15;5140:16;5128:7;;5127:8;5126:31;5125:47;5104:150;;;::::0;-1:-1:-1;;;5104:150:34;;58075:2:82;5104:150:34::1;::::0;::::1;58057:21:82::0;58114:2;58094:18;;;58087:30;58153:34;58133:18;;;58126:62;58224:26;58204:18;;;58197:54;58268:19;;5104:150:34::1;57873:420:82::0;5104:150:34::1;5264:7;:25:::0;;;5304:37:::1;::::0;7748:25:82;;;5313:10:34::1;::::0;5304:37:::1;::::0;7736:2:82;7721:18;5304:37:34::1;7602:177:82::0;6024:360:34;-1:-1:-1;;;;;6127:40:34;;6106:160;;;;-1:-1:-1;;;6106:160:34;;58500:2:82;6106:160:34;;;58482:21:82;58539:2;58519:18;;;58512:30;58578:34;58558:18;;;58551:62;58649:34;58629:18;;;58622:62;-1:-1:-1;;;58700:19:82;;;58693:40;58750:19;;6106:160:34;58298:477:82;6106:160:34;6299:14;;6281:52;;;-1:-1:-1;;;;;6299:14:34;;;59040:34:82;;59110:15;;;59105:2;59090:18;;59083:43;6281:52:34;;58975:18:82;6281:52:34;;;;;;;6343:14;:34;;-1:-1:-1;;;;;;6343:34:34;-1:-1:-1;;;;;6343:34:34;;;;;;;;;;6024:360::o;2943:441::-;3077:14;;-1:-1:-1;;;;;3077:14:34;3069:37;:79;;;;-1:-1:-1;;;;;;3110:38:34;;;;3069:79;3048:197;;;;-1:-1:-1;;;3048:197:34;;59339:2:82;3048:197:34;;;59321:21:82;59378:2;59358:18;;;59351:30;59417:34;59397:18;;;59390:62;59488:34;59468:18;;;59461:62;-1:-1:-1;;;59539:19:82;;;59532:38;59587:19;;3048:197:34;59137:475:82;3048:197:34;3255:7;:26;;;3296:36;;7748:25:82;;;3303:10:34;;3296:36;;7736:2:82;7721:18;3296:36:34;;;;;;;3342:35;3361:15;3342:18;:35::i;:::-;2943:441;;:::o;2673:187:0:-;2765:6;;;-1:-1:-1;;;;;2781:17:0;;;-1:-1:-1;;;;;;2781:17:0;;;;;;;2813:40;;2765:6;;;2781:17;2765:6;;2813:40;;2746:16;;2813:40;2736:124;2673:187;:::o;7084:580:52:-;-1:-1:-1;;;;;;;;;;;;;;;;;7184:23:52;;:::i;:::-;7228:3;;7217:14;;:8;7252:3;;;;7241:8;;;:14;7265:8;;;;:12;;;-1:-1:-1;;7452:1:52;7446:4;7217:14;7436:1;7429:4;7422:5;7418:16;7407:53;7396:64;-1:-1:-1;7396:64:52;7557:48;;;;7530:75;;7557:48;7582:9;7530:75;;7632:7;7624:33;;;;-1:-1:-1;;;7624:33:52;;59819:2:82;7624:33:52;;;59801:21:82;59858:2;59838:18;;;59831:30;-1:-1:-1;;;59877:18:82;;;59870:43;59930:18;;7624:33:52;59617:337:82;7624:33:52;7174:490;;7084:580;;;;:::o;4823:615::-;-1:-1:-1;;;;;;;;;;;;;;;;;4926:23:52;;:::i;:::-;4970:4;;4959:15;;:8;4995:4;;;;4984:8;;;:15;5020:4;;5009:8;;;;:15;;;;5045:4;;;;5034:8;;;:15;-1:-1:-1;;5225:1:52;5219:4;4959:15;5209:1;5202:4;5195:5;5191:16;5180:53;5169:64;-1:-1:-1;5169:64:52;5330:48;;;;5303:75;5406:7;5398:33;;;;-1:-1:-1;;;5398:33:52;;60161:2:82;5398:33:52;;;60143:21:82;60200:2;60180:18;;;60173:30;-1:-1:-1;;;60219:18:82;;;60212:43;60272:18;;5398:33:52;59959:337:82;4070:128:52;4119:14;;:::i;:::-;-1:-1:-1;4152:39:52;;;;;;;;3635:77;4152:39;;;;;;3752:77;4152:39;;;;;;;;;;;;;;3869:77;4152:39;;3986:77;4152:39;;;;;;;;;;;;;;;4070:128::o;11044:451::-;-1:-1:-1;;;;;;;;;;;;;;;;;11123:12:52;;;11185:24;-1:-1:-1;;;;;;;;;;;11193:2:52;11185:24;:::i;:::-;11173:36;;11220:239;11259:13;11270:1;11259:10;:13::i;:::-;11247:25;;-1:-1:-1;11247:25:52;-1:-1:-1;;;;;;;;;;;;11336:1:52;11333;11326:24;11318:4;:32;11314:92;;;11378:13;;;;;;;;;;;;;;;;;;;;11044:451;-1:-1:-1;;;11044:451:52:o;11314:92::-;-1:-1:-1;;;;;;;;;;;11434:1:52;11431;11424:24;11420:28;;11220:239;;9189:1112;9397:31;;;;;;;;;;;;;;;;;;9438;;;;;;;;;;;;;;;;9375:4;;;;9397:31;9480:24;;:::i;:::-;9520:9;9515:302;9539:1;9535;:5;9515:302;;;9561:9;9573:5;:1;9577;9573:5;:::i;:::-;9561:17;;9607:2;9610:1;9607:5;;;;;;;:::i;:::-;;;;;:7;9592:5;9598;:1;9607:7;9598:5;:::i;:::-;9592:12;;;;;;;:::i;:::-;;;;:22;9643:2;9646:1;9643:5;;;;;;;:::i;:::-;;;;;:7;;;9628:5;9634:1;9638;9634:5;;;;:::i;:::-;9628:12;;;;;;;:::i;:::-;;;;:22;9679:2;9682:1;9679:5;;;;;;;:::i;:::-;;;;;:7;:10;9664:5;9670;:1;9674;9670:5;:::i;:::-;9664:12;;;;;;;:::i;:::-;;;;:25;9718:2;9721:1;9718:5;;;;;;;:::i;:::-;;;;;:7;9726:1;9718:10;;;;9703:5;9709;:1;9713;9709:5;:::i;:::-;9703:12;;;;;;;:::i;:::-;;;;:25;9757:2;9760:1;9757:5;;;;;;;:::i;:::-;;;;;:7;;;9765:1;9757:10;;;;;;;:::i;:::-;;;;;9742:5;9748;:1;9752;9748:5;:::i;:::-;9742:12;;;;;;;:::i;:::-;;;;:25;9796:2;9799:1;9796:5;;;;;;;:::i;:::-;;;;;:7;;;9804:1;9796:10;;;;;;;:::i;:::-;;;;;9781:5;9787;:1;9791;9787:5;:::i;:::-;9781:12;;;;;;;:::i;:::-;;;;:25;-1:-1:-1;9542:3:52;;;;:::i;:::-;;;;9515:302;;;;9827:21;;:::i;:::-;9858:12;10032:4;10027:3;10012:13;10005:5;10002:1;9990:10;9979:58;10282:6;;9968:69;;10282:11;;;;-1:-1:-1;10265:29:52;;-1:-1:-1;;;;;;;;;;9189:1112:52:o;3308:360:53:-;3419:7;3438:14;3455:44;3481:17;3455:25;:44::i;:::-;3438:61;;3541:6;3524:13;3519:18;;:1;:18;;3518:29;3510:127;;;;-1:-1:-1;;;3510:127:53;;60503:2:82;3510:127:53;;;60485:21:82;60542:2;60522:18;;;60515:30;60581:34;60561:18;;;60554:62;60652:33;60632:18;;;60625:61;60703:19;;3510:127:53;60301:427:82;6797:406:53;6853:6;;6897:209;6904:5;;6897:209;;6931:5;6935:1;6931;:5;:::i;:::-;6925:12;;;;7020:7;;;;:::i;:::-;;;;6897:209;;5698:1197:52;-1:-1:-1;;;;;;;;;;;;;;;;;5824:4:52;5820:1;:8;;;5812:37;;;;-1:-1:-1;;;5812:37:52;;61137:2:82;5812:37:52;;;61119:21:82;61176:2;61156:18;;;61149:30;-1:-1:-1;;;61195:18:82;;;61188:46;61251:18;;5812:37:52;60935:340:82;5812:37:52;5893:1;:6;;5898:1;5893:6;5890:44;;;-1:-1:-1;5922:1:52;5915:8;;5890:44;6016:19;;;;;;;;;5989:24;6016:19;;;;;;;;;6145:1;;6208;;6337:481;6348:1;6343:6;;:1;:6;;;6337:481;;6493:1;6483:6;;;;;;;6482:12;;:17;6478:84;;;6529:14;6534:3;6539;6529:4;:14::i;:::-;6523:20;;6478:84;6644:14;6649:3;6654;6644:4;:14::i;:::-;6638:20;-1:-1:-1;6765:7:52;6771:1;6765:7;;;;;6790:3;6337:481;;;-1:-1:-1;6885:3:52;;5698:1197;-1:-1:-1;;;;;5698:1197:52:o;4461:295::-;-1:-1:-1;;;;;;;;;;;;;;;;;4600:3:52;;:8;:20;;;;-1:-1:-1;4612:3:52;;;;:8;4600:20;4596:154;;;-1:-1:-1;;4643:13:52;;;;;;;;;-1:-1:-1;4643:13:52;;;;;;;;4461:295::o;4596:154::-;4694:45;;;;;;;;4702:1;:3;;;4694:45;;;;-1:-1:-1;;;;;;;;;;;4721:1:52;:3;;;:16;;;;:::i;:::-;4707:31;;-1:-1:-1;;;;;;;;;;;4707:31:52;:::i;:::-;4694:45;;4687:52;4461:295;-1:-1:-1;;4461:295:52:o;4596:154::-;4461:295;;;:::o;1599:130:0:-;1513:6;;-1:-1:-1;;;;;1513:6:0;929:10:3;1662:23:0;1654:68;;;;-1:-1:-1;;;1654:68:0;;61482:2:82;1654:68:0;;;61464:21:82;;;61501:18;;;61494:30;61560:34;61540:18;;;61533:62;61612:18;;1654:68:0;61280:356:82;5465:1257:53;5532:12;5650:15;5728:23;5764:20;5777:6;5764:12;:20::i;:::-;5754:31;;-1:-1:-1;;;;;5754:31:53;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;5754:31:53;;5728:57;;5841:18;6093:9;6088:601;6122:10;:17;6109:10;:30;6108:45;;;;;6149:3;6145:1;:7;6108:45;6088:601;;;6252:1;:6;;;-1:-1:-1;6339:16:53;;;:21;6335:344;;6524:1;6511:16;;6486:10;6497;6486:22;;;;;;;;:::i;:::-;;;;:41;-1:-1:-1;;;;;6486:41:53;;;;;;;;;6650:12;;;;;6335:344;6155:3;;;:::i;:::-;;;6088:601;;;-1:-1:-1;6705:10:53;;5465:1257;-1:-1:-1;;;;5465:1257:53:o;11616:433:52:-;11670:7;;;-1:-1:-1;;;;;;;;;;;11801:1:52;-1:-1:-1;;;;;;;;;;;11785:1:52;-1:-1:-1;;;;;;;;;;;11769:1:52;11766;11759:24;11752:47;11745:70;11730:85;;11912:9;11924:91;11931:4;11937:65;-1:-1:-1;;;;;;;;;;;11924:6:52;:91::i;:::-;12034:4;;11912:103;;-1:-1:-1;11616:433:52;;-1:-1:-1;;;11616:433:52:o;1188:1693:53:-;1278:7;571:3;1409:17;:24;:49;;1401:142;;;;-1:-1:-1;;;1401:142:53;;61843:2:82;1401:142:53;;;61825:21:82;61882:2;61862:18;;;61855:30;;;61921:34;61901:18;;;61894:62;61992:34;61972:18;;;61965:62;-1:-1:-1;;;62043:19:82;;;62036:35;62088:19;;1401:142:53;61641:472:82;1401:142:53;1619:24;;1615:77;;-1:-1:-1;1679:1:53;;1188:1693;-1:-1:-1;1188:1693:53:o;1615:77::-;1770:14;1873:15;2180:17;2198:1;2180:20;;;;;;;;:::i;:::-;;;;;2169:1;2180:20;;;;;2169:32;;;;-1:-1:-1;2284:568:53;2308:17;:24;2304:1;:28;2284:568;;;2480:17;2498:1;2480:20;;;;;;;;:::i;:::-;;;;;2469:1;2480:20;;;;;2469:32;;-1:-1:-1;2665:16:53;;;2657:100;;;;-1:-1:-1;;;2657:100:53;;62320:2:82;2657:100:53;;;62302:21:82;62359:2;62339:18;;;62332:30;62398:34;62378:18;;;62371:62;62469:34;62449:18;;;62442:62;-1:-1:-1;;;62520:19:82;;;62513:38;62568:19;;2657:100:53;62118:475:82;2657:100:53;2824:16;;;;2334:3;;;:::i;:::-;;;2284:568;;;-1:-1:-1;2868:6:53;;1188:1693;-1:-1:-1;;;1188:1693:53:o;12055:874:52:-;12146:14;12172:12;12194:24;;:::i;:::-;12228:20;;:::i;:::-;12269:4;12258:15;;;12341:8;;;:15;;;12425:8;;;:15;;;12509:8;;;:16;;;12535:8;;;:20;;;12565:8;;;:19;;;12673:6;12667:4;12258:15;12571:1;12650:4;12643:5;12639:16;12628:58;12617:69;-1:-1:-1;12617:69:52;12783:48;;;;12756:75;12858:7;12850:46;;;;-1:-1:-1;;;12850:46:52;;62800:2:82;12850:46:52;;;62782:21:82;62839:2;62819:18;;;62812:30;62878:28;62858:18;;;62851:56;62924:18;;12850:46:52;62598:350:82;12850:46:52;-1:-1:-1;12913:9:52;;;-1:-1:-1;;;;;12055:874:52:o;-1:-1:-1:-;;;;;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;;;;;:::i;:::-;;;;;;;:::i;:::-;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;14:121:82:-;99:10;92:5;88:22;81:5;78:33;68:61;;125:1;122;115:12;140:132;207:20;;236:30;207:20;236:30;:::i;277:347::-;328:8;338:6;392:3;385:4;377:6;373:17;369:27;359:55;;410:1;407;400:12;359:55;-1:-1:-1;433:20:82;;-1:-1:-1;;;;;465:30:82;;462:50;;;508:1;505;498:12;462:50;545:4;537:6;533:17;521:29;;597:3;590:4;581:6;573;569:19;565:30;562:39;559:59;;;614:1;611;604:12;559:59;277:347;;;;;:::o;629:689::-;741:6;749;757;765;818:3;806:9;797:7;793:23;789:33;786:53;;;835:1;832;825:12;786:53;873:3;862:9;858:19;896:7;892:2;889:15;886:35;;;917:1;914;907:12;886:35;940:9;;-1:-1:-1;971:16:82;996:30;971:16;996:30;:::i;:::-;1045:5;-1:-1:-1;1101:3:82;1086:19;;1073:33;-1:-1:-1;;;;;1118:30:82;;1115:50;;;1161:1;1158;1151:12;1115:50;1200:58;1250:7;1241:6;1230:9;1226:22;1200:58;:::i;:::-;629:689;;;;-1:-1:-1;1277:8:82;-1:-1:-1;;;;629:689:82:o;1323:148::-;-1:-1:-1;;;;;1415:31:82;;1405:42;;1395:70;;1461:1;1458;1451:12;1476:288;1559:6;1612:2;1600:9;1591:7;1587:23;1583:32;1580:52;;;1628:1;1625;1618:12;1580:52;1667:9;1654:23;1686:48;1728:5;1686:48;:::i;1769:180::-;1828:6;1881:2;1869:9;1860:7;1856:23;1852:32;1849:52;;;1897:1;1894;1887:12;1849:52;-1:-1:-1;1920:23:82;;1769:180;-1:-1:-1;1769:180:82:o;1954:922::-;2073:6;2081;2089;2097;2105;2158:3;2146:9;2137:7;2133:23;2129:33;2126:53;;;2175:1;2172;2165:12;2126:53;2214:9;2201:23;2233:48;2275:5;2233:48;:::i;:::-;2300:5;-1:-1:-1;2357:2:82;2342:18;;2329:32;2370:50;2329:32;2370:50;:::i;:::-;2439:7;-1:-1:-1;2498:2:82;2483:18;;2470:32;2511:50;2470:32;2511:50;:::i;:::-;2580:7;-1:-1:-1;2639:2:82;2624:18;;2611:32;2652:50;2611:32;2652:50;:::i;:::-;2721:7;-1:-1:-1;2780:3:82;2765:19;;2752:33;2794:50;2752:33;2794:50;:::i;:::-;2863:7;2853:17;;;1954:922;;;;;;;;:::o;2881:153::-;2938:5;2983:3;2974:6;2969:3;2965:16;2961:26;2958:46;;;3000:1;2997;2990:12;2958:46;-1:-1:-1;3022:6:82;2881:153;-1:-1:-1;2881:153:82:o;3039:160::-;3104:5;3149:2;3140:6;3135:3;3131:16;3127:25;3124:45;;;3165:1;3162;3155:12;3204:634;3380:6;3388;3396;3449:3;3437:9;3428:7;3424:23;3420:33;3417:53;;;3466:1;3463;3456:12;3417:53;3506:9;3493:23;-1:-1:-1;;;;;3531:6:82;3528:30;3525:50;;;3571:1;3568;3561:12;3525:50;3594:64;3650:7;3641:6;3630:9;3626:22;3594:64;:::i;:::-;3584:74;;;3677:68;3737:7;3732:2;3721:9;3717:18;3677:68;:::i;:::-;3667:78;;3764:68;3824:7;3819:2;3808:9;3804:18;3764:68;:::i;:::-;3754:78;;3204:634;;;;;:::o;3843:127::-;3904:10;3899:3;3895:20;3892:1;3885:31;3935:4;3932:1;3925:15;3959:4;3956:1;3949:15;3975:257;4047:4;4041:11;;;4079:17;;-1:-1:-1;;;;;4111:34:82;;4147:22;;;4108:62;4105:88;;;4173:18;;:::i;:::-;4209:4;4202:24;3975:257;:::o;4237:255::-;4309:2;4303:9;4351:6;4339:19;;-1:-1:-1;;;;;4373:34:82;;4409:22;;;4370:62;4367:88;;;4435:18;;:::i;4497:275::-;4568:2;4562:9;4633:2;4614:13;;-1:-1:-1;;4610:27:82;4598:40;;-1:-1:-1;;;;;4653:34:82;;4689:22;;;4650:62;4647:88;;;4715:18;;:::i;:::-;4751:2;4744:22;4497:275;;-1:-1:-1;4497:275:82:o;4777:282::-;4831:5;4879:4;4867:9;4862:3;4858:19;4854:30;4851:50;;;4897:1;4894;4887:12;4851:50;4919:22;;:::i;:::-;4910:31;;4977:9;4964:23;4957:5;4950:38;5048:2;5037:9;5033:18;5020:32;5015:2;5008:5;5004:14;4997:56;4777:282;;;;:::o;5064:646::-;5114:5;5167:3;5160:4;5152:6;5148:17;5144:27;5134:55;;5185:1;5182;5175:12;5134:55;5218:2;5212:9;5260:2;5252:6;5248:15;5329:6;5317:10;5314:22;-1:-1:-1;;;;;5281:10:82;5278:34;5275:62;5272:88;;;5340:18;;:::i;:::-;5380:10;5376:2;5369:22;;5411:6;5452:2;5444:6;5440:15;5478:3;5470:6;5467:15;5464:35;;;5495:1;5492;5485:12;5464:35;5519:6;5534:146;5550:6;5545:3;5542:15;5534:146;;;5618:17;;5606:30;;5665:4;5656:14;;;;5567;5534:146;;5715:320;5769:5;5817:4;5805:9;5800:3;5796:19;5792:30;5789:50;;;5835:1;5832;5825:12;5789:50;5857:22;;:::i;:::-;5848:31;;5902:40;5938:3;5927:9;5902:40;:::i;:::-;5895:5;5888:55;5977:51;6024:3;6017:4;6006:9;6002:20;5977:51;:::i;:::-;5970:4;5963:5;5959:16;5952:77;5715:320;;;;:::o;6040:530::-;6204:6;6212;6220;6228;6281:3;6269:9;6260:7;6256:23;6252:33;6249:53;;;6298:1;6295;6288:12;6249:53;6334:9;6321:23;6311:33;;6363:54;6409:7;6404:2;6393:9;6389:18;6363:54;:::i;:::-;6353:64;;6436:54;6482:7;6477:2;6466:9;6462:18;6436:54;:::i;:::-;6426:64;;6509:55;6556:7;6550:3;6539:9;6535:19;6509:55;:::i;:::-;6499:65;;6040:530;;;;;;;:::o;7352:245::-;7410:6;7463:2;7451:9;7442:7;7438:23;7434:32;7431:52;;;7479:1;7476;7469:12;7431:52;7518:9;7505:23;7537:30;7561:5;7537:30;:::i;7784:183::-;7844:4;-1:-1:-1;;;;;7869:6:82;7866:30;7863:56;;;7899:18;;:::i;:::-;-1:-1:-1;7944:1:82;7940:14;7956:4;7936:25;;7784:183::o;7972:1171::-;8095:6;8103;8156:2;8144:9;8135:7;8131:23;8127:32;8124:52;;;8172:1;8169;8162:12;8124:52;8211:9;8198:23;8230:48;8272:5;8230:48;:::i;:::-;8297:5;-1:-1:-1;8321:2:82;8359:18;;;8346:32;-1:-1:-1;;;;;8390:30:82;;8387:50;;;8433:1;8430;8423:12;8387:50;8456:22;;8509:4;8501:13;;8497:27;-1:-1:-1;8487:55:82;;8538:1;8535;8528:12;8487:55;8574:2;8561:16;8597:60;8613:43;8653:2;8613:43;:::i;:::-;8597:60;:::i;:::-;8691:15;;;8773:1;8769:10;;;;8761:19;;8757:28;;;8722:12;;;;8797:19;;;8794:39;;;8829:1;8826;8819:12;8794:39;8853:11;;;;8873:240;8889:6;8884:3;8881:15;8873:240;;;8971:3;8958:17;8988:50;9030:7;8988:50;:::i;:::-;9051:20;;8906:12;;;;9091;;;;8873:240;;;9132:5;9122:15;;;;;;;7972:1171;;;;;:::o;9148:435::-;9201:3;9239:5;9233:12;9266:6;9261:3;9254:19;9292:4;9321:2;9316:3;9312:12;9305:19;;9358:2;9351:5;9347:14;9379:1;9389:169;9403:6;9400:1;9397:13;9389:169;;;9464:13;;9452:26;;9498:12;;;;9533:15;;;;9425:1;9418:9;9389:169;;;-1:-1:-1;9574:3:82;;9148:435;-1:-1:-1;;;;;9148:435:82:o;9588:261::-;9767:2;9756:9;9749:21;9730:4;9787:56;9839:2;9828:9;9824:18;9816:6;9787:56;:::i;9854:1017::-;9969:6;9977;9985;10038:2;10026:9;10017:7;10013:23;10009:32;10006:52;;;10054:1;10051;10044:12;10006:52;10093:9;10080:23;10112:48;10154:5;10112:48;:::i;:::-;10179:5;-1:-1:-1;10203:2:82;10241:18;;;10228:32;-1:-1:-1;;;;;10309:14:82;;;10306:34;;;10336:1;10333;10326:12;10306:34;10374:6;10363:9;10359:22;10349:32;;10419:7;10412:4;10408:2;10404:13;10400:27;10390:55;;10441:1;10438;10431:12;10390:55;10477:2;10464:16;10499:2;10495;10492:10;10489:36;;;10505:18;;:::i;:::-;10547:53;10590:2;10571:13;;-1:-1:-1;;10567:27:82;10563:36;;10547:53;:::i;:::-;10534:66;;10623:2;10616:5;10609:17;10663:7;10658:2;10653;10649;10645:11;10641:20;10638:33;10635:53;;;10684:1;10681;10674:12;10635:53;10739:2;10734;10730;10726:11;10721:2;10714:5;10710:14;10697:45;10783:1;10778:2;10773;10766:5;10762:14;10758:23;10751:34;;10804:5;10794:15;;;;;10828:37;10861:2;10850:9;10846:18;10828:37;:::i;10876:1336::-;10947:3;10985:5;10979:12;11012:6;11007:3;11000:19;11038:4;11079:2;11074:3;11070:12;11104:11;11131;11124:18;;11181:6;11178:1;11174:14;11167:5;11163:26;11151:38;;11223:2;11216:5;11212:14;11244:1;11265;11275:911;11291:6;11286:3;11283:15;11275:911;;;11356:16;;;11344:29;;11396:13;;11468:9;;11490:22;;;11576:11;;;;11534:13;;;;11611:1;11625:455;11641:8;11636:3;11633:17;11625:455;;;11714:15;;11764:9;;-1:-1:-1;;;;;11760:35:82;11746:50;;11842:11;;;11836:18;11820:14;;;11813:42;11882:4;11936:11;;;11930:18;-1:-1:-1;;;;;11926:51:82;11910:14;;;11903:75;12049:17;;;;12015:4;12004:16;;;;11792:1;11660:11;11625:455;;;-1:-1:-1;;12164:12:82;;;;12101:5;-1:-1:-1;;12129:15:82;;;;11317:1;11308:11;11275:911;;;-1:-1:-1;12202:4:82;;10876:1336;-1:-1:-1;;;;;;;;10876:1336:82:o;12217:381::-;12498:2;12487:9;12480:21;12461:4;12518:74;12588:2;12577:9;12573:18;12565:6;12518:74;:::i;12603:118::-;12689:5;12682:13;12675:21;12668:5;12665:32;12655:60;;12711:1;12708;12701:12;12726:241;12782:6;12835:2;12823:9;12814:7;12810:23;12806:32;12803:52;;;12851:1;12848;12841:12;12803:52;12890:9;12877:23;12909:28;12931:5;12909:28;:::i;12972:689::-;13033:5;13086:3;13079:4;13071:6;13067:17;13063:27;13053:55;;13104:1;13101;13094:12;13053:55;13140:6;13127:20;13166:4;13190:60;13206:43;13246:2;13206:43;:::i;13190:60::-;13284:15;;;13370:1;13366:10;;;;13354:23;;13350:32;;;13315:12;;;;13394:15;;;13391:35;;;13422:1;13419;13412:12;13391:35;13458:2;13450:6;13446:15;13470:162;13486:6;13481:3;13478:15;13470:162;;;13554:35;13585:3;13580;13554:35;:::i;:::-;13542:48;;13610:12;;;;13512:4;13503:14;13470:162;;;-1:-1:-1;13650:5:82;12972:689;-1:-1:-1;;;;;;12972:689:82:o;13666:915::-;13902:6;13910;13918;13926;13979:3;13967:9;13958:7;13954:23;13950:33;13947:53;;;13996:1;13993;13986:12;13947:53;14036:9;14023:23;-1:-1:-1;;;;;14106:2:82;14098:6;14095:14;14092:34;;;14122:1;14119;14112:12;14092:34;14145:64;14201:7;14192:6;14181:9;14177:22;14145:64;:::i;:::-;14135:74;;14228:68;14288:7;14283:2;14272:9;14268:18;14228:68;:::i;:::-;14218:78;;14315:68;14375:7;14370:2;14359:9;14355:18;14315:68;:::i;:::-;14305:78;;14436:3;14425:9;14421:19;14408:33;14392:49;;14466:2;14456:8;14453:16;14450:36;;;14482:1;14479;14472:12;14450:36;;14505:70;14567:7;14556:8;14545:9;14541:24;14505:70;:::i;:::-;14495:80;;;13666:915;;;;;;;:::o;14586:662::-;14640:5;14693:3;14686:4;14678:6;14674:17;14670:27;14660:55;;14711:1;14708;14701:12;14660:55;14747:6;14734:20;14773:4;14797:60;14813:43;14853:2;14813:43;:::i;14797:60::-;14891:15;;;14977:1;14973:10;;;;14961:23;;14957:32;;;14922:12;;;;15001:15;;;14998:35;;;15029:1;15026;15019:12;14998:35;15065:2;15057:6;15053:15;15077:142;15093:6;15088:3;15085:15;15077:142;;;15159:17;;15147:30;;15197:12;;;;15110;;15077:142;;15253:530;15376:6;15384;15437:2;15425:9;15416:7;15412:23;15408:32;15405:52;;;15453:1;15450;15443:12;15405:52;15492:9;15479:23;15511:48;15553:5;15511:48;:::i;:::-;15578:5;-1:-1:-1;15634:2:82;15619:18;;15606:32;-1:-1:-1;;;;;15650:30:82;;15647:50;;;15693:1;15690;15683:12;15647:50;15716:61;15769:7;15760:6;15749:9;15745:22;15716:61;:::i;:::-;15706:71;;;15253:530;;;;;:::o;15788:658::-;15959:2;16011:21;;;16081:13;;15984:18;;;16103:22;;;15930:4;;15959:2;16182:15;;;;16156:2;16141:18;;;15930:4;16225:195;16239:6;16236:1;16233:13;16225:195;;;16304:13;;-1:-1:-1;;;;;16300:39:82;16288:52;;16395:15;;;;16360:12;;;;16336:1;16254:9;16225:195;;;-1:-1:-1;16437:3:82;;15788:658;-1:-1:-1;;;;;;15788:658:82:o;16451:367::-;16514:8;16524:6;16578:3;16571:4;16563:6;16559:17;16555:27;16545:55;;16596:1;16593;16586:12;16545:55;-1:-1:-1;16619:20:82;;-1:-1:-1;;;;;16651:30:82;;16648:50;;;16694:1;16691;16684:12;16648:50;16731:4;16723:6;16719:17;16707:29;;16791:3;16784:4;16774:6;16771:1;16767:14;16759:6;16755:27;16751:38;16748:47;16745:67;;;16808:1;16805;16798:12;16823:1067;16976:6;16984;16992;17000;17008;17016;17069:3;17057:9;17048:7;17044:23;17040:33;17037:53;;;17086:1;17083;17076:12;17037:53;17125:9;17112:23;17144:48;17186:5;17144:48;:::i;:::-;17211:5;-1:-1:-1;17268:2:82;17253:18;;17240:32;17281;17240;17281;:::i;:::-;17332:7;-1:-1:-1;17390:2:82;17375:18;;17362:32;-1:-1:-1;;;;;17443:14:82;;;17440:34;;;17470:1;17467;17460:12;17440:34;17509:58;17559:7;17550:6;17539:9;17535:22;17509:58;:::i;:::-;17586:8;;-1:-1:-1;17483:84:82;-1:-1:-1;17674:2:82;17659:18;;17646:32;;-1:-1:-1;17690:16:82;;;17687:36;;;17719:1;17716;17709:12;17687:36;;17758:72;17822:7;17811:8;17800:9;17796:24;17758:72;:::i;:::-;16823:1067;;;;-1:-1:-1;16823:1067:82;;-1:-1:-1;16823:1067:82;;17849:8;;16823:1067;-1:-1:-1;;;16823:1067:82:o;17895:451::-;17947:3;17985:5;17979:12;18012:6;18007:3;18000:19;18038:4;18067:2;18062:3;18058:12;18051:19;;18104:2;18097:5;18093:14;18125:1;18135:186;18149:6;18146:1;18143:13;18135:186;;;18214:13;;18229:10;18210:30;18198:43;;18261:12;;;;18296:15;;;;18171:1;18164:9;18135:186;;18351:1491;18523:4;18552:2;18581;18570:9;18563:21;18619:6;18613:13;18662:4;18657:2;18646:9;18642:18;18635:32;18690:62;18747:3;18736:9;18732:19;18718:12;18690:62;:::i;:::-;18676:76;;18801:2;18793:6;18789:15;18783:22;18828:2;18824:7;18895:2;18883:9;18875:6;18871:22;18867:31;18862:2;18851:9;18847:18;18840:59;18922:51;18966:6;18950:14;18922:51;:::i;:::-;18908:65;;19022:2;19014:6;19010:15;19004:22;18982:44;;19090:2;19078:9;19070:6;19066:22;19062:31;19057:2;19046:9;19042:18;19035:59;19117:51;19161:6;19145:14;19117:51;:::i;:::-;19217:2;19205:15;;19199:22;19263;;;19259:31;;19252:4;19237:20;;19230:61;19340:21;;19370:22;;;19103:65;;-1:-1:-1;19505:23:82;;;-1:-1:-1;19408:15:82;;;;19466:1;19462:14;;;19450:27;;19446:36;;19546:1;19556:257;19570:6;19567:1;19564:13;19556:257;;;19656:2;19647:6;19639;19635:19;19631:28;19626:3;19619:41;19683:50;19726:6;19717;19711:13;19683:50;:::i;:::-;19756:15;;;;19791:12;;;;19673:60;-1:-1:-1;19592:1:82;19585:9;19556:257;;;-1:-1:-1;19830:6:82;18351:1491;-1:-1:-1;;;;;;;;;18351:1491:82:o;19847:878::-;19961:6;19969;19977;19985;19993;20046:2;20034:9;20025:7;20021:23;20017:32;20014:52;;;20062:1;20059;20052:12;20014:52;20101:9;20088:23;20120:30;20144:5;20120:30;:::i;:::-;20169:5;-1:-1:-1;20225:2:82;20210:18;;20197:32;-1:-1:-1;;;;;20278:14:82;;;20275:34;;;20305:1;20302;20295:12;20275:34;20344:70;20406:7;20397:6;20386:9;20382:22;20344:70;:::i;:::-;20433:8;;-1:-1:-1;20318:96:82;-1:-1:-1;20521:2:82;20506:18;;20493:32;;-1:-1:-1;20537:16:82;;;20534:36;;;20566:1;20563;20556:12;20534:36;;20605:60;20657:7;20646:8;20635:9;20631:24;20605:60;:::i;:::-;19847:878;;;;-1:-1:-1;19847:878:82;;-1:-1:-1;20684:8:82;;20579:86;19847:878;-1:-1:-1;;;19847:878:82:o;20730:114::-;20814:4;20807:5;20803:16;20796:5;20793:27;20783:55;;20834:1;20831;20824:12;20849:243;20906:6;20959:2;20947:9;20938:7;20934:23;20930:32;20927:52;;;20975:1;20972;20965:12;20927:52;21014:9;21001:23;21033:29;21056:5;21033:29;:::i;21289:669::-;21420:6;21428;21436;21489:2;21477:9;21468:7;21464:23;21460:32;21457:52;;;21505:1;21502;21495:12;21457:52;21544:9;21531:23;21563:48;21605:5;21563:48;:::i;:::-;21630:5;-1:-1:-1;21686:2:82;21671:18;;21658:32;-1:-1:-1;;;;;21702:30:82;;21699:50;;;21745:1;21742;21735:12;21699:50;21768:61;21821:7;21812:6;21801:9;21797:22;21768:61;:::i;:::-;21758:71;;;21881:2;21870:9;21866:18;21853:32;21894;21918:7;21894:32;:::i;:::-;21945:7;21935:17;;;21289:669;;;;;:::o;21963:632::-;22134:2;22186:21;;;22256:13;;22159:18;;;22278:22;;;22105:4;;22134:2;22357:15;;;;22331:2;22316:18;;;22105:4;22400:169;22414:6;22411:1;22408:13;22400:169;;;22475:13;;22463:26;;22544:15;;;;22509:12;;;;22436:1;22429:9;22400:169;;23015:313;23082:6;23090;23143:2;23131:9;23122:7;23118:23;23114:32;23111:52;;;23159:1;23156;23149:12;23111:52;23198:9;23185:23;23217:30;23241:5;23217:30;:::i;:::-;23266:5;23318:2;23303:18;;;;23290:32;;-1:-1:-1;;;23015:313:82:o;23803:735::-;23856:5;23909:3;23902:4;23894:6;23890:17;23886:27;23876:55;;23927:1;23924;23917:12;23876:55;23963:6;23950:20;23989:4;24013:60;24029:43;24069:2;24029:43;:::i;24013:60::-;24107:15;;;24193:1;24189:10;;;;24177:23;;24173:32;;;24138:12;;;;24217:15;;;24214:35;;;24245:1;24242;24235:12;24214:35;24281:2;24273:6;24269:15;24293:216;24309:6;24304:3;24301:15;24293:216;;;24389:3;24376:17;24406:30;24430:5;24406:30;:::i;:::-;24449:18;;24487:12;;;;24326;;24293:216;;24543:908;24606:5;24659:3;24652:4;24644:6;24640:17;24636:27;24626:55;;24677:1;24674;24667:12;24626:55;24713:6;24700:20;24739:4;24763:60;24779:43;24819:2;24779:43;:::i;24763:60::-;24857:15;;;24943:1;24939:10;;;;24927:23;;24923:32;;;24888:12;;;;24967:15;;;24964:35;;;24995:1;24992;24985:12;24964:35;25031:2;25023:6;25019:15;25043:379;25059:6;25054:3;25051:15;25043:379;;;25145:3;25132:17;-1:-1:-1;;;;;25168:11:82;25165:35;25162:125;;;25241:1;25270:2;25266;25259:14;25162:125;25312:67;25375:3;25370:2;25356:11;25348:6;25344:24;25340:33;25312:67;:::i;:::-;25300:80;;-1:-1:-1;25400:12:82;;;;25076;;25043:379;;25456:1566;25530:5;25578:6;25566:9;25561:3;25557:19;25553:32;25550:52;;;25598:1;25595;25588:12;25550:52;25620:22;;:::i;:::-;25611:31;;25678:9;25665:23;-1:-1:-1;;;;;25748:2:82;25740:6;25737:14;25734:34;;;25764:1;25761;25754:12;25734:34;25791:56;25843:3;25834:6;25823:9;25819:22;25791:56;:::i;:::-;25784:5;25777:71;25901:2;25890:9;25886:18;25873:32;25857:48;;25930:2;25920:8;25917:16;25914:36;;;25946:1;25943;25936:12;25914:36;25982:66;26044:3;26033:8;26022:9;26018:24;25982:66;:::i;:::-;25977:2;25970:5;25966:14;25959:90;26102:2;26091:9;26087:18;26074:32;26058:48;;26131:2;26121:8;26118:16;26115:36;;;26147:1;26144;26137:12;26115:36;26183:66;26245:3;26234:8;26223:9;26219:24;26183:66;:::i;:::-;26178:2;26171:5;26167:14;26160:90;26282:50;26328:3;26323:2;26312:9;26308:18;26282:50;:::i;:::-;26277:2;26270:5;26266:14;26259:74;26367:51;26414:3;26408;26397:9;26393:19;26367:51;:::i;:::-;26360:4;26353:5;26349:16;26342:77;26472:3;26461:9;26457:19;26444:33;26428:49;;26502:2;26492:8;26489:16;26486:36;;;26518:1;26515;26508:12;26486:36;26556:58;26610:3;26599:8;26588:9;26584:24;26556:58;:::i;:::-;26549:4;26542:5;26538:16;26531:84;26668:3;26657:9;26653:19;26640:33;26624:49;;26698:2;26688:8;26685:16;26682:36;;;26714:1;26711;26704:12;26682:36;26752:58;26806:3;26795:8;26784:9;26780:24;26752:58;:::i;:::-;26745:4;26738:5;26734:16;26727:84;26864:3;26853:9;26849:19;26836:33;26820:49;;26894:2;26884:8;26881:16;26878:36;;;26910:1;26907;26900:12;26878:36;;26947:68;27011:3;27000:8;26989:9;26985:24;26947:68;:::i;:::-;26941:3;26934:5;26930:15;26923:93;;25456:1566;;;;:::o;27027:896::-;27169:6;27177;27185;27193;27201;27254:3;27242:9;27233:7;27229:23;27225:33;27222:53;;;27271:1;27268;27261:12;27222:53;27307:9;27294:23;27284:33;;27368:2;27357:9;27353:18;27340:32;-1:-1:-1;;;;;27432:2:82;27424:6;27421:14;27418:34;;;27448:1;27445;27438:12;27418:34;27487:58;27537:7;27528:6;27517:9;27513:22;27487:58;:::i;:::-;27564:8;;-1:-1:-1;27461:84:82;-1:-1:-1;27649:2:82;27634:18;;27621:32;;-1:-1:-1;27662:30:82;27621:32;27662:30;:::i;:::-;27711:5;;-1:-1:-1;27769:2:82;27754:18;;27741:32;;27785:16;;;27782:36;;;27814:1;27811;27804:12;27782:36;;27837:80;27909:7;27898:8;27887:9;27883:24;27837:80;:::i;:::-;27827:90;;;27027:896;;;;;;;;:::o;27928:467::-;27980:3;28018:5;28012:12;28045:6;28040:3;28033:19;28071:4;28100:2;28095:3;28091:12;28084:19;;28137:2;28130:5;28126:14;28158:1;28168:202;28182:6;28179:1;28176:13;28168:202;;;28247:13;;-1:-1:-1;;;;;28243:46:82;28231:59;;28310:12;;;;28345:15;;;;28204:1;28197:9;28168:202;;28400:645;28629:2;28618:9;28611:21;28592:4;28667:6;28661:13;28710:2;28705;28694:9;28690:18;28683:30;28736:62;28793:3;28782:9;28778:19;28764:12;28736:62;:::i;:::-;28722:76;;28847:4;28839:6;28835:17;28829:24;28921:2;28917:7;28905:9;28897:6;28893:22;28889:36;28884:2;28873:9;28869:18;28862:64;28943:51;28987:6;28971:14;28943:51;:::i;:::-;28935:59;;;;29032:6;29025:4;29014:9;29010:20;29003:36;28400:645;;;;;:::o;29282:774::-;29463:6;29471;29479;29532:3;29520:9;29511:7;29507:23;29503:33;29500:53;;;29549:1;29546;29539:12;29500:53;29589:9;29576:23;-1:-1:-1;;;;;29659:2:82;29651:6;29648:14;29645:34;;;29675:1;29672;29665:12;29645:34;29698:64;29754:7;29745:6;29734:9;29730:22;29698:64;:::i;:::-;29688:74;;29781:68;29841:7;29836:2;29825:9;29821:18;29781:68;:::i;:::-;29771:78;;29902:2;29891:9;29887:18;29874:32;29858:48;;29931:2;29921:8;29918:16;29915:36;;;29947:1;29944;29937:12;29915:36;;29970:80;30042:7;30031:8;30020:9;30016:24;29970:80;:::i;:::-;29960:90;;;29282:774;;;;;:::o;30061:501::-;30167:6;30175;30183;30236:2;30224:9;30215:7;30211:23;30207:32;30204:52;;;30252:1;30249;30242:12;30204:52;30291:9;30278:23;30310:48;30352:5;30310:48;:::i;:::-;30377:5;-1:-1:-1;30429:2:82;30414:18;;30401:32;;-1:-1:-1;30485:2:82;30470:18;;30457:32;30498;30457;30498;:::i;30567:452::-;30876:6;30865:9;30858:25;30919:2;30914;30903:9;30899:18;30892:30;30839:4;30939:74;31009:2;30998:9;30994:18;30986:6;30939:74;:::i;:::-;30931:82;30567:452;-1:-1:-1;;;;30567:452:82:o;31528:1306::-;31695:2;31747:21;;;31787:13;;31666:4;;31695:2;31666:4;31720:18;;;31901:167;31915:4;31912:1;31909:11;31901:167;;;31974:13;;31962:26;;32043:15;;;;31935:1;31928:9;;;;;32008:12;;31901:167;;;31905:3;;;32133:10;32127:2;32119:6;32115:15;32109:22;32105:39;32099:3;32088:9;32084:19;32077:68;32192:4;32184:6;32180:17;32174:24;32217:6;32260:2;32254:3;32243:9;32239:19;32232:31;32292:12;32286:19;32342:6;32336:3;32325:9;32321:19;32314:35;32369:1;32379:157;32395:6;32390:3;32387:15;32379:157;;;32497:22;;;32493:31;;32487:38;32460:19;;;32481:3;32456:29;32449:77;32412:12;;32379:157;;;32556:6;32551:3;32548:15;32545:94;;;32627:1;32621:3;32612:6;32601:9;32597:22;32593:32;32586:43;32545:94;-1:-1:-1;32688:4:82;32676:17;;32670:24;6924:10;6913:22;;32737:18;;;6901:35;32670:24;-1:-1:-1;32817:2:82;32796:15;-1:-1:-1;;32792:29:82;32777:45;;;;32824:3;32773:55;;31528:1306;-1:-1:-1;;;;;31528:1306:82:o;32839:127::-;32900:10;32895:3;32891:20;32888:1;32881:31;32931:4;32928:1;32921:15;32955:4;32952:1;32945:15;32971:228;33010:3;33038:10;33075:2;33072:1;33068:10;33105:2;33102:1;33098:10;33136:3;33132:2;33128:12;33123:3;33120:21;33117:47;;;33144:18;;:::i;:::-;33180:13;;32971:228;-1:-1:-1;;;;32971:228:82:o;33204:268::-;33274:6;33327:2;33315:9;33306:7;33302:23;33298:32;33295:52;;;33343:1;33340;33333:12;33295:52;33375:9;33369:16;33394:48;33436:5;33394:48;:::i;33477:406::-;33679:2;33661:21;;;33718:2;33698:18;;;33691:30;33757:34;33752:2;33737:18;;33730:62;-1:-1:-1;;;33823:2:82;33808:18;;33801:40;33873:3;33858:19;;33477:406::o;33888:245::-;33955:6;34008:2;33996:9;33987:7;33983:23;33979:32;33976:52;;;34024:1;34021;34014:12;33976:52;34056:9;34050:16;34075:28;34097:5;34075:28;:::i;34138:404::-;34340:2;34322:21;;;34379:2;34359:18;;;34352:30;34418:34;34413:2;34398:18;;34391:62;-1:-1:-1;;;34484:2:82;34469:18;;34462:38;34532:3;34517:19;;34138:404::o;35586:265::-;35689:5;35676:19;35704:32;35728:7;35704:32;:::i;:::-;35770:10;35757:24;35745:37;;35838:4;35827:16;;;35814:30;35798:14;;35791:54;35586:265::o;35856:591::-;36162:3;36147:19;;36175:58;36151:9;36215:6;36175:58;:::i;:::-;36268:6;36255:20;36284:30;36308:5;36284:30;:::i;:::-;36361:10;36350:22;36345:2;36330:18;;36323:50;36434:4;36422:17;;;;36409:31;36404:2;36389:18;;;36382:59;35856:591;;-1:-1:-1;35856:591:82:o;36452:439::-;36731:10;36719:23;;36701:42;;36688:3;36673:19;;36793:4;36785:6;36780:2;36765:18;;36752:46;36832:3;36817:19;;;;36867:18;;;;36452:439;;-1:-1:-1;;36452:439:82:o;36896:127::-;36957:10;36952:3;36948:20;36945:1;36938:31;36988:4;36985:1;36978:15;37012:4;37009:1;37002:15;37873:209;37905:1;37931;37921:132;;37975:10;37970:3;37966:20;37963:1;37956:31;38010:4;38007:1;38000:15;38038:4;38035:1;38028:15;37921:132;-1:-1:-1;38067:9:82;;37873:209::o;38087:184::-;38157:6;38210:2;38198:9;38189:7;38185:23;38181:32;38178:52;;;38226:1;38223;38216:12;38178:52;-1:-1:-1;38249:16:82;;38087:184;-1:-1:-1;38087:184:82:o;38276:135::-;38315:3;-1:-1:-1;;38336:17:82;;38333:43;;;38356:18;;:::i;:::-;-1:-1:-1;38403:1:82;38392:13;;38276:135::o;39583:881::-;39678:6;39709:2;39752;39740:9;39731:7;39727:23;39723:32;39720:52;;;39768:1;39765;39758:12;39720:52;39801:9;39795:16;-1:-1:-1;;;;;39826:6:82;39823:30;39820:50;;;39866:1;39863;39856:12;39820:50;39889:22;;39942:4;39934:13;;39930:27;-1:-1:-1;39920:55:82;;39971:1;39968;39961:12;39920:55;40000:2;39994:9;40023:60;40039:43;40079:2;40039:43;:::i;40023:60::-;40117:15;;;40199:1;40195:10;;;;40187:19;;40183:28;;;40148:12;;;;40223:19;;;40220:39;;;40255:1;40252;40245:12;40220:39;40279:11;;;;40299:135;40315:6;40310:3;40307:15;40299:135;;;40381:10;;40369:23;;40332:12;;;;40412;;;;40299:135;;;40453:5;39583:881;-1:-1:-1;;;;;;;39583:881:82:o;40815:296::-;40884:6;40937:2;40925:9;40916:7;40912:23;40908:32;40905:52;;;40953:1;40950;40943:12;40905:52;40985:9;40979:16;-1:-1:-1;;;;;41028:5:82;41024:38;41017:5;41014:49;41004:77;;41077:1;41074;41067:12;41617:644;41865:10;41860:3;41856:20;41847:6;41842:3;41838:16;41834:43;41829:3;41822:56;41804:3;41909:1;41904:3;41900:11;41940:6;41934:13;41989:4;42028:2;42020:6;42016:15;42049:1;42059:175;42073:6;42070:1;42067:13;42059:175;;;42136:13;;42122:28;;42172:14;;;;42209:15;;;;42095:1;42088:9;42059:175;;;-1:-1:-1;42250:5:82;;41617:644;-1:-1:-1;;;;;;;41617:644:82:o;42266:354::-;42354:19;;;42336:3;-1:-1:-1;;;;;42385:31:82;;42382:51;;;42429:1;42426;42419:12;42382:51;42465:6;42462:1;42458:14;42517:8;42510:5;42503:4;42498:3;42494:14;42481:45;42594:1;42549:18;;42569:4;42545:29;42583:13;;;-1:-1:-1;42545:29:82;;42266:354;-1:-1:-1;;42266:354:82:o;42625:374::-;42852:10;42844:6;42840:23;42829:9;42822:42;42900:2;42895;42884:9;42880:18;42873:30;42803:4;42920:73;42989:2;42978:9;42974:18;42966:6;42958;42920:73;:::i;:::-;42912:81;42625:374;-1:-1:-1;;;;;42625:374:82:o;43004:954::-;43098:6;43129:2;43172;43160:9;43151:7;43147:23;43143:32;43140:52;;;43188:1;43185;43178:12;43140:52;43221:9;43215:16;-1:-1:-1;;;;;43246:6:82;43243:30;43240:50;;;43286:1;43283;43276:12;43240:50;43309:22;;43362:4;43354:13;;43350:27;-1:-1:-1;43340:55:82;;43391:1;43388;43381:12;43340:55;43420:2;43414:9;43443:60;43459:43;43499:2;43459:43;:::i;43443:60::-;43537:15;;;43619:1;43615:10;;;;43607:19;;43603:28;;;43568:12;;;;43643:19;;;43640:39;;;43675:1;43672;43665:12;43640:39;43699:11;;;;43719:209;43735:6;43730:3;43727:15;43719:209;;;43808:3;43802:10;43825:30;43849:5;43825:30;:::i;:::-;43868:18;;43752:12;;;;43906;;;;43719:209;;43963:266;44051:6;44046:3;44039:19;44103:6;44096:5;44089:4;44084:3;44080:14;44067:43;-1:-1:-1;44155:1:82;44130:16;;;44148:4;44126:27;;;44119:38;;;;44211:2;44190:15;;;-1:-1:-1;;44186:29:82;44177:39;;;44173:50;;43963:266::o;44234:330::-;44429:10;44421:6;44417:23;44406:9;44399:42;44477:2;44472;44461:9;44457:18;44450:30;44380:4;44497:61;44554:2;44543:9;44539:18;44531:6;44523;44497:61;:::i;44937:290::-;45007:6;45060:2;45048:9;45039:7;45035:23;45031:32;45028:52;;;45076:1;45073;45066:12;45028:52;45102:16;;-1:-1:-1;;;;;45147:31:82;;45137:42;;45127:70;;45193:1;45190;45183:12;45733:249;45802:6;45855:2;45843:9;45834:7;45830:23;45826:32;45823:52;;;45871:1;45868;45861:12;45823:52;45903:9;45897:16;45922:30;45946:5;45922:30;:::i;45987:175::-;46024:3;46068:4;46061:5;46057:16;46097:4;46088:7;46085:17;46082:43;;;46105:18;;:::i;:::-;46154:1;46141:15;;45987:175;-1:-1:-1;;45987:175:82:o;46167:331::-;46351:2;46340:9;46333:21;46314:4;46371:61;46428:2;46417:9;46413:18;46405:6;46397;46371:61;:::i;:::-;46363:69;;46480:10;46472:6;46468:23;46463:2;46452:9;46448:18;46441:51;46167:331;;;;;;:::o;46503:561::-;46786:10;46778:6;46774:23;46763:9;46756:42;46834:2;46829;46818:9;46814:18;46807:30;46737:4;46860:73;46929:2;46918:9;46914:18;46906:6;46898;46860:73;:::i;:::-;46981:9;46973:6;46969:22;46964:2;46953:9;46949:18;46942:50;47009:49;47051:6;47043;47035;47009:49;:::i;:::-;47001:57;46503:561;-1:-1:-1;;;;;;;;46503:561:82:o;47069:347::-;47286:10;47278:6;47274:23;47263:9;47256:42;47334:2;47329;47318:9;47314:18;47307:30;47237:4;47354:56;47406:2;47395:9;47391:18;47383:6;47354:56;:::i;49682:247::-;49750:6;49803:2;49791:9;49782:7;49778:23;49774:32;49771:52;;;49819:1;49816;49809:12;49771:52;49851:9;49845:16;49870:29;49893:5;49870:29;:::i;49934:125::-;49974:4;50002:1;49999;49996:8;49993:34;;;50007:18;;:::i;:::-;-1:-1:-1;50044:9:82;;49934:125::o;50686:128::-;50726:3;50757:1;50753:6;50750:1;50747:13;50744:39;;;50763:18;;:::i;:::-;-1:-1:-1;50799:9:82;;50686:128::o;51746:294::-;51816:6;51869:2;51857:9;51848:7;51844:23;51840:32;51837:52;;;51885:1;51882;51875:12;51837:52;51911:16;;-1:-1:-1;;51956:35:82;;51946:46;;51936:74;;52006:1;52003;51996:12;53038:237;53077:4;-1:-1:-1;;;;;53182:10:82;;;;53152;;53204:12;;;53201:38;;;53219:18;;:::i;:::-;53256:13;;53038:237;-1:-1:-1;;;53038:237:82:o;54182:521::-;54259:4;54265:6;54325:11;54312:25;54419:2;54415:7;54404:8;54388:14;54384:29;54380:43;54360:18;54356:68;54346:96;;54438:1;54435;54428:12;54346:96;54465:33;;54517:20;;;-1:-1:-1;;;;;;54549:30:82;;54546:50;;;54592:1;54589;54582:12;54546:50;54625:4;54613:17;;-1:-1:-1;54656:14:82;54652:27;;;54642:38;;54639:58;;;54693:1;54690;54683:12;54708:1165;54885:2;54874:9;54867:21;54938:4;54930:6;54925:2;54914:9;54910:18;54897:46;54848:4;54977:3;54966:9;54962:19;55000:1;55021:2;55017;55010:14;55071:4;55063:6;55059:17;55046:31;55086:30;55110:5;55086:30;:::i;:::-;55147:10;55136:22;55125:34;;;55219:3;55207:16;;55194:30;;55275:14;55271:27;;;-1:-1:-1;;55267:41:82;55243:66;;55233:96;;55324:2;55320;55313:14;55233:96;55353:31;;;;55407:21;;-1:-1:-1;;;;;55440:30:82;;55437:52;;;55484:2;55480;55473:14;55437:52;55533:6;55517:14;55513:27;55505:6;55501:40;55498:62;;;55555:2;55551;55544:14;55498:62;55579:6;55569:16;;55622:2;55616:3;55605:9;55601:19;55594:31;55648:72;55715:3;55704:9;55700:19;55692:6;55687:2;55678:7;55674:16;55648:72;:::i;:::-;55634:86;;;55749:35;55779:3;55771:6;55767:16;55749:35;:::i;:::-;6924:10;6913:22;;55825:18;;;6901:35;55793:51;6848:94;55878:274;56076:2;56061:18;;56088:58;56065:9;56128:6;56088:58;:::i;56157:278::-;56196:7;-1:-1:-1;;;;;56281:2:82;56278:1;56274:10;56311:2;56308:1;56304:10;56367:3;56363:2;56359:12;56354:3;56351:21;56344:3;56337:11;56330:19;56326:47;56323:73;;;56376:18;;:::i;:::-;56416:13;;56157:278;-1:-1:-1;;;;56157:278:82:o;56440:168::-;56480:7;56546:1;56542;56538:6;56534:14;56531:1;56528:21;56523:1;56516:9;56509:17;56505:45;56502:71;;;56553:18;;:::i;:::-;-1:-1:-1;56593:9:82;;56440:168::o;56613:509::-;56917:3;56902:19;;56930:58;56906:9;56970:6;56930:58;:::i;:::-;57043:10;57034:6;57028:13;57024:30;57019:2;57008:9;57004:18;56997:58;57109:4;57101:6;57097:17;57091:24;57086:2;57075:9;57071:18;57064:52;56613:509;;;;;:::o;60733:197::-;60771:3;60799:6;60840:2;60833:5;60829:14;60867:2;60858:7;60855:15;60852:41;;;60873:18;;:::i;:::-;60922:1;60909:15;;60733:197;-1:-1:-1;;;60733:197:82:o",
    linkReferences: {},
    immutableReferences: {
      "58120": [
        { start: 738, length: 32 },
        { start: 1585, length: 32 },
        { start: 13287, length: 32 },
      ],
      "7401": [
        { start: 1300, length: 32 },
        { start: 5021, length: 32 },
        { start: 10144, length: 32 },
        { start: 10521, length: 32 },
        { start: 11091, length: 32 },
      ],
      "7404": [
        { start: 1261, length: 32 },
        { start: 11886, length: 32 },
        { start: 12336, length: 32 },
      ],
      "7407": [
        { start: 1203, length: 32 },
        { start: 5753, length: 32 },
        { start: 11416, length: 32 },
      ],
      "7410": [
        { start: 1530, length: 32 },
        { start: 10934, length: 32 },
      ],
    },
  },
  methodIdentifiers: {
    "TASK_CHALLENGE_WINDOW_BLOCK()": "f63c5bab",
    "TASK_RESPONSE_WINDOW_BLOCK()": "1ad43189",
    "aggregator()": "245a7bfc",
    "allTaskHashes(uint32)": "2d89f6fc",
    "allTaskResponses(uint32)": "2cb223d5",
    "blsApkRegistry()": "5df45946",
    "challengeInstances(uint32,uint256)": "67d6be44",
    "checkSignatures(bytes32,bytes,uint32,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]))":
      "6efb4636",
    "confirmChallenge((uint256[5],uint32,bytes,uint32),(uint32,uint256),(uint32,bytes32),(uint256,uint256)[])":
      "480bab6b",
    "createNewTask(uint256[5],uint32,bytes)": "0627721e",
    "delegation()": "df5cf723",
    "generator()": "7afa1eed",
    "getBatchOperatorFromId(address,bytes32[])": "4d2b57fe",
    "getBatchOperatorId(address,address[])": "31b36bd9",
    "getCheckSignaturesIndices(address,uint32,bytes,bytes32[])": "4f739f74",
    "getOperatorState(address,bytes,uint32)": "3563b0d1",
    "getOperatorState(address,bytes32,uint32)": "cefdc1d4",
    "getQuorumBitmapsAtBlockNumber(address,bytes32[],uint32)": "5c155662",
    "getTaskResponseWindowBlock()": "f5c9899d",
    "initialize(address,address,address,address,address)": "1459457a",
    "latestTaskNum()": "8b00ce7c",
    "owner()": "8da5cb5b",
    "pause(uint256)": "136439dd",
    "pauseAll()": "595c6a67",
    "paused()": "5c975abb",
    "paused(uint8)": "5ac86ab7",
    "pauserRegistry()": "886f1195",
    "proveResultAccurate(uint32,uint256[],bytes)": "58a7cd26",
    "raiseChallenger((uint256[5],uint32,bytes,uint32),(uint32,uint256),(uint32,bytes32))":
      "162d8a3f",
    "registryCoordinator()": "6d14a987",
    "renounceOwnership()": "715018a6",
    "respondToTask((uint256[5],uint32,bytes,uint32),(uint32,uint256),(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]))":
      "bdeea6cc",
    "setPauserRegistry(address)": "10d67a2f",
    "setStaleStakesForbidden(bool)": "416c7e5e",
    "stakeRegistry()": "68304835",
    "staleStakesForbidden()": "b98d0908",
    "taskNumber()": "72d18e8d",
    "transferOwnership(address)": "f2fde38b",
    "trySignatureAndApkVerification(bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256))":
      "171f1d5b",
    "unpause(uint256)": "fabc1cbc",
  },
  rawMetadata:
    '{"compiler":{"version":"0.8.12+commit.f00d7308"},"language":"Solidity","output":{"abi":[{"inputs":[{"internalType":"contract IRegistryCoordinator","name":"_registryCoordinator","type":"address"},{"internalType":"uint32","name":"_taskResponseWindowBlock","type":"uint32"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint8","name":"version","type":"uint8"}],"name":"Initialized","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint32","name":"taskIndex","type":"uint32"},{"components":[{"internalType":"uint256[5]","name":"inputs","type":"uint256[5]"},{"internalType":"uint32","name":"taskCreatedBlock","type":"uint32"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"uint32","name":"quorumThresholdPercentage","type":"uint32"}],"indexed":false,"internalType":"struct IOmronTaskManager.Task","name":"task","type":"tuple"}],"name":"NewTaskCreated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":false,"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"Paused","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"contract IPauserRegistry","name":"pauserRegistry","type":"address"},{"indexed":false,"internalType":"contract IPauserRegistry","name":"newPauserRegistry","type":"address"}],"name":"PauserRegistrySet","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bool","name":"value","type":"bool"}],"name":"StaleStakesForbiddenUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint32","name":"taskIndex","type":"uint32"}],"name":"TaskChallenged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint32","name":"taskIndex","type":"uint32"},{"indexed":true,"internalType":"address","name":"challenger","type":"address"}],"name":"TaskChallengedSuccessfully","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint32","name":"taskIndex","type":"uint32"},{"indexed":true,"internalType":"address","name":"prover","type":"address"}],"name":"TaskChallengedUnsuccessfully","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint32","name":"taskIndex","type":"uint32"}],"name":"TaskCompleted","type":"event"},{"anonymous":false,"inputs":[{"components":[{"internalType":"uint32","name":"referenceTaskIndex","type":"uint32"},{"internalType":"uint256","name":"output","type":"uint256"}],"indexed":false,"internalType":"struct IOmronTaskManager.TaskResponse","name":"taskResponse","type":"tuple"},{"components":[{"internalType":"uint32","name":"taskResponsedBlock","type":"uint32"},{"internalType":"bytes32","name":"hashOfNonSigners","type":"bytes32"}],"indexed":false,"internalType":"struct IOmronTaskManager.TaskResponseMetadata","name":"taskResponseMetadata","type":"tuple"}],"name":"TaskResponded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":false,"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"Unpaused","type":"event"},{"inputs":[],"name":"TASK_CHALLENGE_WINDOW_BLOCK","outputs":[{"internalType":"uint32","name":"","type":"uint32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"TASK_RESPONSE_WINDOW_BLOCK","outputs":[{"internalType":"uint32","name":"","type":"uint32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"aggregator","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint32","name":"","type":"uint32"}],"name":"allTaskHashes","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint32","name":"","type":"uint32"}],"name":"allTaskResponses","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"blsApkRegistry","outputs":[{"internalType":"contract IBLSApkRegistry","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint32","name":"id","type":"uint32"},{"internalType":"uint256","name":"index","type":"uint256"}],"name":"challengeInstances","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"msgHash","type":"bytes32"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"uint32","name":"referenceBlockNumber","type":"uint32"},{"components":[{"internalType":"uint32[]","name":"nonSignerQuorumBitmapIndices","type":"uint32[]"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point[]","name":"nonSignerPubkeys","type":"tuple[]"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point[]","name":"quorumApks","type":"tuple[]"},{"components":[{"internalType":"uint256[2]","name":"X","type":"uint256[2]"},{"internalType":"uint256[2]","name":"Y","type":"uint256[2]"}],"internalType":"struct BN254.G2Point","name":"apkG2","type":"tuple"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"sigma","type":"tuple"},{"internalType":"uint32[]","name":"quorumApkIndices","type":"uint32[]"},{"internalType":"uint32[]","name":"totalStakeIndices","type":"uint32[]"},{"internalType":"uint32[][]","name":"nonSignerStakeIndices","type":"uint32[][]"}],"internalType":"struct IBLSSignatureChecker.NonSignerStakesAndSignature","name":"params","type":"tuple"}],"name":"checkSignatures","outputs":[{"components":[{"internalType":"uint96[]","name":"signedStakeForQuorum","type":"uint96[]"},{"internalType":"uint96[]","name":"totalStakeForQuorum","type":"uint96[]"}],"internalType":"struct IBLSSignatureChecker.QuorumStakeTotals","name":"","type":"tuple"},{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"components":[{"internalType":"uint256[5]","name":"inputs","type":"uint256[5]"},{"internalType":"uint32","name":"taskCreatedBlock","type":"uint32"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"uint32","name":"quorumThresholdPercentage","type":"uint32"}],"internalType":"struct IOmronTaskManager.Task","name":"task","type":"tuple"},{"components":[{"internalType":"uint32","name":"referenceTaskIndex","type":"uint32"},{"internalType":"uint256","name":"output","type":"uint256"}],"internalType":"struct IOmronTaskManager.TaskResponse","name":"taskResponse","type":"tuple"},{"components":[{"internalType":"uint32","name":"taskResponsedBlock","type":"uint32"},{"internalType":"bytes32","name":"hashOfNonSigners","type":"bytes32"}],"internalType":"struct IOmronTaskManager.TaskResponseMetadata","name":"taskResponseMetadata","type":"tuple"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point[]","name":"pubkeysOfNonSigningOperators","type":"tuple[]"}],"name":"confirmChallenge","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256[5]","name":"inputs","type":"uint256[5]"},{"internalType":"uint32","name":"quorumThresholdPercentage","type":"uint32"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"}],"name":"createNewTask","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"delegation","outputs":[{"internalType":"contract IDelegationManager","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"generator","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"contract IRegistryCoordinator","name":"registryCoordinator","type":"address"},{"internalType":"bytes32[]","name":"operatorIds","type":"bytes32[]"}],"name":"getBatchOperatorFromId","outputs":[{"internalType":"address[]","name":"operators","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"contract IRegistryCoordinator","name":"registryCoordinator","type":"address"},{"internalType":"address[]","name":"operators","type":"address[]"}],"name":"getBatchOperatorId","outputs":[{"internalType":"bytes32[]","name":"operatorIds","type":"bytes32[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"contract IRegistryCoordinator","name":"registryCoordinator","type":"address"},{"internalType":"uint32","name":"referenceBlockNumber","type":"uint32"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"bytes32[]","name":"nonSignerOperatorIds","type":"bytes32[]"}],"name":"getCheckSignaturesIndices","outputs":[{"components":[{"internalType":"uint32[]","name":"nonSignerQuorumBitmapIndices","type":"uint32[]"},{"internalType":"uint32[]","name":"quorumApkIndices","type":"uint32[]"},{"internalType":"uint32[]","name":"totalStakeIndices","type":"uint32[]"},{"internalType":"uint32[][]","name":"nonSignerStakeIndices","type":"uint32[][]"}],"internalType":"struct OperatorStateRetriever.CheckSignaturesIndices","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"contract IRegistryCoordinator","name":"registryCoordinator","type":"address"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"uint32","name":"blockNumber","type":"uint32"}],"name":"getOperatorState","outputs":[{"components":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bytes32","name":"operatorId","type":"bytes32"},{"internalType":"uint96","name":"stake","type":"uint96"}],"internalType":"struct OperatorStateRetriever.Operator[][]","name":"","type":"tuple[][]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"contract IRegistryCoordinator","name":"registryCoordinator","type":"address"},{"internalType":"bytes32","name":"operatorId","type":"bytes32"},{"internalType":"uint32","name":"blockNumber","type":"uint32"}],"name":"getOperatorState","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"components":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bytes32","name":"operatorId","type":"bytes32"},{"internalType":"uint96","name":"stake","type":"uint96"}],"internalType":"struct OperatorStateRetriever.Operator[][]","name":"","type":"tuple[][]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"contract IRegistryCoordinator","name":"registryCoordinator","type":"address"},{"internalType":"bytes32[]","name":"operatorIds","type":"bytes32[]"},{"internalType":"uint32","name":"blockNumber","type":"uint32"}],"name":"getQuorumBitmapsAtBlockNumber","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getTaskResponseWindowBlock","outputs":[{"internalType":"uint32","name":"","type":"uint32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"contract IPauserRegistry","name":"_pauserRegistry","type":"address"},{"internalType":"address","name":"initialOwner","type":"address"},{"internalType":"address","name":"_aggregator","type":"address"},{"internalType":"address","name":"_generator","type":"address"},{"internalType":"address","name":"_inferenceDB","type":"address"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"latestTaskNum","outputs":[{"internalType":"uint32","name":"","type":"uint32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"pause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"pauseAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint8","name":"index","type":"uint8"}],"name":"paused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"paused","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pauserRegistry","outputs":[{"internalType":"contract IPauserRegistry","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint32","name":"taskId","type":"uint32"},{"internalType":"uint256[]","name":"instances","type":"uint256[]"},{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"proveResultAccurate","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"components":[{"internalType":"uint256[5]","name":"inputs","type":"uint256[5]"},{"internalType":"uint32","name":"taskCreatedBlock","type":"uint32"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"uint32","name":"quorumThresholdPercentage","type":"uint32"}],"internalType":"struct IOmronTaskManager.Task","name":"task","type":"tuple"},{"components":[{"internalType":"uint32","name":"referenceTaskIndex","type":"uint32"},{"internalType":"uint256","name":"output","type":"uint256"}],"internalType":"struct IOmronTaskManager.TaskResponse","name":"taskResponse","type":"tuple"},{"components":[{"internalType":"uint32","name":"taskResponsedBlock","type":"uint32"},{"internalType":"bytes32","name":"hashOfNonSigners","type":"bytes32"}],"internalType":"struct IOmronTaskManager.TaskResponseMetadata","name":"taskResponseMetadata","type":"tuple"}],"name":"raiseChallenger","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"registryCoordinator","outputs":[{"internalType":"contract IRegistryCoordinator","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"components":[{"internalType":"uint256[5]","name":"inputs","type":"uint256[5]"},{"internalType":"uint32","name":"taskCreatedBlock","type":"uint32"},{"internalType":"bytes","name":"quorumNumbers","type":"bytes"},{"internalType":"uint32","name":"quorumThresholdPercentage","type":"uint32"}],"internalType":"struct IOmronTaskManager.Task","name":"task","type":"tuple"},{"components":[{"internalType":"uint32","name":"referenceTaskIndex","type":"uint32"},{"internalType":"uint256","name":"output","type":"uint256"}],"internalType":"struct IOmronTaskManager.TaskResponse","name":"taskResponse","type":"tuple"},{"components":[{"internalType":"uint32[]","name":"nonSignerQuorumBitmapIndices","type":"uint32[]"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point[]","name":"nonSignerPubkeys","type":"tuple[]"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point[]","name":"quorumApks","type":"tuple[]"},{"components":[{"internalType":"uint256[2]","name":"X","type":"uint256[2]"},{"internalType":"uint256[2]","name":"Y","type":"uint256[2]"}],"internalType":"struct BN254.G2Point","name":"apkG2","type":"tuple"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"sigma","type":"tuple"},{"internalType":"uint32[]","name":"quorumApkIndices","type":"uint32[]"},{"internalType":"uint32[]","name":"totalStakeIndices","type":"uint32[]"},{"internalType":"uint32[][]","name":"nonSignerStakeIndices","type":"uint32[][]"}],"internalType":"struct IBLSSignatureChecker.NonSignerStakesAndSignature","name":"nonSignerStakesAndSignature","type":"tuple"}],"name":"respondToTask","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"contract IPauserRegistry","name":"newPauserRegistry","type":"address"}],"name":"setPauserRegistry","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bool","name":"value","type":"bool"}],"name":"setStaleStakesForbidden","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"stakeRegistry","outputs":[{"internalType":"contract IStakeRegistry","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"staleStakesForbidden","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"taskNumber","outputs":[{"internalType":"uint32","name":"","type":"uint32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"msgHash","type":"bytes32"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"apk","type":"tuple"},{"components":[{"internalType":"uint256[2]","name":"X","type":"uint256[2]"},{"internalType":"uint256[2]","name":"Y","type":"uint256[2]"}],"internalType":"struct BN254.G2Point","name":"apkG2","type":"tuple"},{"components":[{"internalType":"uint256","name":"X","type":"uint256"},{"internalType":"uint256","name":"Y","type":"uint256"}],"internalType":"struct BN254.G1Point","name":"sigma","type":"tuple"}],"name":"trySignatureAndApkVerification","outputs":[{"internalType":"bool","name":"pairingSuccessful","type":"bool"},{"internalType":"bool","name":"siganatureIsValid","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"newPausedStatus","type":"uint256"}],"name":"unpause","outputs":[],"stateMutability":"nonpayable","type":"function"}],"devdoc":{"kind":"dev","methods":{"checkSignatures(bytes32,bytes,uint32,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]))":{"details":"Before signature verification, the function verifies operator stake information.  This includes ensuring that the provided `referenceBlockNumber` is correct, i.e., ensure that the stake returned from the specified block number is recent enough and that the stake is either the most recent update for the total stake (of the operator) or latest before the referenceBlockNumber.NOTE: Be careful to ensure `msgHash` is collision-resistant! This method does not hash  `msgHash` in any way, so if an attacker is able to pass in an arbitrary value, they may be able to tamper with signature verification.","params":{"msgHash":"is the hash being signed","params":"is the struct containing information on nonsigners, stakes, quorum apks, and the aggregate signature","quorumNumbers":"is the bytes array of quorum numbers that are being signed for","referenceBlockNumber":"is the block number at which the stake information is being verified"},"returns":{"_0":"quorumStakeTotals is the struct containing the total and signed stake for each quorum","_1":"signatoryRecordHash is the hash of the signatory record, which is used for fraud proofs"}},"getBatchOperatorFromId(address,bytes32[])":{"details":"if an operator is not registered, the operator address will be 0","params":{"operators":"is the array of operatorIds to get corresponding operator addresses for","registryCoordinator":"is the AVS registry coordinator to fetch the operator information from"}},"getBatchOperatorId(address,address[])":{"details":"if an operator is not registered, the operatorId will be 0","params":{"operators":"is the array of operator address to get corresponding operatorIds for","registryCoordinator":"is the AVS registry coordinator to fetch the operator information from"}},"getCheckSignaturesIndices(address,uint32,bytes,bytes32[])":{"params":{"nonSignerOperatorIds":"are the ids of the nonsigning operators","quorumNumbers":"are the ids of the quorums to get the operator state for","referenceBlockNumber":"is the block number to get the indices for","registryCoordinator":"is the registry coordinator to fetch the AVS registry information from"},"returns":{"_0":"1) the indices of the quorumBitmaps for each of the operators in the @param nonSignerOperatorIds array at the given blocknumber         2) the indices of the total stakes entries for the given quorums at the given blocknumber         3) the indices of the stakes of each of the nonsigners in each of the quorums they were a             part of (for each nonsigner, an array of length the number of quorums they were a part of            that are also part of the provided quorumNumbers) at the given blocknumber         4) the indices of the quorum apks for each of the provided quorums at the given blocknumber"}},"getOperatorState(address,bytes,uint32)":{"params":{"blockNumber":"is the block number to get the operator state for","quorumNumbers":"are the ids of the quorums to get the operator state for","registryCoordinator":"is the registry coordinator to fetch the AVS registry information from"},"returns":{"_0":"2d array of Operators. For each quorum, an ordered list of Operators"}},"getOperatorState(address,bytes32,uint32)":{"params":{"blockNumber":"is the block number to get the operator state for","operatorId":"the id of the operator to fetch the quorums lists ","registryCoordinator":"is the registry coordinator to fetch the AVS registry information from"},"returns":{"_0":"1) the quorumBitmap of the operator at the given blockNumber         2) 2d array of Operator structs. For each quorum the provided operator             was a part of at `blockNumber`, an ordered list of operators."}},"getQuorumBitmapsAtBlockNumber(address,bytes32[],uint32)":{"params":{"blockNumber":"is the block number to get the quorumBitmaps for","operatorIds":"are the ids of the operators to get the quorumBitmaps for","registryCoordinator":"is the AVS registry coordinator to fetch the operator information from"}},"owner()":{"details":"Returns the address of the current owner."},"pause(uint256)":{"details":"This function can only pause functionality, and thus cannot \'unflip\' any bit in `_paused` from 1 to 0.","params":{"newPausedStatus":"represents the new value for `_paused` to take, which means it may flip several bits at once."}},"renounceOwnership()":{"details":"Leaves the contract without owner. It will not be possible to call `onlyOwner` functions anymore. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby removing any functionality that is only available to the owner."},"setStaleStakesForbidden(bool)":{"params":{"value":"to toggle staleStakesForbidden"}},"transferOwnership(address)":{"details":"Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner."},"trySignatureAndApkVerification(bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256))":{"params":{"apk":"is the claimed G1 public key","apkG2":"is provided G2 public key","msgHash":"is the hash being signed","sigma":"is the G1 point signature"},"returns":{"pairingSuccessful":"is true if the pairing precompile call was successful","siganatureIsValid":"is true if the signature is valid"}},"unpause(uint256)":{"details":"This function can only unpause functionality, and thus cannot \'flip\' any bit in `_paused` from 0 to 1.","params":{"newPausedStatus":"represents the new value for `_paused` to take, which means it may flip several bits at once."}}},"version":1},"userdoc":{"events":{"Paused(address,uint256)":{"notice":"Emitted when the pause is triggered by `account`, and changed to `newPausedStatus`."},"PauserRegistrySet(address,address)":{"notice":"Emitted when the `pauserRegistry` is set to `newPauserRegistry`."},"StaleStakesForbiddenUpdate(bool)":{"notice":"Emitted when `staleStakesForbiddenUpdate` is set"},"Unpaused(address,uint256)":{"notice":"Emitted when the pause is lifted by `account`, and changed to `newPausedStatus`."}},"kind":"user","methods":{"checkSignatures(bytes32,bytes,uint32,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]))":{"notice":"This function is called by disperser when it has aggregated all the signatures of the operators that are part of the quorum for a particular taskNumber and is asserting them into onchain. The function checks that the claim for aggregated signatures are valid. The thesis of this procedure entails: - getting the aggregated pubkey of all registered nodes at the time of pre-commit by the disperser (represented by apk in the parameters), - subtracting the pubkeys of all the signers not in the quorum (nonSignerPubkeys) and storing  the output in apk to get aggregated pubkey of all operators that are part of quorum. - use this aggregated pubkey to verify the aggregated signature under BLS scheme. "},"getBatchOperatorFromId(address,bytes32[])":{"notice":"This function returns the operator addresses for each of the operators in the operatorIds array"},"getBatchOperatorId(address,address[])":{"notice":"This function returns the operatorIds for each of the operators in the operators array"},"getCheckSignaturesIndices(address,uint32,bytes,bytes32[])":{"notice":"this is called by the AVS operator to get the relevant indices for the checkSignatures function if they are not running an indexer    "},"getOperatorState(address,bytes,uint32)":{"notice":"returns the ordered list of operators (id and stake) for each quorum. The AVS coordinator  may call this function directly to get the operator state for a given block number"},"getOperatorState(address,bytes32,uint32)":{"notice":"This function is intended to to be called by AVS operators every time a new task is created (i.e.) the AVS coordinator makes a request to AVS operators. Since all of the crucial information is kept onchain,  operators don\'t need to run indexers to fetch the data."},"getQuorumBitmapsAtBlockNumber(address,bytes32[],uint32)":{"notice":"this function returns the quorumBitmaps for each of the operators in the operatorIds array at the given blocknumber"},"getTaskResponseWindowBlock()":{"notice":"Returns the TASK_RESPONSE_WINDOW_BLOCK"},"pause(uint256)":{"notice":"This function is used to pause an EigenLayer contract\'s functionality. It is permissioned to the `pauser` address, which is expected to be a low threshold multisig."},"pauseAll()":{"notice":"Alias for `pause(type(uint256).max)`."},"paused()":{"notice":"Returns the current paused status as a uint256."},"paused(uint8)":{"notice":"Returns \'true\' if the `indexed`th bit of `_paused` is 1, and \'false\' otherwise"},"pauserRegistry()":{"notice":"Address of the `PauserRegistry` contract that this contract defers to for determining access control (for pausing)."},"setPauserRegistry(address)":{"notice":"Allows the unpauser to set a new pauser registry"},"setStaleStakesForbidden(bool)":{"notice":"RegistryCoordinator owner can either enforce or not that operator stakes are staler than the delegation.minWithdrawalDelayBlocks() window."},"staleStakesForbidden()":{"notice":"If true, check the staleness of the operator stakes and that its within the delegation withdrawalDelayBlocks window."},"taskNumber()":{"notice":"Returns the current \'taskNumber\' for the middleware"},"trySignatureAndApkVerification(bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256))":{"notice":"trySignatureAndApkVerification verifies a BLS aggregate signature and the veracity of a calculated G1 Public key"},"unpause(uint256)":{"notice":"This function is used to unpause an EigenLayer contract\'s functionality. It is permissioned to the `unpauser` address, which is expected to be a high threshold multisig or governance contract."}},"version":1}},"settings":{"compilationTarget":{"src/OmronTaskManager.sol":"OmronTaskManager"},"evmVersion":"london","libraries":{},"metadata":{"bytecodeHash":"ipfs"},"optimizer":{"enabled":true,"runs":200},"remappings":[":@eigenlayer-middleware/=lib/eigenlayer-middleware/",":@eigenlayer-scripts/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/script/",":@eigenlayer/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/",":@omron/=src/",":@openzeppelin-upgrades/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/",":@openzeppelin/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/",":ds-test/=lib/eigenlayer-middleware/lib/ds-test/src/",":eigenlayer-contracts/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/",":eigenlayer-middleware/=lib/eigenlayer-middleware/",":forge-std/=lib/forge-std/src/",":openzeppelin-contracts-upgradeable/=lib/eigenlayer-middleware/lib/openzeppelin-contracts-upgradeable/",":openzeppelin-contracts/=lib/eigenlayer-middleware/lib/openzeppelin-contracts/"]},"sources":{"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol":{"keccak256":"0x247c62047745915c0af6b955470a72d1696ebad4352d7d3011aef1a2463cd888","license":"MIT","urls":["bzz-raw://d7fc8396619de513c96b6e00301b88dd790e83542aab918425633a5f7297a15a","dweb:/ipfs/QmXbP4kiZyp7guuS7xe8KaybnwkRPGrBc2Kbi3vhcTfpxb"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol":{"keccak256":"0x0203dcadc5737d9ef2c211d6fa15d18ebc3b30dfa51903b64870b01a062b0b4e","license":"MIT","urls":["bzz-raw://6eb2fd1e9894dbe778f4b8131adecebe570689e63cf892f4e21257bfe1252497","dweb:/ipfs/QmXgUGNfZvrn6N2miv3nooSs7Jm34A41qz94fu2GtDFcx8"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/utils/AddressUpgradeable.sol":{"keccak256":"0x611aa3f23e59cfdd1863c536776407b3e33d695152a266fa7cfb34440a29a8a3","license":"MIT","urls":["bzz-raw://9b4b2110b7f2b3eb32951bc08046fa90feccffa594e1176cb91cdfb0e94726b4","dweb:/ipfs/QmSxLwYjicf9zWFuieRc8WQwE4FisA1Um5jp1iSa731TGt"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol":{"keccak256":"0x963ea7f0b48b032eef72fe3a7582edf78408d6f834115b9feadd673a4d5bd149","license":"MIT","urls":["bzz-raw://d6520943ea55fdf5f0bafb39ed909f64de17051bc954ff3e88c9e5621412c79c","dweb:/ipfs/QmWZ4rAKTQbNG2HxGs46AcTXShsVytKeLs7CUCdCSv5N7a"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/access/Ownable.sol":{"keccak256":"0xa94b34880e3c1b0b931662cb1c09e5dfa6662f31cba80e07c5ee71cd135c9673","license":"MIT","urls":["bzz-raw://40fb1b5102468f783961d0af743f91b9980cf66b50d1d12009f6bb1869cea4d2","dweb:/ipfs/QmYqEbJML4jB1GHbzD4cUZDtJg5wVwNm3vDJq1GbyDus8y"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/interfaces/IERC1271.sol":{"keccak256":"0x0705a4b1b86d7b0bd8432118f226ba139c44b9dcaba0a6eafba2dd7d0639c544","license":"MIT","urls":["bzz-raw://c45b821ef9e882e57c256697a152e108f0f2ad6997609af8904cae99c9bd422e","dweb:/ipfs/QmRKCJW6jjzR5UYZcLpGnhEJ75UVbH6EHkEa49sWx2SKng"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/proxy/beacon/IBeacon.sol":{"keccak256":"0xd50a3421ac379ccb1be435fa646d66a65c986b4924f0849839f08692f39dde61","license":"MIT","urls":["bzz-raw://ada1e030c0231db8d143b44ce92b4d1158eedb087880cad6d8cc7bd7ebe7b354","dweb:/ipfs/QmWZ2NHZweRpz1U9GF6R1h65ri76dnX7fNxLBeM2t5N5Ce"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol":{"keccak256":"0x9750c6b834f7b43000631af5cc30001c5f547b3ceb3635488f140f60e897ea6b","license":"MIT","urls":["bzz-raw://5a7d5b1ef5d8d5889ad2ed89d8619c09383b80b72ab226e0fe7bde1636481e34","dweb:/ipfs/QmebXWgtEfumQGBdVeM6c71McLixYXQP5Bk6kKXuoY4Bmr"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/Address.sol":{"keccak256":"0xd6153ce99bcdcce22b124f755e72553295be6abcd63804cfdffceb188b8bef10","license":"MIT","urls":["bzz-raw://35c47bece3c03caaa07fab37dd2bb3413bfbca20db7bd9895024390e0a469487","dweb:/ipfs/QmPGWT2x3QHcKxqe6gRmAkdakhbaRgx3DLzcakHz5M4eXG"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/Context.sol":{"keccak256":"0xe2e337e6dde9ef6b680e07338c493ebea1b5fd09b43424112868e9cc1706bca7","license":"MIT","urls":["bzz-raw://6df0ddf21ce9f58271bdfaa85cde98b200ef242a05a3f85c2bc10a8294800a92","dweb:/ipfs/QmRK2Y5Yc6BK7tGKkgsgn3aJEQGi5aakeSPZvS65PV8Xp3"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/Strings.sol":{"keccak256":"0xaf159a8b1923ad2a26d516089bceca9bdeaeacd04be50983ea00ba63070f08a3","license":"MIT","urls":["bzz-raw://6f2cf1c531122bc7ca96b8c8db6a60deae60441e5223065e792553d4849b5638","dweb:/ipfs/QmPBdJmBBABMDCfyDjCbdxgiqRavgiSL88SYPGibgbPas9"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol":{"keccak256":"0x84ac2d2f343df1e683da7a12bbcf70db542a7a7a0cea90a5d70fcb5e5d035481","license":"MIT","urls":["bzz-raw://73ae8e0c6f975052973265113d762629002ce33987b1933c2a378667e2816f2f","dweb:/ipfs/QmQAootkVfoe4PLaYbT4Xob2dJRm3bZfbCffEHRbCYXNPF"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/cryptography/draft-EIP712.sol":{"keccak256":"0x6688fad58b9ec0286d40fa957152e575d5d8bd4c3aa80985efdb11b44f776ae7","license":"MIT","urls":["bzz-raw://8bc00ab7f133cdaafd212a5cc6a16c8d37319721105d130c8e5af0c4e8f170ba","dweb:/ipfs/QmVmf6LVMfFiEkvKYLzSv3bGHzymEW93AcUuFrNUdY3NtT"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IBeaconChainOracle.sol":{"keccak256":"0x0fef07aa6179c77198f1514e12e628aa1c876e04f9c181ec853a322179e5be00","license":"BUSL-1.1","urls":["bzz-raw://51438325876cc2d4c77f58488a7e27b488015d1b663c50be6a5cafbd73b9c983","dweb:/ipfs/QmViCuGoYZzi6wtXA8PPKigqVv3KMuNxEVQ1Td9dGqjL18"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol":{"keccak256":"0xe1626408822f552043f945d9aea18c5cbf878ef160da55e6f98706ed3a2acc07","license":"BUSL-1.1","urls":["bzz-raw://426f6dddc040f2040f48dd4236c4201a3c978b4417ec3b4bd1004f8a48b29aaa","dweb:/ipfs/QmWgY46nZiw1KQYNYMrJDTz7S9Y4KhyUoM9zVD92Mkf51S"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IETHPOSDeposit.sol":{"keccak256":"0x2e60e5f4b0da0a0a4e2a07c63141120998559970c21deac743ea0c64a60a880c","license":"CC0-1.0","urls":["bzz-raw://e635c346bde5b7ade9bcf35bc733081520cb86015be4fbc6e761e6e9482c4c91","dweb:/ipfs/QmRoeazEnbFn5SPSWAkoFK2gSN9DMp3hJAnrLWuL2sKutz"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IEigenPod.sol":{"keccak256":"0xb50c36ad96b6679bb80fd8331f949cbfbcba0f529026e1421a4d2bae64396eba","license":"BUSL-1.1","urls":["bzz-raw://5719181d780120f1e688c0da276992a8caf185815917f453b3550537c31ed4cc","dweb:/ipfs/QmYprRC5ZEXhz3zAUND5E8Xjn6s5TL8ZF8QbnndVq7aVPR"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IEigenPodManager.sol":{"keccak256":"0xd8a64dbed03d3a5cdbefe1af75968f2dde07f973749c2ef5197bf7187c3e448c","license":"BUSL-1.1","urls":["bzz-raw://27ccc7c1fd9352e9f9b357c9063d255dc0ed9583f43db09f786ac7497d7846b8","dweb:/ipfs/QmeJzuJkE9m2NUNwZSp4tGZEZmih1LeucePup8hzMVDRbG"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IPausable.sol":{"keccak256":"0x98cffc894842947377e24c1d375813a1120dd73a84c29782ab68404e109cb34f","license":"BUSL-1.1","urls":["bzz-raw://b3474f6c350ceaee57cbdfb08fb48835d0c6e81ae8ebfbb9667899584a139324","dweb:/ipfs/QmWELKtksdtWxQbqAccd8yGyhKqrgPZXTADKR7BuT27Zg5"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol":{"keccak256":"0x9de8dd682bc0d812bbd6583c0231cbf35448d5eff58b74a93efa64cb9a768c49","license":"BUSL-1.1","urls":["bzz-raw://c00d6c675b9c72b092d287fe85fd37782588df32b8eb59ab4c7db7a86be25e7d","dweb:/ipfs/QmeYokY3HhAdbBaCPdHg3PgQEdRCDFEJy3Wf7VtgHBkQSx"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IPaymentCoordinator.sol":{"keccak256":"0xf08a2adc0ec07ec0fd4711adab5f8a72fbb0dcafb62ee032b983b51b167348e6","license":"BUSL-1.1","urls":["bzz-raw://702dda0a24d175188d87a9d924ce8244ad626b01aba547e9a4368811991e8950","dweb:/ipfs/QmUgsi4BUcqXvtcLTQvk3wfVHCaCGwKc2gNixx4bz58csY"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol":{"keccak256":"0x5e52482a31d94401a8502f3014c4aada1142b4450fc0596dff8e1866a85fe092","license":"BUSL-1.1","urls":["bzz-raw://17dc326c9361bc1453379f26545963557b2883b0c88bc07d4477e04dbcc0cc8c","dweb:/ipfs/QmZXT7A816W5JH2ymirE2ETaJttqztFCsEL22AV8oEfCK9"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/ISlasher.sol":{"keccak256":"0x45dfaa2cfdde87f48a6ee38bb6fb739847aef7cf3f6137bdcd8c8a330559ec79","license":"BUSL-1.1","urls":["bzz-raw://1b7f6bd75b42fcaa91ceb7140cb2c41926a1fe6ee2d3161e4fe6186b181ba232","dweb:/ipfs/QmZjbdKiSs33C9i3GDc3sdD39Pz4YPkDoKftowoUF4kHmY"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol":{"keccak256":"0xc530c6a944b70051fd0dac0222de9a4b5baadeaf94ad194daac6ad8d2ace7420","license":"BUSL-1.1","urls":["bzz-raw://3767df0364ce835b52e786d2851431eb9223fe4747602107505477e162231d73","dweb:/ipfs/QmZkH5bKUygQrJomndNaQqkefVRW4rRefCa8HPJ5HMczxJ"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol":{"keccak256":"0xccd308b0996295c92058b70c3b83563c009c074cb6815329c5f35e1b1a0571f4","license":"BUSL-1.1","urls":["bzz-raw://cd1050445ff7aeb588b3b037aab53e2d92c3abd4620e94dfc95aca870e71b821","dweb:/ipfs/QmUC96Ctwn3KQr6VSqHPpAVJ5qEUSVnugaxiZ8gfXygW92"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/BeaconChainProofs.sol":{"keccak256":"0x70d89b05c1c5f47b74a07fbb5a2c05e606fed494e749ea98a9915b7be73df377","license":"BUSL-1.1","urls":["bzz-raw://db1d3bfaee69aef53c8b12b492a17584e6d1ac94610cb8b38aad33e1cdd81af7","dweb:/ipfs/QmfVsMTj1hcf9fMEm5RzvtcBN4dMcAKFBgUUDsNDr5XFpq"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/EIP1271SignatureUtils.sol":{"keccak256":"0xe92d584c47c5828e026a8082af3da38a853e3942c4da7deb705d6470a41afab3","license":"BUSL-1.1","urls":["bzz-raw://1c436c578781fd7d3dffdb24e906819422819f5e9a71d39ee63166a3d5cb3373","dweb:/ipfs/QmP7bJhYqLpwqk2Xq4tqDCUMi2nFAhxxW3Pz36ctE1sbdD"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/Endian.sol":{"keccak256":"0xf3b72653ba2567a978d4612703fa5f71c5fcd015d8dac7818468f22772d90a9d","license":"BUSL-1.1","urls":["bzz-raw://cee9d09370d968138d775c39525db4cd0768d60d17be7685519de12444e7dd2f","dweb:/ipfs/QmUdGh8wpMei3edKiEWA6S96s9dRt4ekZKJ4nau356X8xQ"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/Merkle.sol":{"keccak256":"0x606eabfdc2241dab76f7c6e6754324ae9eb12b0a5068984d2c11e2cd2fa94d98","license":"MIT","urls":["bzz-raw://a69c88393e9cf58ab066b75c75134b8c7cd51c242b726767cd8ec7e7d8351916","dweb:/ipfs/QmaNMz951WD5JZeQs5yav29mZn2E6fvdFm5u3moMupRzSM"]},"lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/permissions/Pausable.sol":{"keccak256":"0xce8ee0ab28f2bce9e94aa19fffe55bebef080327632ac98ff3ab14994b369bc0","license":"BUSL-1.1","urls":["bzz-raw://5c7e2be97a8840fa2a0434077a36136553a84efd9bff4b46712ce9fddb813a6a","dweb:/ipfs/QmZKvgPxLAbGo1CqTA4AX6MCDPFLSSNt43ZKWRjvvzFp7S"]},"lib/eigenlayer-middleware/src/BLSApkRegistry.sol":{"keccak256":"0xdd88bbc550a146eef8c694dd855db4e305594a4507967eeb31f2d26c47246e1d","license":"BUSL-1.1","urls":["bzz-raw://b0bd9b34c723c85b51edda8ae3dd4b796ea96afb5d98facf91738fd5c689e0b1","dweb:/ipfs/Qma8kTBq2SjSSNXxAf3e4ddkT3AkBZHHgxMFL4mjfZJv9B"]},"lib/eigenlayer-middleware/src/BLSApkRegistryStorage.sol":{"keccak256":"0xf61107c6cf909dc5745f6718b0e93ce2c4bdd947112bb3a18246d350b46edef3","license":"BUSL-1.1","urls":["bzz-raw://b15007adf4937aeb7540d79fb566086d7510f36545a6d9d57c46fdd4f0625122","dweb:/ipfs/QmVQa9GbCVcVCa9DHaQrNZpnVe1G6wznhctuPgTQLTTcVA"]},"lib/eigenlayer-middleware/src/BLSSignatureChecker.sol":{"keccak256":"0xc4d7b99bfe89cff16037cac6e2487f3fb859dec130eeba4ff2e17138957afad2","license":"BUSL-1.1","urls":["bzz-raw://7fac14f0ab8b2d34bbaff8a4ac5f7b350619d8981ed92f63fbde27a995541b12","dweb:/ipfs/QmYkt3j2f4Aa9ntRPcQQ4WJxetyvKBKg9H8nEKH72wDShk"]},"lib/eigenlayer-middleware/src/OperatorStateRetriever.sol":{"keccak256":"0x5573c9b7416d08e8b2f3e2e238ca4ba7a0c0fd4e6c6f8d4f7eca5487f26a042a","license":"BUSL-1.1","urls":["bzz-raw://98c9e6ec2b3478f3a962d57e280ddb69a93be7035ed7a4cdb775d29b763053af","dweb:/ipfs/QmaMHNFsddfP7fKxaVwn8foWqwp7ySwaD5Lof19bsmsdvg"]},"lib/eigenlayer-middleware/src/RegistryCoordinator.sol":{"keccak256":"0x2258bdf153fd9f76e37b07478089aa6f48e7c94d0c70bd91a2fc168040c3e2d4","license":"BUSL-1.1","urls":["bzz-raw://4ecf341836454b24d75f559082a68942445024b56f41ceb8ecbfa99eb0f8c0c8","dweb:/ipfs/Qmbhno4vPwbLz9ioUsygbhyLX6mD3oLx145T6ZBXmKnbRU"]},"lib/eigenlayer-middleware/src/RegistryCoordinatorStorage.sol":{"keccak256":"0x75cde4bc83b4f19a95b9447c9faf5aadbf4c579d7acb6ab0cfaef1b674777130","license":"BUSL-1.1","urls":["bzz-raw://46aca5d4c2ca28e58486279fa33117f070129435dbd6ade35903d576a5aac1da","dweb:/ipfs/QmUnobvB1qDf9LCCuN89DqLW3mCTmx3nzdzeUjj9BVQctQ"]},"lib/eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol":{"keccak256":"0xc07a5edfd95ab4f16f16a8dc8e76eadf4b0e90fe49db90540d01daaad86898c5","license":"BUSL-1.1","urls":["bzz-raw://52b53266450a53da641e82d8ae3be93c5e09f8342b4ea0cc96bb9038d8406354","dweb:/ipfs/QmVuoiQyqPTLCGnyt8zDaxiyaj4ETdgTGKv4MDHWzqEDjp"]},"lib/eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol":{"keccak256":"0x91c233280d6707404c65b7989c3fec6997c40cb3ab7d6c2e3f021102a0e2750d","license":"BUSL-1.1","urls":["bzz-raw://f2033dbb94acab37f3505734d8aad1481fbceedaa4742871f07506243a195aeb","dweb:/ipfs/QmXWJNkhUxfMhGfuFWw4UAU6nvw9qP9aswisQJLnZUUCzs"]},"lib/eigenlayer-middleware/src/interfaces/IIndexRegistry.sol":{"keccak256":"0x83b2d56aacf27e65c4959a832c5de573e013908c044f6e48ea8284ac5282ae2b","license":"BUSL-1.1","urls":["bzz-raw://877af382587e96bb39bcc6db8bb5e4b871db5025c52347d4bee9afeaa4a6cc8d","dweb:/ipfs/QmdnhsQCChzq2o5NgbeT3JxSsEcMm1PC9QW6zenZNPjD9F"]},"lib/eigenlayer-middleware/src/interfaces/IRegistry.sol":{"keccak256":"0x51426a17fb7e54bd3720e2890104e97a8559a13ff248b3d6b840916751c143d3","license":"BUSL-1.1","urls":["bzz-raw://01f91289e6100d528cb8b318cb14ff22a0bc52882c9d4db41585e030cc9ddc25","dweb:/ipfs/Qmb22nqGrsrtNovHRwbMCvDHGENuxAgrWu3Db4p7Er2MHY"]},"lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol":{"keccak256":"0xdd8effb082c1d5957d5ff43d7c59497b32866a6d82bcc7d5efa49ea9bc9b3385","license":"BUSL-1.1","urls":["bzz-raw://cb33a2a1446585b13b7a509e880c60d658d2d2522ec48a9f02e30d2cff54002d","dweb:/ipfs/QmVNG8ZPZkXzNEadPdTj1uBYLiZdCnYfsE5iGU6nJcJXiD"]},"lib/eigenlayer-middleware/src/interfaces/IServiceManager.sol":{"keccak256":"0xde4d88891f393ef6f3fefabd71a5b3d1305c7373f9e33a13ad30683c1d8c63a5","license":"BUSL-1.1","urls":["bzz-raw://fbab45eee5148f1764396a76a427cbec1dff1854ee4a2ecf7e47acf129f4797d","dweb:/ipfs/Qmc2yBePJYJMwmXPryAodK2MmBjJx2xYa3Na4pWNUwc3rJ"]},"lib/eigenlayer-middleware/src/interfaces/IServiceManagerUI.sol":{"keccak256":"0x365761699b4a5b7360ee6c75f12606eefc4b0394754c8b8e1e1eefec0cba7ffb","license":"BUSL-1.1","urls":["bzz-raw://ac14ada180b66cbbc5f9de0d6b4bb87b5946d2a9569ae88f2f62aaca47e879a8","dweb:/ipfs/QmcN9xKYF24naNWc6cYah9suz5gJSYE9nPj9eZFWnP35VX"]},"lib/eigenlayer-middleware/src/interfaces/ISocketUpdater.sol":{"keccak256":"0x2f209d4556d493b7b9d30a48eb98b9ee17f823ff677e9c656ebd6ed454b3626e","license":"BUSL-1.1","urls":["bzz-raw://1f1a7930cf8acf19684bcc39ea958d354586846f6dac0fd6a10e69a30eebea25","dweb:/ipfs/QmNsuiA1KKx22mf2YxYvK8CeuqUsgPNfUR7ijvnpBZWETw"]},"lib/eigenlayer-middleware/src/interfaces/IStakeRegistry.sol":{"keccak256":"0x1b8b4d757c1b804bc4cf6fbbf8bf8f89ebdeb30a31014751fe7d01deb9d513d4","license":"BUSL-1.1","urls":["bzz-raw://984bf2777b898ed187d28997f9783f5c293a1a1848e3e9aa470ce9183d454c97","dweb:/ipfs/Qme3aTpBrkLu8wYHFMZbCfhXHoZ1M6SpXkeC237T9BuU5B"]},"lib/eigenlayer-middleware/src/libraries/BN254.sol":{"keccak256":"0xb428c8d0c3b325507a88a61a80115493eb88606ccc19ed64a31e11294ab853b3","license":"MIT","urls":["bzz-raw://d7b6fb935bfe0494e6ff970c8f30a86d5f4cf5c3e0967300c28cd383c043acae","dweb:/ipfs/QmUHfFZaVjLPXhkBmcxrZhAHZaSFQDqXtrLGpjGBQBa5Ki"]},"lib/eigenlayer-middleware/src/libraries/BitmapUtils.sol":{"keccak256":"0x44315ac460be30a6b18fd4df4d1b8afb46653bf4dc06ca9f93c32353fd0605c5","license":"MIT","urls":["bzz-raw://da14f2ead3a375b02afd09d4a02edddf7b63a88945746b96789b2473184fdb04","dweb:/ipfs/QmRqcjxa2Vv2MrLdPeAwsktXdWTirapEDsRbJCyYRtKT6g"]},"src/IInferenceDB.sol":{"keccak256":"0x62da4685a54a9e252952087513987b328023ceb8a1d5d39b3b0316054f11339e","license":"MIT","urls":["bzz-raw://19da4e25d55b01ab81311a05aed03ac306a246e3e774121d6f653f26266561a7","dweb:/ipfs/QmdJUvxSMGZiRyimtMm4dm4y78GhcohmaBVRoxnE4D69JZ"]},"src/IOmronTaskManager.sol":{"keccak256":"0xd5bdcfbd09074fafd90c8e9412e39338e666d38e775173a641ae6936d078fe63","license":"UNLICENSED","urls":["bzz-raw://2edacc185b771c389cb6118243e91e2f86a63ce1f78d72d579dc8af7f8b8c1dc","dweb:/ipfs/QmWTPdZkCDGztkWSKuF221KFA1puMeY3GLy35M7GDbzbBf"]},"src/IZKVerifier.sol":{"keccak256":"0xff584258944883c1c8ef9492747167ab430df7daac904ef24b025652a9706591","license":"MIT","urls":["bzz-raw://2bd77c9673b57ae96bb85fe1c4e08c9bc52678972a08456f3a38595c2f43c9e8","dweb:/ipfs/QmUHBoFgA5A8aKY2HiMuwWGjnn3TEWdgG4kCEy3oURxfpS"]},"src/OmronTaskManager.sol":{"keccak256":"0x5a2773d5c2b3c90495526e3629c8a7cb059b106aeb63d3bd538e27f098bdb72b","license":"UNLICENSED","urls":["bzz-raw://665cfe4e60812e9f63d622deada4a900d361f1c5d807862def150b434c61abcb","dweb:/ipfs/Qmc3Py651aaAtVTQSshaHSpojYHzxLRqyzqdtYuoXt6DML"]}},"version":1}',
  metadata: {
    compiler: { version: "0.8.12+commit.f00d7308" },
    language: "Solidity",
    output: {
      abi: [
        {
          inputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "_registryCoordinator",
              type: "address",
            },
            {
              internalType: "uint32",
              name: "_taskResponseWindowBlock",
              type: "uint32",
            },
          ],
          stateMutability: "nonpayable",
          type: "constructor",
        },
        {
          inputs: [
            {
              internalType: "uint8",
              name: "version",
              type: "uint8",
              indexed: false,
            },
          ],
          type: "event",
          name: "Initialized",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "uint32",
              name: "taskIndex",
              type: "uint32",
              indexed: true,
            },
            {
              internalType: "struct IOmronTaskManager.Task",
              name: "task",
              type: "tuple",
              components: [
                {
                  internalType: "uint256[5]",
                  name: "inputs",
                  type: "uint256[5]",
                },
                {
                  internalType: "uint32",
                  name: "taskCreatedBlock",
                  type: "uint32",
                },
                { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
                {
                  internalType: "uint32",
                  name: "quorumThresholdPercentage",
                  type: "uint32",
                },
              ],
              indexed: false,
            },
          ],
          type: "event",
          name: "NewTaskCreated",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "address",
              name: "previousOwner",
              type: "address",
              indexed: true,
            },
            {
              internalType: "address",
              name: "newOwner",
              type: "address",
              indexed: true,
            },
          ],
          type: "event",
          name: "OwnershipTransferred",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "address",
              name: "account",
              type: "address",
              indexed: true,
            },
            {
              internalType: "uint256",
              name: "newPausedStatus",
              type: "uint256",
              indexed: false,
            },
          ],
          type: "event",
          name: "Paused",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "contract IPauserRegistry",
              name: "pauserRegistry",
              type: "address",
              indexed: false,
            },
            {
              internalType: "contract IPauserRegistry",
              name: "newPauserRegistry",
              type: "address",
              indexed: false,
            },
          ],
          type: "event",
          name: "PauserRegistrySet",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "bool",
              name: "value",
              type: "bool",
              indexed: false,
            },
          ],
          type: "event",
          name: "StaleStakesForbiddenUpdate",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "uint32",
              name: "taskIndex",
              type: "uint32",
              indexed: true,
            },
          ],
          type: "event",
          name: "TaskChallenged",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "uint32",
              name: "taskIndex",
              type: "uint32",
              indexed: true,
            },
            {
              internalType: "address",
              name: "challenger",
              type: "address",
              indexed: true,
            },
          ],
          type: "event",
          name: "TaskChallengedSuccessfully",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "uint32",
              name: "taskIndex",
              type: "uint32",
              indexed: true,
            },
            {
              internalType: "address",
              name: "prover",
              type: "address",
              indexed: true,
            },
          ],
          type: "event",
          name: "TaskChallengedUnsuccessfully",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "uint32",
              name: "taskIndex",
              type: "uint32",
              indexed: true,
            },
          ],
          type: "event",
          name: "TaskCompleted",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "struct IOmronTaskManager.TaskResponse",
              name: "taskResponse",
              type: "tuple",
              components: [
                {
                  internalType: "uint32",
                  name: "referenceTaskIndex",
                  type: "uint32",
                },
                { internalType: "uint256", name: "output", type: "uint256" },
              ],
              indexed: false,
            },
            {
              internalType: "struct IOmronTaskManager.TaskResponseMetadata",
              name: "taskResponseMetadata",
              type: "tuple",
              components: [
                {
                  internalType: "uint32",
                  name: "taskResponsedBlock",
                  type: "uint32",
                },
                {
                  internalType: "bytes32",
                  name: "hashOfNonSigners",
                  type: "bytes32",
                },
              ],
              indexed: false,
            },
          ],
          type: "event",
          name: "TaskResponded",
          anonymous: false,
        },
        {
          inputs: [
            {
              internalType: "address",
              name: "account",
              type: "address",
              indexed: true,
            },
            {
              internalType: "uint256",
              name: "newPausedStatus",
              type: "uint256",
              indexed: false,
            },
          ],
          type: "event",
          name: "Unpaused",
          anonymous: false,
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "TASK_CHALLENGE_WINDOW_BLOCK",
          outputs: [{ internalType: "uint32", name: "", type: "uint32" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "TASK_RESPONSE_WINDOW_BLOCK",
          outputs: [{ internalType: "uint32", name: "", type: "uint32" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "aggregator",
          outputs: [{ internalType: "address", name: "", type: "address" }],
        },
        {
          inputs: [{ internalType: "uint32", name: "", type: "uint32" }],
          stateMutability: "view",
          type: "function",
          name: "allTaskHashes",
          outputs: [{ internalType: "bytes32", name: "", type: "bytes32" }],
        },
        {
          inputs: [{ internalType: "uint32", name: "", type: "uint32" }],
          stateMutability: "view",
          type: "function",
          name: "allTaskResponses",
          outputs: [{ internalType: "bytes32", name: "", type: "bytes32" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "blsApkRegistry",
          outputs: [
            {
              internalType: "contract IBLSApkRegistry",
              name: "",
              type: "address",
            },
          ],
        },
        {
          inputs: [
            { internalType: "uint32", name: "id", type: "uint32" },
            { internalType: "uint256", name: "index", type: "uint256" },
          ],
          stateMutability: "view",
          type: "function",
          name: "challengeInstances",
          outputs: [{ internalType: "uint256", name: "", type: "uint256" }],
        },
        {
          inputs: [
            { internalType: "bytes32", name: "msgHash", type: "bytes32" },
            { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
            {
              internalType: "uint32",
              name: "referenceBlockNumber",
              type: "uint32",
            },
            {
              internalType:
                "struct IBLSSignatureChecker.NonSignerStakesAndSignature",
              name: "params",
              type: "tuple",
              components: [
                {
                  internalType: "uint32[]",
                  name: "nonSignerQuorumBitmapIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "struct BN254.G1Point[]",
                  name: "nonSignerPubkeys",
                  type: "tuple[]",
                  components: [
                    { internalType: "uint256", name: "X", type: "uint256" },
                    { internalType: "uint256", name: "Y", type: "uint256" },
                  ],
                },
                {
                  internalType: "struct BN254.G1Point[]",
                  name: "quorumApks",
                  type: "tuple[]",
                  components: [
                    { internalType: "uint256", name: "X", type: "uint256" },
                    { internalType: "uint256", name: "Y", type: "uint256" },
                  ],
                },
                {
                  internalType: "struct BN254.G2Point",
                  name: "apkG2",
                  type: "tuple",
                  components: [
                    {
                      internalType: "uint256[2]",
                      name: "X",
                      type: "uint256[2]",
                    },
                    {
                      internalType: "uint256[2]",
                      name: "Y",
                      type: "uint256[2]",
                    },
                  ],
                },
                {
                  internalType: "struct BN254.G1Point",
                  name: "sigma",
                  type: "tuple",
                  components: [
                    { internalType: "uint256", name: "X", type: "uint256" },
                    { internalType: "uint256", name: "Y", type: "uint256" },
                  ],
                },
                {
                  internalType: "uint32[]",
                  name: "quorumApkIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "uint32[]",
                  name: "totalStakeIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "uint32[][]",
                  name: "nonSignerStakeIndices",
                  type: "uint32[][]",
                },
              ],
            },
          ],
          stateMutability: "view",
          type: "function",
          name: "checkSignatures",
          outputs: [
            {
              internalType: "struct IBLSSignatureChecker.QuorumStakeTotals",
              name: "",
              type: "tuple",
              components: [
                {
                  internalType: "uint96[]",
                  name: "signedStakeForQuorum",
                  type: "uint96[]",
                },
                {
                  internalType: "uint96[]",
                  name: "totalStakeForQuorum",
                  type: "uint96[]",
                },
              ],
            },
            { internalType: "bytes32", name: "", type: "bytes32" },
          ],
        },
        {
          inputs: [
            {
              internalType: "struct IOmronTaskManager.Task",
              name: "task",
              type: "tuple",
              components: [
                {
                  internalType: "uint256[5]",
                  name: "inputs",
                  type: "uint256[5]",
                },
                {
                  internalType: "uint32",
                  name: "taskCreatedBlock",
                  type: "uint32",
                },
                { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
                {
                  internalType: "uint32",
                  name: "quorumThresholdPercentage",
                  type: "uint32",
                },
              ],
            },
            {
              internalType: "struct IOmronTaskManager.TaskResponse",
              name: "taskResponse",
              type: "tuple",
              components: [
                {
                  internalType: "uint32",
                  name: "referenceTaskIndex",
                  type: "uint32",
                },
                { internalType: "uint256", name: "output", type: "uint256" },
              ],
            },
            {
              internalType: "struct IOmronTaskManager.TaskResponseMetadata",
              name: "taskResponseMetadata",
              type: "tuple",
              components: [
                {
                  internalType: "uint32",
                  name: "taskResponsedBlock",
                  type: "uint32",
                },
                {
                  internalType: "bytes32",
                  name: "hashOfNonSigners",
                  type: "bytes32",
                },
              ],
            },
            {
              internalType: "struct BN254.G1Point[]",
              name: "pubkeysOfNonSigningOperators",
              type: "tuple[]",
              components: [
                { internalType: "uint256", name: "X", type: "uint256" },
                { internalType: "uint256", name: "Y", type: "uint256" },
              ],
            },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "confirmChallenge",
        },
        {
          inputs: [
            { internalType: "uint256[5]", name: "inputs", type: "uint256[5]" },
            {
              internalType: "uint32",
              name: "quorumThresholdPercentage",
              type: "uint32",
            },
            { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "createNewTask",
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "delegation",
          outputs: [
            {
              internalType: "contract IDelegationManager",
              name: "",
              type: "address",
            },
          ],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "generator",
          outputs: [{ internalType: "address", name: "", type: "address" }],
        },
        {
          inputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "registryCoordinator",
              type: "address",
            },
            {
              internalType: "bytes32[]",
              name: "operatorIds",
              type: "bytes32[]",
            },
          ],
          stateMutability: "view",
          type: "function",
          name: "getBatchOperatorFromId",
          outputs: [
            { internalType: "address[]", name: "operators", type: "address[]" },
          ],
        },
        {
          inputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "registryCoordinator",
              type: "address",
            },
            { internalType: "address[]", name: "operators", type: "address[]" },
          ],
          stateMutability: "view",
          type: "function",
          name: "getBatchOperatorId",
          outputs: [
            {
              internalType: "bytes32[]",
              name: "operatorIds",
              type: "bytes32[]",
            },
          ],
        },
        {
          inputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "registryCoordinator",
              type: "address",
            },
            {
              internalType: "uint32",
              name: "referenceBlockNumber",
              type: "uint32",
            },
            { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
            {
              internalType: "bytes32[]",
              name: "nonSignerOperatorIds",
              type: "bytes32[]",
            },
          ],
          stateMutability: "view",
          type: "function",
          name: "getCheckSignaturesIndices",
          outputs: [
            {
              internalType:
                "struct OperatorStateRetriever.CheckSignaturesIndices",
              name: "",
              type: "tuple",
              components: [
                {
                  internalType: "uint32[]",
                  name: "nonSignerQuorumBitmapIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "uint32[]",
                  name: "quorumApkIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "uint32[]",
                  name: "totalStakeIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "uint32[][]",
                  name: "nonSignerStakeIndices",
                  type: "uint32[][]",
                },
              ],
            },
          ],
        },
        {
          inputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "registryCoordinator",
              type: "address",
            },
            { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
            { internalType: "uint32", name: "blockNumber", type: "uint32" },
          ],
          stateMutability: "view",
          type: "function",
          name: "getOperatorState",
          outputs: [
            {
              internalType: "struct OperatorStateRetriever.Operator[][]",
              name: "",
              type: "tuple[][]",
              components: [
                { internalType: "address", name: "operator", type: "address" },
                {
                  internalType: "bytes32",
                  name: "operatorId",
                  type: "bytes32",
                },
                { internalType: "uint96", name: "stake", type: "uint96" },
              ],
            },
          ],
        },
        {
          inputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "registryCoordinator",
              type: "address",
            },
            { internalType: "bytes32", name: "operatorId", type: "bytes32" },
            { internalType: "uint32", name: "blockNumber", type: "uint32" },
          ],
          stateMutability: "view",
          type: "function",
          name: "getOperatorState",
          outputs: [
            { internalType: "uint256", name: "", type: "uint256" },
            {
              internalType: "struct OperatorStateRetriever.Operator[][]",
              name: "",
              type: "tuple[][]",
              components: [
                { internalType: "address", name: "operator", type: "address" },
                {
                  internalType: "bytes32",
                  name: "operatorId",
                  type: "bytes32",
                },
                { internalType: "uint96", name: "stake", type: "uint96" },
              ],
            },
          ],
        },
        {
          inputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "registryCoordinator",
              type: "address",
            },
            {
              internalType: "bytes32[]",
              name: "operatorIds",
              type: "bytes32[]",
            },
            { internalType: "uint32", name: "blockNumber", type: "uint32" },
          ],
          stateMutability: "view",
          type: "function",
          name: "getQuorumBitmapsAtBlockNumber",
          outputs: [{ internalType: "uint256[]", name: "", type: "uint256[]" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "getTaskResponseWindowBlock",
          outputs: [{ internalType: "uint32", name: "", type: "uint32" }],
        },
        {
          inputs: [
            {
              internalType: "contract IPauserRegistry",
              name: "_pauserRegistry",
              type: "address",
            },
            { internalType: "address", name: "initialOwner", type: "address" },
            { internalType: "address", name: "_aggregator", type: "address" },
            { internalType: "address", name: "_generator", type: "address" },
            { internalType: "address", name: "_inferenceDB", type: "address" },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "initialize",
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "latestTaskNum",
          outputs: [{ internalType: "uint32", name: "", type: "uint32" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "owner",
          outputs: [{ internalType: "address", name: "", type: "address" }],
        },
        {
          inputs: [
            {
              internalType: "uint256",
              name: "newPausedStatus",
              type: "uint256",
            },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "pause",
        },
        {
          inputs: [],
          stateMutability: "nonpayable",
          type: "function",
          name: "pauseAll",
        },
        {
          inputs: [{ internalType: "uint8", name: "index", type: "uint8" }],
          stateMutability: "view",
          type: "function",
          name: "paused",
          outputs: [{ internalType: "bool", name: "", type: "bool" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "paused",
          outputs: [{ internalType: "uint256", name: "", type: "uint256" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "pauserRegistry",
          outputs: [
            {
              internalType: "contract IPauserRegistry",
              name: "",
              type: "address",
            },
          ],
        },
        {
          inputs: [
            { internalType: "uint32", name: "taskId", type: "uint32" },
            { internalType: "uint256[]", name: "instances", type: "uint256[]" },
            { internalType: "bytes", name: "proof", type: "bytes" },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "proveResultAccurate",
        },
        {
          inputs: [
            {
              internalType: "struct IOmronTaskManager.Task",
              name: "task",
              type: "tuple",
              components: [
                {
                  internalType: "uint256[5]",
                  name: "inputs",
                  type: "uint256[5]",
                },
                {
                  internalType: "uint32",
                  name: "taskCreatedBlock",
                  type: "uint32",
                },
                { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
                {
                  internalType: "uint32",
                  name: "quorumThresholdPercentage",
                  type: "uint32",
                },
              ],
            },
            {
              internalType: "struct IOmronTaskManager.TaskResponse",
              name: "taskResponse",
              type: "tuple",
              components: [
                {
                  internalType: "uint32",
                  name: "referenceTaskIndex",
                  type: "uint32",
                },
                { internalType: "uint256", name: "output", type: "uint256" },
              ],
            },
            {
              internalType: "struct IOmronTaskManager.TaskResponseMetadata",
              name: "taskResponseMetadata",
              type: "tuple",
              components: [
                {
                  internalType: "uint32",
                  name: "taskResponsedBlock",
                  type: "uint32",
                },
                {
                  internalType: "bytes32",
                  name: "hashOfNonSigners",
                  type: "bytes32",
                },
              ],
            },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "raiseChallenger",
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "registryCoordinator",
          outputs: [
            {
              internalType: "contract IRegistryCoordinator",
              name: "",
              type: "address",
            },
          ],
        },
        {
          inputs: [],
          stateMutability: "nonpayable",
          type: "function",
          name: "renounceOwnership",
        },
        {
          inputs: [
            {
              internalType: "struct IOmronTaskManager.Task",
              name: "task",
              type: "tuple",
              components: [
                {
                  internalType: "uint256[5]",
                  name: "inputs",
                  type: "uint256[5]",
                },
                {
                  internalType: "uint32",
                  name: "taskCreatedBlock",
                  type: "uint32",
                },
                { internalType: "bytes", name: "quorumNumbers", type: "bytes" },
                {
                  internalType: "uint32",
                  name: "quorumThresholdPercentage",
                  type: "uint32",
                },
              ],
            },
            {
              internalType: "struct IOmronTaskManager.TaskResponse",
              name: "taskResponse",
              type: "tuple",
              components: [
                {
                  internalType: "uint32",
                  name: "referenceTaskIndex",
                  type: "uint32",
                },
                { internalType: "uint256", name: "output", type: "uint256" },
              ],
            },
            {
              internalType:
                "struct IBLSSignatureChecker.NonSignerStakesAndSignature",
              name: "nonSignerStakesAndSignature",
              type: "tuple",
              components: [
                {
                  internalType: "uint32[]",
                  name: "nonSignerQuorumBitmapIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "struct BN254.G1Point[]",
                  name: "nonSignerPubkeys",
                  type: "tuple[]",
                  components: [
                    { internalType: "uint256", name: "X", type: "uint256" },
                    { internalType: "uint256", name: "Y", type: "uint256" },
                  ],
                },
                {
                  internalType: "struct BN254.G1Point[]",
                  name: "quorumApks",
                  type: "tuple[]",
                  components: [
                    { internalType: "uint256", name: "X", type: "uint256" },
                    { internalType: "uint256", name: "Y", type: "uint256" },
                  ],
                },
                {
                  internalType: "struct BN254.G2Point",
                  name: "apkG2",
                  type: "tuple",
                  components: [
                    {
                      internalType: "uint256[2]",
                      name: "X",
                      type: "uint256[2]",
                    },
                    {
                      internalType: "uint256[2]",
                      name: "Y",
                      type: "uint256[2]",
                    },
                  ],
                },
                {
                  internalType: "struct BN254.G1Point",
                  name: "sigma",
                  type: "tuple",
                  components: [
                    { internalType: "uint256", name: "X", type: "uint256" },
                    { internalType: "uint256", name: "Y", type: "uint256" },
                  ],
                },
                {
                  internalType: "uint32[]",
                  name: "quorumApkIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "uint32[]",
                  name: "totalStakeIndices",
                  type: "uint32[]",
                },
                {
                  internalType: "uint32[][]",
                  name: "nonSignerStakeIndices",
                  type: "uint32[][]",
                },
              ],
            },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "respondToTask",
        },
        {
          inputs: [
            {
              internalType: "contract IPauserRegistry",
              name: "newPauserRegistry",
              type: "address",
            },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "setPauserRegistry",
        },
        {
          inputs: [{ internalType: "bool", name: "value", type: "bool" }],
          stateMutability: "nonpayable",
          type: "function",
          name: "setStaleStakesForbidden",
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "stakeRegistry",
          outputs: [
            {
              internalType: "contract IStakeRegistry",
              name: "",
              type: "address",
            },
          ],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "staleStakesForbidden",
          outputs: [{ internalType: "bool", name: "", type: "bool" }],
        },
        {
          inputs: [],
          stateMutability: "view",
          type: "function",
          name: "taskNumber",
          outputs: [{ internalType: "uint32", name: "", type: "uint32" }],
        },
        {
          inputs: [
            { internalType: "address", name: "newOwner", type: "address" },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "transferOwnership",
        },
        {
          inputs: [
            { internalType: "bytes32", name: "msgHash", type: "bytes32" },
            {
              internalType: "struct BN254.G1Point",
              name: "apk",
              type: "tuple",
              components: [
                { internalType: "uint256", name: "X", type: "uint256" },
                { internalType: "uint256", name: "Y", type: "uint256" },
              ],
            },
            {
              internalType: "struct BN254.G2Point",
              name: "apkG2",
              type: "tuple",
              components: [
                { internalType: "uint256[2]", name: "X", type: "uint256[2]" },
                { internalType: "uint256[2]", name: "Y", type: "uint256[2]" },
              ],
            },
            {
              internalType: "struct BN254.G1Point",
              name: "sigma",
              type: "tuple",
              components: [
                { internalType: "uint256", name: "X", type: "uint256" },
                { internalType: "uint256", name: "Y", type: "uint256" },
              ],
            },
          ],
          stateMutability: "view",
          type: "function",
          name: "trySignatureAndApkVerification",
          outputs: [
            { internalType: "bool", name: "pairingSuccessful", type: "bool" },
            { internalType: "bool", name: "siganatureIsValid", type: "bool" },
          ],
        },
        {
          inputs: [
            {
              internalType: "uint256",
              name: "newPausedStatus",
              type: "uint256",
            },
          ],
          stateMutability: "nonpayable",
          type: "function",
          name: "unpause",
        },
      ],
      devdoc: {
        kind: "dev",
        methods: {
          "checkSignatures(bytes32,bytes,uint32,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]))":
            {
              details:
                "Before signature verification, the function verifies operator stake information.  This includes ensuring that the provided `referenceBlockNumber` is correct, i.e., ensure that the stake returned from the specified block number is recent enough and that the stake is either the most recent update for the total stake (of the operator) or latest before the referenceBlockNumber.NOTE: Be careful to ensure `msgHash` is collision-resistant! This method does not hash  `msgHash` in any way, so if an attacker is able to pass in an arbitrary value, they may be able to tamper with signature verification.",
              params: {
                msgHash: "is the hash being signed",
                params:
                  "is the struct containing information on nonsigners, stakes, quorum apks, and the aggregate signature",
                quorumNumbers:
                  "is the bytes array of quorum numbers that are being signed for",
                referenceBlockNumber:
                  "is the block number at which the stake information is being verified",
              },
              returns: {
                _0: "quorumStakeTotals is the struct containing the total and signed stake for each quorum",
                _1: "signatoryRecordHash is the hash of the signatory record, which is used for fraud proofs",
              },
            },
          "getBatchOperatorFromId(address,bytes32[])": {
            details:
              "if an operator is not registered, the operator address will be 0",
            params: {
              operators:
                "is the array of operatorIds to get corresponding operator addresses for",
              registryCoordinator:
                "is the AVS registry coordinator to fetch the operator information from",
            },
          },
          "getBatchOperatorId(address,address[])": {
            details:
              "if an operator is not registered, the operatorId will be 0",
            params: {
              operators:
                "is the array of operator address to get corresponding operatorIds for",
              registryCoordinator:
                "is the AVS registry coordinator to fetch the operator information from",
            },
          },
          "getCheckSignaturesIndices(address,uint32,bytes,bytes32[])": {
            params: {
              nonSignerOperatorIds: "are the ids of the nonsigning operators",
              quorumNumbers:
                "are the ids of the quorums to get the operator state for",
              referenceBlockNumber:
                "is the block number to get the indices for",
              registryCoordinator:
                "is the registry coordinator to fetch the AVS registry information from",
            },
            returns: {
              _0: "1) the indices of the quorumBitmaps for each of the operators in the @param nonSignerOperatorIds array at the given blocknumber         2) the indices of the total stakes entries for the given quorums at the given blocknumber         3) the indices of the stakes of each of the nonsigners in each of the quorums they were a             part of (for each nonsigner, an array of length the number of quorums they were a part of            that are also part of the provided quorumNumbers) at the given blocknumber         4) the indices of the quorum apks for each of the provided quorums at the given blocknumber",
            },
          },
          "getOperatorState(address,bytes,uint32)": {
            params: {
              blockNumber: "is the block number to get the operator state for",
              quorumNumbers:
                "are the ids of the quorums to get the operator state for",
              registryCoordinator:
                "is the registry coordinator to fetch the AVS registry information from",
            },
            returns: {
              _0: "2d array of Operators. For each quorum, an ordered list of Operators",
            },
          },
          "getOperatorState(address,bytes32,uint32)": {
            params: {
              blockNumber: "is the block number to get the operator state for",
              operatorId: "the id of the operator to fetch the quorums lists ",
              registryCoordinator:
                "is the registry coordinator to fetch the AVS registry information from",
            },
            returns: {
              _0: "1) the quorumBitmap of the operator at the given blockNumber         2) 2d array of Operator structs. For each quorum the provided operator             was a part of at `blockNumber`, an ordered list of operators.",
            },
          },
          "getQuorumBitmapsAtBlockNumber(address,bytes32[],uint32)": {
            params: {
              blockNumber: "is the block number to get the quorumBitmaps for",
              operatorIds:
                "are the ids of the operators to get the quorumBitmaps for",
              registryCoordinator:
                "is the AVS registry coordinator to fetch the operator information from",
            },
          },
          "owner()": { details: "Returns the address of the current owner." },
          "pause(uint256)": {
            details:
              "This function can only pause functionality, and thus cannot 'unflip' any bit in `_paused` from 1 to 0.",
            params: {
              newPausedStatus:
                "represents the new value for `_paused` to take, which means it may flip several bits at once.",
            },
          },
          "renounceOwnership()": {
            details:
              "Leaves the contract without owner. It will not be possible to call `onlyOwner` functions anymore. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby removing any functionality that is only available to the owner.",
          },
          "setStaleStakesForbidden(bool)": {
            params: { value: "to toggle staleStakesForbidden" },
          },
          "transferOwnership(address)": {
            details:
              "Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner.",
          },
          "trySignatureAndApkVerification(bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256))":
            {
              params: {
                apk: "is the claimed G1 public key",
                apkG2: "is provided G2 public key",
                msgHash: "is the hash being signed",
                sigma: "is the G1 point signature",
              },
              returns: {
                pairingSuccessful:
                  "is true if the pairing precompile call was successful",
                siganatureIsValid: "is true if the signature is valid",
              },
            },
          "unpause(uint256)": {
            details:
              "This function can only unpause functionality, and thus cannot 'flip' any bit in `_paused` from 0 to 1.",
            params: {
              newPausedStatus:
                "represents the new value for `_paused` to take, which means it may flip several bits at once.",
            },
          },
        },
        version: 1,
      },
      userdoc: {
        kind: "user",
        methods: {
          "checkSignatures(bytes32,bytes,uint32,(uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]))":
            {
              notice:
                "This function is called by disperser when it has aggregated all the signatures of the operators that are part of the quorum for a particular taskNumber and is asserting them into onchain. The function checks that the claim for aggregated signatures are valid. The thesis of this procedure entails: - getting the aggregated pubkey of all registered nodes at the time of pre-commit by the disperser (represented by apk in the parameters), - subtracting the pubkeys of all the signers not in the quorum (nonSignerPubkeys) and storing  the output in apk to get aggregated pubkey of all operators that are part of quorum. - use this aggregated pubkey to verify the aggregated signature under BLS scheme. ",
            },
          "getBatchOperatorFromId(address,bytes32[])": {
            notice:
              "This function returns the operator addresses for each of the operators in the operatorIds array",
          },
          "getBatchOperatorId(address,address[])": {
            notice:
              "This function returns the operatorIds for each of the operators in the operators array",
          },
          "getCheckSignaturesIndices(address,uint32,bytes,bytes32[])": {
            notice:
              "this is called by the AVS operator to get the relevant indices for the checkSignatures function if they are not running an indexer    ",
          },
          "getOperatorState(address,bytes,uint32)": {
            notice:
              "returns the ordered list of operators (id and stake) for each quorum. The AVS coordinator  may call this function directly to get the operator state for a given block number",
          },
          "getOperatorState(address,bytes32,uint32)": {
            notice:
              "This function is intended to to be called by AVS operators every time a new task is created (i.e.) the AVS coordinator makes a request to AVS operators. Since all of the crucial information is kept onchain,  operators don't need to run indexers to fetch the data.",
          },
          "getQuorumBitmapsAtBlockNumber(address,bytes32[],uint32)": {
            notice:
              "this function returns the quorumBitmaps for each of the operators in the operatorIds array at the given blocknumber",
          },
          "getTaskResponseWindowBlock()": {
            notice: "Returns the TASK_RESPONSE_WINDOW_BLOCK",
          },
          "pause(uint256)": {
            notice:
              "This function is used to pause an EigenLayer contract's functionality. It is permissioned to the `pauser` address, which is expected to be a low threshold multisig.",
          },
          "pauseAll()": { notice: "Alias for `pause(type(uint256).max)`." },
          "paused()": {
            notice: "Returns the current paused status as a uint256.",
          },
          "paused(uint8)": {
            notice:
              "Returns 'true' if the `indexed`th bit of `_paused` is 1, and 'false' otherwise",
          },
          "pauserRegistry()": {
            notice:
              "Address of the `PauserRegistry` contract that this contract defers to for determining access control (for pausing).",
          },
          "setPauserRegistry(address)": {
            notice: "Allows the unpauser to set a new pauser registry",
          },
          "setStaleStakesForbidden(bool)": {
            notice:
              "RegistryCoordinator owner can either enforce or not that operator stakes are staler than the delegation.minWithdrawalDelayBlocks() window.",
          },
          "staleStakesForbidden()": {
            notice:
              "If true, check the staleness of the operator stakes and that its within the delegation withdrawalDelayBlocks window.",
          },
          "taskNumber()": {
            notice: "Returns the current 'taskNumber' for the middleware",
          },
          "trySignatureAndApkVerification(bytes32,(uint256,uint256),(uint256[2],uint256[2]),(uint256,uint256))":
            {
              notice:
                "trySignatureAndApkVerification verifies a BLS aggregate signature and the veracity of a calculated G1 Public key",
            },
          "unpause(uint256)": {
            notice:
              "This function is used to unpause an EigenLayer contract's functionality. It is permissioned to the `unpauser` address, which is expected to be a high threshold multisig or governance contract.",
          },
        },
        version: 1,
      },
    },
    settings: {
      remappings: [
        "@eigenlayer-middleware/=lib/eigenlayer-middleware/",
        "@eigenlayer-scripts/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/script/",
        "@eigenlayer/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/",
        "@omron/=src/",
        "@openzeppelin-upgrades/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/",
        "@openzeppelin/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/",
        "ds-test/=lib/eigenlayer-middleware/lib/ds-test/src/",
        "eigenlayer-contracts/=lib/eigenlayer-middleware/lib/eigenlayer-contracts/",
        "eigenlayer-middleware/=lib/eigenlayer-middleware/",
        "forge-std/=lib/forge-std/src/",
        "openzeppelin-contracts-upgradeable/=lib/eigenlayer-middleware/lib/openzeppelin-contracts-upgradeable/",
        "openzeppelin-contracts/=lib/eigenlayer-middleware/lib/openzeppelin-contracts/",
      ],
      optimizer: { enabled: true, runs: 200 },
      metadata: { bytecodeHash: "ipfs" },
      compilationTarget: { "src/OmronTaskManager.sol": "OmronTaskManager" },
      evmVersion: "london",
      libraries: {},
    },
    sources: {
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol":
        {
          keccak256:
            "0x247c62047745915c0af6b955470a72d1696ebad4352d7d3011aef1a2463cd888",
          urls: [
            "bzz-raw://d7fc8396619de513c96b6e00301b88dd790e83542aab918425633a5f7297a15a",
            "dweb:/ipfs/QmXbP4kiZyp7guuS7xe8KaybnwkRPGrBc2Kbi3vhcTfpxb",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol":
        {
          keccak256:
            "0x0203dcadc5737d9ef2c211d6fa15d18ebc3b30dfa51903b64870b01a062b0b4e",
          urls: [
            "bzz-raw://6eb2fd1e9894dbe778f4b8131adecebe570689e63cf892f4e21257bfe1252497",
            "dweb:/ipfs/QmXgUGNfZvrn6N2miv3nooSs7Jm34A41qz94fu2GtDFcx8",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/utils/AddressUpgradeable.sol":
        {
          keccak256:
            "0x611aa3f23e59cfdd1863c536776407b3e33d695152a266fa7cfb34440a29a8a3",
          urls: [
            "bzz-raw://9b4b2110b7f2b3eb32951bc08046fa90feccffa594e1176cb91cdfb0e94726b4",
            "dweb:/ipfs/QmSxLwYjicf9zWFuieRc8WQwE4FisA1Um5jp1iSa731TGt",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol":
        {
          keccak256:
            "0x963ea7f0b48b032eef72fe3a7582edf78408d6f834115b9feadd673a4d5bd149",
          urls: [
            "bzz-raw://d6520943ea55fdf5f0bafb39ed909f64de17051bc954ff3e88c9e5621412c79c",
            "dweb:/ipfs/QmWZ4rAKTQbNG2HxGs46AcTXShsVytKeLs7CUCdCSv5N7a",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/access/Ownable.sol":
        {
          keccak256:
            "0xa94b34880e3c1b0b931662cb1c09e5dfa6662f31cba80e07c5ee71cd135c9673",
          urls: [
            "bzz-raw://40fb1b5102468f783961d0af743f91b9980cf66b50d1d12009f6bb1869cea4d2",
            "dweb:/ipfs/QmYqEbJML4jB1GHbzD4cUZDtJg5wVwNm3vDJq1GbyDus8y",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/interfaces/IERC1271.sol":
        {
          keccak256:
            "0x0705a4b1b86d7b0bd8432118f226ba139c44b9dcaba0a6eafba2dd7d0639c544",
          urls: [
            "bzz-raw://c45b821ef9e882e57c256697a152e108f0f2ad6997609af8904cae99c9bd422e",
            "dweb:/ipfs/QmRKCJW6jjzR5UYZcLpGnhEJ75UVbH6EHkEa49sWx2SKng",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/proxy/beacon/IBeacon.sol":
        {
          keccak256:
            "0xd50a3421ac379ccb1be435fa646d66a65c986b4924f0849839f08692f39dde61",
          urls: [
            "bzz-raw://ada1e030c0231db8d143b44ce92b4d1158eedb087880cad6d8cc7bd7ebe7b354",
            "dweb:/ipfs/QmWZ2NHZweRpz1U9GF6R1h65ri76dnX7fNxLBeM2t5N5Ce",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol":
        {
          keccak256:
            "0x9750c6b834f7b43000631af5cc30001c5f547b3ceb3635488f140f60e897ea6b",
          urls: [
            "bzz-raw://5a7d5b1ef5d8d5889ad2ed89d8619c09383b80b72ab226e0fe7bde1636481e34",
            "dweb:/ipfs/QmebXWgtEfumQGBdVeM6c71McLixYXQP5Bk6kKXuoY4Bmr",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/Address.sol":
        {
          keccak256:
            "0xd6153ce99bcdcce22b124f755e72553295be6abcd63804cfdffceb188b8bef10",
          urls: [
            "bzz-raw://35c47bece3c03caaa07fab37dd2bb3413bfbca20db7bd9895024390e0a469487",
            "dweb:/ipfs/QmPGWT2x3QHcKxqe6gRmAkdakhbaRgx3DLzcakHz5M4eXG",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/Context.sol":
        {
          keccak256:
            "0xe2e337e6dde9ef6b680e07338c493ebea1b5fd09b43424112868e9cc1706bca7",
          urls: [
            "bzz-raw://6df0ddf21ce9f58271bdfaa85cde98b200ef242a05a3f85c2bc10a8294800a92",
            "dweb:/ipfs/QmRK2Y5Yc6BK7tGKkgsgn3aJEQGi5aakeSPZvS65PV8Xp3",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/Strings.sol":
        {
          keccak256:
            "0xaf159a8b1923ad2a26d516089bceca9bdeaeacd04be50983ea00ba63070f08a3",
          urls: [
            "bzz-raw://6f2cf1c531122bc7ca96b8c8db6a60deae60441e5223065e792553d4849b5638",
            "dweb:/ipfs/QmPBdJmBBABMDCfyDjCbdxgiqRavgiSL88SYPGibgbPas9",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol":
        {
          keccak256:
            "0x84ac2d2f343df1e683da7a12bbcf70db542a7a7a0cea90a5d70fcb5e5d035481",
          urls: [
            "bzz-raw://73ae8e0c6f975052973265113d762629002ce33987b1933c2a378667e2816f2f",
            "dweb:/ipfs/QmQAootkVfoe4PLaYbT4Xob2dJRm3bZfbCffEHRbCYXNPF",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/lib/openzeppelin-contracts/contracts/utils/cryptography/draft-EIP712.sol":
        {
          keccak256:
            "0x6688fad58b9ec0286d40fa957152e575d5d8bd4c3aa80985efdb11b44f776ae7",
          urls: [
            "bzz-raw://8bc00ab7f133cdaafd212a5cc6a16c8d37319721105d130c8e5af0c4e8f170ba",
            "dweb:/ipfs/QmVmf6LVMfFiEkvKYLzSv3bGHzymEW93AcUuFrNUdY3NtT",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IBeaconChainOracle.sol":
        {
          keccak256:
            "0x0fef07aa6179c77198f1514e12e628aa1c876e04f9c181ec853a322179e5be00",
          urls: [
            "bzz-raw://51438325876cc2d4c77f58488a7e27b488015d1b663c50be6a5cafbd73b9c983",
            "dweb:/ipfs/QmViCuGoYZzi6wtXA8PPKigqVv3KMuNxEVQ1Td9dGqjL18",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol":
        {
          keccak256:
            "0xe1626408822f552043f945d9aea18c5cbf878ef160da55e6f98706ed3a2acc07",
          urls: [
            "bzz-raw://426f6dddc040f2040f48dd4236c4201a3c978b4417ec3b4bd1004f8a48b29aaa",
            "dweb:/ipfs/QmWgY46nZiw1KQYNYMrJDTz7S9Y4KhyUoM9zVD92Mkf51S",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IETHPOSDeposit.sol":
        {
          keccak256:
            "0x2e60e5f4b0da0a0a4e2a07c63141120998559970c21deac743ea0c64a60a880c",
          urls: [
            "bzz-raw://e635c346bde5b7ade9bcf35bc733081520cb86015be4fbc6e761e6e9482c4c91",
            "dweb:/ipfs/QmRoeazEnbFn5SPSWAkoFK2gSN9DMp3hJAnrLWuL2sKutz",
          ],
          license: "CC0-1.0",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IEigenPod.sol":
        {
          keccak256:
            "0xb50c36ad96b6679bb80fd8331f949cbfbcba0f529026e1421a4d2bae64396eba",
          urls: [
            "bzz-raw://5719181d780120f1e688c0da276992a8caf185815917f453b3550537c31ed4cc",
            "dweb:/ipfs/QmYprRC5ZEXhz3zAUND5E8Xjn6s5TL8ZF8QbnndVq7aVPR",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IEigenPodManager.sol":
        {
          keccak256:
            "0xd8a64dbed03d3a5cdbefe1af75968f2dde07f973749c2ef5197bf7187c3e448c",
          urls: [
            "bzz-raw://27ccc7c1fd9352e9f9b357c9063d255dc0ed9583f43db09f786ac7497d7846b8",
            "dweb:/ipfs/QmeJzuJkE9m2NUNwZSp4tGZEZmih1LeucePup8hzMVDRbG",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IPausable.sol":
        {
          keccak256:
            "0x98cffc894842947377e24c1d375813a1120dd73a84c29782ab68404e109cb34f",
          urls: [
            "bzz-raw://b3474f6c350ceaee57cbdfb08fb48835d0c6e81ae8ebfbb9667899584a139324",
            "dweb:/ipfs/QmWELKtksdtWxQbqAccd8yGyhKqrgPZXTADKR7BuT27Zg5",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol":
        {
          keccak256:
            "0x9de8dd682bc0d812bbd6583c0231cbf35448d5eff58b74a93efa64cb9a768c49",
          urls: [
            "bzz-raw://c00d6c675b9c72b092d287fe85fd37782588df32b8eb59ab4c7db7a86be25e7d",
            "dweb:/ipfs/QmeYokY3HhAdbBaCPdHg3PgQEdRCDFEJy3Wf7VtgHBkQSx",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IPaymentCoordinator.sol":
        {
          keccak256:
            "0xf08a2adc0ec07ec0fd4711adab5f8a72fbb0dcafb62ee032b983b51b167348e6",
          urls: [
            "bzz-raw://702dda0a24d175188d87a9d924ce8244ad626b01aba547e9a4368811991e8950",
            "dweb:/ipfs/QmUgsi4BUcqXvtcLTQvk3wfVHCaCGwKc2gNixx4bz58csY",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol":
        {
          keccak256:
            "0x5e52482a31d94401a8502f3014c4aada1142b4450fc0596dff8e1866a85fe092",
          urls: [
            "bzz-raw://17dc326c9361bc1453379f26545963557b2883b0c88bc07d4477e04dbcc0cc8c",
            "dweb:/ipfs/QmZXT7A816W5JH2ymirE2ETaJttqztFCsEL22AV8oEfCK9",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/ISlasher.sol":
        {
          keccak256:
            "0x45dfaa2cfdde87f48a6ee38bb6fb739847aef7cf3f6137bdcd8c8a330559ec79",
          urls: [
            "bzz-raw://1b7f6bd75b42fcaa91ceb7140cb2c41926a1fe6ee2d3161e4fe6186b181ba232",
            "dweb:/ipfs/QmZjbdKiSs33C9i3GDc3sdD39Pz4YPkDoKftowoUF4kHmY",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol":
        {
          keccak256:
            "0xc530c6a944b70051fd0dac0222de9a4b5baadeaf94ad194daac6ad8d2ace7420",
          urls: [
            "bzz-raw://3767df0364ce835b52e786d2851431eb9223fe4747602107505477e162231d73",
            "dweb:/ipfs/QmZkH5bKUygQrJomndNaQqkefVRW4rRefCa8HPJ5HMczxJ",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol":
        {
          keccak256:
            "0xccd308b0996295c92058b70c3b83563c009c074cb6815329c5f35e1b1a0571f4",
          urls: [
            "bzz-raw://cd1050445ff7aeb588b3b037aab53e2d92c3abd4620e94dfc95aca870e71b821",
            "dweb:/ipfs/QmUC96Ctwn3KQr6VSqHPpAVJ5qEUSVnugaxiZ8gfXygW92",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/BeaconChainProofs.sol":
        {
          keccak256:
            "0x70d89b05c1c5f47b74a07fbb5a2c05e606fed494e749ea98a9915b7be73df377",
          urls: [
            "bzz-raw://db1d3bfaee69aef53c8b12b492a17584e6d1ac94610cb8b38aad33e1cdd81af7",
            "dweb:/ipfs/QmfVsMTj1hcf9fMEm5RzvtcBN4dMcAKFBgUUDsNDr5XFpq",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/EIP1271SignatureUtils.sol":
        {
          keccak256:
            "0xe92d584c47c5828e026a8082af3da38a853e3942c4da7deb705d6470a41afab3",
          urls: [
            "bzz-raw://1c436c578781fd7d3dffdb24e906819422819f5e9a71d39ee63166a3d5cb3373",
            "dweb:/ipfs/QmP7bJhYqLpwqk2Xq4tqDCUMi2nFAhxxW3Pz36ctE1sbdD",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/Endian.sol":
        {
          keccak256:
            "0xf3b72653ba2567a978d4612703fa5f71c5fcd015d8dac7818468f22772d90a9d",
          urls: [
            "bzz-raw://cee9d09370d968138d775c39525db4cd0768d60d17be7685519de12444e7dd2f",
            "dweb:/ipfs/QmUdGh8wpMei3edKiEWA6S96s9dRt4ekZKJ4nau356X8xQ",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/libraries/Merkle.sol":
        {
          keccak256:
            "0x606eabfdc2241dab76f7c6e6754324ae9eb12b0a5068984d2c11e2cd2fa94d98",
          urls: [
            "bzz-raw://a69c88393e9cf58ab066b75c75134b8c7cd51c242b726767cd8ec7e7d8351916",
            "dweb:/ipfs/QmaNMz951WD5JZeQs5yav29mZn2E6fvdFm5u3moMupRzSM",
          ],
          license: "MIT",
        },
      "lib/eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/permissions/Pausable.sol":
        {
          keccak256:
            "0xce8ee0ab28f2bce9e94aa19fffe55bebef080327632ac98ff3ab14994b369bc0",
          urls: [
            "bzz-raw://5c7e2be97a8840fa2a0434077a36136553a84efd9bff4b46712ce9fddb813a6a",
            "dweb:/ipfs/QmZKvgPxLAbGo1CqTA4AX6MCDPFLSSNt43ZKWRjvvzFp7S",
          ],
          license: "BUSL-1.1",
        },
      "lib/eigenlayer-middleware/src/BLSApkRegistry.sol": {
        keccak256:
          "0xdd88bbc550a146eef8c694dd855db4e305594a4507967eeb31f2d26c47246e1d",
        urls: [
          "bzz-raw://b0bd9b34c723c85b51edda8ae3dd4b796ea96afb5d98facf91738fd5c689e0b1",
          "dweb:/ipfs/Qma8kTBq2SjSSNXxAf3e4ddkT3AkBZHHgxMFL4mjfZJv9B",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/BLSApkRegistryStorage.sol": {
        keccak256:
          "0xf61107c6cf909dc5745f6718b0e93ce2c4bdd947112bb3a18246d350b46edef3",
        urls: [
          "bzz-raw://b15007adf4937aeb7540d79fb566086d7510f36545a6d9d57c46fdd4f0625122",
          "dweb:/ipfs/QmVQa9GbCVcVCa9DHaQrNZpnVe1G6wznhctuPgTQLTTcVA",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/BLSSignatureChecker.sol": {
        keccak256:
          "0xc4d7b99bfe89cff16037cac6e2487f3fb859dec130eeba4ff2e17138957afad2",
        urls: [
          "bzz-raw://7fac14f0ab8b2d34bbaff8a4ac5f7b350619d8981ed92f63fbde27a995541b12",
          "dweb:/ipfs/QmYkt3j2f4Aa9ntRPcQQ4WJxetyvKBKg9H8nEKH72wDShk",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/OperatorStateRetriever.sol": {
        keccak256:
          "0x5573c9b7416d08e8b2f3e2e238ca4ba7a0c0fd4e6c6f8d4f7eca5487f26a042a",
        urls: [
          "bzz-raw://98c9e6ec2b3478f3a962d57e280ddb69a93be7035ed7a4cdb775d29b763053af",
          "dweb:/ipfs/QmaMHNFsddfP7fKxaVwn8foWqwp7ySwaD5Lof19bsmsdvg",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/RegistryCoordinator.sol": {
        keccak256:
          "0x2258bdf153fd9f76e37b07478089aa6f48e7c94d0c70bd91a2fc168040c3e2d4",
        urls: [
          "bzz-raw://4ecf341836454b24d75f559082a68942445024b56f41ceb8ecbfa99eb0f8c0c8",
          "dweb:/ipfs/Qmbhno4vPwbLz9ioUsygbhyLX6mD3oLx145T6ZBXmKnbRU",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/RegistryCoordinatorStorage.sol": {
        keccak256:
          "0x75cde4bc83b4f19a95b9447c9faf5aadbf4c579d7acb6ab0cfaef1b674777130",
        urls: [
          "bzz-raw://46aca5d4c2ca28e58486279fa33117f070129435dbd6ade35903d576a5aac1da",
          "dweb:/ipfs/QmUnobvB1qDf9LCCuN89DqLW3mCTmx3nzdzeUjj9BVQctQ",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol": {
        keccak256:
          "0xc07a5edfd95ab4f16f16a8dc8e76eadf4b0e90fe49db90540d01daaad86898c5",
        urls: [
          "bzz-raw://52b53266450a53da641e82d8ae3be93c5e09f8342b4ea0cc96bb9038d8406354",
          "dweb:/ipfs/QmVuoiQyqPTLCGnyt8zDaxiyaj4ETdgTGKv4MDHWzqEDjp",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol": {
        keccak256:
          "0x91c233280d6707404c65b7989c3fec6997c40cb3ab7d6c2e3f021102a0e2750d",
        urls: [
          "bzz-raw://f2033dbb94acab37f3505734d8aad1481fbceedaa4742871f07506243a195aeb",
          "dweb:/ipfs/QmXWJNkhUxfMhGfuFWw4UAU6nvw9qP9aswisQJLnZUUCzs",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IIndexRegistry.sol": {
        keccak256:
          "0x83b2d56aacf27e65c4959a832c5de573e013908c044f6e48ea8284ac5282ae2b",
        urls: [
          "bzz-raw://877af382587e96bb39bcc6db8bb5e4b871db5025c52347d4bee9afeaa4a6cc8d",
          "dweb:/ipfs/QmdnhsQCChzq2o5NgbeT3JxSsEcMm1PC9QW6zenZNPjD9F",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IRegistry.sol": {
        keccak256:
          "0x51426a17fb7e54bd3720e2890104e97a8559a13ff248b3d6b840916751c143d3",
        urls: [
          "bzz-raw://01f91289e6100d528cb8b318cb14ff22a0bc52882c9d4db41585e030cc9ddc25",
          "dweb:/ipfs/Qmb22nqGrsrtNovHRwbMCvDHGENuxAgrWu3Db4p7Er2MHY",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol": {
        keccak256:
          "0xdd8effb082c1d5957d5ff43d7c59497b32866a6d82bcc7d5efa49ea9bc9b3385",
        urls: [
          "bzz-raw://cb33a2a1446585b13b7a509e880c60d658d2d2522ec48a9f02e30d2cff54002d",
          "dweb:/ipfs/QmVNG8ZPZkXzNEadPdTj1uBYLiZdCnYfsE5iGU6nJcJXiD",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IServiceManager.sol": {
        keccak256:
          "0xde4d88891f393ef6f3fefabd71a5b3d1305c7373f9e33a13ad30683c1d8c63a5",
        urls: [
          "bzz-raw://fbab45eee5148f1764396a76a427cbec1dff1854ee4a2ecf7e47acf129f4797d",
          "dweb:/ipfs/Qmc2yBePJYJMwmXPryAodK2MmBjJx2xYa3Na4pWNUwc3rJ",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IServiceManagerUI.sol": {
        keccak256:
          "0x365761699b4a5b7360ee6c75f12606eefc4b0394754c8b8e1e1eefec0cba7ffb",
        urls: [
          "bzz-raw://ac14ada180b66cbbc5f9de0d6b4bb87b5946d2a9569ae88f2f62aaca47e879a8",
          "dweb:/ipfs/QmcN9xKYF24naNWc6cYah9suz5gJSYE9nPj9eZFWnP35VX",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/ISocketUpdater.sol": {
        keccak256:
          "0x2f209d4556d493b7b9d30a48eb98b9ee17f823ff677e9c656ebd6ed454b3626e",
        urls: [
          "bzz-raw://1f1a7930cf8acf19684bcc39ea958d354586846f6dac0fd6a10e69a30eebea25",
          "dweb:/ipfs/QmNsuiA1KKx22mf2YxYvK8CeuqUsgPNfUR7ijvnpBZWETw",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/interfaces/IStakeRegistry.sol": {
        keccak256:
          "0x1b8b4d757c1b804bc4cf6fbbf8bf8f89ebdeb30a31014751fe7d01deb9d513d4",
        urls: [
          "bzz-raw://984bf2777b898ed187d28997f9783f5c293a1a1848e3e9aa470ce9183d454c97",
          "dweb:/ipfs/Qme3aTpBrkLu8wYHFMZbCfhXHoZ1M6SpXkeC237T9BuU5B",
        ],
        license: "BUSL-1.1",
      },
      "lib/eigenlayer-middleware/src/libraries/BN254.sol": {
        keccak256:
          "0xb428c8d0c3b325507a88a61a80115493eb88606ccc19ed64a31e11294ab853b3",
        urls: [
          "bzz-raw://d7b6fb935bfe0494e6ff970c8f30a86d5f4cf5c3e0967300c28cd383c043acae",
          "dweb:/ipfs/QmUHfFZaVjLPXhkBmcxrZhAHZaSFQDqXtrLGpjGBQBa5Ki",
        ],
        license: "MIT",
      },
      "lib/eigenlayer-middleware/src/libraries/BitmapUtils.sol": {
        keccak256:
          "0x44315ac460be30a6b18fd4df4d1b8afb46653bf4dc06ca9f93c32353fd0605c5",
        urls: [
          "bzz-raw://da14f2ead3a375b02afd09d4a02edddf7b63a88945746b96789b2473184fdb04",
          "dweb:/ipfs/QmRqcjxa2Vv2MrLdPeAwsktXdWTirapEDsRbJCyYRtKT6g",
        ],
        license: "MIT",
      },
      "src/IInferenceDB.sol": {
        keccak256:
          "0x62da4685a54a9e252952087513987b328023ceb8a1d5d39b3b0316054f11339e",
        urls: [
          "bzz-raw://19da4e25d55b01ab81311a05aed03ac306a246e3e774121d6f653f26266561a7",
          "dweb:/ipfs/QmdJUvxSMGZiRyimtMm4dm4y78GhcohmaBVRoxnE4D69JZ",
        ],
        license: "MIT",
      },
      "src/IOmronTaskManager.sol": {
        keccak256:
          "0xd5bdcfbd09074fafd90c8e9412e39338e666d38e775173a641ae6936d078fe63",
        urls: [
          "bzz-raw://2edacc185b771c389cb6118243e91e2f86a63ce1f78d72d579dc8af7f8b8c1dc",
          "dweb:/ipfs/QmWTPdZkCDGztkWSKuF221KFA1puMeY3GLy35M7GDbzbBf",
        ],
        license: "UNLICENSED",
      },
      "src/IZKVerifier.sol": {
        keccak256:
          "0xff584258944883c1c8ef9492747167ab430df7daac904ef24b025652a9706591",
        urls: [
          "bzz-raw://2bd77c9673b57ae96bb85fe1c4e08c9bc52678972a08456f3a38595c2f43c9e8",
          "dweb:/ipfs/QmUHBoFgA5A8aKY2HiMuwWGjnn3TEWdgG4kCEy3oURxfpS",
        ],
        license: "MIT",
      },
      "src/OmronTaskManager.sol": {
        keccak256:
          "0x5a2773d5c2b3c90495526e3629c8a7cb059b106aeb63d3bd538e27f098bdb72b",
        urls: [
          "bzz-raw://665cfe4e60812e9f63d622deada4a900d361f1c5d807862def150b434c61abcb",
          "dweb:/ipfs/Qmc3Py651aaAtVTQSshaHSpojYHzxLRqyzqdtYuoXt6DML",
        ],
        license: "UNLICENSED",
      },
    },
    version: 1,
  },
  id: 80,
};
