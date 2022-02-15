// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancegetter

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
)

// BalancegetterMetaData contains all meta data concerning the Balancegetter contract.
var BalancegetterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"readBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BalancegetterABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancegetterMetaData.ABI instead.
var BalancegetterABI = BalancegetterMetaData.ABI

// Balancegetter is an auto generated Go binding around an Ethereum contract.
type Balancegetter struct {
	BalancegetterCaller     // Read-only binding to the contract
	BalancegetterTransactor // Write-only binding to the contract
	BalancegetterFilterer   // Log filterer for contract events
}

// BalancegetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancegetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancegetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancegetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancegetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancegetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancegetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancegetterSession struct {
	Contract     *Balancegetter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancegetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancegetterCallerSession struct {
	Contract *BalancegetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BalancegetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancegetterTransactorSession struct {
	Contract     *BalancegetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BalancegetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancegetterRaw struct {
	Contract *Balancegetter // Generic contract binding to access the raw methods on
}

// BalancegetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancegetterCallerRaw struct {
	Contract *BalancegetterCaller // Generic read-only contract binding to access the raw methods on
}

// BalancegetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancegetterTransactorRaw struct {
	Contract *BalancegetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancegetter creates a new instance of Balancegetter, bound to a specific deployed contract.
func NewBalancegetter(address common.Address, backend bind.ContractBackend) (*Balancegetter, error) {
	contract, err := bindBalancegetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancegetter{BalancegetterCaller: BalancegetterCaller{contract: contract}, BalancegetterTransactor: BalancegetterTransactor{contract: contract}, BalancegetterFilterer: BalancegetterFilterer{contract: contract}}, nil
}

// NewBalancegetterCaller creates a new read-only instance of Balancegetter, bound to a specific deployed contract.
func NewBalancegetterCaller(address common.Address, caller bind.ContractCaller) (*BalancegetterCaller, error) {
	contract, err := bindBalancegetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancegetterCaller{contract: contract}, nil
}

// NewBalancegetterTransactor creates a new write-only instance of Balancegetter, bound to a specific deployed contract.
func NewBalancegetterTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancegetterTransactor, error) {
	contract, err := bindBalancegetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancegetterTransactor{contract: contract}, nil
}

// NewBalancegetterFilterer creates a new log filterer instance of Balancegetter, bound to a specific deployed contract.
func NewBalancegetterFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancegetterFilterer, error) {
	contract, err := bindBalancegetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancegetterFilterer{contract: contract}, nil
}

// bindBalancegetter binds a generic wrapper to an already deployed contract.
func bindBalancegetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancegetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancegetter *BalancegetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancegetter.Contract.BalancegetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancegetter *BalancegetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancegetter.Contract.BalancegetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancegetter *BalancegetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancegetter.Contract.BalancegetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancegetter *BalancegetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancegetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancegetter *BalancegetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancegetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancegetter *BalancegetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancegetter.Contract.contract.Transact(opts, method, params...)
}

// ReadBalances is a paid mutator transaction binding the contract method 0x8f82a651.
//
// Solidity: function readBalances(address token, address[] addresses) returns(uint256[])
func (_Balancegetter *BalancegetterTransactor) ReadBalances(opts *bind.TransactOpts, token common.Address, addresses []common.Address) (*types.Transaction, error) {
	return _Balancegetter.contract.Transact(opts, "readBalances", token, addresses)
}

// ReadBalances is a paid mutator transaction binding the contract method 0x8f82a651.
//
// Solidity: function readBalances(address token, address[] addresses) returns(uint256[])
func (_Balancegetter *BalancegetterSession) ReadBalances(token common.Address, addresses []common.Address) (*types.Transaction, error) {
	return _Balancegetter.Contract.ReadBalances(&_Balancegetter.TransactOpts, token, addresses)
}

// ReadBalances is a paid mutator transaction binding the contract method 0x8f82a651.
//
// Solidity: function readBalances(address token, address[] addresses) returns(uint256[])
func (_Balancegetter *BalancegetterTransactorSession) ReadBalances(token common.Address, addresses []common.Address) (*types.Transaction, error) {
	return _Balancegetter.Contract.ReadBalances(&_Balancegetter.TransactOpts, token, addresses)
}
