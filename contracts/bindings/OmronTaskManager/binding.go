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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeData\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskProven\",\"type\":\"uint8\",\"internalType\":\"enumIOmronTaskManager.ChallengeStatus\"},{\"name\":\"timeChallenged\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeInstances\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_zkVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveResultAccurate\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"instances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"raiseChallenger\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIOmronTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOmronTaskManager.Task\",\"components\":[{\"name\":\"inputs\",\"type\":\"uint256[5]\",\"internalType\":\"uint256[5]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallenged\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOmronTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"output\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOmronTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006268380380620062688339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615f71620002f760003960008181610348015281816106d2015261342b01526000818161069b0152612afa01526000818161053c015281816116770152612cdc01526000818161058e01528181612eb201526130740152600081816105b50152818161133b015281816127e40152818161295d0152612b970152615f716000f3fe608060405234801561001057600080fd5b506004361061025e5760003560e01c80635c975abb11610146578063886f1195116100c3578063cefdc1d411610087578063cefdc1d414610675578063df5cf72314610696578063f2fde38b146106bd578063f5c9899d146106d0578063f63c5bab146106f6578063fabc1cbc146106fe57600080fd5b8063886f1195146106215780638b00ce7c146106345780638da5cb5b14610644578063b98d090814610655578063bdeea6cc1461066257600080fd5b80636d14a9871161010a5780636d14a987146105b05780636efb4636146105d7578063715018a6146105f857806372d18e8d146106005780637afa1eed1461060e57600080fd5b80635c975abb1461050c5780635decc3f5146105145780635df459461461053757806367d6be441461055e578063683048351461058957600080fd5b80632d89f6fc116101df5780634d2b57fe116101a35780634d2b57fe1461045e5780634f739f741461047e57806358a7cd261461049e578063595c6a67146104b15780635ac86ab7146104b95780635c155662146104ec57600080fd5b80632d89f6fc146103d857806331b36bd9146103f85780633563b0d114610418578063416c7e5e14610438578063480bab6b1461044b57600080fd5b8063162d8a3f11610226578063162d8a3f14610306578063171f1d5b146103195780631ad4318914610343578063245a7bfc1461037f5780632cb223d5146103aa57600080fd5b80630627721e1461026357806310d67a2f146102785780631151d0b91461028b578063136439dd146102e05780631459457a146102f3575b600080fd5b6102766102713660046147d0565b610711565b005b61027661028636600461484c565b61085a565b6102c8610299366004614869565b60cd60205260009081526040902080546001909101546001600160a01b03821691600160a01b900460ff169083565b6040516102d79392919061489c565b60405180910390f35b6102766102ee3660046148db565b610916565b6102766103013660046148f4565b610a55565b610276610314366004614990565b610bb4565b61032c610327366004614b54565b610bfd565b6040805192151583529015156020830152016102d7565b61036a7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102d7565b60cf54610392906001600160a01b031681565b6040516001600160a01b0390911681526020016102d7565b6103ca6103b8366004614869565b60cb6020526000908152604090205481565b6040519081526020016102d7565b6103ca6103e6366004614869565b60ca6020526000908152604090205481565b61040b610406366004614bc8565b610d87565b6040516102d79190614cb6565b61042b610426366004614cd0565b610ea3565b6040516102d79190614e22565b610276610446366004614e43565b611339565b610276610459366004614ecf565b6114ae565b61047161046c366004614fb0565b611813565b6040516102d79190614fff565b61049161048c366004615090565b611928565b6040516102d7919061515a565b6102766104ac366004615215565b61204e565b6102766121a2565b6104dc6104c73660046152a6565b606654600160ff9092169190911b9081161490565b60405190151581526020016102d7565b6104ff6104fa3660046152c3565b612269565b6040516102d79190615326565b6066546103ca565b6104dc610522366004614869565b60cc6020526000908152604090205460ff1681565b6103927f000000000000000000000000000000000000000000000000000000000000000081565b6103ca61056c36600461535e565b60ce60209081526000928352604080842090915290825290205481565b6103927f000000000000000000000000000000000000000000000000000000000000000081565b6103927f000000000000000000000000000000000000000000000000000000000000000081565b6105ea6105e5366004615597565b612431565b6040516102d7929190615657565b610276613329565b60c95463ffffffff1661036a565b60d054610392906001600160a01b031681565b606554610392906001600160a01b031681565b60c95461036a9063ffffffff1681565b6033546001600160a01b0316610392565b6097546104dc9060ff1681565b6102766106703660046156a0565b61333d565b610688610683366004615714565b6135f5565b6040516102d792919061574b565b6103927f000000000000000000000000000000000000000000000000000000000000000081565b6102766106cb36600461484c565b613787565b7f000000000000000000000000000000000000000000000000000000000000000061036a565b61036a606481565b61027661070c3660046148db565b6137fd565b60d0546001600160a01b0316331461072857600080fd5b610730614645565b6040805160a081810190925290869060059083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f8501839004830281018301909152838152908490849081908401838280828437600092019190915250505050604080830191909152516107ba90829060200161576c565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f70b3096d1238f88c5789bee6d4b63455de2e1c40ed61a383d87b24e5f54977349061081d90849061576c565b60405180910390a260c9546108399063ffffffff166001615838565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d19190615860565b6001600160a01b0316336001600160a01b03161461090a5760405162461bcd60e51b81526004016109019061587d565b60405180910390fd5b61091381613959565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561095e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098291906158c7565b61099e5760405162461bcd60e51b8152600401610901906158e4565b60665481811614610a175760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610901565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff1615808015610a755750600054600160ff909116105b80610a8f5750303b158015610a8f575060005460ff166001145b610af25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610901565b6000805460ff191660011790558015610b15576000805461ff0019166101001790555b610b20866000613a50565b610b2985613b3a565b60cf80546001600160a01b038087166001600160a01b03199283161790925560d0805486841690831617905560d18054928516929091169190911790558015610bac576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6000610bc36020840184614869565b60405190915063ffffffff8216907f03db1f7b7bfeecd3e8b085c7922e488a65d21504b70d2b1d330abeb07c9f214a90600090a250505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610c4557610c4561592c565b60200201518951600160200201518a60200151600060028110610c6a57610c6a61592c565b60200201518b60200151600160028110610c8657610c8661592c565b602090810291909101518c518d830151604051610ce39a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d069190615942565b9050610d79610d1f610d188884613b8c565b8690613c23565b610d27613cb7565b610d6f610d6085610d5a604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613b8c565b610d698c613d77565b90613c23565b886201d4c0613e07565b909890975095505050505050565b606081516001600160401b03811115610da257610da26149ef565b604051908082528060200260200182016040528015610dcb578160200160208202803683370190505b50905060005b8251811015610e9c57836001600160a01b03166313542a4e848381518110610dfb57610dfb61592c565b60200260200101516040518263ffffffff1660e01b8152600401610e2e91906001600160a01b0391909116815260200190565b602060405180830381865afa158015610e4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6f9190615964565b828281518110610e8157610e8161592c565b6020908102919091010152610e958161597d565b9050610dd1565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ee5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f099190615860565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6f9190615860565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd59190615860565b9050600086516001600160401b03811115610ff257610ff26149ef565b60405190808252806020026020018201604052801561102557816020015b60608152602001906001900390816110105790505b50905060005b875181101561132d5760008882815181106110485761104861592c565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa1580156110a9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110d19190810190615998565b905080516001600160401b038111156110ec576110ec6149ef565b60405190808252806020026020018201604052801561113757816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161110a5790505b5084848151811061114a5761114a61592c565b602002602001018190525060005b8151811015611317576040518060600160405280876001600160a01b03166347b314e885858151811061118d5761118d61592c565b60200260200101516040518263ffffffff1660e01b81526004016111b391815260200190565b602060405180830381865afa1580156111d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111f49190615860565b6001600160a01b031681526020018383815181106112145761121461592c565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106112425761124261592c565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa15801561129e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c29190615a28565b6001600160601b03168152508585815181106112e0576112e061592c565b602002602001015182815181106112f9576112f961592c565b6020026020010181905250808061130f9061597d565b915050611158565b50505080806113259061597d565b91505061102b565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611397573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113bb9190615860565b6001600160a01b0316336001600160a01b0316146114675760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610901565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b60006114bd6020850185614869565b9050600163ffffffff8216600090815260cd6020526040902054600160a01b900460ff1660038111156114f2576114f2614886565b148015611516575063ffffffff8116600090815260cd602052604090206001015442115b61151f57600080fd5b600082516001600160401b0381111561153a5761153a6149ef565b604051908082528060200260200182016040528015611563578160200160208202803683370190505b50905060005b83518110156115d5576115a68482815181106115875761158761592c565b6020026020010151805160009081526020918201519091526040902090565b8282815181106115b8576115b861592c565b6020908102919091010152806115cd8161597d565b915050611569565b5060006115e860c0880160a08901614869565b826040516020016115fa929190615a51565b6040516020818303038152906040528051906020012090508460200135811461162257600080fd5b600084516001600160401b0381111561163d5761163d6149ef565b604051908082528060200260200182016040528015611666578160200160208202803683370190505b50905060005b8551811015611759577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106116b6576116b661592c565b60200260200101516040518263ffffffff1660e01b81526004016116dc91815260200190565b602060405180830381865afa1580156116f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171d9190615860565b82828151811061172f5761172f61592c565b6001600160a01b0390921660209283029190910190910152806117518161597d565b91505061166c565b5063ffffffff8416600090815260cd6020526040902060010154421161177e57600080fd5b600163ffffffff8516600090815260cd6020526040902054600160a01b900460ff1660038111156117b1576117b1614886565b146117bb57600080fd5b63ffffffff8416600081815260cd6020526040808220805460ff60a01b1916600360a01b179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505050565b606081516001600160401b0381111561182e5761182e6149ef565b604051908082528060200260200182016040528015611857578160200160208202803683370190505b50905060005b8251811015610e9c57836001600160a01b031663296bb0648483815181106118875761188761592c565b60200260200101516040518263ffffffff1660e01b81526004016118ad91815260200190565b602060405180830381865afa1580156118ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ee9190615860565b8282815181106119005761190061592c565b6001600160a01b03909216602092830291909101909101526119218161597d565b905061185d565b6119536040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611993573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119b79190615860565b90506119e46040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611a14908b9089908990600401615acf565b600060405180830381865afa158015611a31573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a599190810190615af8565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611a8b908b908b908b90600401615baf565b600060405180830381865afa158015611aa8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ad09190810190615af8565b6040820152856001600160401b03811115611aed57611aed6149ef565b604051908082528060200260200182016040528015611b2057816020015b6060815260200190600190039081611b0b5790505b50606082015260005b60ff8116871115611f5f576000856001600160401b03811115611b4e57611b4e6149ef565b604051908082528060200260200182016040528015611b77578160200160208202803683370190505b5083606001518360ff1681518110611b9157611b9161592c565b602002602001018190525060005b86811015611e5f5760008c6001600160a01b03166304ec63518a8a85818110611bca57611bca61592c565b905060200201358e88600001518681518110611be857611be861592c565b60200260200101516040518463ffffffff1660e01b8152600401611c259392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611c42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c669190615bcf565b90506001600160c01b038116611d0a5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610901565b8a8a8560ff16818110611d1f57611d1f61592c565b6001600160c01b03841692013560f81c9190911c600190811614159050611e4c57856001600160a01b031663dd9846b98a8a85818110611d6157611d6161592c565b905060200201358d8d8860ff16818110611d7d57611d7d61592c565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611dd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611df79190615bf8565b85606001518560ff1681518110611e1057611e1061592c565b60200260200101518481518110611e2957611e2961592c565b63ffffffff9092166020928302919091019091015282611e488161597d565b9350505b5080611e578161597d565b915050611b9f565b506000816001600160401b03811115611e7a57611e7a6149ef565b604051908082528060200260200182016040528015611ea3578160200160208202803683370190505b50905060005b82811015611f245784606001518460ff1681518110611eca57611eca61592c565b60200260200101518181518110611ee357611ee361592c565b6020026020010151828281518110611efd57611efd61592c565b63ffffffff9092166020928302919091019091015280611f1c8161597d565b915050611ea9565b508084606001518460ff1681518110611f3f57611f3f61592c565b602002602001018190525050508080611f5790615c15565b915050611b29565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc49190615860565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611ff7908b908b908e90600401615c35565b600060405180830381865afa158015612014573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261203c9190810190615af8565b60208301525098975050505050505050565b63ffffffff8516600090815260cd60205260408120905b60068110156120c8578585828181106120805761208061592c565b63ffffffff8a16600090815260ce602090815260408083208784528252909120549102929092013590911490506120b657600080fd5b806120c08161597d565b915050612065565b5060d154604051631e8e1e1360e01b81526001600160a01b0390911690631e8e1e13906120ff90869086908a908a90600401615c5f565b6020604051808303816000875af115801561211e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061214291906158c7565b61214b57600080fd5b8054600160a11b60ff60a01b1982161782556040516001600160a01b039091169063ffffffff8816907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a3505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156121ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061220e91906158c7565b61222a5760405162461bcd60e51b8152600401610901906158e4565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161229b929190615c86565b600060405180830381865afa1580156122b8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122e09190810190615af8565b9050600084516001600160401b038111156122fd576122fd6149ef565b604051908082528060200260200182016040528015612326578160200160208202803683370190505b50905060005b855181101561242757866001600160a01b03166304ec63518783815181106123565761235661592c565b6020026020010151878685815181106123715761237161592c565b60200260200101516040518463ffffffff1660e01b81526004016123ae9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156123cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123ef9190615bcf565b6001600160c01b031682828151811061240a5761240a61592c565b60209081029190910101528061241f8161597d565b91505061232c565b5095945050505050565b60408051808201909152606080825260208201526000846124a85760405162461bcd60e51b81526020600482015260376024820152600080516020615f1c83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610901565b604083015151851480156124c0575060a08301515185145b80156124d0575060c08301515185145b80156124e0575060e08301515185145b61254a5760405162461bcd60e51b81526020600482015260416024820152600080516020615f1c83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610901565b825151602084015151146125c25760405162461bcd60e51b815260206004820152604460248201819052600080516020615f1c833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610901565b4363ffffffff168463ffffffff16106126315760405162461bcd60e51b815260206004820152603c6024820152600080516020615f1c83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610901565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115612672576126726149ef565b60405190808252806020026020018201604052801561269b578160200160208202803683370190505b506020820152866001600160401b038111156126b9576126b96149ef565b6040519080825280602002602001820160405280156126e2578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115612716576127166149ef565b60405190808252806020026020018201604052801561273f578160200160208202803683370190505b5081526020860151516001600160401b0381111561275f5761275f6149ef565b604051908082528060200260200182016040528015612788578160200160208202803683370190505b508160200181905250600061285a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612831573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128559190615ca5565b61402b565b905060005b876020015151811015612ad657612885886020015182815181106115875761158761592c565b8360200151828151811061289b5761289b61592c565b6020908102919091010152801561295b5760208301516128bc600183615cc2565b815181106128cc576128cc61592c565b602002602001015160001c836020015182815181106128ed576128ed61592c565b602002602001015160001c1161295b576040805162461bcd60e51b8152602060048201526024810191909152600080516020615f1c83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610901565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106129a0576129a061592c565b60200260200101518b8b6000015185815181106129bf576129bf61592c565b60200260200101516040518463ffffffff1660e01b81526004016129fc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612a19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a3d9190615bcf565b6001600160c01b031683600001518281518110612a5c57612a5c61592c565b602002602001018181525050612ac2610d18612a968486600001518581518110612a8857612a8861592c565b6020026020010151166140be565b8a602001518481518110612aac57612aac61592c565b60200260200101516140e990919063ffffffff16565b945080612ace8161597d565b91505061285f565b5050612ae1836141cd565b60975490935060ff16600081612af8576000612b7a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b7a9190615964565b905060005b8a8110156131f8578215612cda578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612bd657612bd661592c565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612c16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c3a9190615964565b612c449190615cd9565b11612cda5760405162461bcd60e51b81526020600482015260666024820152600080516020615f1c83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610901565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612d1b57612d1b61592c565b9050013560f81c60f81b60f81c8c8c60a001518581518110612d3f57612d3f61592c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612d9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dbf9190615cf1565b6001600160401b031916612de28a6040015183815181106115875761158761592c565b67ffffffffffffffff191614612e7e5760405162461bcd60e51b81526020600482015260616024820152600080516020615f1c83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610901565b612eae89604001518281518110612e9757612e9761592c565b602002602001015187613c2390919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612ef157612ef161592c565b9050013560f81c60f81b60f81c8c8c60c001518581518110612f1557612f1561592c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612f71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f959190615a28565b85602001518281518110612fab57612fab61592c565b6001600160601b03909216602092830291909101820152850151805182908110612fd757612fd761592c565b602002602001015185600001518281518110612ff557612ff561592c565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156131e35761306d8660000151828151811061303f5761303f61592c565b60200260200101518f8f868181106130595761305961592c565b600192013560f81c9290921c811614919050565b156131d1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106130b3576130b361592c565b9050013560f81c60f81b60f81c8e896020015185815181106130d7576130d761592c565b60200260200101518f60e0015188815181106130f5576130f561592c565b6020026020010151878151811061310e5761310e61592c565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015613172573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131969190615a28565b87518051859081106131aa576131aa61592c565b602002602001018181516131be9190615d1c565b6001600160601b03169052506001909101905b806131db8161597d565b915050613019565b505080806131f09061597d565b915050612b7f565b5050506000806132128c868a606001518b60800151610bfd565b91509150816132835760405162461bcd60e51b81526020600482015260436024820152600080516020615f1c83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610901565b806132e45760405162461bcd60e51b81526020600482015260396024820152600080516020615f1c83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610901565b505060008782602001516040516020016132ff929190615a51565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b613331614268565b61333b6000613b3a565b565b60cf546001600160a01b0316331461335457600080fd5b600061336660c0850160a08601614869565b905036600061337860c0870187615d44565b90925090506000613390610100880160e08901614869565b905060ca60006133a36020890189614869565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016133cf9190615d8a565b60405160208183030381529060405280519060200120146133ef57600080fd5b600060cb8161340160208a018a614869565b63ffffffff1663ffffffff168152602001908152602001600020541461342657600080fd5b6134507f000000000000000000000000000000000000000000000000000000000000000085615838565b63ffffffff164363ffffffff16111561346857600080fd5b60008660405160200161347b9190615e51565b6040516020818303038152906040528051906020012090506000806134a38387878a8c612431565b9150915060005b85811015613537578460ff16836020015182815181106134cc576134cc61592c565b60200260200101516134de9190615e5f565b6001600160601b03166064846000015183815181106134ff576134ff61592c565b60200260200101516001600160601b031661351a9190615e8e565b101561352557600080fd5b8061352f8161597d565b9150506134aa565b5060408051808201825263ffffffff43168152602080820184905291519091613564918c91849101615ead565b6040516020818303038152906040528051906020012060cb60008c60000160208101906135919190614869565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a826040516135e0929190615ead565b60405180910390a15050505050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106136305761363061592c565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e9061366c9088908690600401615c86565b600060405180830381865afa158015613689573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526136b19190810190615af8565b6000815181106136c3576136c361592c565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561372f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137539190615bcf565b6001600160c01b031690506000613769826142c2565b9050816137778a838a610ea3565b9550955050505050935093915050565b61378f614268565b6001600160a01b0381166137f45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610901565b61091381613b3a565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613850573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138749190615860565b6001600160a01b0316336001600160a01b0316146138a45760405162461bcd60e51b81526004016109019061587d565b6066541981196066541916146139225760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610901565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610a4a565b6001600160a01b0381166139e75760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610901565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b0316158015613a7157506001600160a01b03821615155b613af35760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610901565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613b3682613959565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040805180820190915260008082526020820152613ba8614673565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613bdb57613bdd565bfe5b5080613c1b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610901565b505092915050565b6040805180820190915260008082526020820152613c3f614691565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613bdb575080613c1b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610901565b613cbf6146af565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613da7600080516020615efc83398151915286615942565b90505b613db38161438e565b9093509150600080516020615efc833981519152828309831415613ded576040805180820190915290815260208101919091529392505050565b600080516020615efc833981519152600182089050613daa565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613e396146d4565b60005b6002811015613ffe576000613e52826006615e8e565b9050848260028110613e6657613e6661592c565b60200201515183613e78836000615cd9565b600c8110613e8857613e8861592c565b6020020152848260028110613e9f57613e9f61592c565b60200201516020015183826001613eb69190615cd9565b600c8110613ec657613ec661592c565b6020020152838260028110613edd57613edd61592c565b6020020151515183613ef0836002615cd9565b600c8110613f0057613f0061592c565b6020020152838260028110613f1757613f1761592c565b6020020151516001602002015183613f30836003615cd9565b600c8110613f4057613f4061592c565b6020020152838260028110613f5757613f5761592c565b602002015160200151600060028110613f7257613f7261592c565b602002015183613f83836004615cd9565b600c8110613f9357613f9361592c565b6020020152838260028110613faa57613faa61592c565b602002015160200151600160028110613fc557613fc561592c565b602002015183613fd6836005615cd9565b600c8110613fe657613fe661592c565b60200201525080613ff68161597d565b915050613e3c565b506140076146f3565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008061403784614410565b9050808360ff166001901b116140b55760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610901565b90505b92915050565b6000805b82156140b8576140d3600184615cc2565b90921691806140e181615ed9565b9150506140c2565b60408051808201909152600080825260208201526102008261ffff16106141455760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610901565b8161ffff16600114156141595750816140b8565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16106141c257600161ffff871660ff83161c811614156141a5576141a28484613c23565b93505b6141af8384613c23565b92506201fffe600192831b169101614175565b509195945050505050565b604080518082019091526000808252602082015281511580156141f257506020820151155b15614210575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615efc83398151915284602001516142439190615942565b61425b90600080516020615efc833981519152615cc2565b905292915050565b919050565b6033546001600160a01b0316331461333b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610901565b60606000806142d0846140be565b61ffff166001600160401b038111156142eb576142eb6149ef565b6040519080825280601f01601f191660200182016040528015614315576020820181803683370190505b5090506000805b82518210801561432d575061010081105b15614384576001811b935085841615614374578060f81b8383815181106143565761435661592c565b60200101906001600160f81b031916908160001a9053508160010191505b61437d8161597d565b905061431c565b5090949350505050565b60008080600080516020615efc8339815191526003600080516020615efc83398151915286600080516020615efc833981519152888909090890506000614404827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615efc83398151915261459d565b91959194509092505050565b6000610100825111156144995760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610901565b81516144a757506000919050565b600080836000815181106144bd576144bd61592c565b0160200151600160f89190911c81901b92505b8451811015614594578481815181106144eb576144eb61592c565b0160200151600160f89190911c1b91508282116145805760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610901565b9181179161458d8161597d565b90506144d0565b50909392505050565b6000806145a86146f3565b6145b0614711565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613bdb57508261463a5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610901565b505195945050505050565b604051806080016040528061465861472f565b81526000602082018190526060604083018190529091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806146c261474d565b81526020016146cf61474d565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b63ffffffff8116811461091357600080fd5b80356142638161476b565b60008083601f84011261479a57600080fd5b5081356001600160401b038111156147b157600080fd5b6020830191508360208285010111156147c957600080fd5b9250929050565b60008060008060e085870312156147e657600080fd5b60a08501868111156147f757600080fd5b859450356148048161476b565b925060c08501356001600160401b0381111561481f57600080fd5b61482b87828801614788565b95989497509550505050565b6001600160a01b038116811461091357600080fd5b60006020828403121561485e57600080fd5b81356140b581614837565b60006020828403121561487b57600080fd5b81356140b58161476b565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b038416815260608101600484106148ca57634e487b7160e01b600052602160045260246000fd5b602082019390935260400152919050565b6000602082840312156148ed57600080fd5b5035919050565b600080600080600060a0868803121561490c57600080fd5b853561491781614837565b9450602086013561492781614837565b9350604086013561493781614837565b9250606086013561494781614837565b9150608086013561495781614837565b809150509295509295909350565b6000610100828403121561497857600080fd5b50919050565b60006040828403121561497857600080fd5b600080600060a084860312156149a557600080fd5b83356001600160401b038111156149bb57600080fd5b6149c786828701614965565b9350506149d7856020860161497e565b91506149e6856060860161497e565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614a2757614a276149ef565b60405290565b60405161010081016001600160401b0381118282101715614a2757614a276149ef565b604051601f8201601f191681016001600160401b0381118282101715614a7857614a786149ef565b604052919050565b600060408284031215614a9257600080fd5b614a9a614a05565b9050813581526020820135602082015292915050565b600082601f830112614ac157600080fd5b604051604081018181106001600160401b0382111715614ae357614ae36149ef565b8060405250806040840185811115614afa57600080fd5b845b818110156141c2578035835260209283019201614afc565b600060808284031215614b2657600080fd5b614b2e614a05565b9050614b3a8383614ab0565b8152614b498360408401614ab0565b602082015292915050565b6000806000806101208587031215614b6b57600080fd5b84359350614b7c8660208701614a80565b9250614b8b8660608701614b14565b9150614b9a8660e08701614a80565b905092959194509250565b60006001600160401b03821115614bbe57614bbe6149ef565b5060051b60200190565b60008060408385031215614bdb57600080fd5b8235614be681614837565b91506020838101356001600160401b03811115614c0257600080fd5b8401601f81018613614c1357600080fd5b8035614c26614c2182614ba5565b614a50565b81815260059190911b82018301908381019088831115614c4557600080fd5b928401925b82841015614c6c578335614c5d81614837565b82529284019290840190614c4a565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614cab57815187529582019590820190600101614c8f565b509495945050505050565b602081526000614cc96020830184614c7b565b9392505050565b600080600060608486031215614ce557600080fd5b8335614cf081614837565b92506020848101356001600160401b0380821115614d0d57600080fd5b818701915087601f830112614d2157600080fd5b813581811115614d3357614d336149ef565b614d45601f8201601f19168501614a50565b91508082528884828501011115614d5b57600080fd5b80848401858401376000848284010152508094505050506149e66040850161477d565b600081518084526020808501808196508360051b810191508286016000805b86811015614e14578385038a52825180518087529087019087870190845b81811015614dff57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614dbb565b50509a87019a95505091850191600101614d9d565b509298975050505050505050565b602081526000614cc96020830184614d7e565b801515811461091357600080fd5b600060208284031215614e5557600080fd5b81356140b581614e35565b600082601f830112614e7157600080fd5b81356020614e81614c2183614ba5565b82815260069290921b84018101918181019086841115614ea057600080fd5b8286015b84811015614ec457614eb68882614a80565b835291830191604001614ea4565b509695505050505050565b60008060008060c08587031215614ee557600080fd5b84356001600160401b0380821115614efc57600080fd5b614f0888838901614965565b9550614f17886020890161497e565b9450614f26886060890161497e565b935060a0870135915080821115614f3c57600080fd5b50614f4987828801614e60565b91505092959194509250565b600082601f830112614f6657600080fd5b81356020614f76614c2183614ba5565b82815260059290921b84018101918181019086841115614f9557600080fd5b8286015b84811015614ec45780358352918301918301614f99565b60008060408385031215614fc357600080fd5b8235614fce81614837565b915060208301356001600160401b03811115614fe957600080fd5b614ff585828601614f55565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b818110156150405783516001600160a01b03168352928401929184019160010161501b565b50909695505050505050565b60008083601f84011261505e57600080fd5b5081356001600160401b0381111561507557600080fd5b6020830191508360208260051b85010111156147c957600080fd5b600080600080600080608087890312156150a957600080fd5b86356150b481614837565b955060208701356150c48161476b565b945060408701356001600160401b03808211156150e057600080fd5b6150ec8a838b01614788565b9096509450606089013591508082111561510557600080fd5b5061511289828a0161504c565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614cab57815163ffffffff1687529582019590820190600101615138565b60006020808352835160808285015261517660a0850182615124565b905081850151601f19808684030160408701526151938383615124565b925060408701519150808684030160608701526151b08383615124565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561520757848783030184526151f5828751615124565b958801959388019391506001016151db565b509998505050505050505050565b60008060008060006060868803121561522d57600080fd5b85356152388161476b565b945060208601356001600160401b038082111561525457600080fd5b61526089838a0161504c565b9096509450604088013591508082111561527957600080fd5b5061528688828901614788565b969995985093965092949392505050565b60ff8116811461091357600080fd5b6000602082840312156152b857600080fd5b81356140b581615297565b6000806000606084860312156152d857600080fd5b83356152e381614837565b925060208401356001600160401b038111156152fe57600080fd5b61530a86828701614f55565b925050604084013561531b8161476b565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b8181101561504057835183529284019291840191600101615342565b6000806040838503121561537157600080fd5b823561537c8161476b565b946020939093013593505050565b600082601f83011261539b57600080fd5b813560206153ab614c2183614ba5565b82815260059290921b840181019181810190868411156153ca57600080fd5b8286015b84811015614ec45780356153e18161476b565b83529183019183016153ce565b600082601f8301126153ff57600080fd5b8135602061540f614c2183614ba5565b82815260059290921b8401810191818101908684111561542e57600080fd5b8286015b84811015614ec45780356001600160401b038111156154515760008081fd5b61545f8986838b010161538a565b845250918301918301615432565b6000610180828403121561548057600080fd5b615488614a2d565b905081356001600160401b03808211156154a157600080fd5b6154ad8583860161538a565b835260208401359150808211156154c357600080fd5b6154cf85838601614e60565b602084015260408401359150808211156154e857600080fd5b6154f485838601614e60565b60408401526155068560608601614b14565b60608401526155188560e08601614a80565b608084015261012084013591508082111561553257600080fd5b61553e8583860161538a565b60a084015261014084013591508082111561555857600080fd5b6155648583860161538a565b60c084015261016084013591508082111561557e57600080fd5b5061558b848285016153ee565b60e08301525092915050565b6000806000806000608086880312156155af57600080fd5b8535945060208601356001600160401b03808211156155cd57600080fd5b6155d989838a01614788565b9096509450604088013591506155ee8261476b565b9092506060870135908082111561560457600080fd5b506156118882890161546d565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614cab5781516001600160601b031687529582019590820190600101615632565b6040815260008351604080840152615672608084018261561e565b90506020850151603f1984830301606085015261568f828261561e565b925050508260208301529392505050565b6000806000608084860312156156b557600080fd5b83356001600160401b03808211156156cc57600080fd5b6156d887838801614965565b94506156e7876020880161497e565b935060608601359150808211156156fd57600080fd5b5061570a8682870161546d565b9150509250925092565b60008060006060848603121561572957600080fd5b833561573481614837565b925060208401359150604084013561531b8161476b565b8281526040602082015260006157646040830184614d7e565b949350505050565b6020808252825160009190828483015b600582101561579b57825181529183019160019190910190830161577c565b50505063ffffffff818501511660c084015260408401516101008060e086015281518061012087015260005b818110156157e457838101850151878201610140015284016157c7565b818111156157f757600061014083890101525b50606087015163ffffffff8116878401529350601f01601f1916949094016101400195945050505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff80831681851680830382111561585757615857615822565b01949350505050565b60006020828403121561587257600080fd5b81516140b581614837565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156158d957600080fd5b81516140b581614e35565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261595f57634e487b7160e01b600052601260045260246000fd5b500690565b60006020828403121561597657600080fd5b5051919050565b600060001982141561599157615991615822565b5060010190565b600060208083850312156159ab57600080fd5b82516001600160401b038111156159c157600080fd5b8301601f810185136159d257600080fd5b80516159e0614c2182614ba5565b81815260059190911b820183019083810190878311156159ff57600080fd5b928401925b82841015615a1d57835182529284019290840190615a04565b979650505050505050565b600060208284031215615a3a57600080fd5b81516001600160601b03811681146140b557600080fd5b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615a8c57815185529382019390820190600101615a70565b5092979650505050505050565b81835260006001600160fb1b03831115615ab257600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff84168152604060208201526000615aef604083018486615a99565b95945050505050565b60006020808385031215615b0b57600080fd5b82516001600160401b03811115615b2157600080fd5b8301601f81018513615b3257600080fd5b8051615b40614c2182614ba5565b81815260059190911b82018301908381019087831115615b5f57600080fd5b928401925b82841015615a1d578351615b778161476b565b82529284019290840190615b64565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615aef604083018486615b86565b600060208284031215615be157600080fd5b81516001600160c01b03811681146140b557600080fd5b600060208284031215615c0a57600080fd5b81516140b58161476b565b600060ff821660ff811415615c2c57615c2c615822565b60010192915050565b604081526000615c49604083018587615b86565b905063ffffffff83166020830152949350505050565b604081526000615c73604083018688615b86565b8281036020840152615a1d818587615a99565b63ffffffff831681526040602082015260006157646040830184614c7b565b600060208284031215615cb757600080fd5b81516140b581615297565b600082821015615cd457615cd4615822565b500390565b60008219821115615cec57615cec615822565b500190565b600060208284031215615d0357600080fd5b815167ffffffffffffffff19811681146140b557600080fd5b60006001600160601b0383811690831681811015615d3c57615d3c615822565b039392505050565b6000808335601e19843603018112615d5b57600080fd5b8301803591506001600160401b03821115615d7557600080fd5b6020019150368190038213156147c957600080fd5b6020815260a0826020830137600060c08201600080825260a0850135615daf8161476b565b63ffffffff1690915260c08401359036859003601e19018212615dd0578081fd5b9084019081356001600160401b03811115615de9578182fd5b803603861315615df7578182fd5b61010091508160e0860152615e1461012086018260208601615b86565b925050615e2360e0860161477d565b63ffffffff811685830152614384565b8035615e3e8161476b565b63ffffffff168252602090810135910152565b604081016140b88284615e33565b60006001600160601b0380831681851681830481118215151615615e8557615e85615822565b02949350505050565b6000816000190483118215151615615ea857615ea8615822565b500290565b60808101615ebb8285615e33565b63ffffffff8351166040830152602083015160608301529392505050565b600061ffff80831681811415615ef157615ef1615822565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220b165875e2ad09abb534991f3a32b15c36c72df50ee2e25ae1f0a4a8f73f4e6ef64736f6c634300080c0033",
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

// ChallengeData is a free data retrieval call binding the contract method 0x1151d0b9.
//
// Solidity: function challengeData(uint32 ) view returns(address challenger, uint8 taskProven, uint256 timeChallenged)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) ChallengeData(opts *bind.CallOpts, arg0 uint32) (struct {
	Challenger     common.Address
	TaskProven     uint8
	TimeChallenged *big.Int
}, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "challengeData", arg0)

	outstruct := new(struct {
		Challenger     common.Address
		TaskProven     uint8
		TimeChallenged *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenger = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TaskProven = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.TimeChallenged = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChallengeData is a free data retrieval call binding the contract method 0x1151d0b9.
//
// Solidity: function challengeData(uint32 ) view returns(address challenger, uint8 taskProven, uint256 timeChallenged)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) ChallengeData(arg0 uint32) (struct {
	Challenger     common.Address
	TaskProven     uint8
	TimeChallenged *big.Int
}, error) {
	return _ContractOmronTaskManager.Contract.ChallengeData(&_ContractOmronTaskManager.CallOpts, arg0)
}

// ChallengeData is a free data retrieval call binding the contract method 0x1151d0b9.
//
// Solidity: function challengeData(uint32 ) view returns(address challenger, uint8 taskProven, uint256 timeChallenged)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) ChallengeData(arg0 uint32) (struct {
	Challenger     common.Address
	TaskProven     uint8
	TimeChallenged *big.Int
}, error) {
	return _ContractOmronTaskManager.Contract.ChallengeData(&_ContractOmronTaskManager.CallOpts, arg0)
}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 , uint256 ) view returns(uint256)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCaller) ChallengeInstances(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractOmronTaskManager.contract.Call(opts, &out, "challengeInstances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 , uint256 ) view returns(uint256)
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) ChallengeInstances(arg0 uint32, arg1 *big.Int) (*big.Int, error) {
	return _ContractOmronTaskManager.Contract.ChallengeInstances(&_ContractOmronTaskManager.CallOpts, arg0, arg1)
}

// ChallengeInstances is a free data retrieval call binding the contract method 0x67d6be44.
//
// Solidity: function challengeInstances(uint32 , uint256 ) view returns(uint256)
func (_ContractOmronTaskManager *ContractOmronTaskManagerCallerSession) ChallengeInstances(arg0 uint32, arg1 *big.Int) (*big.Int, error) {
	return _ContractOmronTaskManager.Contract.ChallengeInstances(&_ContractOmronTaskManager.CallOpts, arg0, arg1)
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

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x480bab6b.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) ConfirmChallenge(opts *bind.TransactOpts, task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "confirmChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x480bab6b.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) ConfirmChallenge(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.ConfirmChallenge(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// ConfirmChallenge is a paid mutator transaction binding the contract method 0x480bab6b.
//
// Solidity: function confirmChallenge((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) ConfirmChallenge(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.ConfirmChallenge(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
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

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) ProveResultAccurate(opts *bind.TransactOpts, taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "proveResultAccurate", taskId, instances, proof)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) ProveResultAccurate(taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.ProveResultAccurate(&_ContractOmronTaskManager.TransactOpts, taskId, instances, proof)
}

// ProveResultAccurate is a paid mutator transaction binding the contract method 0x58a7cd26.
//
// Solidity: function proveResultAccurate(uint32 taskId, uint256[] instances, bytes proof) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) ProveResultAccurate(taskId uint32, instances []*big.Int, proof []byte) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.ProveResultAccurate(&_ContractOmronTaskManager.TransactOpts, taskId, instances, proof)
}

// RaiseChallenger is a paid mutator transaction binding the contract method 0x162d8a3f.
//
// Solidity: function raiseChallenger((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactor) RaiseChallenger(opts *bind.TransactOpts, task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractOmronTaskManager.contract.Transact(opts, "raiseChallenger", task, taskResponse, taskResponseMetadata)
}

// RaiseChallenger is a paid mutator transaction binding the contract method 0x162d8a3f.
//
// Solidity: function raiseChallenger((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerSession) RaiseChallenger(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RaiseChallenger(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
}

// RaiseChallenger is a paid mutator transaction binding the contract method 0x162d8a3f.
//
// Solidity: function raiseChallenger((uint256[5],uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata) returns()
func (_ContractOmronTaskManager *ContractOmronTaskManagerTransactorSession) RaiseChallenger(task IOmronTaskManagerTask, taskResponse IOmronTaskManagerTaskResponse, taskResponseMetadata IOmronTaskManagerTaskResponseMetadata) (*types.Transaction, error) {
	return _ContractOmronTaskManager.Contract.RaiseChallenger(&_ContractOmronTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata)
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
