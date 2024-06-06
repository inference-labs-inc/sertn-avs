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

// ITaskStructTask is an auto generated low-level Go binding around an user-defined struct.
type ITaskStructTask struct {
	Inputs                    [5]*big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
	ProvenOnResponse          bool
}

// ITaskStructTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type ITaskStructTaskResponse struct {
	ReferenceTaskIndex uint32
	Output             *big.Int
}

// ITaskStructTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type ITaskStructTaskResponseMetadata struct {
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeInstances\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"provenOnResponce\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_inferenceDB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveResultAccurate\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"raiseChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTaskWithProof\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskRespondedWithProof\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162006205380380620062058339810160408190526200003591620001ed565b80806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001ed565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001ed565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b39190620001ed565b6001600160a01b031660e05250506097805460ff1916600117905562000214565b6001600160a01b0381168114620001ea57600080fd5b50565b6000602082840312156200020057600080fd5b81516200020d81620001d4565b9392505050565b60805160a05160c05160e051615f786200028d600039600081816105870152612a8b01526000818161045e015281816108be0152612c6d01526000818161049801528181612e4301526130050152600081816104bf0152818161165e01528181612775015281816128ee0152612b280152615f786000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80635ac86ab71161012557806379e3fb33116100ad578063b98d09081161007c578063b98d090814610554578063cefdc1d414610561578063df5cf72314610582578063f2fde38b146105a9578063fabc1cbc146105bc57600080fd5b806379e3fb331461050a5780637afa1eed1461051d578063886f1195146105305780638da5cb5b1461054357600080fd5b806367d6be44116100f457806367d6be441461048057806368304835146104935780636d14a987146104ba5780636efb4636146104e1578063715018a61461050257600080fd5b80635ac86ab7146103fe5780635c155662146104315780635c975abb146104515780635df459461461045957600080fd5b80632d89f6fc116101a8578063416c7e5e11610177578063416c7e5e146103905780634d2b57fe146103a35780634f739f74146103c357806358a7cd26146103e3578063595c6a67146103f657600080fd5b80632d89f6fc1461031d57806331b36bd91461033d5780633563b0d11461035d578063369bcec51461037d57600080fd5b80631459457a116101ef5780631459457a1461026f578063171f1d5b14610282578063245a7bfc146102b15780632cb223d5146102dc5780632ced7fa61461030a57600080fd5b8063027e4a371461022157806309b880971461023657806310d67a2f14610249578063136439dd1461025c575b600080fd5b61023461022f3660046146a9565b6105cf565b005b6102346102443660046148ac565b610755565b610234610257366004614947565b610a41565b61023461026a366004614964565b610afd565b61023461027d36600461497d565b610c3c565b610295610290366004614a84565b610d9b565b6040805192151583529015156020830152015b60405180910390f35b60cb546102c4906001600160a01b031681565b6040516001600160a01b0390911681526020016102a8565b6102fc6102ea366004614ad5565b60ca6020526000908152604090205481565b6040519081526020016102a8565b610234610318366004614b36565b610f25565b6102fc61032b366004614ad5565b60c96020526000908152604090205481565b61035061034b366004614bde565b610ffd565b6040516102a89190614cc7565b61037061036b366004614cda565b611119565b6040516102a89190614e35565b61023461038b366004614e48565b6115b1565b61023461039e366004614e9e565b61165c565b6103b66103b1366004614f16565b6117d1565b6040516102a89190614f65565b6103d66103d1366004614fb2565b6118e6565b6040516102a8919061506a565b6102346103f1366004615125565b61200c565b6102346120b0565b61042161040c3660046151b6565b606654600160ff9092169190911b9081161490565b60405190151581526020016102a8565b61044461043f3660046151d3565b612177565b6040516102a89190615236565b6066546102fc565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b6102fc61048e36600461526e565b61233f565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b6104f46104ef3660046154a7565b6123c2565b6040516102a8929190615567565b6102346132ba565b6102346105183660046155b0565b6132ce565b60cc546102c4906001600160a01b031681565b6065546102c4906001600160a01b031681565b6033546001600160a01b03166102c4565b6097546104219060ff1681565b61057461056f366004615624565b6134b7565b6040516102a892919061565b565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b6102346105b7366004614947565b613649565b6102346105ca366004614964565b6136bf565b60cc546001600160a01b031633146105e657600080fd5b6105ee6144fe565b6040805160a081810190925290879060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908616606083015260408051601f86018390048302810183019091528481529085908590819084018382808284376000920182905250604080870195909552861515608087015260cd54945190946001600160a01b03169350633b7d2cee9250610695915085906020016156c9565b604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004016106c991815260200190565b6020604051808303816000875af11580156106e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070c9190615752565b90508063ffffffff167f02fbf40079c9022bc174c39bce245cecd33b518a66e442ae48699be1ab5671ee8360405161074491906156c9565b60405180910390a250505050505050565b60006107646020850185614ad5565b9050600082516001600160401b0381111561078157610781614754565b6040519080825280602002602001820160405280156107aa578160200160208202803683370190505b50905060005b835181101561081c576107ed8482815181106107ce576107ce61576f565b6020026020010151805160009081526020918201519091526040902090565b8282815181106107ff576107ff61576f565b6020908102919091010152806108148161579b565b9150506107b0565b50600061082f60c0880160a08901614ad5565b826040516020016108419291906157b6565b6040516020818303038152906040528051906020012090508460200135811461086957600080fd5b600084516001600160401b0381111561088457610884614754565b6040519080825280602002602001820160405280156108ad578160200160208202803683370190505b50905060005b85518110156109a0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106108fd576108fd61576f565b60200260200101516040518263ffffffff1660e01b815260040161092391815260200190565b602060405180830381865afa158015610940573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096491906157fe565b8282815181106109765761097661576f565b6001600160a01b0390921660209283029190910190910152806109988161579b565b9150506108b3565b5060cd5460405163959b127960e01b815263ffffffff861660048201526001600160a01b039091169063959b127990602401600060405180830381600087803b1580156109ec57600080fd5b505af1158015610a00573d6000803e3d6000fd5b505060405133925063ffffffff871691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab891906157fe565b6001600160a01b0316336001600160a01b031614610af15760405162461bcd60e51b8152600401610ae89061581b565b60405180910390fd5b610afa8161381b565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610b45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b699190615865565b610b855760405162461bcd60e51b8152600401610ae890615882565b60665481811614610bfe5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610ae8565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff1615808015610c5c5750600054600160ff909116105b80610c765750303b158015610c76575060005460ff166001145b610cd95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610ae8565b6000805460ff191660011790558015610cfc576000805461ff0019166101001790555b610d07866000613912565b610d10856139fc565b60cb80546001600160a01b038087166001600160a01b03199283161790925560cc805486841690831617905560cd8054928516929091169190911790558015610d93576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610de357610de361576f565b60200201518951600160200201518a60200151600060028110610e0857610e0861576f565b60200201518b60200151600160028110610e2457610e2461576f565b602090810291909101518c518d830151604051610e819a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610ea491906158ca565b9050610f17610ebd610eb68884613a4e565b8690613ae5565b610ec5613b79565b610f0d610efe85610ef8604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613a4e565b610f078c613c39565b90613ae5565b886201d4c0613cc9565b909890975095505050505050565b60cd54604051631676bfd360e11b81526001600160a01b0390911690632ced7fa690610f5f90899089908990899089908990600401615a00565b600060405180830381600087803b158015610f7957600080fd5b505af1158015610f8d573d6000803e3d6000fd5b50339250505063ffffffff86167f4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada8686610fc8600182615a54565b818110610fd757610fd761576f565b90506020020135604051610fed91815260200190565b60405180910390a3505050505050565b606081516001600160401b0381111561101857611018614754565b604051908082528060200260200182016040528015611041578160200160208202803683370190505b50905060005b825181101561111257836001600160a01b03166313542a4e8483815181106110715761107161576f565b60200260200101516040518263ffffffff1660e01b81526004016110a491906001600160a01b0391909116815260200190565b602060405180830381865afa1580156110c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e59190615a6b565b8282815181106110f7576110f761576f565b602090810291909101015261110b8161579b565b9050611047565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561115b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117f91906157fe565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e591906157fe565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611227573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124b91906157fe565b9050600086516001600160401b0381111561126857611268614754565b60405190808252806020026020018201604052801561129b57816020015b60608152602001906001900390816112865790505b50905060005b87518110156115a35760008882815181106112be576112be61576f565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa15801561131f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526113479190810190615a84565b905080516001600160401b0381111561136257611362614754565b6040519080825280602002602001820160405280156113ad57816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816113805790505b508484815181106113c0576113c061576f565b602002602001018190525060005b815181101561158d576040518060600160405280876001600160a01b03166347b314e88585815181106114035761140361576f565b60200260200101516040518263ffffffff1660e01b815260040161142991815260200190565b602060405180830381865afa158015611446573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061146a91906157fe565b6001600160a01b0316815260200183838151811061148a5761148a61576f565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106114b8576114b861576f565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015611514573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115389190615b14565b6001600160601b03168152508585815181106115565761155661576f565b6020026020010151828151811061156f5761156f61576f565b602002602001018190525080806115859061579b565b9150506113ce565b505050808061159b9061579b565b9150506112a1565b5093505050505b9392505050565b60cd54604051630b5238df60e41b81526001600160a01b039091169063b5238df0906115e590869086908690600401615b5b565b600060405180830381600087803b1580156115ff57600080fd5b505af1158015611613573d6000803e3d6000fd5b50611625925050506020830183614ad5565b63ffffffff167f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a60405160405180910390a2505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116de91906157fe565b6001600160a01b0316336001600160a01b03161461178a5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610ae8565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b606081516001600160401b038111156117ec576117ec614754565b604051908082528060200260200182016040528015611815578160200160208202803683370190505b50905060005b825181101561111257836001600160a01b031663296bb0648483815181106118455761184561576f565b60200260200101516040518263ffffffff1660e01b815260040161186b91815260200190565b602060405180830381865afa158015611888573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ac91906157fe565b8282815181106118be576118be61576f565b6001600160a01b03909216602092830291909101909101526118df8161579b565b905061181b565b6119116040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611951573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061197591906157fe565b90506119a26040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906119d2908b9089908990600401615ba6565b600060405180830381865afa1580156119ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a179190810190615bcf565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611a49908b908b908b90600401615c5d565b600060405180830381865afa158015611a66573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a8e9190810190615bcf565b6040820152856001600160401b03811115611aab57611aab614754565b604051908082528060200260200182016040528015611ade57816020015b6060815260200190600190039081611ac95790505b50606082015260005b60ff8116871115611f1d576000856001600160401b03811115611b0c57611b0c614754565b604051908082528060200260200182016040528015611b35578160200160208202803683370190505b5083606001518360ff1681518110611b4f57611b4f61576f565b602002602001018190525060005b86811015611e1d5760008c6001600160a01b03166304ec63518a8a85818110611b8857611b8861576f565b905060200201358e88600001518681518110611ba657611ba661576f565b60200260200101516040518463ffffffff1660e01b8152600401611be39392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611c00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c249190615c7d565b90506001600160c01b038116611cc85760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610ae8565b8a8a8560ff16818110611cdd57611cdd61576f565b6001600160c01b03841692013560f81c9190911c600190811614159050611e0a57856001600160a01b031663dd9846b98a8a85818110611d1f57611d1f61576f565b905060200201358d8d8860ff16818110611d3b57611d3b61576f565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611d91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611db59190615752565b85606001518560ff1681518110611dce57611dce61576f565b60200260200101518481518110611de757611de761576f565b63ffffffff9092166020928302919091019091015282611e068161579b565b9350505b5080611e158161579b565b915050611b5d565b506000816001600160401b03811115611e3857611e38614754565b604051908082528060200260200182016040528015611e61578160200160208202803683370190505b50905060005b82811015611ee25784606001518460ff1681518110611e8857611e8861576f565b60200260200101518181518110611ea157611ea161576f565b6020026020010151828281518110611ebb57611ebb61576f565b63ffffffff9092166020928302919091019091015280611eda8161579b565b915050611e67565b508084606001518460ff1681518110611efd57611efd61576f565b602002602001018190525050508080611f1590615ca6565b915050611ae7565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f8291906157fe565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611fb5908b908b908e90600401615cc6565b600060405180830381865afa158015611fd2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ffa9190810190615bcf565b60208301525098975050505050505050565b60cd54604051632c53e69360e11b81526001600160a01b03909116906358a7cd26906120449088908890889088908890600401615cf0565b600060405180830381600087803b15801561205e57600080fd5b505af1158015612072573d6000803e3d6000fd5b505060405133925063ffffffff881691507ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156120f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061211c9190615865565b6121385760405162461bcd60e51b8152600401610ae890615882565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016121a9929190615d2f565b600060405180830381865afa1580156121c6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526121ee9190810190615bcf565b9050600084516001600160401b0381111561220b5761220b614754565b604051908082528060200260200182016040528015612234578160200160208202803683370190505b50905060005b855181101561233557866001600160a01b03166304ec63518783815181106122645761226461576f565b60200260200101518786858151811061227f5761227f61576f565b60200260200101516040518463ffffffff1660e01b81526004016122bc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156122d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122fd9190615c7d565b6001600160c01b03168282815181106123185761231861576f565b60209081029190910101528061232d8161579b565b91505061223a565b5095945050505050565b60cd546040516319f5af9160e21b815263ffffffff84166004820152602481018390526000916001600160a01b0316906367d6be4490604401602060405180830381865afa158015612395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123b99190615a6b565b90505b92915050565b60408051808201909152606080825260208201526000846124395760405162461bcd60e51b81526020600482015260376024820152600080516020615f2383398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610ae8565b60408301515185148015612451575060a08301515185145b8015612461575060c08301515185145b8015612471575060e08301515185145b6124db5760405162461bcd60e51b81526020600482015260416024820152600080516020615f2383398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610ae8565b825151602084015151146125535760405162461bcd60e51b815260206004820152604460248201819052600080516020615f23833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610ae8565b4363ffffffff168463ffffffff16106125c25760405162461bcd60e51b815260206004820152603c6024820152600080516020615f2383398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610ae8565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561260357612603614754565b60405190808252806020026020018201604052801561262c578160200160208202803683370190505b506020820152866001600160401b0381111561264a5761264a614754565b604051908082528060200260200182016040528015612673578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156126a7576126a7614754565b6040519080825280602002602001820160405280156126d0578160200160208202803683370190505b5081526020860151516001600160401b038111156126f0576126f0614754565b604051908082528060200260200182016040528015612719578160200160208202803683370190505b50816020018190525060006127eb8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156127c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127e69190615d4e565b613eed565b905060005b876020015151811015612a6757612816886020015182815181106107ce576107ce61576f565b8360200151828151811061282c5761282c61576f565b602090810291909101015280156128ec57602083015161284d600183615a54565b8151811061285d5761285d61576f565b602002602001015160001c8360200151828151811061287e5761287e61576f565b602002602001015160001c116128ec576040805162461bcd60e51b8152602060048201526024810191909152600080516020615f2383398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610ae8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106129315761293161576f565b60200260200101518b8b6000015185815181106129505761295061576f565b60200260200101516040518463ffffffff1660e01b815260040161298d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156129aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129ce9190615c7d565b6001600160c01b0316836000015182815181106129ed576129ed61576f565b602002602001018181525050612a53610eb6612a278486600001518581518110612a1957612a1961576f565b602002602001015116613f77565b8a602001518481518110612a3d57612a3d61576f565b6020026020010151613fa290919063ffffffff16565b945080612a5f8161579b565b9150506127f0565b5050612a7283614086565b60975490935060ff16600081612a89576000612b0b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ae7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b0b9190615a6b565b905060005b8a811015613189578215612c6b578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612b6757612b6761576f565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612ba7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bcb9190615a6b565b612bd59190615d6b565b11612c6b5760405162461bcd60e51b81526020600482015260666024820152600080516020615f2383398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610ae8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612cac57612cac61576f565b9050013560f81c60f81b60f81c8c8c60a001518581518110612cd057612cd061576f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612d2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d509190615d83565b6001600160401b031916612d738a6040015183815181106107ce576107ce61576f565b67ffffffffffffffff191614612e0f5760405162461bcd60e51b81526020600482015260616024820152600080516020615f2383398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610ae8565b612e3f89604001518281518110612e2857612e2861576f565b602002602001015187613ae590919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612e8257612e8261576f565b9050013560f81c60f81b60f81c8c8c60c001518581518110612ea657612ea661576f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612f02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f269190615b14565b85602001518281518110612f3c57612f3c61576f565b6001600160601b03909216602092830291909101820152850151805182908110612f6857612f6861576f565b602002602001015185600001518281518110612f8657612f8661576f565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561317457612ffe86600001518281518110612fd057612fd061576f565b60200260200101518f8f86818110612fea57612fea61576f565b600192013560f81c9290921c811614919050565b15613162577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106130445761304461576f565b9050013560f81c60f81b60f81c8e896020015185815181106130685761306861576f565b60200260200101518f60e0015188815181106130865761308661576f565b6020026020010151878151811061309f5761309f61576f565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015613103573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131279190615b14565b875180518590811061313b5761313b61576f565b6020026020010181815161314f9190615dae565b6001600160601b03169052506001909101905b8061316c8161579b565b915050612faa565b505080806131819061579b565b915050612b10565b5050506000806131a38c868a606001518b60800151610d9b565b91509150816132145760405162461bcd60e51b81526020600482015260436024820152600080516020615f2383398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610ae8565b806132755760405162461bcd60e51b81526020600482015260396024820152600080516020615f2383398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610ae8565b505060008782602001516040516020016132909291906157b6565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6132c2614121565b6132cc60006139fc565b565b60cb546001600160a01b031633146132e557600080fd5b60006132f760c0850160a08601614ad5565b905036600061330960c0870187615dd6565b90925090506000613321610100880160e08901614ad5565b90506000866040516020016133369190615e1c565b60405160208183030381529060405280519060200120905060008061335e8387878a8c6123c2565b9150915060005b858110156133f2578460ff16836020015182815181106133875761338761576f565b60200260200101516133999190615e2a565b6001600160601b03166064846000015183815181106133ba576133ba61576f565b60200260200101516001600160601b03166133d59190615e59565b10156133e057600080fd5b806133ea8161579b565b915050613365565b5060408051808201825263ffffffff431681526020810183905260cd549151636a3a9c6160e01b815290916001600160a01b031690636a3a9c619061343f908e908e908690600401615e78565b600060405180830381600087803b15801561345957600080fd5b505af115801561346d573d6000803e3d6000fd5b505050507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a826040516134a2929190615eb6565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106134f2576134f261576f565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061352e9088908690600401615d2f565b600060405180830381865afa15801561354b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526135739190810190615bcf565b6000815181106135855761358561576f565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156135f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136159190615c7d565b6001600160c01b03169050600061362b8261417b565b9050816136398a838a611119565b9550955050505050935093915050565b613651614121565b6001600160a01b0381166136b65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610ae8565b610afa816139fc565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613712573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061373691906157fe565b6001600160a01b0316336001600160a01b0316146137665760405162461bcd60e51b8152600401610ae89061581b565b6066541981196066541916146137e45760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610ae8565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610c31565b6001600160a01b0381166138a95760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610ae8565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b031615801561393357506001600160a01b03821615155b6139b55760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610ae8565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26139f88261381b565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613a6a614533565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613a9d57613a9f565bfe5b5080613add5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610ae8565b505092915050565b6040805180820190915260008082526020820152613b01614551565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613a9d575080613add5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610ae8565b613b8161456f565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613c69600080516020615f03833981519152866158ca565b90505b613c7581614247565b9093509150600080516020615f03833981519152828309831415613caf576040805180820190915290815260208101919091529392505050565b600080516020615f03833981519152600182089050613c6c565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613cfb614594565b60005b6002811015613ec0576000613d14826006615e59565b9050848260028110613d2857613d2861576f565b60200201515183613d3a836000615d6b565b600c8110613d4a57613d4a61576f565b6020020152848260028110613d6157613d6161576f565b60200201516020015183826001613d789190615d6b565b600c8110613d8857613d8861576f565b6020020152838260028110613d9f57613d9f61576f565b6020020151515183613db2836002615d6b565b600c8110613dc257613dc261576f565b6020020152838260028110613dd957613dd961576f565b6020020151516001602002015183613df2836003615d6b565b600c8110613e0257613e0261576f565b6020020152838260028110613e1957613e1961576f565b602002015160200151600060028110613e3457613e3461576f565b602002015183613e45836004615d6b565b600c8110613e5557613e5561576f565b6020020152838260028110613e6c57613e6c61576f565b602002015160200151600160028110613e8757613e8761576f565b602002015183613e98836005615d6b565b600c8110613ea857613ea861576f565b60200201525080613eb88161579b565b915050613cfe565b50613ec96145b3565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613ef9846142c9565b9050808360ff166001901b116123b95760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610ae8565b6000805b82156123bc57613f8c600184615a54565b9092169180613f9a81615ee0565b915050613f7b565b60408051808201909152600080825260208201526102008261ffff1610613ffe5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610ae8565b8161ffff16600114156140125750816123bc565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061407b57600161ffff871660ff83161c8116141561405e5761405b8484613ae5565b93505b6140688384613ae5565b92506201fffe600192831b16910161402e565b509195945050505050565b604080518082019091526000808252602082015281511580156140ab57506020820151155b156140c9575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615f0383398151915284602001516140fc91906158ca565b61411490600080516020615f03833981519152615a54565b905292915050565b919050565b6033546001600160a01b031633146132cc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610ae8565b606060008061418984613f77565b61ffff166001600160401b038111156141a4576141a4614754565b6040519080825280601f01601f1916602001820160405280156141ce576020820181803683370190505b5090506000805b8251821080156141e6575061010081105b1561423d576001811b93508584161561422d578060f81b83838151811061420f5761420f61576f565b60200101906001600160f81b031916908160001a9053508160010191505b6142368161579b565b90506141d5565b5090949350505050565b60008080600080516020615f038339815191526003600080516020615f0383398151915286600080516020615f038339815191528889090908905060006142bd827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615f03833981519152614456565b91959194509092505050565b6000610100825111156143525760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610ae8565b815161436057506000919050565b600080836000815181106143765761437661576f565b0160200151600160f89190911c81901b92505b845181101561444d578481815181106143a4576143a461576f565b0160200151600160f89190911c1b91508282116144395760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610ae8565b918117916144468161579b565b9050614389565b50909392505050565b6000806144616145b3565b6144696145d1565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613a9d5750826144f35760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610ae8565b505195945050505050565b6040518060a001604052806145116145ef565b8152600060208201819052606060408301819052820181905260809091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061458261460d565b815260200161458f61460d565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b63ffffffff81168114610afa57600080fd5b803561411c8161462b565b60008083601f84011261465a57600080fd5b5081356001600160401b0381111561467157600080fd5b60208301915083602082850101111561468957600080fd5b9250929050565b8015158114610afa57600080fd5b803561411c81614690565b600080600080600061010086880312156146c257600080fd5b60a08601878111156146d357600080fd5b869550356146e08161462b565b935060c08601356001600160401b038111156146fb57600080fd5b61470788828901614648565b90945092505060e086013561471b81614690565b809150509295509295909350565b6000610120828403121561473c57600080fd5b50919050565b60006040828403121561473c57600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561478c5761478c614754565b60405290565b60405161010081016001600160401b038111828210171561478c5761478c614754565b604051601f8201601f191681016001600160401b03811182821017156147dd576147dd614754565b604052919050565b60006001600160401b038211156147fe576147fe614754565b5060051b60200190565b60006040828403121561481a57600080fd5b61482261476a565b9050813581526020820135602082015292915050565b600082601f83011261484957600080fd5b8135602061485e614859836147e5565b6147b5565b82815260069290921b8401810191818101908684111561487d57600080fd5b8286015b848110156148a1576148938882614808565b835291830191604001614881565b509695505050505050565b60008060008060c085870312156148c257600080fd5b84356001600160401b03808211156148d957600080fd5b6148e588838901614729565b95506148f48860208901614742565b94506149038860608901614742565b935060a087013591508082111561491957600080fd5b5061492687828801614838565b91505092959194509250565b6001600160a01b0381168114610afa57600080fd5b60006020828403121561495957600080fd5b81356123b981614932565b60006020828403121561497657600080fd5b5035919050565b600080600080600060a0868803121561499557600080fd5b85356149a081614932565b945060208601356149b081614932565b935060408601356149c081614932565b925060608601356149d081614932565b9150608086013561471b81614932565b600082601f8301126149f157600080fd5b604051604081018181106001600160401b0382111715614a1357614a13614754565b8060405250806040840185811115614a2a57600080fd5b845b8181101561407b578035835260209283019201614a2c565b600060808284031215614a5657600080fd5b614a5e61476a565b9050614a6a83836149e0565b8152614a7983604084016149e0565b602082015292915050565b6000806000806101208587031215614a9b57600080fd5b84359350614aac8660208701614808565b9250614abb8660608701614a44565b9150614aca8660e08701614808565b905092959194509250565b600060208284031215614ae757600080fd5b81356123b98161462b565b60008083601f840112614b0457600080fd5b5081356001600160401b03811115614b1b57600080fd5b6020830191508360208260051b850101111561468957600080fd5b60008060008060008060808789031215614b4f57600080fd5b86356001600160401b0380821115614b6657600080fd5b614b728a838b01614729565b975060208901359150614b848261462b565b90955060408801359080821115614b9a57600080fd5b614ba68a838b01614af2565b90965094506060890135915080821115614bbf57600080fd5b50614bcc89828a01614648565b979a9699509497509295939492505050565b60008060408385031215614bf157600080fd5b8235614bfc81614932565b91506020838101356001600160401b03811115614c1857600080fd5b8401601f81018613614c2957600080fd5b8035614c37614859826147e5565b81815260059190911b82018301908381019088831115614c5657600080fd5b928401925b82841015614c7d578335614c6e81614932565b82529284019290840190614c5b565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614cbc57815187529582019590820190600101614ca0565b509495945050505050565b6020815260006123b96020830184614c8c565b600080600060608486031215614cef57600080fd5b8335614cfa81614932565b92506020848101356001600160401b0380821115614d1757600080fd5b818701915087601f830112614d2b57600080fd5b813581811115614d3d57614d3d614754565b614d4f601f8201601f191685016147b5565b91508082528884828501011115614d6557600080fd5b8084840185840137600084828401015250809450505050614d886040850161463d565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614e27578385038a52825180518087529087019087870190845b81811015614e1257835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614dce565b50509a87019a95505091850191600101614db0565b509298975050505050505050565b6020815260006123b96020830184614d91565b600080600060a08486031215614e5d57600080fd5b83356001600160401b03811115614e7357600080fd5b614e7f86828701614729565b935050614e8f8560208601614742565b9150614d888560608601614742565b600060208284031215614eb057600080fd5b81356123b981614690565b600082601f830112614ecc57600080fd5b81356020614edc614859836147e5565b82815260059290921b84018101918181019086841115614efb57600080fd5b8286015b848110156148a15780358352918301918301614eff565b60008060408385031215614f2957600080fd5b8235614f3481614932565b915060208301356001600160401b03811115614f4f57600080fd5b614f5b85828601614ebb565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614fa65783516001600160a01b031683529284019291840191600101614f81565b50909695505050505050565b60008060008060008060808789031215614fcb57600080fd5b8635614fd681614932565b95506020870135614fe68161462b565b945060408701356001600160401b038082111561500257600080fd5b61500e8a838b01614648565b9096509450606089013591508082111561502757600080fd5b50614bcc89828a01614af2565b600081518084526020808501945080840160005b83811015614cbc57815163ffffffff1687529582019590820190600101615048565b60006020808352835160808285015261508660a0850182615034565b905081850151601f19808684030160408701526150a38383615034565b925060408701519150808684030160608701526150c08383615034565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156151175784878303018452615105828751615034565b958801959388019391506001016150eb565b509998505050505050505050565b60008060008060006060868803121561513d57600080fd5b85356151488161462b565b945060208601356001600160401b038082111561516457600080fd5b61517089838a01614af2565b9096509450604088013591508082111561518957600080fd5b5061519688828901614648565b969995985093965092949392505050565b60ff81168114610afa57600080fd5b6000602082840312156151c857600080fd5b81356123b9816151a7565b6000806000606084860312156151e857600080fd5b83356151f381614932565b925060208401356001600160401b0381111561520e57600080fd5b61521a86828701614ebb565b925050604084013561522b8161462b565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614fa657835183529284019291840191600101615252565b6000806040838503121561528157600080fd5b823561528c8161462b565b946020939093013593505050565b600082601f8301126152ab57600080fd5b813560206152bb614859836147e5565b82815260059290921b840181019181810190868411156152da57600080fd5b8286015b848110156148a15780356152f18161462b565b83529183019183016152de565b600082601f83011261530f57600080fd5b8135602061531f614859836147e5565b82815260059290921b8401810191818101908684111561533e57600080fd5b8286015b848110156148a15780356001600160401b038111156153615760008081fd5b61536f8986838b010161529a565b845250918301918301615342565b6000610180828403121561539057600080fd5b615398614792565b905081356001600160401b03808211156153b157600080fd5b6153bd8583860161529a565b835260208401359150808211156153d357600080fd5b6153df85838601614838565b602084015260408401359150808211156153f857600080fd5b61540485838601614838565b60408401526154168560608601614a44565b60608401526154288560e08601614808565b608084015261012084013591508082111561544257600080fd5b61544e8583860161529a565b60a084015261014084013591508082111561546857600080fd5b6154748583860161529a565b60c084015261016084013591508082111561548e57600080fd5b5061549b848285016152fe565b60e08301525092915050565b6000806000806000608086880312156154bf57600080fd5b8535945060208601356001600160401b03808211156154dd57600080fd5b6154e989838a01614648565b9096509450604088013591506154fe8261462b565b9092506060870135908082111561551457600080fd5b506155218882890161537d565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614cbc5781516001600160601b031687529582019590820190600101615542565b6040815260008351604080840152615582608084018261552e565b90506020850151603f1984830301606085015261559f828261552e565b925050508260208301529392505050565b6000806000608084860312156155c557600080fd5b83356001600160401b03808211156155dc57600080fd5b6155e887838801614729565b94506155f78760208801614742565b9350606086013591508082111561560d57600080fd5b5061561a8682870161537d565b9150509250925092565b60008060006060848603121561563957600080fd5b833561564481614932565b925060208401359150604084013561522b8161462b565b8281526040602082015260006156746040830184614d91565b949350505050565b6000815180845260005b818110156156a257602081850181015186830182015201615686565b818111156156b4576000602083870101525b50601f01601f19169290920160200192915050565b6020808252825160009190828483015b60058210156156f85782518152918301916001919091019083016156d9565b50505083015163ffffffff1660c0830152604083015161012060e0840181905261572661014085018361567c565b9150606085015161574061010086018263ffffffff169052565b5060808501518015158583015261423d565b60006020828403121561576457600080fd5b81516123b98161462b565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156157af576157af615785565b5060010190565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156157f1578151855293820193908201906001016157d5565b5092979650505050505050565b60006020828403121561581057600080fd5b81516123b981614932565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561587757600080fd5b81516123b981614690565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000826158e757634e487b7160e01b600052601260045260246000fd5b500690565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061012060a083853760a08401600080825260a08501356159368161462b565b63ffffffff1690915260c08401359036859003601e19018212615957578081fd5b9084019081356001600160401b03811115615970578182fd5b80360386131561597e578182fd5b8360c088015261599484880182602086016158ec565b93505050506159a560e0840161463d565b63ffffffff1660e08501526101006159be84820161469e565b8015158683015261423d565b81835260006001600160fb1b038311156159e357600080fd5b8260051b8083602087013760009401602001938452509192915050565b608081526000615a136080830189615915565b63ffffffff881660208401528281036040840152615a328187896159ca565b90508281036060840152615a478185876158ec565b9998505050505050505050565b600082821015615a6657615a66615785565b500390565b600060208284031215615a7d57600080fd5b5051919050565b60006020808385031215615a9757600080fd5b82516001600160401b03811115615aad57600080fd5b8301601f81018513615abe57600080fd5b8051615acc614859826147e5565b81815260059190911b82018301908381019087831115615aeb57600080fd5b928401925b82841015615b0957835182529284019290840190615af0565b979650505050505050565b600060208284031215615b2657600080fd5b81516001600160601b03811681146123b957600080fd5b8035615b488161462b565b63ffffffff168252602090810135910152565b60a081526000615b6e60a0830186615915565b9050615b7d6020830185615b3d565b8235615b888161462b565b63ffffffff1660608301526020929092013560809091015292915050565b63ffffffff84168152604060208201526000615bc66040830184866159ca565b95945050505050565b60006020808385031215615be257600080fd5b82516001600160401b03811115615bf857600080fd5b8301601f81018513615c0957600080fd5b8051615c17614859826147e5565b81815260059190911b82018301908381019087831115615c3657600080fd5b928401925b82841015615b09578351615c4e8161462b565b82529284019290840190615c3b565b63ffffffff84168152604060208201526000615bc66040830184866158ec565b600060208284031215615c8f57600080fd5b81516001600160c01b03811681146123b957600080fd5b600060ff821660ff811415615cbd57615cbd615785565b60010192915050565b604081526000615cda6040830185876158ec565b905063ffffffff83166020830152949350505050565b63ffffffff86168152606060208201526000615d106060830186886159ca565b8281036040840152615d238185876158ec565b98975050505050505050565b63ffffffff831681526040602082015260006156746040830184614c8c565b600060208284031215615d6057600080fd5b81516123b9816151a7565b60008219821115615d7e57615d7e615785565b500190565b600060208284031215615d9557600080fd5b815167ffffffffffffffff19811681146123b957600080fd5b60006001600160601b0383811690831681811015615dce57615dce615785565b039392505050565b6000808335601e19843603018112615ded57600080fd5b8301803591506001600160401b03821115615e0757600080fd5b60200191503681900382131561468957600080fd5b604081016123bc8284615b3d565b60006001600160601b0380831681851681830481118215151615615e5057615e50615785565b02949350505050565b6000816000190483118215151615615e7357615e73615785565b500290565b60a081526000615e8b60a0830186615915565b9050615e9a6020830185615b3d565b825163ffffffff16606083015260208301516080830152615674565b60808101615ec48285615b3d565b825163ffffffff166040830152602083015160608301526115aa565b600061ffff80831681811415615ef857615ef8615785565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122021ee74f0fa27c9705e624a35ab88b5dd3a0d4d362830bb90abb84d350b90dc0164736f6c634300080c0033",
}

// ContractZklayerTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractZklayerTaskManagerMetaData.ABI instead.
var ContractZklayerTaskManagerABI = ContractZklayerTaskManagerMetaData.ABI

// ContractZklayerTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractZklayerTaskManagerMetaData.Bin instead.
var ContractZklayerTaskManagerBin = ContractZklayerTaskManagerMetaData.Bin

// DeployContractZklayerTaskManager deploys a new Ethereum contract, binding an instance of ContractZklayerTaskManager to it.
func DeployContractZklayerTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address) (common.Address, *types.Transaction, *ContractZklayerTaskManager, error) {
	parsed, err := ContractZklayerTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractZklayerTaskManagerBin), backend, _registryCoordinator)
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

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x09b88097.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) ConfirmChallenge(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "confirmChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x09b88097.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) ConfirmChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ConfirmChallenge(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x09b88097.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) ConfirmChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.ConfirmChallenge(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x027e4a37.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "createNewTask", inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x027e4a37.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.CreateNewTask(&_ContractZklayerTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x027e4a37.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.CreateNewTask(&_ContractZklayerTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce)
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

// RaiseChallenge is a paid mutator transaction binding the contract method 0x369bcec5.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) RaiseChallenge(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "raiseChallenge", task, taskResponse, taskResponseMetadata)
}

// RaiseChallenge is a paid mutator transaction binding the contract method 0x369bcec5.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) RaiseChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RaiseChallenge(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
}

// RaiseChallenge is a paid mutator transaction binding the contract method 0x369bcec5.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) RaiseChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RaiseChallenge(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
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

// RespondToTask is a paid mutator transaction binding the contract method 0x79e3fb33.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x79e3fb33.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) RespondToTask(task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RespondToTask(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x79e3fb33.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) RespondToTask(task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RespondToTask(&_ContractZklayerTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x2ced7fa6.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactor) RespondToTaskWithProof(opts *bind.TransactOpts, task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.contract.Transact(opts, "respondToTaskWithProof", task, taskIndex, instances, proof)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x2ced7fa6.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerSession) RespondToTaskWithProof(task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RespondToTaskWithProof(&_ContractZklayerTaskManager.TransactOpts, task, taskIndex, instances, proof)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x2ced7fa6.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerTransactorSession) RespondToTaskWithProof(task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractZklayerTaskManager.Contract.RespondToTaskWithProof(&_ContractZklayerTaskManager.TransactOpts, task, taskIndex, instances, proof)
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
	Task      ITaskStructTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x02fbf40079c9022bc174c39bce245cecd33b518a66e442ae48699be1ab5671ee.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool) task)
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

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x02fbf40079c9022bc174c39bce245cecd33b518a66e442ae48699be1ab5671ee.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool) task)
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x02fbf40079c9022bc174c39bce245cecd33b518a66e442ae48699be1ab5671ee.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool) task)
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
	TaskResponse         ITaskStructTaskResponse
	TaskResponseMetadata ITaskStructTaskResponseMetadata
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

// ContractZklayerTaskManagerTaskRespondedWithProofIterator is returned from FilterTaskRespondedWithProof and is used to iterate over the raw logs and unpacked data for TaskRespondedWithProof events raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskRespondedWithProofIterator struct {
	Event *ContractZklayerTaskManagerTaskRespondedWithProof // Event containing the contract specifics and raw log

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
func (it *ContractZklayerTaskManagerTaskRespondedWithProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerTaskManagerTaskRespondedWithProof)
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
		it.Event = new(ContractZklayerTaskManagerTaskRespondedWithProof)
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
func (it *ContractZklayerTaskManagerTaskRespondedWithProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerTaskManagerTaskRespondedWithProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerTaskManagerTaskRespondedWithProof represents a TaskRespondedWithProof event raised by the ContractZklayerTaskManager contract.
type ContractZklayerTaskManagerTaskRespondedWithProof struct {
	TaskIndex uint32
	Output    *big.Int
	Prover    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskRespondedWithProof is a free log retrieval operation binding the contract event 0x4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada.
//
// Solidity: event TaskRespondedWithProof(uint32 indexed taskIndex, uint256 output, address indexed prover)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) FilterTaskRespondedWithProof(opts *bind.FilterOpts, taskIndex []uint32, prover []common.Address) (*ContractZklayerTaskManagerTaskRespondedWithProofIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.FilterLogs(opts, "TaskRespondedWithProof", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerTaskManagerTaskRespondedWithProofIterator{contract: _ContractZklayerTaskManager.contract, event: "TaskRespondedWithProof", logs: logs, sub: sub}, nil
}

// WatchTaskRespondedWithProof is a free log subscription operation binding the contract event 0x4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada.
//
// Solidity: event TaskRespondedWithProof(uint32 indexed taskIndex, uint256 output, address indexed prover)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) WatchTaskRespondedWithProof(opts *bind.WatchOpts, sink chan<- *ContractZklayerTaskManagerTaskRespondedWithProof, taskIndex []uint32, prover []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractZklayerTaskManager.contract.WatchLogs(opts, "TaskRespondedWithProof", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerTaskManagerTaskRespondedWithProof)
				if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskRespondedWithProof", log); err != nil {
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

// ParseTaskRespondedWithProof is a log parse operation binding the contract event 0x4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada.
//
// Solidity: event TaskRespondedWithProof(uint32 indexed taskIndex, uint256 output, address indexed prover)
func (_ContractZklayerTaskManager *ContractZklayerTaskManagerFilterer) ParseTaskRespondedWithProof(log types.Log) (*ContractZklayerTaskManagerTaskRespondedWithProof, error) {
	event := new(ContractZklayerTaskManagerTaskRespondedWithProof)
	if err := _ContractZklayerTaskManager.contract.UnpackLog(event, "TaskRespondedWithProof", log); err != nil {
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
