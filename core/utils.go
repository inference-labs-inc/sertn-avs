package core

import (
	"fmt"
	"math/big"
	"math/rand"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"
	cstaskmanager "github.com/inference-labs-inc/sertn-avs/contracts/bindings/SertnTaskManager"
	"golang.org/x/crypto/sha3"
)

// this hardcodes abi.encode() for cstaskmanager.ITaskStructTaskResponse
// unclear why abigen doesn't provide this out of the box...
func AbiEncodeTaskResponse(h *cstaskmanager.ITaskStructTaskResponse) ([]byte, error) {

	// The order here has to match the field ordering of cstaskmanager.ITaskStructTaskResponse
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "referenceTaskIndex",
			Type: "uint32",
		},
		{
			Name: "output",
			Type: "uint256",
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	bytes, err := arguments.Pack(h)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GetTaskResponseDigest returns the hash of the TaskResponse, which is what operators sign over
func GetTaskResponseDigest(h *cstaskmanager.ITaskStructTaskResponse) ([32]byte, error) {

	encodeTaskResponseByte, err := AbiEncodeTaskResponse(h)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

// BINDING UTILS - conversion from contract structs to golang structs

// BN254.sol is a library, so bindings for G1 Points and G2 Points are only generated
// in every contract that imports that library. Thus the output here will need to be
// type casted if G1Point is needed to interface with another contract (eg: BLSPublicKeyCompendium.sol)
func ConvertToBN254G1Point(input *bls.G1Point) cstaskmanager.BN254G1Point {
	output := cstaskmanager.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) cstaskmanager.BN254G2Point {
	output := cstaskmanager.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}

func FormatBigIntInputsToString(rawInputs [5]*big.Int) string {
	var inputs []string
	for i := 0; i < 5; i++ {
		inputs = append(inputs, rawInputs[i].String())
	}
	return strings.Join(inputs, " ")
}

func FormatFloatInputsToString(rawInputs [5]float64) string {
	var inputs []string
	for i := 0; i < 5; i++ {
		inputs = append(inputs, strconv.FormatFloat(rawInputs[i], 'f', -1, 64))
	}
	return strings.Join(inputs, " ")
}

func RelativeUrl(pathFromRoot string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error getting relative url", pathFromRoot)
	}
	dirname := filepath.Dir(filename)
	dirname = filepath.Join(dirname, "../", pathFromRoot)
	return dirname
}

// TODO:(opnun) Do in golang
func FormatInputsForChain(rawInputs [5]float64) [5]*big.Int {
	inputString := FormatFloatInputsToString(rawInputs)
	formattedInputs := FormatFloatStringForChain(inputString)
	return formattedInputs
}

func FormatFloatStringForChain(inputString string) [5]*big.Int {
	var formattedInputs [5]*big.Int

	cmd := exec.Command("python", RelativeUrl("python/format.py"), "-i", inputString)

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error: formatting inputs from chain", err, inputString)
	}
	output := strings.Split(string(stdout), "\n")[0]

	for i := 0; i < 5; i++ {
		input := strings.Split(output, " ")[i]
		temp, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Error: parsing integer from string", "Error:", err, "Input: ", input)
		}
		formattedInputs[i] = big.NewInt(temp)
	}
	return formattedInputs
}

func RandomInputs() [5]*big.Int {
	var inputs [5]float64
	for i := 0; i < 5; i++ {
		r := rand.Float64() * 2
		inputs[i] = r
	}
	return FormatInputsForChain(inputs)
}
