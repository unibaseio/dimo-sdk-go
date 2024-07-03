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
	Root   []byte
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

// ModelMetaData contains all meta data concerning the Model contract.
var ModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_b\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"AddModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_s\",\"type\":\"bool\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_rt\",\"type\":\"bytes\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mn\",\"type\":\"string\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getModel\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"}],\"internalType\":\"structIModel.Info\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_mi\",\"type\":\"uint64\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051611f10380380611f10833981810160405281019061003191906102bb565b61004d61004261015760201b60201c565b61015e60201b60201c565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061009561021f565b600281908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816100d19190610520565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff021916908315150217905550606082015181600201908161014d9190610651565b5050505050610720565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6040518060800160405280606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f15158152602001606081525090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61028a82610261565b9050919050565b61029a81610280565b81146102a4575f80fd5b50565b5f815190506102b581610291565b92915050565b5f602082840312156102d0576102cf61025d565b5b5f6102dd848285016102a7565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061036157607f821691505b6020821081036103745761037361031d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103d67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261039b565b6103e0868361039b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61042461041f61041a846103f8565b610401565b6103f8565b9050919050565b5f819050919050565b61043d8361040a565b6104516104498261042b565b8484546103a7565b825550505050565b5f90565b610465610459565b610470818484610434565b505050565b5b81811015610493576104885f8261045d565b600181019050610476565b5050565b601f8211156104d8576104a98161037a565b6104b28461038c565b810160208510156104c1578190505b6104d56104cd8561038c565b830182610475565b50505b505050565b5f82821c905092915050565b5f6104f85f19846008026104dd565b1980831691505092915050565b5f61051083836104e9565b9150826002028217905092915050565b610529826102e6565b67ffffffffffffffff811115610542576105416102f0565b5b61054c825461034a565b610557828285610497565b5f60209050601f831160018114610588575f8415610576578287015190505b6105808582610505565b8655506105e7565b601f1984166105968661037a565b5f5b828110156105bd57848901518255600182019150602085019450602081019050610598565b868310156105da57848901516105d6601f8916826104e9565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f819050815f5260205f209050919050565b601f82111561064c5761061d816105f9565b6106268461038c565b81016020851015610635578190505b6106496106418561038c565b830182610475565b50505b505050565b61065a826105ef565b67ffffffffffffffff811115610673576106726102f0565b5b61067d825461034a565b61068882828561060b565b5f60209050601f8311600181146106b9575f84156106a7578287015190505b6106b18582610505565b865550610718565b601f1984166106c7866105f9565b5f5b828110156106ee578489015182556001820191506020850194506020810190506106c9565b8683101561070b5784890151610707601f8916826104e9565b8355505b6001600288020188555050505b505050505050565b6117e38061072d5f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c8063d370a37d11610064578063d370a37d14610132578063d48dca2d14610162578063ea1bbe351461017e578063f2fde38b146101ae578063f804843d146101ca5761009c565b80635af3cf0e146100a0578063715018a6146100bc57806376cdb03b146100c65780638da5cb5b146100e45780639f5591ab14610102575b5f80fd5b6100ba60048036038101906100b59190610cb5565b6101e6565b005b6100c4610423565b005b6100ce610436565b6040516100db9190610d6a565b60405180910390f35b6100ec61045b565b6040516100f99190610d6a565b60405180910390f35b61011c60048036038101906101179190610dc0565b610482565b6040516101299190610f47565b60405180910390f35b61014c60048036038101906101479190610dc0565b6106b5565b6040516101599190610d6a565b60405180910390f35b61017c60048036038101906101779190610f91565b61070a565b005b61019860048036038101906101939190610fcf565b61080d565b6040516101a59190611025565b60405180910390f35b6101c860048036038101906101c39190611068565b610847565b005b6101e460048036038101906101df9190610dc0565b6108c9565b005b5f6003836040516101f791906110cd565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614610263576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025a9061113d565b60405180910390fd5b5f6002805490509050610274610a8c565b83815f018190525033816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001819052506001816040019015159081151581525050600281908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f01908161030a919061135e565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160020190816103869190611485565b5050508160038560405161039a91906110cd565b90815260200160405180910390205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f1e0210396cbbc2a1ccae7d7b17e27a3fea0e56962c68454eda6fa9682bc7fe47838660405161041592919061158c565b60405180910390a250505050565b61042b610946565b6104345f6109c4565b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61048a610a8c565b604051806080016040528060028467ffffffffffffffff16815481106104b3576104b26115ba565b5b905f5260205f2090600302015f0180546104cc90611188565b80601f01602080910402602001604051908101604052809291908181526020018280546104f890611188565b80156105435780601f1061051a57610100808354040283529160200191610543565b820191905f5260205f20905b81548152906001019060200180831161052657829003601f168201915b5050505050815260200160028467ffffffffffffffff168154811061056b5761056a6115ba565b5b905f5260205f2090600302016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028467ffffffffffffffff16815481106105d3576105d26115ba565b5b905f5260205f20906003020160010160149054906101000a900460ff161515815260200160028467ffffffffffffffff1681548110610615576106146115ba565b5b905f5260205f209060030201600201805461062f90611188565b80601f016020809104026020016040519081016040528092919081815260200182805461065b90611188565b80156106a65780601f1061067d576101008083540402835291602001916106a6565b820191905f5260205f20905b81548152906001019060200180831161068957829003601f168201915b50505050508152509050919050565b5f60028267ffffffffffffffff16815481106106d4576106d36115ba565b5b905f5260205f2090600302016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660028367ffffffffffffffff168154811061073f5761073e6115ba565b5b905f5260205f2090600302016001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bb90611631565b60405180910390fd5b8060028367ffffffffffffffff16815481106107e3576107e26115ba565b5b905f5260205f20906003020160010160146101000a81548160ff0219169083151502179055505050565b5f60038260405161081e91906110cd565b90815260200160405180910390205f9054906101000a900467ffffffffffffffff169050919050565b61084f610946565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b4906116bf565b60405180910390fd5b6108c6816109c4565b50565b60028167ffffffffffffffff16815481106108e7576108e66115ba565b5b905f5260205f20906003020160010160149054906101000a900460ff16610943576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093a90611727565b60405180910390fd5b50565b61094e610a85565b73ffffffffffffffffffffffffffffffffffffffff1661096c61045b565b73ffffffffffffffffffffffffffffffffffffffff16146109c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b99061178f565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b6040518060800160405280606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f15158152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b2982610ae3565b810181811067ffffffffffffffff82111715610b4857610b47610af3565b5b80604052505050565b5f610b5a610aca565b9050610b668282610b20565b919050565b5f67ffffffffffffffff821115610b8557610b84610af3565b5b610b8e82610ae3565b9050602081019050919050565b828183375f83830152505050565b5f610bbb610bb684610b6b565b610b51565b905082815260208101848484011115610bd757610bd6610adf565b5b610be2848285610b9b565b509392505050565b5f82601f830112610bfe57610bfd610adb565b5b8135610c0e848260208601610ba9565b91505092915050565b5f67ffffffffffffffff821115610c3157610c30610af3565b5b610c3a82610ae3565b9050602081019050919050565b5f610c59610c5484610c17565b610b51565b905082815260208101848484011115610c7557610c74610adf565b5b610c80848285610b9b565b509392505050565b5f82601f830112610c9c57610c9b610adb565b5b8135610cac848260208601610c47565b91505092915050565b5f8060408385031215610ccb57610cca610ad3565b5b5f83013567ffffffffffffffff811115610ce857610ce7610ad7565b5b610cf485828601610bea565b925050602083013567ffffffffffffffff811115610d1557610d14610ad7565b5b610d2185828601610c88565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d5482610d2b565b9050919050565b610d6481610d4a565b82525050565b5f602082019050610d7d5f830184610d5b565b92915050565b5f67ffffffffffffffff82169050919050565b610d9f81610d83565b8114610da9575f80fd5b50565b5f81359050610dba81610d96565b92915050565b5f60208284031215610dd557610dd4610ad3565b5b5f610de284828501610dac565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e22578082015181840152602081019050610e07565b5f8484015250505050565b5f610e3782610deb565b610e418185610df5565b9350610e51818560208601610e05565b610e5a81610ae3565b840191505092915050565b610e6e81610d4a565b82525050565b5f8115159050919050565b610e8881610e74565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610eb282610e8e565b610ebc8185610e98565b9350610ecc818560208601610e05565b610ed581610ae3565b840191505092915050565b5f608083015f8301518482035f860152610efa8282610e2d565b9150506020830151610f0f6020860182610e65565b506040830151610f226040860182610e7f565b5060608301518482036060860152610f3a8282610ea8565b9150508091505092915050565b5f6020820190508181035f830152610f5f8184610ee0565b905092915050565b610f7081610e74565b8114610f7a575f80fd5b50565b5f81359050610f8b81610f67565b92915050565b5f8060408385031215610fa757610fa6610ad3565b5b5f610fb485828601610dac565b9250506020610fc585828601610f7d565b9150509250929050565b5f60208284031215610fe457610fe3610ad3565b5b5f82013567ffffffffffffffff81111561100157611000610ad7565b5b61100d84828501610bea565b91505092915050565b61101f81610d83565b82525050565b5f6020820190506110385f830184611016565b92915050565b61104781610d4a565b8114611051575f80fd5b50565b5f813590506110628161103e565b92915050565b5f6020828403121561107d5761107c610ad3565b5b5f61108a84828501611054565b91505092915050565b5f81905092915050565b5f6110a782610deb565b6110b18185611093565b93506110c1818560208601610e05565b80840191505092915050565b5f6110d8828461109d565b915081905092915050565b5f82825260208201905092915050565b7f65786973740000000000000000000000000000000000000000000000000000005f82015250565b5f6111276005836110e3565b9150611132826110f3565b602082019050919050565b5f6020820190508181035f8301526111548161111b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061119f57607f821691505b6020821081036111b2576111b161115b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026112147fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826111d9565b61121e86836111d9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61126261125d61125884611236565b61123f565b611236565b9050919050565b5f819050919050565b61127b83611248565b61128f61128782611269565b8484546111e5565b825550505050565b5f90565b6112a3611297565b6112ae818484611272565b505050565b5b818110156112d1576112c65f8261129b565b6001810190506112b4565b5050565b601f821115611316576112e7816111b8565b6112f0846111ca565b810160208510156112ff578190505b61131361130b856111ca565b8301826112b3565b50505b505050565b5f82821c905092915050565b5f6113365f198460080261131b565b1980831691505092915050565b5f61134e8383611327565b9150826002028217905092915050565b61136782610deb565b67ffffffffffffffff8111156113805761137f610af3565b5b61138a8254611188565b6113958282856112d5565b5f60209050601f8311600181146113c6575f84156113b4578287015190505b6113be8582611343565b865550611425565b601f1984166113d4866111b8565b5f5b828110156113fb578489015182556001820191506020850194506020810190506113d6565b868310156114185784890151611414601f891682611327565b8355505b6001600288020188555050505b505050505050565b5f819050815f5260205f209050919050565b601f821115611480576114518161142d565b61145a846111ca565b81016020851015611469578190505b61147d611475856111ca565b8301826112b3565b50505b505050565b61148e82610e8e565b67ffffffffffffffff8111156114a7576114a6610af3565b5b6114b18254611188565b6114bc82828561143f565b5f60209050601f8311600181146114ed575f84156114db578287015190505b6114e58582611343565b86555061154c565b601f1984166114fb8661142d565b5f5b82811015611522578489015182556001820191506020850194506020810190506114fd565b8683101561153f578489015161153b601f891682611327565b8355505b6001600288020188555050505b505050505050565b5f61155e82610deb565b61156881856110e3565b9350611578818560208601610e05565b61158181610ae3565b840191505092915050565b5f60408201905061159f5f830185611016565b81810360208301526115b18184611554565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f696e7600000000000000000000000000000000000000000000000000000000005f82015250565b5f61161b6003836110e3565b9150611626826115e7565b602082019050919050565b5f6020820190508181035f8301526116488161160f565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6116a96026836110e3565b91506116b48261164f565b604082019050919050565b5f6020820190508181035f8301526116d68161169d565b9050919050565b7f6e6f2061637469766500000000000000000000000000000000000000000000005f82015250565b5f6117116009836110e3565b915061171c826116dd565b602082019050919050565b5f6020820190508181035f83015261173e81611705565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6117796020836110e3565b915061178482611745565b602082019050919050565b5f6020820190508181035f8301526117a68161176d565b905091905056fea2646970667358221220347659fad05bb46d72d066d9431e7ce0af85a7b2077c0464cd41d6b71669428e64736f6c634300081a0033",
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
// Solidity: function getModel(uint64 _mi) view returns((string,address,bool,bytes) _info)
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
// Solidity: function getModel(uint64 _mi) view returns((string,address,bool,bytes) _info)
func (_Model *ModelSession) GetModel(_mi uint64) (IModelInfo, error) {
	return _Model.Contract.GetModel(&_Model.CallOpts, _mi)
}

// GetModel is a free data retrieval call binding the contract method 0x9f5591ab.
//
// Solidity: function getModel(uint64 _mi) view returns((string,address,bool,bytes) _info)
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

// Add is a paid mutator transaction binding the contract method 0x5af3cf0e.
//
// Solidity: function add(string _mn, bytes _rt) returns()
func (_Model *ModelTransactor) Add(opts *bind.TransactOpts, _mn string, _rt []byte) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "add", _mn, _rt)
}

// Add is a paid mutator transaction binding the contract method 0x5af3cf0e.
//
// Solidity: function add(string _mn, bytes _rt) returns()
func (_Model *ModelSession) Add(_mn string, _rt []byte) (*types.Transaction, error) {
	return _Model.Contract.Add(&_Model.TransactOpts, _mn, _rt)
}

// Add is a paid mutator transaction binding the contract method 0x5af3cf0e.
//
// Solidity: function add(string _mn, bytes _rt) returns()
func (_Model *ModelTransactorSession) Add(_mn string, _rt []byte) (*types.Transaction, error) {
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
