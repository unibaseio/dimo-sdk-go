// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package epoch

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

// EpochMetaData contains all meta data concerning the Epoch contract.
var EpochMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_e\",\"type\":\"uint64\"}],\"name\":\"SetEpoch\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"check\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"}],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slots\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405261012c60015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015610039575f80fd5b506100426100cf565b5f815f018181525050334260014361005a9190610120565b4060405160200161006d93929190610211565b60405160208183030381529060405280519060200120816020018181525050600281908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015550505061024d565b60405180604001604052805f81526020015f80191681525090565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61012a826100ea565b9150610135836100ea565b925082820390508181111561014d5761014c6100f3565b5b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61017c82610153565b9050919050565b5f8160601b9050919050565b5f61019982610183565b9050919050565b5f6101aa8261018f565b9050919050565b6101c26101bd82610172565b6101a0565b82525050565b5f819050919050565b6101e26101dd826100ea565b6101c8565b82525050565b5f819050919050565b5f819050919050565b61020b610206826101e8565b6101f1565b82525050565b5f61021c82866101b1565b60148201915061022c82856101d1565b60208201915061023c82846101fa565b602082019150819050949350505050565b6106a58061025a5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806312a02c821461004e57806348547d691461007f578063919840ad1461009d5780639fa6a6e3146100a7575b5f80fd5b610068600480360381019061006391906103c3565b6100c5565b60405161007692919061041e565b60405180910390f35b61008761012b565b6040516100949190610454565b60405180910390f35b6100a5610144565b005b6100af61014e565b6040516100bc9190610454565b60405180910390f35b5f8060028367ffffffffffffffff16815481106100e5576100e461046d565b5b905f5260205f2090600202015f015460028467ffffffffffffffff16815481106101125761011161046d565b5b905f5260205f2090600202016001015491509150915091565b60015f9054906101000a900467ffffffffffffffff1681565b61014c610168565b565b600160089054906101000a900467ffffffffffffffff1681565b5a5f8082825461017891906104c7565b9250508190555060015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff166002600160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16815481106101d3576101d261046d565b5b905f5260205f2090600202015f0154436101ed91906104fa565b10610365576101fa610367565b43815f0181815250506002600160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16815481106102375761023661046d565b5b905f5260205f20906002020160010154334260014361025691906104fa565b405f5460405160200161026d9594939291906105e2565b60405160208183030381529060405280519060200120816020018181525050600281908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015550506001600881819054906101000a900467ffffffffffffffff16809291906102ef90610640565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507f9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75600160089054906101000a900467ffffffffffffffff1660405161035b9190610454565b60405180910390a1505b565b60405180604001604052805f81526020015f80191681525090565b5f80fd5b5f67ffffffffffffffff82169050919050565b6103a281610386565b81146103ac575f80fd5b50565b5f813590506103bd81610399565b92915050565b5f602082840312156103d8576103d7610382565b5b5f6103e5848285016103af565b91505092915050565b5f819050919050565b610400816103ee565b82525050565b5f819050919050565b61041881610406565b82525050565b5f6040820190506104315f8301856103f7565b61043e602083018461040f565b9392505050565b61044e81610386565b82525050565b5f6020820190506104675f830184610445565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6104d1826103ee565b91506104dc836103ee565b92508282019050808211156104f4576104f361049a565b5b92915050565b5f610504826103ee565b915061050f836103ee565b92508282039050818111156105275761052661049a565b5b92915050565b5f819050919050565b61054761054282610406565b61052d565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105768261054d565b9050919050565b5f8160601b9050919050565b5f6105938261057d565b9050919050565b5f6105a482610589565b9050919050565b6105bc6105b78261056c565b61059a565b82525050565b5f819050919050565b6105dc6105d7826103ee565b6105c2565b82525050565b5f6105ed8288610536565b6020820191506105fd82876105ab565b60148201915061060d82866105cb565b60208201915061061d8285610536565b60208201915061062d82846105cb565b6020820191508190509695505050505050565b5f61064a82610386565b915067ffffffffffffffff82036106645761066361049a565b5b60018201905091905056fea2646970667358221220133d92f564f220213ba9e03f101523e92fdf202f00d989c4cafbb5bb9392a8c664736f6c634300081a0033",
}

// EpochABI is the input ABI used to generate the binding from.
// Deprecated: Use EpochMetaData.ABI instead.
var EpochABI = EpochMetaData.ABI

// EpochBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EpochMetaData.Bin instead.
var EpochBin = EpochMetaData.Bin

// DeployEpoch deploys a new Ethereum contract, binding an instance of Epoch to it.
func DeployEpoch(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Epoch, error) {
	parsed, err := EpochMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EpochBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Epoch{EpochCaller: EpochCaller{contract: contract}, EpochTransactor: EpochTransactor{contract: contract}, EpochFilterer: EpochFilterer{contract: contract}}, nil
}

// Epoch is an auto generated Go binding around an Ethereum contract.
type Epoch struct {
	EpochCaller     // Read-only binding to the contract
	EpochTransactor // Write-only binding to the contract
	EpochFilterer   // Log filterer for contract events
}

// EpochCaller is an auto generated read-only Go binding around an Ethereum contract.
type EpochCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EpochTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EpochFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EpochSession struct {
	Contract     *Epoch            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EpochCallerSession struct {
	Contract *EpochCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EpochTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EpochTransactorSession struct {
	Contract     *EpochTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochRaw is an auto generated low-level Go binding around an Ethereum contract.
type EpochRaw struct {
	Contract *Epoch // Generic contract binding to access the raw methods on
}

// EpochCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EpochCallerRaw struct {
	Contract *EpochCaller // Generic read-only contract binding to access the raw methods on
}

// EpochTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EpochTransactorRaw struct {
	Contract *EpochTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEpoch creates a new instance of Epoch, bound to a specific deployed contract.
func NewEpoch(address common.Address, backend bind.ContractBackend) (*Epoch, error) {
	contract, err := bindEpoch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Epoch{EpochCaller: EpochCaller{contract: contract}, EpochTransactor: EpochTransactor{contract: contract}, EpochFilterer: EpochFilterer{contract: contract}}, nil
}

// NewEpochCaller creates a new read-only instance of Epoch, bound to a specific deployed contract.
func NewEpochCaller(address common.Address, caller bind.ContractCaller) (*EpochCaller, error) {
	contract, err := bindEpoch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EpochCaller{contract: contract}, nil
}

// NewEpochTransactor creates a new write-only instance of Epoch, bound to a specific deployed contract.
func NewEpochTransactor(address common.Address, transactor bind.ContractTransactor) (*EpochTransactor, error) {
	contract, err := bindEpoch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EpochTransactor{contract: contract}, nil
}

// NewEpochFilterer creates a new log filterer instance of Epoch, bound to a specific deployed contract.
func NewEpochFilterer(address common.Address, filterer bind.ContractFilterer) (*EpochFilterer, error) {
	contract, err := bindEpoch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EpochFilterer{contract: contract}, nil
}

// bindEpoch binds a generic wrapper to an already deployed contract.
func bindEpoch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EpochMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Epoch *EpochRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Epoch.Contract.EpochCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Epoch *EpochRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.Contract.EpochTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Epoch *EpochRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Epoch.Contract.EpochTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Epoch *EpochCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Epoch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Epoch *EpochTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Epoch *EpochTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Epoch.Contract.contract.Transact(opts, method, params...)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochCaller) Current(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "current")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochSession) Current() (uint64, error) {
	return _Epoch.Contract.Current(&_Epoch.CallOpts)
}

// Current is a free data retrieval call binding the contract method 0x9fa6a6e3.
//
// Solidity: function current() view returns(uint64)
func (_Epoch *EpochCallerSession) Current() (uint64, error) {
	return _Epoch.Contract.Current(&_Epoch.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochCaller) GetEpoch(opts *bind.CallOpts, _epoch uint64) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "getEpoch", _epoch)

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochSession) GetEpoch(_epoch uint64) (*big.Int, [32]byte, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// GetEpoch is a free data retrieval call binding the contract method 0x12a02c82.
//
// Solidity: function getEpoch(uint64 _epoch) view returns(uint256, bytes32)
func (_Epoch *EpochCallerSession) GetEpoch(_epoch uint64) (*big.Int, [32]byte, error) {
	return _Epoch.Contract.GetEpoch(&_Epoch.CallOpts, _epoch)
}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochCaller) Slots(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Epoch.contract.Call(opts, &out, "slots")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochSession) Slots() (uint64, error) {
	return _Epoch.Contract.Slots(&_Epoch.CallOpts)
}

// Slots is a free data retrieval call binding the contract method 0x48547d69.
//
// Solidity: function slots() view returns(uint64)
func (_Epoch *EpochCallerSession) Slots() (uint64, error) {
	return _Epoch.Contract.Slots(&_Epoch.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochTransactor) Check(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Epoch.contract.Transact(opts, "check")
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochSession) Check() (*types.Transaction, error) {
	return _Epoch.Contract.Check(&_Epoch.TransactOpts)
}

// Check is a paid mutator transaction binding the contract method 0x919840ad.
//
// Solidity: function check() returns()
func (_Epoch *EpochTransactorSession) Check() (*types.Transaction, error) {
	return _Epoch.Contract.Check(&_Epoch.TransactOpts)
}

// EpochSetEpochIterator is returned from FilterSetEpoch and is used to iterate over the raw logs and unpacked data for SetEpoch events raised by the Epoch contract.
type EpochSetEpochIterator struct {
	Event *EpochSetEpoch // Event containing the contract specifics and raw log

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
func (it *EpochSetEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochSetEpoch)
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
		it.Event = new(EpochSetEpoch)
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
func (it *EpochSetEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochSetEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochSetEpoch represents a SetEpoch event raised by the Epoch contract.
type EpochSetEpoch struct {
	E   uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetEpoch is a free log retrieval operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_Epoch *EpochFilterer) FilterSetEpoch(opts *bind.FilterOpts) (*EpochSetEpochIterator, error) {

	logs, sub, err := _Epoch.contract.FilterLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return &EpochSetEpochIterator{contract: _Epoch.contract, event: "SetEpoch", logs: logs, sub: sub}, nil
}

// WatchSetEpoch is a free log subscription operation binding the contract event 0x9d4ccb161ea809df14334516cc3070025c80baddb8e8364afaca6a6fe31bfd75.
//
// Solidity: event SetEpoch(uint64 _e)
func (_Epoch *EpochFilterer) WatchSetEpoch(opts *bind.WatchOpts, sink chan<- *EpochSetEpoch) (event.Subscription, error) {

	logs, sub, err := _Epoch.contract.WatchLogs(opts, "SetEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochSetEpoch)
				if err := _Epoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
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
func (_Epoch *EpochFilterer) ParseSetEpoch(log types.Log) (*EpochSetEpoch, error) {
	event := new(EpochSetEpoch)
	if err := _Epoch.contract.UnpackLog(event, "SetEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
