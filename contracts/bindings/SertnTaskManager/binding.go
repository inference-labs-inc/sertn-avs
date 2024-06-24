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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeInstances\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"provenOnResponce\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveResultAccurate\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"raiseChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTaskWithProof\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"taskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNewAggregator\",\"inputs\":[{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNewInferenceDB\",\"inputs\":[{\"name\":\"_inferenceDB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"provenOnResponse\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modelVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structITaskStruct.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskRespondedWithProof\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200604f3803806200604f8339810160408190526200003591620001ed565b80806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b59190620001ed565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620001ed565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b39190620001ed565b6001600160a01b031660e05250506097805460ff1916600117905562000214565b6001600160a01b0381168114620001ea57600080fd5b50565b6000602082840312156200020057600080fd5b81516200020d81620001d4565b9392505050565b60805160a05160c05160e051615dc26200028d6000396000818161057f015261268401526000818161043001528181611a82015261286601526000818161046a01528181612a3c0152612bfe01526000818161049101528181610f280152818161236e015281816124e701526127210152615dc26000f3fe608060405234801561001057600080fd5b50600436106102275760003560e01c80635df45946116101305780638da5cb5b116100b8578063df5cf7231161007c578063df5cf7231461057a578063ed78be6c146105a1578063f1e1011d146105b4578063f2fde38b146105c7578063fabc1cbc146105da57600080fd5b80638da5cb5b146105155780639e8e568914610526578063b953bace14610539578063b98d09081461054c578063cefdc1d41461055957600080fd5b80636efb4636116100ff5780636efb4636146104b3578063715018a6146104d45780637afa1eed146104dc57806384cbfa5c146104ef578063886f11951461050257600080fd5b80635df459461461042b57806367d6be441461045257806368304835146104655780636d14a9871461048c57600080fd5b806344650275116101b357806358a7cd261161018257806358a7cd26146103b5578063595c6a67146103c85780635ac86ab7146103d05780635c155662146104035780635c975abb1461042357600080fd5b8063446502751461034f5780634d2b57fe146103625780634f739f74146103825780635325ada4146103a257600080fd5b80632cb223d5116101fa5780632cb223d5146102ae5780632d89f6fc146102dc57806331b36bd9146102fc5780633563b0d11461031c578063416c7e5e1461033c57600080fd5b806310d67a2f1461022c578063136439dd14610241578063171f1d5b14610254578063245a7bfc14610283575b600080fd5b61023f61023a3660046144a1565b6105ed565b005b61023f61024f3660046144be565b6106a9565b61026761026236600461463c565b6107e8565b6040805192151583529015156020830152015b60405180910390f35b60cb54610296906001600160a01b031681565b6040516001600160a01b03909116815260200161027a565b6102ce6102bc3660046146aa565b60ca6020526000908152604090205481565b60405190815260200161027a565b6102ce6102ea3660046146aa565b60c96020526000908152604090205481565b61030f61030a3660046146f5565b610972565b60405161027a91906147e3565b61032f61032a3660046147f6565b610a8e565b60405161027a9190614951565b61023f61034a36600461497d565b610f26565b61023f61035d3660046144a1565b61109b565b610375610370366004614a00565b6110de565b60405161027a9190614a4f565b610395610390366004614b28565b6111f3565b60405161027a9190614bf2565b61023f6103b0366004614d3c565b611919565b61023f6103c3366004614dc2565b611c05565b61023f611ca9565b6103f36103de366004614e53565b606654600160ff9092169190911b9081161490565b604051901515815260200161027a565b610416610411366004614e70565b611d70565b60405161027a9190614ed3565b6066546102ce565b6102967f000000000000000000000000000000000000000000000000000000000000000081565b6102ce610460366004614f0b565b611f38565b6102967f000000000000000000000000000000000000000000000000000000000000000081565b6102967f000000000000000000000000000000000000000000000000000000000000000081565b6104c66104c1366004615144565b611fbb565b60405161027a929190615204565b61023f612eb3565b60cc54610296906001600160a01b031681565b61023f6104fd36600461524d565b612ec7565b606554610296906001600160a01b031681565b6033546001600160a01b0316610296565b61023f6105343660046152c1565b6130b0565b61023f610547366004615357565b613188565b6097546103f39060ff1681565b61056c6105673660046153ad565b613233565b60405161027a9291906153e4565b6102967f000000000000000000000000000000000000000000000000000000000000000081565b61023f6105af366004615405565b6133c5565b61023f6105c23660046144a1565b613554565b61023f6105d53660046144a1565b61358d565b61023f6105e83660046144be565b613603565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610640573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106649190615497565b6001600160a01b0316336001600160a01b03161461069d5760405162461bcd60e51b8152600401610694906154b4565b60405180910390fd5b6106a68161375f565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071591906154fe565b6107315760405162461bcd60e51b81526004016106949061551b565b606654818116146107aa5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610694565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061083057610830615563565b60200201518951600160200201518a6020015160006002811061085557610855615563565b60200201518b6020015160016002811061087157610871615563565b602090810291909101518c518d8301516040516108ce9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108f19190615579565b905061096461090a6109038884613856565b86906138ed565b610912613981565b61095a61094b85610945604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613856565b6109548c613a41565b906138ed565b886201d4c0613ad1565b909890975095505050505050565b606081516001600160401b0381111561098d5761098d6144d7565b6040519080825280602002602001820160405280156109b6578160200160208202803683370190505b50905060005b8251811015610a8757836001600160a01b03166313542a4e8483815181106109e6576109e6615563565b60200260200101516040518263ffffffff1660e01b8152600401610a1991906001600160a01b0391909116815260200190565b602060405180830381865afa158015610a36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5a919061559b565b828281518110610a6c57610a6c615563565b6020908102919091010152610a80816155ca565b90506109bc565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ad0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af49190615497565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5a9190615497565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc09190615497565b9050600086516001600160401b03811115610bdd57610bdd6144d7565b604051908082528060200260200182016040528015610c1057816020015b6060815260200190600190039081610bfb5790505b50905060005b8751811015610f18576000888281518110610c3357610c33615563565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c94573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610cbc91908101906155e5565b905080516001600160401b03811115610cd757610cd76144d7565b604051908082528060200260200182016040528015610d2257816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610cf55790505b50848481518110610d3557610d35615563565b602002602001018190525060005b8151811015610f02576040518060600160405280876001600160a01b03166347b314e8858581518110610d7857610d78615563565b60200260200101516040518263ffffffff1660e01b8152600401610d9e91815260200190565b602060405180830381865afa158015610dbb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ddf9190615497565b6001600160a01b03168152602001838381518110610dff57610dff615563565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610e2d57610e2d615563565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610e89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ead9190615675565b6001600160601b0316815250858581518110610ecb57610ecb615563565b60200260200101518281518110610ee457610ee4615563565b60200260200101819052508080610efa906155ca565b915050610d43565b5050508080610f10906155ca565b915050610c16565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa89190615497565b6001600160a01b0316336001600160a01b0316146110545760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610694565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60cb546001600160a01b031633146110b257600080fd5b60cb80546001600160a01b039092166001600160a01b0319928316811790915560cc8054909216179055565b606081516001600160401b038111156110f9576110f96144d7565b604051908082528060200260200182016040528015611122578160200160208202803683370190505b50905060005b8251811015610a8757836001600160a01b031663296bb06484838151811061115257611152615563565b60200260200101516040518263ffffffff1660e01b815260040161117891815260200190565b602060405180830381865afa158015611195573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b99190615497565b8282815181106111cb576111cb615563565b6001600160a01b03909216602092830291909101909101526111ec816155ca565b9050611128565b61121e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561125e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112829190615497565b90506112af6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906112df908b90899089906004016156d4565b600060405180830381865afa1580156112fc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261132491908101906156fd565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611356908b908b908b906004016157b4565b600060405180830381865afa158015611373573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261139b91908101906156fd565b6040820152856001600160401b038111156113b8576113b86144d7565b6040519080825280602002602001820160405280156113eb57816020015b60608152602001906001900390816113d65790505b50606082015260005b60ff811687111561182a576000856001600160401b03811115611419576114196144d7565b604051908082528060200260200182016040528015611442578160200160208202803683370190505b5083606001518360ff168151811061145c5761145c615563565b602002602001018190525060005b8681101561172a5760008c6001600160a01b03166304ec63518a8a8581811061149557611495615563565b905060200201358e886000015186815181106114b3576114b3615563565b60200260200101516040518463ffffffff1660e01b81526004016114f09392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561150d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153191906157d4565b90506001600160c01b0381166115d55760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610694565b8a8a8560ff168181106115ea576115ea615563565b6001600160c01b03841692013560f81c9190911c60019081161415905061171757856001600160a01b031663dd9846b98a8a8581811061162c5761162c615563565b905060200201358d8d8860ff1681811061164857611648615563565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561169e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116c291906157fd565b85606001518560ff16815181106116db576116db615563565b602002602001015184815181106116f4576116f4615563565b63ffffffff9092166020928302919091019091015282611713816155ca565b9350505b5080611722816155ca565b91505061146a565b506000816001600160401b03811115611745576117456144d7565b60405190808252806020026020018201604052801561176e578160200160208202803683370190505b50905060005b828110156117ef5784606001518460ff168151811061179557611795615563565b602002602001015181815181106117ae576117ae615563565b60200260200101518282815181106117c8576117c8615563565b63ffffffff90921660209283029190910190910152806117e7816155ca565b915050611774565b508084606001518460ff168151811061180a5761180a615563565b6020026020010181905250505080806118229061581a565b9150506113f4565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561186b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061188f9190615497565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906118c2908b908b908e9060040161583a565b600060405180830381865afa1580156118df573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261190791908101906156fd565b60208301525098975050505050505050565b600061192860208501856146aa565b9050600082516001600160401b03811115611945576119456144d7565b60405190808252806020026020018201604052801561196e578160200160208202803683370190505b50905060005b83518110156119e0576119b184828151811061199257611992615563565b6020026020010151805160009081526020918201519091526040902090565b8282815181106119c3576119c3615563565b6020908102919091010152806119d8816155ca565b915050611974565b5060006119f360c0880160a089016146aa565b82604051602001611a05929190615864565b60405160208183030381529060405280519060200120905084602001358114611a2d57600080fd5b600084516001600160401b03811115611a4857611a486144d7565b604051908082528060200260200182016040528015611a71578160200160208202803683370190505b50905060005b8551811015611b64577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110611ac157611ac1615563565b60200260200101516040518263ffffffff1660e01b8152600401611ae791815260200190565b602060405180830381865afa158015611b04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b289190615497565b828281518110611b3a57611b3a615563565b6001600160a01b039092166020928302919091019091015280611b5c816155ca565b915050611a77565b5060cd5460405163959b127960e01b815263ffffffff861660048201526001600160a01b039091169063959b127990602401600060405180830381600087803b158015611bb057600080fd5b505af1158015611bc4573d6000803e3d6000fd5b505060405133925063ffffffff871691507fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a35050505050505050565b60cd54604051632c53e69360e11b81526001600160a01b03909116906358a7cd2690611c3d90889088908890889088906004016158ac565b600060405180830381600087803b158015611c5757600080fd5b505af1158015611c6b573d6000803e3d6000fd5b505060405133925063ffffffff881691507ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611cf1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d1591906154fe565b611d315760405162461bcd60e51b81526004016106949061551b565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611da29291906158eb565b600060405180830381865afa158015611dbf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611de791908101906156fd565b9050600084516001600160401b03811115611e0457611e046144d7565b604051908082528060200260200182016040528015611e2d578160200160208202803683370190505b50905060005b8551811015611f2e57866001600160a01b03166304ec6351878381518110611e5d57611e5d615563565b602002602001015187868581518110611e7857611e78615563565b60200260200101516040518463ffffffff1660e01b8152600401611eb59392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611ed2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ef691906157d4565b6001600160c01b0316828281518110611f1157611f11615563565b602090810291909101015280611f26816155ca565b915050611e33565b5095945050505050565b60cd546040516319f5af9160e21b815263ffffffff84166004820152602481018390526000916001600160a01b0316906367d6be4490604401602060405180830381865afa158015611f8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fb2919061559b565b90505b92915050565b60408051808201909152606080825260208201526000846120325760405162461bcd60e51b81526020600482015260376024820152600080516020615d6d83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610694565b6040830151518514801561204a575060a08301515185145b801561205a575060c08301515185145b801561206a575060e08301515185145b6120d45760405162461bcd60e51b81526020600482015260416024820152600080516020615d6d83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610694565b8251516020840151511461214c5760405162461bcd60e51b815260206004820152604460248201819052600080516020615d6d833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610694565b4363ffffffff168463ffffffff16106121bb5760405162461bcd60e51b815260206004820152603c6024820152600080516020615d6d83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610694565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156121fc576121fc6144d7565b604051908082528060200260200182016040528015612225578160200160208202803683370190505b506020820152866001600160401b03811115612243576122436144d7565b60405190808252806020026020018201604052801561226c578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156122a0576122a06144d7565b6040519080825280602002602001820160405280156122c9578160200160208202803683370190505b5081526020860151516001600160401b038111156122e9576122e96144d7565b604051908082528060200260200182016040528015612312578160200160208202803683370190505b50816020018190525060006123e48a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156123bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123df919061590a565b613cf5565b905060005b8760200151518110156126605761240f8860200151828151811061199257611992615563565b8360200151828151811061242557612425615563565b602090810291909101015280156124e5576020830151612446600183615927565b8151811061245657612456615563565b602002602001015160001c8360200151828151811061247757612477615563565b602002602001015160001c116124e5576040805162461bcd60e51b8152602060048201526024810191909152600080516020615d6d83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610694565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061252a5761252a615563565b60200260200101518b8b60000151858151811061254957612549615563565b60200260200101516040518463ffffffff1660e01b81526004016125869392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156125a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c791906157d4565b6001600160c01b0316836000015182815181106125e6576125e6615563565b60200260200101818152505061264c610903612620848660000151858151811061261257612612615563565b602002602001015116613d7f565b8a60200151848151811061263657612636615563565b6020026020010151613daa90919063ffffffff16565b945080612658816155ca565b9150506123e9565b505061266b83613e8e565b60975490935060ff16600081612682576000612704565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156126e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612704919061559b565b905060005b8a811015612d82578215612864578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061276057612760615563565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156127a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127c4919061559b565b6127ce919061593e565b116128645760405162461bcd60e51b81526020600482015260666024820152600080516020615d6d83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610694565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106128a5576128a5615563565b9050013560f81c60f81b60f81c8c8c60a0015185815181106128c9576128c9615563565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612925573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129499190615956565b6001600160401b03191661296c8a60400151838151811061199257611992615563565b67ffffffffffffffff191614612a085760405162461bcd60e51b81526020600482015260616024820152600080516020615d6d83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610694565b612a3889604001518281518110612a2157612a21615563565b6020026020010151876138ed90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612a7b57612a7b615563565b9050013560f81c60f81b60f81c8c8c60c001518581518110612a9f57612a9f615563565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612afb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b1f9190615675565b85602001518281518110612b3557612b35615563565b6001600160601b03909216602092830291909101820152850151805182908110612b6157612b61615563565b602002602001015185600001518281518110612b7f57612b7f615563565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612d6d57612bf786600001518281518110612bc957612bc9615563565b60200260200101518f8f86818110612be357612be3615563565b600192013560f81c9290921c811614919050565b15612d5b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612c3d57612c3d615563565b9050013560f81c60f81b60f81c8e89602001518581518110612c6157612c61615563565b60200260200101518f60e001518881518110612c7f57612c7f615563565b60200260200101518781518110612c9857612c98615563565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612cfc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d209190615675565b8751805185908110612d3457612d34615563565b60200260200101818151612d489190615981565b6001600160601b03169052506001909101905b80612d65816155ca565b915050612ba3565b50508080612d7a906155ca565b915050612709565b505050600080612d9c8c868a606001518b608001516107e8565b9150915081612e0d5760405162461bcd60e51b81526020600482015260436024820152600080516020615d6d83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610694565b80612e6e5760405162461bcd60e51b81526020600482015260396024820152600080516020615d6d83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610694565b50506000878260200151604051602001612e89929190615864565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612ebb613f29565b612ec56000613f83565b565b60cb546001600160a01b03163314612ede57600080fd5b6000612ef060c0850160a086016146aa565b9050366000612f0260c08701876159a9565b90925090506000612f1a610100880160e089016146aa565b9050600086604051602001612f2f9190615a0d565b604051602081830303815290604052805190602001209050600080612f578387878a8c611fbb565b9150915060005b85811015612feb578460ff1683602001518281518110612f8057612f80615563565b6020026020010151612f929190615a1b565b6001600160601b0316606484600001518381518110612fb357612fb3615563565b60200260200101516001600160601b0316612fce9190615a4a565b1015612fd957600080fd5b80612fe3816155ca565b915050612f5e565b5060408051808201825263ffffffff431681526020810183905260cd549151634cba41a760e01b815290916001600160a01b031690634cba41a790613038908e908e908690600401615b39565b600060405180830381600087803b15801561305257600080fd5b505af1158015613066573d6000803e3d6000fd5b505050507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a8260405161309b929190615b77565b60405180910390a15050505050505050505050565b60cd54604051639e8e568960e01b81526001600160a01b0390911690639e8e5689906130ea90899089908990899089908990600401615ba1565b600060405180830381600087803b15801561310457600080fd5b505af1158015613118573d6000803e3d6000fd5b50339250505063ffffffff86167f4bfe5edf6d6b0679adcb2253cb622737a0c58a5f67f97765ebc420c10ea47ada8686613153600182615927565b81811061316257613162615563565b9050602002013560405161317891815260200190565b60405180910390a3505050505050565b60cd54604051630c0db69560e31b81526001600160a01b039091169063606db4a8906131bc90869086908690600401615bf5565b600060405180830381600087803b1580156131d657600080fd5b505af11580156131ea573d6000803e3d6000fd5b506131fc9250505060208301836146aa565b63ffffffff167f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a60405160405180910390a2505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061326e5761326e615563565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906132aa90889086906004016158eb565b600060405180830381865afa1580156132c7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526132ef91908101906156fd565b60008151811061330157613301615563565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561336d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061339191906157d4565b6001600160c01b0316905060006133a782613fd5565b9050816133b58a838a610a8e565b9550955050505050935093915050565b60cc546001600160a01b031633146133dc57600080fd5b6133e4614358565b6040805160a081810190925290889060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908716606083015260408051601f8701839004830281018301909152858152908690869081908401838280828437600092018290525060408087019590955287151560808701526001600160a01b0380881660a088015260cd5495519195169350633b7d2cee925061349391508590602001615c8d565b604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004016134c791815260200190565b6020604051808303816000875af11580156134e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061350a91906157fd565b90508063ffffffff167f8fdbcbca31148444d0448386d6e130db10f657ce3f4350386df06f78e71d3ec0836040516135429190615c8d565b60405180910390a25050505050505050565b60cc546001600160a01b0316331461356b57600080fd5b60cd80546001600160a01b0319166001600160a01b0392909216919091179055565b613595613f29565b6001600160a01b0381166135fa5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610694565b6106a681613f83565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613656573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061367a9190615497565b6001600160a01b0316336001600160a01b0316146136aa5760405162461bcd60e51b8152600401610694906154b4565b6066541981196066541916146137285760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610694565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107dd565b6001600160a01b0381166137ed5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610694565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613872614394565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156138a5576138a7565bfe5b50806138e55760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610694565b505092915050565b60408051808201909152600080825260208201526139096143b2565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156138a55750806138e55760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610694565b6139896143d0565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613a71600080516020615d4d83398151915286615579565b90505b613a7d816140a1565b9093509150600080516020615d4d833981519152828309831415613ab7576040805180820190915290815260208101919091529392505050565b600080516020615d4d833981519152600182089050613a74565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613b036143f5565b60005b6002811015613cc8576000613b1c826006615a4a565b9050848260028110613b3057613b30615563565b60200201515183613b4283600061593e565b600c8110613b5257613b52615563565b6020020152848260028110613b6957613b69615563565b60200201516020015183826001613b80919061593e565b600c8110613b9057613b90615563565b6020020152838260028110613ba757613ba7615563565b6020020151515183613bba83600261593e565b600c8110613bca57613bca615563565b6020020152838260028110613be157613be1615563565b6020020151516001602002015183613bfa83600361593e565b600c8110613c0a57613c0a615563565b6020020152838260028110613c2157613c21615563565b602002015160200151600060028110613c3c57613c3c615563565b602002015183613c4d83600461593e565b600c8110613c5d57613c5d615563565b6020020152838260028110613c7457613c74615563565b602002015160200151600160028110613c8f57613c8f615563565b602002015183613ca083600561593e565b600c8110613cb057613cb0615563565b60200201525080613cc0816155ca565b915050613b06565b50613cd1614414565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613d0184614123565b9050808360ff166001901b11611fb25760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610694565b6000805b8215611fb557613d94600184615927565b9092169180613da281615d2a565b915050613d83565b60408051808201909152600080825260208201526102008261ffff1610613e065760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610694565b8161ffff1660011415613e1a575081611fb5565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613e8357600161ffff871660ff83161c81161415613e6657613e6384846138ed565b93505b613e7083846138ed565b92506201fffe600192831b169101613e36565b509195945050505050565b60408051808201909152600080825260208201528151158015613eb357506020820151155b15613ed1575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615d4d8339815191528460200151613f049190615579565b613f1c90600080516020615d4d833981519152615927565b905292915050565b919050565b6033546001600160a01b03163314612ec55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610694565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613fe384613d7f565b61ffff166001600160401b03811115613ffe57613ffe6144d7565b6040519080825280601f01601f191660200182016040528015614028576020820181803683370190505b5090506000805b825182108015614040575061010081105b15614097576001811b935085841615614087578060f81b83838151811061406957614069615563565b60200101906001600160f81b031916908160001a9053508160010191505b614090816155ca565b905061402f565b5090949350505050565b60008080600080516020615d4d8339815191526003600080516020615d4d83398151915286600080516020615d4d833981519152888909090890506000614117827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615d4d8339815191526142b0565b91959194509092505050565b6000610100825111156141ac5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610694565b81516141ba57506000919050565b600080836000815181106141d0576141d0615563565b0160200151600160f89190911c81901b92505b84518110156142a7578481815181106141fe576141fe615563565b0160200151600160f89190911c1b91508282116142935760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610694565b918117916142a0816155ca565b90506141e3565b50909392505050565b6000806142bb614414565b6142c3614432565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156138a557508261434d5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610694565b505195945050505050565b6040518060c0016040528061436b614450565b815260006020820181905260606040830181905282018190526080820181905260a09091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806143e361446e565b81526020016143f061446e565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106a657600080fd5b6000602082840312156144b357600080fd5b8135611fb28161448c565b6000602082840312156144d057600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561450f5761450f6144d7565b60405290565b60405161010081016001600160401b038111828210171561450f5761450f6144d7565b604051601f8201601f191681016001600160401b0381118282101715614560576145606144d7565b604052919050565b60006040828403121561457a57600080fd5b6145826144ed565b9050813581526020820135602082015292915050565b600082601f8301126145a957600080fd5b604051604081018181106001600160401b03821117156145cb576145cb6144d7565b80604052508060408401858111156145e257600080fd5b845b81811015613e835780358352602092830192016145e4565b60006080828403121561460e57600080fd5b6146166144ed565b90506146228383614598565b81526146318360408401614598565b602082015292915050565b600080600080610120858703121561465357600080fd5b843593506146648660208701614568565b925061467386606087016145fc565b91506146828660e08701614568565b905092959194509250565b63ffffffff811681146106a657600080fd5b8035613f248161468d565b6000602082840312156146bc57600080fd5b8135611fb28161468d565b60006001600160401b038211156146e0576146e06144d7565b5060051b60200190565b8035613f248161448c565b6000806040838503121561470857600080fd5b82356147138161448c565b91506020838101356001600160401b0381111561472f57600080fd5b8401601f8101861361474057600080fd5b803561475361474e826146c7565b614538565b81815260059190911b8201830190838101908883111561477257600080fd5b928401925b8284101561479957833561478a8161448c565b82529284019290840190614777565b80955050505050509250929050565b600081518084526020808501945080840160005b838110156147d8578151875295820195908201906001016147bc565b509495945050505050565b602081526000611fb260208301846147a8565b60008060006060848603121561480b57600080fd5b83356148168161448c565b92506020848101356001600160401b038082111561483357600080fd5b818701915087601f83011261484757600080fd5b813581811115614859576148596144d7565b61486b601f8201601f19168501614538565b9150808252888482850101111561488157600080fd5b80848401858401376000848284010152508094505050506148a46040850161469f565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614943578385038a52825180518087529087019087870190845b8181101561492e57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016148ea565b50509a87019a955050918501916001016148cc565b509298975050505050505050565b602081526000611fb260208301846148ad565b80151581146106a657600080fd5b8035613f2481614964565b60006020828403121561498f57600080fd5b8135611fb281614964565b600082601f8301126149ab57600080fd5b813560206149bb61474e836146c7565b82815260059290921b840181019181810190868411156149da57600080fd5b8286015b848110156149f557803583529183019183016149de565b509695505050505050565b60008060408385031215614a1357600080fd5b8235614a1e8161448c565b915060208301356001600160401b03811115614a3957600080fd5b614a458582860161499a565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614a905783516001600160a01b031683529284019291840191600101614a6b565b50909695505050505050565b60008083601f840112614aae57600080fd5b5081356001600160401b03811115614ac557600080fd5b602083019150836020828501011115614add57600080fd5b9250929050565b60008083601f840112614af657600080fd5b5081356001600160401b03811115614b0d57600080fd5b6020830191508360208260051b8501011115614add57600080fd5b60008060008060008060808789031215614b4157600080fd5b8635614b4c8161448c565b95506020870135614b5c8161468d565b945060408701356001600160401b0380821115614b7857600080fd5b614b848a838b01614a9c565b90965094506060890135915080821115614b9d57600080fd5b50614baa89828a01614ae4565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b838110156147d857815163ffffffff1687529582019590820190600101614bd0565b600060208083528351608082850152614c0e60a0850182614bbc565b905081850151601f1980868403016040870152614c2b8383614bbc565b92506040870151915080868403016060870152614c488383614bbc565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614c9f5784878303018452614c8d828751614bbc565b95880195938801939150600101614c73565b509998505050505050505050565b60006101408284031215614cc057600080fd5b50919050565b600060408284031215614cc057600080fd5b600082601f830112614ce957600080fd5b81356020614cf961474e836146c7565b82815260069290921b84018101918181019086841115614d1857600080fd5b8286015b848110156149f557614d2e8882614568565b835291830191604001614d1c565b60008060008060c08587031215614d5257600080fd5b84356001600160401b0380821115614d6957600080fd5b614d7588838901614cad565b9550614d848860208901614cc6565b9450614d938860608901614cc6565b935060a0870135915080821115614da957600080fd5b50614db687828801614cd8565b91505092959194509250565b600080600080600060608688031215614dda57600080fd5b8535614de58161468d565b945060208601356001600160401b0380821115614e0157600080fd5b614e0d89838a01614ae4565b90965094506040880135915080821115614e2657600080fd5b50614e3388828901614a9c565b969995985093965092949392505050565b60ff811681146106a657600080fd5b600060208284031215614e6557600080fd5b8135611fb281614e44565b600080600060608486031215614e8557600080fd5b8335614e908161448c565b925060208401356001600160401b03811115614eab57600080fd5b614eb78682870161499a565b9250506040840135614ec88161468d565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614a9057835183529284019291840191600101614eef565b60008060408385031215614f1e57600080fd5b8235614f298161468d565b946020939093013593505050565b600082601f830112614f4857600080fd5b81356020614f5861474e836146c7565b82815260059290921b84018101918181019086841115614f7757600080fd5b8286015b848110156149f5578035614f8e8161468d565b8352918301918301614f7b565b600082601f830112614fac57600080fd5b81356020614fbc61474e836146c7565b82815260059290921b84018101918181019086841115614fdb57600080fd5b8286015b848110156149f55780356001600160401b03811115614ffe5760008081fd5b61500c8986838b0101614f37565b845250918301918301614fdf565b6000610180828403121561502d57600080fd5b615035614515565b905081356001600160401b038082111561504e57600080fd5b61505a85838601614f37565b8352602084013591508082111561507057600080fd5b61507c85838601614cd8565b6020840152604084013591508082111561509557600080fd5b6150a185838601614cd8565b60408401526150b385606086016145fc565b60608401526150c58560e08601614568565b60808401526101208401359150808211156150df57600080fd5b6150eb85838601614f37565b60a084015261014084013591508082111561510557600080fd5b61511185838601614f37565b60c084015261016084013591508082111561512b57600080fd5b5061513884828501614f9b565b60e08301525092915050565b60008060008060006080868803121561515c57600080fd5b8535945060208601356001600160401b038082111561517a57600080fd5b61518689838a01614a9c565b90965094506040880135915061519b8261468d565b909250606087013590808211156151b157600080fd5b506151be8882890161501a565b9150509295509295909350565b600081518084526020808501945080840160005b838110156147d85781516001600160601b0316875295820195908201906001016151df565b604081526000835160408084015261521f60808401826151cb565b90506020850151603f1984830301606085015261523c82826151cb565b925050508260208301529392505050565b60008060006080848603121561526257600080fd5b83356001600160401b038082111561527957600080fd5b61528587838801614cad565b94506152948760208801614cc6565b935060608601359150808211156152aa57600080fd5b506152b78682870161501a565b9150509250925092565b600080600080600080608087890312156152da57600080fd5b86356001600160401b03808211156152f157600080fd5b6152fd8a838b01614cad565b97506020890135915061530f8261468d565b9095506040880135908082111561532557600080fd5b6153318a838b01614ae4565b9096509450606089013591508082111561534a57600080fd5b50614baa89828a01614a9c565b600080600060a0848603121561536c57600080fd5b83356001600160401b0381111561538257600080fd5b61538e86828701614cad565b93505061539e8560208601614cc6565b91506148a48560608601614cc6565b6000806000606084860312156153c257600080fd5b83356153cd8161448c565b9250602084013591506040840135614ec88161468d565b8281526040602082015260006153fd60408301846148ad565b949350505050565b600080600080600080610120878903121561541f57600080fd5b60a087018881111561543057600080fd5b8796503561543d8161468d565b945060c08701356001600160401b0381111561545857600080fd5b61546489828a01614a9c565b90955093505060e087013561547881614964565b91506101008701356154898161448c565b809150509295509295509295565b6000602082840312156154a957600080fd5b8151611fb28161448c565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561551057600080fd5b8151611fb281614964565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261559657634e487b7160e01b600052601260045260246000fd5b500690565b6000602082840312156155ad57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60006000198214156155de576155de6155b4565b5060010190565b600060208083850312156155f857600080fd5b82516001600160401b0381111561560e57600080fd5b8301601f8101851361561f57600080fd5b805161562d61474e826146c7565b81815260059190911b8201830190838101908783111561564c57600080fd5b928401925b8284101561566a57835182529284019290840190615651565b979650505050505050565b60006020828403121561568757600080fd5b81516001600160601b0381168114611fb257600080fd5b81835260006001600160fb1b038311156156b757600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff841681526040602082015260006156f460408301848661569e565b95945050505050565b6000602080838503121561571057600080fd5b82516001600160401b0381111561572657600080fd5b8301601f8101851361573757600080fd5b805161574561474e826146c7565b81815260059190911b8201830190838101908783111561576457600080fd5b928401925b8284101561566a57835161577c8161468d565b82529284019290840190615769565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006156f460408301848661578b565b6000602082840312156157e657600080fd5b81516001600160c01b0381168114611fb257600080fd5b60006020828403121561580f57600080fd5b8151611fb28161468d565b600060ff821660ff811415615831576158316155b4565b60010192915050565b60408152600061584e60408301858761578b565b905063ffffffff83166020830152949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561589f57815185529382019390820190600101615883565b5092979650505050505050565b63ffffffff861681526060602082015260006158cc60608301868861569e565b82810360408401526158df81858761578b565b98975050505050505050565b63ffffffff831681526040602082015260006153fd60408301846147a8565b60006020828403121561591c57600080fd5b8151611fb281614e44565b600082821015615939576159396155b4565b500390565b60008219821115615951576159516155b4565b500190565b60006020828403121561596857600080fd5b815167ffffffffffffffff1981168114611fb257600080fd5b60006001600160601b03838116908316818110156159a1576159a16155b4565b039392505050565b6000808335601e198436030181126159c057600080fd5b8301803591506001600160401b038211156159da57600080fd5b602001915036819003821315614add57600080fd5b80356159fa8161468d565b63ffffffff168252602090810135910152565b60408101611fb582846159ef565b60006001600160601b0380831681851681830481118215151615615a4157615a416155b4565b02949350505050565b6000816000190483118215151615615a6457615a646155b4565b500290565b600061014060a083853760a08401600080825260a0850135615a8a8161468d565b63ffffffff1690915260c08401359036859003601e19018212615aab578081fd5b9084019081356001600160401b03811115615ac4578182fd5b803603861315615ad2578182fd5b8360c0880152615ae8848801826020860161578b565b9350505050615af960e0840161469f565b63ffffffff1660e0850152610100615b12848201614972565b151590850152610120615b268482016146ea565b6001600160a01b03811686830152614097565b60a081526000615b4c60a0830186615a69565b9050615b5b60208301856159ef565b825163ffffffff166060830152602083015160808301526153fd565b60808101615b8582856159ef565b825163ffffffff16604083015260208301516060830152610f1f565b608081526000615bb46080830189615a69565b63ffffffff881660208401528281036040840152615bd381878961569e565b90508281036060840152615be881858761578b565b9998505050505050505050565b60a081526000615c0860a0830186615a69565b9050615c1760208301856159ef565b8235615c228161468d565b63ffffffff1660608301526020929092013560809091015292915050565b6000815180845260005b81811015615c6657602081850181015186830182015201615c4a565b81811115615c78576000602083870101525b50601f01601f19169290920160200192915050565b6020808252825160009190828483015b6005821015615cbc578251815291830191600191909101908301615c9d565b50505083015163ffffffff1660c0830152604083015161014060e08401819052615cea610160850183615c40565b91506060850151615d0461010086018263ffffffff169052565b506080850151151561012085015260a08501516001600160a01b03811682860152614097565b600061ffff80831681811415615d4257615d426155b4565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220be5b6c39eba5e9c6f341cce3841fe94ce8e6794a84be21f5c14d70d05994057f64736f6c634300080c0033",
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
