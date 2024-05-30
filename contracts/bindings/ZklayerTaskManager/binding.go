// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractZklayerTaskManager

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

// IZklayerTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IZklayerTaskManagerTask struct {
	Inputs                    [5]*big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IZklayerTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IZklayerTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Output             *big.Int
}

// IZklayerTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IZklayerTaskManagerTaskResponseMetadata struct {
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

// ContractZklayerTaskManagerMetaData contains all meta data concerning the ContractZklayerTaskManager contract.
var ContractZklayerTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeInstances\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_inferenceDB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveResultAccurate\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"raiseChallenger\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIZklayerTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIZklayerTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIZklayerTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIZklayerTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006232380380620062328339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615f3b620002f7600039600081816102e20152818161063101526133e70152600081816105fa0152612ab60152600081816104b3015281816116790152612c980152600081816104ed01528181612e6e01526130300152600081816105140152818161139d015281816127a0015281816129190152612b530152615f3b6000f3fe608060405234801561001057600080fd5b50600436106102485760003560e01c80635c975abb1161013b5780638b00ce7c116100b8578063df5cf7231161007c578063df5cf723146105f5578063f2fde38b1461061c578063f5c9899d1461062f578063f63c5bab14610655578063fabc1cbc1461065d57600080fd5b80638b00ce7c146105935780638da5cb5b146105a3578063b98d0908146105b4578063bdeea6cc146105c1578063cefdc1d4146105d457600080fd5b80636efb4636116100ff5780636efb463614610536578063715018a61461055757806372d18e8d1461055f5780637afa1eed1461056d578063886f11951461058057600080fd5b80635c975abb146104a65780635df45946146104ae57806367d6be44146104d557806368304835146104e85780636d14a9871461050f57600080fd5b806331b36bd9116101c95780634f739f741161018d5780634f739f741461041857806358a7cd2614610438578063595c6a671461044b5780635ac86ab7146104535780635c1556621461048657600080fd5b806331b36bd9146103925780633563b0d1146103b2578063416c7e5e146103d2578063480bab6b146103e55780634d2b57fe146103f857600080fd5b8063171f1d5b11610210578063171f1d5b146102ae5780631ad43189146102dd578063245a7bfc146103195780632cb223d5146103445780632d89f6fc1461037257600080fd5b80630627721e1461024d57806310d67a2f14610262578063136439dd146102755780631459457a14610288578063162d8a3f1461029b575b600080fd5b61026061025b366004614783565b610670565b005b6102606102703660046147ff565b6107b9565b61026061028336600461481c565b610875565b610260610296366004614835565b6109b4565b6102606102a93660046148d1565b610b13565b6102c16102bc366004614a95565b610c5f565b6040805192151583529015156020830152015b60405180910390f35b6103047f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102d4565b60cc5461032c906001600160a01b031681565b6040516001600160a01b0390911681526020016102d4565b610364610352366004614ae6565b60cb6020526000908152604090205481565b6040519081526020016102d4565b610364610380366004614ae6565b60ca6020526000908152604090205481565b6103a56103a0366004614b26565b610de9565b6040516102d49190614c14565b6103c56103c0366004614c27565b610f05565b6040516102d49190614d79565b6102606103e0366004614d9a565b61139b565b6102606103f3366004614e26565b611510565b61040b610406366004614f07565b6117fc565b6040516102d49190614f56565b61042b610426366004614fe7565b611911565b6040516102d491906150b1565b61026061044636600461516c565b612037565b6102606120db565b6104766104613660046151fd565b606654600160ff9092169190911b9081161490565b60405190151581526020016102d4565b61049961049436600461521a565b6121a2565b6040516102d4919061527d565b606654610364565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b6103646104e33660046152b5565b61236a565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b6105496105443660046154ee565b6123ed565b6040516102d49291906155ae565b6102606132e5565b60c95463ffffffff16610304565b60cd5461032c906001600160a01b031681565b60655461032c906001600160a01b031681565b60c9546103049063ffffffff1681565b6033546001600160a01b031661032c565b6097546104769060ff1681565b6102606105cf3660046155f7565b6132f9565b6105e76105e236600461566b565b6135b1565b6040516102d49291906156a2565b61032c7f000000000000000000000000000000000000000000000000000000000000000081565b61026061062a3660046147ff565b613743565b7f0000000000000000000000000000000000000000000000000000000000000000610304565b610304606481565b61026061066b36600461481c565b6137b9565b60cd546001600160a01b0316331461068757600080fd5b61068f6145f8565b6040805160a081810190925290869060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f8501839004830281018301909152838152908490849081908401838280828437600092019190915250505050604080830191909152516107199082906020016156c3565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f54977349061077c9084906156c3565b60405180910390a260c9546107989063ffffffff16600161578f565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561080c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083091906157b7565b6001600160a01b0316336001600160a01b0316146108695760405162461bcd60e51b8152600401610860906157d4565b60405180910390fd5b61087281613915565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156108bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e1919061581e565b6108fd5760405162461bcd60e51b81526004016108609061583b565b606654818116146109765760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610860565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156109d45750600054600160ff909116105b806109ee5750303b1580156109ee575060005460ff166001145b610a515760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610860565b6000805460ff191660011790558015610a74576000805461ff0019166101001790555b610a7f866000613a0c565b610a8885613af6565b60cc80546001600160a01b038087166001600160a01b03199283161790925560cd805486841690831617905560ce8054928516929091169190911790558015610b0b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6000610b226020840184614ae6565b63ffffffff8116600090815260cb6020526040902054909150610b4457600080fd5b8282604051602001610b579291906158a1565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb90935291205414610b8d57600080fd5b6064610b9c6020840184614ae6565b610ba6919061578f565b63ffffffff164363ffffffff161115610bbe57600080fd5b60ce54604051631150ef3360e21b81526001600160a01b0390911690634543bccc90610bf690849088906020890135906004016158d7565b600060405180830381600087803b158015610c1057600080fd5b505af1158015610c24573d6000803e3d6000fd5b505060405163ffffffff841692507f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a9150600090a250505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610ca757610ca76158fc565b60200201518951600160200201518a60200151600060028110610ccc57610ccc6158fc565b60200201518b60200151600160028110610ce857610ce86158fc565b602090810291909101518c518d830151604051610d459a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d689190615912565b9050610ddb610d81610d7a8884613b48565b8690613bdf565b610d89613c73565b610dd1610dc285610dbc604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613b48565b610dcb8c613d33565b90613bdf565b886201d4c0613dc3565b909890975095505050505050565b606081516001600160401b03811115610e0457610e04614930565b604051908082528060200260200182016040528015610e2d578160200160208202803683370190505b50905060005b8251811015610efe57836001600160a01b03166313542a4e848381518110610e5d57610e5d6158fc565b60200260200101516040518263ffffffff1660e01b8152600401610e9091906001600160a01b0391909116815260200190565b602060405180830381865afa158015610ead573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed19190615934565b828281518110610ee357610ee36158fc565b6020908102919091010152610ef78161594d565b9050610e33565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6b91906157b7565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd191906157b7565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611013573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103791906157b7565b9050600086516001600160401b0381111561105457611054614930565b60405190808252806020026020018201604052801561108757816020015b60608152602001906001900390816110725790505b50905060005b875181101561138f5760008882815181106110aa576110aa6158fc565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa15801561110b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111339190810190615968565b905080516001600160401b0381111561114e5761114e614930565b60405190808252806020026020018201604052801561119957816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161116c5790505b508484815181106111ac576111ac6158fc565b602002602001018190525060005b8151811015611379576040518060600160405280876001600160a01b03166347b314e88585815181106111ef576111ef6158fc565b60200260200101516040518263ffffffff1660e01b815260040161121591815260200190565b602060405180830381865afa158015611232573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125691906157b7565b6001600160a01b03168152602001838381518110611276576112766158fc565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106112a4576112a46158fc565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611300573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132491906159f8565b6001600160601b0316815250858581518110611342576113426158fc565b6020026020010151828151811061135b5761135b6158fc565b602002602001018190525080806113719061594d565b9150506111ba565b50505080806113879061594d565b91505061108d565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141d91906157b7565b6001600160a01b0316336001600160a01b0316146114c95760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610860565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b600061151f6020850185614ae6565b9050600082516001600160401b0381111561153c5761153c614930565b604051908082528060200260200182016040528015611565578160200160208202803683370190505b50905060005b83518110156115d7576115a8848281518110611589576115896158fc565b6020026020010151805160009081526020918201519091526040902090565b8282815181106115ba576115ba6158fc565b6020908102919091010152806115cf8161594d565b91505061156b565b5060006115ea60c0880160a08901614ae6565b826040516020016115fc929190615a21565b6040516020818303038152906040528051906020012090508460200135811461162457600080fd5b600084516001600160401b0381111561163f5761163f614930565b604051908082528060200260200182016040528015611668578160200160208202803683370190505b50905060005b855181101561175b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106116b8576116b86158fc565b60200260200101516040518263ffffffff1660e01b81526004016116de91815260200190565b602060405180830381865afa1580156116fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171f91906157b7565b828281518110611731576117316158fc565b6001600160a01b0390921660209283029190910190910152806117538161594d565b91505061166e565b5060ce5460405163959b127960e01b815263ffffffff861660048201526001600160a01b039091169063959b127990602401600060405180830381600087803b1580156117a757600080fd5b505af11580156117bb573d6000803e3d6000fd5b505060405133925063ffffffff871691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050565b606081516001600160401b0381111561181757611817614930565b604051908082528060200260200182016040528015611840578160200160208202803683370190505b50905060005b8251811015610efe57836001600160a01b031663296bb064848381518110611870576118706158fc565b60200260200101516040518263ffffffff1660e01b815260040161189691815260200190565b602060405180830381865afa1580156118b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118d791906157b7565b8282815181106118e9576118e96158fc565b6001600160a01b039092166020928302919091019091015261190a8161594d565b9050611846565b61193c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561197c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119a091906157b7565b90506119cd6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906119fd908b9089908990600401615a9f565b600060405180830381865afa158015611a1a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a429190810190615ac8565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611a74908b908b908b90600401615b7f565b600060405180830381865afa158015611a91573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ab99190810190615ac8565b6040820152856001600160401b03811115611ad657611ad6614930565b604051908082528060200260200182016040528015611b0957816020015b6060815260200190600190039081611af45790505b50606082015260005b60ff8116871115611f48576000856001600160401b03811115611b3757611b37614930565b604051908082528060200260200182016040528015611b60578160200160208202803683370190505b5083606001518360ff1681518110611b7a57611b7a6158fc565b602002602001018190525060005b86811015611e485760008c6001600160a01b03166304ec63518a8a85818110611bb357611bb36158fc565b905060200201358e88600001518681518110611bd157611bd16158fc565b60200260200101516040518463ffffffff1660e01b8152600401611c0e9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611c2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4f9190615b9f565b90506001600160c01b038116611cf35760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610860565b8a8a8560ff16818110611d0857611d086158fc565b6001600160c01b03841692013560f81c9190911c600190811614159050611e3557856001600160a01b031663dd9846b98a8a85818110611d4a57611d4a6158fc565b905060200201358d8d8860ff16818110611d6657611d666158fc565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611dbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de09190615bc8565b85606001518560ff1681518110611df957611df96158fc565b60200260200101518481518110611e1257611e126158fc565b63ffffffff9092166020928302919091019091015282611e318161594d565b9350505b5080611e408161594d565b915050611b88565b506000816001600160401b03811115611e6357611e63614930565b604051908082528060200260200182016040528015611e8c578160200160208202803683370190505b50905060005b82811015611f0d5784606001518460ff1681518110611eb357611eb36158fc565b60200260200101518181518110611ecc57611ecc6158fc565b6020026020010151828281518110611ee657611ee66158fc565b63ffffffff9092166020928302919091019091015280611f058161594d565b915050611e92565b508084606001518460ff1681518110611f2857611f286158fc565b602002602001018190525050508080611f4090615be5565b915050611b12565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fad91906157b7565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611fe0908b908b908e90600401615c05565b600060405180830381865afa158015611ffd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120259190810190615ac8565b60208301525098975050505050505050565b60ce54604051632c53e69360e11b81526001600160a01b03909116906358a7cd269061206f9088908890889088908890600401615c2f565b600060405180830381600087803b15801561208957600080fd5b505af115801561209d573d6000803e3d6000fd5b505060405133925063ffffffff881691507ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015612123573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612147919061581e565b6121635760405162461bcd60e51b81526004016108609061583b565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016121d4929190615c6e565b600060405180830381865afa1580156121f1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122199190810190615ac8565b9050600084516001600160401b0381111561223657612236614930565b60405190808252806020026020018201604052801561225f578160200160208202803683370190505b50905060005b855181101561236057866001600160a01b03166304ec635187838151811061228f5761228f6158fc565b6020026020010151878685815181106122aa576122aa6158fc565b60200260200101516040518463ffffffff1660e01b81526004016122e79392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612304573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123289190615b9f565b6001600160c01b0316828281518110612343576123436158fc565b6020908102919091010152806123588161594d565b915050612265565b5095945050505050565b60ce546040516319f5af9160e21b815263ffffffff84166004820152602481018390526000916001600160a01b0316906367d6be4490604401602060405180830381865afa1580156123c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123e49190615934565b90505b92915050565b60408051808201909152606080825260208201526000846124645760405162461bcd60e51b81526020600482015260376024820152600080516020615ee683398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610860565b6040830151518514801561247c575060a08301515185145b801561248c575060c08301515185145b801561249c575060e08301515185145b6125065760405162461bcd60e51b81526020600482015260416024820152600080516020615ee683398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610860565b8251516020840151511461257e5760405162461bcd60e51b815260206004820152604460248201819052600080516020615ee6833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610860565b4363ffffffff168463ffffffff16106125ed5760405162461bcd60e51b815260206004820152603c6024820152600080516020615ee683398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610860565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561262e5761262e614930565b604051908082528060200260200182016040528015612657578160200160208202803683370190505b506020820152866001600160401b0381111561267557612675614930565b60405190808252806020026020018201604052801561269e578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156126d2576126d2614930565b6040519080825280602002602001820160405280156126fb578160200160208202803683370190505b5081526020860151516001600160401b0381111561271b5761271b614930565b604051908082528060200260200182016040528015612744578160200160208202803683370190505b50816020018190525060006128168a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156127ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128119190615c8d565b613fe7565b905060005b876020015151811015612a925761284188602001518281518110611589576115896158fc565b83602001518281518110612857576128576158fc565b60209081029190910101528015612917576020830151612878600183615caa565b81518110612888576128886158fc565b602002602001015160001c836020015182815181106128a9576128a96158fc565b602002602001015160001c11612917576040805162461bcd60e51b8152602060048201526024810191909152600080516020615ee683398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610860565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061295c5761295c6158fc565b60200260200101518b8b60000151858151811061297b5761297b6158fc565b60200260200101516040518463ffffffff1660e01b81526004016129b89392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156129d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129f99190615b9f565b6001600160c01b031683600001518281518110612a1857612a186158fc565b602002602001018181525050612a7e610d7a612a528486600001518581518110612a4457612a446158fc565b602002602001015116614071565b8a602001518481518110612a6857612a686158fc565b602002602001015161409c90919063ffffffff16565b945080612a8a8161594d565b91505061281b565b5050612a9d83614180565b60975490935060ff16600081612ab4576000612b36565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b369190615934565b905060005b8a8110156131b4578215612c96578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612b9257612b926158fc565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612bd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf69190615934565b612c009190615cc1565b11612c965760405162461bcd60e51b81526020600482015260666024820152600080516020615ee683398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610860565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612cd757612cd76158fc565b9050013560f81c60f81b60f81c8c8c60a001518581518110612cfb57612cfb6158fc565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612d57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d7b9190615cd9565b6001600160401b031916612d9e8a604001518381518110611589576115896158fc565b67ffffffffffffffff191614612e3a5760405162461bcd60e51b81526020600482015260616024820152600080516020615ee683398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610860565b612e6a89604001518281518110612e5357612e536158fc565b602002602001015187613bdf90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612ead57612ead6158fc565b9050013560f81c60f81b60f81c8c8c60c001518581518110612ed157612ed16158fc565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612f2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f5191906159f8565b85602001518281518110612f6757612f676158fc565b6001600160601b03909216602092830291909101820152850151805182908110612f9357612f936158fc565b602002602001015185600001518281518110612fb157612fb16158fc565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561319f5761302986600001518281518110612ffb57612ffb6158fc565b60200260200101518f8f86818110613015576130156158fc565b600192013560f81c9290921c811614919050565b1561318d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061306f5761306f6158fc565b9050013560f81c60f81b60f81c8e89602001518581518110613093576130936158fc565b60200260200101518f60e0015188815181106130b1576130b16158fc565b602002602001015187815181106130ca576130ca6158fc565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561312e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061315291906159f8565b8751805185908110613166576131666158fc565b6020026020010181815161317a9190615d04565b6001600160601b03169052506001909101905b806131978161594d565b915050612fd5565b505080806131ac9061594d565b915050612b3b565b5050506000806131ce8c868a606001518b60800151610c5f565b915091508161323f5760405162461bcd60e51b81526020600482015260436024820152600080516020615ee683398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610860565b806132a05760405162461bcd60e51b81526020600482015260396024820152600080516020615ee683398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610860565b505060008782602001516040516020016132bb929190615a21565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6132ed61421b565b6132f76000613af6565b565b60cc546001600160a01b0316331461331057600080fd5b600061332260c0850160a08601614ae6565b905036600061333460c0870187615d2c565b9092509050600061334c610100880160e08901614ae6565b905060ca600061335f6020890189614ae6565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161338b9190615d72565b60405160208183030381529060405280519060200120146133ab57600080fd5b600060cb816133bd60208a018a614ae6565b63ffffffff1663ffffffff16815260200190815260200160002054146133e257600080fd5b61340c7f00000000000000000000000000000000000000000000000000000000000000008561578f565b63ffffffff164363ffffffff16111561342457600080fd5b6000866040516020016134379190615e1b565b60405160208183030381529060405280519060200120905060008061345f8387878a8c6123ed565b9150915060005b858110156134f3578460ff1683602001518281518110613488576134886158fc565b602002602001015161349a9190615e29565b6001600160601b03166064846000015183815181106134bb576134bb6158fc565b60200260200101516001600160601b03166134d69190615e58565b10156134e157600080fd5b806134eb8161594d565b915050613466565b5060408051808201825263ffffffff43168152602080820184905291519091613520918c91849101615e77565b6040516020818303038152906040528051906020012060cb60008c600001602081019061354d9190614ae6565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a8260405161359c929190615e77565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106135ec576135ec6158fc565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906136289088908690600401615c6e565b600060405180830381865afa158015613645573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261366d9190810190615ac8565b60008151811061367f5761367f6158fc565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156136eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061370f9190615b9f565b6001600160c01b03169050600061372582614275565b9050816137338a838a610f05565b9550955050505050935093915050565b61374b61421b565b6001600160a01b0381166137b05760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610860565b61087281613af6565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561380c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061383091906157b7565b6001600160a01b0316336001600160a01b0316146138605760405162461bcd60e51b8152600401610860906157d4565b6066541981196066541916146138de5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610860565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016109a9565b6001600160a01b0381166139a35760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610860565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b0316158015613a2d57506001600160a01b03821615155b613aaf5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610860565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613af282613915565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613b64614626565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613b9757613b99565bfe5b5080613bd75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610860565b505092915050565b6040805180820190915260008082526020820152613bfb614644565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613b97575080613bd75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610860565b613c7b614662565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613d63600080516020615ec683398151915286615912565b90505b613d6f81614341565b9093509150600080516020615ec6833981519152828309831415613da9576040805180820190915290815260208101919091529392505050565b600080516020615ec6833981519152600182089050613d66565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613df5614687565b60005b6002811015613fba576000613e0e826006615e58565b9050848260028110613e2257613e226158fc565b60200201515183613e34836000615cc1565b600c8110613e4457613e446158fc565b6020020152848260028110613e5b57613e5b6158fc565b60200201516020015183826001613e729190615cc1565b600c8110613e8257613e826158fc565b6020020152838260028110613e9957613e996158fc565b6020020151515183613eac836002615cc1565b600c8110613ebc57613ebc6158fc565b6020020152838260028110613ed357613ed36158fc565b6020020151516001602002015183613eec836003615cc1565b600c8110613efc57613efc6158fc565b6020020152838260028110613f1357613f136158fc565b602002015160200151600060028110613f2e57613f2e6158fc565b602002015183613f3f836004615cc1565b600c8110613f4f57613f4f6158fc565b6020020152838260028110613f6657613f666158fc565b602002015160200151600160028110613f8157613f816158fc565b602002015183613f92836005615cc1565b600c8110613fa257613fa26158fc565b60200201525080613fb28161594d565b915050613df8565b50613fc36146a6565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613ff3846143c3565b9050808360ff166001901b116123e45760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610860565b6000805b82156123e757614086600184615caa565b909216918061409481615ea3565b915050614075565b60408051808201909152600080825260208201526102008261ffff16106140f85760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610860565b8161ffff166001141561410c5750816123e7565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061417557600161ffff871660ff83161c81161415614158576141558484613bdf565b93505b6141628384613bdf565b92506201fffe600192831b169101614128565b509195945050505050565b604080518082019091526000808252602082015281511580156141a557506020820151155b156141c3575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615ec683398151915284602001516141f69190615912565b61420e90600080516020615ec6833981519152615caa565b905292915050565b919050565b6033546001600160a01b031633146132f75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610860565b606060008061428384614071565b61ffff166001600160401b0381111561429e5761429e614930565b6040519080825280601f01601f1916602001820160405280156142c8576020820181803683370190505b5090506000805b8251821080156142e0575061010081105b15614337576001811b935085841615614327578060f81b838381518110614309576143096158fc565b60200101906001600160f81b031916908160001a9053508160010191505b6143308161594d565b90506142cf565b5090949350505050565b60008080600080516020615ec68339815191526003600080516020615ec683398151915286600080516020615ec68339815191528889090908905060006143b7827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615ec6833981519152614550565b91959194509092505050565b60006101008251111561444c5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610860565b815161445a57506000919050565b60008083600081518110614470576144706158fc565b0160200151600160f89190911c81901b92505b84518110156145475784818151811061449e5761449e6158fc565b0160200151600160f89190911c1b91508282116145335760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610860565b918117916145408161594d565b9050614483565b50909392505050565b60008061455b6146a6565b6145636146c4565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613b975750826145ed5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610860565b505195945050505050565b604051806080016040528061460b6146e2565b81526000602082018190526060604083018190529091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280614675614700565b8152602001614682614700565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461087257600080fd5b80356142168161471e565b60008083601f84011261474d57600080fd5b5081356001600160401b0381111561476457600080fd5b60208301915083602082850101111561477c57600080fd5b9250929050565b60008060008060e0858703121561479957600080fd5b60a08501868111156147aa57600080fd5b859450356147b78161471e565b925060c08501356001600160401b038111156147d257600080fd5b6147de8782880161473b565b95989497509550505050565b6001600160a01b038116811461087257600080fd5b60006020828403121561481157600080fd5b81356123e4816147ea565b60006020828403121561482e57600080fd5b5035919050565b600080600080600060a0868803121561484d57600080fd5b8535614858816147ea565b94506020860135614868816147ea565b93506040860135614878816147ea565b92506060860135614888816147ea565b91506080860135614898816147ea565b809150509295509295909350565b600061010082840312156148b957600080fd5b50919050565b6000604082840312156148b957600080fd5b600080600060a084860312156148e657600080fd5b83356001600160401b038111156148fc57600080fd5b614908868287016148a6565b93505061491885602086016148bf565b915061492785606086016148bf565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561496857614968614930565b60405290565b60405161010081016001600160401b038111828210171561496857614968614930565b604051601f8201601f191681016001600160401b03811182821017156149b9576149b9614930565b604052919050565b6000604082840312156149d357600080fd5b6149db614946565b9050813581526020820135602082015292915050565b600082601f830112614a0257600080fd5b604051604081018181106001600160401b0382111715614a2457614a24614930565b8060405250806040840185811115614a3b57600080fd5b845b81811015614175578035835260209283019201614a3d565b600060808284031215614a6757600080fd5b614a6f614946565b9050614a7b83836149f1565b8152614a8a83604084016149f1565b602082015292915050565b6000806000806101208587031215614aac57600080fd5b84359350614abd86602087016149c1565b9250614acc8660608701614a55565b9150614adb8660e087016149c1565b905092959194509250565b600060208284031215614af857600080fd5b81356123e48161471e565b60006001600160401b03821115614b1c57614b1c614930565b5060051b60200190565b60008060408385031215614b3957600080fd5b8235614b44816147ea565b91506020838101356001600160401b03811115614b6057600080fd5b8401601f81018613614b7157600080fd5b8035614b84614b7f82614b03565b614991565b81815260059190911b82018301908381019088831115614ba357600080fd5b928401925b82841015614bca578335614bbb816147ea565b82529284019290840190614ba8565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614c0957815187529582019590820190600101614bed565b509495945050505050565b6020815260006123e46020830184614bd9565b600080600060608486031215614c3c57600080fd5b8335614c47816147ea565b92506020848101356001600160401b0380821115614c6457600080fd5b818701915087601f830112614c7857600080fd5b813581811115614c8a57614c8a614930565b614c9c601f8201601f19168501614991565b91508082528884828501011115614cb257600080fd5b808484018584013760008482840101525080945050505061492760408501614730565b600081518084526020808501808196508360051b810191508286016000805b86811015614d6b578385038a52825180518087529087019087870190845b81811015614d5657835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614d12565b50509a87019a95505091850191600101614cf4565b509298975050505050505050565b6020815260006123e46020830184614cd5565b801515811461087257600080fd5b600060208284031215614dac57600080fd5b81356123e481614d8c565b600082601f830112614dc857600080fd5b81356020614dd8614b7f83614b03565b82815260069290921b84018101918181019086841115614df757600080fd5b8286015b84811015614e1b57614e0d88826149c1565b835291830191604001614dfb565b509695505050505050565b60008060008060c08587031215614e3c57600080fd5b84356001600160401b0380821115614e5357600080fd5b614e5f888389016148a6565b9550614e6e88602089016148bf565b9450614e7d88606089016148bf565b935060a0870135915080821115614e9357600080fd5b50614ea087828801614db7565b91505092959194509250565b600082601f830112614ebd57600080fd5b81356020614ecd614b7f83614b03565b82815260059290921b84018101918181019086841115614eec57600080fd5b8286015b84811015614e1b5780358352918301918301614ef0565b60008060408385031215614f1a57600080fd5b8235614f25816147ea565b915060208301356001600160401b03811115614f4057600080fd5b614f4c85828601614eac565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614f975783516001600160a01b031683529284019291840191600101614f72565b50909695505050505050565b60008083601f840112614fb557600080fd5b5081356001600160401b03811115614fcc57600080fd5b6020830191508360208260051b850101111561477c57600080fd5b6000806000806000806080878903121561500057600080fd5b863561500b816147ea565b9550602087013561501b8161471e565b945060408701356001600160401b038082111561503757600080fd5b6150438a838b0161473b565b9096509450606089013591508082111561505c57600080fd5b5061506989828a01614fa3565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614c0957815163ffffffff168752958201959082019060010161508f565b6000602080835283516080828501526150cd60a085018261507b565b905081850151601f19808684030160408701526150ea838361507b565b92506040870151915080868403016060870152615107838361507b565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561515e578487830301845261514c82875161507b565b95880195938801939150600101615132565b509998505050505050505050565b60008060008060006060868803121561518457600080fd5b853561518f8161471e565b945060208601356001600160401b03808211156151ab57600080fd5b6151b789838a01614fa3565b909650945060408801359150808211156151d057600080fd5b506151dd8882890161473b565b969995985093965092949392505050565b60ff8116811461087257600080fd5b60006020828403121561520f57600080fd5b81356123e4816151ee565b60008060006060848603121561522f57600080fd5b833561523a816147ea565b925060208401356001600160401b0381111561525557600080fd5b61526186828701614eac565b92505060408401356152728161471e565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614f9757835183529284019291840191600101615299565b600080604083850312156152c857600080fd5b82356152d38161471e565b946020939093013593505050565b600082601f8301126152f257600080fd5b81356020615302614b7f83614b03565b82815260059290921b8401810191818101908684111561532157600080fd5b8286015b84811015614e1b5780356153388161471e565b8352918301918301615325565b600082601f83011261535657600080fd5b81356020615366614b7f83614b03565b82815260059290921b8401810191818101908684111561538557600080fd5b8286015b84811015614e1b5780356001600160401b038111156153a85760008081fd5b6153b68986838b01016152e1565b845250918301918301615389565b600061018082840312156153d757600080fd5b6153df61496e565b905081356001600160401b03808211156153f857600080fd5b615404858386016152e1565b8352602084013591508082111561541a57600080fd5b61542685838601614db7565b6020840152604084013591508082111561543f57600080fd5b61544b85838601614db7565b604084015261545d8560608601614a55565b606084015261546f8560e086016149c1565b608084015261012084013591508082111561548957600080fd5b615495858386016152e1565b60a08401526101408401359150808211156154af57600080fd5b6154bb858386016152e1565b60c08401526101608401359150808211156154d557600080fd5b506154e284828501615345565b60e08301525092915050565b60008060008060006080868803121561550657600080fd5b8535945060208601356001600160401b038082111561552457600080fd5b61553089838a0161473b565b9096509450604088013591506155458261471e565b9092506060870135908082111561555b57600080fd5b50615568888289016153c4565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614c095781516001600160601b031687529582019590820190600101615589565b60408152600083516040808401526155c96080840182615575565b90506020850151603f198483030160608501526155e68282615575565b925050508260208301529392505050565b60008060006080848603121561560c57600080fd5b83356001600160401b038082111561562357600080fd5b61562f878388016148a6565b945061563e87602088016148bf565b9350606086013591508082111561565457600080fd5b50615661868287016153c4565b9150509250925092565b60008060006060848603121561568057600080fd5b833561568b816147ea565b92506020840135915060408401356152728161471e565b8281526040602082015260006156bb6040830184614cd5565b949350505050565b6020808252825160009190828483015b60058210156156f25782518152918301916001919091019083016156d3565b50505063ffffffff818501511660c084015260408401516101008060e086015281518061012087015260005b8181101561573b578381018501518782016101400152840161571e565b8181111561574e57600061014083890101525b50606087015163ffffffff8116878401529350601f01601f1916949094016101400195945050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168185168083038211156157ae576157ae615779565b01949350505050565b6000602082840312156157c957600080fd5b81516123e4816147ea565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561583057600080fd5b81516123e481614d8c565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b803561588e8161471e565b63ffffffff168252602090810135910152565b608081016158af8285615883565b82356158ba8161471e565b63ffffffff16604083015260209290920135606090910152919050565b63ffffffff8416815260e0810160a084602084013760c0919091019190915292915050565b634e487b7160e01b600052603260045260246000fd5b60008261592f57634e487b7160e01b600052601260045260246000fd5b500690565b60006020828403121561594657600080fd5b5051919050565b600060001982141561596157615961615779565b5060010190565b6000602080838503121561597b57600080fd5b82516001600160401b0381111561599157600080fd5b8301601f810185136159a257600080fd5b80516159b0614b7f82614b03565b81815260059190911b820183019083810190878311156159cf57600080fd5b928401925b828410156159ed578351825292840192908401906159d4565b979650505050505050565b600060208284031215615a0a57600080fd5b81516001600160601b03811681146123e457600080fd5b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615a5c57815185529382019390820190600101615a40565b5092979650505050505050565b81835260006001600160fb1b03831115615a8257600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff84168152604060208201526000615abf604083018486615a69565b95945050505050565b60006020808385031215615adb57600080fd5b82516001600160401b03811115615af157600080fd5b8301601f81018513615b0257600080fd5b8051615b10614b7f82614b03565b81815260059190911b82018301908381019087831115615b2f57600080fd5b928401925b828410156159ed578351615b478161471e565b82529284019290840190615b34565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615abf604083018486615b56565b600060208284031215615bb157600080fd5b81516001600160c01b03811681146123e457600080fd5b600060208284031215615bda57600080fd5b81516123e48161471e565b600060ff821660ff811415615bfc57615bfc615779565b60010192915050565b604081526000615c19604083018587615b56565b905063ffffffff83166020830152949350505050565b63ffffffff86168152606060208201526000615c4f606083018688615a69565b8281036040840152615c62818587615b56565b98975050505050505050565b63ffffffff831681526040602082015260006156bb6040830184614bd9565b600060208284031215615c9f57600080fd5b81516123e4816151ee565b600082821015615cbc57615cbc615779565b500390565b60008219821115615cd457615cd4615779565b500190565b600060208284031215615ceb57600080fd5b815167ffffffffffffffff19811681146123e457600080fd5b60006001600160601b0383811690831681811015615d2457615d24615779565b039392505050565b6000808335601e19843603018112615d4357600080fd5b8301803591506001600160401b03821115615d5d57600080fd5b60200191503681900382131561477c57600080fd5b6020815260a0826020830137600060c08201600080825260a0850135615d978161471e565b63ffffffff1690915260c08401359036859003601e19018212615db8578081fd5b9084019081356001600160401b03811115615dd1578182fd5b803603861315615ddf578182fd5b61010091508160e0860152615dfc61012086018260208601615b56565b925050615e0b60e08601614730565b63ffffffff811685830152614337565b604081016123e78284615883565b60006001600160601b0380831681851681830481118215151615615e4f57615e4f615779565b02949350505050565b6000816000190483118215151615615e7257615e72615779565b500290565b60808101615e858285615883565b63ffffffff8351166040830152602083015160608301529392505050565b600061ffff80831681811415615ebb57615ebb615779565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220ce90fcc9ef344cb7d2c095c5ea90c6742c9054cac4c8f6583b3a31317a4e9be464736f6c634300080c0033",
}

// ContractZklayerTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractZklayerTaskManagerMetaData.ABI instead.
var ContractZklayerTaskManagerABI = ContractZklayerTaskManagerMetaData.ABI

// ContractZklayerTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractZklayerTaskManagerMetaData.Bin instead.
var ContractZklayerTaskManagerBin = ContractZklayerTaskManagerMetaData.Bin

// DeployContractZklayerTaskManager deploys a new Ethereum contract, binding an instance of ContractZklayerTaskManager to it.
func DeployContractZklayerTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractZklayerTaskManager, error) {
	parsed, err := ContractZklayerTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractZklayerTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractZklayerTaskManager{ContractZklayerTaskManagerCaller: ContractZklayerTaskManagerCaller{contract: contract}, ContractZklayerTaskManagerTransactor: ContractZklayerTaskManagerTransactor{contract: contract}, ContractZklayerTaskManagerFilterer: ContractZklayerTaskManagerFilterer{contract: contract}}, nil
}

// ContractZklayerTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractZklayerTaskManager struct {
	ContractZklayerTaskManagerCaller     // Read-only binding to the contract
	ContractZklayerTaskManagerTransactor // Write-only binding to the contract
	ContractZklayerTaskManagerFilterer   // Log filterer for contract events
}

// ContractZklayerTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractZklayerTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZklayerTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractZklayerTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZklayerTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractZklayerTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZklayerTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractZklayerTaskManagerSession struct {
	Contract     *ContractZklayerTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractZklayerTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractZklayerTaskManagerCallerSession struct {
	Contract *ContractZklayerTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractZklayerTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractZklayerTaskManagerTransactorSession struct {
	Contract     *ContractZklayerTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractZklayerTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractZklayerTaskManagerRaw struct {
	Contract *ContractZklayerTaskManager // Generic contract binding to access the raw methods on
}

// ContractZklayerTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractZklayerTaskManagerCallerRaw struct {
	Contract *ContractZklayerTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractZklayerTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractZklayerTaskManagerTransactorRaw struct {
	Contract *ContractZklayerTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractZklayerTaskManager creates a new instance of ContractZklayerTaskManager, bound to a specific deployed contract.
func NewContractZklayerTaskManager(address common.Address, backend bind.ContractBackend) (*ContractZklayerTaskManager, error) {
	contract, err := bindContractZklayerTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManager{ContractZklayerTaskManagerCaller: ContractZklayerTaskManagerCaller{contract: contract}, ContractZklayerTaskManagerTransactor: ContractZklayerTaskManagerTransactor{contract: contract}, ContractZklayerTaskManagerFilterer: ContractZklayerTaskManagerFilterer{contract: contract}}, nil
}

// NewContractZklayerTaskManagerCaller creates a new read-only instance of ContractZklayerTaskManager, bound to a specific deployed contract.
func NewContractZklayerTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractZklayerTaskManagerCaller, error) {
	contract, err := bindContractZklayerTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerCaller{contract: contract}, nil
}

// NewContractZklayerTaskManagerTransactor creates a new write-only instance of ContractZklayerTaskManager, bound to a specific deployed contract.
func NewContractZklayerTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractZklayerTaskManagerTransactor, error) {
	contract, err := bindContractZklayerTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerTransactor{contract: contract}, nil
}

// NewContractZklayerTaskManagerFilterer creates a new log filterer instance of ContractZklayerTaskManager, bound to a specific deployed contract.
func NewContractZklayerTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractZklayerTaskManagerFilterer, error) {
	contract, err := bindContractZklayerTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerFilterer{contract: contract}, nil
}

// bindContractZklayerTaskManager binds a generic wrapper to an already deployed contract.
func bindContractZklayerTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractZklayerTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZklayerTaskManager.Contract.ContractZklayerTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ContractZklayerTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ContractZklayerTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZklayerTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZklayerTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractZklayerTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractZklayerTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractZklayerTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Aggregator(&_ContractZklayerTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Aggregator(&_ContractZklayerTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractZklayerTaskManager.Contract.AllTaskHashes(&_ContractZklayerTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractZklayerTaskManager.Contract.AllTaskHashes(&_ContractZklayerTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractZklayerTaskManager.Contract.AllTaskResponses(&_ContractZklayerTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractZklayerTaskManager.Contract.AllTaskResponses(&_ContractZklayerTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.BlsApkRegistry(&_ContractZklayerTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.BlsApkRegistry(&_ContractZklayerTaskManager.CallOpts)
}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 id, uint256 index) view returns(uint256)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) ChallengeInstances(opts *bind.CallOpts, id uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "challengeInstances", id, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 id, uint256 index) view returns(uint256)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) ChallengeInstances(id uint32, index *big.Int) (*big.Int, error) {
	return _ContractZklayerTaskManager.Contract.ChallengeInstances(&_ContractZklayerTaskManager.CallOpts, id, index)
}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 id, uint256 index) view returns(uint256)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) ChallengeInstances(id uint32, index *big.Int) (*big.Int, error) {
	return _ContractZklayerTaskManager.Contract.ChallengeInstances(&_ContractZklayerTaskManager.CallOpts, id, index)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZklayerTaskManager.Contract.CheckSignatures(&_ContractZklayerTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractZklayerTaskManager.Contract.CheckSignatures(&_ContractZklayerTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Delegation(&_ContractZklayerTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Delegation(&_ContractZklayerTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Generator() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Generator(&_ContractZklayerTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Generator(&_ContractZklayerTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractZklayerTaskManager.Contract.GetBatchOperatorFromId(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractZklayerTaskManager.Contract.GetBatchOperatorFromId(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractZklayerTaskManager.Contract.GetBatchOperatorId(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractZklayerTaskManager.Contract.GetBatchOperatorId(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZklayerTaskManager.Contract.GetCheckSignaturesIndices(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractZklayerTaskManager.Contract.GetCheckSignaturesIndices(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZklayerTaskManager.Contract.GetOperatorState(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractZklayerTaskManager.Contract.GetOperatorState(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZklayerTaskManager.Contract.GetOperatorState0(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractZklayerTaskManager.Contract.GetOperatorState0(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZklayerTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractZklayerTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractZklayerTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractZklayerTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractZklayerTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.LatestTaskNum(&_ContractZklayerTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.LatestTaskNum(&_ContractZklayerTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Owner() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Owner(&_ContractZklayerTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.Owner(&_ContractZklayerTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractZklayerTaskManager.Contract.Paused(&_ContractZklayerTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractZklayerTaskManager.Contract.Paused(&_ContractZklayerTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractZklayerTaskManager.Contract.Paused0(&_ContractZklayerTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractZklayerTaskManager.Contract.Paused0(&_ContractZklayerTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.PauserRegistry(&_ContractZklayerTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.PauserRegistry(&_ContractZklayerTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.RegistryCoordinator(&_ContractZklayerTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.RegistryCoordinator(&_ContractZklayerTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.StakeRegistry(&_ContractZklayerTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractZklayerTaskManager.Contract.StakeRegistry(&_ContractZklayerTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZklayerTaskManager.Contract.StaleStakesForbidden(&_ContractZklayerTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractZklayerTaskManager.Contract.StaleStakesForbidden(&_ContractZklayerTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.TaskNumber(&_ContractZklayerTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractZklayerTaskManager.Contract.TaskNumber(&_ContractZklayerTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractZklayerTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZklayerTaskManager.Contract.TrySignatureAndApkVerification(&_ContractZklayerTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractZklayerTaskManager.Contract.TrySignatureAndApkVerification(&_ContractZklayerTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x480bab6b.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) ConfirmChallenge(opts *bind.TransactOpts, task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, taskResponseMetadata IZklayerTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "confirmChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x480bab6b.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) ConfirmChallenge(task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, taskResponseMetadata IZklayerTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ConfirmChallenge(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x480bab6b.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) ConfirmChallenge(task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, taskResponseMetadata IZklayerTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ConfirmChallenge(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "createNewTask", inputs, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.CreateNewTask(&_ContractZklayerTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.CreateNewTask(&_ContractZklayerTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _inferenceDB) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _inferenceDB common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator, _inferenceDB)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _inferenceDB) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _inferenceDB common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.Initialize(&_ContractZklayerTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _inferenceDB)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _inferenceDB) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _inferenceDB common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.Initialize(&_ContractZklayerTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _inferenceDB)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.Pause(&_ContractZklayerTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.Pause(&_ContractZklayerTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.PauseAll(&_ContractZklayerTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.PauseAll(&_ContractZklayerTaskManager.TransactOpts)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) ProveResultAccurate(opts *bind.TransactOpts, taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "proveResultAccurate", taskId, instances, proof)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) ProveResultAccurate(taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ProveResultAccurate(&_ContractZklayerTaskManager.TransactOpts, taskId, instances, proof)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) ProveResultAccurate(taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ProveResultAccurate(&_ContractZklayerTaskManager.TransactOpts, taskId, instances, proof)
}

// RaiseChallenger is a paid mutator transaction binding the contract method 0x162d8a3f.
//
// Solidity: function raiseChallenger((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) RaiseChallenger(opts *bind.TransactOpts, task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, taskResponseMetadata IZklayerTaskManagerTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "raiseChallenger", task, taskResponse, taskResponseMetadata)
}

// RaiseChallenger is a paid mutator transaction binding the contract method 0x162d8a3f.
//
// Solidity: function raiseChallenger((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) RaiseChallenger(task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, taskResponseMetadata IZklayerTaskManagerTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RaiseChallenger(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
}

// RaiseChallenger is a paid mutator transaction binding the contract method 0x162d8a3f.
//
// Solidity: function raiseChallenger((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) RaiseChallenger(task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, taskResponseMetadata IZklayerTaskManagerTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RaiseChallenger(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RenounceOwnership(&_ContractZklayerTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RenounceOwnership(&_ContractZklayerTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) RespondToTask(task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RespondToTask(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) RespondToTask(task IZklayerTaskManagerTask, taskResponse IZklayerTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RespondToTask(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.SetPauserRegistry(&_ContractZklayerTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.SetPauserRegistry(&_ContractZklayerTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.SetStaleStakesForbidden(&_ContractZklayerTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.SetStaleStakesForbidden(&_ContractZklayerTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.TransferOwnership(&_ContractZklayerTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.TransferOwnership(&_ContractZklayerTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.Unpause(&_ContractZklayerTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.Unpause(&_ContractZklayerTaskManager.TransactOpts, newPausedStatus)
}

// ContractZklayerTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerInitializedIterator struct {
	Event *ContractZklayerTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerInitialized)
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
		it.Event = new(ContractZklayerTaskManagerInitialized)
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
func (it *ContractZklayerTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerInitialized represents a Initialized event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractZklayerTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerInitializedIterator{contract: _ContractZklayerTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerInitialized)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractZklayerTaskManagerInitialized, error) {
	event := new(ContractZklayerTaskManagerInitialized)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerNewTaskCreatedIterator struct {
	Event *ContractZklayerTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerNewTaskCreated)
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
		it.Event = new(ContractZklayerTaskManagerNewTaskCreated)
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
func (it *ContractZklayerTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IZklayerTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32) task)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractZklayerTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerNewTaskCreatedIterator{contract: _ContractZklayerTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32) task)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerNewTaskCreated)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractZklayerTaskManagerNewTaskCreated, error) {
	event := new(ContractZklayerTaskManagerNewTaskCreated)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerOwnershipTransferredIterator struct {
	Event *ContractZklayerTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractZklayerTaskManagerOwnershipTransferred)
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
func (it *ContractZklayerTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractZklayerTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerOwnershipTransferredIterator{contract: _ContractZklayerTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerOwnershipTransferred)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractZklayerTaskManagerOwnershipTransferred, error) {
	event := new(ContractZklayerTaskManagerOwnershipTransferred)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerPausedIterator struct {
	Event *ContractZklayerTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerPaused)
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
		it.Event = new(ContractZklayerTaskManagerPaused)
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
func (it *ContractZklayerTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerPaused represents a Paused event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractZklayerTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerPausedIterator{contract: _ContractZklayerTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerPaused)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParsePaused(log types.Log) (*ContractZklayerTaskManagerPaused, error) {
	event := new(ContractZklayerTaskManagerPaused)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerPauserRegistrySetIterator struct {
	Event *ContractZklayerTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractZklayerTaskManagerPauserRegistrySet)
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
func (it *ContractZklayerTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractZklayerTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerPauserRegistrySetIterator{contract: _ContractZklayerTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerPauserRegistrySet)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractZklayerTaskManagerPauserRegistrySet, error) {
	event := new(ContractZklayerTaskManagerPauserRegistrySet)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractZklayerTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractZklayerTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractZklayerTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractZklayerTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractZklayerTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractZklayerTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractZklayerTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerTaskChallengedIterator is returned from FilterTaskChallenged and is used to iterate over the raw logs and unpacked data for TaskChallenged events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskChallengedIterator struct {
	Event *ContractZklayerTaskManagerTaskChallenged // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerTaskChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerTaskChallenged)
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
		it.Event = new(ContractZklayerTaskManagerTaskChallenged)
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
func (it *ContractZklayerTaskManagerTaskChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerTaskChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerTaskChallenged represents a TaskChallenged event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskChallenged struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskChallenged is a free log retrieval operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterTaskChallenged(opts *bind.FilterOpts, taskIndex []uint32) (*ContractZklayerTaskManagerTaskChallengedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerTaskChallengedIterator{contract: _ContractZklayerTaskManager.contract, event: "TaskChallenged", logs: logs, sub: sub}, nil
}

// WatchTaskChallenged is a free log subscription operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchTaskChallenged(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerTaskChallenged, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerTaskChallenged)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseTaskChallenged(log types.Log) (*ContractZklayerTaskManagerTaskChallenged, error) {
	event := new(ContractZklayerTaskManagerTaskChallenged)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractZklayerTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractZklayerTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractZklayerTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractZklayerTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractZklayerTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerTaskChallengedSuccessfully)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractZklayerTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractZklayerTaskManagerTaskChallengedSuccessfully)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractZklayerTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractZklayerTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractZklayerTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex uint32
	Prover    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed prover)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, prover []common.Address) (*ContractZklayerTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractZklayerTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed prover)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, prover []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed prover)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractZklayerTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractZklayerTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskCompletedIterator struct {
	Event *ContractZklayerTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerTaskCompleted)
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
		it.Event = new(ContractZklayerTaskManagerTaskCompleted)
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
func (it *ContractZklayerTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractZklayerTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerTaskCompletedIterator{contract: _ContractZklayerTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerTaskCompleted)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractZklayerTaskManagerTaskCompleted, error) {
	event := new(ContractZklayerTaskManagerTaskCompleted)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskRespondedIterator struct {
	Event *ContractZklayerTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerTaskResponded)
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
		it.Event = new(ContractZklayerTaskManagerTaskResponded)
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
func (it *ContractZklayerTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerTaskResponded represents a TaskResponded event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskResponded struct {
	TaskResponse         IZklayerTaskManagerTaskResponse
	TaskResponseMetadata IZklayerTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractZklayerTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerTaskRespondedIterator{contract: _ContractZklayerTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerTaskResponded)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractZklayerTaskManagerTaskResponded, error) {
	event := new(ContractZklayerTaskManagerTaskResponded)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerUnpausedIterator struct {
	Event *ContractZklayerTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerUnpaused)
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
		it.Event = new(ContractZklayerTaskManagerUnpaused)
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
func (it *ContractZklayerTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerUnpaused represents a Unpaused event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractZklayerTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerUnpausedIterator{contract: _ContractZklayerTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerUnpaused)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractZklayerTaskManagerUnpaused, error) {
	event := new(ContractZklayerTaskManagerUnpaused)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
