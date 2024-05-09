// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSquaringTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IIncredibleSquaringTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTask struct {
	Inputs                    [5]*big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IIncredibleSquaringTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Output             *big.Int
}

// IIncredibleSquaringTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractIncredibleSquaringTaskManagerMetaData contains all meta data concerning the ContractIncredibleSquaringTaskManager contract.
var ContractIncredibleSquaringTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ProveChallenge\",\"inputs\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskChallengeState\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIncredibleSquaringTaskManager.TASK_CHALLENGED_STATE\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskChallengedBlock\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskChallengedInstances\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"instances\",\"type\":\"uint256[6]\",\"internalType\":\"uint256[6]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162008c1e38038062008c1e8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051618927620002f7600039600081816102dd015281816106a20152615b9601526000818161063b0152614d2d01526000818161049b01528181614f0f015261586f0152600081816104c2015281816150e501526152a701526000818161050f0152818161150f015281816149f801528181614b900152614dca01526189276000f3fe608060405234801561001057600080fd5b50600436106102695760003560e01c806369b1843f11610151578063b030b7b4116100c3578063e7a6c1fa11610087578063e7a6c1fa1461065d578063f2fde38b1461068d578063f5c9899d146106a0578063f63c5bab146106c6578063f8c8765e146106ce578063fabc1cbc146106e157600080fd5b8063b030b7b4146105d5578063b98d0908146105f5578063bdeea6cc14610602578063cefdc1d414610615578063df5cf7231461063657600080fd5b80637afa1eed116101155780637afa1eed146105685780637cfd89681461057b578063886f11951461058e5780638b00ce7c146105a15780638da5cb5b146105b1578063a087d99a146105c257600080fd5b806369b1843f146104f75780636d14a9871461050a5780636efb463614610531578063715018a61461055257806372d18e8d1461055a57600080fd5b80633563b0d1116101ea5780635ac86ab7116101ae5780635ac86ab71461043b5780635c1556621461046e5780635c975abb1461048e5780635df459461461049657806368304835146104bd578063695ef6f9146104e457600080fd5b80633563b0d1146103c0578063416c7e5e146103e05780634d2b57fe146103f35780634f739f7414610413578063595c6a671461043357600080fd5b8063245a7bfc11610231578063245a7bfc1461031457806327e00b3a1461033f5780632cb223d5146103525780632d89f6fc1461038057806331b36bd9146103a057600080fd5b80630627721e1461026e57806310d67a2f14610283578063136439dd14610296578063171f1d5b146102a95780631ad43189146102d8575b600080fd5b61028161027c3660046171a3565b6106f4565b005b61028161029136600461721f565b61088f565b6102816102a436600461723c565b610942565b6102bc6102b73660046173ba565b610a81565b6040805192151583529015156020830152015b60405180910390f35b6102ff7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102cf565b60cf54610327906001600160a01b031681565b6040516001600160a01b0390911681526020016102cf565b61028161034d3660046174cd565b610c0b565b610372610360366004617553565b60cb6020526000908152604090205481565b6040519081526020016102cf565b61037261038e366004617553565b60ca6020526000908152604090205481565b6103b36103ae366004617570565b610f5b565b6040516102cf9190617659565b6103d36103ce3660046176e2565b611077565b6040516102cf91906177e9565b6102816103ee36600461780a565b61150d565b610406610401366004617882565b611682565b6040516102cf91906178d1565b61042661042136600461791e565b611797565b6040516102cf9190617a17565b610281611ebd565b61045e610449366004617ae1565b606654600160ff9092169190911b9081161490565b60405190151581526020016102cf565b61048161047c366004617afe565b611f84565b6040516102cf9190617b45565b606654610372565b6103277f000000000000000000000000000000000000000000000000000000000000000081565b6103277f000000000000000000000000000000000000000000000000000000000000000081565b61045e6104f2366004617b7d565b61214c565b610281610505366004617c23565b614561565b6103277f000000000000000000000000000000000000000000000000000000000000000081565b61054461053f366004617e84565b614645565b6040516102cf929190617f44565b61028161555c565b60c95463ffffffff166102ff565b60d054610327906001600160a01b031681565b610372610589366004617f8d565b615570565b606554610327906001600160a01b031681565b60c9546102ff9063ffffffff1681565b6033546001600160a01b0316610327565b6102816105d03660046174cd565b615595565b6103726105e3366004617553565b60cd6020526000908152604090205481565b60975461045e9060ff1681565b610281610610366004617fb9565b6159a4565b61062861062336600461802d565b615e24565b6040516102cf929190618064565b6103277f000000000000000000000000000000000000000000000000000000000000000081565b61068061066b366004617553565b60cc6020526000908152604090205460ff1681565b6040516102cf919061809b565b61028161069b36600461721f565b615fb6565b7f00000000000000000000000000000000000000000000000000000000000000006102ff565b6102ff606481565b6102816106dc3660046180c3565b61602c565b6102816106ef36600461723c565b61617d565b60d0546001600160a01b0316331461075d5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084015b60405180910390fd5b610765616fc5565b6040805160a081810190925290869060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f8501839004830281018301909152838152908490849081908401838280828437600092019190915250505050604080830191909152516107ef90829060200161811f565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f54977349061085290849061811f565b60405180910390a260c95461086e9063ffffffff1660016181eb565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109069190618213565b6001600160a01b0316336001600160a01b0316146109365760405162461bcd60e51b815260040161075490618230565b61093f816162d9565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561098a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ae919061827a565b6109ca5760405162461bcd60e51b815260040161075490618297565b60665481811614610a435760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610754565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610ac957610ac96182df565b60200201518951600160200201518a60200151600060028110610aee57610aee6182df565b60200201518b60200151600160028110610b0a57610b0a6182df565b602090810291909101518c518d830151604051610b679a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610b8a91906182f5565b9050610bfd610ba3610b9c88846163d0565b8690616467565b610bab6164fb565b610bf3610be485610bde604080518082018252600080825260209182015281518083019092526001825260029082015290565b906163d0565b610bed8c6165bb565b90616467565b886201d4c061664b565b909890975095505050505050565b6000610c1a6020850185617553565b63ffffffff8116600090815260cb6020526040902054909150610c895760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610754565b8383604051602001610c9c929190618335565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb90935291205414610d3b5760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610754565b63ffffffff8116600090815260cc602052604081205460ff166003811115610d6557610d65618085565b14610dd15760405162461bcd60e51b815260206004820152603660248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c6044820152753932b0b23c903132b2b71031b430b63632b733b2b21760511b6064820152608401610754565b6064610de06020850185617553565b610dea91906181eb565b63ffffffff164363ffffffff161115610e6b5760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610754565b610e73616ff3565b60005b6005811015610ec957868160058110610e9157610e916182df565b60200201356004610ea2919061836b565b828260068110610eb457610eb46182df565b6020020152610ec28161838a565b9050610e76565b50610ed96020860135600461836b565b60a082015263ffffffff8216600090815260ce60205260409020610eff90826006617011565b5063ffffffff8216600081815260cd6020908152604080832043905560cc909152808220805460ff19166001179055517f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a9190a2505050505050565b606081516001600160401b03811115610f7657610f76617255565b604051908082528060200260200182016040528015610f9f578160200160208202803683370190505b50905060005b825181101561107057836001600160a01b03166313542a4e848381518110610fcf57610fcf6182df565b60200260200101516040518263ffffffff1660e01b815260040161100291906001600160a01b0391909116815260200190565b602060405180830381865afa15801561101f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061104391906183a5565b828281518110611055576110556182df565b60209081029190910101526110698161838a565b9050610fa5565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110dd9190618213565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa15801561111f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111439190618213565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611185573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a99190618213565b9050600086516001600160401b038111156111c6576111c6617255565b6040519080825280602002602001820160405280156111f957816020015b60608152602001906001900390816111e45790505b50905060005b875181101561150157600088828151811061121c5761121c6182df565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa15801561127d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112a591908101906183be565b905080516001600160401b038111156112c0576112c0617255565b60405190808252806020026020018201604052801561130b57816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816112de5790505b5084848151811061131e5761131e6182df565b602002602001018190525060005b81518110156114eb576040518060600160405280876001600160a01b03166347b314e8858581518110611361576113616182df565b60200260200101516040518263ffffffff1660e01b815260040161138791815260200190565b602060405180830381865afa1580156113a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c89190618213565b6001600160a01b031681526020018383815181106113e8576113e86182df565b60200260200101518152602001896001600160a01b031663fa28c627858581518110611416576114166182df565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611472573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611496919061844e565b6001600160601b03168152508585815181106114b4576114b46182df565b602002602001015182815181106114cd576114cd6182df565b602002602001018190525080806114e39061838a565b91505061132c565b50505080806114f99061838a565b9150506111ff565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561156b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061158f9190618213565b6001600160a01b0316336001600160a01b03161461163b5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610754565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b606081516001600160401b0381111561169d5761169d617255565b6040519080825280602002602001820160405280156116c6578160200160208202803683370190505b50905060005b825181101561107057836001600160a01b031663296bb0648483815181106116f6576116f66182df565b60200260200101516040518263ffffffff1660e01b815260040161171c91815260200190565b602060405180830381865afa158015611739573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175d9190618213565b82828151811061176f5761176f6182df565b6001600160a01b03909216602092830291909101909101526117908161838a565b90506116cc565b6117c26040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611802573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118269190618213565b90506118536040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611883908b9089908990600401618477565b600060405180830381865afa1580156118a0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118c891908101906184c1565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906118fa908b908b908b90600401618578565b600060405180830381865afa158015611917573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261193f91908101906184c1565b6040820152856001600160401b0381111561195c5761195c617255565b60405190808252806020026020018201604052801561198f57816020015b606081526020019060019003908161197a5790505b50606082015260005b60ff8116871115611dce576000856001600160401b038111156119bd576119bd617255565b6040519080825280602002602001820160405280156119e6578160200160208202803683370190505b5083606001518360ff1681518110611a0057611a006182df565b602002602001018190525060005b86811015611cce5760008c6001600160a01b03166304ec63518a8a85818110611a3957611a396182df565b905060200201358e88600001518681518110611a5757611a576182df565b60200260200101516040518463ffffffff1660e01b8152600401611a949392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611ab1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ad591906185a1565b90506001600160c01b038116611b795760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610754565b8a8a8560ff16818110611b8e57611b8e6182df565b6001600160c01b03841692013560f81c9190911c600190811614159050611cbb57856001600160a01b031663dd9846b98a8a85818110611bd057611bd06182df565b905060200201358d8d8860ff16818110611bec57611bec6182df565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611c42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c6691906185ca565b85606001518560ff1681518110611c7f57611c7f6182df565b60200260200101518481518110611c9857611c986182df565b63ffffffff9092166020928302919091019091015282611cb78161838a565b9350505b5080611cc68161838a565b915050611a0e565b506000816001600160401b03811115611ce957611ce9617255565b604051908082528060200260200182016040528015611d12578160200160208202803683370190505b50905060005b82811015611d935784606001518460ff1681518110611d3957611d396182df565b60200260200101518181518110611d5257611d526182df565b6020026020010151828281518110611d6c57611d6c6182df565b63ffffffff9092166020928302919091019091015280611d8b8161838a565b915050611d18565b508084606001518460ff1681518110611dae57611dae6182df565b602002602001018190525050508080611dc6906185e7565b915050611998565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e339190618213565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611e66908b908b908e90600401618607565b600060405180830381865afa158015611e83573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611eab91908101906184c1565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611f05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f29919061827a565b611f455760405162461bcd60e51b815260040161075490618297565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611fb6929190618631565b600060405180830381865afa158015611fd3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ffb91908101906184c1565b9050600084516001600160401b0381111561201857612018617255565b604051908082528060200260200182016040528015612041578160200160208202803683370190505b50905060005b855181101561214257866001600160a01b03166304ec6351878381518110612071576120716182df565b60200260200101518786858151811061208c5761208c6182df565b60200260200101516040518463ffffffff1660e01b81526004016120c99392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156120e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061210a91906185a1565b6001600160c01b0316828281518110612125576121256182df565b60209081029190910101528061213a8161838a565b915050612047565b5095945050505050565b6000612358565b60008060008435602086013588821086169450888110851694508860038a8b858609850908898283099289526020890191909152149290921696604094850196509390940193505050565b60016020536021600090812092830682529190915260200190565b6000838351602085015b602087038110156121e5578781518309808452602093840193909250016121c3565b8781518309915060208352602080840152602060408401528160608401526002880360808401528760a084015260208360c08560055afa851693508251915085602081016020850394505b828110156122585789855185098a84518609908452601f199586019590945090920191612230565b8981518509945089825185099490915292909252509095945050505050565b81604052826060526000604060006080600060065afa9091169392505050565b816040526000604060006060600060075afa90911692915050565b8160c0528260e05260006040608080608060065afa9091169392505050565b8160c0526000604060806060608060075afa90911692915050565b81600052826020526107405160405261076051606052610780516080526107a05160a0528360c0528460e0526107c051610100526107e0516101205261080051610140526108205161016052600060206000610180600060085afa600051921690911695945050505050565b7f1df36213366afb618ed86862f1ee7b020c58e972bf9390420cfe9182467044e16105a081905260066105c08190526000919091526000805160206188b2833981519152907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190610da4358114610d406044351416600116906020610dc4610e845b808210156123fd57813580845286119490941693602092830192909101906123da565b5060649250610dc090506102645b8084101561242a5761241f87848688612153565b94509450945061240b565b5060009182208581068252909152602090810161010083015b808410156124625761245787848688612153565b945094509450612443565b5060009182208581068252909152602090810161247f858261219e565b905061020083015b808410156124a65761249b87848688612153565b945094509450612487565b5060009182208581068252909152602090810161014083015b808410156124de576124d387848688612153565b9450945094506124bf565b5060009182208581068252909152602090810161068083015b8084101561251a57833580845286119490941693602093840193909201916124f7565b50600091822085810682529091526020908101612537858261219e565b905061254586838587612153565b600090812088810690945292909252935091506020905061256885828486612153565b92509250925050507f1df36213366afb618ed86862f1ee7b020c58e972bf9390420cfe9182467044e16105a05260066105c05260076105e0527f300385d5fb6f3ce964dfa52b147e55ac6de38077e8c5fdb0215a31a8c8200001610600527f2822ef9d2d155c2b49f7a010aeec0dae3df9cff80535c8d08c9e954b942e6d6b610620527f1e3cef5e680760be385354b20cfa424b1545b57ee133143343ec8178b2e4ae8b610640527f2f0215559e251f1423df018c84fc874f074743154e03a1f96e57c953a337654c6106605260006106805260006106a05260006106c05260006106e0526001610700526002610720527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2610740527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed610760527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b610780527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6107a0527f186282957db913abd99f91db59fe69922e95040603ef44c0bd7aa3adeef8f5ac6107c0527f17944351223333f260ddc3b4af45191b856689eda9eab5cbcddbbe570ce860d26107e0527f06d971ff4a7467c3ec596ed6efc674572e32fd6f52b721f97e35b0b3d3546753610800527f06ecdb9f9567f59ed2eee36e1e1d58797fd13cc97fafc2910f5e8a12f202fa9a610820527f14ec4237a76469ad7c18091b19eedc2873dfebaecc486453d4e0598ba937f9b1610840527f282383a3c3a440a6798530b1a2e13ed3f63e21ea6e8f6f7c1eab753305db37d8610860527f21476ca20dc61723f34cc443f1acfa760975d5bbe43d36cfb66fd3610490eb1d610880527e4e4dc126d22bd5f1d5b5be3d206157123257735843514a3fc49ea82f3d0df96108a0527f28f450d0af8b274363e93a322a1743c2e28c6d9de317e8cfea3ad0b4be6806fd6108c0527f06fc092381131bb1ccf252cfeff132eefe5dd6f93bf91fda201e77cc051e8e836108e0527f164f02147ee406462971e8f31f436d0cb01c25ee7502129aa476e282f204375c610900527f11caf2aab38b9e43646075f44da61df35603e33128ca9953f1a543421b16704c610920527f164f02147ee406462971e8f31f436d0cb01c25ee7502129aa476e282f204375c610940527f11caf2aab38b9e43646075f44da61df35603e33128ca9953f1a543421b16704c610960527f25790aaef35f3e9299cbc026577ec1d681e406724a7fcc035259707b84cf86ec610980527f0437f8a7db0600965581ffcf1cf19693f8fe9356d77824f283e0f5f59cecde8f6109a0527f04679b415d5c67ce4c63a6a4cfe042113a360f58b8c4388e87d53c31f315ad5c6109c0527f1b9a8ea36466162eb7023e7dc84be27e6cd23d23f218c1362ee5c2f8f5115cda6109e0527f2f4c583af08c14d980f8a823e563695a47ccf4a862d0a6c30eaef37405ed845c610a00527f28e7d460bb7e7e551952c8a3584142498243088262b23742fd9346e2e556ac90610a20527f300cbf49cc37efd0c31add57fe897d63ed3c26aa87222f876107435e900e655c610a40527f0ab01181ce19bdfacff7935d5308290454a0678080f191c6a6b5fe80f637f8ba610a60527f300cbf49cc37efd0c31add57fe897d63ed3c26aa87222f876107435e900e655c610a80527f0ab01181ce19bdfacff7935d5308290454a0678080f191c6a6b5fe80f637f8ba610aa0526000610ac0526000610ae0526000610b00526000610b20527f055e17e3458e64b981f985c6f98ee847e6f4e28e175156437e97a564d350f940610b40527f2688f25934c56569410b1a8e9437fc3135284db7b92092c6d6a6c87b36ac2c22610b60527f12e0db55ccdbfef5019cc8f48cbe39401fc5b5ff962f7cd7780ae727b6c10bd9610b80527f03d75a7e667e932237c8563a3c717dc2b4cfc9170178c1a20e3e381c0297726a610ba0527f11c4ce711f4e784ff61059069d04080b0dbd61a6a1d0c314fb4032d861fd927f610bc0527f1d668de6d3721bf302bb552181b4a75f1dfe1336759bc64321970609d4f9795b610be0527f027b0aed4dbe9caaaf4964567a9d6644724055ed33c4767a116b780ec86fb331610c00527f25f781e7525ab98382eb2910b36961241fcb2dc2c38ac20daceb4fc98ee0c894610c20527f10f425414bdb525830d205811e549ec27fde006b3b70c75ccaeeeadd21800e42610c40527f232a17df37428ffa8ea275b3bf91bf05707a593a194cc90ec95819ef112c6345610c60527f2606c21158c0c13e85a06f8eb551e8af8962c7e1ce6dc35884a681fd307fa81b610c80527f24ab9a2bedb0fca61909c8752fa9e45c1ac83df32164cbaa9d98058a35269bdb610ca0527f07223791e8365e8d392205e91c17389816cefe8d14fc181d40dd3b506cef94de610cc0527f08242ae8b03f6df340be8cee34a24eb869b96f49503f8e58c9451e444e4c5806610ce0527f2e08d81ae6847749289cd5bd8b22b7ff0e08537d96e0252955109336f0bf97b2610d00527f2448ea4d5e34dea9033b5071587a032ec839348114259fc26f6a087c085f7710610d20527f29f4515f6dec507495c7ddffc3054927e5bb6c7d025b82f5a6b1bc738525075a610d40527f27111cfc5b00b4b013cee74e4d65be106d24bb8d8f8e3b2ffd76703e0492562a610d60527f25892cbf1dedad51f5eb616f8fea2e1524447e9b80c1b268d46b3e2e82ceb01f610d80527f1d8d76945d5ed1951385395b176979e1080a7d2aa64fe3f9aa12567b21b85fa7610da0526106805115612e25576106c0516106e05160206106a05102610dc4016020830260028102600382028335838501358386013583870135896020028801895b818a1015612ddc5760208a0199508935811b86019550888a0135811b85019450878a0135811b84019350868a0135811b830192508a81019050612d9e565b50508c60038e8f8788098709088d848509148b169a508c60038e8f8586098509088d828309148b169a5083610ec05282610ee05281610f005280610f2052505050505050505050505b80612e2f57600080fd5b6105e0519250610e40518060005b85811015612e5357848283099150600101612e3d565b50610620519450610f4060066105c0510160200281016105c051612e75576020015b610660515b81831015612e9b578681880386088352868882099050602083019250612e7a565b50856001870384089350838152612eb98660208301610f40886121b9565b9450610f409150856106005185099350610660515b81831015612ef15786878285510986098352868882099050602083019250612ece565b50610f60519650610f80935061100091505b81841015612f1d5785845188089650602084019350612f03565b60009150610dc46105c05160200281015b80821015612f52578788833588510985089350602082019150602086019550612f2e565b505051610f4080516110005194909152610f6091909152610f8052610fa095909552610fc05250610fe092909252610e2051600092907f09226b6e22c6f0ca64ec26aad4c86e715b5f898e5e963f25870e56bbe533e9a2906108c435600281860386818308915086828409925060039150868183089150868284099250600491508681830891505085818309915050610724356106a4356106e4358781830888039150879050818308915050858183099650506108e4359050600281860386818308915086828409925060039150868183089150868284099250600491508681830891505085818309915050610744356106c43561070435878183088803915087905081830891505085818309915050848186848909089550506108c435600181860386818308915086828409925060029150868183089150868284099250600491508681830891505085818309915050610724356106a4356106e435878183098803915087905081830891505085818309915050848186848909089550506108e435600181860386818308915086828409925060029150868183089150868284099250600491508681830891505085818309915050610744356106c43561070435878183098803915087905081830891505085818309915050848186848909089550506108c435600181860386818308915086828409925060039150868183089150868284099250600491508681830891505085818309915050610724356106a4356106e4358703878183088803915087905081830891505085818309915050848186848909089550506108e435600181860386818308915086828409925060039150868183089150868284099250600491508681830891505085818309915050610744356106c435610704358703878183088803915087905081830891505085818309915050848186848909089550506108c43560018186038681830887818509935050600287828208905087818509935050600387828208915050868184099250508503610724358682820891508682820991505085818309915050848186848909089550506108e4356001818603868183088781850993505060028782820890508781850993505060038782820891505086818409925050850361074435868282089150868282099150508581830991505084818684890908955050610904356001818603868183089150868284099250600391508681830891505085818309915050610724356107a43560006106a4356106e43589818309915050888183089150506106c435610704358981830991505088818308915050878183088803915087905081830891505085818309915050848186848909089550506109043560028186038681830891508682840992506003915086818308915050858183099150506107243560006106a4356106e43588818309915050878183089150506106c4356107043588818309915050878183088803915087905081830891505085818309915050848186848909089550506109243560028186038681830891508682840992506003915086818308915050858183099150506107243560016106e435878183099150506107043587818309880391508790508183089150508581830991505084818684890908955050610904356001818603868183088781850993505060028782820891505086818409925050610724356107a4356106e435888185099350506107043588818509935050878382098803925087905082820891505085818309915050848186848909089550506109243560018186038681830891508682840992506002915086818308915050858183099150506107243560006106e435878183089150506107043587818308880391508790508183089150508581830991505084818684890908955050610924356001818603868183089150868284099250600391508681830891505085818309915050610724356107a43560006106e43588818308915050610704358881830891505087818308880391508790508183089150508581830991505084818684890908955050610fc0518485610aa435830986038208905084818684890908955050610b643584858287038784850908610f8051099050848186848909089550508384610ae4358603610b043508610fc05109848186848909089550508384610b44358603610b643508610fc0510984818684890908955050610e0051610de051610ac435610aa4358788858a8b6109643588096106a4350808830991508788858a8b6109843588096106c4350808830991508788858a8b6109a43588096106e4350808830991508788858a8b6109c43588096107043508088309915087610e405184096000528788858a6000516106a4350808820990508786600051096000528788858a6000516106c4350808820990508786600051096000528788858a6000516106e4350808820990508786600051096000528788858a600051610704350808820990508786600051096000528781890383089350505050848586610fa051610f805108830986038208905084818684890908955050610e0051610de051610b2435610b04358788858a8b6109e4358809610724350808830991508788858a8b610a04358809610744350808830991508788858a8b610a24358809610764350808830991508788858a8b610a44358809610784350808830991508788858a600051610724350808820990508786600051096000528788858a600051610744350808820990508786600051096000528788858a600051610764350808820990508786600051096000528788858a600051610784350808820990508786600051096000528781890383089350505050848586610fa051610f805108830986038208905084818684890908955050610e0051610de051610b8435610b64358788858a8b610a643588096107c4350808830991508788858a8b610a84358809610fe0510808830991508788858a6000516107c4350808820990508786600051096000528788858a600051610fe0510808820990508781890383089350505050848586610fa051610f805108830986038208905084818684890908955050610fc05184610ba4358209905084818684890908955050610f805184610ba4358209905084818684890908955050610dc051610de051600060016108643588818309915061076435898282099050610784358a83820992505082935089818b8887090893505088818a878609089250505086828208905060006001610884358981830991506106a4358a82820990506106e4358b8382099250508293508a818c8987090893505089818b888609089250505087838208600094509192508291905080888082610be435098a03850893508889610ba4358b03610bc435088a838609099450505050858682880384088788610f8051610fa0510889036001080991505084818684890908955050610fc05184610c04358209905084818684890908955050610f805184610c04358209905084818684890908955050610dc051610de051600060016108643588818309915061076435898282099050610784358a83820992505082935089818b8887090893505088818a8786090892505050868282089050600060016108a4358981830991506106c4358a8282099050610704358b8382099250508293508a818c8987090893505089818b888609089250505087838208600094509192508291905080888082610c4435098a03850893508889610c04358b03610c2435088a838609099450505050858682880384088788610f8051610fa0510889036001080991505084818684890908955050610fc05184610c64358209905084818684890908955050610f805184610c64358209905084818684890908610dc051610de05191975091506107e4358061080435888181878509089250505086828208905060006108243560018981830991506106a4358a8184099050828b038b818408905060008c81830990508c818408925050610724358c8186099450508b8382099250508a8284089250809350505088818a8785090891505087838208600094509192508291905080888082610ca435098a03850893508889610c64358b03610c8435088a838609099450505050858682880384088788610f8051610fa0510889036001080991505084818684890908955050610fc05184610cc4358209905084818684890908955050610f805184610cc4358209905084818684890908610dc051610de05191975091506107e4358061080435888181878509089250505086828208905060006108443560018981830991506106c4358a8184099050828b038b818408905060008c81830990508c818408925050610744358c8186099450508b8382099250508a8284089250809350505088818a8785090891505087838208600094509192508291905080888082610d0435098a03850893508889610cc4358b03610ce435088a838609099450505050858682880384088788610f8051610fa0510889036001080991505084818684890908955050505081610f60518409611000526106643560005261068435602052610f405192506106246105245b81811015613c9a57613c7b8584612297565b9250613c8d6020830135833585612277565b9250604082039150613c69565b50506000516110205260205161104052610e40519250610620516106405183828609915081610360528461034052838186099150816103205283818309915083818309915083818309915083818309915083818309915050806103005250610ea05192506103806104006103005b81831015613d255784815186038708835260209283019201613d08565b50506103c0516104005250610380516103a051909350829084099250816103e05184099250826104205282600052610380519250816103e0518409610440526103a05161046081905261038051935082908409610480526103c0516001935082908409602052610320516103405190935060008382810386089050836103a05182099050806040528385850383089050836103c05182096060525050610300516103405161036051919450906000848381038708905084858387038808820990508461038051820990508060805284868603840890508485838703850882099050846103c051820990508060a05284868603830890508485848703840882099050846103e051820960c05250506103405161036051909450905060008382810386089050836103c051820990508060e0528385850383089050836103e05182096101005250613e7b9050826101206000846121b9565b9050600051925082610420526104406104a05b80821015613ea85783858351098252602082019150613e8e565b50506020519250610e60516000838461094435870982089050838282099050838461100051870982089050610a846109445b81811015613efa5785868335890987868609089250602082039150613eda565b50506109246107a45b81811015613f235785868335890987868609089250602082039150613f03565b50508382820990508384610d04358709820890508382820990508384610ca4358709820890508382820990508384610c44358709820890508382820990508384610be4358709820890506107846107245b81811015613f945785868335890987868609089250602082039150613f74565b50506107046106845b81811015613fbd5785868335890987868609089250602082039150613f9d565b50506104a0525060405160009350829081906107a4359009840892508182610724356060510984089250816104405184099250826104c052610e6051925060008283610b443560805109820890508283610b043560a05109820890508283610b243560c05109820890508284820990508283610ae43560805109820890508283610aa43560a05109820890508283610ac43560c0510982089050826104605182099050806104e05250610e6051925060008283610cc43560e05109820890508283610ce4356101005109820890508284820990508283610c643560e05109820890508283610c84356101005109820890508284820990508283610c043560e05109820890508283610c24356101005109820890508284820990508283610ba43560e05109820890508283610bc4356101005109820890508284820990508283610b643560e05109820890508283610b843561010051098208905082610480518209610500525060205161052052604051606051909350829084086105405260805160a0519093508290840892508160c05184086105605260e0516101005190935082908408610580526000925060806105205b81851015614188578051855260209485019401614170565b50506141988260806000846121b9565b90508161050051606051099250604060806104e05b818310156141d65784610e8051870995508485825185510987089550601f1992830192016141ad565b5050508261106052610e8051925061052435600052610544356020526141ff610e605182612297565b9050614212611040516110205183612277565b9050610d806108005b8181101561424f57614230610e605184612297565b92506142426020830151835185612277565b925060408203915061421b565b50506103246101645b8181101561428c5761426d610e605184612297565b925061427f6020830135833585612277565b9250604082039150614258565b505061012460245b818110156142c8576142a9610e605184612297565b92506142bb6020830135833585612277565b9250604082039150614294565b5050610164356080526101843560a0526142e882610440518509826122d1565b90506142f960a05160805183612277565b905081610e8051840992506103a4356080526103c43560a05261431f610e6051826122d1565b90506143326103843561036435836122b2565b905061434482610460518509826122d1565b905061435560a05160805183612277565b905081610e8051840992506104e4356080526105043560a0526104a46103a45b818110156143a95761438a610e6051846122d1565b925061439c60208301358335856122b2565b9250604082039150614375565b50506143bb82610480518509826122d1565b90506143cc60a05160805183612277565b9050610700516080526107205160a0526143eb611060518303826122d1565b90506143fc60a05160805183612277565b9050610d2435608052610d443560a05261441b610400518303826122d1565b905061442c60a05160805183612277565b9050610d6435608052610d843560a052614449610ea051826122d1565b905061445a60a05160805183612277565b9050600051611080526020516110a052610d64356110c052610d84356110e052610680511561452e57610ec051600052610ee051602052610f0051604052610f2051606052611080516080526110a05160a0526110c05160c0526110e05160e052816101006000200692506144cf8382612297565b90506144e26110a0516110805183612277565b9050600051611080526020516110a052610f0051600052610f205160205261450a8382612297565b905061451d6110e0516110c05183612277565b90506000516110c0526020516110e0525b6145476110e0516110c0516110a05161108051856122ec565b90508061455357600080fd5b505050600160005292915050565b60006145de83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525063ffffffff8a16815260ce602052604090819020815160c0810192839052945092506006915082845b8154815260200190600101908083116145c157505050505061214c565b90506001811515141561463f57604051339063ffffffff8616907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a363ffffffff8416600090815260cc60205260409020805460ff191660031790555b50505050565b60408051808201909152606080825260208201526000846146bc5760405162461bcd60e51b815260206004820152603760248201526000805160206188d283398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610754565b604083015151851480156146d4575060a08301515185145b80156146e4575060c08301515185145b80156146f4575060e08301515185145b61475e5760405162461bcd60e51b815260206004820152604160248201526000805160206188d283398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610754565b825151602084015151146147d65760405162461bcd60e51b8152602060048201526044602482018190526000805160206188d2833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610754565b4363ffffffff168463ffffffff16106148455760405162461bcd60e51b815260206004820152603c60248201526000805160206188d283398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610754565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561488657614886617255565b6040519080825280602002602001820160405280156148af578160200160208202803683370190505b506020820152866001600160401b038111156148cd576148cd617255565b6040519080825280602002602001820160405280156148f6578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561492a5761492a617255565b604051908082528060200260200182016040528015614953578160200160208202803683370190505b5081526020860151516001600160401b0381111561497357614973617255565b60405190808252806020026020018201604052801561499c578160200160208202803683370190505b5081602001819052506000614a6e8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015614a45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a699190618650565b61686f565b905060005b876020015151811015614d0957614ab888602001518281518110614a9957614a996182df565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110614ace57614ace6182df565b60209081029190910101528015614b8e576020830151614aef60018361866d565b81518110614aff57614aff6182df565b602002602001015160001c83602001518281518110614b2057614b206182df565b602002602001015160001c11614b8e576040805162461bcd60e51b81526020600482015260248101919091526000805160206188d283398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610754565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110614bd357614bd36182df565b60200260200101518b8b600001518581518110614bf257614bf26182df565b60200260200101516040518463ffffffff1660e01b8152600401614c2f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015614c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c7091906185a1565b6001600160c01b031683600001518281518110614c8f57614c8f6182df565b602002602001018181525050614cf5610b9c614cc98486600001518581518110614cbb57614cbb6182df565b602002602001015116616902565b8a602001518481518110614cdf57614cdf6182df565b602002602001015161692d90919063ffffffff16565b945080614d018161838a565b915050614a73565b5050614d1483616a11565b60975490935060ff16600081614d2b576000614dad565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015614d89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614dad91906183a5565b905060005b8a81101561542b578215614f0d578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110614e0957614e096182df565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015614e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614e6d91906183a5565b614e779190618684565b11614f0d5760405162461bcd60e51b815260206004820152606660248201526000805160206188d283398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610754565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110614f4e57614f4e6182df565b9050013560f81c60f81b60f81c8c8c60a001518581518110614f7257614f726182df565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015614fce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ff2919061869c565b6001600160401b0319166150158a604001518381518110614a9957614a996182df565b67ffffffffffffffff1916146150b15760405162461bcd60e51b815260206004820152606160248201526000805160206188d283398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610754565b6150e1896040015182815181106150ca576150ca6182df565b60200260200101518761646790919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110615124576151246182df565b9050013560f81c60f81b60f81c8c8c60c001518581518110615148576151486182df565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156151a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151c8919061844e565b856020015182815181106151de576151de6182df565b6001600160601b0390921660209283029190910182015285015180518290811061520a5761520a6182df565b602002602001015185600001518281518110615228576152286182df565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015615416576152a086600001518281518110615272576152726182df565b60200260200101518f8f8681811061528c5761528c6182df565b600192013560f81c9290921c811614919050565b15615404577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106152e6576152e66182df565b9050013560f81c60f81b60f81c8e8960200151858151811061530a5761530a6182df565b60200260200101518f60e001518881518110615328576153286182df565b60200260200101518781518110615341576153416182df565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156153a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906153c9919061844e565b87518051859081106153dd576153dd6182df565b602002602001018181516153f191906186c7565b6001600160601b03169052506001909101905b8061540e8161838a565b91505061524c565b505080806154239061838a565b915050614db2565b5050506000806154458c868a606001518b60800151610a81565b91509150816154b65760405162461bcd60e51b815260206004820152604360248201526000805160206188d283398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610754565b806155175760405162461bcd60e51b815260206004820152603960248201526000805160206188d283398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610754565b505060008782602001516040516020016155329291906186ef565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b615564616aac565b61556e6000616b06565b565b60ce602052816000526040600020816006811061558c57600080fd5b01549150829050565b60006155a46020850185617553565b9050600163ffffffff8216600090815260cc602052604090205460ff1660038111156155d2576155d2618085565b1461562e5760405162461bcd60e51b815260206004820152602660248201527f54686520726573706f6e7365206973206e6f7420696e20612077616974696e6760448201526520737461746560d01b6064820152608401610754565b63ffffffff8116600090815260cd602052604090205461565090606490618684565b43116156b45760405162461bcd60e51b815260206004820152602d60248201527f546865206f70657261746f72207374696c6c206861732074696d6520746f207360448201526c3ab136b4ba103090383937b7b360991b6064820152608401610754565b600082516001600160401b038111156156cf576156cf617255565b6040519080825280602002602001820160405280156156f8578160200160208202803683370190505b50905060005b835181101561574b5761571c848281518110614a9957614a996182df565b82828151811061572e5761572e6182df565b6020908102919091010152806157438161838a565b9150506156fe565b50600061575e60c0880160a08901617553565b826040516020016157709291906186ef565b6040516020818303038152906040528051906020012090508460200135811461581a5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610754565b600084516001600160401b0381111561583557615835617255565b60405190808252806020026020018201604052801561585e578160200160208202803683370190505b50905060005b8551811015615951577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106158ae576158ae6182df565b60200260200101516040518263ffffffff1660e01b81526004016158d491815260200190565b602060405180830381865afa1580156158f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906159159190618213565b828281518110615927576159276182df565b6001600160a01b0390921660209283029190910190910152806159498161838a565b915050615864565b5063ffffffff8416600081815260cc6020526040808220805460ff19166002179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505050565b60cf546001600160a01b031633146159fe5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610754565b6000615a1060c0850160a08601617553565b9050366000615a2260c0870187618737565b90925090506000615a3a610100880160e08901617553565b905060ca6000615a4d6020890189617553565b63ffffffff1663ffffffff1681526020019081526020016000205487604051602001615a79919061877d565b6040516020818303038152906040528051906020012014615b025760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610754565b600060cb81615b1460208a018a617553565b63ffffffff1663ffffffff1681526020019081526020016000205414615b915760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610754565b615bbb7f0000000000000000000000000000000000000000000000000000000000000000856181eb565b63ffffffff164363ffffffff161115615c2c5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610754565b600086604051602001615c3f9190618826565b604051602081830303815290604052805190602001209050600080615c678387878a8c614645565b9150915060005b85811015615d66578460ff1683602001518281518110615c9057615c906182df565b6020026020010151615ca29190618834565b6001600160601b0316606484600001518381518110615cc357615cc36182df565b60200260200101516001600160601b0316615cde919061836b565b1015615d54576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610754565b80615d5e8161838a565b915050615c6e565b5060408051808201825263ffffffff43168152602080820184905291519091615d93918c91849101618863565b6040516020818303038152906040528051906020012060cb60008c6000016020810190615dc09190617553565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a82604051615e0f929190618863565b60405180910390a15050505050505050505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110615e5f57615e5f6182df565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90615e9b9088908690600401618631565b600060405180830381865afa158015615eb8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615ee091908101906184c1565b600081518110615ef257615ef26182df565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015615f5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615f8291906185a1565b6001600160c01b031690506000615f9882616b58565b905081615fa68a838a611077565b9550955050505050935093915050565b615fbe616aac565b6001600160a01b0381166160235760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610754565b61093f81616b06565b600054610100900460ff161580801561604c5750600054600160ff909116105b806160665750303b158015616066575060005460ff166001145b6160c95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610754565b6000805460ff1916600117905580156160ec576000805461ff0019166101001790555b6160f7856000616c24565b61610084616b06565b60cf80546001600160a01b038086166001600160a01b03199283161790925560d08054928516929091169190911790558015616176576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156161d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906161f49190618213565b6001600160a01b0316336001600160a01b0316146162245760405162461bcd60e51b815260040161075490618230565b6066541981196066541916146162a25760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610754565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610a76565b6001600160a01b0381166163675760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610754565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526163ec61704f565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561641f57616421565bfe5b508061645f5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610754565b505092915050565b604080518082019091526000808252602082015261648361706d565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561641f57508061645f5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610754565b61650361708b565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806165eb6000805160206188b2833981519152866182f5565b90505b6165f781616d0e565b90935091506000805160206188b2833981519152828309831415616631576040805180820190915290815260208101919091529392505050565b6000805160206188b28339815191526001820890506165ee565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061667d6170b0565b60005b600281101561684257600061669682600661836b565b90508482600281106166aa576166aa6182df565b602002015151836166bc836000618684565b600c81106166cc576166cc6182df565b60200201528482600281106166e3576166e36182df565b602002015160200151838260016166fa9190618684565b600c811061670a5761670a6182df565b6020020152838260028110616721576167216182df565b6020020151515183616734836002618684565b600c8110616744576167446182df565b602002015283826002811061675b5761675b6182df565b6020020151516001602002015183616774836003618684565b600c8110616784576167846182df565b602002015283826002811061679b5761679b6182df565b6020020151602001516000600281106167b6576167b66182df565b6020020151836167c7836004618684565b600c81106167d7576167d76182df565b60200201528382600281106167ee576167ee6182df565b602002015160200151600160028110616809576168096182df565b60200201518361681a836005618684565b600c811061682a5761682a6182df565b6020020152508061683a8161838a565b915050616680565b5061684b6170cf565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008061687b84616d90565b9050808360ff166001901b116168f95760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610754565b90505b92915050565b6000805b82156168fc5761691760018461866d565b90921691806169258161888f565b915050616906565b60408051808201909152600080825260208201526102008261ffff16106169895760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610754565b8161ffff166001141561699d5750816168fc565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610616a0657600161ffff871660ff83161c811614156169e9576169e68484616467565b93505b6169f38384616467565b92506201fffe600192831b1691016169b9565b509195945050505050565b60408051808201909152600080825260208201528151158015616a3657506020820151155b15616a54575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206188b28339815191528460200151616a8791906182f5565b616a9f906000805160206188b283398151915261866d565b905292915050565b919050565b6033546001600160a01b0316331461556e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610754565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080616b6684616902565b61ffff166001600160401b03811115616b8157616b81617255565b6040519080825280601f01601f191660200182016040528015616bab576020820181803683370190505b5090506000805b825182108015616bc3575061010081105b15616c1a576001811b935085841615616c0a578060f81b838381518110616bec57616bec6182df565b60200101906001600160f81b031916908160001a9053508160010191505b616c138161838a565b9050616bb2565b5090949350505050565b6065546001600160a01b0316158015616c4557506001600160a01b03821615155b616cc75760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610754565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2616d0a826162d9565b5050565b600080806000805160206188b283398151915260036000805160206188b2833981519152866000805160206188b2833981519152888909090890506000616d84827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206188b2833981519152616f1d565b91959194509092505050565b600061010082511115616e195760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610754565b8151616e2757506000919050565b60008083600081518110616e3d57616e3d6182df565b0160200151600160f89190911c81901b92505b8451811015616f1457848181518110616e6b57616e6b6182df565b0160200151600160f89190911c1b9150828211616f005760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610754565b91811791616f0d8161838a565b9050616e50565b50909392505050565b600080616f286170cf565b616f30616ff3565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561641f575082616fba5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610754565b505195945050505050565b6040518060800160405280616fd86170ed565b81526000602082018190526060604083018190529091015290565b6040518060c001604052806006906020820280368337509192915050565b826006810192821561703f579160200282015b8281111561703f578251825591602001919060010190617024565b5061704b92915061710b565b5090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061709e617120565b81526020016170ab617120565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b5b8082111561704b576000815560010161710c565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461093f57600080fd5b8035616aa78161713e565b60008083601f84011261716d57600080fd5b5081356001600160401b0381111561718457600080fd5b60208301915083602082850101111561719c57600080fd5b9250929050565b60008060008060e085870312156171b957600080fd5b60a08501868111156171ca57600080fd5b859450356171d78161713e565b925060c08501356001600160401b038111156171f257600080fd5b6171fe8782880161715b565b95989497509550505050565b6001600160a01b038116811461093f57600080fd5b60006020828403121561723157600080fd5b81356168f98161720a565b60006020828403121561724e57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561728d5761728d617255565b60405290565b60405161010081016001600160401b038111828210171561728d5761728d617255565b604051601f8201601f191681016001600160401b03811182821017156172de576172de617255565b604052919050565b6000604082840312156172f857600080fd5b61730061726b565b9050813581526020820135602082015292915050565b600082601f83011261732757600080fd5b604051604081018181106001600160401b038211171561734957617349617255565b806040525080604084018581111561736057600080fd5b845b81811015616a06578035835260209283019201617362565b60006080828403121561738c57600080fd5b61739461726b565b90506173a08383617316565b81526173af8360408401617316565b602082015292915050565b60008060008061012085870312156173d157600080fd5b843593506173e286602087016172e6565b92506173f1866060870161737a565b91506174008660e087016172e6565b905092959194509250565b6000610100828403121561741e57600080fd5b50919050565b60006040828403121561741e57600080fd5b60006001600160401b0382111561744f5761744f617255565b5060051b60200190565b600082601f83011261746a57600080fd5b8135602061747f61747a83617436565b6172b6565b82815260069290921b8401810191818101908684111561749e57600080fd5b8286015b848110156174c2576174b488826172e6565b8352918301916040016174a2565b509695505050505050565b60008060008060c085870312156174e357600080fd5b84356001600160401b03808211156174fa57600080fd5b6175068883890161740b565b95506175158860208901617424565b94506175248860608901617424565b935060a087013591508082111561753a57600080fd5b5061754787828801617459565b91505092959194509250565b60006020828403121561756557600080fd5b81356168f98161713e565b6000806040838503121561758357600080fd5b823561758e8161720a565b91506020838101356001600160401b038111156175aa57600080fd5b8401601f810186136175bb57600080fd5b80356175c961747a82617436565b81815260059190911b820183019083810190888311156175e857600080fd5b928401925b8284101561760f5783356176008161720a565b825292840192908401906175ed565b80955050505050509250929050565b600081518084526020808501945080840160005b8381101561764e57815187529582019590820190600101617632565b509495945050505050565b60208152600061766c602083018461761e565b9392505050565b600082601f83011261768457600080fd5b81356001600160401b0381111561769d5761769d617255565b6176b0601f8201601f19166020016172b6565b8181528460208386010111156176c557600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000606084860312156176f757600080fd5b83356177028161720a565b925060208401356001600160401b0381111561771d57600080fd5b61772986828701617673565b925050604084013561773a8161713e565b809150509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156177db578385038a52825180518087529087019087870190845b818110156177c657835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101617782565b50509a87019a95505091850191600101617764565b509298975050505050505050565b60208152600061766c6020830184617745565b801515811461093f57600080fd5b60006020828403121561781c57600080fd5b81356168f9816177fc565b600082601f83011261783857600080fd5b8135602061784861747a83617436565b82815260059290921b8401810191818101908684111561786757600080fd5b8286015b848110156174c2578035835291830191830161786b565b6000806040838503121561789557600080fd5b82356178a08161720a565b915060208301356001600160401b038111156178bb57600080fd5b6178c785828601617827565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b818110156179125783516001600160a01b0316835292840192918401916001016178ed565b50909695505050505050565b6000806000806000806080878903121561793757600080fd5b86356179428161720a565b955060208701356179528161713e565b945060408701356001600160401b038082111561796e57600080fd5b61797a8a838b0161715b565b9096509450606089013591508082111561799357600080fd5b818901915089601f8301126179a757600080fd5b8135818111156179b657600080fd5b8a60208260051b85010111156179cb57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561764e57815163ffffffff16875295820195908201906001016179f5565b600060208083528351608082850152617a3360a08501826179e1565b905081850151601f1980868403016040870152617a5083836179e1565b92506040870151915080868403016060870152617a6d83836179e1565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015617ac45784878303018452617ab28287516179e1565b95880195938801939150600101617a98565b509998505050505050505050565b60ff8116811461093f57600080fd5b600060208284031215617af357600080fd5b81356168f981617ad2565b600080600060608486031215617b1357600080fd5b8335617b1e8161720a565b925060208401356001600160401b03811115617b3957600080fd5b61772986828701617827565b6020808252825182820181905260009190848201906040850190845b8181101561791257835183529284019291840191600101617b61565b60008060e08385031215617b9057600080fd5b82356001600160401b0380821115617ba757600080fd5b617bb386838701617673565b93506020915085603f860112617bc857600080fd5b60405160c081018181108382111715617be357617be3617255565b60405290508060e0860187811115617bfa57600080fd5b8387015b81811015617c155780358452928401928401617bfe565b509497909650945050505050565b600080600060408486031215617c3857600080fd5b8335617c438161713e565b925060208401356001600160401b03811115617c5e57600080fd5b617c6a8682870161715b565b9497909650939450505050565b600082601f830112617c8857600080fd5b81356020617c9861747a83617436565b82815260059290921b84018101918181019086841115617cb757600080fd5b8286015b848110156174c2578035617cce8161713e565b8352918301918301617cbb565b600082601f830112617cec57600080fd5b81356020617cfc61747a83617436565b82815260059290921b84018101918181019086841115617d1b57600080fd5b8286015b848110156174c25780356001600160401b03811115617d3e5760008081fd5b617d4c8986838b0101617c77565b845250918301918301617d1f565b60006101808284031215617d6d57600080fd5b617d75617293565b905081356001600160401b0380821115617d8e57600080fd5b617d9a85838601617c77565b83526020840135915080821115617db057600080fd5b617dbc85838601617459565b60208401526040840135915080821115617dd557600080fd5b617de185838601617459565b6040840152617df3856060860161737a565b6060840152617e058560e086016172e6565b6080840152610120840135915080821115617e1f57600080fd5b617e2b85838601617c77565b60a0840152610140840135915080821115617e4557600080fd5b617e5185838601617c77565b60c0840152610160840135915080821115617e6b57600080fd5b50617e7884828501617cdb565b60e08301525092915050565b600080600080600060808688031215617e9c57600080fd5b8535945060208601356001600160401b0380821115617eba57600080fd5b617ec689838a0161715b565b909650945060408801359150617edb8261713e565b90925060608701359080821115617ef157600080fd5b50617efe88828901617d5a565b9150509295509295909350565b600081518084526020808501945080840160005b8381101561764e5781516001600160601b031687529582019590820190600101617f1f565b6040815260008351604080840152617f5f6080840182617f0b565b90506020850151603f19848303016060850152617f7c8282617f0b565b925050508260208301529392505050565b60008060408385031215617fa057600080fd5b8235617fab8161713e565b946020939093013593505050565b600080600060808486031215617fce57600080fd5b83356001600160401b0380821115617fe557600080fd5b617ff18783880161740b565b94506180008760208801617424565b9350606086013591508082111561801657600080fd5b5061802386828701617d5a565b9150509250925092565b60008060006060848603121561804257600080fd5b833561804d8161720a565b925060208401359150604084013561773a8161713e565b82815260406020820152600061807d6040830184617745565b949350505050565b634e487b7160e01b600052602160045260246000fd5b60208101600483106180bd57634e487b7160e01b600052602160045260246000fd5b91905290565b600080600080608085870312156180d957600080fd5b84356180e48161720a565b935060208501356180f48161720a565b925060408501356181048161720a565b915060608501356181148161720a565b939692955090935050565b6020808252825160009190828483015b600582101561814e57825181529183019160019190910190830161812f565b50505063ffffffff818501511660c084015260408401516101008060e086015281518061012087015260005b81811015618197578381018501518782016101400152840161817a565b818111156181aa57600061014083890101525b50606087015163ffffffff8116878401529350601f01601f1916949094016101400195945050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff80831681851680830382111561820a5761820a6181d5565b01949350505050565b60006020828403121561822557600080fd5b81516168f98161720a565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561828c57600080fd5b81516168f9816177fc565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261831257634e487b7160e01b600052601260045260246000fd5b500690565b80356183228161713e565b63ffffffff168252602090810135910152565b608081016183438285618317565b823561834e8161713e565b63ffffffff16604083015260209290920135606090910152919050565b6000816000190483118215151615618385576183856181d5565b500290565b600060001982141561839e5761839e6181d5565b5060010190565b6000602082840312156183b757600080fd5b5051919050565b600060208083850312156183d157600080fd5b82516001600160401b038111156183e757600080fd5b8301601f810185136183f857600080fd5b805161840661747a82617436565b81815260059190911b8201830190838101908783111561842557600080fd5b928401925b828410156184435783518252928401929084019061842a565b979650505050505050565b60006020828403121561846057600080fd5b81516001600160601b03811681146168f957600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156184a457600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156184d457600080fd5b82516001600160401b038111156184ea57600080fd5b8301601f810185136184fb57600080fd5b805161850961747a82617436565b81815260059190911b8201830190838101908783111561852857600080fd5b928401925b828410156184435783516185408161713e565b8252928401929084019061852d565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff8416815260406020820152600061859860408301848661854f565b95945050505050565b6000602082840312156185b357600080fd5b81516001600160c01b03811681146168f957600080fd5b6000602082840312156185dc57600080fd5b81516168f98161713e565b600060ff821660ff8114156185fe576185fe6181d5565b60010192915050565b60408152600061861b60408301858761854f565b905063ffffffff83166020830152949350505050565b63ffffffff8316815260406020820152600061807d604083018461761e565b60006020828403121561866257600080fd5b81516168f981617ad2565b60008282101561867f5761867f6181d5565b500390565b60008219821115618697576186976181d5565b500190565b6000602082840312156186ae57600080fd5b815167ffffffffffffffff19811681146168f957600080fd5b60006001600160601b03838116908316818110156186e7576186e76181d5565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561872a5781518552938201939082019060010161870e565b5092979650505050505050565b6000808335601e1984360301811261874e57600080fd5b8301803591506001600160401b0382111561876857600080fd5b60200191503681900382131561719c57600080fd5b6020815260a0826020830137600060c08201600080825260a08501356187a28161713e565b63ffffffff1690915260c08401359036859003601e190182126187c3578081fd5b9084019081356001600160401b038111156187dc578182fd5b8036038613156187ea578182fd5b61010091508160e08601526188076101208601826020860161854f565b92505061881660e08601617150565b63ffffffff811685830152616c1a565b604081016168fc8284618317565b60006001600160601b038083168185168183048111821515161561885a5761885a6181d5565b02949350505050565b608081016188718285618317565b63ffffffff8351166040830152602083015160608301529392505050565b600061ffff808316818114156188a7576188a76181d5565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220579744483c9ab13c84d416b1aa8efb3f9b12922ba4cdccdb6d6bf17a2a30ec4d64736f6c634300080c0033",
}

// ContractIncredibleSquaringTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.ABI instead.
var ContractIncredibleSquaringTaskManagerABI = ContractIncredibleSquaringTaskManagerMetaData.ABI

// ContractIncredibleSquaringTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.Bin instead.
var ContractIncredibleSquaringTaskManagerBin = ContractIncredibleSquaringTaskManagerMetaData.Bin

// DeployContractIncredibleSquaringTaskManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSquaringTaskManager to it.
func DeployContractIncredibleSquaringTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractIncredibleSquaringTaskManager, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSquaringTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSquaringTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManager struct {
	ContractIncredibleSquaringTaskManagerCaller     // Read-only binding to the contract
	ContractIncredibleSquaringTaskManagerTransactor // Write-only binding to the contract
	ContractIncredibleSquaringTaskManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSquaringTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSquaringTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSquaringTaskManagerSession struct {
	Contract     *ContractIncredibleSquaringTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSquaringTaskManagerCallerSession struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSquaringTaskManagerTransactorSession struct {
	Contract     *ContractIncredibleSquaringTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerRaw struct {
	Contract *ContractIncredibleSquaringTaskManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCallerRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactorRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSquaringTaskManager creates a new instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSquaringTaskManager, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSquaringTaskManagerCaller creates a new read-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSquaringTaskManagerCaller, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerTransactor creates a new write-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSquaringTaskManagerTransactor, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerFilterer creates a new log filterer instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSquaringTaskManagerFilterer, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSquaringTaskManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSquaringTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskChallengeState is a free data retrieval call binding the contract method 0xe7a6c1fa.
//
// Solidity: function taskChallengeState(uint32 ) view returns(uint8)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskChallengeState(opts *bind.CallOpts, arg0 uint32) (uint8, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskChallengeState", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TaskChallengeState is a free data retrieval call binding the contract method 0xe7a6c1fa.
//
// Solidity: function taskChallengeState(uint32 ) view returns(uint8)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskChallengeState(arg0 uint32) (uint8, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskChallengeState(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskChallengeState is a free data retrieval call binding the contract method 0xe7a6c1fa.
//
// Solidity: function taskChallengeState(uint32 ) view returns(uint8)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskChallengeState(arg0 uint32) (uint8, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskChallengeState(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskChallengedBlock is a free data retrieval call binding the contract method 0xb030b7b4.
//
// Solidity: function taskChallengedBlock(uint32 ) view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskChallengedBlock(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskChallengedBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskChallengedBlock is a free data retrieval call binding the contract method 0xb030b7b4.
//
// Solidity: function taskChallengedBlock(uint32 ) view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskChallengedBlock(arg0 uint32) (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskChallengedBlock(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskChallengedBlock is a free data retrieval call binding the contract method 0xb030b7b4.
//
// Solidity: function taskChallengedBlock(uint32 ) view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskChallengedBlock(arg0 uint32) (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskChallengedBlock(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskChallengedInstances is a free data retrieval call binding the contract method 0x7cfd8968.
//
// Solidity: function taskChallengedInstances(uint32 , uint256 ) view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskChallengedInstances(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskChallengedInstances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskChallengedInstances is a free data retrieval call binding the contract method 0x7cfd8968.
//
// Solidity: function taskChallengedInstances(uint32 , uint256 ) view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskChallengedInstances(arg0 uint32, arg1 *big.Int) (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskChallengedInstances(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0, arg1)
}

// TaskChallengedInstances is a free data retrieval call binding the contract method 0x7cfd8968.
//
// Solidity: function taskChallengedInstances(uint32 , uint256 ) view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskChallengedInstances(arg0 uint32, arg1 *big.Int) (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskChallengedInstances(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0, arg1)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// VerifyProof is a free data retrieval call binding the contract method 0x695ef6f9.
//
// Solidity: function verifyProof(bytes proof, uint256[6] instances) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) VerifyProof(opts *bind.CallOpts, proof []byte, instances [6]*big.Int) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "verifyProof", proof, instances)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x695ef6f9.
//
// Solidity: function verifyProof(bytes proof, uint256[6] instances) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) VerifyProof(proof []byte, instances [6]*big.Int) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.VerifyProof(&_ContractIncredibleSquaringTaskManager.CallOpts, proof, instances)
}

// VerifyProof is a free data retrieval call binding the contract method 0x695ef6f9.
//
// Solidity: function verifyProof(bytes proof, uint256[6] instances) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) VerifyProof(proof []byte, instances [6]*big.Int) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.VerifyProof(&_ContractIncredibleSquaringTaskManager.CallOpts, proof, instances)
}

// ProveChallenge is a paid mutator transaction binding the contract method 0x69b1843f.
//
// Solidity: function ProveChallenge(uint32 referenceTaskIndex, bytes proof) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) ProveChallenge(opts *bind.TransactOpts, referenceTaskIndex uint32, proof []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "ProveChallenge", referenceTaskIndex, proof)
}

// ProveChallenge is a paid mutator transaction binding the contract method 0x69b1843f.
//
// Solidity: function ProveChallenge(uint32 referenceTaskIndex, bytes proof) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) ProveChallenge(referenceTaskIndex uint32, proof []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ProveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, referenceTaskIndex, proof)
}

// ProveChallenge is a paid mutator transaction binding the contract method 0x69b1843f.
//
// Solidity: function ProveChallenge(uint32 referenceTaskIndex, bytes proof) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) ProveChallenge(referenceTaskIndex uint32, proof []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ProveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, referenceTaskIndex, proof)
}

// CheckChallenge is a paid mutator transaction binding the contract method 0xa087d99a.
//
// Solidity: function checkChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) CheckChallenge(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "checkChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// CheckChallenge is a paid mutator transaction binding the contract method 0xa087d99a.
//
// Solidity: function checkChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CheckChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// CheckChallenge is a paid mutator transaction binding the contract method 0xa087d99a.
//
// Solidity: function checkChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) CheckChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "createNewTask", inputs, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x27e00b3a.
//
// Solidity: function raiseAndResolveChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x27e00b3a.
//
// Solidity: function raiseAndResolveChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RaiseAndResolveChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x27e00b3a.
//
// Solidity: function raiseAndResolveChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RaiseAndResolveChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// ContractIncredibleSquaringTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitializedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerInitialized represents a Initialized event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerInitializedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerInitialized)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSquaringTaskManagerInitialized, error) {
	event := new(ContractIncredibleSquaringTaskManagerInitialized)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IIncredibleSquaringTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0x70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractIncredibleSquaringTaskManagerNewTaskCreated, error) {
	event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSquaringTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSquaringTaskManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPaused represents a Paused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePaused(log types.Log) (*ContractIncredibleSquaringTaskManagerPaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerPaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractIncredibleSquaringTaskManagerPauserRegistrySet, error) {
	event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedIterator is returned from FilterTaskChallenged and is used to iterate over the raw logs and unpacked data for TaskChallenged events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallenged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallenged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallenged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallenged represents a TaskChallenged event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallenged struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskChallenged is a free log retrieval operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallenged(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerTaskChallengedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallenged", logs: logs, sub: sub}, nil
}

// WatchTaskChallenged is a free log subscription operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallenged(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallenged, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallenged)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallenged is a log parse operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallenged(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallenged, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallenged)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompletedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskCompletedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskCompleted, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskRespondedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskResponded represents a TaskResponded event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskResponded struct {
	TaskResponse         IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskRespondedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskResponded, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerUnpaused represents a Unpaused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerUnpausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerUnpaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractIncredibleSquaringTaskManagerUnpaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerUnpaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
