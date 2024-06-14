// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSertnServiceManager

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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractSertnServiceManagerMetaData contains all meta data concerning the ContractSertnServiceManager contract.
var ContractSertnServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_sertnTaskManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sertnTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISertnTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b4e38038062001b4e83398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118d76200027760003960008181610175015261093b01526000818161065a015281816107b60152818161084d01528181610c5701528181610ddb0152610e7a0152600081816104850152818161051401528181610594015281816109e901528181610a7f01528181610b950152610d3601526000818161031501526103f301526000818161010c01528181610a3d01528181610adb0152610b5a01526118d76000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639926ee7d116100715780639926ee7d1461015d578063a2bc040014610170578063a364f4da14610197578063a98fb355146101aa578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a6146101445780638da5cb5b1461014c575b600080fd5b6100cc6100c736600461117c565b6101d8565b005b6100e16100dc366004611206565b610460565b6040516100ee919061122a565b60405180910390f35b6100cc610105366004611206565b610930565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ee565b6100cc6109ca565b6033546001600160a01b031661012c565b6100cc61016b36600461132c565b6109de565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6100cc6101a5366004611206565b610a74565b6100cc6101b83660046113d7565b610b3b565b6100e1610b8f565b6100cc6101d3366004611206565b610f59565b6101e0610fcf565b60005b818110156103db578282828181106101fd576101fd611428565b905060200281019061020f919061143e565b610220906040810190602001611206565b6001600160a01b03166323b872dd333086868681811061024257610242611428565b9050602002810190610254919061143e565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf919061146e565b508282828181106102e2576102e2611428565b90506020028101906102f4919061143e565b610305906040810190602001611206565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061034657610346611428565b9050602002810190610358919061143e565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca919061146e565b506103d4816114a6565b90506101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042a908590859060040161155a565b600060405180830381600087803b15801561044457600080fd5b505af1158015610458573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f09190611668565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f9190611681565b90506001600160c01b038116158061061957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061491906116aa565b60ff16155b1561063557505060408051600081526020810190915292915050565b6000610649826001600160c01b0316611029565b90506000805b825181101561071f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061069957610699611428565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107019190611668565b61070b90836116cd565b915080610717816114a6565b91505061064f565b5060008167ffffffffffffffff81111561073b5761073b611277565b604051908082528060200260200182016040528015610764578160200160208202803683370190505b5090506000805b845181101561092357600085828151811061078857610788611428565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108219190611668565b905060005b8181101561090d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906116e5565b600001518686815181106108d5576108d5611428565b6001600160a01b0390921660209283029190910190910152846108f7816114a6565b9550508080610905906114a6565b915050610826565b505050808061091b906114a6565b91505061076b565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109c75760405162461bcd60e51b815260206004820152603160248201527f6f6e6c79536572746e5461736b4d616e616765723a206e6f742066726f6d207360448201527032b93a37103a30b9b59036b0b730b3b2b960791b60648201526084015b60405180910390fd5b50565b6109d2610fcf565b6109dc60006110ec565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a265760405162461bcd60e51b81526004016109be90611744565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042a9085908590600401611809565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610abc5760405162461bcd60e51b81526004016109be90611744565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b2057600080fd5b505af1158015610b34573d6000803e3d6000fd5b5050505050565b610b43610fcf565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610b06908490600401611855565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bf1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1591906116aa565b60ff16905080610c3357505060408051600081526020810190915290565b6000805b82811015610ce857604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610ca6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cca9190611668565b610cd490836116cd565b915080610ce0816114a6565b915050610c37565b5060008167ffffffffffffffff811115610d0457610d04611277565b604051908082528060200260200182016040528015610d2d578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db691906116aa565b60ff16811015610f4f57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4e9190611668565b905060005b81811015610f3a576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610ec8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eec91906116e5565b60000151858581518110610f0257610f02611428565b6001600160a01b039092166020928302919091019091015283610f24816114a6565b9450508080610f32906114a6565b915050610e53565b50508080610f47906114a6565b915050610d34565b5090949350505050565b610f61610fcf565b6001600160a01b038116610fc65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109be565b6109c7816110ec565b6033546001600160a01b031633146109dc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109be565b60606000806110378461114b565b61ffff1667ffffffffffffffff81111561105357611053611277565b6040519080825280601f01601f19166020018201604052801561107d576020820181803683370190505b5090506000805b825182108015611095575061010081105b15610f4f576001811b9350858416156110dc578060f81b8383815181106110be576110be611428565b60200101906001600160f81b031916908160001a9053508160010191505b6110e5816114a6565b9050611084565b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b821561117657611160600184611868565b909216918061116e8161187f565b91505061114f565b92915050565b6000806020838503121561118f57600080fd5b823567ffffffffffffffff808211156111a757600080fd5b818501915085601f8301126111bb57600080fd5b8135818111156111ca57600080fd5b8660208260051b85010111156111df57600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109c757600080fd5b60006020828403121561121857600080fd5b8135611223816111f1565b9392505050565b6020808252825182820181905260009190848201906040850190845b8181101561126b5783516001600160a01b031683529284019291840191600101611246565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112b0576112b0611277565b60405290565b600067ffffffffffffffff808411156112d1576112d1611277565b604051601f8501601f19908116603f011681019082821181831017156112f9576112f9611277565b8160405280935085815286868601111561131257600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561133f57600080fd5b823561134a816111f1565b9150602083013567ffffffffffffffff8082111561136757600080fd5b908401906060828703121561137b57600080fd5b61138361128d565b82358281111561139257600080fd5b83019150601f820187136113a557600080fd5b6113b4878335602085016112b6565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113e957600080fd5b813567ffffffffffffffff81111561140057600080fd5b8201601f8101841361141157600080fd5b611420848235602084016112b6565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261145457600080fd5b9190910192915050565b8035611469816111f1565b919050565b60006020828403121561148057600080fd5b8151801515811461122357600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114ba576114ba611490565b5060010190565b6bffffffffffffffffffffffff811681146109c757600080fd5b8183526000602080850194508260005b8581101561153b5781356114fe816111f1565b6001600160a01b0316875281830135611516816114c1565b6bffffffffffffffffffffffff168784015260409687019691909101906001016114eb565b509495945050505050565b803563ffffffff8116811461146957600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561165a57878303603f190184528135368b9003609e1901811261159f57600080fd5b8a0160a0813536839003601e190181126115b857600080fd5b8201803567ffffffffffffffff8111156115d157600080fd5b8060061b36038413156115e357600080fd5b8287526115f5838801828c85016114db565b9250505061160488830161145e565b6001600160a01b03168886015281870135878601526060611626818401611546565b63ffffffff1690860152608061163d838201611546565b63ffffffff16950194909452509285019290850190600101611579565b509098975050505050505050565b60006020828403121561167a57600080fd5b5051919050565b60006020828403121561169357600080fd5b81516001600160c01b038116811461122357600080fd5b6000602082840312156116bc57600080fd5b815160ff8116811461122357600080fd5b600082198211156116e0576116e0611490565b500190565b6000604082840312156116f757600080fd5b6040516040810181811067ffffffffffffffff8211171561171a5761171a611277565b6040528251611728816111f1565b81526020830151611738816114c1565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117e2576020818501810151868301820152016117c6565b818111156117f4576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b038316815260406020820152600082516060604084015261183460a08401826117bc565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061122360208301846117bc565b60008282101561187a5761187a611490565b500390565b600061ffff8083168181141561189757611897611490565b600101939250505056fea2646970667358221220cca3278028fab6801cffdd75763732c328c5fa88f3de7dcb39d130bfb3d099c564736f6c634300080c0033",
}

// ContractSertnServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSertnServiceManagerMetaData.ABI instead.
var ContractSertnServiceManagerABI = ContractSertnServiceManagerMetaData.ABI

// ContractSertnServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSertnServiceManagerMetaData.Bin instead.
var ContractSertnServiceManagerBin = ContractSertnServiceManagerMetaData.Bin

// DeployContractSertnServiceManager deploys a new Ethereum contract, binding an instance of ContractSertnServiceManager to it.
func DeployContractSertnServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _sertnTaskManager common.Address) (common.Address, *types.Transaction, *ContractSertnServiceManager, error) {
	parsed, err := ContractSertnServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSertnServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _sertnTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSertnServiceManager{ContractSertnServiceManagerCaller: ContractSertnServiceManagerCaller{contract: contract}, ContractSertnServiceManagerTransactor: ContractSertnServiceManagerTransactor{contract: contract}, ContractSertnServiceManagerFilterer: ContractSertnServiceManagerFilterer{contract: contract}}, nil
}

// ContractSertnServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractSertnServiceManager struct {
	ContractSertnServiceManagerCaller     // Read-only binding to the contract
	ContractSertnServiceManagerTransactor // Write-only binding to the contract
	ContractSertnServiceManagerFilterer   // Log filterer for contract events
}

// ContractSertnServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSertnServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSertnServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSertnServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSertnServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSertnServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSertnServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSertnServiceManagerSession struct {
	Contract     *ContractSertnServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractSertnServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSertnServiceManagerCallerSession struct {
	Contract *ContractSertnServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractSertnServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSertnServiceManagerTransactorSession struct {
	Contract     *ContractSertnServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractSertnServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSertnServiceManagerRaw struct {
	Contract *ContractSertnServiceManager // Generic contract binding to access the raw methods on
}

// ContractSertnServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSertnServiceManagerCallerRaw struct {
	Contract *ContractSertnServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSertnServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSertnServiceManagerTransactorRaw struct {
	Contract *ContractSertnServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSertnServiceManager creates a new instance of ContractSertnServiceManager, bound to a specific deployed contract.
func NewContractSertnServiceManager(address common.Address, backend bind.ContractBackend) (*ContractSertnServiceManager, error) {
	contract, err := bindContractSertnServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSertnServiceManager{ContractSertnServiceManagerCaller: ContractSertnServiceManagerCaller{contract: contract}, ContractSertnServiceManagerTransactor: ContractSertnServiceManagerTransactor{contract: contract}, ContractSertnServiceManagerFilterer: ContractSertnServiceManagerFilterer{contract: contract}}, nil
}

// NewContractSertnServiceManagerCaller creates a new read-only instance of ContractSertnServiceManager, bound to a specific deployed contract.
func NewContractSertnServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractSertnServiceManagerCaller, error) {
	contract, err := bindContractSertnServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSertnServiceManagerCaller{contract: contract}, nil
}

// NewContractSertnServiceManagerTransactor creates a new write-only instance of ContractSertnServiceManager, bound to a specific deployed contract.
func NewContractSertnServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSertnServiceManagerTransactor, error) {
	contract, err := bindContractSertnServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSertnServiceManagerTransactor{contract: contract}, nil
}

// NewContractSertnServiceManagerFilterer creates a new log filterer instance of ContractSertnServiceManager, bound to a specific deployed contract.
func NewContractSertnServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSertnServiceManagerFilterer, error) {
	contract, err := bindContractSertnServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSertnServiceManagerFilterer{contract: contract}, nil
}

// bindContractSertnServiceManager binds a generic wrapper to an already deployed contract.
func bindContractSertnServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSertnServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSertnServiceManager *ContractSertnServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSertnServiceManager.Contract.ContractSertnServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSertnServiceManager *ContractSertnServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.ContractSertnServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSertnServiceManager *ContractSertnServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.ContractSertnServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSertnServiceManager *ContractSertnServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSertnServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractSertnServiceManager.Contract.AvsDirectory(&_ContractSertnServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractSertnServiceManager.Contract.AvsDirectory(&_ContractSertnServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractSertnServiceManager *ContractSertnServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractSertnServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractSertnServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractSertnServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractSertnServiceManager *ContractSertnServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractSertnServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractSertnServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractSertnServiceManager *ContractSertnServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractSertnServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractSertnServiceManager.Contract.GetRestakeableStrategies(&_ContractSertnServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractSertnServiceManager *ContractSertnServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractSertnServiceManager.Contract.GetRestakeableStrategies(&_ContractSertnServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) Owner() (common.Address, error) {
	return _ContractSertnServiceManager.Contract.Owner(&_ContractSertnServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractSertnServiceManager.Contract.Owner(&_ContractSertnServiceManager.CallOpts)
}

// SertnTaskManager is a free data retrieval call binding the contract method 0xa2bc0400.
//
// Solidity: function sertnTaskManager() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerCaller) SertnTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSertnServiceManager.contract.Call(opts, &out, "sertnTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SertnTaskManager is a free data retrieval call binding the contract method 0xa2bc0400.
//
// Solidity: function sertnTaskManager() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) SertnTaskManager() (common.Address, error) {
	return _ContractSertnServiceManager.Contract.SertnTaskManager(&_ContractSertnServiceManager.CallOpts)
}

// SertnTaskManager is a free data retrieval call binding the contract method 0xa2bc0400.
//
// Solidity: function sertnTaskManager() view returns(address)
func (_ContractSertnServiceManager *ContractSertnServiceManagerCallerSession) SertnTaskManager() (common.Address, error) {
	return _ContractSertnServiceManager.Contract.SertnTaskManager(&_ContractSertnServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractSertnServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractSertnServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.FreezeOperator(&_ContractSertnServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.FreezeOperator(&_ContractSertnServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractSertnServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.PayForRange(&_ContractSertnServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.PayForRange(&_ContractSertnServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSertnServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.RegisterOperatorToAVS(&_ContractSertnServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.RegisterOperatorToAVS(&_ContractSertnServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSertnServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.RenounceOwnership(&_ContractSertnServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.RenounceOwnership(&_ContractSertnServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.TransferOwnership(&_ContractSertnServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.TransferOwnership(&_ContractSertnServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractSertnServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.UpdateAVSMetadataURI(&_ContractSertnServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractSertnServiceManager *ContractSertnServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractSertnServiceManager.Contract.UpdateAVSMetadataURI(&_ContractSertnServiceManager.TransactOpts, _metadataURI)
}

// ContractSertnServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSertnServiceManager contract.
type ContractSertnServiceManagerInitializedIterator struct {
	Event *ContractSertnServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractSertnServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnServiceManagerInitialized)
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
		it.Event = new(ContractSertnServiceManagerInitialized)
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
func (it *ContractSertnServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnServiceManagerInitialized represents a Initialized event raised by the ContractSertnServiceManager contract.
type ContractSertnServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSertnServiceManager *ContractSertnServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSertnServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractSertnServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSertnServiceManagerInitializedIterator{contract: _ContractSertnServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSertnServiceManager *ContractSertnServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSertnServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSertnServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnServiceManagerInitialized)
				if err := _ContractSertnServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractSertnServiceManager *ContractSertnServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractSertnServiceManagerInitialized, error) {
	event := new(ContractSertnServiceManagerInitialized)
	if err := _ContractSertnServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSertnServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSertnServiceManager contract.
type ContractSertnServiceManagerOwnershipTransferredIterator struct {
	Event *ContractSertnServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractSertnServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSertnServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractSertnServiceManagerOwnershipTransferred)
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
func (it *ContractSertnServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSertnServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSertnServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSertnServiceManager contract.
type ContractSertnServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSertnServiceManager *ContractSertnServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSertnServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSertnServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSertnServiceManagerOwnershipTransferredIterator{contract: _ContractSertnServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSertnServiceManager *ContractSertnServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSertnServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSertnServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSertnServiceManagerOwnershipTransferred)
				if err := _ContractSertnServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractSertnServiceManager *ContractSertnServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSertnServiceManagerOwnershipTransferred, error) {
	event := new(ContractSertnServiceManagerOwnershipTransferred)
	if err := _ContractSertnServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
