// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bls

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BLSABI is the input ABI used to generate the binding from.
const BLSABI = "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_com\",\"type\":\"bytes[]\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_b\",\"type\":\"bytes\"}],\"name\":\"equal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vk\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_rnd\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_commit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_psi\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_y\",\"type\":\"bytes32\"}],\"name\":\"verifyKZG\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BLSFuncSigs maps the 4-byte function signature to its string representation.
var BLSFuncSigs = map[string]string{
	"b2bd8933": "add(bytes[])",
	"235266d2": "equal(bytes,bytes)",
	"d30aa507": "verifyKZG(bytes,bytes32,bytes,bytes,bytes32)",
}

// BLSBin is the compiled bytecode used for deploying new contracts.
var BLSBin = "0x608060405234801561001057600080fd5b50610d24806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063235266d214610046578063b2bd89331461006e578063d30aa5071461008e575b600080fd5b6100596100543660046108e8565b6100a1565b60405190151581526020015b60405180910390f35b61008161007c36600461094c565b610118565b6040516100659190610a33565b61005961009c366004610a66565b6102ad565b6000816040516020016100b49190610b00565b60408051601f19818403018152908290526100d191602001610b00565b60405160208183030381529060405280519060200120836040516020016100f89190610b00565b604051602081830303815290604052805190602001201490505b92915050565b606060008251116101615760405162461bcd60e51b815260206004820152600e60248201526d0d2dcecc2d8d2c840d8cadccee8d60931b60448201526064015b60405180910390fd5b60006101888360008151811061017957610179610b1c565b6020026020010151600061069b565b905060015b83518110156102635760008482815181106101aa576101aa610b1c565b60200260200101515111156102515760006101d085838151811061017957610179610b1c565b6040516312f8d6e160e11b815290915073__$51e475105ae19fccae74d465d2ad61b657$__906325f1adc29061020c9086908590600401610b5b565b608060405180830381865af4158015610229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024d9190610b7e565b9250505b8061025b81610c12565b91505061018d565b50805160208083015160408085015160608087015183519586019690965291840192909252820152608081019190915260a001604051602081830303815290604052915050919050565b6000806102bb87600061069b565b905060006102ca88608061074b565b905060006102da8961017661074b565b905060006102e988600061069b565b905060006102f888600061069b565b604051635f42358760e11b815290915073__$51e475105ae19fccae74d465d2ad61b657$__9063be846b0e906103349088908b90600401610c2b565b608060405180830381865af4158015610351573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103759190610b7e565b604051630eea0f1360e11b815290955073__$51e475105ae19fccae74d465d2ad61b657$__90631dd41e26906103af908890600401610c46565b608060405180830381865af41580156103cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f09190610b7e565b6040516312f8d6e160e11b815290955073__$51e475105ae19fccae74d465d2ad61b657$__906325f1adc29061042c9085908990600401610b5b565b608060405180830381865af4158015610449573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046d9190610b7e565b604051635f42358760e11b815290925073__$51e475105ae19fccae74d465d2ad61b657$__9063be846b0e906104a99084908e90600401610c2b565b608060405180830381865af41580156104c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ea9190610b7e565b6040516312f8d6e160e11b815290955073__$51e475105ae19fccae74d465d2ad61b657$__906325f1adc2906105269085908990600401610b5b565b608060405180830381865af4158015610543573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105679190610b7e565b604051630eea0f1360e11b815290925073__$51e475105ae19fccae74d465d2ad61b657$__90631dd41e26906105a1908490600401610c46565b608060405180830381865af41580156105be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e29190610b7e565b9050600073__$51e475105ae19fccae74d465d2ad61b657$__63357d55dd848785886040518563ffffffff1660e01b81526004016106239493929190610c77565b602060405180830381865af4158015610640573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106649190610caf565b9050600160f81b6001600160f81b031960f883901b1614610686576000610689565b60015b96505050505050505b95945050505050565b6106a36107f4565b60808284516106b29190610cc8565b10156106ef5760405162461bcd60e51b815260206004820152600c60248201526b696e76616c69642073697a6560a01b6044820152606401610158565b6106f76107f4565b60005b6004811015610743576000610710602086610cdb565b94508486015190508083836004811061072b5761072b610b1c565b6020020152508061073b81610c12565b9150506106fa565b509392505050565b610753610812565b6101008284516107639190610cc8565b10156107a05760405162461bcd60e51b815260206004820152600c60248201526b696e76616c69642073697a6560a01b6044820152606401610158565b6107a8610812565b60005b60088110156107435760006107c1602086610cdb565b9450848601519050808383600881106107dc576107dc610b1c565b602002015250806107ec81610c12565b9150506107ab565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561087057610870610831565b604052919050565b600082601f83011261088957600080fd5b813567ffffffffffffffff8111156108a3576108a3610831565b6108b6601f8201601f1916602001610847565b8181528460208386010111156108cb57600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156108fb57600080fd5b823567ffffffffffffffff8082111561091357600080fd5b61091f86838701610878565b9350602085013591508082111561093557600080fd5b5061094285828601610878565b9150509250929050565b6000602080838503121561095f57600080fd5b823567ffffffffffffffff8082111561097757600080fd5b818501915085601f83011261098b57600080fd5b81358181111561099d5761099d610831565b8060051b6109ac858201610847565b91825283810185019185810190898411156109c657600080fd5b86860192505b83831015610a02578235858111156109e45760008081fd5b6109f28b89838a0101610878565b83525091860191908601906109cc565b9998505050505050505050565b60005b83811015610a2a578181015183820152602001610a12565b50506000910152565b6020815260008251806020840152610a52816040850160208701610a0f565b601f01601f19169190910160400192915050565b600080600080600060a08688031215610a7e57600080fd5b853567ffffffffffffffff80821115610a9657600080fd5b610aa289838a01610878565b9650602088013595506040880135915080821115610abf57600080fd5b610acb89838a01610878565b94506060880135915080821115610ae157600080fd5b50610aee88828901610878565b95989497509295608001359392505050565b60008251610b12818460208701610a0f565b9190910192915050565b634e487b7160e01b600052603260045260246000fd5b8060005b6004811015610b55578151845260209384019390910190600101610b36565b50505050565b6101008101610b6a8285610b32565b610b776080830184610b32565b9392505050565b600060808284031215610b9057600080fd5b82601f830112610b9f57600080fd5b6040516080810181811067ffffffffffffffff82111715610bc257610bc2610831565b604052806080840185811115610bd757600080fd5b845b81811015610bf1578051835260209283019201610bd9565b509195945050505050565b634e487b7160e01b600052601160045260246000fd5b600060018201610c2457610c24610bfc565b5060010190565b60a08101610c398285610b32565b8260808301529392505050565b608081016101128284610b32565b8060005b6008811015610b55578151845260209384019390910190600101610c58565b6103008101610c868287610b32565b610c936080830186610c54565b610ca1610180830185610b32565b610692610200830184610c54565b600060208284031215610cc157600080fd5b5051919050565b8181038181111561011257610112610bfc565b8082018082111561011257610112610bfc56fea264697066735822122057ae9e5c791e28263a37f09048955e674fe065aae95ff2bfba0b308da099e7b464736f6c63430008130033"

// DeployBLS deploys a new Ethereum contract, binding an instance of BLS to it.
func DeployBLS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BLS, error) {
	parsed, err := abi.JSON(strings.NewReader(BLSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	bLS12377Addr, _, _, _ := DeployBLS12377(auth, backend)
	BLSBin = strings.Replace(BLSBin, "__$51e475105ae19fccae74d465d2ad61b657$__", bLS12377Addr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BLSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLS{BLSCaller: BLSCaller{contract: contract}, BLSTransactor: BLSTransactor{contract: contract}, BLSFilterer: BLSFilterer{contract: contract}}, nil
}

// BLS is an auto generated Go binding around an Ethereum contract.
type BLS struct {
	BLSCaller     // Read-only binding to the contract
	BLSTransactor // Write-only binding to the contract
	BLSFilterer   // Log filterer for contract events
}

// BLSCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSSession struct {
	Contract     *BLS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSCallerSession struct {
	Contract *BLSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BLSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSTransactorSession struct {
	Contract     *BLSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSRaw struct {
	Contract *BLS // Generic contract binding to access the raw methods on
}

// BLSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSCallerRaw struct {
	Contract *BLSCaller // Generic read-only contract binding to access the raw methods on
}

// BLSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSTransactorRaw struct {
	Contract *BLSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLS creates a new instance of BLS, bound to a specific deployed contract.
func NewBLS(address common.Address, backend bind.ContractBackend) (*BLS, error) {
	contract, err := bindBLS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLS{BLSCaller: BLSCaller{contract: contract}, BLSTransactor: BLSTransactor{contract: contract}, BLSFilterer: BLSFilterer{contract: contract}}, nil
}

// NewBLSCaller creates a new read-only instance of BLS, bound to a specific deployed contract.
func NewBLSCaller(address common.Address, caller bind.ContractCaller) (*BLSCaller, error) {
	contract, err := bindBLS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSCaller{contract: contract}, nil
}

// NewBLSTransactor creates a new write-only instance of BLS, bound to a specific deployed contract.
func NewBLSTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSTransactor, error) {
	contract, err := bindBLS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSTransactor{contract: contract}, nil
}

// NewBLSFilterer creates a new log filterer instance of BLS, bound to a specific deployed contract.
func NewBLSFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSFilterer, error) {
	contract, err := bindBLS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSFilterer{contract: contract}, nil
}

// bindBLS binds a generic wrapper to an already deployed contract.
func bindBLS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BLSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLS *BLSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLS.Contract.BLSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLS *BLSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLS.Contract.BLSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLS *BLSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLS.Contract.BLSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLS *BLSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLS *BLSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLS *BLSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLS.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0xb2bd8933.
//
// Solidity: function add(bytes[] _com) view returns(bytes)
func (_BLS *BLSCaller) Add(opts *bind.CallOpts, _com [][]byte) ([]byte, error) {
	var out []interface{}
	err := _BLS.contract.Call(opts, &out, "add", _com)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0xb2bd8933.
//
// Solidity: function add(bytes[] _com) view returns(bytes)
func (_BLS *BLSSession) Add(_com [][]byte) ([]byte, error) {
	return _BLS.Contract.Add(&_BLS.CallOpts, _com)
}

// Add is a free data retrieval call binding the contract method 0xb2bd8933.
//
// Solidity: function add(bytes[] _com) view returns(bytes)
func (_BLS *BLSCallerSession) Add(_com [][]byte) ([]byte, error) {
	return _BLS.Contract.Add(&_BLS.CallOpts, _com)
}

// Equal is a free data retrieval call binding the contract method 0x235266d2.
//
// Solidity: function equal(bytes _a, bytes _b) pure returns(bool)
func (_BLS *BLSCaller) Equal(opts *bind.CallOpts, _a []byte, _b []byte) (bool, error) {
	var out []interface{}
	err := _BLS.contract.Call(opts, &out, "equal", _a, _b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equal is a free data retrieval call binding the contract method 0x235266d2.
//
// Solidity: function equal(bytes _a, bytes _b) pure returns(bool)
func (_BLS *BLSSession) Equal(_a []byte, _b []byte) (bool, error) {
	return _BLS.Contract.Equal(&_BLS.CallOpts, _a, _b)
}

// Equal is a free data retrieval call binding the contract method 0x235266d2.
//
// Solidity: function equal(bytes _a, bytes _b) pure returns(bool)
func (_BLS *BLSCallerSession) Equal(_a []byte, _b []byte) (bool, error) {
	return _BLS.Contract.Equal(&_BLS.CallOpts, _a, _b)
}

// VerifyKZG is a free data retrieval call binding the contract method 0xd30aa507.
//
// Solidity: function verifyKZG(bytes _vk, bytes32 _rnd, bytes _commit, bytes _psi, bytes32 _y) view returns(bool)
func (_BLS *BLSCaller) VerifyKZG(opts *bind.CallOpts, _vk []byte, _rnd [32]byte, _commit []byte, _psi []byte, _y [32]byte) (bool, error) {
	var out []interface{}
	err := _BLS.contract.Call(opts, &out, "verifyKZG", _vk, _rnd, _commit, _psi, _y)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyKZG is a free data retrieval call binding the contract method 0xd30aa507.
//
// Solidity: function verifyKZG(bytes _vk, bytes32 _rnd, bytes _commit, bytes _psi, bytes32 _y) view returns(bool)
func (_BLS *BLSSession) VerifyKZG(_vk []byte, _rnd [32]byte, _commit []byte, _psi []byte, _y [32]byte) (bool, error) {
	return _BLS.Contract.VerifyKZG(&_BLS.CallOpts, _vk, _rnd, _commit, _psi, _y)
}

// VerifyKZG is a free data retrieval call binding the contract method 0xd30aa507.
//
// Solidity: function verifyKZG(bytes _vk, bytes32 _rnd, bytes _commit, bytes _psi, bytes32 _y) view returns(bool)
func (_BLS *BLSCallerSession) VerifyKZG(_vk []byte, _rnd [32]byte, _commit []byte, _psi []byte, _y [32]byte) (bool, error) {
	return _BLS.Contract.VerifyKZG(&_BLS.CallOpts, _vk, _rnd, _commit, _psi, _y)
}

// BLS12377ABI is the input ABI used to generate the binding from.
const BLS12377ABI = "[{\"inputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"p1\",\"type\":\"bytes32[4]\"},{\"internalType\":\"bytes32[4]\",\"name\":\"p2\",\"type\":\"bytes32[4]\"}],\"name\":\"g1Add\",\"outputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"p\",\"type\":\"bytes32[4]\"},{\"internalType\":\"bytes32\",\"name\":\"scalar\",\"type\":\"bytes32\"}],\"name\":\"g1Mul\",\"outputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"p\",\"type\":\"bytes32[4]\"}],\"name\":\"g1Neg\",\"outputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"\",\"type\":\"bytes32[4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[8]\",\"name\":\"p1\",\"type\":\"bytes32[8]\"},{\"internalType\":\"bytes32[8]\",\"name\":\"p2\",\"type\":\"bytes32[8]\"}],\"name\":\"g2Add\",\"outputs\":[{\"internalType\":\"bytes32[8]\",\"name\":\"\",\"type\":\"bytes32[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[8]\",\"name\":\"p\",\"type\":\"bytes32[8]\"},{\"internalType\":\"bytes32\",\"name\":\"scalar\",\"type\":\"bytes32\"}],\"name\":\"g2Mul\",\"outputs\":[{\"internalType\":\"bytes32[8]\",\"name\":\"\",\"type\":\"bytes32[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[8]\",\"name\":\"p\",\"type\":\"bytes32[8]\"}],\"name\":\"g2Neg\",\"outputs\":[{\"internalType\":\"bytes32[8]\",\"name\":\"\",\"type\":\"bytes32[8]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[4]\",\"name\":\"p1\",\"type\":\"bytes32[4]\"},{\"internalType\":\"bytes32[8]\",\"name\":\"p2\",\"type\":\"bytes32[8]\"},{\"internalType\":\"bytes32[4]\",\"name\":\"p3\",\"type\":\"bytes32[4]\"},{\"internalType\":\"bytes32[8]\",\"name\":\"p4\",\"type\":\"bytes32[8]\"}],\"name\":\"pairing\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BLS12377FuncSigs maps the 4-byte function signature to its string representation.
var BLS12377FuncSigs = map[string]string{
	"25f1adc2": "g1Add(bytes32[4],bytes32[4])",
	"be846b0e": "g1Mul(bytes32[4],bytes32)",
	"1dd41e26": "g1Neg(bytes32[4])",
	"5c7baa45": "g2Add(bytes32[8],bytes32[8])",
	"5d234075": "g2Mul(bytes32[8],bytes32)",
	"e7eb5f61": "g2Neg(bytes32[8])",
	"357d55dd": "pairing(bytes32[4],bytes32[8],bytes32[4],bytes32[8])",
}

// BLS12377Bin is the compiled bytecode used for deploying new contracts.
var BLS12377Bin = "0x610b9f61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100875760003560e01c80635c7baa45116100655780635c7baa45146100e95780635d23407514610109578063be846b0e1461011c578063e7eb5f611461012f57600080fd5b80631dd41e261461008c57806325f1adc2146100b5578063357d55dd146100c8575b600080fd5b61009f61009a3660046108ca565b610142565b6040516100ac91906108ed565b60405180910390f35b61009f6100c336600461091e565b610298565b6100db6100d63660046109b7565b610315565b6040519081526020016100ac565b6100fc6100f7366004610a12565b6104be565b6040516100ac9190610a40565b6100fc610117366004610a69565b6105a2565b61009f61012a366004610a96565b61062c565b6100fc61013d366004610ac1565b610680565b61014a61078e565b606082015160408301516fffffffffffffffffffffffffffffffff8083169260801c91166f170b5d44300000008508c000000000016f1a22d9f300f5138f1ef3622fba0948006f01ae3a4617c510eac63b05c06ca1493b6000868410600181146101b757968403966101c6565b968403600160801b0196600191505b508083039250858310600181146101e45795830395600091506101f3565b958303600160801b0195600191505b509003838110600181146100875793810393505060408051600060208201526001600160801b0319608086901b1660308201520191506102309050565b60405160208183030381529060405261024890610af4565b60408087019190915280516001600160801b0319608085811b8216602084015286901b1660308201520160405160208183030381529060405261028a90610af4565b606086015250929392505050565b6102a061078e565b6102a86107ac565b8351815260208085015182820152604080860151818401526060808701518185015285516080808601919091529286015160a08501529085015160c084015284015160e08301528461010083600a610258fa61030a576040513d6000823e3d81fd5b839150505b92915050565b600061031f6107cb565b60005b60048160ff16101561037457868160ff166004811061034357610343610ade565b6020020151828260ff166018811061035d5761035d610ade565b60200201528061036c81610b31565b915050610322565b5060005b60088160ff1610156103d457858160ff166008811061039957610399610ade565b6020020151826103aa836004610b50565b60ff16601881106103bd576103bd610ade565b6020020152806103cc81610b31565b915050610378565b5060005b60048160ff16101561043457848160ff16600481106103f9576103f9610ade565b60200201518261040a83600c610b50565b60ff166018811061041d5761041d610ade565b60200201528061042c81610b31565b9150506103d8565b5060005b60088160ff16101561049457838160ff166008811061045957610459610ade565b60200201518261046a836010610b50565b60ff166018811061047d5761047d610ade565b60200201528061048c81610b31565b915050610438565b50602086610300836010620274e8fa6104b3576040513d6000823e3d81fd5b505092519392505050565b6104c66107ac565b6104ce6107ea565b60005b60088160ff16101561052357848160ff16600881106104f2576104f2610ade565b6020020151828260ff166010811061050c5761050c610ade565b60200201528061051b81610b31565b9150506104d1565b5060005b60088160ff16101561058357838160ff166008811061054857610548610ade565b602002015182610559836008610b50565b60ff166010811061056c5761056c610ade565b60200201528061057b81610b31565b915050610527565b506101008461020083600d611194fa61030a576040513d6000823e3d81fd5b6105aa6107ac565b6105b2610809565b60005b60088160ff16101561060757848160ff16600881106105d6576105d6610ade565b6020020151828260ff16600981106105f0576105f0610ade565b6020020152806105ff81610b31565b9150506105b5565b506101008082018490528461012083600e61d6d8fa61030a576040513d6000823e3d81fd5b61063461078e565b61063c610828565b8351815260208085015190820152604080850151908201526060808501519082015260808082018490528460a083600b612ee0fa61030a576040513d6000823e3d81fd5b6106886107ac565b60006040518060800160405280846002600881106106a8576106a8610ade565b60200201518152602001846003600881106106c5576106c5610ade565b60200201518152602001846004600881106106e2576106e2610ade565b60200201518152602001846005600881106106ff576106ff610ade565b60200201519052905061071181610142565b6040818101516080868101918252606084015160a088019081528351918201845291518152905160208201529192508101846006602002015181526020018460076008811061076257610762610ade565b60200201519052905061077481610142565b604081015160c08501526060015160e08401525090919050565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061020001604052806010906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261086d57600080fd5b6040516080810181811067ffffffffffffffff8211171561089057610890610846565b6040528060808401858111156108a557600080fd5b845b818110156108bf5780358352602092830192016108a7565b509195945050505050565b6000608082840312156108dc57600080fd5b6108e6838361085c565b9392505050565b60808101818360005b60048110156109155781518352602092830192909101906001016108f6565b50505092915050565b600080610100838503121561093257600080fd5b61093c848461085c565b915061094b846080850161085c565b90509250929050565b600082601f83011261096557600080fd5b60405161010080820182811067ffffffffffffffff8211171561098a5761098a610846565b6040528301818582111561099d57600080fd5b845b828110156108bf57803582526020918201910161099f565b60008060008061030085870312156109ce57600080fd5b6109d8868661085c565b93506109e78660808701610954565b92506109f786610180870161085c565b9150610a07866102008701610954565b905092959194509250565b6000806102008385031215610a2657600080fd5b610a308484610954565b915061094b846101008501610954565b6101008101818360005b6008811015610915578151835260209283019290910190600101610a4a565b6000806101208385031215610a7d57600080fd5b610a878484610954565b94610100939093013593505050565b60008060a08385031215610aa957600080fd5b610ab3848461085c565b946080939093013593505050565b60006101008284031215610ad457600080fd5b6108e68383610954565b634e487b7160e01b600052603260045260246000fd5b80516020808301519190811015610b15576000198160200360031b1b821691505b50919050565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff8103610b4757610b47610b1b565b60010192915050565b60ff818116838216019081111561030f5761030f610b1b56fea26469706673582212203908f2528327660025feffd749d60b146f193198a37e4f2fe5187accc111c90a64736f6c63430008130033"

// DeployBLS12377 deploys a new Ethereum contract, binding an instance of BLS12377 to it.
func DeployBLS12377(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BLS12377, error) {
	parsed, err := abi.JSON(strings.NewReader(BLS12377ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BLS12377Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLS12377{BLS12377Caller: BLS12377Caller{contract: contract}, BLS12377Transactor: BLS12377Transactor{contract: contract}, BLS12377Filterer: BLS12377Filterer{contract: contract}}, nil
}

// BLS12377 is an auto generated Go binding around an Ethereum contract.
type BLS12377 struct {
	BLS12377Caller     // Read-only binding to the contract
	BLS12377Transactor // Write-only binding to the contract
	BLS12377Filterer   // Log filterer for contract events
}

// BLS12377Caller is an auto generated read-only Go binding around an Ethereum contract.
type BLS12377Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLS12377Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BLS12377Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLS12377Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLS12377Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLS12377Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLS12377Session struct {
	Contract     *BLS12377         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLS12377CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLS12377CallerSession struct {
	Contract *BLS12377Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BLS12377TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLS12377TransactorSession struct {
	Contract     *BLS12377Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BLS12377Raw is an auto generated low-level Go binding around an Ethereum contract.
type BLS12377Raw struct {
	Contract *BLS12377 // Generic contract binding to access the raw methods on
}

// BLS12377CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLS12377CallerRaw struct {
	Contract *BLS12377Caller // Generic read-only contract binding to access the raw methods on
}

// BLS12377TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLS12377TransactorRaw struct {
	Contract *BLS12377Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBLS12377 creates a new instance of BLS12377, bound to a specific deployed contract.
func NewBLS12377(address common.Address, backend bind.ContractBackend) (*BLS12377, error) {
	contract, err := bindBLS12377(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLS12377{BLS12377Caller: BLS12377Caller{contract: contract}, BLS12377Transactor: BLS12377Transactor{contract: contract}, BLS12377Filterer: BLS12377Filterer{contract: contract}}, nil
}

// NewBLS12377Caller creates a new read-only instance of BLS12377, bound to a specific deployed contract.
func NewBLS12377Caller(address common.Address, caller bind.ContractCaller) (*BLS12377Caller, error) {
	contract, err := bindBLS12377(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLS12377Caller{contract: contract}, nil
}

// NewBLS12377Transactor creates a new write-only instance of BLS12377, bound to a specific deployed contract.
func NewBLS12377Transactor(address common.Address, transactor bind.ContractTransactor) (*BLS12377Transactor, error) {
	contract, err := bindBLS12377(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLS12377Transactor{contract: contract}, nil
}

// NewBLS12377Filterer creates a new log filterer instance of BLS12377, bound to a specific deployed contract.
func NewBLS12377Filterer(address common.Address, filterer bind.ContractFilterer) (*BLS12377Filterer, error) {
	contract, err := bindBLS12377(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLS12377Filterer{contract: contract}, nil
}

// bindBLS12377 binds a generic wrapper to an already deployed contract.
func bindBLS12377(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BLS12377ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLS12377 *BLS12377Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLS12377.Contract.BLS12377Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLS12377 *BLS12377Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLS12377.Contract.BLS12377Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLS12377 *BLS12377Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLS12377.Contract.BLS12377Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLS12377 *BLS12377CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLS12377.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLS12377 *BLS12377TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLS12377.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLS12377 *BLS12377TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLS12377.Contract.contract.Transact(opts, method, params...)
}

// G1Add is a free data retrieval call binding the contract method 0x25f1adc2.
//
// Solidity: function g1Add(bytes32[4] p1, bytes32[4] p2) view returns(bytes32[4])
func (_BLS12377 *BLS12377Caller) G1Add(opts *bind.CallOpts, p1 [4][32]byte, p2 [4][32]byte) ([4][32]byte, error) {
	var out []interface{}
	err := _BLS12377.contract.Call(opts, &out, "g1Add", p1, p2)

	if err != nil {
		return *new([4][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4][32]byte)).(*[4][32]byte)

	return out0, err

}

// G1Add is a free data retrieval call binding the contract method 0x25f1adc2.
//
// Solidity: function g1Add(bytes32[4] p1, bytes32[4] p2) view returns(bytes32[4])
func (_BLS12377 *BLS12377Session) G1Add(p1 [4][32]byte, p2 [4][32]byte) ([4][32]byte, error) {
	return _BLS12377.Contract.G1Add(&_BLS12377.CallOpts, p1, p2)
}

// G1Add is a free data retrieval call binding the contract method 0x25f1adc2.
//
// Solidity: function g1Add(bytes32[4] p1, bytes32[4] p2) view returns(bytes32[4])
func (_BLS12377 *BLS12377CallerSession) G1Add(p1 [4][32]byte, p2 [4][32]byte) ([4][32]byte, error) {
	return _BLS12377.Contract.G1Add(&_BLS12377.CallOpts, p1, p2)
}

// G1Mul is a free data retrieval call binding the contract method 0xbe846b0e.
//
// Solidity: function g1Mul(bytes32[4] p, bytes32 scalar) view returns(bytes32[4])
func (_BLS12377 *BLS12377Caller) G1Mul(opts *bind.CallOpts, p [4][32]byte, scalar [32]byte) ([4][32]byte, error) {
	var out []interface{}
	err := _BLS12377.contract.Call(opts, &out, "g1Mul", p, scalar)

	if err != nil {
		return *new([4][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4][32]byte)).(*[4][32]byte)

	return out0, err

}

// G1Mul is a free data retrieval call binding the contract method 0xbe846b0e.
//
// Solidity: function g1Mul(bytes32[4] p, bytes32 scalar) view returns(bytes32[4])
func (_BLS12377 *BLS12377Session) G1Mul(p [4][32]byte, scalar [32]byte) ([4][32]byte, error) {
	return _BLS12377.Contract.G1Mul(&_BLS12377.CallOpts, p, scalar)
}

// G1Mul is a free data retrieval call binding the contract method 0xbe846b0e.
//
// Solidity: function g1Mul(bytes32[4] p, bytes32 scalar) view returns(bytes32[4])
func (_BLS12377 *BLS12377CallerSession) G1Mul(p [4][32]byte, scalar [32]byte) ([4][32]byte, error) {
	return _BLS12377.Contract.G1Mul(&_BLS12377.CallOpts, p, scalar)
}

// G1Neg is a free data retrieval call binding the contract method 0x1dd41e26.
//
// Solidity: function g1Neg(bytes32[4] p) pure returns(bytes32[4])
func (_BLS12377 *BLS12377Caller) G1Neg(opts *bind.CallOpts, p [4][32]byte) ([4][32]byte, error) {
	var out []interface{}
	err := _BLS12377.contract.Call(opts, &out, "g1Neg", p)

	if err != nil {
		return *new([4][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4][32]byte)).(*[4][32]byte)

	return out0, err

}

// G1Neg is a free data retrieval call binding the contract method 0x1dd41e26.
//
// Solidity: function g1Neg(bytes32[4] p) pure returns(bytes32[4])
func (_BLS12377 *BLS12377Session) G1Neg(p [4][32]byte) ([4][32]byte, error) {
	return _BLS12377.Contract.G1Neg(&_BLS12377.CallOpts, p)
}

// G1Neg is a free data retrieval call binding the contract method 0x1dd41e26.
//
// Solidity: function g1Neg(bytes32[4] p) pure returns(bytes32[4])
func (_BLS12377 *BLS12377CallerSession) G1Neg(p [4][32]byte) ([4][32]byte, error) {
	return _BLS12377.Contract.G1Neg(&_BLS12377.CallOpts, p)
}

// G2Add is a free data retrieval call binding the contract method 0x5c7baa45.
//
// Solidity: function g2Add(bytes32[8] p1, bytes32[8] p2) view returns(bytes32[8])
func (_BLS12377 *BLS12377Caller) G2Add(opts *bind.CallOpts, p1 [8][32]byte, p2 [8][32]byte) ([8][32]byte, error) {
	var out []interface{}
	err := _BLS12377.contract.Call(opts, &out, "g2Add", p1, p2)

	if err != nil {
		return *new([8][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8][32]byte)).(*[8][32]byte)

	return out0, err

}

// G2Add is a free data retrieval call binding the contract method 0x5c7baa45.
//
// Solidity: function g2Add(bytes32[8] p1, bytes32[8] p2) view returns(bytes32[8])
func (_BLS12377 *BLS12377Session) G2Add(p1 [8][32]byte, p2 [8][32]byte) ([8][32]byte, error) {
	return _BLS12377.Contract.G2Add(&_BLS12377.CallOpts, p1, p2)
}

// G2Add is a free data retrieval call binding the contract method 0x5c7baa45.
//
// Solidity: function g2Add(bytes32[8] p1, bytes32[8] p2) view returns(bytes32[8])
func (_BLS12377 *BLS12377CallerSession) G2Add(p1 [8][32]byte, p2 [8][32]byte) ([8][32]byte, error) {
	return _BLS12377.Contract.G2Add(&_BLS12377.CallOpts, p1, p2)
}

// G2Mul is a free data retrieval call binding the contract method 0x5d234075.
//
// Solidity: function g2Mul(bytes32[8] p, bytes32 scalar) view returns(bytes32[8])
func (_BLS12377 *BLS12377Caller) G2Mul(opts *bind.CallOpts, p [8][32]byte, scalar [32]byte) ([8][32]byte, error) {
	var out []interface{}
	err := _BLS12377.contract.Call(opts, &out, "g2Mul", p, scalar)

	if err != nil {
		return *new([8][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8][32]byte)).(*[8][32]byte)

	return out0, err

}

// G2Mul is a free data retrieval call binding the contract method 0x5d234075.
//
// Solidity: function g2Mul(bytes32[8] p, bytes32 scalar) view returns(bytes32[8])
func (_BLS12377 *BLS12377Session) G2Mul(p [8][32]byte, scalar [32]byte) ([8][32]byte, error) {
	return _BLS12377.Contract.G2Mul(&_BLS12377.CallOpts, p, scalar)
}

// G2Mul is a free data retrieval call binding the contract method 0x5d234075.
//
// Solidity: function g2Mul(bytes32[8] p, bytes32 scalar) view returns(bytes32[8])
func (_BLS12377 *BLS12377CallerSession) G2Mul(p [8][32]byte, scalar [32]byte) ([8][32]byte, error) {
	return _BLS12377.Contract.G2Mul(&_BLS12377.CallOpts, p, scalar)
}

// G2Neg is a free data retrieval call binding the contract method 0xe7eb5f61.
//
// Solidity: function g2Neg(bytes32[8] p) pure returns(bytes32[8])
func (_BLS12377 *BLS12377Caller) G2Neg(opts *bind.CallOpts, p [8][32]byte) ([8][32]byte, error) {
	var out []interface{}
	err := _BLS12377.contract.Call(opts, &out, "g2Neg", p)

	if err != nil {
		return *new([8][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8][32]byte)).(*[8][32]byte)

	return out0, err

}

// G2Neg is a free data retrieval call binding the contract method 0xe7eb5f61.
//
// Solidity: function g2Neg(bytes32[8] p) pure returns(bytes32[8])
func (_BLS12377 *BLS12377Session) G2Neg(p [8][32]byte) ([8][32]byte, error) {
	return _BLS12377.Contract.G2Neg(&_BLS12377.CallOpts, p)
}

// G2Neg is a free data retrieval call binding the contract method 0xe7eb5f61.
//
// Solidity: function g2Neg(bytes32[8] p) pure returns(bytes32[8])
func (_BLS12377 *BLS12377CallerSession) G2Neg(p [8][32]byte) ([8][32]byte, error) {
	return _BLS12377.Contract.G2Neg(&_BLS12377.CallOpts, p)
}

// Pairing is a free data retrieval call binding the contract method 0x357d55dd.
//
// Solidity: function pairing(bytes32[4] p1, bytes32[8] p2, bytes32[4] p3, bytes32[8] p4) view returns(bytes32)
func (_BLS12377 *BLS12377Caller) Pairing(opts *bind.CallOpts, p1 [4][32]byte, p2 [8][32]byte, p3 [4][32]byte, p4 [8][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BLS12377.contract.Call(opts, &out, "pairing", p1, p2, p3, p4)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Pairing is a free data retrieval call binding the contract method 0x357d55dd.
//
// Solidity: function pairing(bytes32[4] p1, bytes32[8] p2, bytes32[4] p3, bytes32[8] p4) view returns(bytes32)
func (_BLS12377 *BLS12377Session) Pairing(p1 [4][32]byte, p2 [8][32]byte, p3 [4][32]byte, p4 [8][32]byte) ([32]byte, error) {
	return _BLS12377.Contract.Pairing(&_BLS12377.CallOpts, p1, p2, p3, p4)
}

// Pairing is a free data retrieval call binding the contract method 0x357d55dd.
//
// Solidity: function pairing(bytes32[4] p1, bytes32[8] p2, bytes32[4] p3, bytes32[8] p4) view returns(bytes32)
func (_BLS12377 *BLS12377CallerSession) Pairing(p1 [4][32]byte, p2 [8][32]byte, p3 [4][32]byte, p4 [8][32]byte) ([32]byte, error) {
	return _BLS12377.Contract.Pairing(&_BLS12377.CallOpts, p1, p2, p3, p4)
}

// IBLSABI is the input ABI used to generate the binding from.
const IBLSABI = "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_com\",\"type\":\"bytes[]\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_b\",\"type\":\"bytes\"}],\"name\":\"equal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vk\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_rnd\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_commit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_psi\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_y\",\"type\":\"bytes32\"}],\"name\":\"verifyKZG\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBLSFuncSigs maps the 4-byte function signature to its string representation.
var IBLSFuncSigs = map[string]string{
	"b2bd8933": "add(bytes[])",
	"235266d2": "equal(bytes,bytes)",
	"d30aa507": "verifyKZG(bytes,bytes32,bytes,bytes,bytes32)",
}

// IBLS is an auto generated Go binding around an Ethereum contract.
type IBLS struct {
	IBLSCaller     // Read-only binding to the contract
	IBLSTransactor // Write-only binding to the contract
	IBLSFilterer   // Log filterer for contract events
}

// IBLSCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBLSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBLSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBLSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBLSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBLSSession struct {
	Contract     *IBLS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBLSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBLSCallerSession struct {
	Contract *IBLSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBLSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBLSTransactorSession struct {
	Contract     *IBLSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBLSRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBLSRaw struct {
	Contract *IBLS // Generic contract binding to access the raw methods on
}

// IBLSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBLSCallerRaw struct {
	Contract *IBLSCaller // Generic read-only contract binding to access the raw methods on
}

// IBLSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBLSTransactorRaw struct {
	Contract *IBLSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBLS creates a new instance of IBLS, bound to a specific deployed contract.
func NewIBLS(address common.Address, backend bind.ContractBackend) (*IBLS, error) {
	contract, err := bindIBLS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBLS{IBLSCaller: IBLSCaller{contract: contract}, IBLSTransactor: IBLSTransactor{contract: contract}, IBLSFilterer: IBLSFilterer{contract: contract}}, nil
}

// NewIBLSCaller creates a new read-only instance of IBLS, bound to a specific deployed contract.
func NewIBLSCaller(address common.Address, caller bind.ContractCaller) (*IBLSCaller, error) {
	contract, err := bindIBLS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBLSCaller{contract: contract}, nil
}

// NewIBLSTransactor creates a new write-only instance of IBLS, bound to a specific deployed contract.
func NewIBLSTransactor(address common.Address, transactor bind.ContractTransactor) (*IBLSTransactor, error) {
	contract, err := bindIBLS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBLSTransactor{contract: contract}, nil
}

// NewIBLSFilterer creates a new log filterer instance of IBLS, bound to a specific deployed contract.
func NewIBLSFilterer(address common.Address, filterer bind.ContractFilterer) (*IBLSFilterer, error) {
	contract, err := bindIBLS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBLSFilterer{contract: contract}, nil
}

// bindIBLS binds a generic wrapper to an already deployed contract.
func bindIBLS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBLSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBLS *IBLSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBLS.Contract.IBLSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBLS *IBLSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBLS.Contract.IBLSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBLS *IBLSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBLS.Contract.IBLSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBLS *IBLSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBLS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBLS *IBLSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBLS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBLS *IBLSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBLS.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0xb2bd8933.
//
// Solidity: function add(bytes[] _com) returns(bytes)
func (_IBLS *IBLSTransactor) Add(opts *bind.TransactOpts, _com [][]byte) (*types.Transaction, error) {
	return _IBLS.contract.Transact(opts, "add", _com)
}

// Add is a paid mutator transaction binding the contract method 0xb2bd8933.
//
// Solidity: function add(bytes[] _com) returns(bytes)
func (_IBLS *IBLSSession) Add(_com [][]byte) (*types.Transaction, error) {
	return _IBLS.Contract.Add(&_IBLS.TransactOpts, _com)
}

// Add is a paid mutator transaction binding the contract method 0xb2bd8933.
//
// Solidity: function add(bytes[] _com) returns(bytes)
func (_IBLS *IBLSTransactorSession) Add(_com [][]byte) (*types.Transaction, error) {
	return _IBLS.Contract.Add(&_IBLS.TransactOpts, _com)
}

// Equal is a paid mutator transaction binding the contract method 0x235266d2.
//
// Solidity: function equal(bytes _a, bytes _b) returns(bool)
func (_IBLS *IBLSTransactor) Equal(opts *bind.TransactOpts, _a []byte, _b []byte) (*types.Transaction, error) {
	return _IBLS.contract.Transact(opts, "equal", _a, _b)
}

// Equal is a paid mutator transaction binding the contract method 0x235266d2.
//
// Solidity: function equal(bytes _a, bytes _b) returns(bool)
func (_IBLS *IBLSSession) Equal(_a []byte, _b []byte) (*types.Transaction, error) {
	return _IBLS.Contract.Equal(&_IBLS.TransactOpts, _a, _b)
}

// Equal is a paid mutator transaction binding the contract method 0x235266d2.
//
// Solidity: function equal(bytes _a, bytes _b) returns(bool)
func (_IBLS *IBLSTransactorSession) Equal(_a []byte, _b []byte) (*types.Transaction, error) {
	return _IBLS.Contract.Equal(&_IBLS.TransactOpts, _a, _b)
}

// VerifyKZG is a paid mutator transaction binding the contract method 0xd30aa507.
//
// Solidity: function verifyKZG(bytes _vk, bytes32 _rnd, bytes _commit, bytes _psi, bytes32 _y) returns(bool)
func (_IBLS *IBLSTransactor) VerifyKZG(opts *bind.TransactOpts, _vk []byte, _rnd [32]byte, _commit []byte, _psi []byte, _y [32]byte) (*types.Transaction, error) {
	return _IBLS.contract.Transact(opts, "verifyKZG", _vk, _rnd, _commit, _psi, _y)
}

// VerifyKZG is a paid mutator transaction binding the contract method 0xd30aa507.
//
// Solidity: function verifyKZG(bytes _vk, bytes32 _rnd, bytes _commit, bytes _psi, bytes32 _y) returns(bool)
func (_IBLS *IBLSSession) VerifyKZG(_vk []byte, _rnd [32]byte, _commit []byte, _psi []byte, _y [32]byte) (*types.Transaction, error) {
	return _IBLS.Contract.VerifyKZG(&_IBLS.TransactOpts, _vk, _rnd, _commit, _psi, _y)
}

// VerifyKZG is a paid mutator transaction binding the contract method 0xd30aa507.
//
// Solidity: function verifyKZG(bytes _vk, bytes32 _rnd, bytes _commit, bytes _psi, bytes32 _y) returns(bool)
func (_IBLS *IBLSTransactorSession) VerifyKZG(_vk []byte, _rnd [32]byte, _commit []byte, _psi []byte, _y [32]byte) (*types.Transaction, error) {
	return _IBLS.Contract.VerifyKZG(&_IBLS.TransactOpts, _vk, _rnd, _commit, _psi, _y)
}
