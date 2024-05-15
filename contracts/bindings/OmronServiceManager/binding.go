// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOmronServiceManager

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

// ContractOmronServiceManagerMetaData contains all meta data concerning the ContractOmronServiceManager contract.
var ContractOmronServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_omronTaskManager\",\"type\":\"address\",\"internalType\":\"contractIOmronTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omronTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmronTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b4c38038062001b4c83398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118d5620002776000396000818161019b015261093b01526000818161065a015281816107b60152818161084d01528181610c6301528181610de70152610e860152600081816104850152818161051401528181610594015281816109f501528181610a8b01528181610ba10152610d4201526000818161031501526103f301526000818161010c01528181610a4901528181610ae70152610b6601526118d56000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639926ee7d116100715780639926ee7d1461015d578063a364f4da14610170578063a98fb35514610183578063cffd9add14610196578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a6146101445780638da5cb5b1461014c575b600080fd5b6100cc6100c736600461117b565b6101d8565b005b6100e16100dc366004611205565b610460565b6040516100ee9190611229565b60405180910390f35b6100cc610105366004611205565b610930565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ee565b6100cc6109d6565b6033546001600160a01b031661012c565b6100cc61016b36600461132b565b6109ea565b6100cc61017e366004611205565b610a80565b6100cc6101913660046113d6565b610b47565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6100e1610b9b565b6100cc6101d3366004611205565b610f65565b6101e0610fdb565b60005b818110156103db578282828181106101fd576101fd611427565b905060200281019061020f919061143d565b610220906040810190602001611205565b6001600160a01b03166323b872dd333086868681811061024257610242611427565b9050602002810190610254919061143d565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf919061146d565b508282828181106102e2576102e2611427565b90506020028101906102f4919061143d565b610305906040810190602001611205565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061034657610346611427565b9050602002810190610358919061143d565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca919061146d565b506103d4816114a5565b90506101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042a9085908590600401611559565b600060405180830381600087803b15801561044457600080fd5b505af1158015610458573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f09190611667565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f9190611680565b90506001600160c01b038116158061061957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061491906116a9565b60ff16155b1561063557505060408051600081526020810190915292915050565b6000610649826001600160c01b0316611035565b90506000805b825181101561071f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061069957610699611427565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107019190611667565b61070b90836116cc565b915080610717816114a5565b91505061064f565b5060008167ffffffffffffffff81111561073b5761073b611276565b604051908082528060200260200182016040528015610764578160200160208202803683370190505b5090506000805b845181101561092357600085828151811061078857610788611427565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108219190611667565b905060005b8181101561090d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906116e4565b600001518686815181106108d5576108d5611427565b6001600160a01b0390921660209283029190910190910152846108f7816114a5565b9550508080610905906114a5565b915050610826565b505050808061091b906114a5565b91505061076b565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109d35760405162461bcd60e51b815260206004820152603d60248201527f6f6e6c794f6d726f6e5461736b4d616e616765723a206e6f742066726f6d206360448201527f72656469626c65207371756172696e67207461736b206d616e6167657200000060648201526084015b60405180910390fd5b50565b6109de610fdb565b6109e860006110f8565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a325760405162461bcd60e51b81526004016109ca90611743565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042a9085908590600401611808565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ac85760405162461bcd60e51b81526004016109ca90611743565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b2c57600080fd5b505af1158015610b40573d6000803e3d6000fd5b5050505050565b610b4f610fdb565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610b12908490600401611853565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bfd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2191906116a9565b60ff16905080610c3f57505060408051600081526020810190915290565b6000805b82811015610cf457604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd69190611667565b610ce090836116cc565b915080610cec816114a5565b915050610c43565b5060008167ffffffffffffffff811115610d1057610d10611276565b604051908082528060200260200182016040528015610d39578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc291906116a9565b60ff16811015610f5b57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5a9190611667565b905060005b81811015610f46576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610ed4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef891906116e4565b60000151858581518110610f0e57610f0e611427565b6001600160a01b039092166020928302919091019091015283610f30816114a5565b9450508080610f3e906114a5565b915050610e5f565b50508080610f53906114a5565b915050610d40565b5090949350505050565b610f6d610fdb565b6001600160a01b038116610fd25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109ca565b6109d3816110f8565b6033546001600160a01b031633146109e85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109ca565b60606000806110438461114a565b61ffff1667ffffffffffffffff81111561105f5761105f611276565b6040519080825280601f01601f191660200182016040528015611089576020820181803683370190505b5090506000805b8251821080156110a1575061010081105b15610f5b576001811b9350858416156110e8578060f81b8383815181106110ca576110ca611427565b60200101906001600160f81b031916908160001a9053508160010191505b6110f1816114a5565b9050611090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b82156111755761115f600184611866565b909216918061116d8161187d565b91505061114e565b92915050565b6000806020838503121561118e57600080fd5b823567ffffffffffffffff808211156111a657600080fd5b818501915085601f8301126111ba57600080fd5b8135818111156111c957600080fd5b8660208260051b85010111156111de57600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109d357600080fd5b60006020828403121561121757600080fd5b8135611222816111f0565b9392505050565b6020808252825182820181905260009190848201906040850190845b8181101561126a5783516001600160a01b031683529284019291840191600101611245565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112af576112af611276565b60405290565b600067ffffffffffffffff808411156112d0576112d0611276565b604051601f8501601f19908116603f011681019082821181831017156112f8576112f8611276565b8160405280935085815286868601111561131157600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561133e57600080fd5b8235611349816111f0565b9150602083013567ffffffffffffffff8082111561136657600080fd5b908401906060828703121561137a57600080fd5b61138261128c565b82358281111561139157600080fd5b83019150601f820187136113a457600080fd5b6113b3878335602085016112b5565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113e857600080fd5b813567ffffffffffffffff8111156113ff57600080fd5b8201601f8101841361141057600080fd5b61141f848235602084016112b5565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261145357600080fd5b9190910192915050565b8035611468816111f0565b919050565b60006020828403121561147f57600080fd5b8151801515811461122257600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114b9576114b961148f565b5060010190565b6bffffffffffffffffffffffff811681146109d357600080fd5b8183526000602080850194508260005b8581101561153a5781356114fd816111f0565b6001600160a01b0316875281830135611515816114c0565b6bffffffffffffffffffffffff168784015260409687019691909101906001016114ea565b509495945050505050565b803563ffffffff8116811461146857600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561165957878303603f190184528135368b9003609e1901811261159e57600080fd5b8a0160a0813536839003601e190181126115b757600080fd5b8201803567ffffffffffffffff8111156115d057600080fd5b8060061b36038413156115e257600080fd5b8287526115f4838801828c85016114da565b9250505061160388830161145d565b6001600160a01b03168886015281870135878601526060611625818401611545565b63ffffffff1690860152608061163c838201611545565b63ffffffff16950194909452509285019290850190600101611578565b509098975050505050505050565b60006020828403121561167957600080fd5b5051919050565b60006020828403121561169257600080fd5b81516001600160c01b038116811461122257600080fd5b6000602082840312156116bb57600080fd5b815160ff8116811461122257600080fd5b600082198211156116df576116df61148f565b500190565b6000604082840312156116f657600080fd5b6040516040810181811067ffffffffffffffff8211171561171957611719611276565b6040528251611727816111f0565b81526020830151611737816114c0565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117e1576020818501810151868301820152016117c5565b818111156117f3576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261183260a08401826117bb565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061122260208301846117bb565b6000828210156118785761187861148f565b500390565b600061ffff808316818114156118955761189561148f565b600101939250505056fea26469706673582212205a3106c58ae42cd1ffc31cc4fffb6dbd9b31b21c3402cdb0809b7a766264387064736f6c634300080c0033",
}

// ContractOmronServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOmronServiceManagerMetaData.ABI instead.
var ContractOmronServiceManagerABI = ContractOmronServiceManagerMetaData.ABI

// ContractOmronServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOmronServiceManagerMetaData.Bin instead.
var ContractOmronServiceManagerBin = ContractOmronServiceManagerMetaData.Bin

// DeployContractOmronServiceManager deploys a new Ethereum contract, binding an instance of ContractOmronServiceManager to it.
func DeployContractOmronServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _omronTaskManager common.Address) (common.Address, *types.Transaction, *ContractOmronServiceManager, error) {
	parsed, err := ContractOmronServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOmronServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _omronTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOmronServiceManager{ContractOmronServiceManagerCaller: ContractOmronServiceManagerCaller{contract: contract}, ContractOmronServiceManagerTransactor: ContractOmronServiceManagerTransactor{contract: contract}, ContractOmronServiceManagerFilterer: ContractOmronServiceManagerFilterer{contract: contract}}, nil
}

// ContractOmronServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractOmronServiceManager struct {
	ContractOmronServiceManagerCaller     // Read-only binding to the contract
	ContractOmronServiceManagerTransactor // Write-only binding to the contract
	ContractOmronServiceManagerFilterer   // Log filterer for contract events
}

// ContractOmronServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOmronServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOmronServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOmronServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOmronServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOmronServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOmronServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOmronServiceManagerSession struct {
	Contract     *ContractOmronServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractOmronServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOmronServiceManagerCallerSession struct {
	Contract *ContractOmronServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractOmronServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOmronServiceManagerTransactorSession struct {
	Contract     *ContractOmronServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractOmronServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOmronServiceManagerRaw struct {
	Contract *ContractOmronServiceManager // Generic contract binding to access the raw methods on
}

// ContractOmronServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOmronServiceManagerCallerRaw struct {
	Contract *ContractOmronServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOmronServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOmronServiceManagerTransactorRaw struct {
	Contract *ContractOmronServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOmronServiceManager creates a new instance of ContractOmronServiceManager, bound to a specific deployed contract.
func NewContractOmronServiceManager(address common.Address, backend bind.ContractBackend) (*ContractOmronServiceManager, error) {
	contract, err := bindContractOmronServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOmronServiceManager{ContractOmronServiceManagerCaller: ContractOmronServiceManagerCaller{contract: contract}, ContractOmronServiceManagerTransactor: ContractOmronServiceManagerTransactor{contract: contract}, ContractOmronServiceManagerFilterer: ContractOmronServiceManagerFilterer{contract: contract}}, nil
}

// NewContractOmronServiceManagerCaller creates a new read-only instance of ContractOmronServiceManager, bound to a specific deployed contract.
func NewContractOmronServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractOmronServiceManagerCaller, error) {
	contract, err := bindContractOmronServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOmronServiceManagerCaller{contract: contract}, nil
}

// NewContractOmronServiceManagerTransactor creates a new write-only instance of ContractOmronServiceManager, bound to a specific deployed contract.
func NewContractOmronServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOmronServiceManagerTransactor, error) {
	contract, err := bindContractOmronServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOmronServiceManagerTransactor{contract: contract}, nil
}

// NewContractOmronServiceManagerFilterer creates a new log filterer instance of ContractOmronServiceManager, bound to a specific deployed contract.
func NewContractOmronServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOmronServiceManagerFilterer, error) {
	contract, err := bindContractOmronServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOmronServiceManagerFilterer{contract: contract}, nil
}

// bindContractOmronServiceManager binds a generic wrapper to an already deployed contract.
func bindContractOmronServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOmronServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOmronServiceManager *ContractOmronServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOmronServiceManager.Contract.ContractOmronServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOmronServiceManager *ContractOmronServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.ContractOmronServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOmronServiceManager *ContractOmronServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.ContractOmronServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOmronServiceManager *ContractOmronServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOmronServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractOmronServiceManager.Contract.AvsDirectory(&_ContractOmronServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractOmronServiceManager.Contract.AvsDirectory(&_ContractOmronServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractOmronServiceManager *ContractOmronServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOmronServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractOmronServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractOmronServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractOmronServiceManager *ContractOmronServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractOmronServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractOmronServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractOmronServiceManager *ContractOmronServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOmronServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractOmronServiceManager.Contract.GetRestakeableStrategies(&_ContractOmronServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractOmronServiceManager *ContractOmronServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractOmronServiceManager.Contract.GetRestakeableStrategies(&_ContractOmronServiceManager.CallOpts)
}

// OmronTaskManager is a free data retrieval call binding the contract method 0xcffd9add.
//
// Solidity: function omronTaskManager() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerCaller) OmronTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronServiceManager.contract.Call(opts, &out, "omronTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OmronTaskManager is a free data retrieval call binding the contract method 0xcffd9add.
//
// Solidity: function omronTaskManager() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) OmronTaskManager() (common.Address, error) {
	return _ContractOmronServiceManager.Contract.OmronTaskManager(&_ContractOmronServiceManager.CallOpts)
}

// OmronTaskManager is a free data retrieval call binding the contract method 0xcffd9add.
//
// Solidity: function omronTaskManager() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerCallerSession) OmronTaskManager() (common.Address, error) {
	return _ContractOmronServiceManager.Contract.OmronTaskManager(&_ContractOmronServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractOmronServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) Owner() (common.Address, error) {
	return _ContractOmronServiceManager.Contract.Owner(&_ContractOmronServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractOmronServiceManager *ContractOmronServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractOmronServiceManager.Contract.Owner(&_ContractOmronServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractOmronServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractOmronServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.FreezeOperator(&_ContractOmronServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.FreezeOperator(&_ContractOmronServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractOmronServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.PayForRange(&_ContractOmronServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.PayForRange(&_ContractOmronServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractOmronServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.RegisterOperatorToAVS(&_ContractOmronServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.RegisterOperatorToAVS(&_ContractOmronServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOmronServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.RenounceOwnership(&_ContractOmronServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.RenounceOwnership(&_ContractOmronServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.TransferOwnership(&_ContractOmronServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.TransferOwnership(&_ContractOmronServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractOmronServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.UpdateAVSMetadataURI(&_ContractOmronServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractOmronServiceManager *ContractOmronServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractOmronServiceManager.Contract.UpdateAVSMetadataURI(&_ContractOmronServiceManager.TransactOpts, _metadataURI)
}

// ContractOmronServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractOmronServiceManager contract.
type ContractOmronServiceManagerInitializedIterator struct {
	Event *ContractOmronServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractOmronServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronServiceManagerInitialized)
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
		it.Event = new(ContractOmronServiceManagerInitialized)
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
func (it *ContractOmronServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronServiceManagerInitialized represents a Initialized event raised by the ContractOmronServiceManager contract.
type ContractOmronServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOmronServiceManager *ContractOmronServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractOmronServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractOmronServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractOmronServiceManagerInitializedIterator{contract: _ContractOmronServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractOmronServiceManager *ContractOmronServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractOmronServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractOmronServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronServiceManagerInitialized)
				if err := _ContractOmronServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractOmronServiceManager *ContractOmronServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractOmronServiceManagerInitialized, error) {
	event := new(ContractOmronServiceManagerInitialized)
	if err := _ContractOmronServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOmronServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractOmronServiceManager contract.
type ContractOmronServiceManagerOwnershipTransferredIterator struct {
	Event *ContractOmronServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOmronServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOmronServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractOmronServiceManagerOwnershipTransferred)
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
func (it *ContractOmronServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOmronServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOmronServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractOmronServiceManager contract.
type ContractOmronServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOmronServiceManager *ContractOmronServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOmronServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOmronServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOmronServiceManagerOwnershipTransferredIterator{contract: _ContractOmronServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractOmronServiceManager *ContractOmronServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOmronServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractOmronServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOmronServiceManagerOwnershipTransferred)
				if err := _ContractOmronServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractOmronServiceManager *ContractOmronServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOmronServiceManagerOwnershipTransferred, error) {
	event := new(ContractOmronServiceManagerOwnershipTransferred)
	if err := _ContractOmronServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
