// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.29;
import {IStrategy} from "../lib/eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface ISertnServiceManager {
    error IncorrectAVS();
    error IncorrectOperatorSetIds();
    error NotModelId();
    error NotAggregator();
    error NotOperator();
    error NotTaskManager();
    error InvalidTaskManager();
    error AggregatorAlreadyExists();
    error NoProofOnResponse();
    error TaskCouldNotBeSent();
    error IncorrectOperator();
    error TaskNotExpired();
    error TaskExpired();
    error NotPausedLongEnough();
    error NotRegisteredAggregatorOrOperator();
    error TransferFailed();
    error BountyAlreadySet();
    error NotInSubmittedTasks();
    error MaxBlocksTooLong();
    error NoPermission();
    error NotRegisteredToModel();
    error InvalidAggregator();
    error NotSertnRegistrar();
    error ZeroAddress();

    event NewOperator(address opAddr_);
    event NewModels(uint256[] modelId_);
    event NewStrategies(address[] newSupportedTokens_);
    event NewTask(address indexed opAddr_, bytes indexed taskId_);
    event TaskResponded(
        uint256 indexed model,
        bytes indexed taskId,
        TaskResponse task
    );
    event UpForSlashing(address indexed operator, bytes indexed taskId);
    event ProofRequested(address indexed operator, bytes indexed taskId);
    event OperatorSlashed(address indexed operator, bytes indexed taskId);
    event ModelUpdated(uint256 indexed modelId, OperatorModel operatorModel);
    event OpInfoChanged(address indexed _operator, bytes _opInfo);
    event OperatorDeleted(address indexed _operator, uint32[] opSetIds);
    struct Model {
        string title_;
        string description_;
        address modelVerifier_;
        address[] operators_;
    }

    struct Operator {
        uint256[] models_;
        bytes32[] computeUnits_;
        bytes[] openTasks_;
        bytes[] submittedTasks_;
        bytes[] proofRequests_;
        uint256[] allocatedEth_;
        uint256 allocatedSer_;
        uint256[2] proofRequestExponents_;
        uint32 pausedBlock_;
    }

    struct OperatorModel {
        address operator_;
        uint256 modelId_;
        uint32 maxBlocks_;
        IStrategy[] ethStrategies_;
        uint256[] ethShares_;
        uint256 baseFee_;
        //May want to swap out in future for an operator's unique token.  For now will just use SER
        uint256 maxSer_;
        bytes32 computeType_;
        bool proveOnResponse_;
        bool available_;
    }

    struct Task {
        uint256 operatorModelId_;
        bytes inputs_;
        uint256 poc_;
        uint256 startTime_;
        uint32 startingBlock_;
        bool proveOnResponse_;
        address user_;
        uint256 nonce_;
    }

    struct TaskResponse {
        bytes taskId_;
        bytes output_;
        bool proven_;
    }

    /**
     * @notice Remove an aggregator from the service manager
     * @param _aggregator The address of the aggregator to remove
     */
    function removeAggregator(address _aggregator) external;

    /**
     * @notice Task completed
     */
    function taskCompleted(
        address _operator,
        uint256 _fee,
        IStrategy _strategy,
        IERC20 _token
    ) external;

    // State Variables
    function isAggregator(address) external view returns (bool);

    function operatorNodeModelIds(
        address,
        uint256,
        uint256
    ) external view returns (bool);

    function operatorNodeComputeUnits(
        address,
        uint256
    ) external view returns (uint256);

    function operatorNodeCount(address) external view returns (uint256);

    function opInfo(address) external view returns (bytes memory);

    function aggregators(uint256) external view returns (address);
}
