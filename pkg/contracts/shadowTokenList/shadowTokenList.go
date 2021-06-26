// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shadowTokenList

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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BasicTokenFuncSigs maps the 4-byte function signature to its string representation.
var BasicTokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
var BasicTokenBin = "0x608060405234801561001057600080fd5b50610220806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806318160ddd1461004657806370a0823114610060578063a9059cbb14610086575b600080fd5b61004e6100c6565b60408051918252519081900360200190f35b61004e6004803603602081101561007657600080fd5b50356001600160a01b03166100cc565b6100b26004803603604081101561009c57600080fd5b506001600160a01b0381351690602001356100e7565b604080519115158252519081900360200190f35b60015490565b6001600160a01b031660009081526020819052604090205490565b60006001600160a01b0383166100fc57600080fd5b3360009081526020819052604090205482111561011857600080fd5b33600090815260208190526040902054610138908363ffffffff6101c616565b33600090815260208190526040808220929092556001600160a01b0385168152205461016a908363ffffffff6101d816565b6001600160a01b038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6000828211156101d257fe5b50900390565b818101828110156101e557fe5b9291505056fea265627a7a7231582006b84ed7f67dc48835992392beca294187a893121b7fb94d73dacd2c107f1aff64736f6c63430005100032"

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasicToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasicToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) ParseTransfer(log types.Log) (*BasicTokenTransfer, error) {
	event := new(BasicTokenTransfer)
	if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BasicFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BasicFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) ParseTransfer(log types.Log) (*ERC20BasicTransfer, error) {
	event := new(ERC20BasicTransfer)
	if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenABI is the input ABI used to generate the binding from.
const MintableTokenABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MintableTokenFuncSigs maps the 4-byte function signature to its string representation.
var MintableTokenFuncSigs = map[string]string{
	"42966c68": "burn(uint256)",
	"40c10f19": "mint(address,uint256)",
}

// MintableToken is an auto generated Go binding around an Ethereum contract.
type MintableToken struct {
	MintableTokenCaller     // Read-only binding to the contract
	MintableTokenTransactor // Write-only binding to the contract
	MintableTokenFilterer   // Log filterer for contract events
}

// MintableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintableTokenSession struct {
	Contract     *MintableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintableTokenCallerSession struct {
	Contract *MintableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MintableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintableTokenTransactorSession struct {
	Contract     *MintableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MintableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintableTokenRaw struct {
	Contract *MintableToken // Generic contract binding to access the raw methods on
}

// MintableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintableTokenCallerRaw struct {
	Contract *MintableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MintableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintableTokenTransactorRaw struct {
	Contract *MintableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintableToken creates a new instance of MintableToken, bound to a specific deployed contract.
func NewMintableToken(address common.Address, backend bind.ContractBackend) (*MintableToken, error) {
	contract, err := bindMintableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}, MintableTokenFilterer: MintableTokenFilterer{contract: contract}}, nil
}

// NewMintableTokenCaller creates a new read-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenCaller(address common.Address, caller bind.ContractCaller) (*MintableTokenCaller, error) {
	contract, err := bindMintableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenCaller{contract: contract}, nil
}

// NewMintableTokenTransactor creates a new write-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MintableTokenTransactor, error) {
	contract, err := bindMintableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransactor{contract: contract}, nil
}

// NewMintableTokenFilterer creates a new log filterer instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MintableTokenFilterer, error) {
	contract, err := bindMintableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintableTokenFilterer{contract: contract}, nil
}

// bindMintableToken binds a generic wrapper to an already deployed contract.
func bindMintableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.MintableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ) returns(bool)
func (_MintableToken *MintableTokenTransactor) Burn(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "burn", arg0)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ) returns(bool)
func (_MintableToken *MintableTokenSession) Burn(arg0 *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Burn(&_MintableToken.TransactOpts, arg0)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Burn(arg0 *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Burn(&_MintableToken.TransactOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns(bool)
func (_MintableToken *MintableTokenTransactor) Mint(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "mint", arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns(bool)
func (_MintableToken *MintableTokenSession) Mint(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Mint(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, arg0, arg1)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610150806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b1461005f575b600080fd5b610043610087565b604080516001600160a01b039092168252519081900360200190f35b6100856004803603602081101561007557600080fd5b50356001600160a01b0316610096565b005b6000546001600160a01b031681565b6000546001600160a01b031633146100ad57600080fd5b6001600160a01b0381166100c057600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820cfac7fdae03aaf4a89685721765e67184869b60564afbeb329e84ccc2f8ddda364736f6c63430005100032"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

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
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// PausableBin is the compiled bytecode used for deploying new contracts.
var PausableBin = "0x6080604052600080546001600160a81b0319163317905561027a806100256000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633f4ba83a1461005c5780635c975abb146100665780638456cb59146100825780638da5cb5b1461008a578063f2fde38b146100ae575b600080fd5b6100646100d4565b005b61006e610137565b604080519115158252519081900360200190f35b610064610147565b6100926101b1565b604080516001600160a01b039092168252519081900360200190f35b610064600480360360208110156100c457600080fd5b50356001600160a01b03166101c0565b6000546001600160a01b031633146100eb57600080fd5b600054600160a01b900460ff1661010157600080fd5b6000805460ff60a01b191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b600054600160a01b900460ff1681565b6000546001600160a01b0316331461015e57600080fd5b600054600160a01b900460ff161561017557600080fd5b6000805460ff60a01b1916600160a01b1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b6000546001600160a01b031681565b6000546001600160a01b031633146101d757600080fd5b6001600160a01b0381166101ea57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a7231582000c5445cc2043ce07bc289e90e76ee92eec886925526a2e82eda58b2a9cff69064736f6c63430005100032"

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pausable *PausableFilterer) ParseOwnershipTransferred(log types.Log) (*PausableOwnershipTransferred, error) {
	event := new(PausableOwnershipTransferred)
	if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) ParsePause(log types.Log) (*PausablePause, error) {
	event := new(PausablePause)
	if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) ParseUnpause(log types.Log) (*PausableUnpause, error) {
	event := new(PausableUnpause)
	if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582068ae7531a7892b41a74930c03accd39619c8a70f2337549b7c797879b7e250c464736f6c63430005100032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ShadowTokenABI is the input ABI used to generate the binding from.
const ShadowTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_coToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"coToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ShadowTokenFuncSigs maps the 4-byte function signature to its string representation.
var ShadowTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"7f9864f7": "coToken()",
	"313ce567": "decimals()",
	"66188463": "decreaseApproval(address,uint256)",
	"d73dd623": "increaseApproval(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"07546172": "minter()",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// ShadowTokenBin is the compiled bytecode used for deploying new contracts.
var ShadowTokenBin = "0x60806040526003805460ff60a01b191690553480156200001e57600080fd5b506040516200108138038062001081833981810160405260a08110156200004457600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200007057600080fd5b9083019060208201858111156200008657600080fd5b8251640100000000811182820188101715620000a157600080fd5b82525081516020918201929091019080838360005b83811015620000d0578181015183820152602001620000b6565b50505050905090810190601f168015620000fe5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200012257600080fd5b9083019060208201858111156200013857600080fd5b82516401000000008111828201881017156200015357600080fd5b82525081516020918201929091019080838360005b838110156200018257818101518382015260200162000168565b50505050905090810190601f168015620001b05780820380516001836020036101000a031916815260200191505b5060405260209081015160038054336001600160a01b0319918216179091556005805482166001600160a01b038b8116919091179091556004805490921690891617905585519093506200020b925060069186019062000270565b5081516200022190600790602085019062000270565b506008805460ff191660ff83161790556040516001600160a01b038616907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a2505050505062000315565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002b357805160ff1916838001178555620002e3565b82800160010185558215620002e3579182015b82811115620002e3578251825591602001919060010190620002c6565b50620002f1929150620002f5565b5090565b6200031291905b80821115620002f15760008155600101620002fc565b90565b610d5c80620003256000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806366188463116100ad57806395d89b411161007157806395d89b4114610345578063a9059cbb1461034d578063d73dd62314610379578063dd62ed3e146103a5578063f2fde38b146103d35761012c565b806366188463146102db57806370a08231146103075780637f9864f71461032d5780638456cb59146103355780638da5cb5b1461033d5761012c565b8063313ce567116100f4578063313ce567146102625780633f4ba83a1461028057806340c10f191461028a57806342966c68146102b65780635c975abb146102d35761012c565b806306fdde031461013157806307546172146101ae578063095ea7b3146101d257806318160ddd1461021257806323b872dd1461022c575b600080fd5b6101396103f9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017357818101518382015260200161015b565b50505050905090810190601f1680156101a05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101b6610487565b604080516001600160a01b039092168252519081900360200190f35b6101fe600480360360408110156101e857600080fd5b506001600160a01b038135169060200135610496565b604080519115158252519081900360200190f35b61021a6104fc565b60408051918252519081900360200190f35b6101fe6004803603606081101561024257600080fd5b506001600160a01b03813581169160208101359091169060400135610502565b61026a610665565b6040805160ff9092168252519081900360200190f35b61028861066e565b005b6101fe600480360360408110156102a057600080fd5b506001600160a01b0381351690602001356106d3565b6101fe600480360360208110156102cc57600080fd5b5035610804565b6101fe6108cb565b6101fe600480360360408110156102f157600080fd5b506001600160a01b0381351690602001356108db565b61021a6004803603602081101561031d57600080fd5b50356001600160a01b03166109cb565b6101b66109e6565b6102886109f5565b6101b6610a61565b610139610a70565b6101fe6004803603604081101561036357600080fd5b506001600160a01b038135169060200135610acb565b6101fe6004803603604081101561038f57600080fd5b506001600160a01b038135169060200135610b98565b61021a600480360360408110156103bb57600080fd5b506001600160a01b0381358116916020013516610c31565b610288600480360360208110156103e957600080fd5b50356001600160a01b0316610c5c565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561047f5780601f106104545761010080835404028352916020019161047f565b820191906000526020600020905b81548152906001019060200180831161046257829003601f168201915b505050505081565b6005546001600160a01b031681565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b60006001600160a01b03831661051757600080fd5b6001600160a01b03841660009081526020819052604090205482111561053c57600080fd5b6001600160a01b038416600090815260026020908152604080832033845290915290205482111561056c57600080fd5b6001600160a01b038416600090815260208190526040902054610595908363ffffffff610ce216565b6001600160a01b0380861660009081526020819052604080822093909355908516815220546105ca908363ffffffff610cf416565b6001600160a01b0380851660009081526020818152604080832094909455918716815260028252828120338252909152205461060c908363ffffffff610ce216565b6001600160a01b0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610d08833981519152929181900390910190a35060019392505050565b60085460ff1681565b6003546001600160a01b0316331461068557600080fd5b600354600160a01b900460ff1661069b57600080fd5b6003805460ff60a01b191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b6005546000906001600160a01b03163314610726576040805162461bcd60e51b815260206004820152600e60248201526d3737ba103a34329036b4b73a32b960911b604482015290519081900360640190fd5b600354600160a01b900460ff161561073d57600080fd5b600154610750908363ffffffff610cf416565b6001556001600160a01b03831660009081526020819052604090205461077c908363ffffffff610cf416565b6001600160a01b03841660008181526020818152604091829020939093558051858152905191927f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe92918290030190a26040805183815290516001600160a01b03851691600091600080516020610d088339815191529181900360200190a350600192915050565b3360009081526020819052604081205482111561082057600080fd5b600154610833908363ffffffff610ce216565b60015533600090815260208190526040902054610856908363ffffffff610ce216565b3360008181526020818152604091829020939093558051858152905191927f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df792918290030190a26040805183815290516000913391600080516020610d088339815191529181900360200190a3506001919050565b600354600160a01b900460ff1681565b3360009081526002602090815260408083206001600160a01b038616845290915281205480831115610930573360009081526002602090815260408083206001600160a01b0388168452909152812055610965565b610940818463ffffffff610ce216565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526020819052604090205490565b6004546001600160a01b031681565b6003546001600160a01b03163314610a0c57600080fd5b600354600160a01b900460ff1615610a2357600080fd5b6003805460ff60a01b1916600160a01b1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b6003546001600160a01b031681565b6007805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561047f5780601f106104545761010080835404028352916020019161047f565b60006001600160a01b038316610ae057600080fd5b33600090815260208190526040902054821115610afc57600080fd5b33600090815260208190526040902054610b1c908363ffffffff610ce216565b33600090815260208190526040808220929092556001600160a01b03851681522054610b4e908363ffffffff610cf416565b6001600160a01b03841660008181526020818152604091829020939093558051858152905191923392600080516020610d088339815191529281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610bcc908363ffffffff610cf416565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6003546001600160a01b03163314610c7357600080fd5b6001600160a01b038116610c8657600080fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b600082821115610cee57fe5b50900390565b81810182811015610d0157fe5b9291505056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa265627a7a723158208f534a7473ce0fe3184de99e306787634a9c6b6e689f3f05a8b332349fe0962964736f6c63430005100032"

// DeployShadowToken deploys a new Ethereum contract, binding an instance of ShadowToken to it.
func DeployShadowToken(auth *bind.TransactOpts, backend bind.ContractBackend, _minter common.Address, _coToken common.Address, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *ShadowToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ShadowTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ShadowTokenBin), backend, _minter, _coToken, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShadowToken{ShadowTokenCaller: ShadowTokenCaller{contract: contract}, ShadowTokenTransactor: ShadowTokenTransactor{contract: contract}, ShadowTokenFilterer: ShadowTokenFilterer{contract: contract}}, nil
}

// ShadowToken is an auto generated Go binding around an Ethereum contract.
type ShadowToken struct {
	ShadowTokenCaller     // Read-only binding to the contract
	ShadowTokenTransactor // Write-only binding to the contract
	ShadowTokenFilterer   // Log filterer for contract events
}

// ShadowTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShadowTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShadowTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShadowTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShadowTokenSession struct {
	Contract     *ShadowToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShadowTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShadowTokenCallerSession struct {
	Contract *ShadowTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ShadowTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShadowTokenTransactorSession struct {
	Contract     *ShadowTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ShadowTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShadowTokenRaw struct {
	Contract *ShadowToken // Generic contract binding to access the raw methods on
}

// ShadowTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShadowTokenCallerRaw struct {
	Contract *ShadowTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ShadowTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShadowTokenTransactorRaw struct {
	Contract *ShadowTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShadowToken creates a new instance of ShadowToken, bound to a specific deployed contract.
func NewShadowToken(address common.Address, backend bind.ContractBackend) (*ShadowToken, error) {
	contract, err := bindShadowToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShadowToken{ShadowTokenCaller: ShadowTokenCaller{contract: contract}, ShadowTokenTransactor: ShadowTokenTransactor{contract: contract}, ShadowTokenFilterer: ShadowTokenFilterer{contract: contract}}, nil
}

// NewShadowTokenCaller creates a new read-only instance of ShadowToken, bound to a specific deployed contract.
func NewShadowTokenCaller(address common.Address, caller bind.ContractCaller) (*ShadowTokenCaller, error) {
	contract, err := bindShadowToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenCaller{contract: contract}, nil
}

// NewShadowTokenTransactor creates a new write-only instance of ShadowToken, bound to a specific deployed contract.
func NewShadowTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ShadowTokenTransactor, error) {
	contract, err := bindShadowToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenTransactor{contract: contract}, nil
}

// NewShadowTokenFilterer creates a new log filterer instance of ShadowToken, bound to a specific deployed contract.
func NewShadowTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ShadowTokenFilterer, error) {
	contract, err := bindShadowToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenFilterer{contract: contract}, nil
}

// bindShadowToken binds a generic wrapper to an already deployed contract.
func bindShadowToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShadowTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShadowToken *ShadowTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShadowToken.Contract.ShadowTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShadowToken *ShadowTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowToken.Contract.ShadowTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShadowToken *ShadowTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShadowToken.Contract.ShadowTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShadowToken *ShadowTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShadowToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShadowToken *ShadowTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShadowToken *ShadowTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShadowToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ShadowToken *ShadowTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ShadowToken *ShadowTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ShadowToken.Contract.Allowance(&_ShadowToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ShadowToken *ShadowTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ShadowToken.Contract.Allowance(&_ShadowToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_ShadowToken *ShadowTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_ShadowToken *ShadowTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ShadowToken.Contract.BalanceOf(&_ShadowToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_ShadowToken *ShadowTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ShadowToken.Contract.BalanceOf(&_ShadowToken.CallOpts, _owner)
}

// CoToken is a free data retrieval call binding the contract method 0x7f9864f7.
//
// Solidity: function coToken() view returns(address)
func (_ShadowToken *ShadowTokenCaller) CoToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "coToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CoToken is a free data retrieval call binding the contract method 0x7f9864f7.
//
// Solidity: function coToken() view returns(address)
func (_ShadowToken *ShadowTokenSession) CoToken() (common.Address, error) {
	return _ShadowToken.Contract.CoToken(&_ShadowToken.CallOpts)
}

// CoToken is a free data retrieval call binding the contract method 0x7f9864f7.
//
// Solidity: function coToken() view returns(address)
func (_ShadowToken *ShadowTokenCallerSession) CoToken() (common.Address, error) {
	return _ShadowToken.Contract.CoToken(&_ShadowToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShadowToken *ShadowTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShadowToken *ShadowTokenSession) Decimals() (uint8, error) {
	return _ShadowToken.Contract.Decimals(&_ShadowToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShadowToken *ShadowTokenCallerSession) Decimals() (uint8, error) {
	return _ShadowToken.Contract.Decimals(&_ShadowToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ShadowToken *ShadowTokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ShadowToken *ShadowTokenSession) Minter() (common.Address, error) {
	return _ShadowToken.Contract.Minter(&_ShadowToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ShadowToken *ShadowTokenCallerSession) Minter() (common.Address, error) {
	return _ShadowToken.Contract.Minter(&_ShadowToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShadowToken *ShadowTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShadowToken *ShadowTokenSession) Name() (string, error) {
	return _ShadowToken.Contract.Name(&_ShadowToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShadowToken *ShadowTokenCallerSession) Name() (string, error) {
	return _ShadowToken.Contract.Name(&_ShadowToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowToken *ShadowTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowToken *ShadowTokenSession) Owner() (common.Address, error) {
	return _ShadowToken.Contract.Owner(&_ShadowToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowToken *ShadowTokenCallerSession) Owner() (common.Address, error) {
	return _ShadowToken.Contract.Owner(&_ShadowToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ShadowToken *ShadowTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ShadowToken *ShadowTokenSession) Paused() (bool, error) {
	return _ShadowToken.Contract.Paused(&_ShadowToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ShadowToken *ShadowTokenCallerSession) Paused() (bool, error) {
	return _ShadowToken.Contract.Paused(&_ShadowToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShadowToken *ShadowTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShadowToken *ShadowTokenSession) Symbol() (string, error) {
	return _ShadowToken.Contract.Symbol(&_ShadowToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShadowToken *ShadowTokenCallerSession) Symbol() (string, error) {
	return _ShadowToken.Contract.Symbol(&_ShadowToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShadowToken *ShadowTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShadowToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShadowToken *ShadowTokenSession) TotalSupply() (*big.Int, error) {
	return _ShadowToken.Contract.TotalSupply(&_ShadowToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShadowToken *ShadowTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ShadowToken.Contract.TotalSupply(&_ShadowToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Approve(&_ShadowToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Approve(&_ShadowToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool)
func (_ShadowToken *ShadowTokenTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool)
func (_ShadowToken *ShadowTokenSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Burn(&_ShadowToken.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool)
func (_ShadowToken *ShadowTokenTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Burn(&_ShadowToken.TransactOpts, _amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_ShadowToken *ShadowTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_ShadowToken *ShadowTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.DecreaseApproval(&_ShadowToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_ShadowToken *ShadowTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.DecreaseApproval(&_ShadowToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_ShadowToken *ShadowTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_ShadowToken *ShadowTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.IncreaseApproval(&_ShadowToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_ShadowToken *ShadowTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.IncreaseApproval(&_ShadowToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ShadowToken *ShadowTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ShadowToken *ShadowTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Mint(&_ShadowToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ShadowToken *ShadowTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Mint(&_ShadowToken.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ShadowToken *ShadowTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ShadowToken *ShadowTokenSession) Pause() (*types.Transaction, error) {
	return _ShadowToken.Contract.Pause(&_ShadowToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ShadowToken *ShadowTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _ShadowToken.Contract.Pause(&_ShadowToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Transfer(&_ShadowToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.Transfer(&_ShadowToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.TransferFrom(&_ShadowToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ShadowToken *ShadowTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ShadowToken.Contract.TransferFrom(&_ShadowToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShadowToken *ShadowTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShadowToken *ShadowTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShadowToken.Contract.TransferOwnership(&_ShadowToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShadowToken *ShadowTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShadowToken.Contract.TransferOwnership(&_ShadowToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ShadowToken *ShadowTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ShadowToken *ShadowTokenSession) Unpause() (*types.Transaction, error) {
	return _ShadowToken.Contract.Unpause(&_ShadowToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ShadowToken *ShadowTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _ShadowToken.Contract.Unpause(&_ShadowToken.TransactOpts)
}

// ShadowTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ShadowToken contract.
type ShadowTokenApprovalIterator struct {
	Event *ShadowTokenApproval // Event containing the contract specifics and raw log

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
func (it *ShadowTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenApproval)
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
		it.Event = new(ShadowTokenApproval)
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
func (it *ShadowTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenApproval represents a Approval event raised by the ShadowToken contract.
type ShadowTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShadowToken *ShadowTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ShadowTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenApprovalIterator{contract: _ShadowToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShadowToken *ShadowTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ShadowTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenApproval)
				if err := _ShadowToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShadowToken *ShadowTokenFilterer) ParseApproval(log types.Log) (*ShadowTokenApproval, error) {
	event := new(ShadowTokenApproval)
	if err := _ShadowToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the ShadowToken contract.
type ShadowTokenBurnedIterator struct {
	Event *ShadowTokenBurned // Event containing the contract specifics and raw log

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
func (it *ShadowTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenBurned)
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
		it.Event = new(ShadowTokenBurned)
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
func (it *ShadowTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenBurned represents a Burned event raised by the ShadowToken contract.
type ShadowTokenBurned struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_ShadowToken *ShadowTokenFilterer) FilterBurned(opts *bind.FilterOpts, from []common.Address) (*ShadowTokenBurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenBurnedIterator{contract: _ShadowToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_ShadowToken *ShadowTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *ShadowTokenBurned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenBurned)
				if err := _ShadowToken.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 amount)
func (_ShadowToken *ShadowTokenFilterer) ParseBurned(log types.Log) (*ShadowTokenBurned, error) {
	event := new(ShadowTokenBurned)
	if err := _ShadowToken.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the ShadowToken contract.
type ShadowTokenMintedIterator struct {
	Event *ShadowTokenMinted // Event containing the contract specifics and raw log

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
func (it *ShadowTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenMinted)
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
		it.Event = new(ShadowTokenMinted)
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
func (it *ShadowTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenMinted represents a Minted event raised by the ShadowToken contract.
type ShadowTokenMinted struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_ShadowToken *ShadowTokenFilterer) FilterMinted(opts *bind.FilterOpts, to []common.Address) (*ShadowTokenMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenMintedIterator{contract: _ShadowToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_ShadowToken *ShadowTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *ShadowTokenMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenMinted)
				if err := _ShadowToken.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed to, uint256 amount)
func (_ShadowToken *ShadowTokenFilterer) ParseMinted(log types.Log) (*ShadowTokenMinted, error) {
	event := new(ShadowTokenMinted)
	if err := _ShadowToken.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the ShadowToken contract.
type ShadowTokenMinterAddedIterator struct {
	Event *ShadowTokenMinterAdded // Event containing the contract specifics and raw log

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
func (it *ShadowTokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenMinterAdded)
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
		it.Event = new(ShadowTokenMinterAdded)
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
func (it *ShadowTokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenMinterAdded represents a MinterAdded event raised by the ShadowToken contract.
type ShadowTokenMinterAdded struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_ShadowToken *ShadowTokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, minter []common.Address) (*ShadowTokenMinterAddedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenMinterAddedIterator{contract: _ShadowToken.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_ShadowToken *ShadowTokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *ShadowTokenMinterAdded, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenMinterAdded)
				if err := _ShadowToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed minter)
func (_ShadowToken *ShadowTokenFilterer) ParseMinterAdded(log types.Log) (*ShadowTokenMinterAdded, error) {
	event := new(ShadowTokenMinterAdded)
	if err := _ShadowToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ShadowToken contract.
type ShadowTokenMinterRemovedIterator struct {
	Event *ShadowTokenMinterRemoved // Event containing the contract specifics and raw log

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
func (it *ShadowTokenMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenMinterRemoved)
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
		it.Event = new(ShadowTokenMinterRemoved)
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
func (it *ShadowTokenMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenMinterRemoved represents a MinterRemoved event raised by the ShadowToken contract.
type ShadowTokenMinterRemoved struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_ShadowToken *ShadowTokenFilterer) FilterMinterRemoved(opts *bind.FilterOpts, minter []common.Address) (*ShadowTokenMinterRemovedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "MinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenMinterRemovedIterator{contract: _ShadowToken.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_ShadowToken *ShadowTokenFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ShadowTokenMinterRemoved, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "MinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenMinterRemoved)
				if err := _ShadowToken.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed minter)
func (_ShadowToken *ShadowTokenFilterer) ParseMinterRemoved(log types.Log) (*ShadowTokenMinterRemoved, error) {
	event := new(ShadowTokenMinterRemoved)
	if err := _ShadowToken.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShadowToken contract.
type ShadowTokenOwnershipTransferredIterator struct {
	Event *ShadowTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShadowTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenOwnershipTransferred)
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
		it.Event = new(ShadowTokenOwnershipTransferred)
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
func (it *ShadowTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenOwnershipTransferred represents a OwnershipTransferred event raised by the ShadowToken contract.
type ShadowTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShadowToken *ShadowTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShadowTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenOwnershipTransferredIterator{contract: _ShadowToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShadowToken *ShadowTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShadowTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenOwnershipTransferred)
				if err := _ShadowToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShadowToken *ShadowTokenFilterer) ParseOwnershipTransferred(log types.Log) (*ShadowTokenOwnershipTransferred, error) {
	event := new(ShadowTokenOwnershipTransferred)
	if err := _ShadowToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ShadowToken contract.
type ShadowTokenPauseIterator struct {
	Event *ShadowTokenPause // Event containing the contract specifics and raw log

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
func (it *ShadowTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenPause)
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
		it.Event = new(ShadowTokenPause)
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
func (it *ShadowTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenPause represents a Pause event raised by the ShadowToken contract.
type ShadowTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ShadowToken *ShadowTokenFilterer) FilterPause(opts *bind.FilterOpts) (*ShadowTokenPauseIterator, error) {

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ShadowTokenPauseIterator{contract: _ShadowToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ShadowToken *ShadowTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ShadowTokenPause) (event.Subscription, error) {

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenPause)
				if err := _ShadowToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ShadowToken *ShadowTokenFilterer) ParsePause(log types.Log) (*ShadowTokenPause, error) {
	event := new(ShadowTokenPause)
	if err := _ShadowToken.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ShadowToken contract.
type ShadowTokenTransferIterator struct {
	Event *ShadowTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ShadowTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenTransfer)
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
		it.Event = new(ShadowTokenTransfer)
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
func (it *ShadowTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenTransfer represents a Transfer event raised by the ShadowToken contract.
type ShadowTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShadowToken *ShadowTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShadowTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenTransferIterator{contract: _ShadowToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShadowToken *ShadowTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ShadowTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenTransfer)
				if err := _ShadowToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShadowToken *ShadowTokenFilterer) ParseTransfer(log types.Log) (*ShadowTokenTransfer, error) {
	event := new(ShadowTokenTransfer)
	if err := _ShadowToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ShadowToken contract.
type ShadowTokenUnpauseIterator struct {
	Event *ShadowTokenUnpause // Event containing the contract specifics and raw log

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
func (it *ShadowTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenUnpause)
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
		it.Event = new(ShadowTokenUnpause)
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
func (it *ShadowTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenUnpause represents a Unpause event raised by the ShadowToken contract.
type ShadowTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ShadowToken *ShadowTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*ShadowTokenUnpauseIterator, error) {

	logs, sub, err := _ShadowToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ShadowTokenUnpauseIterator{contract: _ShadowToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ShadowToken *ShadowTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ShadowTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _ShadowToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenUnpause)
				if err := _ShadowToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ShadowToken *ShadowTokenFilterer) ParseUnpause(log types.Log) (*ShadowTokenUnpause, error) {
	event := new(ShadowTokenUnpause)
	if err := _ShadowToken.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowTokenListManagerABI is the input ABI used to generate the binding from.
const ShadowTokenListManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_coToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"addShadowToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"}],\"name\":\"setMinAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenList\",\"outputs\":[{\"internalType\":\"contractTokenList\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ShadowTokenListManagerFuncSigs maps the 4-byte function signature to its string representation.
var ShadowTokenListManagerFuncSigs = map[string]string{
	"3f47a17f": "addShadowToken(address,address,string,string,uint8,uint256,uint256)",
	"8da5cb5b": "owner()",
	"5fa7b584": "removeToken(address)",
	"fb52b065": "setMaxAmount(address,uint256)",
	"d5708d5a": "setMinAmount(address,uint256)",
	"9e2c58ca": "tokenList()",
	"f2fde38b": "transferOwnership(address)",
	"0900f010": "upgrade(address)",
}

// ShadowTokenListManagerBin is the compiled bytecode used for deploying new contracts.
var ShadowTokenListManagerBin = "0x608060405234801561001057600080fd5b506040516118ab3803806118ab8339818101604052602081101561003357600080fd5b5051600080546001600160a01b03199081163317909155600180546001600160a01b0390931692909116919091179055611839806100726000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80639e2c58ca1161005b5780639e2c58ca14610267578063d5708d5a1461026f578063f2fde38b1461029b578063fb52b065146102c157610088565b80630900f0101461008d5780633f47a17f146100b55780635fa7b5841461021d5780638da5cb5b14610243575b600080fd5b6100b3600480360360208110156100a357600080fd5b50356001600160a01b03166102ed565b005b610209600480360360e08110156100cb57600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156100ff57600080fd5b82018360208201111561011157600080fd5b8035906020019184600183028401116401000000008311171561013357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561018657600080fd5b82018360208201111561019857600080fd5b803590602001918460018302840111640100000000831117156101ba57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505060ff833516935050506020810135906040013561036d565b604080519115158252519081900360200190f35b6102096004803603602081101561023357600080fd5b50356001600160a01b0316610546565b61024b6105df565b604080516001600160a01b039092168252519081900360200190f35b61024b6105ee565b6100b36004803603604081101561028557600080fd5b506001600160a01b0381351690602001356105fd565b6100b3600480360360208110156102b157600080fd5b50356001600160a01b0316610685565b6100b3600480360360408110156102d757600080fd5b506001600160a01b03813516906020013561070a565b6000546001600160a01b0316331461030457600080fd5b6001546040805163f2fde38b60e01b81526001600160a01b0384811660048301529151919092169163f2fde38b91602480830192600092919082900301818387803b15801561035257600080fd5b505af1158015610366573d6000803e3d6000fd5b5050505050565b600080546001600160a01b0316331461038557600080fd5b6000888888888860405161039890610776565b6001600160a01b038087168252851660208083019190915260ff8316608083015260a06040830181815286519184019190915285519091606084019160c085019188019080838360005b838110156103fa5781810151838201526020016103e2565b50505050905090810190601f1680156104275780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561045a578181015183820152602001610442565b50505050905090810190601f1680156104875780820380516001836020036101000a031916815260200191505b50975050505050505050604051809103906000f0801580156104ad573d6000803e3d6000fd5b506001546040805163dee1f2af60e01b81526001600160a01b0380851660048301526024820189905260448201889052915193945091169163dee1f2af916064808201926020929091908290030181600087803b15801561050d57600080fd5b505af1158015610521573d6000803e3d6000fd5b505050506040513d602081101561053757600080fd5b50519998505050505050505050565b600080546001600160a01b0316331461055e57600080fd5b600154604080516317e9ed6160e21b81526001600160a01b03858116600483015291519190921691635fa7b5849160248083019260209291908290030181600087803b1580156105ad57600080fd5b505af11580156105c1573d6000803e3d6000fd5b505050506040513d60208110156105d757600080fd5b505192915050565b6000546001600160a01b031681565b6001546001600160a01b031681565b6000546001600160a01b0316331461061457600080fd5b60015460408051636ab846ad60e11b81526001600160a01b038581166004830152602482018590529151919092169163d5708d5a91604480830192600092919082900301818387803b15801561066957600080fd5b505af115801561067d573d6000803e3d6000fd5b505050505050565b6000546001600160a01b0316331461069c57600080fd5b6001600160a01b0381166106af57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316331461072157600080fd5b6001546040805163fb52b06560e01b81526001600160a01b038581166004830152602482018590529151919092169163fb52b06591604480830192600092919082900301818387803b15801561066957600080fd5b611081806107848339019056fe60806040526003805460ff60a01b191690553480156200001e57600080fd5b506040516200108138038062001081833981810160405260a08110156200004457600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200007057600080fd5b9083019060208201858111156200008657600080fd5b8251640100000000811182820188101715620000a157600080fd5b82525081516020918201929091019080838360005b83811015620000d0578181015183820152602001620000b6565b50505050905090810190601f168015620000fe5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200012257600080fd5b9083019060208201858111156200013857600080fd5b82516401000000008111828201881017156200015357600080fd5b82525081516020918201929091019080838360005b838110156200018257818101518382015260200162000168565b50505050905090810190601f168015620001b05780820380516001836020036101000a031916815260200191505b5060405260209081015160038054336001600160a01b0319918216179091556005805482166001600160a01b038b8116919091179091556004805490921690891617905585519093506200020b925060069186019062000270565b5081516200022190600790602085019062000270565b506008805460ff191660ff83161790556040516001600160a01b038616907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a2505050505062000315565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002b357805160ff1916838001178555620002e3565b82800160010185558215620002e3579182015b82811115620002e3578251825591602001919060010190620002c6565b50620002f1929150620002f5565b5090565b6200031291905b80821115620002f15760008155600101620002fc565b90565b610d5c80620003256000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806366188463116100ad57806395d89b411161007157806395d89b4114610345578063a9059cbb1461034d578063d73dd62314610379578063dd62ed3e146103a5578063f2fde38b146103d35761012c565b806366188463146102db57806370a08231146103075780637f9864f71461032d5780638456cb59146103355780638da5cb5b1461033d5761012c565b8063313ce567116100f4578063313ce567146102625780633f4ba83a1461028057806340c10f191461028a57806342966c68146102b65780635c975abb146102d35761012c565b806306fdde031461013157806307546172146101ae578063095ea7b3146101d257806318160ddd1461021257806323b872dd1461022c575b600080fd5b6101396103f9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017357818101518382015260200161015b565b50505050905090810190601f1680156101a05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101b6610487565b604080516001600160a01b039092168252519081900360200190f35b6101fe600480360360408110156101e857600080fd5b506001600160a01b038135169060200135610496565b604080519115158252519081900360200190f35b61021a6104fc565b60408051918252519081900360200190f35b6101fe6004803603606081101561024257600080fd5b506001600160a01b03813581169160208101359091169060400135610502565b61026a610665565b6040805160ff9092168252519081900360200190f35b61028861066e565b005b6101fe600480360360408110156102a057600080fd5b506001600160a01b0381351690602001356106d3565b6101fe600480360360208110156102cc57600080fd5b5035610804565b6101fe6108cb565b6101fe600480360360408110156102f157600080fd5b506001600160a01b0381351690602001356108db565b61021a6004803603602081101561031d57600080fd5b50356001600160a01b03166109cb565b6101b66109e6565b6102886109f5565b6101b6610a61565b610139610a70565b6101fe6004803603604081101561036357600080fd5b506001600160a01b038135169060200135610acb565b6101fe6004803603604081101561038f57600080fd5b506001600160a01b038135169060200135610b98565b61021a600480360360408110156103bb57600080fd5b506001600160a01b0381358116916020013516610c31565b610288600480360360208110156103e957600080fd5b50356001600160a01b0316610c5c565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561047f5780601f106104545761010080835404028352916020019161047f565b820191906000526020600020905b81548152906001019060200180831161046257829003601f168201915b505050505081565b6005546001600160a01b031681565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b60006001600160a01b03831661051757600080fd5b6001600160a01b03841660009081526020819052604090205482111561053c57600080fd5b6001600160a01b038416600090815260026020908152604080832033845290915290205482111561056c57600080fd5b6001600160a01b038416600090815260208190526040902054610595908363ffffffff610ce216565b6001600160a01b0380861660009081526020819052604080822093909355908516815220546105ca908363ffffffff610cf416565b6001600160a01b0380851660009081526020818152604080832094909455918716815260028252828120338252909152205461060c908363ffffffff610ce216565b6001600160a01b0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610d08833981519152929181900390910190a35060019392505050565b60085460ff1681565b6003546001600160a01b0316331461068557600080fd5b600354600160a01b900460ff1661069b57600080fd5b6003805460ff60a01b191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b6005546000906001600160a01b03163314610726576040805162461bcd60e51b815260206004820152600e60248201526d3737ba103a34329036b4b73a32b960911b604482015290519081900360640190fd5b600354600160a01b900460ff161561073d57600080fd5b600154610750908363ffffffff610cf416565b6001556001600160a01b03831660009081526020819052604090205461077c908363ffffffff610cf416565b6001600160a01b03841660008181526020818152604091829020939093558051858152905191927f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe92918290030190a26040805183815290516001600160a01b03851691600091600080516020610d088339815191529181900360200190a350600192915050565b3360009081526020819052604081205482111561082057600080fd5b600154610833908363ffffffff610ce216565b60015533600090815260208190526040902054610856908363ffffffff610ce216565b3360008181526020818152604091829020939093558051858152905191927f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df792918290030190a26040805183815290516000913391600080516020610d088339815191529181900360200190a3506001919050565b600354600160a01b900460ff1681565b3360009081526002602090815260408083206001600160a01b038616845290915281205480831115610930573360009081526002602090815260408083206001600160a01b0388168452909152812055610965565b610940818463ffffffff610ce216565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526020819052604090205490565b6004546001600160a01b031681565b6003546001600160a01b03163314610a0c57600080fd5b600354600160a01b900460ff1615610a2357600080fd5b6003805460ff60a01b1916600160a01b1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b6003546001600160a01b031681565b6007805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561047f5780601f106104545761010080835404028352916020019161047f565b60006001600160a01b038316610ae057600080fd5b33600090815260208190526040902054821115610afc57600080fd5b33600090815260208190526040902054610b1c908363ffffffff610ce216565b33600090815260208190526040808220929092556001600160a01b03851681522054610b4e908363ffffffff610cf416565b6001600160a01b03841660008181526020818152604091829020939093558051858152905191923392600080516020610d088339815191529281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610bcc908363ffffffff610cf416565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6003546001600160a01b03163314610c7357600080fd5b6001600160a01b038116610c8657600080fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b600082821115610cee57fe5b50900390565b81810182811015610d0157fe5b9291505056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa265627a7a723158208f534a7473ce0fe3184de99e306787634a9c6b6e689f3f05a8b332349fe0962964736f6c63430005100032a265627a7a72315820c08f0a7a97d26b7ffe2f3dd9e66988ffc24294661d33afcda70295e2001235f664736f6c63430005100032"

// DeployShadowTokenListManager deploys a new Ethereum contract, binding an instance of ShadowTokenListManager to it.
func DeployShadowTokenListManager(auth *bind.TransactOpts, backend bind.ContractBackend, _addr common.Address) (common.Address, *types.Transaction, *ShadowTokenListManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ShadowTokenListManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ShadowTokenListManagerBin), backend, _addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShadowTokenListManager{ShadowTokenListManagerCaller: ShadowTokenListManagerCaller{contract: contract}, ShadowTokenListManagerTransactor: ShadowTokenListManagerTransactor{contract: contract}, ShadowTokenListManagerFilterer: ShadowTokenListManagerFilterer{contract: contract}}, nil
}

// ShadowTokenListManager is an auto generated Go binding around an Ethereum contract.
type ShadowTokenListManager struct {
	ShadowTokenListManagerCaller     // Read-only binding to the contract
	ShadowTokenListManagerTransactor // Write-only binding to the contract
	ShadowTokenListManagerFilterer   // Log filterer for contract events
}

// ShadowTokenListManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShadowTokenListManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowTokenListManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShadowTokenListManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowTokenListManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShadowTokenListManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowTokenListManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShadowTokenListManagerSession struct {
	Contract     *ShadowTokenListManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ShadowTokenListManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShadowTokenListManagerCallerSession struct {
	Contract *ShadowTokenListManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ShadowTokenListManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShadowTokenListManagerTransactorSession struct {
	Contract     *ShadowTokenListManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ShadowTokenListManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShadowTokenListManagerRaw struct {
	Contract *ShadowTokenListManager // Generic contract binding to access the raw methods on
}

// ShadowTokenListManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShadowTokenListManagerCallerRaw struct {
	Contract *ShadowTokenListManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ShadowTokenListManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShadowTokenListManagerTransactorRaw struct {
	Contract *ShadowTokenListManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShadowTokenListManager creates a new instance of ShadowTokenListManager, bound to a specific deployed contract.
func NewShadowTokenListManager(address common.Address, backend bind.ContractBackend) (*ShadowTokenListManager, error) {
	contract, err := bindShadowTokenListManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenListManager{ShadowTokenListManagerCaller: ShadowTokenListManagerCaller{contract: contract}, ShadowTokenListManagerTransactor: ShadowTokenListManagerTransactor{contract: contract}, ShadowTokenListManagerFilterer: ShadowTokenListManagerFilterer{contract: contract}}, nil
}

// NewShadowTokenListManagerCaller creates a new read-only instance of ShadowTokenListManager, bound to a specific deployed contract.
func NewShadowTokenListManagerCaller(address common.Address, caller bind.ContractCaller) (*ShadowTokenListManagerCaller, error) {
	contract, err := bindShadowTokenListManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenListManagerCaller{contract: contract}, nil
}

// NewShadowTokenListManagerTransactor creates a new write-only instance of ShadowTokenListManager, bound to a specific deployed contract.
func NewShadowTokenListManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ShadowTokenListManagerTransactor, error) {
	contract, err := bindShadowTokenListManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenListManagerTransactor{contract: contract}, nil
}

// NewShadowTokenListManagerFilterer creates a new log filterer instance of ShadowTokenListManager, bound to a specific deployed contract.
func NewShadowTokenListManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ShadowTokenListManagerFilterer, error) {
	contract, err := bindShadowTokenListManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenListManagerFilterer{contract: contract}, nil
}

// bindShadowTokenListManager binds a generic wrapper to an already deployed contract.
func bindShadowTokenListManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShadowTokenListManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShadowTokenListManager *ShadowTokenListManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShadowTokenListManager.Contract.ShadowTokenListManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShadowTokenListManager *ShadowTokenListManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.ShadowTokenListManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShadowTokenListManager *ShadowTokenListManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.ShadowTokenListManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShadowTokenListManager *ShadowTokenListManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShadowTokenListManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowTokenListManager *ShadowTokenListManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowTokenListManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowTokenListManager *ShadowTokenListManagerSession) Owner() (common.Address, error) {
	return _ShadowTokenListManager.Contract.Owner(&_ShadowTokenListManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowTokenListManager *ShadowTokenListManagerCallerSession) Owner() (common.Address, error) {
	return _ShadowTokenListManager.Contract.Owner(&_ShadowTokenListManager.CallOpts)
}

// TokenList is a free data retrieval call binding the contract method 0x9e2c58ca.
//
// Solidity: function tokenList() view returns(address)
func (_ShadowTokenListManager *ShadowTokenListManagerCaller) TokenList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowTokenListManager.contract.Call(opts, &out, "tokenList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenList is a free data retrieval call binding the contract method 0x9e2c58ca.
//
// Solidity: function tokenList() view returns(address)
func (_ShadowTokenListManager *ShadowTokenListManagerSession) TokenList() (common.Address, error) {
	return _ShadowTokenListManager.Contract.TokenList(&_ShadowTokenListManager.CallOpts)
}

// TokenList is a free data retrieval call binding the contract method 0x9e2c58ca.
//
// Solidity: function tokenList() view returns(address)
func (_ShadowTokenListManager *ShadowTokenListManagerCallerSession) TokenList() (common.Address, error) {
	return _ShadowTokenListManager.Contract.TokenList(&_ShadowTokenListManager.CallOpts)
}

// AddShadowToken is a paid mutator transaction binding the contract method 0x3f47a17f.
//
// Solidity: function addShadowToken(address _minter, address _coToken, string _name, string _symbol, uint8 _decimals, uint256 _min, uint256 _max) returns(bool success_)
func (_ShadowTokenListManager *ShadowTokenListManagerTransactor) AddShadowToken(opts *bind.TransactOpts, _minter common.Address, _coToken common.Address, _name string, _symbol string, _decimals uint8, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.contract.Transact(opts, "addShadowToken", _minter, _coToken, _name, _symbol, _decimals, _min, _max)
}

// AddShadowToken is a paid mutator transaction binding the contract method 0x3f47a17f.
//
// Solidity: function addShadowToken(address _minter, address _coToken, string _name, string _symbol, uint8 _decimals, uint256 _min, uint256 _max) returns(bool success_)
func (_ShadowTokenListManager *ShadowTokenListManagerSession) AddShadowToken(_minter common.Address, _coToken common.Address, _name string, _symbol string, _decimals uint8, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.AddShadowToken(&_ShadowTokenListManager.TransactOpts, _minter, _coToken, _name, _symbol, _decimals, _min, _max)
}

// AddShadowToken is a paid mutator transaction binding the contract method 0x3f47a17f.
//
// Solidity: function addShadowToken(address _minter, address _coToken, string _name, string _symbol, uint8 _decimals, uint256 _min, uint256 _max) returns(bool success_)
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorSession) AddShadowToken(_minter common.Address, _coToken common.Address, _name string, _symbol string, _decimals uint8, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.AddShadowToken(&_ShadowTokenListManager.TransactOpts, _minter, _coToken, _name, _symbol, _decimals, _min, _max)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns(bool success_)
func (_ShadowTokenListManager *ShadowTokenListManagerTransactor) RemoveToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.contract.Transact(opts, "removeToken", _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns(bool success_)
func (_ShadowTokenListManager *ShadowTokenListManagerSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.RemoveToken(&_ShadowTokenListManager.TransactOpts, _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns(bool success_)
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.RemoveToken(&_ShadowTokenListManager.TransactOpts, _token)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0xfb52b065.
//
// Solidity: function setMaxAmount(address _token, uint256 _maxAmount) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactor) SetMaxAmount(opts *bind.TransactOpts, _token common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.contract.Transact(opts, "setMaxAmount", _token, _maxAmount)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0xfb52b065.
//
// Solidity: function setMaxAmount(address _token, uint256 _maxAmount) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerSession) SetMaxAmount(_token common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.SetMaxAmount(&_ShadowTokenListManager.TransactOpts, _token, _maxAmount)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0xfb52b065.
//
// Solidity: function setMaxAmount(address _token, uint256 _maxAmount) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorSession) SetMaxAmount(_token common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.SetMaxAmount(&_ShadowTokenListManager.TransactOpts, _token, _maxAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0xd5708d5a.
//
// Solidity: function setMinAmount(address _token, uint256 _minAmount) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactor) SetMinAmount(opts *bind.TransactOpts, _token common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.contract.Transact(opts, "setMinAmount", _token, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0xd5708d5a.
//
// Solidity: function setMinAmount(address _token, uint256 _minAmount) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerSession) SetMinAmount(_token common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.SetMinAmount(&_ShadowTokenListManager.TransactOpts, _token, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0xd5708d5a.
//
// Solidity: function setMinAmount(address _token, uint256 _minAmount) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorSession) SetMinAmount(_token common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.SetMinAmount(&_ShadowTokenListManager.TransactOpts, _token, _minAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.TransferOwnership(&_ShadowTokenListManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.TransferOwnership(&_ShadowTokenListManager.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _newManager) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactor) Upgrade(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.contract.Transact(opts, "upgrade", _newManager)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _newManager) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerSession) Upgrade(_newManager common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.Upgrade(&_ShadowTokenListManager.TransactOpts, _newManager)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _newManager) returns()
func (_ShadowTokenListManager *ShadowTokenListManagerTransactorSession) Upgrade(_newManager common.Address) (*types.Transaction, error) {
	return _ShadowTokenListManager.Contract.Upgrade(&_ShadowTokenListManager.TransactOpts, _newManager)
}

// ShadowTokenListManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShadowTokenListManager contract.
type ShadowTokenListManagerOwnershipTransferredIterator struct {
	Event *ShadowTokenListManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShadowTokenListManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowTokenListManagerOwnershipTransferred)
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
		it.Event = new(ShadowTokenListManagerOwnershipTransferred)
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
func (it *ShadowTokenListManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowTokenListManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowTokenListManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ShadowTokenListManager contract.
type ShadowTokenListManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShadowTokenListManager *ShadowTokenListManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShadowTokenListManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShadowTokenListManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShadowTokenListManagerOwnershipTransferredIterator{contract: _ShadowTokenListManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShadowTokenListManager *ShadowTokenListManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShadowTokenListManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShadowTokenListManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowTokenListManagerOwnershipTransferred)
				if err := _ShadowTokenListManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShadowTokenListManager *ShadowTokenListManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ShadowTokenListManagerOwnershipTransferred, error) {
	event := new(ShadowTokenListManagerOwnershipTransferred)
	if err := _ShadowTokenListManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StandardTokenFuncSigs maps the 4-byte function signature to its string representation.
var StandardTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"66188463": "decreaseApproval(address,uint256)",
	"d73dd623": "increaseApproval(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
var StandardTokenBin = "0x608060405234801561001057600080fd5b506106de806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a9059cbb1461016f578063d73dd6231461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063661884631461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561025b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610261565b6100b96004803603604081101561013357600080fd5b506001600160a01b0381351690602001356103d6565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166104c6565b6100b96004803603604081101561018557600080fd5b506001600160a01b0381351690602001356104e1565b6100b9600480360360408110156101b157600080fd5b506001600160a01b0381351690602001356105c0565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610659565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b60006001600160a01b03831661027657600080fd5b6001600160a01b03841660009081526020819052604090205482111561029b57600080fd5b6001600160a01b03841660009081526002602090815260408083203384529091529020548211156102cb57600080fd5b6001600160a01b0384166000908152602081905260409020546102f4908363ffffffff61068416565b6001600160a01b038086166000908152602081905260408082209390935590851681522054610329908363ffffffff61069616565b6001600160a01b0380851660009081526020818152604080832094909455918716815260028252828120338252909152205461036b908363ffffffff61068416565b6001600160a01b03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b3360009081526002602090815260408083206001600160a01b03861684529091528120548083111561042b573360009081526002602090815260408083206001600160a01b0388168452909152812055610460565b61043b818463ffffffff61068416565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526020819052604090205490565b60006001600160a01b0383166104f657600080fd5b3360009081526020819052604090205482111561051257600080fd5b33600090815260208190526040902054610532908363ffffffff61068416565b33600090815260208190526040808220929092556001600160a01b03851681522054610564908363ffffffff61069616565b6001600160a01b038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b03861684529091528120546105f4908363ffffffff61069616565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561069057fe5b50900390565b818101828110156106a357fe5b9291505056fea265627a7a72315820ef5032d1fc82ce3a045f349ac0d275c886bf70d9576b56ab72ee07b2788ffa8864736f6c63430005100032"

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) ParseApproval(log types.Log) (*StandardTokenApproval, error) {
	event := new(StandardTokenApproval)
	if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) ParseTransfer(log types.Log) (*StandardTokenTransfer, error) {
	event := new(StandardTokenTransfer)
	if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenListABI is the input ABI used to generate the binding from.
const TokenListABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"name\":\"TokenUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mins\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_maxs\",\"type\":\"uint256[]\"}],\"name\":\"addTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"limit\",\"type\":\"uint8\"}],\"name\":\"getActiveItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"items_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_item\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_item\",\"type\":\"address\"}],\"name\":\"isExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAmount_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfActive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"}],\"name\":\"setMaxAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"}],\"name\":\"setMinAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenListFuncSigs maps the 4-byte function signature to its string representation.
var TokenListFuncSigs = map[string]string{
	"dee1f2af": "addToken(address,uint256,uint256)",
	"73d432ce": "addTokens(address[],uint256[],uint256[])",
	"06661abd": "count()",
	"f7cb1312": "getActiveItems(uint256,uint8)",
	"9f8a13d7": "isActive(address)",
	"babcc539": "isAllowed(address)",
	"0013eb4b": "isExist(address)",
	"94dee0a4": "maxAmount(address)",
	"4d0a32db": "minAmount(address)",
	"593f6969": "numOfActive()",
	"8da5cb5b": "owner()",
	"5fa7b584": "removeToken(address)",
	"fb52b065": "setMaxAmount(address,uint256)",
	"d5708d5a": "setMinAmount(address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// TokenListBin is the compiled bytecode used for deploying new contracts.
var TokenListBin = "0x6080604052600080546001600160a01b03191633179055610cc1806100256000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c806394dee0a411610097578063dee1f2af11610066578063dee1f2af14610410578063f2fde38b14610442578063f7cb131214610468578063fb52b065146104e9576100f4565b806394dee0a4146103705780639f8a13d714610396578063babcc539146103bc578063d5708d5a146103e2576100f4565b8063593f6969116100d3578063593f6969146101735780635fa7b5841461017b57806373d432ce146101a15780638da5cb5b1461034c576100f4565b806213eb4b146100f957806306661abd146101335780634d0a32db1461014d575b600080fd5b61011f6004803603602081101561010f57600080fd5b50356001600160a01b0316610515565b604080519115158252519081900360200190f35b61013b610537565b60408051918252519081900360200190f35b61013b6004803603602081101561016357600080fd5b50356001600160a01b031661053d565b61013b610569565b61011f6004803603602081101561019157600080fd5b50356001600160a01b031661056f565b61011f600480360360608110156101b757600080fd5b8101906020810181356401000000008111156101d257600080fd5b8201836020820111156101e457600080fd5b8035906020019184602083028401116401000000008311171561020657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561025657600080fd5b82018360208201111561026857600080fd5b8035906020019184602083028401116401000000008311171561028a57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156102da57600080fd5b8201836020820111156102ec57600080fd5b8035906020019184602083028401116401000000008311171561030e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506105d1945050505050565b6103546106ab565b604080516001600160a01b039092168252519081900360200190f35b61013b6004803603602081101561038657600080fd5b50356001600160a01b03166106ba565b61011f600480360360208110156103ac57600080fd5b50356001600160a01b03166106e9565b61011f600480360360208110156103d257600080fd5b50356001600160a01b031661070c565b61040e600480360360408110156103f857600080fd5b506001600160a01b03813516906020013561071d565b005b61011f6004803603606081101561042657600080fd5b506001600160a01b0381351690602081013590604001356107d1565b61040e6004803603602081101561045857600080fd5b50356001600160a01b03166108c8565b61048e6004803603604081101561047e57600080fd5b508035906020013560ff1661094d565b6040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156104d45781810151838201526020016104bc565b50505050905001935050505060405180910390f35b61040e600480360360408110156104ff57600080fd5b506001600160a01b038135169060200135610a6b565b6001600160a01b03811660009081526003602052604090205460ff165b919050565b60025490565b600061054882610515565b1561053257506001600160a01b031660009081526004602052604090205490565b60015490565b600080546001600160a01b0316331461058757600080fd5b61059082610b12565b15610532576040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd390600090a2506001919050565b600080546001600160a01b031633146105e957600080fd5b825184511480156105fb575081518351145b610641576040805162461bcd60e51b8152602060048201526012602482015271696e76616c696420706172616d657465727360701b604482015290519081900360640190fd5b60005b84518110156106a35761069185828151811061065c57fe5b602002602001015185838151811061067057fe5b602002602001015185848151811061068457fe5b60200260200101516107d1565b1561069b57600191505b600101610644565b509392505050565b6000546001600160a01b031681565b60006106c582610515565b1561053257506001600160a01b031660009081526004602052604090206001015490565b6001600160a01b0316600090815260036020526040902054610100900460ff1690565b6000610717826106e9565b92915050565b6000546001600160a01b0316331461073457600080fd5b61073d82610515565b610780576040805162461bcd60e51b815260206004820152600f60248201526e1d1bdad95b881b9bdd081859191959608a1b604482015290519081900360640190fd5b6001600160a01b0382166000908152600460205260409020600101548111156107a857600080fd5b600081116107b557600080fd5b6001600160a01b03909116600090815260046020526040902055565b600080546001600160a01b031633146107e957600080fd5b6107f284610b95565b156108c15760008311801561080657508282115b61084c576040805162461bcd60e51b8152602060048201526012602482015271696e76616c696420706172616d657465727360701b604482015290519081900360640190fd5b60408051808201825284815260208082018581526001600160a01b0388166000818152600484528590209351845590516001909301929092558251868152908101859052825191927fa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc92918290030190a25060015b9392505050565b6000546001600160a01b031633146108df57600080fd5b6001600160a01b0381166108f257600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60025460009060609084108015610966575060ff831615155b61096f57600080fd5b8260ff1660405190808252806020026020018201604052801561099c578160200160208202803883390190505b50905060005b8360ff16811015610a6357600254858201106109bd57610a63565b600360006002838801815481106109d057fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff6101009091041615610a5b57600281860181548110610a1257fe5b9060005260206000200160009054906101000a90046001600160a01b0316828481518110610a3c57fe5b6001600160a01b03909216602092830291909101909101526001909201915b6001016109a2565b509250929050565b6000546001600160a01b03163314610a8257600080fd5b610a8b82610515565b610ace576040805162461bcd60e51b815260206004820152600f60248201526e1d1bdad95b881b9bdd081859191959608a1b604482015290519081900360640190fd5b6001600160a01b038216600090815260046020526040902054811015610af357600080fd5b6001600160a01b03909116600090815260046020526040902060010155565b6001600160a01b03811660009081526003602052604081205460ff168015610b5757506001600160a01b038216600090815260036020526040902054610100900460ff165b15610b8d5750600180546000190181556001600160a01b0382166000908152600360205260409020805461ff0019169055610532565b506000919050565b6001600160a01b038116600090815260036020526040812054610100900460ff1615610bc357506000610532565b6001600160a01b03821660009081526003602052604090205460ff16610c2f57600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b0384161790555b50600180548101815560408051808201825282815260208082018481526001600160a01b0386166000908152600390925292902090518154925115156101000261ff001991151560ff19909416939093171691909117905591905056fea265627a7a72315820e100cce9011f43425183eab79c5bcd27b624a0227db84f5b7764f482d7d9e38764736f6c63430005100032"

// DeployTokenList deploys a new Ethereum contract, binding an instance of TokenList to it.
func DeployTokenList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenList, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenList{TokenListCaller: TokenListCaller{contract: contract}, TokenListTransactor: TokenListTransactor{contract: contract}, TokenListFilterer: TokenListFilterer{contract: contract}}, nil
}

// TokenList is an auto generated Go binding around an Ethereum contract.
type TokenList struct {
	TokenListCaller     // Read-only binding to the contract
	TokenListTransactor // Write-only binding to the contract
	TokenListFilterer   // Log filterer for contract events
}

// TokenListCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenListSession struct {
	Contract     *TokenList        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenListCallerSession struct {
	Contract *TokenListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenListTransactorSession struct {
	Contract     *TokenListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenListRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenListRaw struct {
	Contract *TokenList // Generic contract binding to access the raw methods on
}

// TokenListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenListCallerRaw struct {
	Contract *TokenListCaller // Generic read-only contract binding to access the raw methods on
}

// TokenListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenListTransactorRaw struct {
	Contract *TokenListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenList creates a new instance of TokenList, bound to a specific deployed contract.
func NewTokenList(address common.Address, backend bind.ContractBackend) (*TokenList, error) {
	contract, err := bindTokenList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenList{TokenListCaller: TokenListCaller{contract: contract}, TokenListTransactor: TokenListTransactor{contract: contract}, TokenListFilterer: TokenListFilterer{contract: contract}}, nil
}

// NewTokenListCaller creates a new read-only instance of TokenList, bound to a specific deployed contract.
func NewTokenListCaller(address common.Address, caller bind.ContractCaller) (*TokenListCaller, error) {
	contract, err := bindTokenList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenListCaller{contract: contract}, nil
}

// NewTokenListTransactor creates a new write-only instance of TokenList, bound to a specific deployed contract.
func NewTokenListTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenListTransactor, error) {
	contract, err := bindTokenList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenListTransactor{contract: contract}, nil
}

// NewTokenListFilterer creates a new log filterer instance of TokenList, bound to a specific deployed contract.
func NewTokenListFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenListFilterer, error) {
	contract, err := bindTokenList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenListFilterer{contract: contract}, nil
}

// bindTokenList binds a generic wrapper to an already deployed contract.
func bindTokenList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenList *TokenListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenList.Contract.TokenListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenList *TokenListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenList.Contract.TokenListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenList *TokenListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenList.Contract.TokenListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenList *TokenListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenList *TokenListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenList *TokenListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenList.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_TokenList *TokenListCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_TokenList *TokenListSession) Count() (*big.Int, error) {
	return _TokenList.Contract.Count(&_TokenList.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_TokenList *TokenListCallerSession) Count() (*big.Int, error) {
	return _TokenList.Contract.Count(&_TokenList.CallOpts)
}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_TokenList *TokenListCaller) GetActiveItems(opts *bind.CallOpts, offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "getActiveItems", offset, limit)

	outstruct := new(struct {
		Count *big.Int
		Items []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Items = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_TokenList *TokenListSession) GetActiveItems(offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	return _TokenList.Contract.GetActiveItems(&_TokenList.CallOpts, offset, limit)
}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_TokenList *TokenListCallerSession) GetActiveItems(offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	return _TokenList.Contract.GetActiveItems(&_TokenList.CallOpts, offset, limit)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_TokenList *TokenListCaller) IsActive(opts *bind.CallOpts, _item common.Address) (bool, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "isActive", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_TokenList *TokenListSession) IsActive(_item common.Address) (bool, error) {
	return _TokenList.Contract.IsActive(&_TokenList.CallOpts, _item)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_TokenList *TokenListCallerSession) IsActive(_item common.Address) (bool, error) {
	return _TokenList.Contract.IsActive(&_TokenList.CallOpts, _item)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address _token) view returns(bool)
func (_TokenList *TokenListCaller) IsAllowed(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "isAllowed", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address _token) view returns(bool)
func (_TokenList *TokenListSession) IsAllowed(_token common.Address) (bool, error) {
	return _TokenList.Contract.IsAllowed(&_TokenList.CallOpts, _token)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address _token) view returns(bool)
func (_TokenList *TokenListCallerSession) IsAllowed(_token common.Address) (bool, error) {
	return _TokenList.Contract.IsAllowed(&_TokenList.CallOpts, _token)
}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_TokenList *TokenListCaller) IsExist(opts *bind.CallOpts, _item common.Address) (bool, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "isExist", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_TokenList *TokenListSession) IsExist(_item common.Address) (bool, error) {
	return _TokenList.Contract.IsExist(&_TokenList.CallOpts, _item)
}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_TokenList *TokenListCallerSession) IsExist(_item common.Address) (bool, error) {
	return _TokenList.Contract.IsExist(&_TokenList.CallOpts, _item)
}

// MaxAmount is a free data retrieval call binding the contract method 0x94dee0a4.
//
// Solidity: function maxAmount(address _token) view returns(uint256 maxAmount_)
func (_TokenList *TokenListCaller) MaxAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "maxAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAmount is a free data retrieval call binding the contract method 0x94dee0a4.
//
// Solidity: function maxAmount(address _token) view returns(uint256 maxAmount_)
func (_TokenList *TokenListSession) MaxAmount(_token common.Address) (*big.Int, error) {
	return _TokenList.Contract.MaxAmount(&_TokenList.CallOpts, _token)
}

// MaxAmount is a free data retrieval call binding the contract method 0x94dee0a4.
//
// Solidity: function maxAmount(address _token) view returns(uint256 maxAmount_)
func (_TokenList *TokenListCallerSession) MaxAmount(_token common.Address) (*big.Int, error) {
	return _TokenList.Contract.MaxAmount(&_TokenList.CallOpts, _token)
}

// MinAmount is a free data retrieval call binding the contract method 0x4d0a32db.
//
// Solidity: function minAmount(address _token) view returns(uint256 minAmount_)
func (_TokenList *TokenListCaller) MinAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "minAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x4d0a32db.
//
// Solidity: function minAmount(address _token) view returns(uint256 minAmount_)
func (_TokenList *TokenListSession) MinAmount(_token common.Address) (*big.Int, error) {
	return _TokenList.Contract.MinAmount(&_TokenList.CallOpts, _token)
}

// MinAmount is a free data retrieval call binding the contract method 0x4d0a32db.
//
// Solidity: function minAmount(address _token) view returns(uint256 minAmount_)
func (_TokenList *TokenListCallerSession) MinAmount(_token common.Address) (*big.Int, error) {
	return _TokenList.Contract.MinAmount(&_TokenList.CallOpts, _token)
}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_TokenList *TokenListCaller) NumOfActive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "numOfActive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_TokenList *TokenListSession) NumOfActive() (*big.Int, error) {
	return _TokenList.Contract.NumOfActive(&_TokenList.CallOpts)
}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_TokenList *TokenListCallerSession) NumOfActive() (*big.Int, error) {
	return _TokenList.Contract.NumOfActive(&_TokenList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenList *TokenListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenList *TokenListSession) Owner() (common.Address, error) {
	return _TokenList.Contract.Owner(&_TokenList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenList *TokenListCallerSession) Owner() (common.Address, error) {
	return _TokenList.Contract.Owner(&_TokenList.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xdee1f2af.
//
// Solidity: function addToken(address _token, uint256 _min, uint256 _max) returns(bool success_)
func (_TokenList *TokenListTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _TokenList.contract.Transact(opts, "addToken", _token, _min, _max)
}

// AddToken is a paid mutator transaction binding the contract method 0xdee1f2af.
//
// Solidity: function addToken(address _token, uint256 _min, uint256 _max) returns(bool success_)
func (_TokenList *TokenListSession) AddToken(_token common.Address, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.AddToken(&_TokenList.TransactOpts, _token, _min, _max)
}

// AddToken is a paid mutator transaction binding the contract method 0xdee1f2af.
//
// Solidity: function addToken(address _token, uint256 _min, uint256 _max) returns(bool success_)
func (_TokenList *TokenListTransactorSession) AddToken(_token common.Address, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.AddToken(&_TokenList.TransactOpts, _token, _min, _max)
}

// AddTokens is a paid mutator transaction binding the contract method 0x73d432ce.
//
// Solidity: function addTokens(address[] _tokens, uint256[] _mins, uint256[] _maxs) returns(bool success_)
func (_TokenList *TokenListTransactor) AddTokens(opts *bind.TransactOpts, _tokens []common.Address, _mins []*big.Int, _maxs []*big.Int) (*types.Transaction, error) {
	return _TokenList.contract.Transact(opts, "addTokens", _tokens, _mins, _maxs)
}

// AddTokens is a paid mutator transaction binding the contract method 0x73d432ce.
//
// Solidity: function addTokens(address[] _tokens, uint256[] _mins, uint256[] _maxs) returns(bool success_)
func (_TokenList *TokenListSession) AddTokens(_tokens []common.Address, _mins []*big.Int, _maxs []*big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.AddTokens(&_TokenList.TransactOpts, _tokens, _mins, _maxs)
}

// AddTokens is a paid mutator transaction binding the contract method 0x73d432ce.
//
// Solidity: function addTokens(address[] _tokens, uint256[] _mins, uint256[] _maxs) returns(bool success_)
func (_TokenList *TokenListTransactorSession) AddTokens(_tokens []common.Address, _mins []*big.Int, _maxs []*big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.AddTokens(&_TokenList.TransactOpts, _tokens, _mins, _maxs)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns(bool success_)
func (_TokenList *TokenListTransactor) RemoveToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _TokenList.contract.Transact(opts, "removeToken", _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns(bool success_)
func (_TokenList *TokenListSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _TokenList.Contract.RemoveToken(&_TokenList.TransactOpts, _token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _token) returns(bool success_)
func (_TokenList *TokenListTransactorSession) RemoveToken(_token common.Address) (*types.Transaction, error) {
	return _TokenList.Contract.RemoveToken(&_TokenList.TransactOpts, _token)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0xfb52b065.
//
// Solidity: function setMaxAmount(address _token, uint256 _maxAmount) returns()
func (_TokenList *TokenListTransactor) SetMaxAmount(opts *bind.TransactOpts, _token common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _TokenList.contract.Transact(opts, "setMaxAmount", _token, _maxAmount)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0xfb52b065.
//
// Solidity: function setMaxAmount(address _token, uint256 _maxAmount) returns()
func (_TokenList *TokenListSession) SetMaxAmount(_token common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.SetMaxAmount(&_TokenList.TransactOpts, _token, _maxAmount)
}

// SetMaxAmount is a paid mutator transaction binding the contract method 0xfb52b065.
//
// Solidity: function setMaxAmount(address _token, uint256 _maxAmount) returns()
func (_TokenList *TokenListTransactorSession) SetMaxAmount(_token common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.SetMaxAmount(&_TokenList.TransactOpts, _token, _maxAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0xd5708d5a.
//
// Solidity: function setMinAmount(address _token, uint256 _minAmount) returns()
func (_TokenList *TokenListTransactor) SetMinAmount(opts *bind.TransactOpts, _token common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _TokenList.contract.Transact(opts, "setMinAmount", _token, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0xd5708d5a.
//
// Solidity: function setMinAmount(address _token, uint256 _minAmount) returns()
func (_TokenList *TokenListSession) SetMinAmount(_token common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.SetMinAmount(&_TokenList.TransactOpts, _token, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0xd5708d5a.
//
// Solidity: function setMinAmount(address _token, uint256 _minAmount) returns()
func (_TokenList *TokenListTransactorSession) SetMinAmount(_token common.Address, _minAmount *big.Int) (*types.Transaction, error) {
	return _TokenList.Contract.SetMinAmount(&_TokenList.TransactOpts, _token, _minAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenList *TokenListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenList *TokenListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenList.Contract.TransferOwnership(&_TokenList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenList *TokenListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenList.Contract.TransferOwnership(&_TokenList.TransactOpts, newOwner)
}

// TokenListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenList contract.
type TokenListOwnershipTransferredIterator struct {
	Event *TokenListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenListOwnershipTransferred)
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
		it.Event = new(TokenListOwnershipTransferred)
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
func (it *TokenListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenListOwnershipTransferred represents a OwnershipTransferred event raised by the TokenList contract.
type TokenListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenList *TokenListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenListOwnershipTransferredIterator{contract: _TokenList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenList *TokenListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenListOwnershipTransferred)
				if err := _TokenList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenList *TokenListFilterer) ParseOwnershipTransferred(log types.Log) (*TokenListOwnershipTransferred, error) {
	event := new(TokenListOwnershipTransferred)
	if err := _TokenList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenListTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the TokenList contract.
type TokenListTokenAddedIterator struct {
	Event *TokenListTokenAdded // Event containing the contract specifics and raw log

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
func (it *TokenListTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenListTokenAdded)
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
		it.Event = new(TokenListTokenAdded)
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
func (it *TokenListTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenListTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenListTokenAdded represents a TokenAdded event raised by the TokenList contract.
type TokenListTokenAdded struct {
	Token     common.Address
	MinAmount *big.Int
	MaxAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(address indexed token, uint256 minAmount, uint256 maxAmount)
func (_TokenList *TokenListFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*TokenListTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenList.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenListTokenAddedIterator{contract: _TokenList.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(address indexed token, uint256 minAmount, uint256 maxAmount)
func (_TokenList *TokenListFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *TokenListTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenList.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenListTokenAdded)
				if err := _TokenList.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(address indexed token, uint256 minAmount, uint256 maxAmount)
func (_TokenList *TokenListFilterer) ParseTokenAdded(log types.Log) (*TokenListTokenAdded, error) {
	event := new(TokenListTokenAdded)
	if err := _TokenList.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenListTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the TokenList contract.
type TokenListTokenRemovedIterator struct {
	Event *TokenListTokenRemoved // Event containing the contract specifics and raw log

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
func (it *TokenListTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenListTokenRemoved)
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
		it.Event = new(TokenListTokenRemoved)
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
func (it *TokenListTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenListTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenListTokenRemoved represents a TokenRemoved event raised by the TokenList contract.
type TokenListTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_TokenList *TokenListFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*TokenListTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenList.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenListTokenRemovedIterator{contract: _TokenList.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_TokenList *TokenListFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *TokenListTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenList.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenListTokenRemoved)
				if err := _TokenList.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_TokenList *TokenListFilterer) ParseTokenRemoved(log types.Log) (*TokenListTokenRemoved, error) {
	event := new(TokenListTokenRemoved)
	if err := _TokenList.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenListTokenUpdatedIterator is returned from FilterTokenUpdated and is used to iterate over the raw logs and unpacked data for TokenUpdated events raised by the TokenList contract.
type TokenListTokenUpdatedIterator struct {
	Event *TokenListTokenUpdated // Event containing the contract specifics and raw log

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
func (it *TokenListTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenListTokenUpdated)
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
		it.Event = new(TokenListTokenUpdated)
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
func (it *TokenListTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenListTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenListTokenUpdated represents a TokenUpdated event raised by the TokenList contract.
type TokenListTokenUpdated struct {
	Token     common.Address
	MinAmount *big.Int
	MaxAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenUpdated is a free log retrieval operation binding the contract event 0x5d4804fe0ec949f552f757bfb400c951422d44c10c004077ecd19a9d1f503562.
//
// Solidity: event TokenUpdated(address indexed token, uint256 minAmount, uint256 maxAmount)
func (_TokenList *TokenListFilterer) FilterTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*TokenListTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenList.contract.FilterLogs(opts, "TokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenListTokenUpdatedIterator{contract: _TokenList.contract, event: "TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenUpdated is a free log subscription operation binding the contract event 0x5d4804fe0ec949f552f757bfb400c951422d44c10c004077ecd19a9d1f503562.
//
// Solidity: event TokenUpdated(address indexed token, uint256 minAmount, uint256 maxAmount)
func (_TokenList *TokenListFilterer) WatchTokenUpdated(opts *bind.WatchOpts, sink chan<- *TokenListTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenList.contract.WatchLogs(opts, "TokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenListTokenUpdated)
				if err := _TokenList.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
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

// ParseTokenUpdated is a log parse operation binding the contract event 0x5d4804fe0ec949f552f757bfb400c951422d44c10c004077ecd19a9d1f503562.
//
// Solidity: event TokenUpdated(address indexed token, uint256 minAmount, uint256 maxAmount)
func (_TokenList *TokenListFilterer) ParseTokenUpdated(log types.Log) (*TokenListTokenUpdated, error) {
	event := new(TokenListTokenUpdated)
	if err := _TokenList.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniqueAppendOnlyAddressListABI is the input ABI used to generate the binding from.
const UniqueAppendOnlyAddressListABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"limit\",\"type\":\"uint8\"}],\"name\":\"getActiveItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"items_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_item\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_item\",\"type\":\"address\"}],\"name\":\"isExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfActive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniqueAppendOnlyAddressListFuncSigs maps the 4-byte function signature to its string representation.
var UniqueAppendOnlyAddressListFuncSigs = map[string]string{
	"06661abd": "count()",
	"f7cb1312": "getActiveItems(uint256,uint8)",
	"9f8a13d7": "isActive(address)",
	"0013eb4b": "isExist(address)",
	"593f6969": "numOfActive()",
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// UniqueAppendOnlyAddressListBin is the compiled bytecode used for deploying new contracts.
var UniqueAppendOnlyAddressListBin = "0x6080604052600080546001600160a01b03191633179055610404806100256000396000f3fe608060405234801561001057600080fd5b506004361061007c5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100dd5780639f8a13d714610101578063f2fde38b14610127578063f7cb13121461014f5761007c565b806213eb4b1461008157806306661abd146100bb578063593f6969146100d5575b600080fd5b6100a76004803603602081101561009757600080fd5b50356001600160a01b03166101d0565b604080519115158252519081900360200190f35b6100c36101ee565b60408051918252519081900360200190f35b6100c36101f4565b6100e56101fa565b604080516001600160a01b039092168252519081900360200190f35b6100a76004803603602081101561011757600080fd5b50356001600160a01b0316610209565b61014d6004803603602081101561013d57600080fd5b50356001600160a01b031661022c565b005b6101756004803603604081101561016557600080fd5b508035906020013560ff166102b1565b6040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156101bb5781810151838201526020016101a3565b50505050905001935050505060405180910390f35b6001600160a01b031660009081526003602052604090205460ff1690565b60025490565b60015490565b6000546001600160a01b031681565b6001600160a01b0316600090815260036020526040902054610100900460ff1690565b6000546001600160a01b0316331461024357600080fd5b6001600160a01b03811661025657600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600254600090606090841080156102ca575060ff831615155b6102d357600080fd5b8260ff16604051908082528060200260200182016040528015610300578160200160208202803883390190505b50905060005b8360ff168110156103c75760025485820110610321576103c7565b6003600060028388018154811061033457fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff61010090910416156103bf5760028186018154811061037657fe5b9060005260206000200160009054906101000a90046001600160a01b03168284815181106103a057fe5b6001600160a01b03909216602092830291909101909101526001909201915b600101610306565b50925092905056fea265627a7a7231582009e4d18275903642c8a6bd28652d49ea1b3a5bbf50af33ed89576bd6fc73436a64736f6c63430005100032"

// DeployUniqueAppendOnlyAddressList deploys a new Ethereum contract, binding an instance of UniqueAppendOnlyAddressList to it.
func DeployUniqueAppendOnlyAddressList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniqueAppendOnlyAddressList, error) {
	parsed, err := abi.JSON(strings.NewReader(UniqueAppendOnlyAddressListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniqueAppendOnlyAddressListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniqueAppendOnlyAddressList{UniqueAppendOnlyAddressListCaller: UniqueAppendOnlyAddressListCaller{contract: contract}, UniqueAppendOnlyAddressListTransactor: UniqueAppendOnlyAddressListTransactor{contract: contract}, UniqueAppendOnlyAddressListFilterer: UniqueAppendOnlyAddressListFilterer{contract: contract}}, nil
}

// UniqueAppendOnlyAddressList is an auto generated Go binding around an Ethereum contract.
type UniqueAppendOnlyAddressList struct {
	UniqueAppendOnlyAddressListCaller     // Read-only binding to the contract
	UniqueAppendOnlyAddressListTransactor // Write-only binding to the contract
	UniqueAppendOnlyAddressListFilterer   // Log filterer for contract events
}

// UniqueAppendOnlyAddressListCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniqueAppendOnlyAddressListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniqueAppendOnlyAddressListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniqueAppendOnlyAddressListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniqueAppendOnlyAddressListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniqueAppendOnlyAddressListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniqueAppendOnlyAddressListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniqueAppendOnlyAddressListSession struct {
	Contract     *UniqueAppendOnlyAddressList // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniqueAppendOnlyAddressListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniqueAppendOnlyAddressListCallerSession struct {
	Contract *UniqueAppendOnlyAddressListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// UniqueAppendOnlyAddressListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniqueAppendOnlyAddressListTransactorSession struct {
	Contract     *UniqueAppendOnlyAddressListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// UniqueAppendOnlyAddressListRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniqueAppendOnlyAddressListRaw struct {
	Contract *UniqueAppendOnlyAddressList // Generic contract binding to access the raw methods on
}

// UniqueAppendOnlyAddressListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniqueAppendOnlyAddressListCallerRaw struct {
	Contract *UniqueAppendOnlyAddressListCaller // Generic read-only contract binding to access the raw methods on
}

// UniqueAppendOnlyAddressListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniqueAppendOnlyAddressListTransactorRaw struct {
	Contract *UniqueAppendOnlyAddressListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniqueAppendOnlyAddressList creates a new instance of UniqueAppendOnlyAddressList, bound to a specific deployed contract.
func NewUniqueAppendOnlyAddressList(address common.Address, backend bind.ContractBackend) (*UniqueAppendOnlyAddressList, error) {
	contract, err := bindUniqueAppendOnlyAddressList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniqueAppendOnlyAddressList{UniqueAppendOnlyAddressListCaller: UniqueAppendOnlyAddressListCaller{contract: contract}, UniqueAppendOnlyAddressListTransactor: UniqueAppendOnlyAddressListTransactor{contract: contract}, UniqueAppendOnlyAddressListFilterer: UniqueAppendOnlyAddressListFilterer{contract: contract}}, nil
}

// NewUniqueAppendOnlyAddressListCaller creates a new read-only instance of UniqueAppendOnlyAddressList, bound to a specific deployed contract.
func NewUniqueAppendOnlyAddressListCaller(address common.Address, caller bind.ContractCaller) (*UniqueAppendOnlyAddressListCaller, error) {
	contract, err := bindUniqueAppendOnlyAddressList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniqueAppendOnlyAddressListCaller{contract: contract}, nil
}

// NewUniqueAppendOnlyAddressListTransactor creates a new write-only instance of UniqueAppendOnlyAddressList, bound to a specific deployed contract.
func NewUniqueAppendOnlyAddressListTransactor(address common.Address, transactor bind.ContractTransactor) (*UniqueAppendOnlyAddressListTransactor, error) {
	contract, err := bindUniqueAppendOnlyAddressList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniqueAppendOnlyAddressListTransactor{contract: contract}, nil
}

// NewUniqueAppendOnlyAddressListFilterer creates a new log filterer instance of UniqueAppendOnlyAddressList, bound to a specific deployed contract.
func NewUniqueAppendOnlyAddressListFilterer(address common.Address, filterer bind.ContractFilterer) (*UniqueAppendOnlyAddressListFilterer, error) {
	contract, err := bindUniqueAppendOnlyAddressList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniqueAppendOnlyAddressListFilterer{contract: contract}, nil
}

// bindUniqueAppendOnlyAddressList binds a generic wrapper to an already deployed contract.
func bindUniqueAppendOnlyAddressList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniqueAppendOnlyAddressListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniqueAppendOnlyAddressList.Contract.UniqueAppendOnlyAddressListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniqueAppendOnlyAddressList.Contract.UniqueAppendOnlyAddressListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniqueAppendOnlyAddressList.Contract.UniqueAppendOnlyAddressListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniqueAppendOnlyAddressList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniqueAppendOnlyAddressList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniqueAppendOnlyAddressList.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniqueAppendOnlyAddressList.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListSession) Count() (*big.Int, error) {
	return _UniqueAppendOnlyAddressList.Contract.Count(&_UniqueAppendOnlyAddressList.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCallerSession) Count() (*big.Int, error) {
	return _UniqueAppendOnlyAddressList.Contract.Count(&_UniqueAppendOnlyAddressList.CallOpts)
}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCaller) GetActiveItems(opts *bind.CallOpts, offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	var out []interface{}
	err := _UniqueAppendOnlyAddressList.contract.Call(opts, &out, "getActiveItems", offset, limit)

	outstruct := new(struct {
		Count *big.Int
		Items []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Items = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListSession) GetActiveItems(offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	return _UniqueAppendOnlyAddressList.Contract.GetActiveItems(&_UniqueAppendOnlyAddressList.CallOpts, offset, limit)
}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCallerSession) GetActiveItems(offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	return _UniqueAppendOnlyAddressList.Contract.GetActiveItems(&_UniqueAppendOnlyAddressList.CallOpts, offset, limit)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCaller) IsActive(opts *bind.CallOpts, _item common.Address) (bool, error) {
	var out []interface{}
	err := _UniqueAppendOnlyAddressList.contract.Call(opts, &out, "isActive", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListSession) IsActive(_item common.Address) (bool, error) {
	return _UniqueAppendOnlyAddressList.Contract.IsActive(&_UniqueAppendOnlyAddressList.CallOpts, _item)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCallerSession) IsActive(_item common.Address) (bool, error) {
	return _UniqueAppendOnlyAddressList.Contract.IsActive(&_UniqueAppendOnlyAddressList.CallOpts, _item)
}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCaller) IsExist(opts *bind.CallOpts, _item common.Address) (bool, error) {
	var out []interface{}
	err := _UniqueAppendOnlyAddressList.contract.Call(opts, &out, "isExist", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListSession) IsExist(_item common.Address) (bool, error) {
	return _UniqueAppendOnlyAddressList.Contract.IsExist(&_UniqueAppendOnlyAddressList.CallOpts, _item)
}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCallerSession) IsExist(_item common.Address) (bool, error) {
	return _UniqueAppendOnlyAddressList.Contract.IsExist(&_UniqueAppendOnlyAddressList.CallOpts, _item)
}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCaller) NumOfActive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniqueAppendOnlyAddressList.contract.Call(opts, &out, "numOfActive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListSession) NumOfActive() (*big.Int, error) {
	return _UniqueAppendOnlyAddressList.Contract.NumOfActive(&_UniqueAppendOnlyAddressList.CallOpts)
}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCallerSession) NumOfActive() (*big.Int, error) {
	return _UniqueAppendOnlyAddressList.Contract.NumOfActive(&_UniqueAppendOnlyAddressList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniqueAppendOnlyAddressList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListSession) Owner() (common.Address, error) {
	return _UniqueAppendOnlyAddressList.Contract.Owner(&_UniqueAppendOnlyAddressList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListCallerSession) Owner() (common.Address, error) {
	return _UniqueAppendOnlyAddressList.Contract.Owner(&_UniqueAppendOnlyAddressList.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UniqueAppendOnlyAddressList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniqueAppendOnlyAddressList.Contract.TransferOwnership(&_UniqueAppendOnlyAddressList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniqueAppendOnlyAddressList.Contract.TransferOwnership(&_UniqueAppendOnlyAddressList.TransactOpts, newOwner)
}

// UniqueAppendOnlyAddressListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniqueAppendOnlyAddressList contract.
type UniqueAppendOnlyAddressListOwnershipTransferredIterator struct {
	Event *UniqueAppendOnlyAddressListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniqueAppendOnlyAddressListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniqueAppendOnlyAddressListOwnershipTransferred)
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
		it.Event = new(UniqueAppendOnlyAddressListOwnershipTransferred)
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
func (it *UniqueAppendOnlyAddressListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniqueAppendOnlyAddressListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniqueAppendOnlyAddressListOwnershipTransferred represents a OwnershipTransferred event raised by the UniqueAppendOnlyAddressList contract.
type UniqueAppendOnlyAddressListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniqueAppendOnlyAddressListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniqueAppendOnlyAddressList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniqueAppendOnlyAddressListOwnershipTransferredIterator{contract: _UniqueAppendOnlyAddressList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniqueAppendOnlyAddressListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniqueAppendOnlyAddressList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniqueAppendOnlyAddressListOwnershipTransferred)
				if err := _UniqueAppendOnlyAddressList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UniqueAppendOnlyAddressList *UniqueAppendOnlyAddressListFilterer) ParseOwnershipTransferred(log types.Log) (*UniqueAppendOnlyAddressListOwnershipTransferred, error) {
	event := new(UniqueAppendOnlyAddressListOwnershipTransferred)
	if err := _UniqueAppendOnlyAddressList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
