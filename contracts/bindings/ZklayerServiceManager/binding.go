// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractZklayerServiceManager

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

// ContractZklayerServiceManagerMetaData contains all meta data concerning the ContractZklayerServiceManager contract.
var ContractZklayerServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_zklayerTaskManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zklayerTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIZklayerTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b4338038062001b4383398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118cd620002766000396000818160be015261093b01526000818161065a015281816107b60152818161084d01528181610c5b01528181610ddf0152610e7e0152600081816104850152818161051401528181610594015281816109ed01528181610a8301528181610b990152610d3a01526000818161031501526103f301526000818161014701528181610a4101528181610adf0152610b5e01526118cd6000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638da5cb5b116100715780638da5cb5b146101735780639926ee7d14610184578063a364f4da14610197578063a98fb355146101aa578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80630584bfb6146100b95780631b445516146100fd57806333cfb7b71461011257806338c8ee64146101325780636b3aa72e14610145578063715018a61461016b575b600080fd5b6100e07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61011061010b366004611173565b6101d8565b005b6101256101203660046111fd565b610460565b6040516100f49190611221565b6101106101403660046111fd565b610930565b7f00000000000000000000000000000000000000000000000000000000000000006100e0565b6101106109ce565b6033546001600160a01b03166100e0565b610110610192366004611323565b6109e2565b6101106101a53660046111fd565b610a78565b6101106101b83660046113ce565b610b3f565b610125610b93565b6101106101d33660046111fd565b610f5d565b6101e0610fd3565b60005b818110156103db578282828181106101fd576101fd61141f565b905060200281019061020f9190611435565b6102209060408101906020016111fd565b6001600160a01b03166323b872dd33308686868181106102425761024261141f565b90506020028101906102549190611435565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf9190611465565b508282828181106102e2576102e261141f565b90506020028101906102f49190611435565b6103059060408101906020016111fd565b6001600160a01b031663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008585858181106103465761034661141f565b90506020028101906103589190611435565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca9190611465565b506103d48161149d565b90506101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042a9085908590600401611551565b600060405180830381600087803b15801561044457600080fd5b505af1158015610458573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f0919061165f565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f9190611678565b90506001600160c01b038116158061061957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061491906116a1565b60ff16155b1561063557505060408051600081526020810190915292915050565b6000610649826001600160c01b031661102d565b90506000805b825181101561071f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106106995761069961141f565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610701919061165f565b61070b90836116c4565b9150806107178161149d565b91505061064f565b5060008167ffffffffffffffff81111561073b5761073b61126e565b604051908082528060200260200182016040528015610764578160200160208202803683370190505b5090506000805b84518110156109235760008582815181106107885761078861141f565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610821919061165f565b905060005b8181101561090d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906116dc565b600001518686815181106108d5576108d561141f565b6001600160a01b0390921660209283029190910190910152846108f78161149d565b95505080806109059061149d565b915050610826565b505050808061091b9061149d565b91505061076b565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109cb5760405162461bcd60e51b815260206004820152603560248201527f6f6e6c795a6b6c617965725461736b4d616e616765723a206e6f742066726f6d604482015274103d35b630bcb2b9103a30b9b59036b0b730b3b2b960591b60648201526084015b60405180910390fd5b50565b6109d6610fd3565b6109e060006110f0565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a2a5760405162461bcd60e51b81526004016109c29061173b565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042a9085908590600401611800565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ac05760405162461bcd60e51b81526004016109c29061173b565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b2457600080fd5b505af1158015610b38573d6000803e3d6000fd5b5050505050565b610b47610fd3565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610b0a90849060040161184b565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bf5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1991906116a1565b60ff16905080610c3757505060408051600081526020810190915290565b6000805b82811015610cec57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610caa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cce919061165f565b610cd890836116c4565b915080610ce48161149d565b915050610c3b565b5060008167ffffffffffffffff811115610d0857610d0861126e565b604051908082528060200260200182016040528015610d31578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dba91906116a1565b60ff16811015610f5357604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e52919061165f565b905060005b81811015610f3e576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610ecc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef091906116dc565b60000151858581518110610f0657610f0661141f565b6001600160a01b039092166020928302919091019091015283610f288161149d565b9450508080610f369061149d565b915050610e57565b50508080610f4b9061149d565b915050610d38565b5090949350505050565b610f65610fd3565b6001600160a01b038116610fca5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109c2565b6109cb816110f0565b6033546001600160a01b031633146109e05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109c2565b606060008061103b84611142565b61ffff1667ffffffffffffffff8111156110575761105761126e565b6040519080825280601f01601f191660200182016040528015611081576020820181803683370190505b5090506000805b825182108015611099575061010081105b15610f53576001811b9350858416156110e0578060f81b8383815181106110c2576110c261141f565b60200101906001600160f81b031916908160001a9053508160010191505b6110e98161149d565b9050611088565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b821561116d5761115760018461185e565b909216918061116581611875565b915050611146565b92915050565b6000806020838503121561118657600080fd5b823567ffffffffffffffff8082111561119e57600080fd5b818501915085601f8301126111b257600080fd5b8135818111156111c157600080fd5b8660208260051b85010111156111d657600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109cb57600080fd5b60006020828403121561120f57600080fd5b813561121a816111e8565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156112625783516001600160a01b03168352928401929184019160010161123d565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112a7576112a761126e565b60405290565b600067ffffffffffffffff808411156112c8576112c861126e565b604051601f8501601f19908116603f011681019082821181831017156112f0576112f061126e565b8160405280935085815286868601111561130957600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561133657600080fd5b8235611341816111e8565b9150602083013567ffffffffffffffff8082111561135e57600080fd5b908401906060828703121561137257600080fd5b61137a611284565b82358281111561138957600080fd5b83019150601f8201871361139c57600080fd5b6113ab878335602085016112ad565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113e057600080fd5b813567ffffffffffffffff8111156113f757600080fd5b8201601f8101841361140857600080fd5b611417848235602084016112ad565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261144b57600080fd5b9190910192915050565b8035611460816111e8565b919050565b60006020828403121561147757600080fd5b8151801515811461121a57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114b1576114b1611487565b5060010190565b6bffffffffffffffffffffffff811681146109cb57600080fd5b8183526000602080850194508260005b858110156115325781356114f5816111e8565b6001600160a01b031687528183013561150d816114b8565b6bffffffffffffffffffffffff168784015260409687019691909101906001016114e2565b509495945050505050565b803563ffffffff8116811461146057600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561165157878303603f190184528135368b9003609e1901811261159657600080fd5b8a0160a0813536839003601e190181126115af57600080fd5b8201803567ffffffffffffffff8111156115c857600080fd5b8060061b36038413156115da57600080fd5b8287526115ec838801828c85016114d2565b925050506115fb888301611455565b6001600160a01b0316888601528187013587860152606061161d81840161153d565b63ffffffff1690860152608061163483820161153d565b63ffffffff16950194909452509285019290850190600101611570565b509098975050505050505050565b60006020828403121561167157600080fd5b5051919050565b60006020828403121561168a57600080fd5b81516001600160c01b038116811461121a57600080fd5b6000602082840312156116b357600080fd5b815160ff8116811461121a57600080fd5b600082198211156116d7576116d7611487565b500190565b6000604082840312156116ee57600080fd5b6040516040810181811067ffffffffffffffff821117156117115761171161126e565b604052825161171f816111e8565b8152602083015161172f816114b8565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117d9576020818501810151868301820152016117bd565b818111156117eb576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261182a60a08401826117b3565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061121a60208301846117b3565b60008282101561187057611870611487565b500390565b600061ffff8083168181141561188d5761188d611487565b600101939250505056fea26469706673582212200c9bedc26b61f43431385e999a0539999f4c39f8eb6e652ce3d50e562e26c38664736f6c634300080c0033",
}

// ContractZklayerServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractZklayerServiceManagerMetaData.ABI instead.
var ContractZklayerServiceManagerABI = ContractZklayerServiceManagerMetaData.ABI

// ContractZklayerServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractZklayerServiceManagerMetaData.Bin instead.
var ContractZklayerServiceManagerBin = ContractZklayerServiceManagerMetaData.Bin

// DeployContractZklayerServiceManager deploys a new Ethereum contract, binding an instance of ContractZklayerServiceManager to it.
func DeployContractZklayerServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _zklayerTaskManager common.Address) (common.Address, *types.Transaction, *ContractZklayerServiceManager, error) {
	parsed, err := ContractZklayerServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractZklayerServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _zklayerTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractZklayerServiceManager{ContractZklayerServiceManagerCaller: ContractZklayerServiceManagerCaller{contract: contract}, ContractZklayerServiceManagerTransactor: ContractZklayerServiceManagerTransactor{contract: contract}, ContractZklayerServiceManagerFilterer: ContractZklayerServiceManagerFilterer{contract: contract}}, nil
}

// ContractZklayerServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractZklayerServiceManager struct {
	ContractZklayerServiceManagerCaller     // Read-only binding to the contract
	ContractZklayerServiceManagerTransactor // Write-only binding to the contract
	ContractZklayerServiceManagerFilterer   // Log filterer for contract events
}

// ContractZklayerServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractZklayerServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZklayerServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractZklayerServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZklayerServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractZklayerServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractZklayerServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractZklayerServiceManagerSession struct {
	Contract     *ContractZklayerServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractZklayerServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractZklayerServiceManagerCallerSession struct {
	Contract *ContractZklayerServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractZklayerServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractZklayerServiceManagerTransactorSession struct {
	Contract     *ContractZklayerServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractZklayerServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractZklayerServiceManagerRaw struct {
	Contract *ContractZklayerServiceManager // Generic contract binding to access the raw methods on
}

// ContractZklayerServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractZklayerServiceManagerCallerRaw struct {
	Contract *ContractZklayerServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractZklayerServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractZklayerServiceManagerTransactorRaw struct {
	Contract *ContractZklayerServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractZklayerServiceManager creates a new instance of ContractZklayerServiceManager, bound to a specific deployed contract.
func NewContractZklayerServiceManager(address common.Address, backend bind.ContractBackend) (*ContractZklayerServiceManager, error) {
	contract, err := bindContractZklayerServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerServiceManager{ContractZklayerServiceManagerCaller: ContractZklayerServiceManagerCaller{contract: contract}, ContractZklayerServiceManagerTransactor: ContractZklayerServiceManagerTransactor{contract: contract}, ContractZklayerServiceManagerFilterer: ContractZklayerServiceManagerFilterer{contract: contract}}, nil
}

// NewContractZklayerServiceManagerCaller creates a new read-only instance of ContractZklayerServiceManager, bound to a specific deployed contract.
func NewContractZklayerServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractZklayerServiceManagerCaller, error) {
	contract, err := bindContractZklayerServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerServiceManagerCaller{contract: contract}, nil
}

// NewContractZklayerServiceManagerTransactor creates a new write-only instance of ContractZklayerServiceManager, bound to a specific deployed contract.
func NewContractZklayerServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractZklayerServiceManagerTransactor, error) {
	contract, err := bindContractZklayerServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerServiceManagerTransactor{contract: contract}, nil
}

// NewContractZklayerServiceManagerFilterer creates a new log filterer instance of ContractZklayerServiceManager, bound to a specific deployed contract.
func NewContractZklayerServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractZklayerServiceManagerFilterer, error) {
	contract, err := bindContractZklayerServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerServiceManagerFilterer{contract: contract}, nil
}

// bindContractZklayerServiceManager binds a generic wrapper to an already deployed contract.
func bindContractZklayerServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractZklayerServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZklayerServiceManager.Contract.ContractZklayerServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.ContractZklayerServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.ContractZklayerServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractZklayerServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractZklayerServiceManager.Contract.AvsDirectory(&_ContractZklayerServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractZklayerServiceManager.Contract.AvsDirectory(&_ContractZklayerServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZklayerServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractZklayerServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractZklayerServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractZklayerServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractZklayerServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractZklayerServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractZklayerServiceManager.Contract.GetRestakeableStrategies(&_ContractZklayerServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractZklayerServiceManager.Contract.GetRestakeableStrategies(&_ContractZklayerServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) Owner() (common.Address, error) {
	return _ContractZklayerServiceManager.Contract.Owner(&_ContractZklayerServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractZklayerServiceManager.Contract.Owner(&_ContractZklayerServiceManager.CallOpts)
}

// ZklayerTaskManager is a free data retrieval call binding the contract method 0x0584bfb6.
//
// Solidity: function zklayerTaskManager() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCaller) ZklayerTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractZklayerServiceManager.contract.Call(opts, &out, "zklayerTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZklayerTaskManager is a free data retrieval call binding the contract method 0x0584bfb6.
//
// Solidity: function zklayerTaskManager() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) ZklayerTaskManager() (common.Address, error) {
	return _ContractZklayerServiceManager.Contract.ZklayerTaskManager(&_ContractZklayerServiceManager.CallOpts)
}

// ZklayerTaskManager is a free data retrieval call binding the contract method 0x0584bfb6.
//
// Solidity: function zklayerTaskManager() view returns(address)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerCallerSession) ZklayerTaskManager() (common.Address, error) {
	return _ContractZklayerServiceManager.Contract.ZklayerTaskManager(&_ContractZklayerServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractZklayerServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractZklayerServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.FreezeOperator(&_ContractZklayerServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.FreezeOperator(&_ContractZklayerServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.PayForRange(&_ContractZklayerServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.PayForRange(&_ContractZklayerServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.RegisterOperatorToAVS(&_ContractZklayerServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.RegisterOperatorToAVS(&_ContractZklayerServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.RenounceOwnership(&_ContractZklayerServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.RenounceOwnership(&_ContractZklayerServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.TransferOwnership(&_ContractZklayerServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.TransferOwnership(&_ContractZklayerServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.UpdateAVSMetadataURI(&_ContractZklayerServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractZklayerServiceManager.Contract.UpdateAVSMetadataURI(&_ContractZklayerServiceManager.TransactOpts, _metadataURI)
}

// ContractZklayerServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractZklayerServiceManager contract.
type ContractZklayerServiceManagerInitializedIterator struct {
	Event *ContractZklayerServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractZklayerServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerServiceManagerInitialized)
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
		it.Event = new(ContractZklayerServiceManagerInitialized)
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
func (it *ContractZklayerServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerServiceManagerInitialized represents a Initialized event raised by the ContractZklayerServiceManager contract.
type ContractZklayerServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractZklayerServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractZklayerServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractZklayerServiceManagerInitializedIterator{contract: _ContractZklayerServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractZklayerServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractZklayerServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerServiceManagerInitialized)
				if err := _ContractZklayerServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractZklayerServiceManagerInitialized, error) {
	event := new(ContractZklayerServiceManagerInitialized)
	if err := _ContractZklayerServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractZklayerServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractZklayerServiceManager contract.
type ContractZklayerServiceManagerOwnershipTransferredIterator struct {
	Event *ContractZklayerServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractZklayerServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractZklayerServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractZklayerServiceManagerOwnershipTransferred)
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
func (it *ContractZklayerServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractZklayerServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractZklayerServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractZklayerServiceManager contract.
type ContractZklayerServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractZklayerServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZklayerServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractZklayerServiceManagerOwnershipTransferredIterator{contract: _ContractZklayerServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractZklayerServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractZklayerServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractZklayerServiceManagerOwnershipTransferred)
				if err := _ContractZklayerServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractZklayerServiceManager *ContractZklayerServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractZklayerServiceManagerOwnershipTransferred, error) {
	event := new(ContractZklayerServiceManagerOwnershipTransferred)
	if err := _ContractZklayerServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
