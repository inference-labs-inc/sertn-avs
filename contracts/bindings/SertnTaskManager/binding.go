// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSertnTaskManager

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

// ContractSertnTaskManagerMetaData contains all meta data concerning the ContractSertnTaskManager contract.
var ContractSertnTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeInstances\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"provenOnResponce\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveResultAccurate\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"raiseChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTaskWithProof\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"taskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNewAggregator\",\"inputs\":[{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNewInferenceDB\",\"inputs\":[{\"name\":\"_inferenceDB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskRespondedWithProof\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005ff338038062005ff38339810160408190526200003591620001ed565b80806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001ed565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001ed565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b39190620001ed565b6001600160a01b031660e05250506097805460ff1916600117905562000214565b6001600160a01b0381168114620001ea57600080fd5b50565b6000602082840312156200020057600080fd5b81516200020d81620001d4565b9392505050565b60805160a05160c05160e051615d666200028d60003960008181610592015261298d015260008181610469015281816108dc0152612b6f0152600081816104a301528181612d450152612f070152600081816104ca0152818161151d01528181612677015281816127f00152612a2a0152615d666000f3fe608060405234801561001057600080fd5b50600436106102275760003560e01c80635ac86ab71161013057806379e3fb33116100b8578063cefdc1d41161007c578063cefdc1d41461056c578063df5cf7231461058d578063f1e1011d146105b4578063f2fde38b146105c7578063fabc1cbc146105da57600080fd5b806379e3fb33146105155780637afa1eed14610528578063886f11951461053b5780638da5cb5b1461054e578063b98d09081461055f57600080fd5b806367d6be44116100ff57806367d6be441461048b578063683048351461049e5780636d14a987146104c55780636efb4636146104ec578063715018a61461050d57600080fd5b80635ac86ab7146104095780635c1556621461043c5780635c975abb1461045c5780635df459461461046457600080fd5b806331b36bd9116101b35780634465027511610182578063446502751461039b5780634d2b57fe146103ae5780634f739f74146103ce57806358a7cd26146103ee578063595c6a671461040157600080fd5b806331b36bd9146103355780633563b0d114610355578063369bcec514610375578063416c7e5e1461038857600080fd5b8063171f1d5b116101fa578063171f1d5b1461027a578063245a7bfc146102a95780632cb223d5146102d45780632ced7fa6146103025780632d89f6fc1461031557600080fd5b8063027e4a371461022c57806309b880971461024157806310d67a2f14610254578063136439dd14610267575b600080fd5b61023f61023a3660046144fa565b6105ed565b005b61023f61024f3660046146fd565b610773565b61023f610262366004614798565b610a5f565b61023f6102753660046147b5565b610b1b565b61028d610288366004614872565b610c5a565b6040805192151583529015156020830152015b60405180910390f35b60cb546102bc906001600160a01b031681565b6040516001600160a01b0390911681526020016102a0565b6102f46102e23660046148c3565b60ca6020526000908152604090205481565b6040519081526020016102a0565b61023f610310366004614924565b610de4565b6102f46103233660046148c3565b60c96020526000908152604090205481565b6103486103433660046149cc565b610ebc565b6040516102a09190614ab5565b610368610363366004614ac8565b610fd8565b6040516102a09190614c23565b61023f610383366004614c36565b611470565b61023f610396366004614c8c565b61151b565b61023f6103a9366004614798565b611690565b6103c16103bc366004614d04565b6116d3565b6040516102a09190614d53565b6103e16103dc366004614da0565b6117e8565b6040516102a09190614e58565b61023f6103fc366004614f13565b611f0e565b61023f611fb2565b61042c610417366004614fa4565b606654600160ff9092169190911b9081161490565b60405190151581526020016102a0565b61044f61044a366004614fc1565b612079565b6040516102a09190615024565b6066546102f4565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102f461049936600461505c565b612241565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6104ff6104fa366004615295565b6122c4565b6040516102a0929190615355565b61023f6131bc565b61023f61052336600461539e565b6131d0565b60cc546102bc906001600160a01b031681565b6065546102bc906001600160a01b031681565b6033546001600160a01b03166102bc565b60975461042c9060ff1681565b61057f61057a366004615412565b6133b9565b6040516102a0929190615449565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b61023f6105c2366004614798565b61354b565b61023f6105d5366004614798565b613584565b61023f6105e83660046147b5565b6135fa565b60cc546001600160a01b0316331461060457600080fd5b61060c61434f565b6040805160a081810190925290879060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908616606083015260408051601f86018390048302810183019091528481529085908590819084018382808284376000920182905250604080870195909552861515608087015260cd54945190946001600160a01b03169350633b7d2cee92506106b3915085906020016154b7565b604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004016106e791815260200190565b6020604051808303816000875af1158015610706573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072a9190615540565b90508063ffffffff167f02fbf40079c9022bc174c39bce245cecd33b518a66e442ae48699be1ab5671ee8360405161076291906154b7565b60405180910390a250505050505050565b600061078260208501856148c3565b9050600082516001600160401b0381111561079f5761079f6145a5565b6040519080825280602002602001820160405280156107c8578160200160208202803683370190505b50905060005b835181101561083a5761080b8482815181106107ec576107ec61555d565b6020026020010151805160009081526020918201519091526040902090565b82828151811061081d5761081d61555d565b60209081029190910101528061083281615589565b9150506107ce565b50600061084d60c0880160a089016148c3565b8260405160200161085f9291906155a4565b6040516020818303038152906040528051906020012090508460200135811461088757600080fd5b600084516001600160401b038111156108a2576108a26145a5565b6040519080825280602002602001820160405280156108cb578160200160208202803683370190505b50905060005b85518110156109be577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae685838151811061091b5761091b61555d565b60200260200101516040518263ffffffff1660e01b815260040161094191815260200190565b602060405180830381865afa15801561095e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098291906155ec565b8282815181106109945761099461555d565b6001600160a01b0390921660209283029190910190910152806109b681615589565b9150506108d1565b5060cd5460405163959b127960e01b815263ffffffff861660048201526001600160a01b039091169063959b127990602401600060405180830381600087803b158015610a0a57600080fd5b505af1158015610a1e573d6000803e3d6000fd5b505060405133925063ffffffff871691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ab2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad691906155ec565b6001600160a01b0316336001600160a01b031614610b0f5760405162461bcd60e51b8152600401610b0690615609565b60405180910390fd5b610b1881613756565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610b63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b879190615653565b610ba35760405162461bcd60e51b8152600401610b0690615670565b60665481811614610c1c5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610b06565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610ca257610ca261555d565b60200201518951600160200201518a60200151600060028110610cc757610cc761555d565b60200201518b60200151600160028110610ce357610ce361555d565b602090810291909101518c518d830151604051610d409a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d6391906156b8565b9050610dd6610d7c610d75888461384d565b86906138e4565b610d84613978565b610dcc610dbd85610db7604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061384d565b610dc68c613a38565b906138e4565b886201d4c0613ac8565b909890975095505050505050565b60cd54604051631676bfd360e11b81526001600160a01b0390911690632ced7fa690610e1e908990899089908990899089906004016157ee565b600060405180830381600087803b158015610e3857600080fd5b505af1158015610e4c573d6000803e3d6000fd5b50339250505063ffffffff86167f4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada8686610e87600182615842565b818110610e9657610e9661555d565b90506020020135604051610eac91815260200190565b60405180910390a3505050505050565b606081516001600160401b03811115610ed757610ed76145a5565b604051908082528060200260200182016040528015610f00578160200160208202803683370190505b50905060005b8251811015610fd157836001600160a01b03166313542a4e848381518110610f3057610f3061555d565b60200260200101516040518263ffffffff1660e01b8152600401610f6391906001600160a01b0391909116815260200190565b602060405180830381865afa158015610f80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa49190615859565b828281518110610fb657610fb661555d565b6020908102919091010152610fca81615589565b9050610f06565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561101a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103e91906155ec565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015611080573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a491906155ec565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110a91906155ec565b9050600086516001600160401b03811115611127576111276145a5565b60405190808252806020026020018201604052801561115a57816020015b60608152602001906001900390816111455790505b50905060005b875181101561146257600088828151811061117d5761117d61555d565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa1580156111de573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112069190810190615872565b905080516001600160401b03811115611221576112216145a5565b60405190808252806020026020018201604052801561126c57816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161123f5790505b5084848151811061127f5761127f61555d565b602002602001018190525060005b815181101561144c576040518060600160405280876001600160a01b03166347b314e88585815181106112c2576112c261555d565b60200260200101516040518263ffffffff1660e01b81526004016112e891815260200190565b602060405180830381865afa158015611305573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132991906155ec565b6001600160a01b031681526020018383815181106113495761134961555d565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106113775761137761555d565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa1580156113d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f79190615902565b6001600160601b03168152508585815181106114155761141561555d565b6020026020010151828151811061142e5761142e61555d565b6020026020010181905250808061144490615589565b91505061128d565b505050808061145a90615589565b915050611160565b5093505050505b9392505050565b60cd54604051630b5238df60e41b81526001600160a01b039091169063b5238df0906114a490869086908690600401615949565b600060405180830381600087803b1580156114be57600080fd5b505af11580156114d2573d6000803e3d6000fd5b506114e49250505060208301836148c3565b63ffffffff167f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a60405160405180910390a2505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611579573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159d91906155ec565b6001600160a01b0316336001600160a01b0316146116495760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610b06565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60cb546001600160a01b031633146116a757600080fd5b60cb80546001600160a01b039092166001600160a01b0319928316811790915560cc8054909216179055565b606081516001600160401b038111156116ee576116ee6145a5565b604051908082528060200260200182016040528015611717578160200160208202803683370190505b50905060005b8251811015610fd157836001600160a01b031663296bb0648483815181106117475761174761555d565b60200260200101516040518263ffffffff1660e01b815260040161176d91815260200190565b602060405180830381865afa15801561178a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ae91906155ec565b8282815181106117c0576117c061555d565b6001600160a01b03909216602092830291909101909101526117e181615589565b905061171d565b6118136040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611853573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187791906155ec565b90506118a46040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906118d4908b9089908990600401615994565b600060405180830381865afa1580156118f1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261191991908101906159bd565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061194b908b908b908b90600401615a4b565b600060405180830381865afa158015611968573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261199091908101906159bd565b6040820152856001600160401b038111156119ad576119ad6145a5565b6040519080825280602002602001820160405280156119e057816020015b60608152602001906001900390816119cb5790505b50606082015260005b60ff8116871115611e1f576000856001600160401b03811115611a0e57611a0e6145a5565b604051908082528060200260200182016040528015611a37578160200160208202803683370190505b5083606001518360ff1681518110611a5157611a5161555d565b602002602001018190525060005b86811015611d1f5760008c6001600160a01b03166304ec63518a8a85818110611a8a57611a8a61555d565b905060200201358e88600001518681518110611aa857611aa861555d565b60200260200101516040518463ffffffff1660e01b8152600401611ae59392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611b02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b269190615a6b565b90506001600160c01b038116611bca5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610b06565b8a8a8560ff16818110611bdf57611bdf61555d565b6001600160c01b03841692013560f81c9190911c600190811614159050611d0c57856001600160a01b031663dd9846b98a8a85818110611c2157611c2161555d565b905060200201358d8d8860ff16818110611c3d57611c3d61555d565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611c93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cb79190615540565b85606001518560ff1681518110611cd057611cd061555d565b60200260200101518481518110611ce957611ce961555d565b63ffffffff9092166020928302919091019091015282611d0881615589565b9350505b5080611d1781615589565b915050611a5f565b506000816001600160401b03811115611d3a57611d3a6145a5565b604051908082528060200260200182016040528015611d63578160200160208202803683370190505b50905060005b82811015611de45784606001518460ff1681518110611d8a57611d8a61555d565b60200260200101518181518110611da357611da361555d565b6020026020010151828281518110611dbd57611dbd61555d565b63ffffffff9092166020928302919091019091015280611ddc81615589565b915050611d69565b508084606001518460ff1681518110611dff57611dff61555d565b602002602001018190525050508080611e1790615a94565b9150506119e9565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e8491906155ec565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611eb7908b908b908e90600401615ab4565b600060405180830381865afa158015611ed4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611efc91908101906159bd565b60208301525098975050505050505050565b60cd54604051632c53e69360e11b81526001600160a01b03909116906358a7cd2690611f469088908890889088908890600401615ade565b600060405180830381600087803b158015611f6057600080fd5b505af1158015611f74573d6000803e3d6000fd5b505060405133925063ffffffff881691507ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611ffa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061201e9190615653565b61203a5760405162461bcd60e51b8152600401610b0690615670565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016120ab929190615b1d565b600060405180830381865afa1580156120c8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120f091908101906159bd565b9050600084516001600160401b0381111561210d5761210d6145a5565b604051908082528060200260200182016040528015612136578160200160208202803683370190505b50905060005b855181101561223757866001600160a01b03166304ec63518783815181106121665761216661555d565b6020026020010151878685815181106121815761218161555d565b60200260200101516040518463ffffffff1660e01b81526004016121be9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156121db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ff9190615a6b565b6001600160c01b031682828151811061221a5761221a61555d565b60209081029190910101528061222f81615589565b91505061213c565b5095945050505050565b60cd546040516319f5af9160e21b815263ffffffff84166004820152602481018390526000916001600160a01b0316906367d6be4490604401602060405180830381865afa158015612297573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122bb9190615859565b90505b92915050565b604080518082019091526060808252602082015260008461233b5760405162461bcd60e51b81526020600482015260376024820152600080516020615d1183398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610b06565b60408301515185148015612353575060a08301515185145b8015612363575060c08301515185145b8015612373575060e08301515185145b6123dd5760405162461bcd60e51b81526020600482015260416024820152600080516020615d1183398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610b06565b825151602084015151146124555760405162461bcd60e51b815260206004820152604460248201819052600080516020615d11833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610b06565b4363ffffffff168463ffffffff16106124c45760405162461bcd60e51b815260206004820152603c6024820152600080516020615d1183398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610b06565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115612505576125056145a5565b60405190808252806020026020018201604052801561252e578160200160208202803683370190505b506020820152866001600160401b0381111561254c5761254c6145a5565b604051908082528060200260200182016040528015612575578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156125a9576125a96145a5565b6040519080825280602002602001820160405280156125d2578160200160208202803683370190505b5081526020860151516001600160401b038111156125f2576125f26145a5565b60405190808252806020026020018201604052801561261b578160200160208202803683370190505b50816020018190525060006126ed8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156126c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126e89190615b3c565b613cec565b905060005b87602001515181101561296957612718886020015182815181106107ec576107ec61555d565b8360200151828151811061272e5761272e61555d565b602090810291909101015280156127ee57602083015161274f600183615842565b8151811061275f5761275f61555d565b602002602001015160001c836020015182815181106127805761278061555d565b602002602001015160001c116127ee576040805162461bcd60e51b8152602060048201526024810191909152600080516020615d1183398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610b06565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106128335761283361555d565b60200260200101518b8b6000015185815181106128525761285261555d565b60200260200101516040518463ffffffff1660e01b815260040161288f9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156128ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128d09190615a6b565b6001600160c01b0316836000015182815181106128ef576128ef61555d565b602002602001018181525050612955610d75612929848660000151858151811061291b5761291b61555d565b602002602001015116613d76565b8a60200151848151811061293f5761293f61555d565b6020026020010151613da190919063ffffffff16565b94508061296181615589565b9150506126f2565b505061297483613e85565b60975490935060ff1660008161298b576000612a0d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129e9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a0d9190615859565b905060005b8a81101561308b578215612b6d578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612a6957612a6961555d565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612aa9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612acd9190615859565b612ad79190615b59565b11612b6d5760405162461bcd60e51b81526020600482015260666024820152600080516020615d1183398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610b06565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612bae57612bae61555d565b9050013560f81c60f81b60f81c8c8c60a001518581518110612bd257612bd261555d565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612c2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c529190615b71565b6001600160401b031916612c758a6040015183815181106107ec576107ec61555d565b67ffffffffffffffff191614612d115760405162461bcd60e51b81526020600482015260616024820152600080516020615d1183398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610b06565b612d4189604001518281518110612d2a57612d2a61555d565b6020026020010151876138e490919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612d8457612d8461555d565b9050013560f81c60f81b60f81c8c8c60c001518581518110612da857612da861555d565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612e04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e289190615902565b85602001518281518110612e3e57612e3e61555d565b6001600160601b03909216602092830291909101820152850151805182908110612e6a57612e6a61555d565b602002602001015185600001518281518110612e8857612e8861555d565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561307657612f0086600001518281518110612ed257612ed261555d565b60200260200101518f8f86818110612eec57612eec61555d565b600192013560f81c9290921c811614919050565b15613064577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612f4657612f4661555d565b9050013560f81c60f81b60f81c8e89602001518581518110612f6a57612f6a61555d565b60200260200101518f60e001518881518110612f8857612f8861555d565b60200260200101518781518110612fa157612fa161555d565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015613005573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130299190615902565b875180518590811061303d5761303d61555d565b602002602001018181516130519190615b9c565b6001600160601b03169052506001909101905b8061306e81615589565b915050612eac565b5050808061308390615589565b915050612a12565b5050506000806130a58c868a606001518b60800151610c5a565b91509150816131165760405162461bcd60e51b81526020600482015260436024820152600080516020615d1183398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610b06565b806131775760405162461bcd60e51b81526020600482015260396024820152600080516020615d1183398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610b06565b505060008782602001516040516020016131929291906155a4565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6131c4613f20565b6131ce6000613f7a565b565b60cb546001600160a01b031633146131e757600080fd5b60006131f960c0850160a086016148c3565b905036600061320b60c0870187615bc4565b90925090506000613223610100880160e089016148c3565b90506000866040516020016132389190615c0a565b6040516020818303038152906040528051906020012090506000806132608387878a8c6122c4565b9150915060005b858110156132f4578460ff16836020015182815181106132895761328961555d565b602002602001015161329b9190615c18565b6001600160601b03166064846000015183815181106132bc576132bc61555d565b60200260200101516001600160601b03166132d79190615c47565b10156132e257600080fd5b806132ec81615589565b915050613267565b5060408051808201825263ffffffff431681526020810183905260cd549151636a3a9c6160e01b815290916001600160a01b031690636a3a9c6190613341908e908e908690600401615c66565b600060405180830381600087803b15801561335b57600080fd5b505af115801561336f573d6000803e3d6000fd5b505050507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a826040516133a4929190615ca4565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106133f4576133f461555d565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906134309088908690600401615b1d565b600060405180830381865afa15801561344d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261347591908101906159bd565b6000815181106134875761348761555d565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156134f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135179190615a6b565b6001600160c01b03169050600061352d82613fcc565b90508161353b8a838a610fd8565b9550955050505050935093915050565b60cc546001600160a01b0316331461356257600080fd5b60cd80546001600160a01b0319166001600160a01b0392909216919091179055565b61358c613f20565b6001600160a01b0381166135f15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610b06565b610b1881613f7a565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561364d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061367191906155ec565b6001600160a01b0316336001600160a01b0316146136a15760405162461bcd60e51b8152600401610b0690615609565b60665419811960665419161461371f5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610b06565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610c4f565b6001600160a01b0381166137e45760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610b06565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613869614384565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561389c5761389e565bfe5b50806138dc5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610b06565b505092915050565b60408051808201909152600080825260208201526139006143a2565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561389c5750806138dc5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610b06565b6139806143c0565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613a68600080516020615cf1833981519152866156b8565b90505b613a7481614098565b9093509150600080516020615cf1833981519152828309831415613aae576040805180820190915290815260208101919091529392505050565b600080516020615cf1833981519152600182089050613a6b565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613afa6143e5565b60005b6002811015613cbf576000613b13826006615c47565b9050848260028110613b2757613b2761555d565b60200201515183613b39836000615b59565b600c8110613b4957613b4961555d565b6020020152848260028110613b6057613b6061555d565b60200201516020015183826001613b779190615b59565b600c8110613b8757613b8761555d565b6020020152838260028110613b9e57613b9e61555d565b6020020151515183613bb1836002615b59565b600c8110613bc157613bc161555d565b6020020152838260028110613bd857613bd861555d565b6020020151516001602002015183613bf1836003615b59565b600c8110613c0157613c0161555d565b6020020152838260028110613c1857613c1861555d565b602002015160200151600060028110613c3357613c3361555d565b602002015183613c44836004615b59565b600c8110613c5457613c5461555d565b6020020152838260028110613c6b57613c6b61555d565b602002015160200151600160028110613c8657613c8661555d565b602002015183613c97836005615b59565b600c8110613ca757613ca761555d565b60200201525080613cb781615589565b915050613afd565b50613cc8614404565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613cf88461411a565b9050808360ff166001901b116122bb5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610b06565b6000805b82156122be57613d8b600184615842565b9092169180613d9981615cce565b915050613d7a565b60408051808201909152600080825260208201526102008261ffff1610613dfd5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610b06565b8161ffff1660011415613e115750816122be565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613e7a57600161ffff871660ff83161c81161415613e5d57613e5a84846138e4565b93505b613e6783846138e4565b92506201fffe600192831b169101613e2d565b509195945050505050565b60408051808201909152600080825260208201528151158015613eaa57506020820151155b15613ec8575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615cf18339815191528460200151613efb91906156b8565b613f1390600080516020615cf1833981519152615842565b905292915050565b919050565b6033546001600160a01b031633146131ce5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610b06565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613fda84613d76565b61ffff166001600160401b03811115613ff557613ff56145a5565b6040519080825280601f01601f19166020018201604052801561401f576020820181803683370190505b5090506000805b825182108015614037575061010081105b1561408e576001811b93508584161561407e578060f81b8383815181106140605761406061555d565b60200101906001600160f81b031916908160001a9053508160010191505b61408781615589565b9050614026565b5090949350505050565b60008080600080516020615cf18339815191526003600080516020615cf183398151915286600080516020615cf183398151915288890909089050600061410e827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615cf18339815191526142a7565b91959194509092505050565b6000610100825111156141a35760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610b06565b81516141b157506000919050565b600080836000815181106141c7576141c761555d565b0160200151600160f89190911c81901b92505b845181101561429e578481815181106141f5576141f561555d565b0160200151600160f89190911c1b915082821161428a5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610b06565b9181179161429781615589565b90506141da565b50909392505050565b6000806142b2614404565b6142ba614422565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561389c5750826143445760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610b06565b505195945050505050565b6040518060a00160405280614362614440565b8152600060208201819052606060408301819052820181905260809091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806143d361445e565b81526020016143e061445e565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b63ffffffff81168114610b1857600080fd5b8035613f1b8161447c565b60008083601f8401126144ab57600080fd5b5081356001600160401b038111156144c257600080fd5b6020830191508360208285010111156144da57600080fd5b9250929050565b8015158114610b1857600080fd5b8035613f1b816144e1565b6000806000806000610100868803121561451357600080fd5b60a086018781111561452457600080fd5b869550356145318161447c565b935060c08601356001600160401b0381111561454c57600080fd5b61455888828901614499565b90945092505060e086013561456c816144e1565b809150509295509295909350565b6000610120828403121561458d57600080fd5b50919050565b60006040828403121561458d57600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156145dd576145dd6145a5565b60405290565b60405161010081016001600160401b03811182821017156145dd576145dd6145a5565b604051601f8201601f191681016001600160401b038111828210171561462e5761462e6145a5565b604052919050565b60006001600160401b0382111561464f5761464f6145a5565b5060051b60200190565b60006040828403121561466b57600080fd5b6146736145bb565b9050813581526020820135602082015292915050565b600082601f83011261469a57600080fd5b813560206146af6146aa83614636565b614606565b82815260069290921b840181019181810190868411156146ce57600080fd5b8286015b848110156146f2576146e48882614659565b8352918301916040016146d2565b509695505050505050565b60008060008060c0858703121561471357600080fd5b84356001600160401b038082111561472a57600080fd5b6147368883890161457a565b95506147458860208901614593565b94506147548860608901614593565b935060a087013591508082111561476a57600080fd5b5061477787828801614689565b91505092959194509250565b6001600160a01b0381168114610b1857600080fd5b6000602082840312156147aa57600080fd5b81356122bb81614783565b6000602082840312156147c757600080fd5b5035919050565b600082601f8301126147df57600080fd5b604051604081018181106001600160401b0382111715614801576148016145a5565b806040525080604084018581111561481857600080fd5b845b81811015613e7a57803583526020928301920161481a565b60006080828403121561484457600080fd5b61484c6145bb565b905061485883836147ce565b815261486783604084016147ce565b602082015292915050565b600080600080610120858703121561488957600080fd5b8435935061489a8660208701614659565b92506148a98660608701614832565b91506148b88660e08701614659565b905092959194509250565b6000602082840312156148d557600080fd5b81356122bb8161447c565b60008083601f8401126148f257600080fd5b5081356001600160401b0381111561490957600080fd5b6020830191508360208260051b85010111156144da57600080fd5b6000806000806000806080878903121561493d57600080fd5b86356001600160401b038082111561495457600080fd5b6149608a838b0161457a565b9750602089013591506149728261447c565b9095506040880135908082111561498857600080fd5b6149948a838b016148e0565b909650945060608901359150808211156149ad57600080fd5b506149ba89828a01614499565b979a9699509497509295939492505050565b600080604083850312156149df57600080fd5b82356149ea81614783565b91506020838101356001600160401b03811115614a0657600080fd5b8401601f81018613614a1757600080fd5b8035614a256146aa82614636565b81815260059190911b82018301908381019088831115614a4457600080fd5b928401925b82841015614a6b578335614a5c81614783565b82529284019290840190614a49565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614aaa57815187529582019590820190600101614a8e565b509495945050505050565b6020815260006122bb6020830184614a7a565b600080600060608486031215614add57600080fd5b8335614ae881614783565b92506020848101356001600160401b0380821115614b0557600080fd5b818701915087601f830112614b1957600080fd5b813581811115614b2b57614b2b6145a5565b614b3d601f8201601f19168501614606565b91508082528884828501011115614b5357600080fd5b8084840185840137600084828401015250809450505050614b766040850161448e565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614c15578385038a52825180518087529087019087870190845b81811015614c0057835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614bbc565b50509a87019a95505091850191600101614b9e565b509298975050505050505050565b6020815260006122bb6020830184614b7f565b600080600060a08486031215614c4b57600080fd5b83356001600160401b03811115614c6157600080fd5b614c6d8682870161457a565b935050614c7d8560208601614593565b9150614b768560608601614593565b600060208284031215614c9e57600080fd5b81356122bb816144e1565b600082601f830112614cba57600080fd5b81356020614cca6146aa83614636565b82815260059290921b84018101918181019086841115614ce957600080fd5b8286015b848110156146f25780358352918301918301614ced565b60008060408385031215614d1757600080fd5b8235614d2281614783565b915060208301356001600160401b03811115614d3d57600080fd5b614d4985828601614ca9565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614d945783516001600160a01b031683529284019291840191600101614d6f565b50909695505050505050565b60008060008060008060808789031215614db957600080fd5b8635614dc481614783565b95506020870135614dd48161447c565b945060408701356001600160401b0380821115614df057600080fd5b614dfc8a838b01614499565b90965094506060890135915080821115614e1557600080fd5b506149ba89828a016148e0565b600081518084526020808501945080840160005b83811015614aaa57815163ffffffff1687529582019590820190600101614e36565b600060208083528351608082850152614e7460a0850182614e22565b905081850151601f1980868403016040870152614e918383614e22565b92506040870151915080868403016060870152614eae8383614e22565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614f055784878303018452614ef3828751614e22565b95880195938801939150600101614ed9565b509998505050505050505050565b600080600080600060608688031215614f2b57600080fd5b8535614f368161447c565b945060208601356001600160401b0380821115614f5257600080fd5b614f5e89838a016148e0565b90965094506040880135915080821115614f7757600080fd5b50614f8488828901614499565b969995985093965092949392505050565b60ff81168114610b1857600080fd5b600060208284031215614fb657600080fd5b81356122bb81614f95565b600080600060608486031215614fd657600080fd5b8335614fe181614783565b925060208401356001600160401b03811115614ffc57600080fd5b61500886828701614ca9565b92505060408401356150198161447c565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614d9457835183529284019291840191600101615040565b6000806040838503121561506f57600080fd5b823561507a8161447c565b946020939093013593505050565b600082601f83011261509957600080fd5b813560206150a96146aa83614636565b82815260059290921b840181019181810190868411156150c857600080fd5b8286015b848110156146f25780356150df8161447c565b83529183019183016150cc565b600082601f8301126150fd57600080fd5b8135602061510d6146aa83614636565b82815260059290921b8401810191818101908684111561512c57600080fd5b8286015b848110156146f25780356001600160401b0381111561514f5760008081fd5b61515d8986838b0101615088565b845250918301918301615130565b6000610180828403121561517e57600080fd5b6151866145e3565b905081356001600160401b038082111561519f57600080fd5b6151ab85838601615088565b835260208401359150808211156151c157600080fd5b6151cd85838601614689565b602084015260408401359150808211156151e657600080fd5b6151f285838601614689565b60408401526152048560608601614832565b60608401526152168560e08601614659565b608084015261012084013591508082111561523057600080fd5b61523c85838601615088565b60a084015261014084013591508082111561525657600080fd5b61526285838601615088565b60c084015261016084013591508082111561527c57600080fd5b50615289848285016150ec565b60e08301525092915050565b6000806000806000608086880312156152ad57600080fd5b8535945060208601356001600160401b03808211156152cb57600080fd5b6152d789838a01614499565b9096509450604088013591506152ec8261447c565b9092506060870135908082111561530257600080fd5b5061530f8882890161516b565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614aaa5781516001600160601b031687529582019590820190600101615330565b6040815260008351604080840152615370608084018261531c565b90506020850151603f1984830301606085015261538d828261531c565b925050508260208301529392505050565b6000806000608084860312156153b357600080fd5b83356001600160401b03808211156153ca57600080fd5b6153d68783880161457a565b94506153e58760208801614593565b935060608601359150808211156153fb57600080fd5b506154088682870161516b565b9150509250925092565b60008060006060848603121561542757600080fd5b833561543281614783565b92506020840135915060408401356150198161447c565b8281526040602082015260006154626040830184614b7f565b949350505050565b6000815180845260005b8181101561549057602081850181015186830182015201615474565b818111156154a2576000602083870101525b50601f01601f19169290920160200192915050565b6020808252825160009190828483015b60058210156154e65782518152918301916001919091019083016154c7565b50505083015163ffffffff1660c0830152604083015161012060e0840181905261551461014085018361546a565b9150606085015161552e61010086018263ffffffff169052565b5060808501518015158583015261408e565b60006020828403121561555257600080fd5b81516122bb8161447c565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561559d5761559d615573565b5060010190565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156155df578151855293820193908201906001016155c3565b5092979650505050505050565b6000602082840312156155fe57600080fd5b81516122bb81614783565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561566557600080fd5b81516122bb816144e1565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000826156d557634e487b7160e01b600052601260045260246000fd5b500690565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061012060a083853760a08401600080825260a08501356157248161447c565b63ffffffff1690915260c08401359036859003601e19018212615745578081fd5b9084019081356001600160401b0381111561575e578182fd5b80360386131561576c578182fd5b8360c088015261578284880182602086016156da565b935050505061579360e0840161448e565b63ffffffff1660e08501526101006157ac8482016144ef565b8015158683015261408e565b81835260006001600160fb1b038311156157d157600080fd5b8260051b8083602087013760009401602001938452509192915050565b6080815260006158016080830189615703565b63ffffffff8816602084015282810360408401526158208187896157b8565b905082810360608401526158358185876156da565b9998505050505050505050565b60008282101561585457615854615573565b500390565b60006020828403121561586b57600080fd5b5051919050565b6000602080838503121561588557600080fd5b82516001600160401b0381111561589b57600080fd5b8301601f810185136158ac57600080fd5b80516158ba6146aa82614636565b81815260059190911b820183019083810190878311156158d957600080fd5b928401925b828410156158f7578351825292840192908401906158de565b979650505050505050565b60006020828403121561591457600080fd5b81516001600160601b03811681146122bb57600080fd5b80356159368161447c565b63ffffffff168252602090810135910152565b60a08152600061595c60a0830186615703565b905061596b602083018561592b565b82356159768161447c565b63ffffffff1660608301526020929092013560809091015292915050565b63ffffffff841681526040602082015260006159b46040830184866157b8565b95945050505050565b600060208083850312156159d057600080fd5b82516001600160401b038111156159e657600080fd5b8301601f810185136159f757600080fd5b8051615a056146aa82614636565b81815260059190911b82018301908381019087831115615a2457600080fd5b928401925b828410156158f7578351615a3c8161447c565b82529284019290840190615a29565b63ffffffff841681526040602082015260006159b46040830184866156da565b600060208284031215615a7d57600080fd5b81516001600160c01b03811681146122bb57600080fd5b600060ff821660ff811415615aab57615aab615573565b60010192915050565b604081526000615ac86040830185876156da565b905063ffffffff83166020830152949350505050565b63ffffffff86168152606060208201526000615afe6060830186886157b8565b8281036040840152615b118185876156da565b98975050505050505050565b63ffffffff831681526040602082015260006154626040830184614a7a565b600060208284031215615b4e57600080fd5b81516122bb81614f95565b60008219821115615b6c57615b6c615573565b500190565b600060208284031215615b8357600080fd5b815167ffffffffffffffff19811681146122bb57600080fd5b60006001600160601b0383811690831681811015615bbc57615bbc615573565b039392505050565b6000808335601e19843603018112615bdb57600080fd5b8301803591506001600160401b03821115615bf557600080fd5b6020019150368190038213156144da57600080fd5b604081016122be828461592b565b60006001600160601b0380831681851681830481118215151615615c3e57615c3e615573565b02949350505050565b6000816000190483118215151615615c6157615c61615573565b500290565b60a081526000615c7960a0830186615703565b9050615c88602083018561592b565b825163ffffffff16606083015260208301516080830152615462565b60808101615cb2828561592b565b825163ffffffff16604083015260208301516060830152611469565b600061ffff80831681811415615ce657615ce6615573565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220dc66ddef94f078488b678ab12b58f79bc5682663eeade83f2a37e4dd90dcf78664736f6c634300080c0033",
}

// ContractSertnTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSertnTaskManagerMetaData.ABI instead.
var ContractSertnTaskManagerABI = ContractSertnTaskManagerMetaData.ABI

// ContractSertnTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSertnTaskManagerMetaData.Bin instead.
var ContractSertnTaskManagerBin = ContractSertnTaskManagerMetaData.Bin

// DeployContractSertnTaskManager deploys a new Ethereum contract, binding an instance of ContractSertnTaskManager to it.
func DeployContractSertnTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address) (common.Address, *types.Transaction, *ContractSertnTaskManager, error) {
	parsed, err := ContractSertnTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSertnTaskManagerBin), backend, _registryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSertnTaskManager{ContractSertnTaskManagerCaller: ContractSertnTaskManagerCaller{contract: contract}, ContractSertnTaskManagerTransactor: ContractSertnTaskManagerTransactor{contract: contract}, ContractSertnTaskManagerFilterer: ContractSertnTaskManagerFilterer{contract: contract}}, nil
}

// ContractSertnTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractSertnTaskManager struct {
	ContractSertnTaskManagerCaller     // Read-only binding to the contract
	ContractSertnTaskManagerTransactor // Write-only binding to the contract
	ContractSertnTaskManagerFilterer   // Log filterer for contract events
}

// ContractSertnTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSertnTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSertnTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSertnTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSertnTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSertnTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSertnTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSertnTaskManagerSession struct {
	Contract     *ContractSertnTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractSertnTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSertnTaskManagerCallerSession struct {
	Contract *ContractSertnTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractSertnTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSertnTaskManagerTransactorSession struct {
	Contract     *ContractSertnTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractSertnTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSertnTaskManagerRaw struct {
	Contract *ContractSertnTaskManager // Generic contract binding to access the raw methods on
}

// ContractSertnTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSertnTaskManagerCallerRaw struct {
	Contract *ContractSertnTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSertnTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSertnTaskManagerTransactorRaw struct {
	Contract *ContractSertnTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSertnTaskManager creates a new instance of ContractSertnTaskManager, bound to a specific deployed contract.
func NewContractSertnTaskManager(address common.Address, backend bind.ContractBackend) (*ContractSertnTaskManager, error) {
	contract, err := bindContractSertnTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManager{ContractSertnTaskManagerCaller: ContractSertnTaskManagerCaller{contract: contract}, ContractSertnTaskManagerTransactor: ContractSertnTaskManagerTransactor{contract: contract}, ContractSertnTaskManagerFilterer: ContractSertnTaskManagerFilterer{contract: contract}}, nil
}

// NewContractSertnTaskManagerCaller creates a new read-only instance of ContractSertnTaskManager, bound to a specific deployed contract.
func NewContractSertnTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractSertnTaskManagerCaller, error) {
	contract, err := bindContractSertnTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerCaller{contract: contract}, nil
}

// NewContractSertnTaskManagerTransactor creates a new write-only instance of ContractSertnTaskManager, bound to a specific deployed contract.
func NewContractSertnTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSertnTaskManagerTransactor, error) {
	contract, err := bindContractSertnTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerTransactor{contract: contract}, nil
}

// NewContractSertnTaskManagerFilterer creates a new log filterer instance of ContractSertnTaskManager, bound to a specific deployed contract.
func NewContractSertnTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSertnTaskManagerFilterer, error) {
	contract, err := bindContractSertnTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerFilterer{contract: contract}, nil
}

// bindContractSertnTaskManager binds a generic wrapper to an already deployed contract.
func bindContractSertnTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSertnTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSertnTaskManager *ContractSertnTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSertnTaskManager.Contract.ContractSertnTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSertnTaskManager *ContractSertnTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ContractSertnTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSertnTaskManager *ContractSertnTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ContractSertnTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSertnTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Aggregator(&_ContractSertnTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Aggregator(&_ContractSertnTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSertnTaskManager.Contract.AllTaskHashes(&_ContractSertnTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSertnTaskManager.Contract.AllTaskHashes(&_ContractSertnTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSertnTaskManager.Contract.AllTaskResponses(&_ContractSertnTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSertnTaskManager.Contract.AllTaskResponses(&_ContractSertnTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.BlsApkRegistry(&_ContractSertnTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.BlsApkRegistry(&_ContractSertnTaskManager.CallOpts)
}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 id, uint256 index) view returns(uint256)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) ChallengeInstances(opts *bind.CallOpts, id uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "challengeInstances", id, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 id, uint256 index) view returns(uint256)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) ChallengeInstances(id uint32, index *big.Int) (*big.Int, error) {
	return _ContractSertnTaskManager.Contract.ChallengeInstances(&_ContractSertnTaskManager.CallOpts, id, index)
}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 id, uint256 index) view returns(uint256)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) ChallengeInstances(id uint32, index *big.Int) (*big.Int, error) {
	return _ContractSertnTaskManager.Contract.ChallengeInstances(&_ContractSertnTaskManager.CallOpts, id, index)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSertnTaskManager.Contract.CheckSignatures(&_ContractSertnTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSertnTaskManager.Contract.CheckSignatures(&_ContractSertnTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Delegation(&_ContractSertnTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Delegation(&_ContractSertnTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Generator() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Generator(&_ContractSertnTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Generator(&_ContractSertnTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractSertnTaskManager.Contract.GetBatchOperatorFromId(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractSertnTaskManager.Contract.GetBatchOperatorFromId(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractSertnTaskManager.Contract.GetBatchOperatorId(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractSertnTaskManager.Contract.GetBatchOperatorId(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSertnTaskManager.Contract.GetCheckSignaturesIndices(&_ContractSertnTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSertnTaskManager.Contract.GetCheckSignaturesIndices(&_ContractSertnTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSertnTaskManager.Contract.GetOperatorState(&_ContractSertnTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSertnTaskManager.Contract.GetOperatorState(&_ContractSertnTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSertnTaskManager.Contract.GetOperatorState0(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSertnTaskManager.Contract.GetOperatorState0(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractSertnTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractSertnTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractSertnTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Owner() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Owner(&_ContractSertnTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.Owner(&_ContractSertnTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractSertnTaskManager.Contract.Paused(&_ContractSertnTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractSertnTaskManager.Contract.Paused(&_ContractSertnTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractSertnTaskManager.Contract.Paused0(&_ContractSertnTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractSertnTaskManager.Contract.Paused0(&_ContractSertnTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.PauserRegistry(&_ContractSertnTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.PauserRegistry(&_ContractSertnTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.RegistryCoordinator(&_ContractSertnTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.RegistryCoordinator(&_ContractSertnTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.StakeRegistry(&_ContractSertnTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractSertnTaskManager.Contract.StakeRegistry(&_ContractSertnTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractSertnTaskManager.Contract.StaleStakesForbidden(&_ContractSertnTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractSertnTaskManager.Contract.StaleStakesForbidden(&_ContractSertnTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractSertnTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSertnTaskManager.Contract.TrySignatureAndApkVerification(&_ContractSertnTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSertnTaskManager *ContractSertnTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSertnTaskManager.Contract.TrySignatureAndApkVerification(&_ContractSertnTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x09b88097.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) ConfirmChallenge(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "confirmChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x09b88097.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) ConfirmChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ConfirmChallenge(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x09b88097.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) ConfirmChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ConfirmChallenge(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x027e4a37.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "createNewTask", inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x027e4a37.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.CreateNewTask(&_ContractSertnTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x027e4a37.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.CreateNewTask(&_ContractSertnTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.Pause(&_ContractSertnTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.Pause(&_ContractSertnTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.PauseAll(&_ContractSertnTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.PauseAll(&_ContractSertnTaskManager.TransactOpts)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) ProveResultAccurate(opts *bind.TransactOpts, taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "proveResultAccurate", taskId, instances, proof)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) ProveResultAccurate(taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ProveResultAccurate(&_ContractSertnTaskManager.TransactOpts, taskId, instances, proof)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) ProveResultAccurate(taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ProveResultAccurate(&_ContractSertnTaskManager.TransactOpts, taskId, instances, proof)
}

// RaiseChallenge is a paid mutator transaction binding the contract method 0x369bcec5.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) RaiseChallenge(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "raiseChallenge", task, taskResponse, taskResponseMetadata)
}

// RaiseChallenge is a paid mutator transaction binding the contract method 0x369bcec5.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RaiseChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RaiseChallenge(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
}

// RaiseChallenge is a paid mutator transaction binding the contract method 0x369bcec5.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) RaiseChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RaiseChallenge(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RenounceOwnership(&_ContractSertnTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RenounceOwnership(&_ContractSertnTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x79e3fb33.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x79e3fb33.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RespondToTask(task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RespondToTask(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x79e3fb33.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) RespondToTask(task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RespondToTask(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x2ced7fa6.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) RespondToTaskWithProof(opts *bind.TransactOpts, task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "respondToTaskWithProof", task, taskIndex, instances, proof)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x2ced7fa6.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RespondToTaskWithProof(task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RespondToTaskWithProof(&_ContractSertnTaskManager.TransactOpts, task, taskIndex, instances, proof)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x2ced7fa6.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) RespondToTaskWithProof(task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RespondToTaskWithProof(&_ContractSertnTaskManager.TransactOpts, task, taskIndex, instances, proof)
}

// SetNewAggregator is a paid mutator transaction binding the contract method 0x44650275.
//
// Solidity: function setNewAggregator(address _aggregator) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) SetNewAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "setNewAggregator", _aggregator)
}

// SetNewAggregator is a paid mutator transaction binding the contract method 0x44650275.
//
// Solidity: function setNewAggregator(address _aggregator) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) SetNewAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetNewAggregator(&_ContractSertnTaskManager.TransactOpts, _aggregator)
}

// SetNewAggregator is a paid mutator transaction binding the contract method 0x44650275.
//
// Solidity: function setNewAggregator(address _aggregator) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) SetNewAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetNewAggregator(&_ContractSertnTaskManager.TransactOpts, _aggregator)
}

// SetNewInferenceDB is a paid mutator transaction binding the contract method 0xf1e1011d.
//
// Solidity: function setNewInferenceDB(address _inferenceDB) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) SetNewInferenceDB(opts *bind.TransactOpts, _inferenceDB common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "setNewInferenceDB", _inferenceDB)
}

// SetNewInferenceDB is a paid mutator transaction binding the contract method 0xf1e1011d.
//
// Solidity: function setNewInferenceDB(address _inferenceDB) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) SetNewInferenceDB(_inferenceDB common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetNewInferenceDB(&_ContractSertnTaskManager.TransactOpts, _inferenceDB)
}

// SetNewInferenceDB is a paid mutator transaction binding the contract method 0xf1e1011d.
//
// Solidity: function setNewInferenceDB(address _inferenceDB) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) SetNewInferenceDB(_inferenceDB common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetNewInferenceDB(&_ContractSertnTaskManager.TransactOpts, _inferenceDB)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetPauserRegistry(&_ContractSertnTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetPauserRegistry(&_ContractSertnTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetStaleStakesForbidden(&_ContractSertnTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.SetStaleStakesForbidden(&_ContractSertnTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.TransferOwnership(&_ContractSertnTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.TransferOwnership(&_ContractSertnTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.Unpause(&_ContractSertnTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.Unpause(&_ContractSertnTaskManager.TransactOpts, newPausedStatus)
}

// ContractSertnTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerInitializedIterator struct {
	Event *ContractSertnTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerInitialized)
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
		it.Event = new(ContractSertnTaskManagerInitialized)
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
func (it *ContractSertnTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerInitialized represents a Initialized event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSertnTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerInitializedIterator{contract: _ContractSertnTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerInitialized)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractSertnTaskManagerInitialized, error) {
	event := new(ContractSertnTaskManagerInitialized)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerNewTaskCreatedIterator struct {
	Event *ContractSertnTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerNewTaskCreated)
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
		it.Event = new(ContractSertnTaskManagerNewTaskCreated)
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
func (it *ContractSertnTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      ITaskStructTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x02fbf40079c9022bc174c39bce245cecd33b518a66e442ae48699be1ab5671ee.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool) task)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractSertnTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerNewTaskCreatedIterator{contract: _ContractSertnTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x02fbf40079c9022bc174c39bce245cecd33b518a66e442ae48699be1ab5671ee.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool) task)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerNewTaskCreated)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractSertnTaskManagerNewTaskCreated, error) {
	event := new(ContractSertnTaskManagerNewTaskCreated)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerOwnershipTransferredIterator struct {
	Event *ContractSertnTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractSertnTaskManagerOwnershipTransferred)
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
func (it *ContractSertnTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSertnTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerOwnershipTransferredIterator{contract: _ContractSertnTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerOwnershipTransferred)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSertnTaskManagerOwnershipTransferred, error) {
	event := new(ContractSertnTaskManagerOwnershipTransferred)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerPausedIterator struct {
	Event *ContractSertnTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerPaused)
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
		it.Event = new(ContractSertnTaskManagerPaused)
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
func (it *ContractSertnTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerPaused represents a Paused event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSertnTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerPausedIterator{contract: _ContractSertnTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerPaused)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParsePaused(log types.Log) (*ContractSertnTaskManagerPaused, error) {
	event := new(ContractSertnTaskManagerPaused)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerPauserRegistrySetIterator struct {
	Event *ContractSertnTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractSertnTaskManagerPauserRegistrySet)
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
func (it *ContractSertnTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractSertnTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerPauserRegistrySetIterator{contract: _ContractSertnTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerPauserRegistrySet)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractSertnTaskManagerPauserRegistrySet, error) {
	event := new(ContractSertnTaskManagerPauserRegistrySet)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractSertnTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractSertnTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractSertnTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractSertnTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractSertnTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractSertnTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractSertnTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerTaskChallengedIterator is returned from FilterTaskChallenged and is used to iterate over the raw logs and unpacked data for TaskChallenged events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskChallengedIterator struct {
	Event *ContractSertnTaskManagerTaskChallenged // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerTaskChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerTaskChallenged)
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
		it.Event = new(ContractSertnTaskManagerTaskChallenged)
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
func (it *ContractSertnTaskManagerTaskChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerTaskChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerTaskChallenged represents a TaskChallenged event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskChallenged struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskChallenged is a free log retrieval operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterTaskChallenged(opts *bind.FilterOpts, taskIndex []uint32) (*ContractSertnTaskManagerTaskChallengedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerTaskChallengedIterator{contract: _ContractSertnTaskManager.contract, event: "TaskChallenged", logs: logs, sub: sub}, nil
}

// WatchTaskChallenged is a free log subscription operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchTaskChallenged(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerTaskChallenged, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerTaskChallenged)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseTaskChallenged(log types.Log) (*ContractSertnTaskManagerTaskChallenged, error) {
	event := new(ContractSertnTaskManagerTaskChallenged)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractSertnTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractSertnTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractSertnTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractSertnTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractSertnTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerTaskChallengedSuccessfully)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractSertnTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractSertnTaskManagerTaskChallengedSuccessfully)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractSertnTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractSertnTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractSertnTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex uint32
	Prover    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed prover)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, prover []common.Address) (*ContractSertnTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractSertnTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed prover)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, prover []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractSertnTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractSertnTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskCompletedIterator struct {
	Event *ContractSertnTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerTaskCompleted)
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
		it.Event = new(ContractSertnTaskManagerTaskCompleted)
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
func (it *ContractSertnTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractSertnTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerTaskCompletedIterator{contract: _ContractSertnTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerTaskCompleted)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractSertnTaskManagerTaskCompleted, error) {
	event := new(ContractSertnTaskManagerTaskCompleted)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskRespondedIterator struct {
	Event *ContractSertnTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerTaskResponded)
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
		it.Event = new(ContractSertnTaskManagerTaskResponded)
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
func (it *ContractSertnTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerTaskResponded represents a TaskResponded event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskResponded struct {
	TaskResponse         ITaskStructTaskResponse
	TaskResponseMetadata ITaskStructTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractSertnTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerTaskRespondedIterator{contract: _ContractSertnTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerTaskResponded)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractSertnTaskManagerTaskResponded, error) {
	event := new(ContractSertnTaskManagerTaskResponded)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerTaskRespondedWithProofIterator is returned from FilterTaskRespondedWithProof and is used to iterate over the raw logs and unpacked data for TaskRespondedWithProof events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskRespondedWithProofIterator struct {
	Event *ContractSertnTaskManagerTaskRespondedWithProof // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerTaskRespondedWithProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerTaskRespondedWithProof)
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
		it.Event = new(ContractSertnTaskManagerTaskRespondedWithProof)
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
func (it *ContractSertnTaskManagerTaskRespondedWithProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerTaskRespondedWithProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerTaskRespondedWithProof represents a TaskRespondedWithProof event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerTaskRespondedWithProof struct {
	TaskIndex uint32
	Output    *big.Int
	Prover    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskRespondedWithProof is a free log retrieval operation binding the contract event 0x4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada.
//
// Solidity: event TaskRespondedWithProof(uint32 indexed taskIndex, uint256 output, address indexed prover)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterTaskRespondedWithProof(opts *bind.FilterOpts, taskIndex []uint32, prover []common.Address) (*ContractSertnTaskManagerTaskRespondedWithProofIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "TaskRespondedWithProof", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerTaskRespondedWithProofIterator{contract: _ContractSertnTaskManager.contract, event: "TaskRespondedWithProof", logs: logs, sub: sub}, nil
}

// WatchTaskRespondedWithProof is a free log subscription operation binding the contract event 0x4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada.
//
// Solidity: event TaskRespondedWithProof(uint32 indexed taskIndex, uint256 output, address indexed prover)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchTaskRespondedWithProof(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerTaskRespondedWithProof, taskIndex []uint32, prover []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "TaskRespondedWithProof", taskIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerTaskRespondedWithProof)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskRespondedWithProof", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseTaskRespondedWithProof(log types.Log) (*ContractSertnTaskManagerTaskRespondedWithProof, error) {
	event := new(ContractSertnTaskManagerTaskRespondedWithProof)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "TaskRespondedWithProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerUnpausedIterator struct {
	Event *ContractSertnTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractSertnTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnTaskManagerUnpaused)
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
		it.Event = new(ContractSertnTaskManagerUnpaused)
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
func (it *ContractSertnTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnTaskManagerUnpaused represents a Unpaused event raised by the ContractSertnTaskManager contract.
type ContractSertnTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSertnTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnTaskManagerUnpausedIterator{contract: _ContractSertnTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSertnTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSertnTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnTaskManagerUnpaused)
				if err := _ContractSertnTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractSertnTaskManager *ContractSertnTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractSertnTaskManagerUnpaused, error) {
	event := new(ContractSertnTaskManagerUnpaused)
	if err := _ContractSertnTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
