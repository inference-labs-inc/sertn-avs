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
	ModelVerifier             common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeInstances\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"provenOnResponce\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveResultAccurate\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"raiseChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTaskWithProof\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNewAggregator\",\"inputs\":[{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNewInferenceDB\",\"inputs\":[{\"name\":\"_inferenceDB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskRespondedWithProof\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162006010380380620060108339810160408190526200003591620001ed565b80806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001ed565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001ed565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b39190620001ed565b6001600160a01b031660e05250506097805460ff1916600117905562000214565b6001600160a01b0381168114620001ea57600080fd5b50565b6000602082840312156200020057600080fd5b81516200020d81620001d4565b9392505050565b60805160a05160c05160e051615d836200028d60003960008181610525015261262a0152600081816103d601528181611a28015261280c015260008181610410015281816129e20152612ba401526000818161043701528181610ece015281816123140152818161248d01526126c70152615d836000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c806367d6be44116101255780639e8e5689116100ad578063df5cf7231161007c578063df5cf72314610520578063ed78be6c14610547578063f1e1011d1461055a578063f2fde38b1461056d578063fabc1cbc1461058057600080fd5b80639e8e5689146104cc578063b953bace146104df578063b98d0908146104f2578063cefdc1d4146104ff57600080fd5b8063715018a6116100f4578063715018a61461047a5780637afa1eed1461048257806384cbfa5c14610495578063886f1195146104a85780638da5cb5b146104bb57600080fd5b806367d6be44146103f8578063683048351461040b5780636d14a987146104325780636efb46361461045957600080fd5b80634d2b57fe116101a8578063595c6a6711610177578063595c6a67146103645780635ac86ab71461036c5780635c1556621461039f5780635c975abb146103bf5780635df45946146103d157600080fd5b80634d2b57fe146102fe5780634f739f741461031e5780635325ada41461033e57806358a7cd261461035157600080fd5b806331b36bd9116101e457806331b36bd9146102985780633563b0d1146102b8578063416c7e5e146102d857806344650275146102eb57600080fd5b806310d67a2f14610216578063136439dd1461022b578063171f1d5b1461023e578063245a7bfc1461026d575b600080fd5b610229610224366004614452565b610593565b005b61022961023936600461446f565b61064f565b61025161024c3660046145ed565b61078e565b6040805192151583529015156020830152015b60405180910390f35b60c954610280906001600160a01b031681565b6040516001600160a01b039091168152602001610264565b6102ab6102a636600461466c565b610918565b604051610264919061475a565b6102cb6102c636600461478a565b610a34565b60405161026491906148e2565b6102296102e636600461490e565b610ecc565b6102296102f9366004614452565b611041565b61031161030c366004614991565b611084565b60405161026491906149e0565b61033161032c366004614ab9565b611199565b6040516102649190614b83565b61022961034c366004614ccd565b6118bf565b61022961035f366004614d53565b611bab565b610229611c4f565b61038f61037a366004614de4565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6103b26103ad366004614e01565b611d16565b6040516102649190614e64565b6066545b604051908152602001610264565b6102807f000000000000000000000000000000000000000000000000000000000000000081565b6103c3610406366004614e9c565b611ede565b6102807f000000000000000000000000000000000000000000000000000000000000000081565b6102807f000000000000000000000000000000000000000000000000000000000000000081565b61046c6104673660046150d5565b611f61565b604051610264929190615195565b610229612e59565b60ca54610280906001600160a01b031681565b6102296104a33660046151de565b612e6d565b606554610280906001600160a01b031681565b6033546001600160a01b0316610280565b6102296104da366004615252565b613056565b6102296104ed3660046152e8565b61312e565b60975461038f9060ff1681565b61051261050d36600461533e565b6131e4565b604051610264929190615375565b6102807f000000000000000000000000000000000000000000000000000000000000000081565b610229610555366004615396565b613376565b610229610568366004614452565b613505565b61022961057b366004614452565b61353e565b61022961058e36600461446f565b6135b4565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060a9190615428565b6001600160a01b0316336001600160a01b0316146106435760405162461bcd60e51b815260040161063a90615445565b60405180910390fd5b61064c81613710565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610697573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106bb919061548f565b6106d75760405162461bcd60e51b815260040161063a906154ac565b606654818116146107505760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161063a565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106107d6576107d66154f4565b60200201518951600160200201518a602001516000600281106107fb576107fb6154f4565b60200201518b60200151600160028110610817576108176154f4565b602090810291909101518c518d8301516040516108749a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610897919061550a565b905061090a6108b06108a98884613807565b869061389e565b6108b8613932565b6109006108f1856108eb604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613807565b6108fa8c6139f2565b9061389e565b886201d4c0613a82565b909890975095505050505050565b606081516001600160401b0381111561093357610933614488565b60405190808252806020026020018201604052801561095c578160200160208202803683370190505b50905060005b8251811015610a2d57836001600160a01b03166313542a4e84838151811061098c5761098c6154f4565b60200260200101516040518263ffffffff1660e01b81526004016109bf91906001600160a01b0391909116815260200190565b602060405180830381865afa1580156109dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a00919061552c565b828281518110610a1257610a126154f4565b6020908102919091010152610a268161555b565b9050610962565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a9a9190615428565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610adc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b009190615428565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b669190615428565b9050600086516001600160401b03811115610b8357610b83614488565b604051908082528060200260200182016040528015610bb657816020015b6060815260200190600190039081610ba15790505b50905060005b8751811015610ebe576000888281518110610bd957610bd96154f4565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c3a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c629190810190615576565b905080516001600160401b03811115610c7d57610c7d614488565b604051908082528060200260200182016040528015610cc857816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610c9b5790505b50848481518110610cdb57610cdb6154f4565b602002602001018190525060005b8151811015610ea8576040518060600160405280876001600160a01b03166347b314e8858581518110610d1e57610d1e6154f4565b60200260200101516040518263ffffffff1660e01b8152600401610d4491815260200190565b602060405180830381865afa158015610d61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d859190615428565b6001600160a01b03168152602001838381518110610da557610da56154f4565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610dd357610dd36154f4565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610e2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e539190615606565b6001600160601b0316815250858581518110610e7157610e716154f4565b60200260200101518281518110610e8a57610e8a6154f4565b60200260200101819052508080610ea09061555b565b915050610ce9565b5050508080610eb69061555b565b915050610bbc565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4e9190615428565b6001600160a01b0316336001600160a01b031614610ffa5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a40161063a565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60c9546001600160a01b0316331461105857600080fd5b60c980546001600160a01b039092166001600160a01b0319928316811790915560ca8054909216179055565b606081516001600160401b0381111561109f5761109f614488565b6040519080825280602002602001820160405280156110c8578160200160208202803683370190505b50905060005b8251811015610a2d57836001600160a01b031663296bb0648483815181106110f8576110f86154f4565b60200260200101516040518263ffffffff1660e01b815260040161111e91815260200190565b602060405180830381865afa15801561113b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115f9190615428565b828281518110611171576111716154f4565b6001600160a01b03909216602092830291909101909101526111928161555b565b90506110ce565b6111c46040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611204573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112289190615428565b90506112556040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611285908b9089908990600401615665565b600060405180830381865afa1580156112a2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112ca919081019061568e565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906112fc908b908b908b90600401615745565b600060405180830381865afa158015611319573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611341919081019061568e565b6040820152856001600160401b0381111561135e5761135e614488565b60405190808252806020026020018201604052801561139157816020015b606081526020019060019003908161137c5790505b50606082015260005b60ff81168711156117d0576000856001600160401b038111156113bf576113bf614488565b6040519080825280602002602001820160405280156113e8578160200160208202803683370190505b5083606001518360ff1681518110611402576114026154f4565b602002602001018190525060005b868110156116d05760008c6001600160a01b03166304ec63518a8a8581811061143b5761143b6154f4565b905060200201358e88600001518681518110611459576114596154f4565b60200260200101516040518463ffffffff1660e01b81526004016114969392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156114b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114d79190615765565b90506001600160c01b03811661157b5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a40161063a565b8a8a8560ff16818110611590576115906154f4565b6001600160c01b03841692013560f81c9190911c6001908116141590506116bd57856001600160a01b031663dd9846b98a8a858181106115d2576115d26154f4565b905060200201358d8d8860ff168181106115ee576115ee6154f4565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611644573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611668919061578e565b85606001518560ff1681518110611681576116816154f4565b6020026020010151848151811061169a5761169a6154f4565b63ffffffff90921660209283029190910190910152826116b98161555b565b9350505b50806116c88161555b565b915050611410565b506000816001600160401b038111156116eb576116eb614488565b604051908082528060200260200182016040528015611714578160200160208202803683370190505b50905060005b828110156117955784606001518460ff168151811061173b5761173b6154f4565b60200260200101518181518110611754576117546154f4565b602002602001015182828151811061176e5761176e6154f4565b63ffffffff909216602092830291909101909101528061178d8161555b565b91505061171a565b508084606001518460ff16815181106117b0576117b06154f4565b6020026020010181905250505080806117c8906157ab565b91505061139a565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611811573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118359190615428565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611868908b908b908e906004016157cb565b600060405180830381865afa158015611885573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118ad919081019061568e565b60208301525098975050505050505050565b60006118ce60208501856157f5565b9050600082516001600160401b038111156118eb576118eb614488565b604051908082528060200260200182016040528015611914578160200160208202803683370190505b50905060005b835181101561198657611957848281518110611938576119386154f4565b6020026020010151805160009081526020918201519091526040902090565b828281518110611969576119696154f4565b60209081029190910101528061197e8161555b565b91505061191a565b50600061199960c0880160a089016157f5565b826040516020016119ab929190615812565b604051602081830303815290604052805190602001209050846020013581146119d357600080fd5b600084516001600160401b038111156119ee576119ee614488565b604051908082528060200260200182016040528015611a17578160200160208202803683370190505b50905060005b8551811015611b0a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110611a6757611a676154f4565b60200260200101516040518263ffffffff1660e01b8152600401611a8d91815260200190565b602060405180830381865afa158015611aaa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ace9190615428565b828281518110611ae057611ae06154f4565b6001600160a01b039092166020928302919091019091015280611b028161555b565b915050611a1d565b5060cb5460405163959b127960e01b815263ffffffff861660048201526001600160a01b039091169063959b127990602401600060405180830381600087803b158015611b5657600080fd5b505af1158015611b6a573d6000803e3d6000fd5b505060405133925063ffffffff871691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050565b60cb54604051632c53e69360e11b81526001600160a01b03909116906358a7cd2690611be3908890889088908890889060040161585a565b600060405180830381600087803b158015611bfd57600080fd5b505af1158015611c11573d6000803e3d6000fd5b505060405133925063ffffffff881691507ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611c97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cbb919061548f565b611cd75760405162461bcd60e51b815260040161063a906154ac565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611d48929190615899565b600060405180830381865afa158015611d65573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d8d919081019061568e565b9050600084516001600160401b03811115611daa57611daa614488565b604051908082528060200260200182016040528015611dd3578160200160208202803683370190505b50905060005b8551811015611ed457866001600160a01b03166304ec6351878381518110611e0357611e036154f4565b602002602001015187868581518110611e1e57611e1e6154f4565b60200260200101516040518463ffffffff1660e01b8152600401611e5b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611e78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e9c9190615765565b6001600160c01b0316828281518110611eb757611eb76154f4565b602090810291909101015280611ecc8161555b565b915050611dd9565b5095945050505050565b60cb546040516319f5af9160e21b815263ffffffff84166004820152602481018390526000916001600160a01b0316906367d6be4490604401602060405180830381865afa158015611f34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f58919061552c565b90505b92915050565b6040805180820190915260608082526020820152600084611fd85760405162461bcd60e51b81526020600482015260376024820152600080516020615d2e83398151915260448201527f7265733a20656d7074792071756f72756d20696e707574000000000000000000606482015260840161063a565b60408301515185148015611ff0575060a08301515185145b8015612000575060c08301515185145b8015612010575060e08301515185145b61207a5760405162461bcd60e51b81526020600482015260416024820152600080516020615d2e83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a40161063a565b825151602084015151146120f25760405162461bcd60e51b815260206004820152604460248201819052600080516020615d2e833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a40161063a565b4363ffffffff168463ffffffff16106121615760405162461bcd60e51b815260206004820152603c6024820152600080516020615d2e83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b00000000606482015260840161063a565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156121a2576121a2614488565b6040519080825280602002602001820160405280156121cb578160200160208202803683370190505b506020820152866001600160401b038111156121e9576121e9614488565b604051908082528060200260200182016040528015612212578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561224657612246614488565b60405190808252806020026020018201604052801561226f578160200160208202803683370190505b5081526020860151516001600160401b0381111561228f5761228f614488565b6040519080825280602002602001820160405280156122b8578160200160208202803683370190505b508160200181905250600061238a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612361573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061238591906158b8565b613ca6565b905060005b876020015151811015612606576123b588602001518281518110611938576119386154f4565b836020015182815181106123cb576123cb6154f4565b6020908102919091010152801561248b5760208301516123ec6001836158d5565b815181106123fc576123fc6154f4565b602002602001015160001c8360200151828151811061241d5761241d6154f4565b602002602001015160001c1161248b576040805162461bcd60e51b8152602060048201526024810191909152600080516020615d2e83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f72746564606482015260840161063a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106124d0576124d06154f4565b60200260200101518b8b6000015185815181106124ef576124ef6154f4565b60200260200101516040518463ffffffff1660e01b815260040161252c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612549573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061256d9190615765565b6001600160c01b03168360000151828151811061258c5761258c6154f4565b6020026020010181815250506125f26108a96125c684866000015185815181106125b8576125b86154f4565b602002602001015116613d30565b8a6020015184815181106125dc576125dc6154f4565b6020026020010151613d5b90919063ffffffff16565b9450806125fe8161555b565b91505061238f565b505061261183613e3f565b60975490935060ff166000816126285760006126aa565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612686573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126aa919061552c565b905060005b8a811015612d2857821561280a578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612706576127066154f4565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612746573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061276a919061552c565b61277491906158ec565b1161280a5760405162461bcd60e51b81526020600482015260666024820152600080516020615d2e83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c40161063a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061284b5761284b6154f4565b9050013560f81c60f81b60f81c8c8c60a00151858151811061286f5761286f6154f4565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156128cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ef9190615904565b6001600160401b0319166129128a604001518381518110611938576119386154f4565b67ffffffffffffffff1916146129ae5760405162461bcd60e51b81526020600482015260616024820152600080516020615d2e83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c40161063a565b6129de896040015182815181106129c7576129c76154f4565b60200260200101518761389e90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612a2157612a216154f4565b9050013560f81c60f81b60f81c8c8c60c001518581518110612a4557612a456154f4565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612aa1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ac59190615606565b85602001518281518110612adb57612adb6154f4565b6001600160601b03909216602092830291909101820152850151805182908110612b0757612b076154f4565b602002602001015185600001518281518110612b2557612b256154f4565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612d1357612b9d86600001518281518110612b6f57612b6f6154f4565b60200260200101518f8f86818110612b8957612b896154f4565b600192013560f81c9290921c811614919050565b15612d01577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612be357612be36154f4565b9050013560f81c60f81b60f81c8e89602001518581518110612c0757612c076154f4565b60200260200101518f60e001518881518110612c2557612c256154f4565b60200260200101518781518110612c3e57612c3e6154f4565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612ca2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cc69190615606565b8751805185908110612cda57612cda6154f4565b60200260200101818151612cee919061592f565b6001600160601b03169052506001909101905b80612d0b8161555b565b915050612b49565b50508080612d209061555b565b9150506126af565b505050600080612d428c868a606001518b6080015161078e565b9150915081612db35760405162461bcd60e51b81526020600482015260436024820152600080516020615d2e83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a40161063a565b80612e145760405162461bcd60e51b81526020600482015260396024820152600080516020615d2e83398151915260448201527f7265733a207369676e617475726520697320696e76616c696400000000000000606482015260840161063a565b50506000878260200151604051602001612e2f929190615812565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612e61613eda565b612e6b6000613f34565b565b60c9546001600160a01b03163314612e8457600080fd5b6000612e9660c0850160a086016157f5565b9050366000612ea860c0870187615957565b90925090506000612ec0610100880160e089016157f5565b9050600086604051602001612ed591906159bb565b604051602081830303815290604052805190602001209050600080612efd8387878a8c611f61565b9150915060005b85811015612f91578460ff1683602001518281518110612f2657612f266154f4565b6020026020010151612f3891906159c9565b6001600160601b0316606484600001518381518110612f5957612f596154f4565b60200260200101516001600160601b0316612f7491906159f8565b1015612f7f57600080fd5b80612f898161555b565b915050612f04565b5060408051808201825263ffffffff431681526020810183905260cb549151634cba41a760e01b815290916001600160a01b031690634cba41a790612fde908e908e908690600401615ae7565b600060405180830381600087803b158015612ff857600080fd5b505af115801561300c573d6000803e3d6000fd5b505050507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a82604051613041929190615b25565b60405180910390a15050505050505050505050565b60cb54604051639e8e568960e01b81526001600160a01b0390911690639e8e56899061309090899089908990899089908990600401615b4f565b600060405180830381600087803b1580156130aa57600080fd5b505af11580156130be573d6000803e3d6000fd5b50339250505063ffffffff86167f4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada86866130f96001826158d5565b818110613108576131086154f4565b9050602002013560405161311e91815260200190565b60405180910390a3505050505050565b60cb54604051630c0db69560e31b81526001600160a01b039091169063606db4a89061316290869086908690600401615ba3565b600060405180830381600087803b15801561317c57600080fd5b505af1158015613190573d6000803e3d6000fd5b506131a29250505060208301836157f5565b63ffffffff167f07e263801d87e960a4d426543eb7bf2bc36cc793f17327f5cc37cd9aabae3a33846040516131d79190615bee565b60405180910390a2505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061321f5761321f6154f4565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061325b9088908690600401615899565b600060405180830381865afa158015613278573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526132a0919081019061568e565b6000815181106132b2576132b26154f4565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561331e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133429190615765565b6001600160c01b03169050600061335882613f86565b9050816133668a838a610a34565b9550955050505050935093915050565b60ca546001600160a01b0316331461338d57600080fd5b613395614309565b6040805160a081810190925290889060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908716606083015260408051601f8701839004830281018301909152858152908690869081908401838280828437600092018290525060408087019590955287151560808701526001600160a01b0380881660a088015260cb5495519195169350633b7d2cee925061344491508590602001615c4e565b604051602081830303815290604052805190602001206040518263ffffffff1660e01b815260040161347891815260200190565b6020604051808303816000875af1158015613497573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134bb919061578e565b90508063ffffffff167f8fdbcbca31148444d0448386d6e130db10f657ce3f4350386df06f78e71d3ec0836040516134f39190615c4e565b60405180910390a25050505050505050565b60ca546001600160a01b0316331461351c57600080fd5b60cb80546001600160a01b0319166001600160a01b0392909216919091179055565b613546613eda565b6001600160a01b0381166135ab5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161063a565b61064c81613f34565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061362b9190615428565b6001600160a01b0316336001600160a01b03161461365b5760405162461bcd60e51b815260040161063a90615445565b6066541981196066541916146136d95760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161063a565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610783565b6001600160a01b03811661379e5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161063a565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613823614345565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561385657613858565bfe5b50806138965760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640161063a565b505092915050565b60408051808201909152600080825260208201526138ba614363565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156138565750806138965760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161063a565b61393a614381565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613a22600080516020615d0e8339815191528661550a565b90505b613a2e81614052565b9093509150600080516020615d0e833981519152828309831415613a68576040805180820190915290815260208101919091529392505050565b600080516020615d0e833981519152600182089050613a25565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613ab46143a6565b60005b6002811015613c79576000613acd8260066159f8565b9050848260028110613ae157613ae16154f4565b60200201515183613af38360006158ec565b600c8110613b0357613b036154f4565b6020020152848260028110613b1a57613b1a6154f4565b60200201516020015183826001613b3191906158ec565b600c8110613b4157613b416154f4565b6020020152838260028110613b5857613b586154f4565b6020020151515183613b6b8360026158ec565b600c8110613b7b57613b7b6154f4565b6020020152838260028110613b9257613b926154f4565b6020020151516001602002015183613bab8360036158ec565b600c8110613bbb57613bbb6154f4565b6020020152838260028110613bd257613bd26154f4565b602002015160200151600060028110613bed57613bed6154f4565b602002015183613bfe8360046158ec565b600c8110613c0e57613c0e6154f4565b6020020152838260028110613c2557613c256154f4565b602002015160200151600160028110613c4057613c406154f4565b602002015183613c518360056158ec565b600c8110613c6157613c616154f4565b60200201525080613c718161555b565b915050613ab7565b50613c826143c5565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613cb2846140d4565b9050808360ff166001901b11611f585760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c756500606482015260840161063a565b6000805b8215611f5b57613d456001846158d5565b9092169180613d5381615ceb565b915050613d34565b60408051808201909152600080825260208201526102008261ffff1610613db75760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b604482015260640161063a565b8161ffff1660011415613dcb575081611f5b565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613e3457600161ffff871660ff83161c81161415613e1757613e14848461389e565b93505b613e21838461389e565b92506201fffe600192831b169101613de7565b509195945050505050565b60408051808201909152600080825260208201528151158015613e6457506020820151155b15613e82575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615d0e8339815191528460200151613eb5919061550a565b613ecd90600080516020615d0e8339815191526158d5565b905292915050565b919050565b6033546001600160a01b03163314612e6b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161063a565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613f9484613d30565b61ffff166001600160401b03811115613faf57613faf614488565b6040519080825280601f01601f191660200182016040528015613fd9576020820181803683370190505b5090506000805b825182108015613ff1575061010081105b15614048576001811b935085841615614038578060f81b83838151811061401a5761401a6154f4565b60200101906001600160f81b031916908160001a9053508160010191505b6140418161555b565b9050613fe0565b5090949350505050565b60008080600080516020615d0e8339815191526003600080516020615d0e83398151915286600080516020615d0e8339815191528889090908905060006140c8827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615d0e833981519152614261565b91959194509092505050565b60006101008251111561415d5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a40161063a565b815161416b57506000919050565b60008083600081518110614181576141816154f4565b0160200151600160f89190911c81901b92505b8451811015614258578481815181106141af576141af6154f4565b0160200151600160f89190911c1b91508282116142445760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a40161063a565b918117916142518161555b565b9050614194565b50909392505050565b60008061426c6143c5565b6142746143e3565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156138565750826142fe5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640161063a565b505195945050505050565b6040518060c0016040528061431c614401565b815260006020820181905260606040830181905282018190526080820181905260a09091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061439461441f565b81526020016143a161441f565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461064c57600080fd5b60006020828403121561446457600080fd5b8135611f588161443d565b60006020828403121561448157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156144c0576144c0614488565b60405290565b60405161010081016001600160401b03811182821017156144c0576144c0614488565b604051601f8201601f191681016001600160401b038111828210171561451157614511614488565b604052919050565b60006040828403121561452b57600080fd5b61453361449e565b9050813581526020820135602082015292915050565b600082601f83011261455a57600080fd5b604051604081018181106001600160401b038211171561457c5761457c614488565b806040525080604084018581111561459357600080fd5b845b81811015613e34578035835260209283019201614595565b6000608082840312156145bf57600080fd5b6145c761449e565b90506145d38383614549565b81526145e28360408401614549565b602082015292915050565b600080600080610120858703121561460457600080fd5b843593506146158660208701614519565b925061462486606087016145ad565b91506146338660e08701614519565b905092959194509250565b60006001600160401b0382111561465757614657614488565b5060051b60200190565b8035613ed58161443d565b6000806040838503121561467f57600080fd5b823561468a8161443d565b91506020838101356001600160401b038111156146a657600080fd5b8401601f810186136146b757600080fd5b80356146ca6146c58261463e565b6144e9565b81815260059190911b820183019083810190888311156146e957600080fd5b928401925b828410156147105783356147018161443d565b825292840192908401906146ee565b80955050505050509250929050565b600081518084526020808501945080840160005b8381101561474f57815187529582019590820190600101614733565b509495945050505050565b602081526000611f58602083018461471f565b63ffffffff8116811461064c57600080fd5b8035613ed58161476d565b60008060006060848603121561479f57600080fd5b83356147aa8161443d565b92506020848101356001600160401b03808211156147c757600080fd5b818701915087601f8301126147db57600080fd5b8135818111156147ed576147ed614488565b6147ff601f8201601f191685016144e9565b9150808252888482850101111561481557600080fd5b80848401858401376000848284010152508094505050506148386040850161477f565b90509250925092565b6000815180845260208085019450848260051b86018286016000805b868110156148d4578484038a52825180518086529087019087860190845b818110156148bf57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b0316908401529289019260609092019160010161487b565b50509a87019a9450509185019160010161485d565b509198975050505050505050565b602081526000611f586020830184614841565b801515811461064c57600080fd5b8035613ed5816148f5565b60006020828403121561492057600080fd5b8135611f58816148f5565b600082601f83011261493c57600080fd5b8135602061494c6146c58361463e565b82815260059290921b8401810191818101908684111561496b57600080fd5b8286015b84811015614986578035835291830191830161496f565b509695505050505050565b600080604083850312156149a457600080fd5b82356149af8161443d565b915060208301356001600160401b038111156149ca57600080fd5b6149d68582860161492b565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614a215783516001600160a01b0316835292840192918401916001016149fc565b50909695505050505050565b60008083601f840112614a3f57600080fd5b5081356001600160401b03811115614a5657600080fd5b602083019150836020828501011115614a6e57600080fd5b9250929050565b60008083601f840112614a8757600080fd5b5081356001600160401b03811115614a9e57600080fd5b6020830191508360208260051b8501011115614a6e57600080fd5b60008060008060008060808789031215614ad257600080fd5b8635614add8161443d565b95506020870135614aed8161476d565b945060408701356001600160401b0380821115614b0957600080fd5b614b158a838b01614a2d565b90965094506060890135915080821115614b2e57600080fd5b50614b3b89828a01614a75565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561474f57815163ffffffff1687529582019590820190600101614b61565b600060208083528351608082850152614b9f60a0850182614b4d565b905081850151601f1980868403016040870152614bbc8383614b4d565b92506040870151915080868403016060870152614bd98383614b4d565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614c305784878303018452614c1e828751614b4d565b95880195938801939150600101614c04565b509998505050505050505050565b60006101408284031215614c5157600080fd5b50919050565b600060408284031215614c5157600080fd5b600082601f830112614c7a57600080fd5b81356020614c8a6146c58361463e565b82815260069290921b84018101918181019086841115614ca957600080fd5b8286015b8481101561498657614cbf8882614519565b835291830191604001614cad565b60008060008060c08587031215614ce357600080fd5b84356001600160401b0380821115614cfa57600080fd5b614d0688838901614c3e565b9550614d158860208901614c57565b9450614d248860608901614c57565b935060a0870135915080821115614d3a57600080fd5b50614d4787828801614c69565b91505092959194509250565b600080600080600060608688031215614d6b57600080fd5b8535614d768161476d565b945060208601356001600160401b0380821115614d9257600080fd5b614d9e89838a01614a75565b90965094506040880135915080821115614db757600080fd5b50614dc488828901614a2d565b969995985093965092949392505050565b60ff8116811461064c57600080fd5b600060208284031215614df657600080fd5b8135611f5881614dd5565b600080600060608486031215614e1657600080fd5b8335614e218161443d565b925060208401356001600160401b03811115614e3c57600080fd5b614e488682870161492b565b9250506040840135614e598161476d565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614a2157835183529284019291840191600101614e80565b60008060408385031215614eaf57600080fd5b8235614eba8161476d565b946020939093013593505050565b600082601f830112614ed957600080fd5b81356020614ee96146c58361463e565b82815260059290921b84018101918181019086841115614f0857600080fd5b8286015b84811015614986578035614f1f8161476d565b8352918301918301614f0c565b600082601f830112614f3d57600080fd5b81356020614f4d6146c58361463e565b82815260059290921b84018101918181019086841115614f6c57600080fd5b8286015b848110156149865780356001600160401b03811115614f8f5760008081fd5b614f9d8986838b0101614ec8565b845250918301918301614f70565b60006101808284031215614fbe57600080fd5b614fc66144c6565b905081356001600160401b0380821115614fdf57600080fd5b614feb85838601614ec8565b8352602084013591508082111561500157600080fd5b61500d85838601614c69565b6020840152604084013591508082111561502657600080fd5b61503285838601614c69565b604084015261504485606086016145ad565b60608401526150568560e08601614519565b608084015261012084013591508082111561507057600080fd5b61507c85838601614ec8565b60a084015261014084013591508082111561509657600080fd5b6150a285838601614ec8565b60c08401526101608401359150808211156150bc57600080fd5b506150c984828501614f2c565b60e08301525092915050565b6000806000806000608086880312156150ed57600080fd5b8535945060208601356001600160401b038082111561510b57600080fd5b61511789838a01614a2d565b90965094506040880135915061512c8261476d565b9092506060870135908082111561514257600080fd5b5061514f88828901614fab565b9150509295509295909350565b600081518084526020808501945080840160005b8381101561474f5781516001600160601b031687529582019590820190600101615170565b60408152600083516040808401526151b0608084018261515c565b90506020850151603f198483030160608501526151cd828261515c565b925050508260208301529392505050565b6000806000608084860312156151f357600080fd5b83356001600160401b038082111561520a57600080fd5b61521687838801614c3e565b94506152258760208801614c57565b9350606086013591508082111561523b57600080fd5b5061524886828701614fab565b9150509250925092565b6000806000806000806080878903121561526b57600080fd5b86356001600160401b038082111561528257600080fd5b61528e8a838b01614c3e565b9750602089013591506152a08261476d565b909550604088013590808211156152b657600080fd5b6152c28a838b01614a75565b909650945060608901359150808211156152db57600080fd5b50614b3b89828a01614a2d565b600080600060a084860312156152fd57600080fd5b83356001600160401b0381111561531357600080fd5b61531f86828701614c3e565b93505061532f8560208601614c57565b91506148388560608601614c57565b60008060006060848603121561535357600080fd5b833561535e8161443d565b9250602084013591506040840135614e598161476d565b82815260406020820152600061538e6040830184614841565b949350505050565b60008060008060008061012087890312156153b057600080fd5b60a08701888111156153c157600080fd5b879650356153ce8161476d565b945060c08701356001600160401b038111156153e957600080fd5b6153f589828a01614a2d565b90955093505060e0870135615409816148f5565b915061010087013561541a8161443d565b809150509295509295509295565b60006020828403121561543a57600080fd5b8151611f588161443d565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156154a157600080fd5b8151611f58816148f5565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261552757634e487b7160e01b600052601260045260246000fd5b500690565b60006020828403121561553e57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600060001982141561556f5761556f615545565b5060010190565b6000602080838503121561558957600080fd5b82516001600160401b0381111561559f57600080fd5b8301601f810185136155b057600080fd5b80516155be6146c58261463e565b81815260059190911b820183019083810190878311156155dd57600080fd5b928401925b828410156155fb578351825292840192908401906155e2565b979650505050505050565b60006020828403121561561857600080fd5b81516001600160601b0381168114611f5857600080fd5b81835260006001600160fb1b0383111561564857600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff8416815260406020820152600061568560408301848661562f565b95945050505050565b600060208083850312156156a157600080fd5b82516001600160401b038111156156b757600080fd5b8301601f810185136156c857600080fd5b80516156d66146c58261463e565b81815260059190911b820183019083810190878311156156f557600080fd5b928401925b828410156155fb57835161570d8161476d565b825292840192908401906156fa565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff8416815260406020820152600061568560408301848661571c565b60006020828403121561577757600080fd5b81516001600160c01b0381168114611f5857600080fd5b6000602082840312156157a057600080fd5b8151611f588161476d565b600060ff821660ff8114156157c2576157c2615545565b60010192915050565b6040815260006157df60408301858761571c565b905063ffffffff83166020830152949350505050565b60006020828403121561580757600080fd5b8135611f588161476d565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561584d57815185529382019390820190600101615831565b5092979650505050505050565b63ffffffff8616815260606020820152600061587a60608301868861562f565b828103604084015261588d81858761571c565b98975050505050505050565b63ffffffff8316815260406020820152600061538e604083018461471f565b6000602082840312156158ca57600080fd5b8151611f5881614dd5565b6000828210156158e7576158e7615545565b500390565b600082198211156158ff576158ff615545565b500190565b60006020828403121561591657600080fd5b815167ffffffffffffffff1981168114611f5857600080fd5b60006001600160601b038381169083168181101561594f5761594f615545565b039392505050565b6000808335601e1984360301811261596e57600080fd5b8301803591506001600160401b0382111561598857600080fd5b602001915036819003821315614a6e57600080fd5b80356159a88161476d565b63ffffffff168252602090810135910152565b60408101611f5b828461599d565b60006001600160601b03808316818516818304811182151516156159ef576159ef615545565b02949350505050565b6000816000190483118215151615615a1257615a12615545565b500290565b600061014060a083853760a08401600080825260a0850135615a388161476d565b63ffffffff1690915260c08401359036859003601e19018212615a59578081fd5b9084019081356001600160401b03811115615a72578182fd5b803603861315615a80578182fd5b8360c0880152615a96848801826020860161571c565b9350505050615aa760e0840161477f565b63ffffffff1660e0850152610100615ac0848201614903565b151590850152610120615ad4848201614661565b6001600160a01b03811686830152614048565b60a081526000615afa60a0830186615a17565b9050615b09602083018561599d565b825163ffffffff1660608301526020830151608083015261538e565b60808101615b33828561599d565b825163ffffffff16604083015260208301516060830152610ec5565b608081526000615b626080830189615a17565b63ffffffff881660208401528281036040840152615b8181878961562f565b90508281036060840152615b9681858761571c565b9998505050505050505050565b60a081526000615bb660a0830186615a17565b9050615bc5602083018561599d565b8235615bd08161476d565b63ffffffff1660608301526020929092013560809091015292915050565b602081526000611f586020830184615a17565b6000815180845260005b81811015615c2757602081850181015186830182015201615c0b565b81811115615c39576000602083870101525b50601f01601f19169290920160200192915050565b6020808252825160009190828483015b6005821015615c7d578251815291830191600191909101908301615c5e565b50505083015163ffffffff1660c0830152604083015161014060e08401819052615cab610160850183615c01565b91506060850151615cc561010086018263ffffffff169052565b506080850151151561012085015260a08501516001600160a01b03811682860152614048565b600061ffff80831681811415615d0357615d03615545565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212201135aae2a2a48db15dd692ca59bd4daf44906894a0904a68855571090e16db2864736f6c634300080c0033",
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

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x5325ada4.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) ConfirmChallenge(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "confirmChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x5325ada4.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) ConfirmChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ConfirmChallenge(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x5325ada4.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) ConfirmChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.ConfirmChallenge(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xed78be6c.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce, address modelVerifier) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool, modelVerifier common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "createNewTask", inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce, modelVerifier)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xed78be6c.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce, address modelVerifier) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool, modelVerifier common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.CreateNewTask(&_ContractSertnTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce, modelVerifier)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xed78be6c.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers, bool provenOnResponce, address modelVerifier) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte, provenOnResponce bool, modelVerifier common.Address) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.CreateNewTask(&_ContractSertnTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers, provenOnResponce, modelVerifier)
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

// RaiseChallenge is a paid mutator transaction binding the contract method 0xb953bace.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) RaiseChallenge(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "raiseChallenge", task, taskResponse, taskResponseMetadata)
}

// RaiseChallenge is a paid mutator transaction binding the contract method 0xb953bace.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RaiseChallenge(task ITaskStructTask, taskResponse ITaskStructTaskResponse, taskResponseMetadata ITaskStructTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RaiseChallenge(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
}

// RaiseChallenge is a paid mutator transaction binding the contract method 0xb953bace.
//
// Solidity: function raiseChallenge((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
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

// RespondToTask is a paid mutator transaction binding the contract method 0x84cbfa5c.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x84cbfa5c.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RespondToTask(task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RespondToTask(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x84cbfa5c.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32,bool,address) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactorSession) RespondToTask(task ITaskStructTask, taskResponse ITaskStructTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RespondToTask(&_ContractSertnTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x9e8e5689.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool,address) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerTransactor) RespondToTaskWithProof(opts *bind.TransactOpts, task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.contract.Transact(opts, "respondToTaskWithProof", task, taskIndex, instances, proof)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x9e8e5689.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool,address) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
func (_ContractSertnTaskManager *ContractSertnTaskManagerSession) RespondToTaskWithProof(task ITaskStructTask, taskIndex uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractSertnTaskManager.Contract.RespondToTaskWithProof(&_ContractSertnTaskManager.TransactOpts, task, taskIndex, instances, proof)
}

// RespondToTaskWithProof is a paid mutator transaction binding the contract method 0x9e8e5689.
//
// Solidity: function respondToTaskWithProof((uint256[5],uint32,bytes,uint32,bool,address) task, uint32 taskIndex, uint256[] instances, bytes proof) returns()
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

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x8fdbcbca31148444d0448386d6e130db10f657ce3f4350386df06f78e71d3ec0.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool,address) task)
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

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x8fdbcbca31148444d0448386d6e130db10f657ce3f4350386df06f78e71d3ec0.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool,address) task)
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x8fdbcbca31148444d0448386d6e130db10f657ce3f4350386df06f78e71d3ec0.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool,address) task)
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
	Task      ITaskStructTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskChallenged is a free log retrieval operation binding the contract event 0x07e263801d87e960a4d426543eb7bf2bc36cc793f17327f5cc37cd9aabae3a33.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool,address) task)
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

// WatchTaskChallenged is a free log subscription operation binding the contract event 0x07e263801d87e960a4d426543eb7bf2bc36cc793f17327f5cc37cd9aabae3a33.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool,address) task)
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

// ParseTaskChallenged is a log parse operation binding the contract event 0x07e263801d87e960a4d426543eb7bf2bc36cc793f17327f5cc37cd9aabae3a33.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32,bool,address) task)
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
