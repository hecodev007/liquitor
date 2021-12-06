// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquitor

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

// LiquidtorABI is the input ABI used to generate the binding from.
const LiquidtorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"LogAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"flag\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"mgs\",\"type\":\"string\"}],\"name\":\"LogLiquite\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"RecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RevenueWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptroller\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"recipientChange\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"waitingPeriodEnd\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"initiateRecipientChange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmRecipientChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_comptrollerAddress\",\"type\":\"address\"}],\"name\":\"setComptroller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"setParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasThreshold\",\"type\":\"uint256\"}],\"name\":\"setGasThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_messages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"_borrowers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_cTokens\",\"type\":\"address[]\"}],\"name\":\"liquidateSNWithPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_borrowers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_cTokens\",\"type\":\"address[]\"}],\"name\":\"liquidateSN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_messages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_repayCToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seizeCToken\",\"type\":\"address\"}],\"name\":\"liquidateSWithPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_repayCToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seizeCToken\",\"type\":\"address\"}],\"name\":\"liquidateS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_repayCToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seizeCToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_repayCToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seizeCToken\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_repayCToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seizeCToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidateTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_repayCToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seizeCToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SwapTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Liquidtor is an auto generated Go binding around an Ethereum contract.
type Liquidtor struct {
	LiquidtorCaller     // Read-only binding to the contract
	LiquidtorTransactor // Write-only binding to the contract
	LiquidtorFilterer   // Log filterer for contract events
}

// LiquidtorCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidtorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidtorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidtorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidtorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidtorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidtorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidtorSession struct {
	Contract     *Liquidtor        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidtorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidtorCallerSession struct {
	Contract *LiquidtorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LiquidtorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidtorTransactorSession struct {
	Contract     *LiquidtorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LiquidtorRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidtorRaw struct {
	Contract *Liquidtor // Generic contract binding to access the raw methods on
}

// LiquidtorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidtorCallerRaw struct {
	Contract *LiquidtorCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidtorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidtorTransactorRaw struct {
	Contract *LiquidtorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidtor creates a new instance of Liquidtor, bound to a specific deployed contract.
func NewLiquidtor(address common.Address, backend bind.ContractBackend) (*Liquidtor, error) {
	contract, err := bindLiquidtor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liquidtor{LiquidtorCaller: LiquidtorCaller{contract: contract}, LiquidtorTransactor: LiquidtorTransactor{contract: contract}, LiquidtorFilterer: LiquidtorFilterer{contract: contract}}, nil
}

// NewLiquidtorCaller creates a new read-only instance of Liquidtor, bound to a specific deployed contract.
func NewLiquidtorCaller(address common.Address, caller bind.ContractCaller) (*LiquidtorCaller, error) {
	contract, err := bindLiquidtor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidtorCaller{contract: contract}, nil
}

// NewLiquidtorTransactor creates a new write-only instance of Liquidtor, bound to a specific deployed contract.
func NewLiquidtorTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidtorTransactor, error) {
	contract, err := bindLiquidtor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidtorTransactor{contract: contract}, nil
}

// NewLiquidtorFilterer creates a new log filterer instance of Liquidtor, bound to a specific deployed contract.
func NewLiquidtorFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidtorFilterer, error) {
	contract, err := bindLiquidtor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidtorFilterer{contract: contract}, nil
}

// bindLiquidtor binds a generic wrapper to an already deployed contract.
func bindLiquidtor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidtorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidtor *LiquidtorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidtor.Contract.LiquidtorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidtor *LiquidtorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidtorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidtor *LiquidtorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidtorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidtor *LiquidtorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidtor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidtor *LiquidtorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidtor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidtor *LiquidtorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidtor.Contract.contract.Transact(opts, method, params...)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Liquidtor *LiquidtorCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidtor.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Liquidtor *LiquidtorSession) Comptroller() (common.Address, error) {
	return _Liquidtor.Contract.Comptroller(&_Liquidtor.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Liquidtor *LiquidtorCallerSession) Comptroller() (common.Address, error) {
	return _Liquidtor.Contract.Comptroller(&_Liquidtor.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address _repayCToken, address _seizeCToken) view returns(address)
func (_Liquidtor *LiquidtorCaller) GetPair(opts *bind.CallOpts, _repayCToken common.Address, _seizeCToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Liquidtor.contract.Call(opts, &out, "getPair", _repayCToken, _seizeCToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address _repayCToken, address _seizeCToken) view returns(address)
func (_Liquidtor *LiquidtorSession) GetPair(_repayCToken common.Address, _seizeCToken common.Address) (common.Address, error) {
	return _Liquidtor.Contract.GetPair(&_Liquidtor.CallOpts, _repayCToken, _seizeCToken)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address _repayCToken, address _seizeCToken) view returns(address)
func (_Liquidtor *LiquidtorCallerSession) GetPair(_repayCToken common.Address, _seizeCToken common.Address) (common.Address, error) {
	return _Liquidtor.Contract.GetPair(&_Liquidtor.CallOpts, _repayCToken, _seizeCToken)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Liquidtor *LiquidtorCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidtor.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Liquidtor *LiquidtorSession) PriceOracle() (common.Address, error) {
	return _Liquidtor.Contract.PriceOracle(&_Liquidtor.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Liquidtor *LiquidtorCallerSession) PriceOracle() (common.Address, error) {
	return _Liquidtor.Contract.PriceOracle(&_Liquidtor.CallOpts)
}

// RecipientChange is a free data retrieval call binding the contract method 0xe4b56e74.
//
// Solidity: function recipientChange() view returns(address recipient, uint256 waitingPeriodEnd, bool pending)
func (_Liquidtor *LiquidtorCaller) RecipientChange(opts *bind.CallOpts) (struct {
	Recipient        common.Address
	WaitingPeriodEnd *big.Int
	Pending          bool
}, error) {
	var out []interface{}
	err := _Liquidtor.contract.Call(opts, &out, "recipientChange")

	outstruct := new(struct {
		Recipient        common.Address
		WaitingPeriodEnd *big.Int
		Pending          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.WaitingPeriodEnd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Pending = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// RecipientChange is a free data retrieval call binding the contract method 0xe4b56e74.
//
// Solidity: function recipientChange() view returns(address recipient, uint256 waitingPeriodEnd, bool pending)
func (_Liquidtor *LiquidtorSession) RecipientChange() (struct {
	Recipient        common.Address
	WaitingPeriodEnd *big.Int
	Pending          bool
}, error) {
	return _Liquidtor.Contract.RecipientChange(&_Liquidtor.CallOpts)
}

// RecipientChange is a free data retrieval call binding the contract method 0xe4b56e74.
//
// Solidity: function recipientChange() view returns(address recipient, uint256 waitingPeriodEnd, bool pending)
func (_Liquidtor *LiquidtorCallerSession) RecipientChange() (struct {
	Recipient        common.Address
	WaitingPeriodEnd *big.Int
	Pending          bool
}, error) {
	return _Liquidtor.Contract.RecipientChange(&_Liquidtor.CallOpts)
}

// SwapTest is a paid mutator transaction binding the contract method 0x775201b9.
//
// Solidity: function SwapTest(address pair, address r, address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorTransactor) SwapTest(opts *bind.TransactOpts, pair common.Address, r common.Address, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "SwapTest", pair, r, _borrower, _repayCToken, _seizeCToken, _amount)
}

// SwapTest is a paid mutator transaction binding the contract method 0x775201b9.
//
// Solidity: function SwapTest(address pair, address r, address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorSession) SwapTest(pair common.Address, r common.Address, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.SwapTest(&_Liquidtor.TransactOpts, pair, r, _borrower, _repayCToken, _seizeCToken, _amount)
}

// SwapTest is a paid mutator transaction binding the contract method 0x775201b9.
//
// Solidity: function SwapTest(address pair, address r, address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorTransactorSession) SwapTest(pair common.Address, r common.Address, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.SwapTest(&_Liquidtor.TransactOpts, pair, r, _borrower, _repayCToken, _seizeCToken, _amount)
}

// ConfirmRecipientChange is a paid mutator transaction binding the contract method 0x1c15d500.
//
// Solidity: function confirmRecipientChange() returns()
func (_Liquidtor *LiquidtorTransactor) ConfirmRecipientChange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "confirmRecipientChange")
}

// ConfirmRecipientChange is a paid mutator transaction binding the contract method 0x1c15d500.
//
// Solidity: function confirmRecipientChange() returns()
func (_Liquidtor *LiquidtorSession) ConfirmRecipientChange() (*types.Transaction, error) {
	return _Liquidtor.Contract.ConfirmRecipientChange(&_Liquidtor.TransactOpts)
}

// ConfirmRecipientChange is a paid mutator transaction binding the contract method 0x1c15d500.
//
// Solidity: function confirmRecipientChange() returns()
func (_Liquidtor *LiquidtorTransactorSession) ConfirmRecipientChange() (*types.Transaction, error) {
	return _Liquidtor.Contract.ConfirmRecipientChange(&_Liquidtor.TransactOpts)
}

// InitiateRecipientChange is a paid mutator transaction binding the contract method 0x102961ca.
//
// Solidity: function initiateRecipientChange(address _recipient) returns(address)
func (_Liquidtor *LiquidtorTransactor) InitiateRecipientChange(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "initiateRecipientChange", _recipient)
}

// InitiateRecipientChange is a paid mutator transaction binding the contract method 0x102961ca.
//
// Solidity: function initiateRecipientChange(address _recipient) returns(address)
func (_Liquidtor *LiquidtorSession) InitiateRecipientChange(_recipient common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.InitiateRecipientChange(&_Liquidtor.TransactOpts, _recipient)
}

// InitiateRecipientChange is a paid mutator transaction binding the contract method 0x102961ca.
//
// Solidity: function initiateRecipientChange(address _recipient) returns(address)
func (_Liquidtor *LiquidtorTransactorSession) InitiateRecipientChange(_recipient common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.InitiateRecipientChange(&_Liquidtor.TransactOpts, _recipient)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Liquidtor *LiquidtorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Liquidtor *LiquidtorSession) Kill() (*types.Transaction, error) {
	return _Liquidtor.Contract.Kill(&_Liquidtor.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Liquidtor *LiquidtorTransactorSession) Kill() (*types.Transaction, error) {
	return _Liquidtor.Contract.Kill(&_Liquidtor.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0xaab3f868.
//
// Solidity: function liquidate(address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorTransactor) Liquidate(opts *bind.TransactOpts, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "liquidate", _borrower, _repayCToken, _seizeCToken, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xaab3f868.
//
// Solidity: function liquidate(address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorSession) Liquidate(_borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.Liquidate(&_Liquidtor.TransactOpts, _borrower, _repayCToken, _seizeCToken, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xaab3f868.
//
// Solidity: function liquidate(address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorTransactorSession) Liquidate(_borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.Liquidate(&_Liquidtor.TransactOpts, _borrower, _repayCToken, _seizeCToken, _amount)
}

// LiquidateS is a paid mutator transaction binding the contract method 0x4693c8e6.
//
// Solidity: function liquidateS(address _borrower, address _repayCToken, address _seizeCToken) returns()
func (_Liquidtor *LiquidtorTransactor) LiquidateS(opts *bind.TransactOpts, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "liquidateS", _borrower, _repayCToken, _seizeCToken)
}

// LiquidateS is a paid mutator transaction binding the contract method 0x4693c8e6.
//
// Solidity: function liquidateS(address _borrower, address _repayCToken, address _seizeCToken) returns()
func (_Liquidtor *LiquidtorSession) LiquidateS(_borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateS(&_Liquidtor.TransactOpts, _borrower, _repayCToken, _seizeCToken)
}

// LiquidateS is a paid mutator transaction binding the contract method 0x4693c8e6.
//
// Solidity: function liquidateS(address _borrower, address _repayCToken, address _seizeCToken) returns()
func (_Liquidtor *LiquidtorTransactorSession) LiquidateS(_borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateS(&_Liquidtor.TransactOpts, _borrower, _repayCToken, _seizeCToken)
}

// LiquidateSN is a paid mutator transaction binding the contract method 0xe40b65b5.
//
// Solidity: function liquidateSN(address[] _borrowers, address[] _cTokens) returns()
func (_Liquidtor *LiquidtorTransactor) LiquidateSN(opts *bind.TransactOpts, _borrowers []common.Address, _cTokens []common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "liquidateSN", _borrowers, _cTokens)
}

// LiquidateSN is a paid mutator transaction binding the contract method 0xe40b65b5.
//
// Solidity: function liquidateSN(address[] _borrowers, address[] _cTokens) returns()
func (_Liquidtor *LiquidtorSession) LiquidateSN(_borrowers []common.Address, _cTokens []common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateSN(&_Liquidtor.TransactOpts, _borrowers, _cTokens)
}

// LiquidateSN is a paid mutator transaction binding the contract method 0xe40b65b5.
//
// Solidity: function liquidateSN(address[] _borrowers, address[] _cTokens) returns()
func (_Liquidtor *LiquidtorTransactorSession) LiquidateSN(_borrowers []common.Address, _cTokens []common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateSN(&_Liquidtor.TransactOpts, _borrowers, _cTokens)
}

// LiquidateSNWithPrice is a paid mutator transaction binding the contract method 0x52fae007.
//
// Solidity: function liquidateSNWithPrice(bytes[] _messages, bytes[] _signatures, string[] _symbols, address[] _borrowers, address[] _cTokens) returns()
func (_Liquidtor *LiquidtorTransactor) LiquidateSNWithPrice(opts *bind.TransactOpts, _messages [][]byte, _signatures [][]byte, _symbols []string, _borrowers []common.Address, _cTokens []common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "liquidateSNWithPrice", _messages, _signatures, _symbols, _borrowers, _cTokens)
}

// LiquidateSNWithPrice is a paid mutator transaction binding the contract method 0x52fae007.
//
// Solidity: function liquidateSNWithPrice(bytes[] _messages, bytes[] _signatures, string[] _symbols, address[] _borrowers, address[] _cTokens) returns()
func (_Liquidtor *LiquidtorSession) LiquidateSNWithPrice(_messages [][]byte, _signatures [][]byte, _symbols []string, _borrowers []common.Address, _cTokens []common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateSNWithPrice(&_Liquidtor.TransactOpts, _messages, _signatures, _symbols, _borrowers, _cTokens)
}

// LiquidateSNWithPrice is a paid mutator transaction binding the contract method 0x52fae007.
//
// Solidity: function liquidateSNWithPrice(bytes[] _messages, bytes[] _signatures, string[] _symbols, address[] _borrowers, address[] _cTokens) returns()
func (_Liquidtor *LiquidtorTransactorSession) LiquidateSNWithPrice(_messages [][]byte, _signatures [][]byte, _symbols []string, _borrowers []common.Address, _cTokens []common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateSNWithPrice(&_Liquidtor.TransactOpts, _messages, _signatures, _symbols, _borrowers, _cTokens)
}

// LiquidateSWithPrice is a paid mutator transaction binding the contract method 0xf263f3d0.
//
// Solidity: function liquidateSWithPrice(bytes[] _messages, bytes[] _signatures, string[] _symbols, address _borrower, address _repayCToken, address _seizeCToken) returns()
func (_Liquidtor *LiquidtorTransactor) LiquidateSWithPrice(opts *bind.TransactOpts, _messages [][]byte, _signatures [][]byte, _symbols []string, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "liquidateSWithPrice", _messages, _signatures, _symbols, _borrower, _repayCToken, _seizeCToken)
}

// LiquidateSWithPrice is a paid mutator transaction binding the contract method 0xf263f3d0.
//
// Solidity: function liquidateSWithPrice(bytes[] _messages, bytes[] _signatures, string[] _symbols, address _borrower, address _repayCToken, address _seizeCToken) returns()
func (_Liquidtor *LiquidtorSession) LiquidateSWithPrice(_messages [][]byte, _signatures [][]byte, _symbols []string, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateSWithPrice(&_Liquidtor.TransactOpts, _messages, _signatures, _symbols, _borrower, _repayCToken, _seizeCToken)
}

// LiquidateSWithPrice is a paid mutator transaction binding the contract method 0xf263f3d0.
//
// Solidity: function liquidateSWithPrice(bytes[] _messages, bytes[] _signatures, string[] _symbols, address _borrower, address _repayCToken, address _seizeCToken) returns()
func (_Liquidtor *LiquidtorTransactorSession) LiquidateSWithPrice(_messages [][]byte, _signatures [][]byte, _symbols []string, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateSWithPrice(&_Liquidtor.TransactOpts, _messages, _signatures, _symbols, _borrower, _repayCToken, _seizeCToken)
}

// LiquidateTest is a paid mutator transaction binding the contract method 0xadc40ce9.
//
// Solidity: function liquidateTest(address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorTransactor) LiquidateTest(opts *bind.TransactOpts, _borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "liquidateTest", _borrower, _repayCToken, _seizeCToken, _amount)
}

// LiquidateTest is a paid mutator transaction binding the contract method 0xadc40ce9.
//
// Solidity: function liquidateTest(address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorSession) LiquidateTest(_borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateTest(&_Liquidtor.TransactOpts, _borrower, _repayCToken, _seizeCToken, _amount)
}

// LiquidateTest is a paid mutator transaction binding the contract method 0xadc40ce9.
//
// Solidity: function liquidateTest(address _borrower, address _repayCToken, address _seizeCToken, uint256 _amount) returns()
func (_Liquidtor *LiquidtorTransactorSession) LiquidateTest(_borrower common.Address, _repayCToken common.Address, _seizeCToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.LiquidateTest(&_Liquidtor.TransactOpts, _borrower, _repayCToken, _seizeCToken, _amount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptrollerAddress) returns()
func (_Liquidtor *LiquidtorTransactor) SetComptroller(opts *bind.TransactOpts, _comptrollerAddress common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "setComptroller", _comptrollerAddress)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptrollerAddress) returns()
func (_Liquidtor *LiquidtorSession) SetComptroller(_comptrollerAddress common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.SetComptroller(&_Liquidtor.TransactOpts, _comptrollerAddress)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptrollerAddress) returns()
func (_Liquidtor *LiquidtorTransactorSession) SetComptroller(_comptrollerAddress common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.SetComptroller(&_Liquidtor.TransactOpts, _comptrollerAddress)
}

// SetGasThreshold is a paid mutator transaction binding the contract method 0xcb3b3ab3.
//
// Solidity: function setGasThreshold(uint256 _gasThreshold) returns()
func (_Liquidtor *LiquidtorTransactor) SetGasThreshold(opts *bind.TransactOpts, _gasThreshold *big.Int) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "setGasThreshold", _gasThreshold)
}

// SetGasThreshold is a paid mutator transaction binding the contract method 0xcb3b3ab3.
//
// Solidity: function setGasThreshold(uint256 _gasThreshold) returns()
func (_Liquidtor *LiquidtorSession) SetGasThreshold(_gasThreshold *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.SetGasThreshold(&_Liquidtor.TransactOpts, _gasThreshold)
}

// SetGasThreshold is a paid mutator transaction binding the contract method 0xcb3b3ab3.
//
// Solidity: function setGasThreshold(uint256 _gasThreshold) returns()
func (_Liquidtor *LiquidtorTransactorSession) SetGasThreshold(_gasThreshold *big.Int) (*types.Transaction, error) {
	return _Liquidtor.Contract.SetGasThreshold(&_Liquidtor.TransactOpts, _gasThreshold)
}

// SetParam is a paid mutator transaction binding the contract method 0x93577832.
//
// Solidity: function setParam(address cETH, address wETH, address router, address factory) returns()
func (_Liquidtor *LiquidtorTransactor) SetParam(opts *bind.TransactOpts, cETH common.Address, wETH common.Address, router common.Address, factory common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "setParam", cETH, wETH, router, factory)
}

// SetParam is a paid mutator transaction binding the contract method 0x93577832.
//
// Solidity: function setParam(address cETH, address wETH, address router, address factory) returns()
func (_Liquidtor *LiquidtorSession) SetParam(cETH common.Address, wETH common.Address, router common.Address, factory common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.SetParam(&_Liquidtor.TransactOpts, cETH, wETH, router, factory)
}

// SetParam is a paid mutator transaction binding the contract method 0x93577832.
//
// Solidity: function setParam(address cETH, address wETH, address router, address factory) returns()
func (_Liquidtor *LiquidtorTransactorSession) SetParam(cETH common.Address, wETH common.Address, router common.Address, factory common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.SetParam(&_Liquidtor.TransactOpts, cETH, wETH, router, factory)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Liquidtor *LiquidtorTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Liquidtor *LiquidtorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquidtor.Contract.UniswapV2Call(&_Liquidtor.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Liquidtor *LiquidtorTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquidtor.Contract.UniswapV2Call(&_Liquidtor.TransactOpts, sender, amount0, amount1, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_Liquidtor *LiquidtorTransactor) Withdraw(opts *bind.TransactOpts, _assetAddress common.Address) (*types.Transaction, error) {
	return _Liquidtor.contract.Transact(opts, "withdraw", _assetAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_Liquidtor *LiquidtorSession) Withdraw(_assetAddress common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.Withdraw(&_Liquidtor.TransactOpts, _assetAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _assetAddress) returns()
func (_Liquidtor *LiquidtorTransactorSession) Withdraw(_assetAddress common.Address) (*types.Transaction, error) {
	return _Liquidtor.Contract.Withdraw(&_Liquidtor.TransactOpts, _assetAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquidtor *LiquidtorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidtor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquidtor *LiquidtorSession) Receive() (*types.Transaction, error) {
	return _Liquidtor.Contract.Receive(&_Liquidtor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquidtor *LiquidtorTransactorSession) Receive() (*types.Transaction, error) {
	return _Liquidtor.Contract.Receive(&_Liquidtor.TransactOpts)
}

// LiquidtorLogAddrIterator is returned from FilterLogAddr and is used to iterate over the raw logs and unpacked data for LogAddr events raised by the Liquidtor contract.
type LiquidtorLogAddrIterator struct {
	Event *LiquidtorLogAddr // Event containing the contract specifics and raw log

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
func (it *LiquidtorLogAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidtorLogAddr)
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
		it.Event = new(LiquidtorLogAddr)
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
func (it *LiquidtorLogAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidtorLogAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidtorLogAddr represents a LogAddr event raised by the Liquidtor contract.
type LiquidtorLogAddr struct {
	Flag string
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddr is a free log retrieval operation binding the contract event 0x6d82518c519b6cd311f41c64b27912b63453c12888e8c04d4f5b1e78703c2d22.
//
// Solidity: event LogAddr(string flag, address addr)
func (_Liquidtor *LiquidtorFilterer) FilterLogAddr(opts *bind.FilterOpts) (*LiquidtorLogAddrIterator, error) {

	logs, sub, err := _Liquidtor.contract.FilterLogs(opts, "LogAddr")
	if err != nil {
		return nil, err
	}
	return &LiquidtorLogAddrIterator{contract: _Liquidtor.contract, event: "LogAddr", logs: logs, sub: sub}, nil
}

// WatchLogAddr is a free log subscription operation binding the contract event 0x6d82518c519b6cd311f41c64b27912b63453c12888e8c04d4f5b1e78703c2d22.
//
// Solidity: event LogAddr(string flag, address addr)
func (_Liquidtor *LiquidtorFilterer) WatchLogAddr(opts *bind.WatchOpts, sink chan<- *LiquidtorLogAddr) (event.Subscription, error) {

	logs, sub, err := _Liquidtor.contract.WatchLogs(opts, "LogAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidtorLogAddr)
				if err := _Liquidtor.contract.UnpackLog(event, "LogAddr", log); err != nil {
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

// ParseLogAddr is a log parse operation binding the contract event 0x6d82518c519b6cd311f41c64b27912b63453c12888e8c04d4f5b1e78703c2d22.
//
// Solidity: event LogAddr(string flag, address addr)
func (_Liquidtor *LiquidtorFilterer) ParseLogAddr(log types.Log) (*LiquidtorLogAddr, error) {
	event := new(LiquidtorLogAddr)
	if err := _Liquidtor.contract.UnpackLog(event, "LogAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidtorLogLiquiteIterator is returned from FilterLogLiquite and is used to iterate over the raw logs and unpacked data for LogLiquite events raised by the Liquidtor contract.
type LiquidtorLogLiquiteIterator struct {
	Event *LiquidtorLogLiquite // Event containing the contract specifics and raw log

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
func (it *LiquidtorLogLiquiteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidtorLogLiquite)
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
		it.Event = new(LiquidtorLogLiquite)
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
func (it *LiquidtorLogLiquiteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidtorLogLiquiteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidtorLogLiquite represents a LogLiquite event raised by the Liquidtor contract.
type LiquidtorLogLiquite struct {
	Flag string
	Mgs  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogLiquite is a free log retrieval operation binding the contract event 0x367728c7d658cab3ffe2d8ac49fe468c32ec4f4605478e8cfb069e784b4805cb.
//
// Solidity: event LogLiquite(string flag, string mgs)
func (_Liquidtor *LiquidtorFilterer) FilterLogLiquite(opts *bind.FilterOpts) (*LiquidtorLogLiquiteIterator, error) {

	logs, sub, err := _Liquidtor.contract.FilterLogs(opts, "LogLiquite")
	if err != nil {
		return nil, err
	}
	return &LiquidtorLogLiquiteIterator{contract: _Liquidtor.contract, event: "LogLiquite", logs: logs, sub: sub}, nil
}

// WatchLogLiquite is a free log subscription operation binding the contract event 0x367728c7d658cab3ffe2d8ac49fe468c32ec4f4605478e8cfb069e784b4805cb.
//
// Solidity: event LogLiquite(string flag, string mgs)
func (_Liquidtor *LiquidtorFilterer) WatchLogLiquite(opts *bind.WatchOpts, sink chan<- *LiquidtorLogLiquite) (event.Subscription, error) {

	logs, sub, err := _Liquidtor.contract.WatchLogs(opts, "LogLiquite")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidtorLogLiquite)
				if err := _Liquidtor.contract.UnpackLog(event, "LogLiquite", log); err != nil {
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

// ParseLogLiquite is a log parse operation binding the contract event 0x367728c7d658cab3ffe2d8ac49fe468c32ec4f4605478e8cfb069e784b4805cb.
//
// Solidity: event LogLiquite(string flag, string mgs)
func (_Liquidtor *LiquidtorFilterer) ParseLogLiquite(log types.Log) (*LiquidtorLogLiquite, error) {
	event := new(LiquidtorLogLiquite)
	if err := _Liquidtor.contract.UnpackLog(event, "LogLiquite", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidtorRecipientChangedIterator is returned from FilterRecipientChanged and is used to iterate over the raw logs and unpacked data for RecipientChanged events raised by the Liquidtor contract.
type LiquidtorRecipientChangedIterator struct {
	Event *LiquidtorRecipientChanged // Event containing the contract specifics and raw log

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
func (it *LiquidtorRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidtorRecipientChanged)
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
		it.Event = new(LiquidtorRecipientChanged)
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
func (it *LiquidtorRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidtorRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidtorRecipientChanged represents a RecipientChanged event raised by the Liquidtor contract.
type LiquidtorRecipientChanged struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecipientChanged is a free log retrieval operation binding the contract event 0xff2d07bd188a9eb41acbc4a7db39e18956c95ab7f54f434d97849bf6206e577c.
//
// Solidity: event RecipientChanged(address recipient)
func (_Liquidtor *LiquidtorFilterer) FilterRecipientChanged(opts *bind.FilterOpts) (*LiquidtorRecipientChangedIterator, error) {

	logs, sub, err := _Liquidtor.contract.FilterLogs(opts, "RecipientChanged")
	if err != nil {
		return nil, err
	}
	return &LiquidtorRecipientChangedIterator{contract: _Liquidtor.contract, event: "RecipientChanged", logs: logs, sub: sub}, nil
}

// WatchRecipientChanged is a free log subscription operation binding the contract event 0xff2d07bd188a9eb41acbc4a7db39e18956c95ab7f54f434d97849bf6206e577c.
//
// Solidity: event RecipientChanged(address recipient)
func (_Liquidtor *LiquidtorFilterer) WatchRecipientChanged(opts *bind.WatchOpts, sink chan<- *LiquidtorRecipientChanged) (event.Subscription, error) {

	logs, sub, err := _Liquidtor.contract.WatchLogs(opts, "RecipientChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidtorRecipientChanged)
				if err := _Liquidtor.contract.UnpackLog(event, "RecipientChanged", log); err != nil {
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

// ParseRecipientChanged is a log parse operation binding the contract event 0xff2d07bd188a9eb41acbc4a7db39e18956c95ab7f54f434d97849bf6206e577c.
//
// Solidity: event RecipientChanged(address recipient)
func (_Liquidtor *LiquidtorFilterer) ParseRecipientChanged(log types.Log) (*LiquidtorRecipientChanged, error) {
	event := new(LiquidtorRecipientChanged)
	if err := _Liquidtor.contract.UnpackLog(event, "RecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidtorRevenueWithdrawnIterator is returned from FilterRevenueWithdrawn and is used to iterate over the raw logs and unpacked data for RevenueWithdrawn events raised by the Liquidtor contract.
type LiquidtorRevenueWithdrawnIterator struct {
	Event *LiquidtorRevenueWithdrawn // Event containing the contract specifics and raw log

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
func (it *LiquidtorRevenueWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidtorRevenueWithdrawn)
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
		it.Event = new(LiquidtorRevenueWithdrawn)
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
func (it *LiquidtorRevenueWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidtorRevenueWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidtorRevenueWithdrawn represents a RevenueWithdrawn event raised by the Liquidtor contract.
type LiquidtorRevenueWithdrawn struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevenueWithdrawn is a free log retrieval operation binding the contract event 0x6c198f0e3e4f68668dfd7b1a689ee110197d822dd8ec93a685138d4b8537ae06.
//
// Solidity: event RevenueWithdrawn(address recipient, address token, uint256 amount)
func (_Liquidtor *LiquidtorFilterer) FilterRevenueWithdrawn(opts *bind.FilterOpts) (*LiquidtorRevenueWithdrawnIterator, error) {

	logs, sub, err := _Liquidtor.contract.FilterLogs(opts, "RevenueWithdrawn")
	if err != nil {
		return nil, err
	}
	return &LiquidtorRevenueWithdrawnIterator{contract: _Liquidtor.contract, event: "RevenueWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRevenueWithdrawn is a free log subscription operation binding the contract event 0x6c198f0e3e4f68668dfd7b1a689ee110197d822dd8ec93a685138d4b8537ae06.
//
// Solidity: event RevenueWithdrawn(address recipient, address token, uint256 amount)
func (_Liquidtor *LiquidtorFilterer) WatchRevenueWithdrawn(opts *bind.WatchOpts, sink chan<- *LiquidtorRevenueWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Liquidtor.contract.WatchLogs(opts, "RevenueWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidtorRevenueWithdrawn)
				if err := _Liquidtor.contract.UnpackLog(event, "RevenueWithdrawn", log); err != nil {
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

// ParseRevenueWithdrawn is a log parse operation binding the contract event 0x6c198f0e3e4f68668dfd7b1a689ee110197d822dd8ec93a685138d4b8537ae06.
//
// Solidity: event RevenueWithdrawn(address recipient, address token, uint256 amount)
func (_Liquidtor *LiquidtorFilterer) ParseRevenueWithdrawn(log types.Log) (*LiquidtorRevenueWithdrawn, error) {
	event := new(LiquidtorRevenueWithdrawn)
	if err := _Liquidtor.contract.UnpackLog(event, "RevenueWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
