// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package model

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

// IModelInfo is an auto generated low-level Go binding around an user-defined struct.
type IModelInfo struct {
	Name   string
	Owner  common.Address
	Active bool
	Root   [32]byte
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"AddModel\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32)
func (_IModel *IModelCaller) GetRoot(opts *bind.CallOpts, _mi uint64) ([32]byte, error) {
	var out []interface{}
	err := _IModel.contract.Call(opts, &out, "getRoot", _mi)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32)
func (_IModel *IModelSession) GetRoot(_mi uint64) ([32]byte, error) {
	return _IModel.Contract.GetRoot(&_IModel.CallOpts, _mi)
}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32)
func (_IModel *IModelCallerSession) GetRoot(_mi uint64) ([32]byte, error) {
	return _IModel.Contract.GetRoot(&_IModel.CallOpts, _mi)
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

// ModelMetaData contains all meta data concerning the Model contract.
var ModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"AddModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_s\",\"type\":\"bool\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_rt\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getModel\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structIModel.Info\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051611be7380380611be7833981810160405281019061003191906102b1565b61004d61004261014b60201b60201c565b61015260201b60201c565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610095610213565b600281908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816100d19190610516565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160020155505050506105e5565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6040518060800160405280606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f151581526020015f80191681525090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61028082610257565b9050919050565b61029081610276565b811461029a575f80fd5b50565b5f815190506102ab81610287565b92915050565b5f602082840312156102c6576102c5610253565b5b5f6102d38482850161029d565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061035757607f821691505b60208210810361036a57610369610313565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103cc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610391565b6103d68683610391565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61041a610415610410846103ee565b6103f7565b6103ee565b9050919050565b5f819050919050565b61043383610400565b61044761043f82610421565b84845461039d565b825550505050565b5f90565b61045b61044f565b61046681848461042a565b505050565b5b818110156104895761047e5f82610453565b60018101905061046c565b5050565b601f8211156104ce5761049f81610370565b6104a884610382565b810160208510156104b7578190505b6104cb6104c385610382565b83018261046b565b50505b505050565b5f82821c905092915050565b5f6104ee5f19846008026104d3565b1980831691505092915050565b5f61050683836104df565b9150826002028217905092915050565b61051f826102dc565b67ffffffffffffffff811115610538576105376102e6565b5b6105428254610340565b61054d82828561048d565b5f60209050601f83116001811461057e575f841561056c578287015190505b61057685826104fb565b8655506105dd565b601f19841661058c86610370565b5f5b828110156105b35784890151825560018201915060208501945060208101905061058e565b868310156105d057848901516105cc601f8916826104df565b8355505b6001600288020188555050505b505050505050565b6115f5806105f25f395ff3fe608060405234801561000f575f80fd5b50600436106100a7575f3560e01c80639f5591ab1161006f5780639f5591ab1461013d578063d370a37d1461016d578063d48dca2d1461019d578063ea1bbe35146101b9578063f2fde38b146101e9578063f804843d14610205576100a7565b8063691de4a4146100ab5780636c37609a146100c7578063715018a6146100f757806376cdb03b146101015780638da5cb5b1461011f575b5f80fd5b6100c560048036038101906100c09190610c2c565b610221565b005b6100e160048036038101906100dc9190610cc3565b610453565b6040516100ee9190610cfd565b60405180910390f35b6100ff610489565b005b61010961049c565b6040516101169190610d55565b60405180910390f35b6101276104c1565b6040516101349190610d55565b60405180910390f35b61015760048036038101906101529190610cc3565b6104e8565b6040516101649190610e80565b60405180910390f35b61018760048036038101906101829190610cc3565b610695565b6040516101949190610d55565b60405180910390f35b6101b760048036038101906101b29190610eca565b6106ea565b005b6101d360048036038101906101ce9190610f08565b6107ed565b6040516101e09190610f5e565b60405180910390f35b61020360048036038101906101fe9190610fa1565b610827565b005b61021f600480360381019061021a9190610cc3565b6108a9565b005b5f6003836040516102329190611006565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461029e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029590611076565b60405180910390fd5b5f60028054905090506102af610a6c565b83815f018190525033816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160600181815250506001816040019015159081151581525050600281908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816103469190611297565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff021916908315150217905550606082015181600201555050816003856040516103ca9190611006565b90815260200160405180910390205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47838660405161044592919061139e565b60405180910390a250505050565b5f60028267ffffffffffffffff1681548110610472576104716113cc565b5b905f5260205f209060030201600201549050919050565b610491610926565b61049a5f6109a4565b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6104f0610a6c565b604051806080016040528060028467ffffffffffffffff1681548110610519576105186113cc565b5b905f5260205f2090600302015f018054610532906110c1565b80601f016020809104026020016040519081016040528092919081815260200182805461055e906110c1565b80156105a95780601f10610580576101008083540402835291602001916105a9565b820191905f5260205f20905b81548152906001019060200180831161058c57829003601f168201915b5050505050815260200160028467ffffffffffffffff16815481106105d1576105d06113cc565b5b905f5260205f2090600302016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028467ffffffffffffffff1681548110610639576106386113cc565b5b905f5260205f20906003020160010160149054906101000a900460ff161515815260200160028467ffffffffffffffff168154811061067b5761067a6113cc565b5b905f5260205f209060030201600201548152509050919050565b5f60028267ffffffffffffffff16815481106106b4576106b36113cc565b5b905f5260205f2090600302016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660028367ffffffffffffffff168154811061071f5761071e6113cc565b5b905f5260205f2090600302016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079b90611443565b60405180910390fd5b8060028367ffffffffffffffff16815481106107c3576107c26113cc565b5b905f5260205f20906003020160010160146101000a81548160ff0219169083151502179055505050565b5f6003826040516107fe9190611006565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff169050919050565b61082f610926565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361089d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610894906114d1565b60405180910390fd5b6108a6816109a4565b50565b60028167ffffffffffffffff16815481106108c7576108c66113cc565b5b905f5260205f20906003020160010160149054906101000a900460ff16610923576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091a90611539565b60405180910390fd5b50565b61092e610a65565b73ffffffffffffffffffffffffffffffffffffffff1661094c6104c1565b73ffffffffffffffffffffffffffffffffffffffff16146109a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610999906115a1565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b6040518060800160405280606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f151581526020015f80191681525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b0b82610ac5565b810181811067ffffffffffffffff82111715610b2a57610b29610ad5565b5b80604052505050565b5f610b3c610aac565b9050610b488282610b02565b919050565b5f67ffffffffffffffff821115610b6757610b66610ad5565b5b610b7082610ac5565b9050602081019050919050565b828183375f83830152505050565b5f610b9d610b9884610b4d565b610b33565b905082815260208101848484011115610bb957610bb8610ac1565b5b610bc4848285610b7d565b509392505050565b5f82601f830112610be057610bdf610abd565b5b8135610bf0848260208601610b8b565b91505092915050565b5f819050919050565b610c0b81610bf9565b8114610c15575f80fd5b50565b5f81359050610c2681610c02565b92915050565b5f8060408385031215610c4257610c41610ab5565b5b5f83013567ffffffffffffffff811115610c5f57610c5e610ab9565b5b610c6b85828601610bcc565b9250506020610c7c85828601610c18565b9150509250929050565b5f67ffffffffffffffff82169050919050565b610ca281610c86565b8114610cac575f80fd5b50565b5f81359050610cbd81610c99565b92915050565b5f60208284031215610cd857610cd7610ab5565b5b5f610ce584828501610caf565b91505092915050565b610cf781610bf9565b82525050565b5f602082019050610d105f830184610cee565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d3f82610d16565b9050919050565b610d4f81610d35565b82525050565b5f602082019050610d685f830184610d46565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610da5578082015181840152602081019050610d8a565b5f8484015250505050565b5f610dba82610d6e565b610dc48185610d78565b9350610dd4818560208601610d88565b610ddd81610ac5565b840191505092915050565b610df181610d35565b82525050565b5f8115159050919050565b610e0b81610df7565b82525050565b610e1a81610bf9565b82525050565b5f608083015f8301518482035f860152610e3a8282610db0565b9150506020830151610e4f6020860182610de8565b506040830151610e626040860182610e02565b506060830151610e756060860182610e11565b508091505092915050565b5f6020820190508181035f830152610e988184610e20565b905092915050565b610ea981610df7565b8114610eb3575f80fd5b50565b5f81359050610ec481610ea0565b92915050565b5f8060408385031215610ee057610edf610ab5565b5b5f610eed85828601610caf565b9250506020610efe85828601610eb6565b9150509250929050565b5f60208284031215610f1d57610f1c610ab5565b5b5f82013567ffffffffffffffff811115610f3a57610f39610ab9565b5b610f4684828501610bcc565b91505092915050565b610f5881610c86565b82525050565b5f602082019050610f715f830184610f4f565b92915050565b610f8081610d35565b8114610f8a575f80fd5b50565b5f81359050610f9b81610f77565b92915050565b5f60208284031215610fb657610fb5610ab5565b5b5f610fc384828501610f8d565b91505092915050565b5f81905092915050565b5f610fe082610d6e565b610fea8185610fcc565b9350610ffa818560208601610d88565b80840191505092915050565b5f6110118284610fd6565b915081905092915050565b5f82825260208201905092915050565b7f65786973740000000000000000000000000000000000000000000000000000005f82015250565b5f61106060058361101c565b915061106b8261102c565b602082019050919050565b5f6020820190508181035f83015261108d81611054565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806110d857607f821691505b6020821081036110eb576110ea611094565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261114d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611112565b6111578683611112565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61119b6111966111918461116f565b611178565b61116f565b9050919050565b5f819050919050565b6111b483611181565b6111c86111c0826111a2565b84845461111e565b825550505050565b5f90565b6111dc6111d0565b6111e78184846111ab565b505050565b5b8181101561120a576111ff5f826111d4565b6001810190506111ed565b5050565b601f82111561124f57611220816110f1565b61122984611103565b81016020851015611238578190505b61124c61124485611103565b8301826111ec565b50505b505050565b5f82821c905092915050565b5f61126f5f1984600802611254565b1980831691505092915050565b5f6112878383611260565b9150826002028217905092915050565b6112a082610d6e565b67ffffffffffffffff8111156112b9576112b8610ad5565b5b6112c382546110c1565b6112ce82828561120e565b5f60209050601f8311600181146112ff575f84156112ed578287015190505b6112f7858261127c565b86555061135e565b601f19841661130d866110f1565b5f5b828110156113345784890151825560018201915060208501945060208101905061130f565b86831015611351578489015161134d601f891682611260565b8355505b6001600288020188555050505b505050505050565b5f61137082610d6e565b61137a818561101c565b935061138a818560208601610d88565b61139381610ac5565b840191505092915050565b5f6040820190506113b15f830185610f4f565b81810360208301526113c38184611366565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f696e7600000000000000000000000000000000000000000000000000000000005f82015250565b5f61142d60038361101c565b9150611438826113f9565b602082019050919050565b5f6020820190508181035f83015261145a81611421565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6114bb60268361101c565b91506114c682611461565b604082019050919050565b5f6020820190508181035f8301526114e8816114af565b9050919050565b7f6e6f2061637469766500000000000000000000000000000000000000000000005f82015250565b5f61152360098361101c565b915061152e826114ef565b602082019050919050565b5f6020820190508181035f83015261155081611517565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61158b60208361101c565b915061159682611557565b602082019050919050565b5f6020820190508181035f8301526115b88161157f565b905091905056fea264697066735822122048bc070d8d42cb54f377a676d1a4e5ab6d42fa928261f63f5e8bac88e215171e64736f6c634300081a0033",
}

// ModelABI is the input ABI used to generate the binding from.
// Deprecated: Use ModelMetaData.ABI instead.
var ModelABI = ModelMetaData.ABI

// ModelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModelMetaData.Bin instead.
var ModelBin = ModelMetaData.Bin

// DeployModel deploys a new Ethereum contract, binding an instance of Model to it.
func DeployModel(auth *bind.TransactOpts, backend bind.ContractBackend, _b common.Address) (common.Address, *types.Transaction, *Model, error) {
	parsed, err := ModelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModelBin), backend, _b)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Model{ModelCaller: ModelCaller{contract: contract}, ModelTransactor: ModelTransactor{contract: contract}, ModelFilterer: ModelFilterer{contract: contract}}, nil
}

// Model is an auto generated Go binding around an Ethereum contract.
type Model struct {
	ModelCaller     // Read-only binding to the contract
	ModelTransactor // Write-only binding to the contract
	ModelFilterer   // Log filterer for contract events
}

// ModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelSession struct {
	Contract     *Model            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelCallerSession struct {
	Contract *ModelCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelTransactorSession struct {
	Contract     *ModelTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelRaw struct {
	Contract *Model // Generic contract binding to access the raw methods on
}

// ModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelCallerRaw struct {
	Contract *ModelCaller // Generic read-only contract binding to access the raw methods on
}

// ModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelTransactorRaw struct {
	Contract *ModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModel creates a new instance of Model, bound to a specific deployed contract.
func NewModel(address common.Address, backend bind.ContractBackend) (*Model, error) {
	contract, err := bindModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Model{ModelCaller: ModelCaller{contract: contract}, ModelTransactor: ModelTransactor{contract: contract}, ModelFilterer: ModelFilterer{contract: contract}}, nil
}

// NewModelCaller creates a new read-only instance of Model, bound to a specific deployed contract.
func NewModelCaller(address common.Address, caller bind.ContractCaller) (*ModelCaller, error) {
	contract, err := bindModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCaller{contract: contract}, nil
}

// NewModelTransactor creates a new write-only instance of Model, bound to a specific deployed contract.
func NewModelTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelTransactor, error) {
	contract, err := bindModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelTransactor{contract: contract}, nil
}

// NewModelFilterer creates a new log filterer instance of Model, bound to a specific deployed contract.
func NewModelFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelFilterer, error) {
	contract, err := bindModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelFilterer{contract: contract}, nil
}

// bindModel binds a generic wrapper to an already deployed contract.
func bindModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Model *ModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Model.Contract.ModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Model *ModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.Contract.ModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Model *ModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Model.Contract.ModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Model *ModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Model.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Model *ModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Model *ModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Model.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Model *ModelCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Model *ModelSession) Bank() (common.Address, error) {
	return _Model.Contract.Bank(&_Model.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Model *ModelCallerSession) Bank() (common.Address, error) {
	return _Model.Contract.Bank(&_Model.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) view returns()
func (_Model *ModelCaller) Check(opts *bind.CallOpts, _mi uint64) error {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "check", _mi)

	if err != nil {
		return err
	}

	return err

}

// Check is a free data retrieval call binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) view returns()
func (_Model *ModelSession) Check(_mi uint64) error {
	return _Model.Contract.Check(&_Model.CallOpts, _mi)
}

// Check is a free data retrieval call binding the contract method 0xf804843d.
//
// Solidity: function check(uint64 _mi) view returns()
func (_Model *ModelCallerSession) Check(_mi uint64) error {
	return _Model.Contract.Check(&_Model.CallOpts, _mi)
}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _mn) view returns(uint64)
func (_Model *ModelCaller) GetIndex(opts *bind.CallOpts, _mn string) (uint64, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getIndex", _mn)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _mn) view returns(uint64)
func (_Model *ModelSession) GetIndex(_mn string) (uint64, error) {
	return _Model.Contract.GetIndex(&_Model.CallOpts, _mn)
}

// GetIndex is a free data retrieval call binding the contract method 0xea1bbe35.
//
// Solidity: function getIndex(string _mn) view returns(uint64)
func (_Model *ModelCallerSession) GetIndex(_mn string) (uint64, error) {
	return _Model.Contract.GetIndex(&_Model.CallOpts, _mn)
}

// GetModel is a free data retrieval call binding the contract method 0x9f5591ab.
//
// Solidity: function getModel(uint64 _mi) view returns((string,address,bool,bytes32) _info)
func (_Model *ModelCaller) GetModel(opts *bind.CallOpts, _mi uint64) (IModelInfo, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getModel", _mi)

	if err != nil {
		return *new(IModelInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IModelInfo)).(*IModelInfo)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x9f5591ab.
//
// Solidity: function getModel(uint64 _mi) view returns((string,address,bool,bytes32) _info)
func (_Model *ModelSession) GetModel(_mi uint64) (IModelInfo, error) {
	return _Model.Contract.GetModel(&_Model.CallOpts, _mi)
}

// GetModel is a free data retrieval call binding the contract method 0x9f5591ab.
//
// Solidity: function getModel(uint64 _mi) view returns((string,address,bool,bytes32) _info)
func (_Model *ModelCallerSession) GetModel(_mi uint64) (IModelInfo, error) {
	return _Model.Contract.GetModel(&_Model.CallOpts, _mi)
}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_Model *ModelCaller) GetOwner(opts *bind.CallOpts, _mi uint64) (common.Address, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getOwner", _mi)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_Model *ModelSession) GetOwner(_mi uint64) (common.Address, error) {
	return _Model.Contract.GetOwner(&_Model.CallOpts, _mi)
}

// GetOwner is a free data retrieval call binding the contract method 0xd370a37d.
//
// Solidity: function getOwner(uint64 _mi) view returns(address)
func (_Model *ModelCallerSession) GetOwner(_mi uint64) (common.Address, error) {
	return _Model.Contract.GetOwner(&_Model.CallOpts, _mi)
}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32)
func (_Model *ModelCaller) GetRoot(opts *bind.CallOpts, _mi uint64) ([32]byte, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getRoot", _mi)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32)
func (_Model *ModelSession) GetRoot(_mi uint64) ([32]byte, error) {
	return _Model.Contract.GetRoot(&_Model.CallOpts, _mi)
}

// GetRoot is a free data retrieval call binding the contract method 0x6c37609a.
//
// Solidity: function getRoot(uint64 _mi) view returns(bytes32)
func (_Model *ModelCallerSession) GetRoot(_mi uint64) ([32]byte, error) {
	return _Model.Contract.GetRoot(&_Model.CallOpts, _mi)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Model *ModelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Model *ModelSession) Owner() (common.Address, error) {
	return _Model.Contract.Owner(&_Model.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Model *ModelCallerSession) Owner() (common.Address, error) {
	return _Model.Contract.Owner(&_Model.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xd48dca2d.
//
// Solidity: function activate(uint64 _mi, bool _s) returns()
func (_Model *ModelTransactor) Activate(opts *bind.TransactOpts, _mi uint64, _s bool) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "activate", _mi, _s)
}

// Activate is a paid mutator transaction binding the contract method 0xd48dca2d.
//
// Solidity: function activate(uint64 _mi, bool _s) returns()
func (_Model *ModelSession) Activate(_mi uint64, _s bool) (*types.Transaction, error) {
	return _Model.Contract.Activate(&_Model.TransactOpts, _mi, _s)
}

// Activate is a paid mutator transaction binding the contract method 0xd48dca2d.
//
// Solidity: function activate(uint64 _mi, bool _s) returns()
func (_Model *ModelTransactorSession) Activate(_mi uint64, _s bool) (*types.Transaction, error) {
	return _Model.Contract.Activate(&_Model.TransactOpts, _mi, _s)
}

// Add is a paid mutator transaction binding the contract method 0x691de4a4.
//
// Solidity: function add(string _mn, bytes32 _rt) returns()
func (_Model *ModelTransactor) Add(opts *bind.TransactOpts, _mn string, _rt [32]byte) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "add", _mn, _rt)
}

// Add is a paid mutator transaction binding the contract method 0x691de4a4.
//
// Solidity: function add(string _mn, bytes32 _rt) returns()
func (_Model *ModelSession) Add(_mn string, _rt [32]byte) (*types.Transaction, error) {
	return _Model.Contract.Add(&_Model.TransactOpts, _mn, _rt)
}

// Add is a paid mutator transaction binding the contract method 0x691de4a4.
//
// Solidity: function add(string _mn, bytes32 _rt) returns()
func (_Model *ModelTransactorSession) Add(_mn string, _rt [32]byte) (*types.Transaction, error) {
	return _Model.Contract.Add(&_Model.TransactOpts, _mn, _rt)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Model *ModelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Model *ModelSession) RenounceOwnership() (*types.Transaction, error) {
	return _Model.Contract.RenounceOwnership(&_Model.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Model *ModelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Model.Contract.RenounceOwnership(&_Model.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Model *ModelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Model *ModelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Model.Contract.TransferOwnership(&_Model.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Model *ModelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Model.Contract.TransferOwnership(&_Model.TransactOpts, newOwner)
}

// ModelAddModelIterator is returned from FilterAddModel and is used to iterate over the raw logs and unpacked data for AddModel events raised by the Model contract.
type ModelAddModelIterator struct {
	Event *ModelAddModel // Event containing the contract specifics and raw log

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
func (it *ModelAddModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelAddModel)
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
		it.Event = new(ModelAddModel)
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
func (it *ModelAddModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelAddModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelAddModel represents a AddModel event raised by the Model contract.
type ModelAddModel struct {
	A   common.Address
	Mi  uint64
	Mn  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddModel is a free log retrieval operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_Model *ModelFilterer) FilterAddModel(opts *bind.FilterOpts, _a []common.Address) (*ModelAddModelIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Model.contract.FilterLogs(opts, "AddModel", _aRule)
	if err != nil {
		return nil, err
	}
	return &ModelAddModelIterator{contract: _Model.contract, event: "AddModel", logs: logs, sub: sub}, nil
}

// WatchAddModel is a free log subscription operation binding the contract event 0x1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47.
//
// Solidity: event AddModel(address indexed _a, uint64 _mi, string _mn)
func (_Model *ModelFilterer) WatchAddModel(opts *bind.WatchOpts, sink chan<- *ModelAddModel, _a []common.Address) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}

	logs, sub, err := _Model.contract.WatchLogs(opts, "AddModel", _aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelAddModel)
				if err := _Model.contract.UnpackLog(event, "AddModel", log); err != nil {
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
func (_Model *ModelFilterer) ParseAddModel(log types.Log) (*ModelAddModel, error) {
	event := new(ModelAddModel)
	if err := _Model.contract.UnpackLog(event, "AddModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Model contract.
type ModelOwnershipTransferredIterator struct {
	Event *ModelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ModelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelOwnershipTransferred)
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
		it.Event = new(ModelOwnershipTransferred)
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
func (it *ModelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelOwnershipTransferred represents a OwnershipTransferred event raised by the Model contract.
type ModelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Model *ModelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ModelOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Model.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ModelOwnershipTransferredIterator{contract: _Model.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Model *ModelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ModelOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Model.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelOwnershipTransferred)
				if err := _Model.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Model *ModelFilterer) ParseOwnershipTransferred(log types.Log) (*ModelOwnershipTransferred, error) {
	event := new(ModelOwnershipTransferred)
	if err := _Model.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
