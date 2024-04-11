// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// HoraceNFTMarketListing is an auto generated low-level Go binding around an user-defined struct.
type HoraceNFTMarketListing struct {
	Price  *big.Int
	Seller common.Address
}

// HoraceNFTMarketMetaData contains all meta data concerning the HoraceNFTMarket contract.
var HoraceNFTMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ErrorNFTHasListed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ErrorNFTHasNotListed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ErrorNFTInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorNFTInvalidPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ErrorNFTNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorNoProceed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BuyNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"CancelListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"old_price\",\"type\":\"uint256\"}],\"name\":\"NftListed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"internalType\":\"structHoraceNFTMarket.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getProceeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"listNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withDrawProceeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HoraceNFTMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use HoraceNFTMarketMetaData.ABI instead.
var HoraceNFTMarketABI = HoraceNFTMarketMetaData.ABI

// HoraceNFTMarket is an auto generated Go binding around an Ethereum contract.
type HoraceNFTMarket struct {
	HoraceNFTMarketCaller     // Read-only binding to the contract
	HoraceNFTMarketTransactor // Write-only binding to the contract
	HoraceNFTMarketFilterer   // Log filterer for contract events
}

// HoraceNFTMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoraceNFTMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoraceNFTMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoraceNFTMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoraceNFTMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoraceNFTMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoraceNFTMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoraceNFTMarketSession struct {
	Contract     *HoraceNFTMarket  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoraceNFTMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoraceNFTMarketCallerSession struct {
	Contract *HoraceNFTMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HoraceNFTMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoraceNFTMarketTransactorSession struct {
	Contract     *HoraceNFTMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HoraceNFTMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoraceNFTMarketRaw struct {
	Contract *HoraceNFTMarket // Generic contract binding to access the raw methods on
}

// HoraceNFTMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoraceNFTMarketCallerRaw struct {
	Contract *HoraceNFTMarketCaller // Generic read-only contract binding to access the raw methods on
}

// HoraceNFTMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoraceNFTMarketTransactorRaw struct {
	Contract *HoraceNFTMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoraceNFTMarket creates a new instance of HoraceNFTMarket, bound to a specific deployed contract.
func NewHoraceNFTMarket(address common.Address, backend bind.ContractBackend) (*HoraceNFTMarket, error) {
	contract, err := bindHoraceNFTMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoraceNFTMarket{HoraceNFTMarketCaller: HoraceNFTMarketCaller{contract: contract}, HoraceNFTMarketTransactor: HoraceNFTMarketTransactor{contract: contract}, HoraceNFTMarketFilterer: HoraceNFTMarketFilterer{contract: contract}}, nil
}

// NewHoraceNFTMarketCaller creates a new read-only instance of HoraceNFTMarket, bound to a specific deployed contract.
func NewHoraceNFTMarketCaller(address common.Address, caller bind.ContractCaller) (*HoraceNFTMarketCaller, error) {
	contract, err := bindHoraceNFTMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoraceNFTMarketCaller{contract: contract}, nil
}

// NewHoraceNFTMarketTransactor creates a new write-only instance of HoraceNFTMarket, bound to a specific deployed contract.
func NewHoraceNFTMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*HoraceNFTMarketTransactor, error) {
	contract, err := bindHoraceNFTMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoraceNFTMarketTransactor{contract: contract}, nil
}

// NewHoraceNFTMarketFilterer creates a new log filterer instance of HoraceNFTMarket, bound to a specific deployed contract.
func NewHoraceNFTMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*HoraceNFTMarketFilterer, error) {
	contract, err := bindHoraceNFTMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoraceNFTMarketFilterer{contract: contract}, nil
}

// bindHoraceNFTMarket binds a generic wrapper to an already deployed contract.
func bindHoraceNFTMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HoraceNFTMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoraceNFTMarket *HoraceNFTMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoraceNFTMarket.Contract.HoraceNFTMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoraceNFTMarket *HoraceNFTMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.HoraceNFTMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoraceNFTMarket *HoraceNFTMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.HoraceNFTMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoraceNFTMarket *HoraceNFTMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoraceNFTMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoraceNFTMarket *HoraceNFTMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoraceNFTMarket *HoraceNFTMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_HoraceNFTMarket *HoraceNFTMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HoraceNFTMarket.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_HoraceNFTMarket *HoraceNFTMarketSession) Owner() (common.Address, error) {
	return _HoraceNFTMarket.Contract.Owner(&_HoraceNFTMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_HoraceNFTMarket *HoraceNFTMarketCallerSession) Owner() (common.Address, error) {
	return _HoraceNFTMarket.Contract.Owner(&_HoraceNFTMarket.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x88700d1c.
//
// Solidity: function getListing(address nftAddress, uint256 tokenId) view returns((uint256,address))
func (_HoraceNFTMarket *HoraceNFTMarketCaller) GetListing(opts *bind.CallOpts, nftAddress common.Address, tokenId *big.Int) (HoraceNFTMarketListing, error) {
	var out []interface{}
	err := _HoraceNFTMarket.contract.Call(opts, &out, "getListing", nftAddress, tokenId)

	if err != nil {
		return *new(HoraceNFTMarketListing), err
	}

	out0 := *abi.ConvertType(out[0], new(HoraceNFTMarketListing)).(*HoraceNFTMarketListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x88700d1c.
//
// Solidity: function getListing(address nftAddress, uint256 tokenId) view returns((uint256,address))
func (_HoraceNFTMarket *HoraceNFTMarketSession) GetListing(nftAddress common.Address, tokenId *big.Int) (HoraceNFTMarketListing, error) {
	return _HoraceNFTMarket.Contract.GetListing(&_HoraceNFTMarket.CallOpts, nftAddress, tokenId)
}

// GetListing is a free data retrieval call binding the contract method 0x88700d1c.
//
// Solidity: function getListing(address nftAddress, uint256 tokenId) view returns((uint256,address))
func (_HoraceNFTMarket *HoraceNFTMarketCallerSession) GetListing(nftAddress common.Address, tokenId *big.Int) (HoraceNFTMarketListing, error) {
	return _HoraceNFTMarket.Contract.GetListing(&_HoraceNFTMarket.CallOpts, nftAddress, tokenId)
}

// GetProceeds is a free data retrieval call binding the contract method 0x30cfeda9.
//
// Solidity: function getProceeds(address seller) view returns(uint256)
func (_HoraceNFTMarket *HoraceNFTMarketCaller) GetProceeds(opts *bind.CallOpts, seller common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HoraceNFTMarket.contract.Call(opts, &out, "getProceeds", seller)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProceeds is a free data retrieval call binding the contract method 0x30cfeda9.
//
// Solidity: function getProceeds(address seller) view returns(uint256)
func (_HoraceNFTMarket *HoraceNFTMarketSession) GetProceeds(seller common.Address) (*big.Int, error) {
	return _HoraceNFTMarket.Contract.GetProceeds(&_HoraceNFTMarket.CallOpts, seller)
}

// GetProceeds is a free data retrieval call binding the contract method 0x30cfeda9.
//
// Solidity: function getProceeds(address seller) view returns(uint256)
func (_HoraceNFTMarket *HoraceNFTMarketCallerSession) GetProceeds(seller common.Address) (*big.Int, error) {
	return _HoraceNFTMarket.Contract.GetProceeds(&_HoraceNFTMarket.CallOpts, seller)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xa82ba76f.
//
// Solidity: function buyNFT(address nftAddress, uint256 tokenId) payable returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactor) BuyNFT(opts *bind.TransactOpts, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.contract.Transact(opts, "buyNFT", nftAddress, tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xa82ba76f.
//
// Solidity: function buyNFT(address nftAddress, uint256 tokenId) payable returns()
func (_HoraceNFTMarket *HoraceNFTMarketSession) BuyNFT(nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.BuyNFT(&_HoraceNFTMarket.TransactOpts, nftAddress, tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xa82ba76f.
//
// Solidity: function buyNFT(address nftAddress, uint256 tokenId) payable returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactorSession) BuyNFT(nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.BuyNFT(&_HoraceNFTMarket.TransactOpts, nftAddress, tokenId)
}

// CancelList is a paid mutator transaction binding the contract method 0x1af71a8e.
//
// Solidity: function cancelList(address nftAddress, uint256 tokenId) returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactor) CancelList(opts *bind.TransactOpts, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.contract.Transact(opts, "cancelList", nftAddress, tokenId)
}

// CancelList is a paid mutator transaction binding the contract method 0x1af71a8e.
//
// Solidity: function cancelList(address nftAddress, uint256 tokenId) returns()
func (_HoraceNFTMarket *HoraceNFTMarketSession) CancelList(nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.CancelList(&_HoraceNFTMarket.TransactOpts, nftAddress, tokenId)
}

// CancelList is a paid mutator transaction binding the contract method 0x1af71a8e.
//
// Solidity: function cancelList(address nftAddress, uint256 tokenId) returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactorSession) CancelList(nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.CancelList(&_HoraceNFTMarket.TransactOpts, nftAddress, tokenId)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address nftAddress, uint256 price, uint256 tokenId) returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactor) ListNFT(opts *bind.TransactOpts, nftAddress common.Address, price *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.contract.Transact(opts, "listNFT", nftAddress, price, tokenId)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address nftAddress, uint256 price, uint256 tokenId) returns()
func (_HoraceNFTMarket *HoraceNFTMarketSession) ListNFT(nftAddress common.Address, price *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.ListNFT(&_HoraceNFTMarket.TransactOpts, nftAddress, price, tokenId)
}

// ListNFT is a paid mutator transaction binding the contract method 0xad05f1b4.
//
// Solidity: function listNFT(address nftAddress, uint256 price, uint256 tokenId) returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactorSession) ListNFT(nftAddress common.Address, price *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.ListNFT(&_HoraceNFTMarket.TransactOpts, nftAddress, price, tokenId)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xf772adf1.
//
// Solidity: function updateListing(address nftAddress, uint256 tokenId, uint256 newPrice) returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactor) UpdateListing(opts *bind.TransactOpts, nftAddress common.Address, tokenId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.contract.Transact(opts, "updateListing", nftAddress, tokenId, newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xf772adf1.
//
// Solidity: function updateListing(address nftAddress, uint256 tokenId, uint256 newPrice) returns()
func (_HoraceNFTMarket *HoraceNFTMarketSession) UpdateListing(nftAddress common.Address, tokenId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.UpdateListing(&_HoraceNFTMarket.TransactOpts, nftAddress, tokenId, newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xf772adf1.
//
// Solidity: function updateListing(address nftAddress, uint256 tokenId, uint256 newPrice) returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactorSession) UpdateListing(nftAddress common.Address, tokenId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.UpdateListing(&_HoraceNFTMarket.TransactOpts, nftAddress, tokenId, newPrice)
}

// WithDrawProceeds is a paid mutator transaction binding the contract method 0xc6890de5.
//
// Solidity: function withDrawProceeds() returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactor) WithDrawProceeds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoraceNFTMarket.contract.Transact(opts, "withDrawProceeds")
}

// WithDrawProceeds is a paid mutator transaction binding the contract method 0xc6890de5.
//
// Solidity: function withDrawProceeds() returns()
func (_HoraceNFTMarket *HoraceNFTMarketSession) WithDrawProceeds() (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.WithDrawProceeds(&_HoraceNFTMarket.TransactOpts)
}

// WithDrawProceeds is a paid mutator transaction binding the contract method 0xc6890de5.
//
// Solidity: function withDrawProceeds() returns()
func (_HoraceNFTMarket *HoraceNFTMarketTransactorSession) WithDrawProceeds() (*types.Transaction, error) {
	return _HoraceNFTMarket.Contract.WithDrawProceeds(&_HoraceNFTMarket.TransactOpts)
}

// HoraceNFTMarketBuyNFTIterator is returned from FilterBuyNFT and is used to iterate over the raw logs and unpacked data for BuyNFT events raised by the HoraceNFTMarket contract.
type HoraceNFTMarketBuyNFTIterator struct {
	Event *HoraceNFTMarketBuyNFT // Event containing the contract specifics and raw log

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
func (it *HoraceNFTMarketBuyNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoraceNFTMarketBuyNFT)
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
		it.Event = new(HoraceNFTMarketBuyNFT)
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
func (it *HoraceNFTMarketBuyNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoraceNFTMarketBuyNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoraceNFTMarketBuyNFT represents a BuyNFT event raised by the HoraceNFTMarket contract.
type HoraceNFTMarketBuyNFT struct {
	Buyer      common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBuyNFT is a free log retrieval operation binding the contract event 0xe35677743c6b7abf01cca38cf0f4b986386e50105ec5fd3f42daeea57acf57a9.
//
// Solidity: event BuyNFT(address indexed buyer, address indexed nftAddress, uint256 indexed tokenId, uint256 price)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) FilterBuyNFT(opts *bind.FilterOpts, buyer []common.Address, nftAddress []common.Address, tokenId []*big.Int) (*HoraceNFTMarketBuyNFTIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HoraceNFTMarket.contract.FilterLogs(opts, "BuyNFT", buyerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HoraceNFTMarketBuyNFTIterator{contract: _HoraceNFTMarket.contract, event: "BuyNFT", logs: logs, sub: sub}, nil
}

// WatchBuyNFT is a free log subscription operation binding the contract event 0xe35677743c6b7abf01cca38cf0f4b986386e50105ec5fd3f42daeea57acf57a9.
//
// Solidity: event BuyNFT(address indexed buyer, address indexed nftAddress, uint256 indexed tokenId, uint256 price)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) WatchBuyNFT(opts *bind.WatchOpts, sink chan<- *HoraceNFTMarketBuyNFT, buyer []common.Address, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HoraceNFTMarket.contract.WatchLogs(opts, "BuyNFT", buyerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoraceNFTMarketBuyNFT)
				if err := _HoraceNFTMarket.contract.UnpackLog(event, "BuyNFT", log); err != nil {
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

// ParseBuyNFT is a log parse operation binding the contract event 0xe35677743c6b7abf01cca38cf0f4b986386e50105ec5fd3f42daeea57acf57a9.
//
// Solidity: event BuyNFT(address indexed buyer, address indexed nftAddress, uint256 indexed tokenId, uint256 price)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) ParseBuyNFT(log types.Log) (*HoraceNFTMarketBuyNFT, error) {
	event := new(HoraceNFTMarketBuyNFT)
	if err := _HoraceNFTMarket.contract.UnpackLog(event, "BuyNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HoraceNFTMarketCancelListingIterator is returned from FilterCancelListing and is used to iterate over the raw logs and unpacked data for CancelListing events raised by the HoraceNFTMarket contract.
type HoraceNFTMarketCancelListingIterator struct {
	Event *HoraceNFTMarketCancelListing // Event containing the contract specifics and raw log

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
func (it *HoraceNFTMarketCancelListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoraceNFTMarketCancelListing)
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
		it.Event = new(HoraceNFTMarketCancelListing)
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
func (it *HoraceNFTMarketCancelListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoraceNFTMarketCancelListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoraceNFTMarketCancelListing represents a CancelListing event raised by the HoraceNFTMarket contract.
type HoraceNFTMarketCancelListing struct {
	Seller     common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelListing is a free log retrieval operation binding the contract event 0x061e64e10e5e57615a6fe38b48c81926f6999a2f5304d3e5e6b0d3cfac4865a5.
//
// Solidity: event CancelListing(address indexed seller, address indexed nftAddress, uint256 indexed tokenId)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) FilterCancelListing(opts *bind.FilterOpts, seller []common.Address, nftAddress []common.Address, tokenId []*big.Int) (*HoraceNFTMarketCancelListingIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HoraceNFTMarket.contract.FilterLogs(opts, "CancelListing", sellerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HoraceNFTMarketCancelListingIterator{contract: _HoraceNFTMarket.contract, event: "CancelListing", logs: logs, sub: sub}, nil
}

// WatchCancelListing is a free log subscription operation binding the contract event 0x061e64e10e5e57615a6fe38b48c81926f6999a2f5304d3e5e6b0d3cfac4865a5.
//
// Solidity: event CancelListing(address indexed seller, address indexed nftAddress, uint256 indexed tokenId)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) WatchCancelListing(opts *bind.WatchOpts, sink chan<- *HoraceNFTMarketCancelListing, seller []common.Address, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HoraceNFTMarket.contract.WatchLogs(opts, "CancelListing", sellerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoraceNFTMarketCancelListing)
				if err := _HoraceNFTMarket.contract.UnpackLog(event, "CancelListing", log); err != nil {
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

// ParseCancelListing is a log parse operation binding the contract event 0x061e64e10e5e57615a6fe38b48c81926f6999a2f5304d3e5e6b0d3cfac4865a5.
//
// Solidity: event CancelListing(address indexed seller, address indexed nftAddress, uint256 indexed tokenId)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) ParseCancelListing(log types.Log) (*HoraceNFTMarketCancelListing, error) {
	event := new(HoraceNFTMarketCancelListing)
	if err := _HoraceNFTMarket.contract.UnpackLog(event, "CancelListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HoraceNFTMarketNftListedIterator is returned from FilterNftListed and is used to iterate over the raw logs and unpacked data for NftListed events raised by the HoraceNFTMarket contract.
type HoraceNFTMarketNftListedIterator struct {
	Event *HoraceNFTMarketNftListed // Event containing the contract specifics and raw log

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
func (it *HoraceNFTMarketNftListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoraceNFTMarketNftListed)
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
		it.Event = new(HoraceNFTMarketNftListed)
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
func (it *HoraceNFTMarketNftListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoraceNFTMarketNftListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoraceNFTMarketNftListed represents a NftListed event raised by the HoraceNFTMarket contract.
type HoraceNFTMarketNftListed struct {
	Seller     common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Price      *big.Int
	OldPrice   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftListed is a free log retrieval operation binding the contract event 0x9990c878e0fcb8aad5c24109a215405fe250d314b641abe27429eb62c0b41756.
//
// Solidity: event NftListed(address indexed seller, address indexed nftAddress, uint256 indexed tokenId, uint256 price, uint256 old_price)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) FilterNftListed(opts *bind.FilterOpts, seller []common.Address, nftAddress []common.Address, tokenId []*big.Int) (*HoraceNFTMarketNftListedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HoraceNFTMarket.contract.FilterLogs(opts, "NftListed", sellerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HoraceNFTMarketNftListedIterator{contract: _HoraceNFTMarket.contract, event: "NftListed", logs: logs, sub: sub}, nil
}

// WatchNftListed is a free log subscription operation binding the contract event 0x9990c878e0fcb8aad5c24109a215405fe250d314b641abe27429eb62c0b41756.
//
// Solidity: event NftListed(address indexed seller, address indexed nftAddress, uint256 indexed tokenId, uint256 price, uint256 old_price)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) WatchNftListed(opts *bind.WatchOpts, sink chan<- *HoraceNFTMarketNftListed, seller []common.Address, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _HoraceNFTMarket.contract.WatchLogs(opts, "NftListed", sellerRule, nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoraceNFTMarketNftListed)
				if err := _HoraceNFTMarket.contract.UnpackLog(event, "NftListed", log); err != nil {
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

// ParseNftListed is a log parse operation binding the contract event 0x9990c878e0fcb8aad5c24109a215405fe250d314b641abe27429eb62c0b41756.
//
// Solidity: event NftListed(address indexed seller, address indexed nftAddress, uint256 indexed tokenId, uint256 price, uint256 old_price)
func (_HoraceNFTMarket *HoraceNFTMarketFilterer) ParseNftListed(log types.Log) (*HoraceNFTMarketNftListed, error) {
	event := new(HoraceNFTMarketNftListed)
	if err := _HoraceNFTMarket.contract.UnpackLog(event, "NftListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
