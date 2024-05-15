// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOmronTaskManager

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

// IOmronTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IOmronTaskManagerTask struct {
	Inputs                    [5]*big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IOmronTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IOmronTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Output             *big.Int
}

// IOmronTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IOmronTaskManagerTaskResponseMetadata struct {
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

// ContractOmronTaskManagerMetaData contains all meta data concerning the ContractOmronTaskManager contract.
var ContractOmronTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_zkVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOmronTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOmronTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOmronTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620062d1380380620062d18339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615fda620002f7600039600081816102b90152818161060501526134f60152600081816105ce01526125930152600081816104870152818161277501526131de0152600081816104ae0152818161294b0152612b0d0152600081816104d50152818161126e0152818161225e015281816123f601526126300152615fda6000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c80635df45946116101305780638da5cb5b116100b8578063df5cf7231161007c578063df5cf723146105c9578063f2fde38b146105f0578063f5c9899d14610603578063f63c5bab14610629578063fabc1cbc1461063157600080fd5b80638da5cb5b146105645780639c3337e014610575578063b98d090814610588578063bdeea6cc14610595578063cefdc1d4146105a857600080fd5b8063715018a6116100ff578063715018a61461051857806372d18e8d146105205780637afa1eed1461052e578063886f1195146105415780638b00ce7c1461055457600080fd5b80635df459461461048257806368304835146104a95780636d14a987146104d05780636efb4636146104f757600080fd5b806331b36bd9116101be578063595c6a6711610182578063595c6a67146103fc5780635ac86ab7146104045780635c155662146104375780635c975abb146104575780635decc3f51461045f57600080fd5b806331b36bd9146103695780633563b0d114610389578063416c7e5e146103a95780634d2b57fe146103bc5780634f739f74146103dc57600080fd5b8063171f1d5b11610205578063171f1d5b146102855780631ad43189146102b4578063245a7bfc146102f05780632cb223d51461031b5780632d89f6fc1461034957600080fd5b80630627721e1461023757806310d67a2f1461024c578063136439dd1461025f5780631459457a14610272575b600080fd5b61024a610245366004614906565b610644565b005b61024a61025a366004614982565b6107df565b61024a61026d36600461499f565b610892565b61024a6102803660046149b8565b6109d1565b610298610293366004614b8e565b610b30565b6040805192151583529015156020830152015b60405180910390f35b6102db7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102ab565b60cd54610303906001600160a01b031681565b6040516001600160a01b0390911681526020016102ab565b61033b610329366004614bdf565b60cb6020526000908152604090205481565b6040519081526020016102ab565b61033b610357366004614bdf565b60ca6020526000908152604090205481565b61037c610377366004614c1f565b610cba565b6040516102ab9190614d0d565b61039c610397366004614d27565b610dd6565b6040516102ab9190614e82565b61024a6103b7366004614ea3565b61126c565b6103cf6103ca366004614f26565b6113e1565b6040516102ab9190614f75565b6103ef6103ea366004615006565b6114f6565b6040516102ab91906150d0565b61024a611c1c565b61042761041236600461519a565b606654600160ff9092169190911b9081161490565b60405190151581526020016102ab565b61044a6104453660046151b7565b611ce3565b6040516102ab919061521a565b60665461033b565b61042761046d366004614bdf565b60cc6020526000908152604090205460ff1681565b6103037f000000000000000000000000000000000000000000000000000000000000000081565b6103037f000000000000000000000000000000000000000000000000000000000000000081565b6103037f000000000000000000000000000000000000000000000000000000000000000081565b61050a6105053660046154c3565b611eab565b6040516102ab929190615583565b61024a612dc2565b60c95463ffffffff166102db565b60ce54610303906001600160a01b031681565b606554610303906001600160a01b031681565b60c9546102db9063ffffffff1681565b6033546001600160a01b0316610303565b61024a6105833660046155f7565b612dd6565b6097546104279060ff1681565b61024a6105a33660046156d3565b613304565b6105bb6105b6366004615747565b61372b565b6040516102ab92919061577e565b6103037f000000000000000000000000000000000000000000000000000000000000000081565b61024a6105fe366004614982565b6138bd565b7f00000000000000000000000000000000000000000000000000000000000000006102db565b6102db606481565b61024a61063f36600461499f565b613933565b60ce546001600160a01b031633146106ad5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084015b60405180910390fd5b6106b561477b565b6040805160a081810190925290869060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f85018390048302810183019091528381529084908490819084018382808284376000920191909152505050506040808301919091525161073f90829060200161579f565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734906107a290849061579f565b60405180910390a260c9546107be9063ffffffff16600161586b565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610832573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108569190615893565b6001600160a01b0316336001600160a01b0316146108865760405162461bcd60e51b81526004016106a4906158b0565b61088f81613a8f565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156108da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fe91906158fa565b61091a5760405162461bcd60e51b81526004016106a490615917565b606654818116146109935760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016106a4565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156109f15750600054600160ff909116105b80610a0b5750303b158015610a0b575060005460ff166001145b610a6e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106a4565b6000805460ff191660011790558015610a91576000805461ff0019166101001790555b610a9c866000613b86565b610aa585613c70565b60cd80546001600160a01b038087166001600160a01b03199283161790925560ce805486841690831617905560cf8054928516929091169190911790558015610b28576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610b7857610b7861595f565b60200201518951600160200201518a60200151600060028110610b9d57610b9d61595f565b60200201518b60200151600160028110610bb957610bb961595f565b602090810291909101518c518d830151604051610c169a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610c399190615975565b9050610cac610c52610c4b8884613cc2565b8690613d59565b610c5a613ded565b610ca2610c9385610c8d604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613cc2565b610c9c8c613ead565b90613d59565b886201d4c0613f3d565b909890975095505050505050565b606081516001600160401b03811115610cd557610cd5614a29565b604051908082528060200260200182016040528015610cfe578160200160208202803683370190505b50905060005b8251811015610dcf57836001600160a01b03166313542a4e848381518110610d2e57610d2e61595f565b60200260200101516040518263ffffffff1660e01b8152600401610d6191906001600160a01b0391909116815260200190565b602060405180830381865afa158015610d7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da29190615997565b828281518110610db457610db461595f565b6020908102919091010152610dc8816159b0565b9050610d04565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3c9190615893565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea29190615893565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ee4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f089190615893565b9050600086516001600160401b03811115610f2557610f25614a29565b604051908082528060200260200182016040528015610f5857816020015b6060815260200190600190039081610f435790505b50905060005b8751811015611260576000888281518110610f7b57610f7b61595f565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610fdc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261100491908101906159cb565b905080516001600160401b0381111561101f5761101f614a29565b60405190808252806020026020018201604052801561106a57816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161103d5790505b5084848151811061107d5761107d61595f565b602002602001018190525060005b815181101561124a576040518060600160405280876001600160a01b03166347b314e88585815181106110c0576110c061595f565b60200260200101516040518263ffffffff1660e01b81526004016110e691815260200190565b602060405180830381865afa158015611103573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111279190615893565b6001600160a01b031681526020018383815181106111475761114761595f565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106111755761117561595f565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa1580156111d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111f59190615a5b565b6001600160601b03168152508585815181106112135761121361595f565b6020026020010151828151811061122c5761122c61595f565b60200260200101819052508080611242906159b0565b91505061108b565b5050508080611258906159b0565b915050610f5e565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ee9190615893565b6001600160a01b0316336001600160a01b03161461139a5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016106a4565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b606081516001600160401b038111156113fc576113fc614a29565b604051908082528060200260200182016040528015611425578160200160208202803683370190505b50905060005b8251811015610dcf57836001600160a01b031663296bb0648483815181106114555761145561595f565b60200260200101516040518263ffffffff1660e01b815260040161147b91815260200190565b602060405180830381865afa158015611498573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114bc9190615893565b8282815181106114ce576114ce61595f565b6001600160a01b03909216602092830291909101909101526114ef816159b0565b905061142b565b6115216040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611561573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115859190615893565b90506115b26040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906115e2908b9089908990600401615aba565b600060405180830381865afa1580156115ff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116279190810190615ae3565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611659908b908b908b90600401615b9a565b600060405180830381865afa158015611676573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261169e9190810190615ae3565b6040820152856001600160401b038111156116bb576116bb614a29565b6040519080825280602002602001820160405280156116ee57816020015b60608152602001906001900390816116d95790505b50606082015260005b60ff8116871115611b2d576000856001600160401b0381111561171c5761171c614a29565b604051908082528060200260200182016040528015611745578160200160208202803683370190505b5083606001518360ff168151811061175f5761175f61595f565b602002602001018190525060005b86811015611a2d5760008c6001600160a01b03166304ec63518a8a858181106117985761179861595f565b905060200201358e886000015186815181106117b6576117b661595f565b60200260200101516040518463ffffffff1660e01b81526004016117f39392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611810573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118349190615bba565b90506001600160c01b0381166118d85760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a4016106a4565b8a8a8560ff168181106118ed576118ed61595f565b6001600160c01b03841692013560f81c9190911c600190811614159050611a1a57856001600160a01b031663dd9846b98a8a8581811061192f5761192f61595f565b905060200201358d8d8860ff1681811061194b5761194b61595f565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156119a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c59190615be3565b85606001518560ff16815181106119de576119de61595f565b602002602001015184815181106119f7576119f761595f565b63ffffffff9092166020928302919091019091015282611a16816159b0565b9350505b5080611a25816159b0565b91505061176d565b506000816001600160401b03811115611a4857611a48614a29565b604051908082528060200260200182016040528015611a71578160200160208202803683370190505b50905060005b82811015611af25784606001518460ff1681518110611a9857611a9861595f565b60200260200101518181518110611ab157611ab161595f565b6020026020010151828281518110611acb57611acb61595f565b63ffffffff9092166020928302919091019091015280611aea816159b0565b915050611a77565b508084606001518460ff1681518110611b0d57611b0d61595f565b602002602001018190525050508080611b2590615c00565b9150506116f7565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b929190615893565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611bc5908b908b908e90600401615c20565b600060405180830381865afa158015611be2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c0a9190810190615ae3565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611c64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c8891906158fa565b611ca45760405162461bcd60e51b81526004016106a490615917565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611d15929190615c4a565b600060405180830381865afa158015611d32573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d5a9190810190615ae3565b9050600084516001600160401b03811115611d7757611d77614a29565b604051908082528060200260200182016040528015611da0578160200160208202803683370190505b50905060005b8551811015611ea157866001600160a01b03166304ec6351878381518110611dd057611dd061595f565b602002602001015187868581518110611deb57611deb61595f565b60200260200101516040518463ffffffff1660e01b8152600401611e289392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611e45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e699190615bba565b6001600160c01b0316828281518110611e8457611e8461595f565b602090810291909101015280611e99816159b0565b915050611da6565b5095945050505050565b6040805180820190915260608082526020820152600084611f225760405162461bcd60e51b81526020600482015260376024820152600080516020615f8583398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016106a4565b60408301515185148015611f3a575060a08301515185145b8015611f4a575060c08301515185145b8015611f5a575060e08301515185145b611fc45760405162461bcd60e51b81526020600482015260416024820152600080516020615f8583398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016106a4565b8251516020840151511461203c5760405162461bcd60e51b815260206004820152604460248201819052600080516020615f85833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016106a4565b4363ffffffff168463ffffffff16106120ab5760405162461bcd60e51b815260206004820152603c6024820152600080516020615f8583398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016106a4565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156120ec576120ec614a29565b604051908082528060200260200182016040528015612115578160200160208202803683370190505b506020820152866001600160401b0381111561213357612133614a29565b60405190808252806020026020018201604052801561215c578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561219057612190614a29565b6040519080825280602002602001820160405280156121b9578160200160208202803683370190505b5081526020860151516001600160401b038111156121d9576121d9614a29565b604051908082528060200260200182016040528015612202578160200160208202803683370190505b50816020018190525060006122d48a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156122ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122cf9190615c69565b614161565b905060005b87602001515181101561256f5761231e886020015182815181106122ff576122ff61595f565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106123345761233461595f565b602090810291909101015280156123f4576020830151612355600183615c86565b815181106123655761236561595f565b602002602001015160001c836020015182815181106123865761238661595f565b602002602001015160001c116123f4576040805162461bcd60e51b8152602060048201526024810191909152600080516020615f8583398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016106a4565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106124395761243961595f565b60200260200101518b8b6000015185815181106124585761245861595f565b60200260200101516040518463ffffffff1660e01b81526004016124959392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156124b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124d69190615bba565b6001600160c01b0316836000015182815181106124f5576124f561595f565b60200260200101818152505061255b610c4b61252f84866000015185815181106125215761252161595f565b6020026020010151166141f4565b8a6020015184815181106125455761254561595f565b602002602001015161421f90919063ffffffff16565b945080612567816159b0565b9150506122d9565b505061257a83614303565b60975490935060ff16600081612591576000612613565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126139190615997565b905060005b8a811015612c91578215612773578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061266f5761266f61595f565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156126af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126d39190615997565b6126dd9190615c9d565b116127735760405162461bcd60e51b81526020600482015260666024820152600080516020615f8583398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016106a4565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106127b4576127b461595f565b9050013560f81c60f81b60f81c8c8c60a0015185815181106127d8576127d861595f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612834573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128589190615cb5565b6001600160401b03191661287b8a6040015183815181106122ff576122ff61595f565b67ffffffffffffffff1916146129175760405162461bcd60e51b81526020600482015260616024820152600080516020615f8583398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016106a4565b612947896040015182815181106129305761293061595f565b602002602001015187613d5990919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061298a5761298a61595f565b9050013560f81c60f81b60f81c8c8c60c0015185815181106129ae576129ae61595f565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612a0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a2e9190615a5b565b85602001518281518110612a4457612a4461595f565b6001600160601b03909216602092830291909101820152850151805182908110612a7057612a7061595f565b602002602001015185600001518281518110612a8e57612a8e61595f565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612c7c57612b0686600001518281518110612ad857612ad861595f565b60200260200101518f8f86818110612af257612af261595f565b600192013560f81c9290921c811614919050565b15612c6a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612b4c57612b4c61595f565b9050013560f81c60f81b60f81c8e89602001518581518110612b7057612b7061595f565b60200260200101518f60e001518881518110612b8e57612b8e61595f565b60200260200101518781518110612ba757612ba761595f565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612c0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c2f9190615a5b565b8751805185908110612c4357612c4361595f565b60200260200101818151612c579190615ce0565b6001600160601b03169052506001909101905b80612c74816159b0565b915050612ab2565b50508080612c89906159b0565b915050612618565b505050600080612cab8c868a606001518b60800151610b30565b9150915081612d1c5760405162461bcd60e51b81526020600482015260436024820152600080516020615f8583398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016106a4565b80612d7d5760405162461bcd60e51b81526020600482015260396024820152600080516020615f8583398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016106a4565b50506000878260200151604051602001612d98929190615d08565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612dca61439e565b612dd46000613c70565b565b6000612de56020890189614bdf565b63ffffffff8116600090815260cb6020526040902054909150612e0757600080fd5b8787604051602001612e1a929190615d6e565b60408051601f19818403018152918152815160209283012063ffffffff8416600090815260cb90935291205414612e5057600080fd5b63ffffffff8116600090815260cc602052604090205460ff1615612e7357600080fd5b6064612e826020890189614bdf565b612e8c919061586b565b63ffffffff164363ffffffff161115612ecb5760405162461bcd60e51b81526020600482015260016024820152603560f81b60448201526064016106a4565b60005b60058160ff161015612f5f578960ff821660058110612eef57612eef61595f565b60200201356004612f009190615da4565b86868360ff16818110612f1557612f1561595f565b9050602002013514612f4d5760405162461bcd60e51b81526020600482015260016024820152601b60f91b60448201526064016106a4565b80612f5781615c00565b915050612ece565b5060cf54604051631e8e1e1360e01b81526001600160a01b0390911690631e8e1e1390612f9690869086908a908a90600401615dc3565b6020604051808303816000875af1158015612fb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fd991906158fa565b6130095760405162461bcd60e51b81526020600482015260016024820152603760f81b60448201526064016106a4565b600061301a60208a01356004615da4565b8686600581811061302d5761302d61595f565b9050602002013514905080151560011515141561307e57604051339063ffffffff8416907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a350506132fa565b600087516001600160401b0381111561309957613099614a29565b6040519080825280602002602001820160405280156130c2578160200160208202803683370190505b50905060005b8851811015613115576130e68982815181106122ff576122ff61595f565b8282815181106130f8576130f861595f565b60209081029190910101528061310d816159b0565b9150506130c8565b50600061312860c08d0160a08e01614bdf565b8260405160200161313a929190615d08565b604051602081830303815290604052805190602001209050896020013581146131895760405162461bcd60e51b81526020600482015260016024820152600760fb1b60448201526064016106a4565b600089516001600160401b038111156131a4576131a4614a29565b6040519080825280602002602001820160405280156131cd578160200160208202803683370190505b50905060005b8a518110156132c0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae685838151811061321d5761321d61595f565b60200260200101516040518263ffffffff1660e01b815260040161324391815260200190565b602060405180830381865afa158015613260573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132849190615893565b8282815181106132965761329661595f565b6001600160a01b0390921660209283029190910190910152806132b8816159b0565b9150506131d3565b50604051339063ffffffff8716907fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec90600090a350505050505b5050505050505050565b60cd546001600160a01b0316331461335e5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016106a4565b600061337060c0850160a08601614bdf565b905036600061338260c0870187615dea565b9092509050600061339a610100880160e08901614bdf565b905060ca60006133ad6020890189614bdf565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016133d99190615e30565b60405160208183030381529060405280519060200120146134625760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e747261637400000060648201526084016106a4565b600060cb8161347460208a018a614bdf565b63ffffffff1663ffffffff16815260200190815260200160002054146134f15760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016106a4565b61351b7f00000000000000000000000000000000000000000000000000000000000000008561586b565b63ffffffff164363ffffffff16111561353357600080fd5b6000866040516020016135469190615ed9565b60405160208183030381529060405280519060200120905060008061356e8387878a8c611eab565b9150915060005b8581101561366d578460ff16836020015182815181106135975761359761595f565b60200260200101516135a99190615ee7565b6001600160601b03166064846000015183815181106135ca576135ca61595f565b60200260200101516001600160601b03166135e59190615da4565b101561365b576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d60648201526084016106a4565b80613665816159b0565b915050613575565b5060408051808201825263ffffffff4316815260208082018490529151909161369a918c91849101615f16565b6040516020818303038152906040528051906020012060cb60008c60000160208101906136c79190614bdf565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a82604051613716929190615f16565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106137665761376661595f565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906137a29088908690600401615c4a565b600060405180830381865afa1580156137bf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526137e79190810190615ae3565b6000815181106137f9576137f961595f565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613865573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138899190615bba565b6001600160c01b03169050600061389f826143f8565b9050816138ad8a838a610dd6565b9550955050505050935093915050565b6138c561439e565b6001600160a01b03811661392a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106a4565b61088f81613c70565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613986573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139aa9190615893565b6001600160a01b0316336001600160a01b0316146139da5760405162461bcd60e51b81526004016106a4906158b0565b606654198119606654191614613a585760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016106a4565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016109c6565b6001600160a01b038116613b1d5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106a4565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b0316158015613ba757506001600160a01b03821615155b613c295760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106a4565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613c6c82613a8f565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613cde6147a9565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613d1157613d13565bfe5b5080613d515760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016106a4565b505092915050565b6040805180820190915260008082526020820152613d756147c7565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613d11575080613d515760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016106a4565b613df56147e5565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613edd600080516020615f6583398151915286615975565b90505b613ee9816144c4565b9093509150600080516020615f65833981519152828309831415613f23576040805180820190915290815260208101919091529392505050565b600080516020615f65833981519152600182089050613ee0565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613f6f61480a565b60005b6002811015614134576000613f88826006615da4565b9050848260028110613f9c57613f9c61595f565b60200201515183613fae836000615c9d565b600c8110613fbe57613fbe61595f565b6020020152848260028110613fd557613fd561595f565b60200201516020015183826001613fec9190615c9d565b600c8110613ffc57613ffc61595f565b60200201528382600281106140135761401361595f565b6020020151515183614026836002615c9d565b600c81106140365761403661595f565b602002015283826002811061404d5761404d61595f565b6020020151516001602002015183614066836003615c9d565b600c81106140765761407661595f565b602002015283826002811061408d5761408d61595f565b6020020151602001516000600281106140a8576140a861595f565b6020020151836140b9836004615c9d565b600c81106140c9576140c961595f565b60200201528382600281106140e0576140e061595f565b6020020151602001516001600281106140fb576140fb61595f565b60200201518361410c836005615c9d565b600c811061411c5761411c61595f565b6020020152508061412c816159b0565b915050613f72565b5061413d614829565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008061416d84614546565b9050808360ff166001901b116141eb5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016106a4565b90505b92915050565b6000805b82156141ee57614209600184615c86565b909216918061421781615f42565b9150506141f8565b60408051808201909152600080825260208201526102008261ffff161061427b5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016106a4565b8161ffff166001141561428f5750816141ee565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16106142f857600161ffff871660ff83161c811614156142db576142d88484613d59565b93505b6142e58384613d59565b92506201fffe600192831b1691016142ab565b509195945050505050565b6040805180820190915260008082526020820152815115801561432857506020820151155b15614346575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615f6583398151915284602001516143799190615975565b61439190600080516020615f65833981519152615c86565b905292915050565b919050565b6033546001600160a01b03163314612dd45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106a4565b6060600080614406846141f4565b61ffff166001600160401b0381111561442157614421614a29565b6040519080825280601f01601f19166020018201604052801561444b576020820181803683370190505b5090506000805b825182108015614463575061010081105b156144ba576001811b9350858416156144aa578060f81b83838151811061448c5761448c61595f565b60200101906001600160f81b031916908160001a9053508160010191505b6144b3816159b0565b9050614452565b5090949350505050565b60008080600080516020615f658339815191526003600080516020615f6583398151915286600080516020615f6583398151915288890909089050600061453a827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615f658339815191526146d3565b91959194509092505050565b6000610100825111156145cf5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016106a4565b81516145dd57506000919050565b600080836000815181106145f3576145f361595f565b0160200151600160f89190911c81901b92505b84518110156146ca578481815181106146215761462161595f565b0160200151600160f89190911c1b91508282116146b65760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016106a4565b918117916146c3816159b0565b9050614606565b50909392505050565b6000806146de614829565b6146e6614847565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613d115750826147705760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016106a4565b505195945050505050565b604051806080016040528061478e614865565b81526000602082018190526060604083018190529091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806147f8614883565b8152602001614805614883565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461088f57600080fd5b8035614399816148a1565b60008083601f8401126148d057600080fd5b5081356001600160401b038111156148e757600080fd5b6020830191508360208285010111156148ff57600080fd5b9250929050565b60008060008060e0858703121561491c57600080fd5b60a085018681111561492d57600080fd5b8594503561493a816148a1565b925060c08501356001600160401b0381111561495557600080fd5b614961878288016148be565b95989497509550505050565b6001600160a01b038116811461088f57600080fd5b60006020828403121561499457600080fd5b81356141eb8161496d565b6000602082840312156149b157600080fd5b5035919050565b600080600080600060a086880312156149d057600080fd5b85356149db8161496d565b945060208601356149eb8161496d565b935060408601356149fb8161496d565b92506060860135614a0b8161496d565b91506080860135614a1b8161496d565b809150509295509295909350565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614a6157614a61614a29565b60405290565b60405161010081016001600160401b0381118282101715614a6157614a61614a29565b604051601f8201601f191681016001600160401b0381118282101715614ab257614ab2614a29565b604052919050565b600060408284031215614acc57600080fd5b614ad4614a3f565b9050813581526020820135602082015292915050565b600082601f830112614afb57600080fd5b604051604081018181106001600160401b0382111715614b1d57614b1d614a29565b8060405250806040840185811115614b3457600080fd5b845b818110156142f8578035835260209283019201614b36565b600060808284031215614b6057600080fd5b614b68614a3f565b9050614b748383614aea565b8152614b838360408401614aea565b602082015292915050565b6000806000806101208587031215614ba557600080fd5b84359350614bb68660208701614aba565b9250614bc58660608701614b4e565b9150614bd48660e08701614aba565b905092959194509250565b600060208284031215614bf157600080fd5b81356141eb816148a1565b60006001600160401b03821115614c1557614c15614a29565b5060051b60200190565b60008060408385031215614c3257600080fd5b8235614c3d8161496d565b91506020838101356001600160401b03811115614c5957600080fd5b8401601f81018613614c6a57600080fd5b8035614c7d614c7882614bfc565b614a8a565b81815260059190911b82018301908381019088831115614c9c57600080fd5b928401925b82841015614cc3578335614cb48161496d565b82529284019290840190614ca1565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614d0257815187529582019590820190600101614ce6565b509495945050505050565b602081526000614d206020830184614cd2565b9392505050565b600080600060608486031215614d3c57600080fd5b8335614d478161496d565b92506020848101356001600160401b0380821115614d6457600080fd5b818701915087601f830112614d7857600080fd5b813581811115614d8a57614d8a614a29565b614d9c601f8201601f19168501614a8a565b91508082528884828501011115614db257600080fd5b8084840185840137600084828401015250809450505050614dd5604085016148b3565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614e74578385038a52825180518087529087019087870190845b81811015614e5f57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614e1b565b50509a87019a95505091850191600101614dfd565b509298975050505050505050565b602081526000614d206020830184614dde565b801515811461088f57600080fd5b600060208284031215614eb557600080fd5b81356141eb81614e95565b600082601f830112614ed157600080fd5b81356020614ee1614c7883614bfc565b82815260059290921b84018101918181019086841115614f0057600080fd5b8286015b84811015614f1b5780358352918301918301614f04565b509695505050505050565b60008060408385031215614f3957600080fd5b8235614f448161496d565b915060208301356001600160401b03811115614f5f57600080fd5b614f6b85828601614ec0565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614fb65783516001600160a01b031683529284019291840191600101614f91565b50909695505050505050565b60008083601f840112614fd457600080fd5b5081356001600160401b03811115614feb57600080fd5b6020830191508360208260051b85010111156148ff57600080fd5b6000806000806000806080878903121561501f57600080fd5b863561502a8161496d565b9550602087013561503a816148a1565b945060408701356001600160401b038082111561505657600080fd5b6150628a838b016148be565b9096509450606089013591508082111561507b57600080fd5b5061508889828a01614fc2565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614d0257815163ffffffff16875295820195908201906001016150ae565b6000602080835283516080828501526150ec60a085018261509a565b905081850151601f1980868403016040870152615109838361509a565b92506040870151915080868403016060870152615126838361509a565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561517d578487830301845261516b82875161509a565b95880195938801939150600101615151565b509998505050505050505050565b60ff8116811461088f57600080fd5b6000602082840312156151ac57600080fd5b81356141eb8161518b565b6000806000606084860312156151cc57600080fd5b83356151d78161496d565b925060208401356001600160401b038111156151f257600080fd5b6151fe86828701614ec0565b925050604084013561520f816148a1565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614fb657835183529284019291840191600101615236565b600082601f83011261526357600080fd5b81356020615273614c7883614bfc565b82815260059290921b8401810191818101908684111561529257600080fd5b8286015b84811015614f1b5780356152a9816148a1565b8352918301918301615296565b600082601f8301126152c757600080fd5b813560206152d7614c7883614bfc565b82815260069290921b840181019181810190868411156152f657600080fd5b8286015b84811015614f1b5761530c8882614aba565b8352918301916040016152fa565b600082601f83011261532b57600080fd5b8135602061533b614c7883614bfc565b82815260059290921b8401810191818101908684111561535a57600080fd5b8286015b84811015614f1b5780356001600160401b0381111561537d5760008081fd5b61538b8986838b0101615252565b84525091830191830161535e565b600061018082840312156153ac57600080fd5b6153b4614a67565b905081356001600160401b03808211156153cd57600080fd5b6153d985838601615252565b835260208401359150808211156153ef57600080fd5b6153fb858386016152b6565b6020840152604084013591508082111561541457600080fd5b615420858386016152b6565b60408401526154328560608601614b4e565b60608401526154448560e08601614aba565b608084015261012084013591508082111561545e57600080fd5b61546a85838601615252565b60a084015261014084013591508082111561548457600080fd5b61549085838601615252565b60c08401526101608401359150808211156154aa57600080fd5b506154b78482850161531a565b60e08301525092915050565b6000806000806000608086880312156154db57600080fd5b8535945060208601356001600160401b03808211156154f957600080fd5b61550589838a016148be565b90965094506040880135915061551a826148a1565b9092506060870135908082111561553057600080fd5b5061553d88828901615399565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614d025781516001600160601b03168752958201959082019060010161555e565b604081526000835160408084015261559e608084018261554a565b90506020850151603f198483030160608501526155bb828261554a565b925050508260208301529392505050565b600061010082840312156155df57600080fd5b50919050565b6000604082840312156155df57600080fd5b600080600080600080600080610100898b03121561561457600080fd5b88356001600160401b038082111561562b57600080fd5b6156378c838d016155cc565b99506156468c60208d016155e5565b98506156558c60608d016155e5565b975060a08b013591508082111561566b57600080fd5b6156778c838d016152b6565b965060c08b013591508082111561568d57600080fd5b6156998c838d01614fc2565b909650945060e08b01359150808211156156b257600080fd5b506156bf8b828c016148be565b999c989b5096995094979396929594505050565b6000806000608084860312156156e857600080fd5b83356001600160401b03808211156156ff57600080fd5b61570b878388016155cc565b945061571a87602088016155e5565b9350606086013591508082111561573057600080fd5b5061573d86828701615399565b9150509250925092565b60008060006060848603121561575c57600080fd5b83356157678161496d565b925060208401359150604084013561520f816148a1565b8281526040602082015260006157976040830184614dde565b949350505050565b6020808252825160009190828483015b60058210156157ce5782518152918301916001919091019083016157af565b50505063ffffffff818501511660c084015260408401516101008060e086015281518061012087015260005b8181101561581757838101850151878201610140015284016157fa565b8181111561582a57600061014083890101525b50606087015163ffffffff8116878401529350601f01601f1916949094016101400195945050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff80831681851680830382111561588a5761588a615855565b01949350505050565b6000602082840312156158a557600080fd5b81516141eb8161496d565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561590c57600080fd5b81516141eb81614e95565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261599257634e487b7160e01b600052601260045260246000fd5b500690565b6000602082840312156159a957600080fd5b5051919050565b60006000198214156159c4576159c4615855565b5060010190565b600060208083850312156159de57600080fd5b82516001600160401b038111156159f457600080fd5b8301601f81018513615a0557600080fd5b8051615a13614c7882614bfc565b81815260059190911b82018301908381019087831115615a3257600080fd5b928401925b82841015615a5057835182529284019290840190615a37565b979650505050505050565b600060208284031215615a6d57600080fd5b81516001600160601b03811681146141eb57600080fd5b81835260006001600160fb1b03831115615a9d57600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff84168152604060208201526000615ada604083018486615a84565b95945050505050565b60006020808385031215615af657600080fd5b82516001600160401b03811115615b0c57600080fd5b8301601f81018513615b1d57600080fd5b8051615b2b614c7882614bfc565b81815260059190911b82018301908381019087831115615b4a57600080fd5b928401925b82841015615a50578351615b62816148a1565b82529284019290840190615b4f565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615ada604083018486615b71565b600060208284031215615bcc57600080fd5b81516001600160c01b03811681146141eb57600080fd5b600060208284031215615bf557600080fd5b81516141eb816148a1565b600060ff821660ff811415615c1757615c17615855565b60010192915050565b604081526000615c34604083018587615b71565b905063ffffffff83166020830152949350505050565b63ffffffff831681526040602082015260006157976040830184614cd2565b600060208284031215615c7b57600080fd5b81516141eb8161518b565b600082821015615c9857615c98615855565b500390565b60008219821115615cb057615cb0615855565b500190565b600060208284031215615cc757600080fd5b815167ffffffffffffffff19811681146141eb57600080fd5b60006001600160601b0383811690831681811015615d0057615d00615855565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615d4357815185529382019390820190600101615d27565b5092979650505050505050565b8035615d5b816148a1565b63ffffffff168252602090810135910152565b60808101615d7c8285615d50565b8235615d87816148a1565b63ffffffff16604083015260209290920135606090910152919050565b6000816000190483118215151615615dbe57615dbe615855565b500290565b604081526000615dd7604083018688615b71565b8281036020840152615a50818587615a84565b6000808335601e19843603018112615e0157600080fd5b8301803591506001600160401b03821115615e1b57600080fd5b6020019150368190038213156148ff57600080fd5b6020815260a0826020830137600060c08201600080825260a0850135615e55816148a1565b63ffffffff1690915260c08401359036859003601e19018212615e76578081fd5b9084019081356001600160401b03811115615e8f578182fd5b803603861315615e9d578182fd5b61010091508160e0860152615eba61012086018260208601615b71565b925050615ec960e086016148b3565b63ffffffff8116858301526144ba565b604081016141ee8284615d50565b60006001600160601b0380831681851681830481118215151615615f0d57615f0d615855565b02949350505050565b60808101615f248285615d50565b63ffffffff8351166040830152602083015160608301529392505050565b600061ffff80831681811415615f5a57615f5a615855565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122065e24b1d7d87ec400b477d4976075473fc4697809b149b7590078049b7ee17b964736f6c634300080c0033",
}

// ContractOmronTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOmronTaskManagerMetaData.ABI instead.
var ContractOmronTaskManagerABI = ContractOmronTaskManagerMetaData.ABI

// ContractOmronTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOmronTaskManagerMetaData.Bin instead.
var ContractOmronTaskManagerBin = ContractOmronTaskManagerMetaData.Bin

// DeployContractOmronTaskManager deploys a new Ethereum contract, binding an instance of ContractOmronTaskManager to it.
func DeployContractOmronTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractOmronTaskManager, error) {
	parsed, err := ContractOmronTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOmronTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOmronTaskManager{ContractOmronTaskManagerCaller: ContractOmronTaskManagerCaller{contract: contract}, ContractOmronTaskManagerTransactor: ContractOmronTaskManagerTransactor{contract: contract}, ContractOmronTaskManagerFilterer: ContractOmronTaskManagerFilterer{contract: contract}}, nil
}

// ContractOmronTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractOmronTaskManager struct {
	ContractOmronTaskManagerCaller     // Read-only binding to the contract
	ContractOmronTaskManagerTransactor // Write-only binding to the contract
	ContractOmronTaskManagerFilterer   // Log filterer for contract events
}

// ContractOmronTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOmronTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOmronTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOmronTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOmronTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOmronTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOmronTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOmronTaskManagerSession struct {
	Contract     *ContractOmronTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractOmronTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOmronTaskManagerCallerSession struct {
	Contract *ContractOmronTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractOmronTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOmronTaskManagerTransactorSession struct {
	Contract     *ContractOmronTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractOmronTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOmronTaskManagerRaw struct {
	Contract *ContractOmronTaskManager // Generic contract binding to access the raw methods on
}

// ContractOmronTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOmronTaskManagerCallerRaw struct {
	Contract *ContractOmronTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOmronTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOmronTaskManagerTransactorRaw struct {
	Contract *ContractOmronTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOmronTaskManager creates a new instance of ContractOmronTaskManager, bound to a specific deployed contract.
func NewContractOmronTaskManager(address common.Address, backend bind.ContractBackend) (*ContractOmronTaskManager, error) {
	contract, err := bindContractOmronTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManager{ContractOmronTaskManagerCaller: ContractOmronTaskManagerCaller{contract: contract}, ContractOmronTaskManagerTransactor: ContractOmronTaskManagerTransactor{contract: contract}, ContractOmronTaskManagerFilterer: ContractOmronTaskManagerFilterer{contract: contract}}, nil
}

// NewContractOmronTaskManagerCaller creates a new read-only instance of ContractOmronTaskManager, bound to a specific deployed contract.
func NewContractOmronTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractOmronTaskManagerCaller, error) {
	contract, err := bindContractOmronTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerCaller{contract: contract}, nil
}

// NewContractOmronTaskManagerTransactor creates a new write-only instance of ContractOmronTaskManager, bound to a specific deployed contract.
func NewContractOmronTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOmronTaskManagerTransactor, error) {
	contract, err := bindContractOmronTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerTransactor{contract: contract}, nil
}

// NewContractOmronTaskManagerFilterer creates a new log filterer instance of ContractOmronTaskManager, bound to a specific deployed contract.
func NewContractOmronTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOmronTaskManagerFilterer, error) {
	contract, err := bindContractOmronTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerFilterer{contract: contract}, nil
}

// bindContractOmronTaskManager binds a generic wrapper to an already deployed contract.
func bindContractOmronTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOmronTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOmronTaskManager *ContractOmronTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOmronTaskManager.Contract.ContractOmronTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOmronTaskManager *ContractOmronTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.ContractOmronTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOmronTaskManager *ContractOmronTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.ContractOmronTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOmronTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractOmronTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractOmronTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractOmronTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractOmronTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOmronTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOmronTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractOmronTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractOmronTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Aggregator(&_ContractOmronTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Aggregator(&_ContractOmronTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOmronTaskManager.Contract.AllTaskHashes(&_ContractOmronTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractOmronTaskManager.Contract.AllTaskHashes(&_ContractOmronTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOmronTaskManager.Contract.AllTaskResponses(&_ContractOmronTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractOmronTaskManager.Contract.AllTaskResponses(&_ContractOmronTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.BlsApkRegistry(&_ContractOmronTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.BlsApkRegistry(&_ContractOmronTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOmronTaskManager.Contract.CheckSignatures(&_ContractOmronTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractOmronTaskManager.Contract.CheckSignatures(&_ContractOmronTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Delegation(&_ContractOmronTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Delegation(&_ContractOmronTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Generator() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Generator(&_ContractOmronTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Generator(&_ContractOmronTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOmronTaskManager.Contract.GetBatchOperatorFromId(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOmronTaskManager.Contract.GetBatchOperatorFromId(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOmronTaskManager.Contract.GetBatchOperatorId(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOmronTaskManager.Contract.GetBatchOperatorId(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOmronTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOmronTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOmronTaskManager.Contract.GetCheckSignaturesIndices(&_ContractOmronTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOmronTaskManager.Contract.GetOperatorState(&_ContractOmronTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOmronTaskManager.Contract.GetOperatorState(&_ContractOmronTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOmronTaskManager.Contract.GetOperatorState0(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOmronTaskManager.Contract.GetOperatorState0(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOmronTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOmronTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOmronTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOmronTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOmronTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractOmronTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractOmronTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractOmronTaskManager.Contract.LatestTaskNum(&_ContractOmronTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractOmronTaskManager.Contract.LatestTaskNum(&_ContractOmronTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Owner() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Owner(&_ContractOmronTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.Owner(&_ContractOmronTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractOmronTaskManager.Contract.Paused(&_ContractOmronTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractOmronTaskManager.Contract.Paused(&_ContractOmronTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractOmronTaskManager.Contract.Paused0(&_ContractOmronTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractOmronTaskManager.Contract.Paused0(&_ContractOmronTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.PauserRegistry(&_ContractOmronTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.PauserRegistry(&_ContractOmronTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.RegistryCoordinator(&_ContractOmronTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.RegistryCoordinator(&_ContractOmronTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.StakeRegistry(&_ContractOmronTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractOmronTaskManager.Contract.StakeRegistry(&_ContractOmronTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractOmronTaskManager.Contract.StaleStakesForbidden(&_ContractOmronTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractOmronTaskManager.Contract.StaleStakesForbidden(&_ContractOmronTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractOmronTaskManager.Contract.TaskNumber(&_ContractOmronTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractOmronTaskManager.Contract.TaskNumber(&_ContractOmronTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractOmronTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractOmronTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractOmronTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractOmronTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOmronTaskManager.Contract.TrySignatureAndApkVerification(&_ContractOmronTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractOmronTaskManager.Contract.TrySignatureAndApkVerification(&_ContractOmronTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "createNewTask", inputs, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.CreateNewTask(&_ContractOmronTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x0627721e.
//
// Solidity: function createNewTask(uint256[5] inputs, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) CreateNewTask(inputs [5]*big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.CreateNewTask(&_ContractOmronTaskManager.TransactOpts, inputs, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _zkVerifier) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _zkVerifier common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator, _zkVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _zkVerifier) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _zkVerifier common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.Initialize(&_ContractOmronTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _zkVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _zkVerifier) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _zkVerifier common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.Initialize(&_ContractOmronTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _zkVerifier)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.Pause(&_ContractOmronTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.Pause(&_ContractOmronTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.PauseAll(&_ContractOmronTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.PauseAll(&_ContractOmronTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x9c3337e0.
//
// Solidity: function raiseAndResolveChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators, uint256[] instances, bytes proof) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators, instances, proof)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x9c3337e0.
//
// Solidity: function raiseAndResolveChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators, uint256[] instances, bytes proof) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) RaiseAndResolveChallenge(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RaiseAndResolveChallenge(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators, instances, proof)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x9c3337e0.
//
// Solidity: function raiseAndResolveChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators, uint256[] instances, bytes proof) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) RaiseAndResolveChallenge(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RaiseAndResolveChallenge(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators, instances, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RenounceOwnership(&_ContractOmronTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RenounceOwnership(&_ContractOmronTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) RespondToTask(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RespondToTask(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xbdeea6cc.
//
// Solidity: function respondToTask((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) RespondToTask(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RespondToTask(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.SetPauserRegistry(&_ContractOmronTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.SetPauserRegistry(&_ContractOmronTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.SetStaleStakesForbidden(&_ContractOmronTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.SetStaleStakesForbidden(&_ContractOmronTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.TransferOwnership(&_ContractOmronTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.TransferOwnership(&_ContractOmronTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.Unpause(&_ContractOmronTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.Unpause(&_ContractOmronTaskManager.TransactOpts, newPausedStatus)
}

// ContractOmronTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerInitializedIterator struct {
	Event *ContractOmronTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerInitialized)
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
		it.Event = new(ContractOmronTaskManagerInitialized)
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
func (it *ContractOmronTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerInitialized represents a Initialized event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOmronTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerInitializedIterator{contract: _ContractOmronTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerInitialized)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractOmronTaskManagerInitialized, error) {
	event := new(ContractOmronTaskManagerInitialized)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerNewTaskCreatedIterator struct {
	Event *ContractOmronTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerNewTaskCreated)
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
		it.Event = new(ContractOmronTaskManagerNewTaskCreated)
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
func (it *ContractOmronTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IOmronTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32) task)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOmronTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerNewTaskCreatedIterator{contract: _ContractOmronTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f5497734.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256[5],uint32,bytes,uint32) task)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerNewTaskCreated)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractOmronTaskManagerNewTaskCreated, error) {
	event := new(ContractOmronTaskManagerNewTaskCreated)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerOwnershipTransferredIterator struct {
	Event *ContractOmronTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractOmronTaskManagerOwnershipTransferred)
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
func (it *ContractOmronTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOmronTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerOwnershipTransferredIterator{contract: _ContractOmronTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerOwnershipTransferred)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOmronTaskManagerOwnershipTransferred, error) {
	event := new(ContractOmronTaskManagerOwnershipTransferred)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerPausedIterator struct {
	Event *ContractOmronTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerPaused)
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
		it.Event = new(ContractOmronTaskManagerPaused)
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
func (it *ContractOmronTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerPaused represents a Paused event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractOmronTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerPausedIterator{contract: _ContractOmronTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerPaused)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParsePaused(log types.Log) (*ContractOmronTaskManagerPaused, error) {
	event := new(ContractOmronTaskManagerPaused)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerPauserRegistrySetIterator struct {
	Event *ContractOmronTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractOmronTaskManagerPauserRegistrySet)
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
func (it *ContractOmronTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractOmronTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerPauserRegistrySetIterator{contract: _ContractOmronTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerPauserRegistrySet)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractOmronTaskManagerPauserRegistrySet, error) {
	event := new(ContractOmronTaskManagerPauserRegistrySet)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractOmronTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractOmronTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractOmronTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractOmronTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractOmronTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractOmronTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractOmronTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerTaskChallengedIterator is returned from FilterTaskChallenged and is used to iterate over the raw logs and unpacked data for TaskChallenged events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskChallengedIterator struct {
	Event *ContractOmronTaskManagerTaskChallenged // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerTaskChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerTaskChallenged)
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
		it.Event = new(ContractOmronTaskManagerTaskChallenged)
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
func (it *ContractOmronTaskManagerTaskChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerTaskChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerTaskChallenged represents a TaskChallenged event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskChallenged struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskChallenged is a free log retrieval operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterTaskChallenged(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOmronTaskManagerTaskChallengedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerTaskChallengedIterator{contract: _ContractOmronTaskManager.contract, event: "TaskChallenged", logs: logs, sub: sub}, nil
}

// WatchTaskChallenged is a free log subscription operation binding the contract event 0x03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a.
//
// Solidity: event TaskChallenged(uint32 indexed taskIndex)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchTaskChallenged(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerTaskChallenged, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "TaskChallenged", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerTaskChallenged)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseTaskChallenged(log types.Log) (*ContractOmronTaskManagerTaskChallenged, error) {
	event := new(ContractOmronTaskManagerTaskChallenged)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractOmronTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractOmronTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractOmronTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractOmronTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractOmronTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerTaskChallengedSuccessfully)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractOmronTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractOmronTaskManagerTaskChallengedSuccessfully)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractOmronTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractOmronTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractOmronTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractOmronTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractOmronTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractOmronTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractOmronTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskCompletedIterator struct {
	Event *ContractOmronTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerTaskCompleted)
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
		it.Event = new(ContractOmronTaskManagerTaskCompleted)
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
func (it *ContractOmronTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractOmronTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerTaskCompletedIterator{contract: _ContractOmronTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerTaskCompleted)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractOmronTaskManagerTaskCompleted, error) {
	event := new(ContractOmronTaskManagerTaskCompleted)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskRespondedIterator struct {
	Event *ContractOmronTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerTaskResponded)
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
		it.Event = new(ContractOmronTaskManagerTaskResponded)
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
func (it *ContractOmronTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerTaskResponded represents a TaskResponded event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerTaskResponded struct {
	TaskResponse         IOmronTaskManagerTaskResponse
	TaskResponseMetadata IOmronTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractOmronTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerTaskRespondedIterator{contract: _ContractOmronTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerTaskResponded)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractOmronTaskManagerTaskResponded, error) {
	event := new(ContractOmronTaskManagerTaskResponded)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerUnpausedIterator struct {
	Event *ContractOmronTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractOmronTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronTaskManagerUnpaused)
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
		it.Event = new(ContractOmronTaskManagerUnpaused)
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
func (it *ContractOmronTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronTaskManagerUnpaused represents a Unpaused event raised by the ContractOmronTaskManager contract.
type ContractOmronTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractOmronTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronTaskManagerUnpausedIterator{contract: _ContractOmronTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractOmronTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractOmronTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronTaskManagerUnpaused)
				if err := _ContractOmronTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractOmronTaskManager *ContractOmronTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractOmronTaskManagerUnpaused, error) {
	event := new(ContractOmronTaskManagerUnpaused)
	if err := _ContractOmronTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
