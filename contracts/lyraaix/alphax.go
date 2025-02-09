// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lyraAiX

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

// lyraAiXCheckInInfo is an auto generated low-level Go binding around an user-defined struct.
type lyraAiXCheckInInfo struct {
	TaskId    uint32
	UserId    uint64
	Timestamp *big.Int
}

// lyraAiXPredictionInfo is an auto generated low-level Go binding around an user-defined struct.
type lyraAiXPredictionInfo struct {
	SignalId    uint32
	UserId      uint64
	Choice      uint8
	HasInvolved bool
}

// lyraAiXMetaData contains all meta data concerning the lyraAiX contract.
var lyraAiXMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structlyraAiX.CheckInInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"CheckinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"signalId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"choice\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"hasInvolved\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structlyraAiX.PredictionInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"SignalPredictionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"checkInResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"taskId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"}],\"name\":\"checkin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"signalId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"choice\",\"type\":\"uint8\"}],\"name\":\"signalPredict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"signalId\",\"type\":\"uint32\"}],\"name\":\"signalPredictionResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// lyraAiXABI is the input ABI used to generate the binding from.
// Deprecated: Use lyraAiXMetaData.ABI instead.
var lyraAiXABI = lyraAiXMetaData.ABI

// lyraAiX is an auto generated Go binding around an Ethereum contract.
type lyraAiX struct {
	lyraAiXCaller     // Read-only binding to the contract
	lyraAiXTransactor // Write-only binding to the contract
	lyraAiXFilterer   // Log filterer for contract events
}

// lyraAiXCaller is an auto generated read-only Go binding around an Ethereum contract.
type lyraAiXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// lyraAiXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type lyraAiXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// lyraAiXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type lyraAiXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// lyraAiXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type lyraAiXSession struct {
	Contract     *lyraAiX           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// lyraAiXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type lyraAiXCallerSession struct {
	Contract *lyraAiXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// lyraAiXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type lyraAiXTransactorSession struct {
	Contract     *lyraAiXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// lyraAiXRaw is an auto generated low-level Go binding around an Ethereum contract.
type lyraAiXRaw struct {
	Contract *lyraAiX // Generic contract binding to access the raw methods on
}

// lyraAiXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type lyraAiXCallerRaw struct {
	Contract *lyraAiXCaller // Generic read-only contract binding to access the raw methods on
}

// lyraAiXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type lyraAiXTransactorRaw struct {
	Contract *lyraAiXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewlyraAiX creates a new instance of lyraAiX, bound to a specific deployed contract.
func NewlyraAiX(address common.Address, backend bind.ContractBackend) (*lyraAiX, error) {
	contract, err := bindlyraAiX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &lyraAiX{lyraAiXCaller: lyraAiXCaller{contract: contract}, lyraAiXTransactor: lyraAiXTransactor{contract: contract}, lyraAiXFilterer: lyraAiXFilterer{contract: contract}}, nil
}

// NewlyraAiXCaller creates a new read-only instance of lyraAiX, bound to a specific deployed contract.
func NewlyraAiXCaller(address common.Address, caller bind.ContractCaller) (*lyraAiXCaller, error) {
	contract, err := bindlyraAiX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &lyraAiXCaller{contract: contract}, nil
}

// NewlyraAiXTransactor creates a new write-only instance of lyraAiX, bound to a specific deployed contract.
func NewlyraAiXTransactor(address common.Address, transactor bind.ContractTransactor) (*lyraAiXTransactor, error) {
	contract, err := bindlyraAiX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &lyraAiXTransactor{contract: contract}, nil
}

// NewlyraAiXFilterer creates a new log filterer instance of lyraAiX, bound to a specific deployed contract.
func NewlyraAiXFilterer(address common.Address, filterer bind.ContractFilterer) (*lyraAiXFilterer, error) {
	contract, err := bindlyraAiX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &lyraAiXFilterer{contract: contract}, nil
}

// bindlyraAiX binds a generic wrapper to an already deployed contract.
func bindlyraAiX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := lyraAiXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_lyraAiX *lyraAiXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _lyraAiX.Contract.lyraAiXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_lyraAiX *lyraAiXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _lyraAiX.Contract.lyraAiXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_lyraAiX *lyraAiXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _lyraAiX.Contract.lyraAiXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_lyraAiX *lyraAiXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _lyraAiX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_lyraAiX *lyraAiXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _lyraAiX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_lyraAiX *lyraAiXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _lyraAiX.Contract.contract.Transact(opts, method, params...)
}

// CheckInResult is a free data retrieval call binding the contract method 0x2c3ae4a7.
//
// Solidity: function checkInResult(address user, uint256 day) view returns(bool)
func (_lyraAiX *lyraAiXCaller) CheckInResult(opts *bind.CallOpts, user common.Address, day *big.Int) (bool, error) {
	var out []interface{}
	err := _lyraAiX.contract.Call(opts, &out, "checkInResult", user, day)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckInResult is a free data retrieval call binding the contract method 0x2c3ae4a7.
//
// Solidity: function checkInResult(address user, uint256 day) view returns(bool)
func (_lyraAiX *lyraAiXSession) CheckInResult(user common.Address, day *big.Int) (bool, error) {
	return _lyraAiX.Contract.CheckInResult(&_lyraAiX.CallOpts, user, day)
}

// CheckInResult is a free data retrieval call binding the contract method 0x2c3ae4a7.
//
// Solidity: function checkInResult(address user, uint256 day) view returns(bool)
func (_lyraAiX *lyraAiXCallerSession) CheckInResult(user common.Address, day *big.Int) (bool, error) {
	return _lyraAiX.Contract.CheckInResult(&_lyraAiX.CallOpts, user, day)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_lyraAiX *lyraAiXCaller) GetCurrentDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _lyraAiX.contract.Call(opts, &out, "getCurrentDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_lyraAiX *lyraAiXSession) GetCurrentDay() (*big.Int, error) {
	return _lyraAiX.Contract.GetCurrentDay(&_lyraAiX.CallOpts)
}

// GetCurrentDay is a free data retrieval call binding the contract method 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (_lyraAiX *lyraAiXCallerSession) GetCurrentDay() (*big.Int, error) {
	return _lyraAiX.Contract.GetCurrentDay(&_lyraAiX.CallOpts)
}

// SignalPredictionResult is a free data retrieval call binding the contract method 0x703a7fa2.
//
// Solidity: function signalPredictionResult(address user, uint32 signalId) view returns(bool, uint32, uint8)
func (_lyraAiX *lyraAiXCaller) SignalPredictionResult(opts *bind.CallOpts, user common.Address, signalId uint32) (bool, uint32, uint8, error) {
	var out []interface{}
	err := _lyraAiX.contract.Call(opts, &out, "signalPredictionResult", user, signalId)

	if err != nil {
		return *new(bool), *new(uint32), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return out0, out1, out2, err

}

// SignalPredictionResult is a free data retrieval call binding the contract method 0x703a7fa2.
//
// Solidity: function signalPredictionResult(address user, uint32 signalId) view returns(bool, uint32, uint8)
func (_lyraAiX *lyraAiXSession) SignalPredictionResult(user common.Address, signalId uint32) (bool, uint32, uint8, error) {
	return _lyraAiX.Contract.SignalPredictionResult(&_lyraAiX.CallOpts, user, signalId)
}

// SignalPredictionResult is a free data retrieval call binding the contract method 0x703a7fa2.
//
// Solidity: function signalPredictionResult(address user, uint32 signalId) view returns(bool, uint32, uint8)
func (_lyraAiX *lyraAiXCallerSession) SignalPredictionResult(user common.Address, signalId uint32) (bool, uint32, uint8, error) {
	return _lyraAiX.Contract.SignalPredictionResult(&_lyraAiX.CallOpts, user, signalId)
}

// Checkin is a paid mutator transaction binding the contract method 0x0655e4ca.
//
// Solidity: function checkin(address user, uint32 taskId, uint64 userId) returns()
func (_lyraAiX *lyraAiXTransactor) Checkin(opts *bind.TransactOpts, user common.Address, taskId uint32, userId uint64) (*types.Transaction, error) {
	return _lyraAiX.contract.Transact(opts, "checkin", user, taskId, userId)
}

// Checkin is a paid mutator transaction binding the contract method 0x0655e4ca.
//
// Solidity: function checkin(address user, uint32 taskId, uint64 userId) returns()
func (_lyraAiX *lyraAiXSession) Checkin(user common.Address, taskId uint32, userId uint64) (*types.Transaction, error) {
	return _lyraAiX.Contract.Checkin(&_lyraAiX.TransactOpts, user, taskId, userId)
}

// Checkin is a paid mutator transaction binding the contract method 0x0655e4ca.
//
// Solidity: function checkin(address user, uint32 taskId, uint64 userId) returns()
func (_lyraAiX *lyraAiXTransactorSession) Checkin(user common.Address, taskId uint32, userId uint64) (*types.Transaction, error) {
	return _lyraAiX.Contract.Checkin(&_lyraAiX.TransactOpts, user, taskId, userId)
}

// SignalPredict is a paid mutator transaction binding the contract method 0x591abd76.
//
// Solidity: function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) returns()
func (_lyraAiX *lyraAiXTransactor) SignalPredict(opts *bind.TransactOpts, user common.Address, signalId uint32, userId uint64, choice uint8) (*types.Transaction, error) {
	return _lyraAiX.contract.Transact(opts, "signalPredict", user, signalId, userId, choice)
}

// SignalPredict is a paid mutator transaction binding the contract method 0x591abd76.
//
// Solidity: function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) returns()
func (_lyraAiX *lyraAiXSession) SignalPredict(user common.Address, signalId uint32, userId uint64, choice uint8) (*types.Transaction, error) {
	return _lyraAiX.Contract.SignalPredict(&_lyraAiX.TransactOpts, user, signalId, userId, choice)
}

// SignalPredict is a paid mutator transaction binding the contract method 0x591abd76.
//
// Solidity: function signalPredict(address user, uint32 signalId, uint64 userId, uint8 choice) returns()
func (_lyraAiX *lyraAiXTransactorSession) SignalPredict(user common.Address, signalId uint32, userId uint64, choice uint8) (*types.Transaction, error) {
	return _lyraAiX.Contract.SignalPredict(&_lyraAiX.TransactOpts, user, signalId, userId, choice)
}

// lyraAiXCheckinEventIterator is returned from FilterCheckinEvent and is used to iterate over the raw logs and unpacked data for CheckinEvent events raised by the lyraAiX contract.
type lyraAiXCheckinEventIterator struct {
	Event *lyraAiXCheckinEvent // Event containing the contract specifics and raw log

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
func (it *lyraAiXCheckinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(lyraAiXCheckinEvent)
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
		it.Event = new(lyraAiXCheckinEvent)
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
func (it *lyraAiXCheckinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *lyraAiXCheckinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// lyraAiXCheckinEvent represents a CheckinEvent event raised by the lyraAiX contract.
type lyraAiXCheckinEvent struct {
	User common.Address
	Info lyraAiXCheckInInfo
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCheckinEvent is a free log retrieval operation binding the contract event 0x2f384bbc453afcdcb37abef73843b9341c84045a682ee5e170a885d16b660dae.
//
// Solidity: event CheckinEvent(address indexed user, (uint32,uint64,uint256) info)
func (_lyraAiX *lyraAiXFilterer) FilterCheckinEvent(opts *bind.FilterOpts, user []common.Address) (*lyraAiXCheckinEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _lyraAiX.contract.FilterLogs(opts, "CheckinEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &lyraAiXCheckinEventIterator{contract: _lyraAiX.contract, event: "CheckinEvent", logs: logs, sub: sub}, nil
}

// WatchCheckinEvent is a free log subscription operation binding the contract event 0x2f384bbc453afcdcb37abef73843b9341c84045a682ee5e170a885d16b660dae.
//
// Solidity: event CheckinEvent(address indexed user, (uint32,uint64,uint256) info)
func (_lyraAiX *lyraAiXFilterer) WatchCheckinEvent(opts *bind.WatchOpts, sink chan<- *lyraAiXCheckinEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _lyraAiX.contract.WatchLogs(opts, "CheckinEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(lyraAiXCheckinEvent)
				if err := _lyraAiX.contract.UnpackLog(event, "CheckinEvent", log); err != nil {
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

// ParseCheckinEvent is a log parse operation binding the contract event 0x2f384bbc453afcdcb37abef73843b9341c84045a682ee5e170a885d16b660dae.
//
// Solidity: event CheckinEvent(address indexed user, (uint32,uint64,uint256) info)
func (_lyraAiX *lyraAiXFilterer) ParseCheckinEvent(log types.Log) (*lyraAiXCheckinEvent, error) {
	event := new(lyraAiXCheckinEvent)
	if err := _lyraAiX.contract.UnpackLog(event, "CheckinEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// lyraAiXSignalPredictionEventIterator is returned from FilterSignalPredictionEvent and is used to iterate over the raw logs and unpacked data for SignalPredictionEvent events raised by the lyraAiX contract.
type lyraAiXSignalPredictionEventIterator struct {
	Event *lyraAiXSignalPredictionEvent // Event containing the contract specifics and raw log

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
func (it *lyraAiXSignalPredictionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(lyraAiXSignalPredictionEvent)
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
		it.Event = new(lyraAiXSignalPredictionEvent)
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
func (it *lyraAiXSignalPredictionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *lyraAiXSignalPredictionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// lyraAiXSignalPredictionEvent represents a SignalPredictionEvent event raised by the lyraAiX contract.
type lyraAiXSignalPredictionEvent struct {
	User common.Address
	Info lyraAiXPredictionInfo
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSignalPredictionEvent is a free log retrieval operation binding the contract event 0xbe0b3fe34eb803bd35bd0e6883956b86e4bc611efcd52b146df987bad9e18ba1.
//
// Solidity: event SignalPredictionEvent(address indexed user, (uint32,uint64,uint8,bool) info)
func (_lyraAiX *lyraAiXFilterer) FilterSignalPredictionEvent(opts *bind.FilterOpts, user []common.Address) (*lyraAiXSignalPredictionEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _lyraAiX.contract.FilterLogs(opts, "SignalPredictionEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &lyraAiXSignalPredictionEventIterator{contract: _lyraAiX.contract, event: "SignalPredictionEvent", logs: logs, sub: sub}, nil
}

// WatchSignalPredictionEvent is a free log subscription operation binding the contract event 0xbe0b3fe34eb803bd35bd0e6883956b86e4bc611efcd52b146df987bad9e18ba1.
//
// Solidity: event SignalPredictionEvent(address indexed user, (uint32,uint64,uint8,bool) info)
func (_lyraAiX *lyraAiXFilterer) WatchSignalPredictionEvent(opts *bind.WatchOpts, sink chan<- *lyraAiXSignalPredictionEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _lyraAiX.contract.WatchLogs(opts, "SignalPredictionEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(lyraAiXSignalPredictionEvent)
				if err := _lyraAiX.contract.UnpackLog(event, "SignalPredictionEvent", log); err != nil {
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

// ParseSignalPredictionEvent is a log parse operation binding the contract event 0xbe0b3fe34eb803bd35bd0e6883956b86e4bc611efcd52b146df987bad9e18ba1.
//
// Solidity: event SignalPredictionEvent(address indexed user, (uint32,uint64,uint8,bool) info)
func (_lyraAiX *lyraAiXFilterer) ParseSignalPredictionEvent(log types.Log) (*lyraAiXSignalPredictionEvent, error) {
	event := new(lyraAiXSignalPredictionEvent)
	if err := _lyraAiX.contract.UnpackLog(event, "SignalPredictionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
