// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package control

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

// ControlMetaData contains all meta data concerning the Control contract.
var ControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Punish\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610e6a380380610e6a833981810160405281019061003291906100db565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100a88261007d565b9050919050565b6100b88161009d565b81146100c357600080fd5b50565b6000815190506100d5816100af565b92915050565b6000602082840312156100f1576100f0610078565b5b60006100ff848285016100c6565b91505092915050565b610d53806101176000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806306f0b4f11461004657806340c10f191461006257806376cdb03b1461007e575b600080fd5b610060600480360381019061005b919061080d565b61009c565b005b61007c60048036038101906100779190610874565b6103e1565b005b610086610717565b60405161009391906108c3565b60405180910390f35b60018360ff16036101b35760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016100fe9061093b565b6020604051808303816000875af115801561011d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101419190610970565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a5906109e9565b60405180910390fd5b6102d0565b60038360ff16036102ca5760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161021590610a55565b6020604051808303816000875af1158015610234573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102589190610970565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102bc90610ac1565b60405180910390fd5b6102cf565b6103db565b5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161032790610b2d565b6020604051808303816000875af1158015610346573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036a9190610970565b73ffffffffffffffffffffffffffffffffffffffff166306f0b4f1858585856040518563ffffffff1660e01b81526004016103a89493929190610b6b565b600060405180830381600087803b1580156103c257600080fd5b505af11580156103d6573d6000803e3d6000fd5b505050505b50505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161043890610bfc565b6020604051808303816000875af1158015610457573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047b9190610970565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610578575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161050690610c68565b6020604051808303816000875af1158015610525573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105499190610970565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80610647575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b81526004016105d590610a55565b6020604051808303816000875af11580156105f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106189190610970565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610686576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067d90610cd4565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff1660e01b81526004016106e1929190610cf4565b600060405180830381600087803b1580156106fb57600080fd5b505af115801561070f573d6000803e3d6000fd5b505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061076b82610740565b9050919050565b61077b81610760565b811461078657600080fd5b50565b60008135905061079881610772565b92915050565b600060ff82169050919050565b6107b48161079e565b81146107bf57600080fd5b50565b6000813590506107d1816107ab565b92915050565b6000819050919050565b6107ea816107d7565b81146107f557600080fd5b50565b600081359050610807816107e1565b92915050565b600080600080608085870312156108275761082661073b565b5b600061083587828801610789565b9450506020610846878288016107c2565b935050604061085787828801610789565b9250506060610868878288016107f8565b91505092959194509250565b6000806040838503121561088b5761088a61073b565b5b600061089985828601610789565b92505060206108aa858286016107f8565b9150509250929050565b6108bd81610760565b82525050565b60006020820190506108d860008301846108b4565b92915050565b600082825260208201905092915050565b7f70726f6f66000000000000000000000000000000000000000000000000000000600082015250565b60006109256005836108de565b9150610930826108ef565b602082019050919050565b6000602082019050818103600083015261095481610918565b9050919050565b60008151905061096a81610772565b92915050565b6000602082840312156109865761098561073b565b5b60006109948482850161095b565b91505092915050565b7f696e76616c696420700000000000000000000000000000000000000000000000600082015250565b60006109d36009836108de565b91506109de8261099d565b602082019050919050565b60006020820190508181036000830152610a02816109c6565b9050919050565b7f7370616365000000000000000000000000000000000000000000000000000000600082015250565b6000610a3f6005836108de565b9150610a4a82610a09565b602082019050919050565b60006020820190508181036000830152610a6e81610a32565b9050919050565b7f696e76616c696420730000000000000000000000000000000000000000000000600082015250565b6000610aab6009836108de565b9150610ab682610a75565b602082019050919050565b60006020820190508181036000830152610ada81610a9e565b9050919050565b7f6e6f646500000000000000000000000000000000000000000000000000000000600082015250565b6000610b176004836108de565b9150610b2282610ae1565b602082019050919050565b60006020820190508181036000830152610b4681610b0a565b9050919050565b610b568161079e565b82525050565b610b65816107d7565b82525050565b6000608082019050610b8060008301876108b4565b610b8d6020830186610b4d565b610b9a60408301856108b4565b610ba76060830184610b5c565b95945050505050565b7f66696c6500000000000000000000000000000000000000000000000000000000600082015250565b6000610be66004836108de565b9150610bf182610bb0565b602082019050919050565b60006020820190508181036000830152610c1581610bd9565b9050919050565b7f6770750000000000000000000000000000000000000000000000000000000000600082015250565b6000610c526003836108de565b9150610c5d82610c1c565b602082019050919050565b60006020820190508181036000830152610c8181610c45565b9050919050565b7f696e76616c6964206d6300000000000000000000000000000000000000000000600082015250565b6000610cbe600a836108de565b9150610cc982610c88565b602082019050919050565b60006020820190508181036000830152610ced81610cb1565b9050919050565b6000604082019050610d0960008301856108b4565b610d166020830184610b5c565b939250505056fea2646970667358221220d59f8a286e6de410b6b65b2ddf55eb750f7534eba8b658072589354db5028c2e64736f6c63430008180033",
}

// ControlABI is the input ABI used to generate the binding from.
// Deprecated: Use ControlMetaData.ABI instead.
var ControlABI = ControlMetaData.ABI

// ControlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ControlMetaData.Bin instead.
var ControlBin = ControlMetaData.Bin

// DeployControl deploys a new Ethereum contract, binding an instance of Control to it.
func DeployControl(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Control, error) {
	parsed, err := ControlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ControlBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Control{ControlCaller: ControlCaller{contract: contract}, ControlTransactor: ControlTransactor{contract: contract}, ControlFilterer: ControlFilterer{contract: contract}}, nil
}

// Control is an auto generated Go binding around an Ethereum contract.
type Control struct {
	ControlCaller     // Read-only binding to the contract
	ControlTransactor // Write-only binding to the contract
	ControlFilterer   // Log filterer for contract events
}

// ControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControlSession struct {
	Contract     *Control          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControlCallerSession struct {
	Contract *ControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControlTransactorSession struct {
	Contract     *ControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControlRaw struct {
	Contract *Control // Generic contract binding to access the raw methods on
}

// ControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControlCallerRaw struct {
	Contract *ControlCaller // Generic read-only contract binding to access the raw methods on
}

// ControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControlTransactorRaw struct {
	Contract *ControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewControl creates a new instance of Control, bound to a specific deployed contract.
func NewControl(address common.Address, backend bind.ContractBackend) (*Control, error) {
	contract, err := bindControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Control{ControlCaller: ControlCaller{contract: contract}, ControlTransactor: ControlTransactor{contract: contract}, ControlFilterer: ControlFilterer{contract: contract}}, nil
}

// NewControlCaller creates a new read-only instance of Control, bound to a specific deployed contract.
func NewControlCaller(address common.Address, caller bind.ContractCaller) (*ControlCaller, error) {
	contract, err := bindControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControlCaller{contract: contract}, nil
}

// NewControlTransactor creates a new write-only instance of Control, bound to a specific deployed contract.
func NewControlTransactor(address common.Address, transactor bind.ContractTransactor) (*ControlTransactor, error) {
	contract, err := bindControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControlTransactor{contract: contract}, nil
}

// NewControlFilterer creates a new log filterer instance of Control, bound to a specific deployed contract.
func NewControlFilterer(address common.Address, filterer bind.ContractFilterer) (*ControlFilterer, error) {
	contract, err := bindControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControlFilterer{contract: contract}, nil
}

// bindControl binds a generic wrapper to an already deployed contract.
func bindControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control *ControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Control.Contract.ControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control *ControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Control.Contract.ControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control *ControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Control.Contract.ControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control *ControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Control.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control *ControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Control.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control *ControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Control.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Control *ControlCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Control.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Control *ControlSession) Bank() (common.Address, error) {
	return _Control.Contract.Bank(&_Control.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Control *ControlCallerSession) Bank() (common.Address, error) {
	return _Control.Contract.Bank(&_Control.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _t, uint256 _m) returns()
func (_Control *ControlTransactor) Mint(opts *bind.TransactOpts, _t common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "mint", _t, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _t, uint256 _m) returns()
func (_Control *ControlSession) Mint(_t common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Mint(&_Control.TransactOpts, _t, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _t, uint256 _m) returns()
func (_Control *ControlTransactorSession) Mint(_t common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Mint(&_Control.TransactOpts, _t, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Control *ControlTransactor) Punish(opts *bind.TransactOpts, _from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.contract.Transact(opts, "punish", _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Control *ControlSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Punish(&_Control.TransactOpts, _from, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _from, uint8 _typ, address _to, uint256 _m) returns()
func (_Control *ControlTransactorSession) Punish(_from common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _Control.Contract.Punish(&_Control.TransactOpts, _from, _typ, _to, _m)
}

// ControlMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Control contract.
type ControlMintIterator struct {
	Event *ControlMint // Event containing the contract specifics and raw log

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
func (it *ControlMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControlMint)
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
		it.Event = new(ControlMint)
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
func (it *ControlMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControlMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControlMint represents a Mint event raised by the Control contract.
type ControlMint struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _a, uint256 _m)
func (_Control *ControlFilterer) FilterMint(opts *bind.FilterOpts, _a []common.Address) (*ControlMintIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Control.contract.FilterLogs(opts, "Mint", _aRule)
	if err != nil {
		return nil, err
	}
	return &ControlMintIterator{contract: _Control.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _a, uint256 _m)
func (_Control *ControlFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ControlMint, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Control.contract.WatchLogs(opts, "Mint", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControlMint)
				if err := _Control.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _a, uint256 _m)
func (_Control *ControlFilterer) ParseMint(log types.Log) (*ControlMint, error) {
	event := new(ControlMint)
	if err := _Control.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControlPunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the Control contract.
type ControlPunishIterator struct {
	Event *ControlPunish // Event containing the contract specifics and raw log

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
func (it *ControlPunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControlPunish)
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
		it.Event = new(ControlPunish)
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
func (it *ControlPunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControlPunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControlPunish represents a Punish event raised by the Control contract.
type ControlPunish struct {
	A   common.Address
	Typ uint8
	To  common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address _to, uint256 _m)
func (_Control *ControlFilterer) FilterPunish(opts *bind.FilterOpts, _a []common.Address) (*ControlPunishIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Control.contract.FilterLogs(opts, "Punish", _aRule)
	if err != nil {
		return nil, err
	}
	return &ControlPunishIterator{contract: _Control.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address _to, uint256 _m)
func (_Control *ControlFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *ControlPunish, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Control.contract.WatchLogs(opts, "Punish", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControlPunish)
				if err := _Control.contract.UnpackLog(event, "Punish", log); err != nil {
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

// ParsePunish is a log parse operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address _to, uint256 _m)
func (_Control *ControlFilterer) ParsePunish(log types.Log) (*ControlPunish, error) {
	event := new(ControlPunish)
	if err := _Control.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
