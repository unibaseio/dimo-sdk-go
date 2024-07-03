// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package space

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

// ISpaceInfo is an auto generated low-level Go binding around an user-defined struct.
type ISpaceInfo struct {
	Name   string
	Model  uint64
	Gpu    uint64
	Start  uint64
	Expire uint64
	Active bool
	Fin    bool
	Owner  common.Address
	Price  *big.Int
}

// IBankMetaData contains all meta data concerning the IBank contract.
var IBankMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_s\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_c\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"TransferIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_c\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"TransferOut\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_s\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_s\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IBankABI is the input ABI used to generate the binding from.
// Deprecated: Use IBankMetaData.ABI instead.
var IBankABI = IBankMetaData.ABI

// IBank is an auto generated Go binding around an Ethereum contract.
type IBank struct {
	IBankCaller     // Read-only binding to the contract
	IBankTransactor // Write-only binding to the contract
	IBankFilterer   // Log filterer for contract events
}

// IBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBankSession struct {
	Contract     *IBank            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBankCallerSession struct {
	Contract *IBankCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBankTransactorSession struct {
	Contract     *IBankTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBankRaw struct {
	Contract *IBank // Generic contract binding to access the raw methods on
}

// IBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBankCallerRaw struct {
	Contract *IBankCaller // Generic read-only contract binding to access the raw methods on
}

// IBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBankTransactorRaw struct {
	Contract *IBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBank creates a new instance of IBank, bound to a specific deployed contract.
func NewIBank(address common.Address, backend bind.ContractBackend) (*IBank, error) {
	contract, err := bindIBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBank{IBankCaller: IBankCaller{contract: contract}, IBankTransactor: IBankTransactor{contract: contract}, IBankFilterer: IBankFilterer{contract: contract}}, nil
}

// NewIBankCaller creates a new read-only instance of IBank, bound to a specific deployed contract.
func NewIBankCaller(address common.Address, caller bind.ContractCaller) (*IBankCaller, error) {
	contract, err := bindIBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBankCaller{contract: contract}, nil
}

// NewIBankTransactor creates a new write-only instance of IBank, bound to a specific deployed contract.
func NewIBankTransactor(address common.Address, transactor bind.ContractTransactor) (*IBankTransactor, error) {
	contract, err := bindIBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBankTransactor{contract: contract}, nil
}

// NewIBankFilterer creates a new log filterer instance of IBank, bound to a specific deployed contract.
func NewIBankFilterer(address common.Address, filterer bind.ContractFilterer) (*IBankFilterer, error) {
	contract, err := bindIBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBankFilterer{contract: contract}, nil
}

// bindIBank binds a generic wrapper to an already deployed contract.
func bindIBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBank *IBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBank.Contract.IBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBank *IBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBank.Contract.IBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBank *IBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBank.Contract.IBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBank *IBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBank *IBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBank *IBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBank.Contract.contract.Transact(opts, method, params...)
}

// Get is a paid mutator transaction binding the contract method 0x693ec85e.
//
// Solidity: function get(string _s) returns(address)
func (_IBank *IBankTransactor) Get(opts *bind.TransactOpts, _s string) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "get", _s)
}

// Get is a paid mutator transaction binding the contract method 0x693ec85e.
//
// Solidity: function get(string _s) returns(address)
func (_IBank *IBankSession) Get(_s string) (*types.Transaction, error) {
	return _IBank.Contract.Get(&_IBank.TransactOpts, _s)
}

// Get is a paid mutator transaction binding the contract method 0x693ec85e.
//
// Solidity: function get(string _s) returns(address)
func (_IBank *IBankTransactorSession) Get(_s string) (*types.Transaction, error) {
	return _IBank.Contract.Get(&_IBank.TransactOpts, _s)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactor) Mint(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "mint", _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) payable returns()
func (_IBank *IBankSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.Mint(&_IBank.TransactOpts, _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactorSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.Mint(&_IBank.TransactOpts, _a, _m)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(string _s, address _a) returns()
func (_IBank *IBankTransactor) Set(opts *bind.TransactOpts, _s string, _a common.Address) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "set", _s, _a)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(string _s, address _a) returns()
func (_IBank *IBankSession) Set(_s string, _a common.Address) (*types.Transaction, error) {
	return _IBank.Contract.Set(&_IBank.TransactOpts, _s, _a)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(string _s, address _a) returns()
func (_IBank *IBankTransactorSession) Set(_s string, _a common.Address) (*types.Transaction, error) {
	return _IBank.Contract.Set(&_IBank.TransactOpts, _s, _a)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactor) TransferIn(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "transferIn", _a, _m)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _a, uint256 _m) payable returns()
func (_IBank *IBankSession) TransferIn(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferIn(&_IBank.TransactOpts, _a, _m)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe8888915.
//
// Solidity: function transferIn(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactorSession) TransferIn(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferIn(&_IBank.TransactOpts, _a, _m)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactor) TransferOut(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "transferOut", _a, _m)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address _a, uint256 _m) payable returns()
func (_IBank *IBankSession) TransferOut(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferOut(&_IBank.TransactOpts, _a, _m)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address _a, uint256 _m) payable returns()
func (_IBank *IBankTransactorSession) TransferOut(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IBank.Contract.TransferOut(&_IBank.TransactOpts, _a, _m)
}

// IBankMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IBank contract.
type IBankMintIterator struct {
	Event *IBankMint // Event containing the contract specifics and raw log

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
func (it *IBankMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankMint)
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
		it.Event = new(IBankMint)
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
func (it *IBankMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankMint represents a Mint event raised by the IBank contract.
type IBankMint struct {
	To  common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) FilterMint(opts *bind.FilterOpts, _to []common.Address) (*IBankMintIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "Mint", _toRule)
	if err != nil {
		return nil, err
	}
	return &IBankMintIterator{contract: _IBank.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IBankMint, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "Mint", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankMint)
				if err := _IBank.contract.UnpackLog(event, "Mint", log); err != nil {
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
// Solidity: event Mint(address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) ParseMint(log types.Log) (*IBankMint, error) {
	event := new(IBankMint)
	if err := _IBank.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBankSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the IBank contract.
type IBankSetIterator struct {
	Event *IBankSet // Event containing the contract specifics and raw log

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
func (it *IBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankSet)
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
		it.Event = new(IBankSet)
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
func (it *IBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankSet represents a Set event raised by the IBank contract.
type IBankSet struct {
	S   string
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x496595ced95720268cf8bc60bae3f35024ff2a130f73ac4e20f5c1eaca35db99.
//
// Solidity: event Set(string _s, address _to)
func (_IBank *IBankFilterer) FilterSet(opts *bind.FilterOpts) (*IBankSetIterator, error) {

	logs, sub, err := _IBank.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &IBankSetIterator{contract: _IBank.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x496595ced95720268cf8bc60bae3f35024ff2a130f73ac4e20f5c1eaca35db99.
//
// Solidity: event Set(string _s, address _to)
func (_IBank *IBankFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *IBankSet) (event.Subscription, error) {

	logs, sub, err := _IBank.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankSet)
				if err := _IBank.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x496595ced95720268cf8bc60bae3f35024ff2a130f73ac4e20f5c1eaca35db99.
//
// Solidity: event Set(string _s, address _to)
func (_IBank *IBankFilterer) ParseSet(log types.Log) (*IBankSet, error) {
	event := new(IBankSet)
	if err := _IBank.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBankTransferInIterator is returned from FilterTransferIn and is used to iterate over the raw logs and unpacked data for TransferIn events raised by the IBank contract.
type IBankTransferInIterator struct {
	Event *IBankTransferIn // Event containing the contract specifics and raw log

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
func (it *IBankTransferInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankTransferIn)
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
		it.Event = new(IBankTransferIn)
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
func (it *IBankTransferInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankTransferInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankTransferIn represents a TransferIn event raised by the IBank contract.
type IBankTransferIn struct {
	C    common.Address
	From common.Address
	M    *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferIn is a free log retrieval operation binding the contract event 0x8ab008cb8444c6d8f2fa3dc0b159b5a86c62b00092219a6c69757805cbf7bde7.
//
// Solidity: event TransferIn(address indexed _c, address indexed _from, uint256 _m)
func (_IBank *IBankFilterer) FilterTransferIn(opts *bind.FilterOpts, _c []common.Address, _from []common.Address) (*IBankTransferInIterator, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "TransferIn", _cRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &IBankTransferInIterator{contract: _IBank.contract, event: "TransferIn", logs: logs, sub: sub}, nil
}

// WatchTransferIn is a free log subscription operation binding the contract event 0x8ab008cb8444c6d8f2fa3dc0b159b5a86c62b00092219a6c69757805cbf7bde7.
//
// Solidity: event TransferIn(address indexed _c, address indexed _from, uint256 _m)
func (_IBank *IBankFilterer) WatchTransferIn(opts *bind.WatchOpts, sink chan<- *IBankTransferIn, _c []common.Address, _from []common.Address) (event.Subscription, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "TransferIn", _cRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankTransferIn)
				if err := _IBank.contract.UnpackLog(event, "TransferIn", log); err != nil {
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

// ParseTransferIn is a log parse operation binding the contract event 0x8ab008cb8444c6d8f2fa3dc0b159b5a86c62b00092219a6c69757805cbf7bde7.
//
// Solidity: event TransferIn(address indexed _c, address indexed _from, uint256 _m)
func (_IBank *IBankFilterer) ParseTransferIn(log types.Log) (*IBankTransferIn, error) {
	event := new(IBankTransferIn)
	if err := _IBank.contract.UnpackLog(event, "TransferIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBankTransferOutIterator is returned from FilterTransferOut and is used to iterate over the raw logs and unpacked data for TransferOut events raised by the IBank contract.
type IBankTransferOutIterator struct {
	Event *IBankTransferOut // Event containing the contract specifics and raw log

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
func (it *IBankTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankTransferOut)
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
		it.Event = new(IBankTransferOut)
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
func (it *IBankTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankTransferOut represents a TransferOut event raised by the IBank contract.
type IBankTransferOut struct {
	C   common.Address
	To  common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransferOut is a free log retrieval operation binding the contract event 0x5d2c285d7e86ec84b7a404ba6d7627ca0a71709acce9c1cc05fada583e3ded81.
//
// Solidity: event TransferOut(address indexed _c, address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) FilterTransferOut(opts *bind.FilterOpts, _c []common.Address, _to []common.Address) (*IBankTransferOutIterator, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "TransferOut", _cRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &IBankTransferOutIterator{contract: _IBank.contract, event: "TransferOut", logs: logs, sub: sub}, nil
}

// WatchTransferOut is a free log subscription operation binding the contract event 0x5d2c285d7e86ec84b7a404ba6d7627ca0a71709acce9c1cc05fada583e3ded81.
//
// Solidity: event TransferOut(address indexed _c, address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) WatchTransferOut(opts *bind.WatchOpts, sink chan<- *IBankTransferOut, _c []common.Address, _to []common.Address) (event.Subscription, error) {

	var _cRule []interface{}
	for _, _cItem := range _c {
		_cRule = append(_cRule, _cItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "TransferOut", _cRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankTransferOut)
				if err := _IBank.contract.UnpackLog(event, "TransferOut", log); err != nil {
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

// ParseTransferOut is a log parse operation binding the contract event 0x5d2c285d7e86ec84b7a404ba6d7627ca0a71709acce9c1cc05fada583e3ded81.
//
// Solidity: event TransferOut(address indexed _c, address indexed _to, uint256 _m)
func (_IBank *IBankFilterer) ParseTransferOut(log types.Log) (*IBankTransferOut, error) {
	event := new(IBankTransferOut)
	if err := _IBank.contract.UnpackLog(event, "TransferOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControlMetaData contains all meta data concerning the IControl contract.
var IControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_ep\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_expire\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_s\",\"type\":\"uint64\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"checkNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IControlMetaData.ABI instead.
var IControlABI = IControlMetaData.ABI

// IControl is an auto generated Go binding around an Ethereum contract.
type IControl struct {
	IControlCaller     // Read-only binding to the contract
	IControlTransactor // Write-only binding to the contract
	IControlFilterer   // Log filterer for contract events
}

// IControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControlSession struct {
	Contract     *IControl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControlCallerSession struct {
	Contract *IControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControlTransactorSession struct {
	Contract     *IControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControlRaw struct {
	Contract *IControl // Generic contract binding to access the raw methods on
}

// IControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControlCallerRaw struct {
	Contract *IControlCaller // Generic read-only contract binding to access the raw methods on
}

// IControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControlTransactorRaw struct {
	Contract *IControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl creates a new instance of IControl, bound to a specific deployed contract.
func NewIControl(address common.Address, backend bind.ContractBackend) (*IControl, error) {
	contract, err := bindIControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl{IControlCaller: IControlCaller{contract: contract}, IControlTransactor: IControlTransactor{contract: contract}, IControlFilterer: IControlFilterer{contract: contract}}, nil
}

// NewIControlCaller creates a new read-only instance of IControl, bound to a specific deployed contract.
func NewIControlCaller(address common.Address, caller bind.ContractCaller) (*IControlCaller, error) {
	contract, err := bindIControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControlCaller{contract: contract}, nil
}

// NewIControlTransactor creates a new write-only instance of IControl, bound to a specific deployed contract.
func NewIControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IControlTransactor, error) {
	contract, err := bindIControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControlTransactor{contract: contract}, nil
}

// NewIControlFilterer creates a new log filterer instance of IControl, bound to a specific deployed contract.
func NewIControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IControlFilterer, error) {
	contract, err := bindIControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControlFilterer{contract: contract}, nil
}

// bindIControl binds a generic wrapper to an already deployed contract.
func bindIControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl *IControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl.Contract.IControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl *IControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl.Contract.IControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl *IControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl.Contract.IControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl *IControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl *IControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl *IControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0xced433d5.
//
// Solidity: function add(address _a, uint64 _ep, uint64 _expire, uint64 _s) returns()
func (_IControl *IControlTransactor) Add(opts *bind.TransactOpts, _a common.Address, _ep uint64, _expire uint64, _s uint64) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "add", _a, _ep, _expire, _s)
}

// Add is a paid mutator transaction binding the contract method 0xced433d5.
//
// Solidity: function add(address _a, uint64 _ep, uint64 _expire, uint64 _s) returns()
func (_IControl *IControlSession) Add(_a common.Address, _ep uint64, _expire uint64, _s uint64) (*types.Transaction, error) {
	return _IControl.Contract.Add(&_IControl.TransactOpts, _a, _ep, _expire, _s)
}

// Add is a paid mutator transaction binding the contract method 0xced433d5.
//
// Solidity: function add(address _a, uint64 _ep, uint64 _expire, uint64 _s) returns()
func (_IControl *IControlTransactorSession) Add(_a common.Address, _ep uint64, _expire uint64, _s uint64) (*types.Transaction, error) {
	return _IControl.Contract.Add(&_IControl.TransactOpts, _a, _ep, _expire, _s)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_IControl *IControlTransactor) CheckNode(opts *bind.TransactOpts, _a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "checkNode", _a, _money)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_IControl *IControlSession) CheckNode(_a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.CheckNode(&_IControl.TransactOpts, _a, _money)
}

// CheckNode is a paid mutator transaction binding the contract method 0xa3f60869.
//
// Solidity: function checkNode(address _a, uint256 _money) returns()
func (_IControl *IControlTransactorSession) CheckNode(_a common.Address, _money *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.CheckNode(&_IControl.TransactOpts, _a, _money)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactor) Lock(opts *bind.TransactOpts, _a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "lock", _a, _typ, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlSession) Lock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Lock(&_IControl.TransactOpts, _a, _typ, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Lock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Lock(&_IControl.TransactOpts, _a, _typ, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) returns()
func (_IControl *IControlTransactor) Mint(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "mint", _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) returns()
func (_IControl *IControlSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Mint(&_IControl.TransactOpts, _a, _m)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _a, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Mint(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Mint(&_IControl.TransactOpts, _a, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _typ, address _to, uint256 _m) returns()
func (_IControl *IControlTransactor) Punish(opts *bind.TransactOpts, _a common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "punish", _a, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _typ, address _to, uint256 _m) returns()
func (_IControl *IControlSession) Punish(_a common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Punish(&_IControl.TransactOpts, _a, _typ, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _typ, address _to, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Punish(_a common.Address, _typ uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Punish(&_IControl.TransactOpts, _a, _typ, _to, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IControl *IControlTransactor) Release(opts *bind.TransactOpts, _a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "release", _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IControl *IControlSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Release(&_IControl.TransactOpts, _a, _m)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _a, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Release(_a common.Address, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Release(&_IControl.TransactOpts, _a, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactor) Unlock(opts *bind.TransactOpts, _a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "unlock", _a, _typ, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlSession) Unlock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Unlock(&_IControl.TransactOpts, _a, _typ, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _typ, uint256 _m) returns()
func (_IControl *IControlTransactorSession) Unlock(_a common.Address, _typ uint8, _m *big.Int) (*types.Transaction, error) {
	return _IControl.Contract.Unlock(&_IControl.TransactOpts, _a, _typ, _m)
}

// IEpochMetaData contains all meta data concerning the IEpoch contract.
var IEpochMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"SetEpoch\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"}],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEpochABI is the input ABI used to generate the binding from.
// Deprecated: Use IEpochMetaData.ABI instead.
var IEpochABI = IEpochMetaData.ABI

// IEpoch is an auto generated Go binding around an Ethereum contract.
type IEpoch struct {
	IEpochCaller     // Read-only binding to the contract
	IEpochTransactor // Write-only binding to the contract
	IEpochFilterer   // Log filterer for contract events
}

// IEpochCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEpochCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEpochTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEpochFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEpochSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEpochSession struct {
	Contract     *IEpoch           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEpochCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEpochCallerSession struct {
	Contract *IEpochCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IEpochTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEpochTransactorSession struct {
	Contract     *IEpochTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEpochRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEpochRaw struct {
	Contract *IEpoch // Generic contract binding to access the raw methods on
}

// IEpochCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEpochCallerRaw struct {
	Contract *IEpochCaller // Generic read-only contract binding to access the raw methods on
}

// IEpochTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEpochTransactorRaw struct {
	Contract *IEpochTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEpoch creates a new instance of IEpoch, bound to a specific deployed contract.
func NewIEpoch(address common.Address, backend bind.ContractBackend) (*IEpoch, error) {
	contract, err := bindIEpoch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEpoch{IEpochCaller: IEpochCaller{contract: contract}, IEpochTransactor: IEpochTransactor{contract: contract}, IEpochFilterer: IEpochFilterer{contract: contract}}, nil
}

// NewIEpochCaller creates a new read-only instance of IEpoch, bound to a specific deployed contract.
func NewIEpochCaller(address common.Address, caller bind.ContractCaller) (*IEpochCaller, error) {
	contract, err := bindIEpoch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEpochCaller{contract: contract}, nil
}

// NewIEpochTransactor creates a new write-only instance of IEpoch, bound to a specific deployed contract.
func NewIEpochTransactor(address common.Address, transactor bind.ContractTransactor) (*IEpochTransactor, error) {
	contract, err := bindIEpoch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEpochTransactor{contract: contract}, nil
}

// NewIEpochFilterer creates a new log filterer instance of IEpoch, bound to a specific deployed contract.
func NewIEpochFilterer(address common.Address, filterer bind.ContractFilterer) (*IEpochFilterer, error) {
	contract, err := bindIEpoch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEpochFilterer{contract: contract}, nil
}

// bindIEpoch binds a generic wrapper to an already deployed contract.
func bindIEpoch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEpochMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEpoch *IEpochRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEpoch.Contract.IEpochCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEpoch *IEpochRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.Contract.IEpochTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEpoch *IEpochRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEpoch.Contract.IEpochTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEpoch *IEpochCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEpoch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEpoch *IEpochTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEpoch *IEpochTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEpoch.Contract.contract.Transact(opts, method, params...)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_IEpoch *IEpochTransactor) Check(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.contract.Transact(opts, "check")
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_IEpoch *IEpochSession) Check() (*types.Transaction, error) {
	return _IEpoch.Contract.Check(&_IEpoch.TransactOpts)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_IEpoch *IEpochTransactorSession) Check() (*types.Transaction, error) {
	return _IEpoch.Contract.Check(&_IEpoch.TransactOpts)
}

// Current is a paid mutator transaction binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() returns(uint64)
func (_IEpoch *IEpochTransactor) Current(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEpoch.contract.Transact(opts, "current")
}

// Current is a paid mutator transaction binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() returns(uint64)
func (_IEpoch *IEpochSession) Current() (*types.Transaction, error) {
	return _IEpoch.Contract.Current(&_IEpoch.TransactOpts)
}

// Current is a paid mutator transaction binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() returns(uint64)
func (_IEpoch *IEpochTransactorSession) Current() (*types.Transaction, error) {
	return _IEpoch.Contract.Current(&_IEpoch.TransactOpts)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) returns(uint256, bytes32)
func (_IEpoch *IEpochTransactor) GetEpoch(opts *bind.TransactOpts, _epoch uint64) (*types.Transaction, error) {
	return _IEpoch.contract.Transact(opts, "getEpoch", _epoch)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) returns(uint256, bytes32)
func (_IEpoch *IEpochSession) GetEpoch(_epoch uint64) (*types.Transaction, error) {
	return _IEpoch.Contract.GetEpoch(&_IEpoch.TransactOpts, _epoch)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) returns(uint256, bytes32)
func (_IEpoch *IEpochTransactorSession) GetEpoch(_epoch uint64) (*types.Transaction, error) {
	return _IEpoch.Contract.GetEpoch(&_IEpoch.TransactOpts, _epoch)
}

// IEpochSetEpochIterator is returned from FilterSetEpoch and is used to iterate over the raw logs and unpacked data for SetEpoch events raised by the IEpoch contract.
type IEpochSetEpochIterator struct {
	Event *IEpochSetEpoch // Event containing the contract specifics and raw log

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
func (it *IEpochSetEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEpochSetEpoch)
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
		it.Event = new(IEpochSetEpoch)
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
func (it *IEpochSetEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEpochSetEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEpochSetEpoch represents a SetEpoch event raised by the IEpoch contract.
type IEpochSetEpoch struct {
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetEpoch is a free log retrieval operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_IEpoch *IEpochFilterer) FilterSetEpoch(opts *bind.FilterOpts) (*IEpochSetEpochIterator, error) {

	logs, sub, err := _IEpoch.contract.FilterLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return &IEpochSetEpochIterator{contract: _IEpoch.contract, event: "SetEpoch", logs: logs, sub: sub}, nil
}

// WatchSetEpoch is a free log subscription operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_IEpoch *IEpochFilterer) WatchSetEpoch(opts *bind.WatchOpts, sink chan<- *IEpochSetEpoch) (event.Subscription, error) {

	logs, sub, err := _IEpoch.contract.WatchLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEpochSetEpoch)
				if err := _IEpoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
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

// ParseSetEpoch is a log parse operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_IEpoch *IEpochFilterer) ParseSetEpoch(log types.Log) (*IEpochSetEpoch, error) {
	event := new(IEpochSetEpoch)
	if err := _IEpoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGPUMetaData contains all meta data concerning the IGPU contract.
var IGPUMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_gn\",\"type\":\"string\"}],\"name\":\"AddGPU\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"}],\"name\":\"rent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGPUABI is the input ABI used to generate the binding from.
// Deprecated: Use IGPUMetaData.ABI instead.
var IGPUABI = IGPUMetaData.ABI

// IGPU is an auto generated Go binding around an Ethereum contract.
type IGPU struct {
	IGPUCaller     // Read-only binding to the contract
	IGPUTransactor // Write-only binding to the contract
	IGPUFilterer   // Log filterer for contract events
}

// IGPUCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGPUCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGPUTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGPUTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGPUFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGPUFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGPUSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGPUSession struct {
	Contract     *IGPU             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGPUCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGPUCallerSession struct {
	Contract *IGPUCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGPUTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGPUTransactorSession struct {
	Contract     *IGPUTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGPURaw is an auto generated low-level Go binding around an Ethereum contract.
type IGPURaw struct {
	Contract *IGPU // Generic contract binding to access the raw methods on
}

// IGPUCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGPUCallerRaw struct {
	Contract *IGPUCaller // Generic read-only contract binding to access the raw methods on
}

// IGPUTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGPUTransactorRaw struct {
	Contract *IGPUTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGPU creates a new instance of IGPU, bound to a specific deployed contract.
func NewIGPU(address common.Address, backend bind.ContractBackend) (*IGPU, error) {
	contract, err := bindIGPU(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGPU{IGPUCaller: IGPUCaller{contract: contract}, IGPUTransactor: IGPUTransactor{contract: contract}, IGPUFilterer: IGPUFilterer{contract: contract}}, nil
}

// NewIGPUCaller creates a new read-only instance of IGPU, bound to a specific deployed contract.
func NewIGPUCaller(address common.Address, caller bind.ContractCaller) (*IGPUCaller, error) {
	contract, err := bindIGPU(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGPUCaller{contract: contract}, nil
}

// NewIGPUTransactor creates a new write-only instance of IGPU, bound to a specific deployed contract.
func NewIGPUTransactor(address common.Address, transactor bind.ContractTransactor) (*IGPUTransactor, error) {
	contract, err := bindIGPU(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGPUTransactor{contract: contract}, nil
}

// NewIGPUFilterer creates a new log filterer instance of IGPU, bound to a specific deployed contract.
func NewIGPUFilterer(address common.Address, filterer bind.ContractFilterer) (*IGPUFilterer, error) {
	contract, err := bindIGPU(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGPUFilterer{contract: contract}, nil
}

// bindIGPU binds a generic wrapper to an already deployed contract.
func bindIGPU(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGPUMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGPU *IGPURaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGPU.Contract.IGPUCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGPU *IGPURaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGPU.Contract.IGPUTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGPU *IGPURaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGPU.Contract.IGPUTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGPU *IGPUCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGPU.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGPU *IGPUTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGPU.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGPU *IGPUTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGPU.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a paid mutator transaction binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _gi) returns(address)
func (_IGPU *IGPUTransactor) GetOwner(opts *bind.TransactOpts, _gi uint64) (*types.Transaction, error) {
	return _IGPU.contract.Transact(opts, "getOwner", _gi)
}

// GetOwner is a paid mutator transaction binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _gi) returns(address)
func (_IGPU *IGPUSession) GetOwner(_gi uint64) (*types.Transaction, error) {
	return _IGPU.Contract.GetOwner(&_IGPU.TransactOpts, _gi)
}

// GetOwner is a paid mutator transaction binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _gi) returns(address)
func (_IGPU *IGPUTransactorSession) GetOwner(_gi uint64) (*types.Transaction, error) {
	return _IGPU.Contract.GetOwner(&_IGPU.TransactOpts, _gi)
}

// Release is a paid mutator transaction binding the contract method 0x41fbbc31.
//
// Solidity: function release(uint64 _gi) returns()
func (_IGPU *IGPUTransactor) Release(opts *bind.TransactOpts, _gi uint64) (*types.Transaction, error) {
	return _IGPU.contract.Transact(opts, "release", _gi)
}

// Release is a paid mutator transaction binding the contract method 0x41fbbc31.
//
// Solidity: function release(uint64 _gi) returns()
func (_IGPU *IGPUSession) Release(_gi uint64) (*types.Transaction, error) {
	return _IGPU.Contract.Release(&_IGPU.TransactOpts, _gi)
}

// Release is a paid mutator transaction binding the contract method 0x41fbbc31.
//
// Solidity: function release(uint64 _gi) returns()
func (_IGPU *IGPUTransactorSession) Release(_gi uint64) (*types.Transaction, error) {
	return _IGPU.Contract.Release(&_IGPU.TransactOpts, _gi)
}

// Rent is a paid mutator transaction binding the contract method 0x71d49a63.
//
// Solidity: function rent(uint64 _gi) returns()
func (_IGPU *IGPUTransactor) Rent(opts *bind.TransactOpts, _gi uint64) (*types.Transaction, error) {
	return _IGPU.contract.Transact(opts, "rent", _gi)
}

// Rent is a paid mutator transaction binding the contract method 0x71d49a63.
//
// Solidity: function rent(uint64 _gi) returns()
func (_IGPU *IGPUSession) Rent(_gi uint64) (*types.Transaction, error) {
	return _IGPU.Contract.Rent(&_IGPU.TransactOpts, _gi)
}

// Rent is a paid mutator transaction binding the contract method 0x71d49a63.
//
// Solidity: function rent(uint64 _gi) returns()
func (_IGPU *IGPUTransactorSession) Rent(_gi uint64) (*types.Transaction, error) {
	return _IGPU.Contract.Rent(&_IGPU.TransactOpts, _gi)
}

// IGPUAddGPUIterator is returned from FilterAddGPU and is used to iterate over the raw logs and unpacked data for AddGPU events raised by the IGPU contract.
type IGPUAddGPUIterator struct {
	Event *IGPUAddGPU // Event containing the contract specifics and raw log

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
func (it *IGPUAddGPUIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGPUAddGPU)
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
		it.Event = new(IGPUAddGPU)
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
func (it *IGPUAddGPUIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGPUAddGPUIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGPUAddGPU represents a AddGPU event raised by the IGPU contract.
type IGPUAddGPU struct {
	A   common.Address
	Gi  uint64
	Gn  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddGPU is a free log retrieval operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _gi, string _gn)
func (_IGPU *IGPUFilterer) FilterAddGPU(opts *bind.FilterOpts, _a []common.Address) (*IGPUAddGPUIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IGPU.contract.FilterLogs(opts, "AddGPU", _aRule)
	if err != nil {
		return nil, err
	}
	return &IGPUAddGPUIterator{contract: _IGPU.contract, event: "AddGPU", logs: logs, sub: sub}, nil
}

// WatchAddGPU is a free log subscription operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _gi, string _gn)
func (_IGPU *IGPUFilterer) WatchAddGPU(opts *bind.WatchOpts, sink chan<- *IGPUAddGPU, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IGPU.contract.WatchLogs(opts, "AddGPU", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGPUAddGPU)
				if err := _IGPU.contract.UnpackLog(event, "AddGPU", log); err != nil {
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

// ParseAddGPU is a log parse operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _gi, string _gn)
func (_IGPU *IGPUFilterer) ParseAddGPU(log types.Log) (*IGPUAddGPU, error) {
	event := new(IGPUAddGPU)
	if err := _IGPU.contract.UnpackLog(event, "AddGPU", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IModelMetaData contains all meta data concerning the IModel contract.
var IModelMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"AddModel\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IModelABI is the input ABI used to generate the binding from.
// Deprecated: Use IModelMetaData.ABI instead.
var IModelABI = IModelMetaData.ABI

// IModel is an auto generated Go binding around an Ethereum contract.
type IModel struct {
	IModelCaller     // Read-only binding to the contract
	IModelTransactor // Write-only binding to the contract
	IModelFilterer   // Log filterer for contract events
}

// IModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IModelSession struct {
	Contract     *IModel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IModelCallerSession struct {
	Contract *IModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IModelTransactorSession struct {
	Contract     *IModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IModelRaw struct {
	Contract *IModel // Generic contract binding to access the raw methods on
}

// IModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IModelCallerRaw struct {
	Contract *IModelCaller // Generic read-only contract binding to access the raw methods on
}

// IModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IModelTransactorRaw struct {
	Contract *IModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIModel creates a new instance of IModel, bound to a specific deployed contract.
func NewIModel(address common.Address, backend bind.ContractBackend) (*IModel, error) {
	contract, err := bindIModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IModel{IModelCaller: IModelCaller{contract: contract}, IModelTransactor: IModelTransactor{contract: contract}, IModelFilterer: IModelFilterer{contract: contract}}, nil
}

// NewIModelCaller creates a new read-only instance of IModel, bound to a specific deployed contract.
func NewIModelCaller(address common.Address, caller bind.ContractCaller) (*IModelCaller, error) {
	contract, err := bindIModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IModelCaller{contract: contract}, nil
}

// NewIModelTransactor creates a new write-only instance of IModel, bound to a specific deployed contract.
func NewIModelTransactor(address common.Address, transactor bind.ContractTransactor) (*IModelTransactor, error) {
	contract, err := bindIModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IModelTransactor{contract: contract}, nil
}

// NewIModelFilterer creates a new log filterer instance of IModel, bound to a specific deployed contract.
func NewIModelFilterer(address common.Address, filterer bind.ContractFilterer) (*IModelFilterer, error) {
	contract, err := bindIModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IModelFilterer{contract: contract}, nil
}

// bindIModel binds a generic wrapper to an already deployed contract.
func bindIModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IModel *IModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IModel.Contract.IModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IModel *IModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IModel.Contract.IModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IModel *IModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IModel.Contract.IModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IModel *IModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IModel *IModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IModel *IModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IModel.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_IModel *IModelCaller) GetOwner(opts *bind.CallOpts, _mi uint64) (common.Address, error) {
	var out []interface{}
	err := _IModel.contract.Call(opts, &out, "getOwner", _mi)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_IModel *IModelSession) GetOwner(_mi uint64) (common.Address, error) {
	return _IModel.Contract.GetOwner(&_IModel.CallOpts, _mi)
}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_IModel *IModelCallerSession) GetOwner(_mi uint64) (common.Address, error) {
	return _IModel.Contract.GetOwner(&_IModel.CallOpts, _mi)
}

// Check is a paid mutator transaction binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) returns()
func (_IModel *IModelTransactor) Check(opts *bind.TransactOpts, _mi uint64) (*types.Transaction, error) {
	return _IModel.contract.Transact(opts, "check", _mi)
}

// Check is a paid mutator transaction binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) returns()
func (_IModel *IModelSession) Check(_mi uint64) (*types.Transaction, error) {
	return _IModel.Contract.Check(&_IModel.TransactOpts, _mi)
}

// Check is a paid mutator transaction binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) returns()
func (_IModel *IModelTransactorSession) Check(_mi uint64) (*types.Transaction, error) {
	return _IModel.Contract.Check(&_IModel.TransactOpts, _mi)
}

// IModelAddModelIterator is returned from FilterAddModel and is used to iterate over the raw logs and unpacked data for AddModel events raised by the IModel contract.
type IModelAddModelIterator struct {
	Event *IModelAddModel // Event containing the contract specifics and raw log

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
func (it *IModelAddModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IModelAddModel)
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
		it.Event = new(IModelAddModel)
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
func (it *IModelAddModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IModelAddModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IModelAddModel represents a AddModel event raised by the IModel contract.
type IModelAddModel struct {
	A   common.Address
	Mi  uint64
	Mn  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddModel is a free log retrieval operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_IModel *IModelFilterer) FilterAddModel(opts *bind.FilterOpts, _a []common.Address) (*IModelAddModelIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IModel.contract.FilterLogs(opts, "AddModel", _aRule)
	if err != nil {
		return nil, err
	}
	return &IModelAddModelIterator{contract: _IModel.contract, event: "AddModel", logs: logs, sub: sub}, nil
}

// WatchAddModel is a free log subscription operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_IModel *IModelFilterer) WatchAddModel(opts *bind.WatchOpts, sink chan<- *IModelAddModel, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _IModel.contract.WatchLogs(opts, "AddModel", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IModelAddModel)
				if err := _IModel.contract.UnpackLog(event, "AddModel", log); err != nil {
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

// ParseAddModel is a log parse operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_IModel *IModelFilterer) ParseAddModel(log types.Log) (*IModelAddModel, error) {
	event := new(IModelAddModel)
	if err := _IModel.contract.UnpackLog(event, "AddModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodeMetaData contains all meta data concerning the INode contract.
var INodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Punish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INodeABI is the input ABI used to generate the binding from.
// Deprecated: Use INodeMetaData.ABI instead.
var INodeABI = INodeMetaData.ABI

// INode is an auto generated Go binding around an Ethereum contract.
type INode struct {
	INodeCaller     // Read-only binding to the contract
	INodeTransactor // Write-only binding to the contract
	INodeFilterer   // Log filterer for contract events
}

// INodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type INodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INodeSession struct {
	Contract     *INode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INodeCallerSession struct {
	Contract *INodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// INodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INodeTransactorSession struct {
	Contract     *INodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type INodeRaw struct {
	Contract *INode // Generic contract binding to access the raw methods on
}

// INodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INodeCallerRaw struct {
	Contract *INodeCaller // Generic read-only contract binding to access the raw methods on
}

// INodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INodeTransactorRaw struct {
	Contract *INodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINode creates a new instance of INode, bound to a specific deployed contract.
func NewINode(address common.Address, backend bind.ContractBackend) (*INode, error) {
	contract, err := bindINode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INode{INodeCaller: INodeCaller{contract: contract}, INodeTransactor: INodeTransactor{contract: contract}, INodeFilterer: INodeFilterer{contract: contract}}, nil
}

// NewINodeCaller creates a new read-only instance of INode, bound to a specific deployed contract.
func NewINodeCaller(address common.Address, caller bind.ContractCaller) (*INodeCaller, error) {
	contract, err := bindINode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INodeCaller{contract: contract}, nil
}

// NewINodeTransactor creates a new write-only instance of INode, bound to a specific deployed contract.
func NewINodeTransactor(address common.Address, transactor bind.ContractTransactor) (*INodeTransactor, error) {
	contract, err := bindINode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INodeTransactor{contract: contract}, nil
}

// NewINodeFilterer creates a new log filterer instance of INode, bound to a specific deployed contract.
func NewINodeFilterer(address common.Address, filterer bind.ContractFilterer) (*INodeFilterer, error) {
	contract, err := bindINode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INodeFilterer{contract: contract}, nil
}

// bindINode binds a generic wrapper to an already deployed contract.
func bindINode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.INodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(uint256)
func (_INode *INodeCaller) Check(opts *bind.CallOpts, _a common.Address, _type uint8) (*big.Int, error) {
	var out []interface{}
	err := _INode.contract.Call(opts, &out, "check", _a, _type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(uint256)
func (_INode *INodeSession) Check(_a common.Address, _type uint8) (*big.Int, error) {
	return _INode.Contract.Check(&_INode.CallOpts, _a, _type)
}

// Check is a free data retrieval call binding the contract method 0x61e728b1.
//
// Solidity: function check(address _a, uint8 _type) view returns(uint256)
func (_INode *INodeCallerSession) Check(_a common.Address, _type uint8) (*big.Int, error) {
	return _INode.Contract.Check(&_INode.CallOpts, _a, _type)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactor) Lock(opts *bind.TransactOpts, _a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "lock", _a, _type, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeSession) Lock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Lock(&_INode.TransactOpts, _a, _type, _m)
}

// Lock is a paid mutator transaction binding the contract method 0x738dddba.
//
// Solidity: function lock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactorSession) Lock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Lock(&_INode.TransactOpts, _a, _type, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _type, address _to, uint256 _m) returns()
func (_INode *INodeTransactor) Punish(opts *bind.TransactOpts, _a common.Address, _type uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "punish", _a, _type, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _type, address _to, uint256 _m) returns()
func (_INode *INodeSession) Punish(_a common.Address, _type uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Punish(&_INode.TransactOpts, _a, _type, _to, _m)
}

// Punish is a paid mutator transaction binding the contract method 0x06f0b4f1.
//
// Solidity: function punish(address _a, uint8 _type, address _to, uint256 _m) returns()
func (_INode *INodeTransactorSession) Punish(_a common.Address, _type uint8, _to common.Address, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Punish(&_INode.TransactOpts, _a, _type, _to, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactor) Unlock(opts *bind.TransactOpts, _a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.contract.Transact(opts, "unlock", _a, _type, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeSession) Unlock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Unlock(&_INode.TransactOpts, _a, _type, _m)
}

// Unlock is a paid mutator transaction binding the contract method 0x2165e20b.
//
// Solidity: function unlock(address _a, uint8 _type, uint256 _m) returns()
func (_INode *INodeTransactorSession) Unlock(_a common.Address, _type uint8, _m *big.Int) (*types.Transaction, error) {
	return _INode.Contract.Unlock(&_INode.TransactOpts, _a, _type, _m)
}

// INodePledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the INode contract.
type INodePledgeIterator struct {
	Event *INodePledge // Event containing the contract specifics and raw log

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
func (it *INodePledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodePledge)
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
		it.Event = new(INodePledge)
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
func (it *INodePledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodePledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodePledge represents a Pledge event raised by the INode contract.
type INodePledge struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) FilterPledge(opts *bind.FilterOpts, _a []common.Address) (*INodePledgeIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.FilterLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return &INodePledgeIterator{contract: _INode.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *INodePledge, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.WatchLogs(opts, "Pledge", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodePledge)
				if err := _INode.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0xedec38b4b62433445c926815411650aa7d417efa9da307f6aa99e69e6f4493ee.
//
// Solidity: event Pledge(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) ParsePledge(log types.Log) (*INodePledge, error) {
	event := new(INodePledge)
	if err := _INode.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodePunishIterator is returned from FilterPunish and is used to iterate over the raw logs and unpacked data for Punish events raised by the INode contract.
type INodePunishIterator struct {
	Event *INodePunish // Event containing the contract specifics and raw log

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
func (it *INodePunishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodePunish)
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
		it.Event = new(INodePunish)
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
func (it *INodePunishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodePunishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodePunish represents a Punish event raised by the INode contract.
type INodePunish struct {
	A     common.Address
	Typ   uint8
	To    common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPunish is a free log retrieval operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_INode *INodeFilterer) FilterPunish(opts *bind.FilterOpts, _a []common.Address, _to []common.Address) (*INodePunishIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _INode.contract.FilterLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &INodePunishIterator{contract: _INode.contract, event: "Punish", logs: logs, sub: sub}, nil
}

// WatchPunish is a free log subscription operation binding the contract event 0xc6ef96923e613455515c6723eff1723445b22427fe442e8bf742e9d29b4b3c32.
//
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_INode *INodeFilterer) WatchPunish(opts *bind.WatchOpts, sink chan<- *INodePunish, _a []common.Address, _to []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _INode.contract.WatchLogs(opts, "Punish", _aRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodePunish)
				if err := _INode.contract.UnpackLog(event, "Punish", log); err != nil {
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
// Solidity: event Punish(address indexed _a, uint8 _typ, address indexed _to, uint256 _money)
func (_INode *INodeFilterer) ParsePunish(log types.Log) (*INodePunish, error) {
	event := new(INodePunish)
	if err := _INode.contract.UnpackLog(event, "Punish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodeSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the INode contract.
type INodeSetIterator struct {
	Event *INodeSet // Event containing the contract specifics and raw log

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
func (it *INodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodeSet)
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
		it.Event = new(INodeSet)
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
func (it *INodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodeSet represents a Set event raised by the INode contract.
type INodeSet struct {
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_INode *INodeFilterer) FilterSet(opts *bind.FilterOpts) (*INodeSetIterator, error) {

	logs, sub, err := _INode.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &INodeSetIterator{contract: _INode.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_INode *INodeFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *INodeSet) (event.Subscription, error) {

	logs, sub, err := _INode.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodeSet)
				if err := _INode.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xc4b70ab905e9fd7aab427fb9e73cae1480cfdc41c22053b20745349a7ef67881.
//
// Solidity: event Set(uint8 _type, uint256 _money)
func (_INode *INodeFilterer) ParseSet(log types.Log) (*INodeSet, error) {
	event := new(INodeSet)
	if err := _INode.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INodeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the INode contract.
type INodeWithdrawIterator struct {
	Event *INodeWithdraw // Event containing the contract specifics and raw log

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
func (it *INodeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INodeWithdraw)
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
		it.Event = new(INodeWithdraw)
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
func (it *INodeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INodeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INodeWithdraw represents a Withdraw event raised by the INode contract.
type INodeWithdraw struct {
	A     common.Address
	Type  uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*INodeWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &INodeWithdrawIterator{contract: _INode.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *INodeWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _INode.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INodeWithdraw)
				if err := _INode.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe3054ec6f352b09a86839beeeceb27ed4684f9164eb61f4a5d6d159ee8789d74.
//
// Solidity: event Withdraw(address indexed _a, uint8 _type, uint256 _money)
func (_INode *INodeFilterer) ParseWithdraw(log types.Log) (*INodeWithdraw, error) {
	event := new(INodeWithdraw)
	if err := _INode.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISpaceMetaData contains all meta data concerning the ISpace contract.
var ISpaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ai\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_an\",\"type\":\"string\"}],\"name\":\"AddGPU\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]",
}

// ISpaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ISpaceMetaData.ABI instead.
var ISpaceABI = ISpaceMetaData.ABI

// ISpace is an auto generated Go binding around an Ethereum contract.
type ISpace struct {
	ISpaceCaller     // Read-only binding to the contract
	ISpaceTransactor // Write-only binding to the contract
	ISpaceFilterer   // Log filterer for contract events
}

// ISpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISpaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISpaceSession struct {
	Contract     *ISpace           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISpaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISpaceCallerSession struct {
	Contract *ISpaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISpaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISpaceTransactorSession struct {
	Contract     *ISpaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISpaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISpaceRaw struct {
	Contract *ISpace // Generic contract binding to access the raw methods on
}

// ISpaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISpaceCallerRaw struct {
	Contract *ISpaceCaller // Generic read-only contract binding to access the raw methods on
}

// ISpaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISpaceTransactorRaw struct {
	Contract *ISpaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISpace creates a new instance of ISpace, bound to a specific deployed contract.
func NewISpace(address common.Address, backend bind.ContractBackend) (*ISpace, error) {
	contract, err := bindISpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISpace{ISpaceCaller: ISpaceCaller{contract: contract}, ISpaceTransactor: ISpaceTransactor{contract: contract}, ISpaceFilterer: ISpaceFilterer{contract: contract}}, nil
}

// NewISpaceCaller creates a new read-only instance of ISpace, bound to a specific deployed contract.
func NewISpaceCaller(address common.Address, caller bind.ContractCaller) (*ISpaceCaller, error) {
	contract, err := bindISpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISpaceCaller{contract: contract}, nil
}

// NewISpaceTransactor creates a new write-only instance of ISpace, bound to a specific deployed contract.
func NewISpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ISpaceTransactor, error) {
	contract, err := bindISpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISpaceTransactor{contract: contract}, nil
}

// NewISpaceFilterer creates a new log filterer instance of ISpace, bound to a specific deployed contract.
func NewISpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ISpaceFilterer, error) {
	contract, err := bindISpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISpaceFilterer{contract: contract}, nil
}

// bindISpace binds a generic wrapper to an already deployed contract.
func bindISpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISpaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISpace *ISpaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISpace.Contract.ISpaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISpace *ISpaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISpace.Contract.ISpaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISpace *ISpaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISpace.Contract.ISpaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISpace *ISpaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISpace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISpace *ISpaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISpace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISpace *ISpaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISpace.Contract.contract.Transact(opts, method, params...)
}

// ISpaceAddGPUIterator is returned from FilterAddGPU and is used to iterate over the raw logs and unpacked data for AddGPU events raised by the ISpace contract.
type ISpaceAddGPUIterator struct {
	Event *ISpaceAddGPU // Event containing the contract specifics and raw log

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
func (it *ISpaceAddGPUIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISpaceAddGPU)
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
		it.Event = new(ISpaceAddGPU)
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
func (it *ISpaceAddGPUIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISpaceAddGPUIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISpaceAddGPU represents a AddGPU event raised by the ISpace contract.
type ISpaceAddGPU struct {
	A   common.Address
	Ai  uint64
	An  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddGPU is a free log retrieval operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _ai, string _an)
func (_ISpace *ISpaceFilterer) FilterAddGPU(opts *bind.FilterOpts, _a []common.Address) (*ISpaceAddGPUIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _ISpace.contract.FilterLogs(opts, "AddGPU", _aRule)
	if err != nil {
		return nil, err
	}
	return &ISpaceAddGPUIterator{contract: _ISpace.contract, event: "AddGPU", logs: logs, sub: sub}, nil
}

// WatchAddGPU is a free log subscription operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _ai, string _an)
func (_ISpace *ISpaceFilterer) WatchAddGPU(opts *bind.WatchOpts, sink chan<- *ISpaceAddGPU, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _ISpace.contract.WatchLogs(opts, "AddGPU", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISpaceAddGPU)
				if err := _ISpace.contract.UnpackLog(event, "AddGPU", log); err != nil {
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

// ParseAddGPU is a log parse operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _ai, string _an)
func (_ISpace *ISpaceFilterer) ParseAddGPU(log types.Log) (*ISpaceAddGPU, error) {
	event := new(ISpaceAddGPU)
	if err := _ISpace.contract.UnpackLog(event, "AddGPU", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISpaceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ISpace contract.
type ISpaceWithdrawIterator struct {
	Event *ISpaceWithdraw // Event containing the contract specifics and raw log

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
func (it *ISpaceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISpaceWithdraw)
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
		it.Event = new(ISpaceWithdraw)
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
func (it *ISpaceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISpaceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISpaceWithdraw represents a Withdraw event raised by the ISpace contract.
type ISpaceWithdraw struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_ISpace *ISpaceFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*ISpaceWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _ISpace.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &ISpaceWithdrawIterator{contract: _ISpace.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_ISpace *ISpaceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ISpaceWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _ISpace.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISpaceWithdraw)
				if err := _ISpace.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_ISpace *ISpaceFilterer) ParseWithdraw(log types.Log) (*ISpaceWithdraw, error) {
	event := new(ISpaceWithdraw)
	if err := _ISpace.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpaceMetaData contains all meta data concerning the Space contract.
var SpaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ai\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_an\",\"type\":\"string\"}],\"name\":\"AddGPU\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_si\",\"type\":\"uint64\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sn\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_p\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sn\",\"type\":\"string\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_si\",\"type\":\"uint64\"}],\"name\":\"getSpace\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"model\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expire\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"fin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structISpace.Info\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minRent\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_si\",\"type\":\"uint64\"}],\"name\":\"shutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_m\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052601e60015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015610038575f80fd5b50604051613853380380613853833981810160405281019061005a9190610317565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506100a1610233565b600281908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f0190816100dd919061057c565b506020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816002015f6101000a81548160ff02191690831515021790555060c08201518160020160016101000a81548160ff02191690831515021790555060e08201518160020160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061010082015181600301555050505061064b565b604051806101200160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81525090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102e6826102bd565b9050919050565b6102f6816102dc565b8114610300575f80fd5b50565b5f81519050610311816102ed565b92915050565b5f6020828403121561032c5761032b6102b9565b5b5f61033984828501610303565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103bd57607f821691505b6020821081036103d0576103cf610379565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104327fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103f7565b61043c86836103f7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61048061047b61047684610454565b61045d565b610454565b9050919050565b5f819050919050565b61049983610466565b6104ad6104a582610487565b848454610403565b825550505050565b5f90565b6104c16104b5565b6104cc818484610490565b505050565b5b818110156104ef576104e45f826104b9565b6001810190506104d2565b5050565b601f82111561053457610505816103d6565b61050e846103e8565b8101602085101561051d578190505b610531610529856103e8565b8301826104d1565b50505b505050565b5f82821c905092915050565b5f6105545f1984600802610539565b1980831691505092915050565b5f61056c8383610545565b9150826002028217905092915050565b61058582610342565b67ffffffffffffffff81111561059e5761059d61034c565b5b6105a882546103a6565b6105b38282856104f3565b5f60209050601f8311600181146105e4575f84156105d2578287015190505b6105dc8582610561565b865550610643565b601f1984166105f2866103d6565b5f5b82811015610619578489015182556001820191506020850194506020810190506105f4565b868310156106365784890151610632601f891682610545565b8355505b6001600288020188555050505b505050505050565b6131fb806106585f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c806399760d121161006457806399760d12146101565780639fa6a6e314610172578063e38c611614610190578063ea1bbe35146101ae578063fc7e4746146101de5761009c565b80631b1683f7146100a05780632e1a7d4d146100bc57806370a08231146100d857806376cdb03b1461010857806391fb07c414610126575b5f80fd5b6100ba60048036038101906100b5919061215b565b6101fa565b005b6100d660048036038101906100d191906121b9565b610c02565b005b6100f260048036038101906100ed919061223e565b610e8e565b6040516100ff9190612278565b60405180910390f35b610110610ed4565b60405161011d91906122a0565b60405180910390f35b610140600480360381019061013b919061215b565b610ef7565b60405161014d919061244c565b60405180910390f35b610170600480360381019061016b9190612598565b61122a565b005b61017a611a9e565b604051610187919061263a565b60405180910390f35b610198611ab7565b6040516101a5919061263a565b60405180910390f35b6101c860048036038101906101c39190612653565b611ad0565b6040516101d5919061263a565b60405180910390f35b6101f860048036038101906101f3919061215b565b611b0a565b005b610202611ec3565b60028167ffffffffffffffff16815481106102205761021f61269a565b5b905f5260205f20906004020160020160019054906101000a900460ff161561027d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027490612721565b60405180910390fd5b5f60028267ffffffffffffffff168154811061029c5761029b61269a565b5b905f5260205f20906004020160010160189054906101000a900467ffffffffffffffff1690508067ffffffffffffffff165f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff16116103c6573373ffffffffffffffffffffffffffffffffffffffff1660028367ffffffffffffffff16815481106103275761032661269a565b5b905f5260205f20906004020160020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a490612789565b60405180910390fd5b5f60149054906101000a900467ffffffffffffffff1690505b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161041e906127f1565b6020604051808303815f875af115801561043a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061045e9190612823565b90505f60028467ffffffffffffffff168154811061047f5761047e61269a565b5b905f5260205f2090600402016002015f9054906101000a900460ff161561087a575f8273ffffffffffffffffffffffffffffffffffffffff1663d370a37d60028767ffffffffffffffff16815481106104db576104da61269a565b5b905f5260205f20906004020160010160089054906101000a900467ffffffffffffffff166040518263ffffffff1660e01b815260040161051b919061263a565b6020604051808303815f875af1158015610537573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055b9190612823565b905060028567ffffffffffffffff168154811061057b5761057a61269a565b5b905f5260205f2090600402016003015460028667ffffffffffffffff16815481106105a9576105a861269a565b5b905f5260205f20906004020160010160109054906101000a900467ffffffffffffffff16856105d8919061287b565b67ffffffffffffffff166105ec91906128b6565b9150600a60098360026105ff91906128b6565b61060991906128b6565b6106139190612924565b60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461065e9190612954565b92505081905550600160055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610710906129d1565b6020604051808303815f875af115801561072c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107509190612823565b73ffffffffffffffffffffffffffffffffffffffff1663d370a37d60028767ffffffffffffffff16815481106107895761078861269a565b5b905f5260205f2090600402016001015f9054906101000a900467ffffffffffffffff166040518263ffffffff1660e01b81526004016107c8919061263a565b602060405180830381865afa1580156107e3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108079190612823565b9050600a82600261081891906128b6565b6108229190612924565b60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461086d9190612954565b92505081905550506108bf565b60028467ffffffffffffffff16815481106108985761089761269a565b5b905f5260205f20906004020160010160109054906101000a900467ffffffffffffffff1692505b8173ffffffffffffffffffffffffffffffffffffffff166341fbbc3160028667ffffffffffffffff16815481106108f9576108f861269a565b5b905f5260205f20906004020160010160089054906101000a900467ffffffffffffffff166040518263ffffffff1660e01b8152600401610939919061263a565b5f604051808303815f87803b158015610950575f80fd5b505af1158015610962573d5f803e3d5ffd5b505050505f60028567ffffffffffffffff16815481106109855761098461269a565b5b905f5260205f209060040201600301548460028767ffffffffffffffff16815481106109b4576109b361269a565b5b905f5260205f20906004020160010160189054906101000a900467ffffffffffffffff166109e2919061287b565b67ffffffffffffffff166109f691906128b6565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c5860028767ffffffffffffffff1681548110610a5157610a5061269a565b5b905f5260205f20906004020160020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610a9f9291906129ef565b5f604051808303815f87803b158015610ab6575f80fd5b505af1158015610ac8573d5f803e3d5ffd5b505050508060045f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610b2790612a60565b6020604051808303815f875af1158015610b43573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b679190612823565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610bae9190612954565b92505081905550600160028667ffffffffffffffff1681548110610bd557610bd461269a565b5b905f5260205f20906004020160020160016101000a81548160ff0219169083151502179055505050505050565b60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610d63575f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401610ca790612ac8565b6020604051808303815f875af1158015610cc3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ce79190612823565b73ffffffffffffffffffffffffffffffffffffffff166361e728b13360036040518363ffffffff1660e01b8152600401610d22929190612b34565b602060405180830381865afa158015610d3d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d619190612b6f565b505b8060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610daf9190612b9a565b925050819055505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166376890c5833836040518363ffffffff1660e01b8152600401610e109291906129ef565b5f604051808303815f87803b158015610e27575f80fd5b505af1158015610e39573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436482604051610e839190612278565b60405180910390a250565b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610eff612087565b60405180610120016040528060028467ffffffffffffffff1681548110610f2957610f2861269a565b5b905f5260205f2090600402015f018054610f4290612bfa565b80601f0160208091040260200160405190810160405280929190818152602001828054610f6e90612bfa565b8015610fb95780601f10610f9057610100808354040283529160200191610fb9565b820191905f5260205f20905b815481529060010190602001808311610f9c57829003601f168201915b5050505050815260200160028467ffffffffffffffff1681548110610fe157610fe061269a565b5b905f5260205f2090600402016001015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160028467ffffffffffffffff16815481106110315761103061269a565b5b905f5260205f20906004020160010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160028467ffffffffffffffff16815481106110825761108161269a565b5b905f5260205f20906004020160010160109054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160028467ffffffffffffffff16815481106110d3576110d261269a565b5b905f5260205f20906004020160010160189054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200160028467ffffffffffffffff16815481106111245761112361269a565b5b905f5260205f2090600402016002015f9054906101000a900460ff161515815260200160028467ffffffffffffffff16815481106111655761116461269a565b5b905f5260205f20906004020160020160019054906101000a900460ff161515815260200160028467ffffffffffffffff16815481106111a7576111a661269a565b5b905f5260205f20906004020160020160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028467ffffffffffffffff16815481106112105761120f61269a565b5b905f5260205f209060040201600301548152509050919050565b611232611ec3565b5f6003866040516112439190612c64565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146112af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a690612cc4565b60405180910390fd5b60015f9054906101000a900467ffffffffffffffff165f60149054906101000a900467ffffffffffffffff166112e59190612ce2565b67ffffffffffffffff168167ffffffffffffffff161161133a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133190612d67565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611390906129d1565b6020604051808303815f875af11580156113ac573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113d09190612823565b73ffffffffffffffffffffffffffffffffffffffff1663f804843d856040518263ffffffff1660e01b8152600401611408919061263a565b5f604051808303815f87803b15801561141f575f80fd5b505af1158015611431573d5f803e3d5ffd5b505050505f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161148d906127f1565b6020604051808303815f875af11580156114a9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114cd9190612823565b90508073ffffffffffffffffffffffffffffffffffffffff166371d49a63856040518263ffffffff1660e01b8152600401611508919061263a565b5f604051808303815f87803b15801561151f575f80fd5b505af1158015611531573d5f803e3d5ffd5b505050505f8173ffffffffffffffffffffffffffffffffffffffff1663d370a37d866040518263ffffffff1660e01b815260040161156f919061263a565b6020604051808303815f875af115801561158b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115af9190612823565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b815260040161160790612ac8565b6020604051808303815f875af1158015611623573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116479190612823565b73ffffffffffffffffffffffffffffffffffffffff166361e728b18260036040518363ffffffff1660e01b8152600401611682929190612b34565b602060405180830381865afa15801561169d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116c19190612b6f565b505f8060149054906101000a900467ffffffffffffffff16846116e4919061287b565b67ffffffffffffffff16856116f991906128b6565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e888891533836040518363ffffffff1660e01b81526004016117559291906129ef565b5f604051808303815f87803b15801561176c575f80fd5b505af115801561177e573d5f803e3d5ffd5b505050505f6002805490509050611793612087565b89815f0181905250338160e0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505f60149054906101000a900467ffffffffffffffff16816060019067ffffffffffffffff16908167ffffffffffffffff168152505085816080019067ffffffffffffffff16908167ffffffffffffffff1681525050868161010001818152505088816020019067ffffffffffffffff16908167ffffffffffffffff168152505087816040019067ffffffffffffffff16908167ffffffffffffffff1681525050600281908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f0190816118af9190612f19565b506020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816002015f6101000a81548160ff02191690831515021790555060c08201518160020160016101000a81548160ff02191690831515021790555060e08201518160020160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610100820151816003015550508160038b604051611a0f9190612c64565b90815260200160405180910390205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167ffea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70838c604051611a8a929190613020565b60405180910390a250505050505050505050565b5f60149054906101000a900467ffffffffffffffff1681565b60015f9054906101000a900467ffffffffffffffff1681565b5f600382604051611ae19190612c64565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff169050919050565b611b12611ec3565b60028167ffffffffffffffff1681548110611b3057611b2f61269a565b5b905f5260205f20906004020160020160019054906101000a900460ff1615611b8d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b8490612721565b60405180910390fd5b60028167ffffffffffffffff1681548110611bab57611baa61269a565b5b905f5260205f2090600402016002015f9054906101000a900460ff1615611c07576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bfe90613098565b60405180910390fd5b5f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff16600360028367ffffffffffffffff1681548110611c4757611c4661269a565b5b905f5260205f20906004020160010160109054906101000a900467ffffffffffffffff16611c759190612ce2565b67ffffffffffffffff1611611cbf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cb690613100565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611d2c906127f1565b6020604051808303815f875af1158015611d48573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d6c9190612823565b73ffffffffffffffffffffffffffffffffffffffff1663d370a37d60028467ffffffffffffffff1681548110611da557611da461269a565b5b905f5260205f20906004020160010160089054906101000a900467ffffffffffffffff166040518263ffffffff1660e01b8152600401611de5919061263a565b6020604051808303815f875af1158015611e01573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e259190612823565b73ffffffffffffffffffffffffffffffffffffffff1614611e7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e7290612789565b60405180910390fd5b600160028267ffffffffffffffff1681548110611e9b57611e9a61269a565b5b905f5260205f2090600402016002015f6101000a81548160ff02191690831515021790555050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff1660e01b8152600401611f1b90613168565b6020604051808303815f875af1158015611f37573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f5b9190612823565b90508073ffffffffffffffffffffffffffffffffffffffff1663919840ad6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015611fa2575f80fd5b505af1158015611fb4573d5f803e3d5ffd5b505050505f8173ffffffffffffffffffffffffffffffffffffffff16639fa6a6e36040518163ffffffff1660e01b81526004016020604051808303815f875af1158015612003573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612027919061319a565b90505f60149054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16111561208357805f60146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b5050565b604051806101200160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b61213a8161211e565b8114612144575f80fd5b50565b5f8135905061215581612131565b92915050565b5f602082840312156121705761216f612116565b5b5f61217d84828501612147565b91505092915050565b5f819050919050565b61219881612186565b81146121a2575f80fd5b50565b5f813590506121b38161218f565b92915050565b5f602082840312156121ce576121cd612116565b5b5f6121db848285016121a5565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61220d826121e4565b9050919050565b61221d81612203565b8114612227575f80fd5b50565b5f8135905061223881612214565b92915050565b5f6020828403121561225357612252612116565b5b5f6122608482850161222a565b91505092915050565b61227281612186565b82525050565b5f60208201905061228b5f830184612269565b92915050565b61229a81612203565b82525050565b5f6020820190506122b35f830184612291565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156122f05780820151818401526020810190506122d5565b5f8484015250505050565b5f601f19601f8301169050919050565b5f612315826122b9565b61231f81856122c3565b935061232f8185602086016122d3565b612338816122fb565b840191505092915050565b61234c8161211e565b82525050565b5f8115159050919050565b61236681612352565b82525050565b61237581612203565b82525050565b61238481612186565b82525050565b5f61012083015f8301518482035f8601526123a5828261230b565b91505060208301516123ba6020860182612343565b5060408301516123cd6040860182612343565b5060608301516123e06060860182612343565b5060808301516123f36080860182612343565b5060a083015161240660a086018261235d565b5060c083015161241960c086018261235d565b5060e083015161242c60e086018261236c565b5061010083015161244161010086018261237b565b508091505092915050565b5f6020820190508181035f830152612464818461238a565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6124aa826122fb565b810181811067ffffffffffffffff821117156124c9576124c8612474565b5b80604052505050565b5f6124db61210d565b90506124e782826124a1565b919050565b5f67ffffffffffffffff82111561250657612505612474565b5b61250f826122fb565b9050602081019050919050565b828183375f83830152505050565b5f61253c612537846124ec565b6124d2565b90508281526020810184848401111561255857612557612470565b5b61256384828561251c565b509392505050565b5f82601f83011261257f5761257e61246c565b5b813561258f84826020860161252a565b91505092915050565b5f805f805f60a086880312156125b1576125b0612116565b5b5f86013567ffffffffffffffff8111156125ce576125cd61211a565b5b6125da8882890161256b565b95505060206125eb88828901612147565b94505060406125fc88828901612147565b935050606061260d888289016121a5565b925050608061261e88828901612147565b9150509295509295909350565b6126348161211e565b82525050565b5f60208201905061264d5f83018461262b565b92915050565b5f6020828403121561266857612667612116565b5b5f82013567ffffffffffffffff8111156126855761268461211a565b5b6126918482850161256b565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b7f66696e69736800000000000000000000000000000000000000000000000000005f82015250565b5f61270b6006836126c7565b9150612716826126d7565b602082019050919050565b5f6020820190508181035f830152612738816126ff565b9050919050565b7f696e7600000000000000000000000000000000000000000000000000000000005f82015250565b5f6127736003836126c7565b915061277e8261273f565b602082019050919050565b5f6020820190508181035f8301526127a081612767565b9050919050565b7f67707500000000000000000000000000000000000000000000000000000000005f82015250565b5f6127db6003836126c7565b91506127e6826127a7565b602082019050919050565b5f6020820190508181035f830152612808816127cf565b9050919050565b5f8151905061281d81612214565b92915050565b5f6020828403121561283857612837612116565b5b5f6128458482850161280f565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6128858261211e565b91506128908361211e565b9250828203905067ffffffffffffffff8111156128b0576128af61284e565b5b92915050565b5f6128c082612186565b91506128cb83612186565b92508282026128d981612186565b915082820484148315176128f0576128ef61284e565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61292e82612186565b915061293983612186565b925082612949576129486128f7565b5b828204905092915050565b5f61295e82612186565b915061296983612186565b92508282019050808211156129815761298061284e565b5b92915050565b7f6d6f64656c0000000000000000000000000000000000000000000000000000005f82015250565b5f6129bb6005836126c7565b91506129c682612987565b602082019050919050565b5f6020820190508181035f8301526129e8816129af565b9050919050565b5f604082019050612a025f830185612291565b612a0f6020830184612269565b9392505050565b7f62617365000000000000000000000000000000000000000000000000000000005f82015250565b5f612a4a6004836126c7565b9150612a5582612a16565b602082019050919050565b5f6020820190508181035f830152612a7781612a3e565b9050919050565b7f6e6f6465000000000000000000000000000000000000000000000000000000005f82015250565b5f612ab26004836126c7565b9150612abd82612a7e565b602082019050919050565b5f6020820190508181035f830152612adf81612aa6565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f612b1e612b19612b1484612ae6565b612afb565b612aef565b9050919050565b612b2e81612b04565b82525050565b5f604082019050612b475f830185612291565b612b546020830184612b25565b9392505050565b5f81519050612b698161218f565b92915050565b5f60208284031215612b8457612b83612116565b5b5f612b9184828501612b5b565b91505092915050565b5f612ba482612186565b9150612baf83612186565b9250828203905081811115612bc757612bc661284e565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680612c1157607f821691505b602082108103612c2457612c23612bcd565b5b50919050565b5f81905092915050565b5f612c3e826122b9565b612c488185612c2a565b9350612c588185602086016122d3565b80840191505092915050565b5f612c6f8284612c34565b915081905092915050565b7f65786973740000000000000000000000000000000000000000000000000000005f82015250565b5f612cae6005836126c7565b9150612cb982612c7a565b602082019050919050565b5f6020820190508181035f830152612cdb81612ca2565b9050919050565b5f612cec8261211e565b9150612cf78361211e565b9250828201905067ffffffffffffffff811115612d1757612d1661284e565b5b92915050565b7f73686f72742072656e74696f6e000000000000000000000000000000000000005f82015250565b5f612d51600d836126c7565b9150612d5c82612d1d565b602082019050919050565b5f6020820190508181035f830152612d7e81612d45565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302612de17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612da6565b612deb8683612da6565b95508019841693508086168417925050509392505050565b5f612e1d612e18612e1384612186565b612afb565b612186565b9050919050565b5f819050919050565b612e3683612e03565b612e4a612e4282612e24565b848454612db2565b825550505050565b5f90565b612e5e612e52565b612e69818484612e2d565b505050565b5b81811015612e8c57612e815f82612e56565b600181019050612e6f565b5050565b601f821115612ed157612ea281612d85565b612eab84612d97565b81016020851015612eba578190505b612ece612ec685612d97565b830182612e6e565b50505b505050565b5f82821c905092915050565b5f612ef15f1984600802612ed6565b1980831691505092915050565b5f612f098383612ee2565b9150826002028217905092915050565b612f22826122b9565b67ffffffffffffffff811115612f3b57612f3a612474565b5b612f458254612bfa565b612f50828285612e90565b5f60209050601f831160018114612f81575f8415612f6f578287015190505b612f798582612efe565b865550612fe0565b601f198416612f8f86612d85565b5f5b82811015612fb657848901518255600182019150602085019450602081019050612f91565b86831015612fd35784890151612fcf601f891682612ee2565b8355505b6001600288020188555050505b505050505050565b5f612ff2826122b9565b612ffc81856126c7565b935061300c8185602086016122d3565b613015816122fb565b840191505092915050565b5f6040820190506130335f83018561262b565b81810360208301526130458184612fe8565b90509392505050565b7f61637469766500000000000000000000000000000000000000000000000000005f82015250565b5f6130826006836126c7565b915061308d8261304e565b602082019050919050565b5f6020820190508181035f8301526130af81613076565b9050919050565b7f746f6f206c6174650000000000000000000000000000000000000000000000005f82015250565b5f6130ea6008836126c7565b91506130f5826130b6565b602082019050919050565b5f6020820190508181035f830152613117816130de565b9050919050565b7f65706f63680000000000000000000000000000000000000000000000000000005f82015250565b5f6131526005836126c7565b915061315d8261311e565b602082019050919050565b5f6020820190508181035f83015261317f81613146565b9050919050565b5f8151905061319481612131565b92915050565b5f602082840312156131af576131ae612116565b5b5f6131bc84828501613186565b9150509291505056fea2646970667358221220c49a2cf6f1d7a746ffb897cd956340d38bcae3a097b2ce05527f11171db10ae764736f6c634300081a0033",
}

// SpaceABI is the input ABI used to generate the binding from.
// Deprecated: Use SpaceMetaData.ABI instead.
var SpaceABI = SpaceMetaData.ABI

// SpaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SpaceMetaData.Bin instead.
var SpaceBin = SpaceMetaData.Bin

// DeploySpace deploys a new Ethereum contract, binding an instance of Space to it.
func DeploySpace(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Space, error) {
	parsed, err := SpaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SpaceBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Space{SpaceCaller: SpaceCaller{contract: contract}, SpaceTransactor: SpaceTransactor{contract: contract}, SpaceFilterer: SpaceFilterer{contract: contract}}, nil
}

// Space is an auto generated Go binding around an Ethereum contract.
type Space struct {
	SpaceCaller     // Read-only binding to the contract
	SpaceTransactor // Write-only binding to the contract
	SpaceFilterer   // Log filterer for contract events
}

// SpaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpaceSession struct {
	Contract     *Space            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpaceCallerSession struct {
	Contract *SpaceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpaceTransactorSession struct {
	Contract     *SpaceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpaceRaw struct {
	Contract *Space // Generic contract binding to access the raw methods on
}

// SpaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpaceCallerRaw struct {
	Contract *SpaceCaller // Generic read-only contract binding to access the raw methods on
}

// SpaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpaceTransactorRaw struct {
	Contract *SpaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpace creates a new instance of Space, bound to a specific deployed contract.
func NewSpace(address common.Address, backend bind.ContractBackend) (*Space, error) {
	contract, err := bindSpace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Space{SpaceCaller: SpaceCaller{contract: contract}, SpaceTransactor: SpaceTransactor{contract: contract}, SpaceFilterer: SpaceFilterer{contract: contract}}, nil
}

// NewSpaceCaller creates a new read-only instance of Space, bound to a specific deployed contract.
func NewSpaceCaller(address common.Address, caller bind.ContractCaller) (*SpaceCaller, error) {
	contract, err := bindSpace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpaceCaller{contract: contract}, nil
}

// NewSpaceTransactor creates a new write-only instance of Space, bound to a specific deployed contract.
func NewSpaceTransactor(address common.Address, transactor bind.ContractTransactor) (*SpaceTransactor, error) {
	contract, err := bindSpace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpaceTransactor{contract: contract}, nil
}

// NewSpaceFilterer creates a new log filterer instance of Space, bound to a specific deployed contract.
func NewSpaceFilterer(address common.Address, filterer bind.ContractFilterer) (*SpaceFilterer, error) {
	contract, err := bindSpace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpaceFilterer{contract: contract}, nil
}

// bindSpace binds a generic wrapper to an already deployed contract.
func bindSpace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SpaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Space *SpaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Space.Contract.SpaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Space *SpaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Space.Contract.SpaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Space *SpaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Space.Contract.SpaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Space *SpaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Space.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Space *SpaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Space.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Space *SpaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Space.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_Space *SpaceCaller) BalanceOf(opts *bind.CallOpts, _a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Space.contract.Call(opts, &out, "balanceOf", _a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_Space *SpaceSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _Space.Contract.BalanceOf(&_Space.CallOpts, _a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_Space *SpaceCallerSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _Space.Contract.BalanceOf(&_Space.CallOpts, _a)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Space *SpaceCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Space.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Space *SpaceSession) Bank() (common.Address, error) {
	return _Space.Contract.Bank(&_Space.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Space *SpaceCallerSession) Bank() (common.Address, error) {
	return _Space.Contract.Bank(&_Space.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Space *SpaceCaller) Current(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Space.contract.Call(opts, &out, "current")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Space *SpaceSession) Current() (uint64, error) {
	return _Space.Contract.Current(&_Space.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Space *SpaceCallerSession) Current() (uint64, error) {
	return _Space.Contract.Current(&_Space.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _sn) view returns(uint64)
func (_Space *SpaceCaller) GetIndex(opts *bind.CallOpts, _sn string) (uint64, error) {
	var out []interface{}
	err := _Space.contract.Call(opts, &out, "getIndex", _sn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _sn) view returns(uint64)
func (_Space *SpaceSession) GetIndex(_sn string) (uint64, error) {
	return _Space.Contract.GetIndex(&_Space.CallOpts, _sn)
}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _sn) view returns(uint64)
func (_Space *SpaceCallerSession) GetIndex(_sn string) (uint64, error) {
	return _Space.Contract.GetIndex(&_Space.CallOpts, _sn)
}

// GetSpace is a free data retrieval call binding the contract method 0x91fb07c4.
//
// Solidity: function getSpace(uint64 _si) view returns((string,uint64,uint64,uint64,uint64,bool,bool,address,uint256) _info)
func (_Space *SpaceCaller) GetSpace(opts *bind.CallOpts, _si uint64) (ISpaceInfo, error) {
	var out []interface{}
	err := _Space.contract.Call(opts, &out, "getSpace", _si)

	if err != nil {
		return *new(ISpaceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISpaceInfo)).(*ISpaceInfo)

	return out0, err

}

// GetSpace is a free data retrieval call binding the contract method 0x91fb07c4.
//
// Solidity: function getSpace(uint64 _si) view returns((string,uint64,uint64,uint64,uint64,bool,bool,address,uint256) _info)
func (_Space *SpaceSession) GetSpace(_si uint64) (ISpaceInfo, error) {
	return _Space.Contract.GetSpace(&_Space.CallOpts, _si)
}

// GetSpace is a free data retrieval call binding the contract method 0x91fb07c4.
//
// Solidity: function getSpace(uint64 _si) view returns((string,uint64,uint64,uint64,uint64,bool,bool,address,uint256) _info)
func (_Space *SpaceCallerSession) GetSpace(_si uint64) (ISpaceInfo, error) {
	return _Space.Contract.GetSpace(&_Space.CallOpts, _si)
}

// MinRent is a free data retrieval call binding the contract method 0xe38c6116.
//
// Solidity: function minRent() view returns(uint64)
func (_Space *SpaceCaller) MinRent(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Space.contract.Call(opts, &out, "minRent")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinRent is a free data retrieval call binding the contract method 0xe38c6116.
//
// Solidity: function minRent() view returns(uint64)
func (_Space *SpaceSession) MinRent() (uint64, error) {
	return _Space.Contract.MinRent(&_Space.CallOpts)
}

// MinRent is a free data retrieval call binding the contract method 0xe38c6116.
//
// Solidity: function minRent() view returns(uint64)
func (_Space *SpaceCallerSession) MinRent() (uint64, error) {
	return _Space.Contract.MinRent(&_Space.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _si) returns()
func (_Space *SpaceTransactor) Activate(opts *bind.TransactOpts, _si uint64) (*types.Transaction, error) {
	return _Space.contract.Transact(opts, "activate", _si)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _si) returns()
func (_Space *SpaceSession) Activate(_si uint64) (*types.Transaction, error) {
	return _Space.Contract.Activate(&_Space.TransactOpts, _si)
}

// Activate is a paid mutator transaction binding the contract method 0xfc7e4746.
//
// Solidity: function activate(uint64 _si) returns()
func (_Space *SpaceTransactorSession) Activate(_si uint64) (*types.Transaction, error) {
	return _Space.Contract.Activate(&_Space.TransactOpts, _si)
}

// Add is a paid mutator transaction binding the contract method 0x99760d12.
//
// Solidity: function add(string _sn, uint64 _mi, uint64 _gi, uint256 _p, uint64 _e) returns()
func (_Space *SpaceTransactor) Add(opts *bind.TransactOpts, _sn string, _mi uint64, _gi uint64, _p *big.Int, _e uint64) (*types.Transaction, error) {
	return _Space.contract.Transact(opts, "add", _sn, _mi, _gi, _p, _e)
}

// Add is a paid mutator transaction binding the contract method 0x99760d12.
//
// Solidity: function add(string _sn, uint64 _mi, uint64 _gi, uint256 _p, uint64 _e) returns()
func (_Space *SpaceSession) Add(_sn string, _mi uint64, _gi uint64, _p *big.Int, _e uint64) (*types.Transaction, error) {
	return _Space.Contract.Add(&_Space.TransactOpts, _sn, _mi, _gi, _p, _e)
}

// Add is a paid mutator transaction binding the contract method 0x99760d12.
//
// Solidity: function add(string _sn, uint64 _mi, uint64 _gi, uint256 _p, uint64 _e) returns()
func (_Space *SpaceTransactorSession) Add(_sn string, _mi uint64, _gi uint64, _p *big.Int, _e uint64) (*types.Transaction, error) {
	return _Space.Contract.Add(&_Space.TransactOpts, _sn, _mi, _gi, _p, _e)
}

// Shutdown is a paid mutator transaction binding the contract method 0x1b1683f7.
//
// Solidity: function shutdown(uint64 _si) returns()
func (_Space *SpaceTransactor) Shutdown(opts *bind.TransactOpts, _si uint64) (*types.Transaction, error) {
	return _Space.contract.Transact(opts, "shutdown", _si)
}

// Shutdown is a paid mutator transaction binding the contract method 0x1b1683f7.
//
// Solidity: function shutdown(uint64 _si) returns()
func (_Space *SpaceSession) Shutdown(_si uint64) (*types.Transaction, error) {
	return _Space.Contract.Shutdown(&_Space.TransactOpts, _si)
}

// Shutdown is a paid mutator transaction binding the contract method 0x1b1683f7.
//
// Solidity: function shutdown(uint64 _si) returns()
func (_Space *SpaceTransactorSession) Shutdown(_si uint64) (*types.Transaction, error) {
	return _Space.Contract.Shutdown(&_Space.TransactOpts, _si)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Space *SpaceTransactor) Withdraw(opts *bind.TransactOpts, _m *big.Int) (*types.Transaction, error) {
	return _Space.contract.Transact(opts, "withdraw", _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Space *SpaceSession) Withdraw(_m *big.Int) (*types.Transaction, error) {
	return _Space.Contract.Withdraw(&_Space.TransactOpts, _m)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _m) returns()
func (_Space *SpaceTransactorSession) Withdraw(_m *big.Int) (*types.Transaction, error) {
	return _Space.Contract.Withdraw(&_Space.TransactOpts, _m)
}

// SpaceAddGPUIterator is returned from FilterAddGPU and is used to iterate over the raw logs and unpacked data for AddGPU events raised by the Space contract.
type SpaceAddGPUIterator struct {
	Event *SpaceAddGPU // Event containing the contract specifics and raw log

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
func (it *SpaceAddGPUIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpaceAddGPU)
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
		it.Event = new(SpaceAddGPU)
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
func (it *SpaceAddGPUIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpaceAddGPUIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpaceAddGPU represents a AddGPU event raised by the Space contract.
type SpaceAddGPU struct {
	A   common.Address
	Ai  uint64
	An  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddGPU is a free log retrieval operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _ai, string _an)
func (_Space *SpaceFilterer) FilterAddGPU(opts *bind.FilterOpts, _a []common.Address) (*SpaceAddGPUIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Space.contract.FilterLogs(opts, "AddGPU", _aRule)
	if err != nil {
		return nil, err
	}
	return &SpaceAddGPUIterator{contract: _Space.contract, event: "AddGPU", logs: logs, sub: sub}, nil
}

// WatchAddGPU is a free log subscription operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _ai, string _an)
func (_Space *SpaceFilterer) WatchAddGPU(opts *bind.WatchOpts, sink chan<- *SpaceAddGPU, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Space.contract.WatchLogs(opts, "AddGPU", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpaceAddGPU)
				if err := _Space.contract.UnpackLog(event, "AddGPU", log); err != nil {
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

// ParseAddGPU is a log parse operation binding the contract event 0xfea931920f2df20d0e447fe0ee025f54206ac7d1266b239bb4a96eb075151f70.
//
// Solidity: event AddGPU(address indexed _a, uint64 _ai, string _an)
func (_Space *SpaceFilterer) ParseAddGPU(log types.Log) (*SpaceAddGPU, error) {
	event := new(SpaceAddGPU)
	if err := _Space.contract.UnpackLog(event, "AddGPU", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpaceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Space contract.
type SpaceWithdrawIterator struct {
	Event *SpaceWithdraw // Event containing the contract specifics and raw log

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
func (it *SpaceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpaceWithdraw)
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
		it.Event = new(SpaceWithdraw)
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
func (it *SpaceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpaceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpaceWithdraw represents a Withdraw event raised by the Space contract.
type SpaceWithdraw struct {
	A   common.Address
	M   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Space *SpaceFilterer) FilterWithdraw(opts *bind.FilterOpts, _a []common.Address) (*SpaceWithdrawIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Space.contract.FilterLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return &SpaceWithdrawIterator{contract: _Space.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Space *SpaceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SpaceWithdraw, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Space.contract.WatchLogs(opts, "Withdraw", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpaceWithdraw)
				if err := _Space.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed _a, uint256 _m)
func (_Space *SpaceFilterer) ParseWithdraw(log types.Log) (*SpaceWithdraw, error) {
	event := new(SpaceWithdraw)
	if err := _Space.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
