// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package og

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

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"}],\"name\":\"addressPrivateSlotsChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRemaining\",\"type\":\"uint256\"}],\"name\":\"addressPublicSlotsChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"privateSaleState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"publicSaleState\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLORS_VARIATIONS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MAX_PURCHASE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"OG_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"isFreeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPrivateSaleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"flipPrivateSaleState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPublicSaleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"isPublicSaleEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"flipPublicSaleState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numberOfTokens\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"addWhitelistAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addressPrivateSaleSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addressPublicSaleSlots\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numberOfTokens\",\"type\":\"uint8\"},{\"internalType\":\"uint8[]\",\"name\":\"colors\",\"type\":\"uint8[]\"}],\"name\":\"mintPrivate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numberOfTokens\",\"type\":\"uint8\"},{\"internalType\":\"uint8[]\",\"name\":\"colors\",\"type\":\"uint8[]\"}],\"name\":\"mintPublic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"color\",\"type\":\"uint8\"}],\"name\":\"mintGiftToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenColor\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"availableForMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// COLORSVARIATIONS is a free data retrieval call binding the contract method 0x26e0032d.
//
// Solidity: function COLORS_VARIATIONS() view returns(uint8)
func (_Token *TokenCaller) COLORSVARIATIONS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "COLORS_VARIATIONS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// COLORSVARIATIONS is a free data retrieval call binding the contract method 0x26e0032d.
//
// Solidity: function COLORS_VARIATIONS() view returns(uint8)
func (_Token *TokenSession) COLORSVARIATIONS() (uint8, error) {
	return _Token.Contract.COLORSVARIATIONS(&_Token.CallOpts)
}

// COLORSVARIATIONS is a free data retrieval call binding the contract method 0x26e0032d.
//
// Solidity: function COLORS_VARIATIONS() view returns(uint8)
func (_Token *TokenCallerSession) COLORSVARIATIONS() (uint8, error) {
	return _Token.Contract.COLORSVARIATIONS(&_Token.CallOpts)
}

// MAXPURCHASE is a free data retrieval call binding the contract method 0x7146bd08.
//
// Solidity: function MAX_PURCHASE() view returns(uint8)
func (_Token *TokenCaller) MAXPURCHASE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "MAX_PURCHASE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXPURCHASE is a free data retrieval call binding the contract method 0x7146bd08.
//
// Solidity: function MAX_PURCHASE() view returns(uint8)
func (_Token *TokenSession) MAXPURCHASE() (uint8, error) {
	return _Token.Contract.MAXPURCHASE(&_Token.CallOpts)
}

// MAXPURCHASE is a free data retrieval call binding the contract method 0x7146bd08.
//
// Solidity: function MAX_PURCHASE() view returns(uint8)
func (_Token *TokenCallerSession) MAXPURCHASE() (uint8, error) {
	return _Token.Contract.MAXPURCHASE(&_Token.CallOpts)
}

// OGPRICE is a free data retrieval call binding the contract method 0x90069b42.
//
// Solidity: function OG_PRICE() view returns(uint256)
func (_Token *TokenCaller) OGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "OG_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OGPRICE is a free data retrieval call binding the contract method 0x90069b42.
//
// Solidity: function OG_PRICE() view returns(uint256)
func (_Token *TokenSession) OGPRICE() (*big.Int, error) {
	return _Token.Contract.OGPRICE(&_Token.CallOpts)
}

// OGPRICE is a free data retrieval call binding the contract method 0x90069b42.
//
// Solidity: function OG_PRICE() view returns(uint256)
func (_Token *TokenCallerSession) OGPRICE() (*big.Int, error) {
	return _Token.Contract.OGPRICE(&_Token.CallOpts)
}

// AddressPrivateSaleSlots is a free data retrieval call binding the contract method 0x8ec33726.
//
// Solidity: function addressPrivateSaleSlots(address _address) view returns(uint256)
func (_Token *TokenCaller) AddressPrivateSaleSlots(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "addressPrivateSaleSlots", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressPrivateSaleSlots is a free data retrieval call binding the contract method 0x8ec33726.
//
// Solidity: function addressPrivateSaleSlots(address _address) view returns(uint256)
func (_Token *TokenSession) AddressPrivateSaleSlots(_address common.Address) (*big.Int, error) {
	return _Token.Contract.AddressPrivateSaleSlots(&_Token.CallOpts, _address)
}

// AddressPrivateSaleSlots is a free data retrieval call binding the contract method 0x8ec33726.
//
// Solidity: function addressPrivateSaleSlots(address _address) view returns(uint256)
func (_Token *TokenCallerSession) AddressPrivateSaleSlots(_address common.Address) (*big.Int, error) {
	return _Token.Contract.AddressPrivateSaleSlots(&_Token.CallOpts, _address)
}

// AddressPublicSaleSlots is a free data retrieval call binding the contract method 0x26f8873e.
//
// Solidity: function addressPublicSaleSlots(address _address) view returns(uint8)
func (_Token *TokenCaller) AddressPublicSaleSlots(opts *bind.CallOpts, _address common.Address) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "addressPublicSaleSlots", _address)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AddressPublicSaleSlots is a free data retrieval call binding the contract method 0x26f8873e.
//
// Solidity: function addressPublicSaleSlots(address _address) view returns(uint8)
func (_Token *TokenSession) AddressPublicSaleSlots(_address common.Address) (uint8, error) {
	return _Token.Contract.AddressPublicSaleSlots(&_Token.CallOpts, _address)
}

// AddressPublicSaleSlots is a free data retrieval call binding the contract method 0x26f8873e.
//
// Solidity: function addressPublicSaleSlots(address _address) view returns(uint8)
func (_Token *TokenCallerSession) AddressPublicSaleSlots(_address common.Address) (uint8, error) {
	return _Token.Contract.AddressPublicSaleSlots(&_Token.CallOpts, _address)
}

// AvailableForMint is a free data retrieval call binding the contract method 0x6204b831.
//
// Solidity: function availableForMint() view returns(uint256)
func (_Token *TokenCaller) AvailableForMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "availableForMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableForMint is a free data retrieval call binding the contract method 0x6204b831.
//
// Solidity: function availableForMint() view returns(uint256)
func (_Token *TokenSession) AvailableForMint() (*big.Int, error) {
	return _Token.Contract.AvailableForMint(&_Token.CallOpts)
}

// AvailableForMint is a free data retrieval call binding the contract method 0x6204b831.
//
// Solidity: function availableForMint() view returns(uint256)
func (_Token *TokenCallerSession) AvailableForMint() (*big.Int, error) {
	return _Token.Contract.AvailableForMint(&_Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.GetApproved(&_Token.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Token *TokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.GetApproved(&_Token.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Token *TokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Token.Contract.IsApprovedForAll(&_Token.CallOpts, owner, operator)
}

// IsFreeze is a free data retrieval call binding the contract method 0xc154d6ca.
//
// Solidity: function isFreeze() view returns(bool)
func (_Token *TokenCaller) IsFreeze(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isFreeze")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFreeze is a free data retrieval call binding the contract method 0xc154d6ca.
//
// Solidity: function isFreeze() view returns(bool)
func (_Token *TokenSession) IsFreeze() (bool, error) {
	return _Token.Contract.IsFreeze(&_Token.CallOpts)
}

// IsFreeze is a free data retrieval call binding the contract method 0xc154d6ca.
//
// Solidity: function isFreeze() view returns(bool)
func (_Token *TokenCallerSession) IsFreeze() (bool, error) {
	return _Token.Contract.IsFreeze(&_Token.CallOpts)
}

// IsPrivateSaleActive is a free data retrieval call binding the contract method 0xd8a9e1e9.
//
// Solidity: function isPrivateSaleActive() view returns(bool)
func (_Token *TokenCaller) IsPrivateSaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isPrivateSaleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrivateSaleActive is a free data retrieval call binding the contract method 0xd8a9e1e9.
//
// Solidity: function isPrivateSaleActive() view returns(bool)
func (_Token *TokenSession) IsPrivateSaleActive() (bool, error) {
	return _Token.Contract.IsPrivateSaleActive(&_Token.CallOpts)
}

// IsPrivateSaleActive is a free data retrieval call binding the contract method 0xd8a9e1e9.
//
// Solidity: function isPrivateSaleActive() view returns(bool)
func (_Token *TokenCallerSession) IsPrivateSaleActive() (bool, error) {
	return _Token.Contract.IsPrivateSaleActive(&_Token.CallOpts)
}

// IsPublicSaleActive is a free data retrieval call binding the contract method 0x1e84c413.
//
// Solidity: function isPublicSaleActive() view returns(bool)
func (_Token *TokenCaller) IsPublicSaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isPublicSaleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSaleActive is a free data retrieval call binding the contract method 0x1e84c413.
//
// Solidity: function isPublicSaleActive() view returns(bool)
func (_Token *TokenSession) IsPublicSaleActive() (bool, error) {
	return _Token.Contract.IsPublicSaleActive(&_Token.CallOpts)
}

// IsPublicSaleActive is a free data retrieval call binding the contract method 0x1e84c413.
//
// Solidity: function isPublicSaleActive() view returns(bool)
func (_Token *TokenCallerSession) IsPublicSaleActive() (bool, error) {
	return _Token.Contract.IsPublicSaleActive(&_Token.CallOpts)
}

// IsPublicSaleEnded is a free data retrieval call binding the contract method 0xeeca3737.
//
// Solidity: function isPublicSaleEnded() view returns(bool)
func (_Token *TokenCaller) IsPublicSaleEnded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isPublicSaleEnded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSaleEnded is a free data retrieval call binding the contract method 0xeeca3737.
//
// Solidity: function isPublicSaleEnded() view returns(bool)
func (_Token *TokenSession) IsPublicSaleEnded() (bool, error) {
	return _Token.Contract.IsPublicSaleEnded(&_Token.CallOpts)
}

// IsPublicSaleEnded is a free data retrieval call binding the contract method 0xeeca3737.
//
// Solidity: function isPublicSaleEnded() view returns(bool)
func (_Token *TokenCallerSession) IsPublicSaleEnded() (bool, error) {
	return _Token.Contract.IsPublicSaleEnded(&_Token.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Token *TokenCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isWhitelisted", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Token *TokenSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Token.Contract.IsWhitelisted(&_Token.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_Token *TokenCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Token.Contract.IsWhitelisted(&_Token.CallOpts, _address)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Token *TokenCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Token *TokenSession) MaxTotalSupply() (*big.Int, error) {
	return _Token.Contract.MaxTotalSupply(&_Token.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Token *TokenCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Token.Contract.MaxTotalSupply(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Token *TokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Token.Contract.OwnerOf(&_Token.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Token *TokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Token.Contract.SupportsInterface(&_Token.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Token *TokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Token *TokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Token.Contract.TokenByIndex(&_Token.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Token *TokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Token.Contract.TokenByIndex(&_Token.CallOpts, index)
}

// TokenColor is a free data retrieval call binding the contract method 0x58ac8eef.
//
// Solidity: function tokenColor(uint256 tokenId) view returns(uint8)
func (_Token *TokenCaller) TokenColor(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokenColor", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenColor is a free data retrieval call binding the contract method 0x58ac8eef.
//
// Solidity: function tokenColor(uint256 tokenId) view returns(uint8)
func (_Token *TokenSession) TokenColor(tokenId *big.Int) (uint8, error) {
	return _Token.Contract.TokenColor(&_Token.CallOpts, tokenId)
}

// TokenColor is a free data retrieval call binding the contract method 0x58ac8eef.
//
// Solidity: function tokenColor(uint256 tokenId) view returns(uint8)
func (_Token *TokenCallerSession) TokenColor(tokenId *big.Int) (uint8, error) {
	return _Token.Contract.TokenColor(&_Token.CallOpts, tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Token *TokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Token *TokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Token.Contract.TokenOfOwnerByIndex(&_Token.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Token *TokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Token.Contract.TokenOfOwnerByIndex(&_Token.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Token.Contract.TokenURI(&_Token.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Token *TokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Token.Contract.TokenURI(&_Token.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// AddWhitelistAddresses is a paid mutator transaction binding the contract method 0xe71159b2.
//
// Solidity: function addWhitelistAddresses(uint8 numberOfTokens, address[] addresses) returns()
func (_Token *TokenTransactor) AddWhitelistAddresses(opts *bind.TransactOpts, numberOfTokens uint8, addresses []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addWhitelistAddresses", numberOfTokens, addresses)
}

// AddWhitelistAddresses is a paid mutator transaction binding the contract method 0xe71159b2.
//
// Solidity: function addWhitelistAddresses(uint8 numberOfTokens, address[] addresses) returns()
func (_Token *TokenSession) AddWhitelistAddresses(numberOfTokens uint8, addresses []common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddWhitelistAddresses(&_Token.TransactOpts, numberOfTokens, addresses)
}

// AddWhitelistAddresses is a paid mutator transaction binding the contract method 0xe71159b2.
//
// Solidity: function addWhitelistAddresses(uint8 numberOfTokens, address[] addresses) returns()
func (_Token *TokenTransactorSession) AddWhitelistAddresses(numberOfTokens uint8, addresses []common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddWhitelistAddresses(&_Token.TransactOpts, numberOfTokens, addresses)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token *TokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, to, tokenId)
}

// FlipPrivateSaleState is a paid mutator transaction binding the contract method 0x3719e3b0.
//
// Solidity: function flipPrivateSaleState() returns()
func (_Token *TokenTransactor) FlipPrivateSaleState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "flipPrivateSaleState")
}

// FlipPrivateSaleState is a paid mutator transaction binding the contract method 0x3719e3b0.
//
// Solidity: function flipPrivateSaleState() returns()
func (_Token *TokenSession) FlipPrivateSaleState() (*types.Transaction, error) {
	return _Token.Contract.FlipPrivateSaleState(&_Token.TransactOpts)
}

// FlipPrivateSaleState is a paid mutator transaction binding the contract method 0x3719e3b0.
//
// Solidity: function flipPrivateSaleState() returns()
func (_Token *TokenTransactorSession) FlipPrivateSaleState() (*types.Transaction, error) {
	return _Token.Contract.FlipPrivateSaleState(&_Token.TransactOpts)
}

// FlipPublicSaleState is a paid mutator transaction binding the contract method 0xa10866ef.
//
// Solidity: function flipPublicSaleState() returns()
func (_Token *TokenTransactor) FlipPublicSaleState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "flipPublicSaleState")
}

// FlipPublicSaleState is a paid mutator transaction binding the contract method 0xa10866ef.
//
// Solidity: function flipPublicSaleState() returns()
func (_Token *TokenSession) FlipPublicSaleState() (*types.Transaction, error) {
	return _Token.Contract.FlipPublicSaleState(&_Token.TransactOpts)
}

// FlipPublicSaleState is a paid mutator transaction binding the contract method 0xa10866ef.
//
// Solidity: function flipPublicSaleState() returns()
func (_Token *TokenTransactorSession) FlipPublicSaleState() (*types.Transaction, error) {
	return _Token.Contract.FlipPublicSaleState(&_Token.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_Token *TokenTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "freeze")
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_Token *TokenSession) Freeze() (*types.Transaction, error) {
	return _Token.Contract.Freeze(&_Token.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_Token *TokenTransactorSession) Freeze() (*types.Transaction, error) {
	return _Token.Contract.Freeze(&_Token.TransactOpts)
}

// MintGiftToken is a paid mutator transaction binding the contract method 0x22f8cedf.
//
// Solidity: function mintGiftToken(address to, string _tokenURI, uint8 color) returns()
func (_Token *TokenTransactor) MintGiftToken(opts *bind.TransactOpts, to common.Address, _tokenURI string, color uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintGiftToken", to, _tokenURI, color)
}

// MintGiftToken is a paid mutator transaction binding the contract method 0x22f8cedf.
//
// Solidity: function mintGiftToken(address to, string _tokenURI, uint8 color) returns()
func (_Token *TokenSession) MintGiftToken(to common.Address, _tokenURI string, color uint8) (*types.Transaction, error) {
	return _Token.Contract.MintGiftToken(&_Token.TransactOpts, to, _tokenURI, color)
}

// MintGiftToken is a paid mutator transaction binding the contract method 0x22f8cedf.
//
// Solidity: function mintGiftToken(address to, string _tokenURI, uint8 color) returns()
func (_Token *TokenTransactorSession) MintGiftToken(to common.Address, _tokenURI string, color uint8) (*types.Transaction, error) {
	return _Token.Contract.MintGiftToken(&_Token.TransactOpts, to, _tokenURI, color)
}

// MintPrivate is a paid mutator transaction binding the contract method 0x4da79d5a.
//
// Solidity: function mintPrivate(uint8 numberOfTokens, uint8[] colors) payable returns()
func (_Token *TokenTransactor) MintPrivate(opts *bind.TransactOpts, numberOfTokens uint8, colors []uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintPrivate", numberOfTokens, colors)
}

// MintPrivate is a paid mutator transaction binding the contract method 0x4da79d5a.
//
// Solidity: function mintPrivate(uint8 numberOfTokens, uint8[] colors) payable returns()
func (_Token *TokenSession) MintPrivate(numberOfTokens uint8, colors []uint8) (*types.Transaction, error) {
	return _Token.Contract.MintPrivate(&_Token.TransactOpts, numberOfTokens, colors)
}

// MintPrivate is a paid mutator transaction binding the contract method 0x4da79d5a.
//
// Solidity: function mintPrivate(uint8 numberOfTokens, uint8[] colors) payable returns()
func (_Token *TokenTransactorSession) MintPrivate(numberOfTokens uint8, colors []uint8) (*types.Transaction, error) {
	return _Token.Contract.MintPrivate(&_Token.TransactOpts, numberOfTokens, colors)
}

// MintPublic is a paid mutator transaction binding the contract method 0xfcc777f5.
//
// Solidity: function mintPublic(uint8 numberOfTokens, uint8[] colors) payable returns()
func (_Token *TokenTransactor) MintPublic(opts *bind.TransactOpts, numberOfTokens uint8, colors []uint8) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintPublic", numberOfTokens, colors)
}

// MintPublic is a paid mutator transaction binding the contract method 0xfcc777f5.
//
// Solidity: function mintPublic(uint8 numberOfTokens, uint8[] colors) payable returns()
func (_Token *TokenSession) MintPublic(numberOfTokens uint8, colors []uint8) (*types.Transaction, error) {
	return _Token.Contract.MintPublic(&_Token.TransactOpts, numberOfTokens, colors)
}

// MintPublic is a paid mutator transaction binding the contract method 0xfcc777f5.
//
// Solidity: function mintPublic(uint8 numberOfTokens, uint8[] colors) payable returns()
func (_Token *TokenTransactorSession) MintPublic(numberOfTokens uint8, colors []uint8) (*types.Transaction, error) {
	return _Token.Contract.MintPublic(&_Token.TransactOpts, numberOfTokens, colors)
}

// RemoveWhitelistAddresses is a paid mutator transaction binding the contract method 0xb83921a6.
//
// Solidity: function removeWhitelistAddresses(address[] addresses) returns()
func (_Token *TokenTransactor) RemoveWhitelistAddresses(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeWhitelistAddresses", addresses)
}

// RemoveWhitelistAddresses is a paid mutator transaction binding the contract method 0xb83921a6.
//
// Solidity: function removeWhitelistAddresses(address[] addresses) returns()
func (_Token *TokenSession) RemoveWhitelistAddresses(addresses []common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveWhitelistAddresses(&_Token.TransactOpts, addresses)
}

// RemoveWhitelistAddresses is a paid mutator transaction binding the contract method 0xb83921a6.
//
// Solidity: function removeWhitelistAddresses(address[] addresses) returns()
func (_Token *TokenTransactorSession) RemoveWhitelistAddresses(addresses []common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveWhitelistAddresses(&_Token.TransactOpts, addresses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Token *TokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Token *TokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom0(&_Token.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Token *TokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Token.Contract.SafeTransferFrom0(&_Token.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Token *TokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Token.Contract.SetApprovalForAll(&_Token.TransactOpts, operator, approved)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Token *TokenTransactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTokenURI", tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Token *TokenSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Token.Contract.SetTokenURI(&_Token.TransactOpts, tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Token *TokenTransactorSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Token.Contract.SetTokenURI(&_Token.TransactOpts, tokenId, _tokenURI)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Token *TokenTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Token *TokenSession) Withdraw() (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Token *TokenTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Token contract.
type TokenApprovalForAllIterator struct {
	Event *TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApprovalForAll)
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
		it.Event = new(TokenApprovalForAll)
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
func (it *TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApprovalForAll represents a ApprovalForAll event raised by the Token contract.
type TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalForAllIterator{contract: _Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApprovalForAll)
				if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Token *TokenFilterer) ParseApprovalForAll(log types.Log) (*TokenApprovalForAll, error) {
	event := new(TokenApprovalForAll)
	if err := _Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAddressPrivateSlotsChangeIterator is returned from FilterAddressPrivateSlotsChange and is used to iterate over the raw logs and unpacked data for AddressPrivateSlotsChange events raised by the Token contract.
type TokenAddressPrivateSlotsChangeIterator struct {
	Event *TokenAddressPrivateSlotsChange // Event containing the contract specifics and raw log

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
func (it *TokenAddressPrivateSlotsChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAddressPrivateSlotsChange)
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
		it.Event = new(TokenAddressPrivateSlotsChange)
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
func (it *TokenAddressPrivateSlotsChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAddressPrivateSlotsChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAddressPrivateSlotsChange represents a AddressPrivateSlotsChange event raised by the Token contract.
type TokenAddressPrivateSlotsChange struct {
	Addr  common.Address
	Slots *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddressPrivateSlotsChange is a free log retrieval operation binding the contract event 0xe7d264d696ff341e3e6a10c22579a979a96db58bda028ef4dca464bfbe6a1f08.
//
// Solidity: event addressPrivateSlotsChange(address addr, uint256 slots)
func (_Token *TokenFilterer) FilterAddressPrivateSlotsChange(opts *bind.FilterOpts) (*TokenAddressPrivateSlotsChangeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "addressPrivateSlotsChange")
	if err != nil {
		return nil, err
	}
	return &TokenAddressPrivateSlotsChangeIterator{contract: _Token.contract, event: "addressPrivateSlotsChange", logs: logs, sub: sub}, nil
}

// WatchAddressPrivateSlotsChange is a free log subscription operation binding the contract event 0xe7d264d696ff341e3e6a10c22579a979a96db58bda028ef4dca464bfbe6a1f08.
//
// Solidity: event addressPrivateSlotsChange(address addr, uint256 slots)
func (_Token *TokenFilterer) WatchAddressPrivateSlotsChange(opts *bind.WatchOpts, sink chan<- *TokenAddressPrivateSlotsChange) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "addressPrivateSlotsChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAddressPrivateSlotsChange)
				if err := _Token.contract.UnpackLog(event, "addressPrivateSlotsChange", log); err != nil {
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

// ParseAddressPrivateSlotsChange is a log parse operation binding the contract event 0xe7d264d696ff341e3e6a10c22579a979a96db58bda028ef4dca464bfbe6a1f08.
//
// Solidity: event addressPrivateSlotsChange(address addr, uint256 slots)
func (_Token *TokenFilterer) ParseAddressPrivateSlotsChange(log types.Log) (*TokenAddressPrivateSlotsChange, error) {
	event := new(TokenAddressPrivateSlotsChange)
	if err := _Token.contract.UnpackLog(event, "addressPrivateSlotsChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAddressPublicSlotsChangeIterator is returned from FilterAddressPublicSlotsChange and is used to iterate over the raw logs and unpacked data for AddressPublicSlotsChange events raised by the Token contract.
type TokenAddressPublicSlotsChangeIterator struct {
	Event *TokenAddressPublicSlotsChange // Event containing the contract specifics and raw log

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
func (it *TokenAddressPublicSlotsChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAddressPublicSlotsChange)
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
		it.Event = new(TokenAddressPublicSlotsChange)
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
func (it *TokenAddressPublicSlotsChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAddressPublicSlotsChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAddressPublicSlotsChange represents a AddressPublicSlotsChange event raised by the Token contract.
type TokenAddressPublicSlotsChange struct {
	Addr           common.Address
	Slots          *big.Int
	TotalRemaining *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressPublicSlotsChange is a free log retrieval operation binding the contract event 0xa21cede0a934fc8df5d3330132f2a959fbe0d70b901e2a17909bb1e1b8feb904.
//
// Solidity: event addressPublicSlotsChange(address addr, uint256 slots, uint256 totalRemaining)
func (_Token *TokenFilterer) FilterAddressPublicSlotsChange(opts *bind.FilterOpts) (*TokenAddressPublicSlotsChangeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "addressPublicSlotsChange")
	if err != nil {
		return nil, err
	}
	return &TokenAddressPublicSlotsChangeIterator{contract: _Token.contract, event: "addressPublicSlotsChange", logs: logs, sub: sub}, nil
}

// WatchAddressPublicSlotsChange is a free log subscription operation binding the contract event 0xa21cede0a934fc8df5d3330132f2a959fbe0d70b901e2a17909bb1e1b8feb904.
//
// Solidity: event addressPublicSlotsChange(address addr, uint256 slots, uint256 totalRemaining)
func (_Token *TokenFilterer) WatchAddressPublicSlotsChange(opts *bind.WatchOpts, sink chan<- *TokenAddressPublicSlotsChange) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "addressPublicSlotsChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAddressPublicSlotsChange)
				if err := _Token.contract.UnpackLog(event, "addressPublicSlotsChange", log); err != nil {
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

// ParseAddressPublicSlotsChange is a log parse operation binding the contract event 0xa21cede0a934fc8df5d3330132f2a959fbe0d70b901e2a17909bb1e1b8feb904.
//
// Solidity: event addressPublicSlotsChange(address addr, uint256 slots, uint256 totalRemaining)
func (_Token *TokenFilterer) ParseAddressPublicSlotsChange(log types.Log) (*TokenAddressPublicSlotsChange, error) {
	event := new(TokenAddressPublicSlotsChange)
	if err := _Token.contract.UnpackLog(event, "addressPublicSlotsChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPrivateSaleStateIterator is returned from FilterPrivateSaleState and is used to iterate over the raw logs and unpacked data for PrivateSaleState events raised by the Token contract.
type TokenPrivateSaleStateIterator struct {
	Event *TokenPrivateSaleState // Event containing the contract specifics and raw log

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
func (it *TokenPrivateSaleStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPrivateSaleState)
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
		it.Event = new(TokenPrivateSaleState)
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
func (it *TokenPrivateSaleStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPrivateSaleStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPrivateSaleState represents a PrivateSaleState event raised by the Token contract.
type TokenPrivateSaleState struct {
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPrivateSaleState is a free log retrieval operation binding the contract event 0xc08856a2421028bf2a88be7559541c7087eb189d1518a79356b4ace74b5252c9.
//
// Solidity: event privateSaleState(bool active)
func (_Token *TokenFilterer) FilterPrivateSaleState(opts *bind.FilterOpts) (*TokenPrivateSaleStateIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "privateSaleState")
	if err != nil {
		return nil, err
	}
	return &TokenPrivateSaleStateIterator{contract: _Token.contract, event: "privateSaleState", logs: logs, sub: sub}, nil
}

// WatchPrivateSaleState is a free log subscription operation binding the contract event 0xc08856a2421028bf2a88be7559541c7087eb189d1518a79356b4ace74b5252c9.
//
// Solidity: event privateSaleState(bool active)
func (_Token *TokenFilterer) WatchPrivateSaleState(opts *bind.WatchOpts, sink chan<- *TokenPrivateSaleState) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "privateSaleState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPrivateSaleState)
				if err := _Token.contract.UnpackLog(event, "privateSaleState", log); err != nil {
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

// ParsePrivateSaleState is a log parse operation binding the contract event 0xc08856a2421028bf2a88be7559541c7087eb189d1518a79356b4ace74b5252c9.
//
// Solidity: event privateSaleState(bool active)
func (_Token *TokenFilterer) ParsePrivateSaleState(log types.Log) (*TokenPrivateSaleState, error) {
	event := new(TokenPrivateSaleState)
	if err := _Token.contract.UnpackLog(event, "privateSaleState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPublicSaleStateIterator is returned from FilterPublicSaleState and is used to iterate over the raw logs and unpacked data for PublicSaleState events raised by the Token contract.
type TokenPublicSaleStateIterator struct {
	Event *TokenPublicSaleState // Event containing the contract specifics and raw log

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
func (it *TokenPublicSaleStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPublicSaleState)
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
		it.Event = new(TokenPublicSaleState)
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
func (it *TokenPublicSaleStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPublicSaleStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPublicSaleState represents a PublicSaleState event raised by the Token contract.
type TokenPublicSaleState struct {
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPublicSaleState is a free log retrieval operation binding the contract event 0xd9bc49251edd474a55897591c0b750b4f5b2609a950c788f8ef8d44b45b0f718.
//
// Solidity: event publicSaleState(bool active)
func (_Token *TokenFilterer) FilterPublicSaleState(opts *bind.FilterOpts) (*TokenPublicSaleStateIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "publicSaleState")
	if err != nil {
		return nil, err
	}
	return &TokenPublicSaleStateIterator{contract: _Token.contract, event: "publicSaleState", logs: logs, sub: sub}, nil
}

// WatchPublicSaleState is a free log subscription operation binding the contract event 0xd9bc49251edd474a55897591c0b750b4f5b2609a950c788f8ef8d44b45b0f718.
//
// Solidity: event publicSaleState(bool active)
func (_Token *TokenFilterer) WatchPublicSaleState(opts *bind.WatchOpts, sink chan<- *TokenPublicSaleState) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "publicSaleState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPublicSaleState)
				if err := _Token.contract.UnpackLog(event, "publicSaleState", log); err != nil {
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

// ParsePublicSaleState is a log parse operation binding the contract event 0xd9bc49251edd474a55897591c0b750b4f5b2609a950c788f8ef8d44b45b0f718.
//
// Solidity: event publicSaleState(bool active)
func (_Token *TokenFilterer) ParsePublicSaleState(log types.Log) (*TokenPublicSaleState, error) {
	event := new(TokenPublicSaleState)
	if err := _Token.contract.UnpackLog(event, "publicSaleState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
