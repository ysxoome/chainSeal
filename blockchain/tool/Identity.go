// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tool

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"account\",\"type\":\"bytes\"}],\"name\":\"AddIdentityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"account\",\"type\":\"bytes\"}],\"name\":\"LoseEfficacyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_identity\",\"type\":\"bytes\"}],\"name\":\"addIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"lastEfficacyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_identity\",\"type\":\"bytes\"}],\"name\":\"queryEfficacyTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_identity\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"queryIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_identity\",\"type\":\"bytes\"}],\"name\":\"revokeIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdentityBin is the compiled bytecode used for deploying new contracts.
var IdentityBin = "0x608060405234801561001057600080fd5b50611486806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063a16ce21d11610066578063a16ce21d14610368578063a77fe1a81461043b578063c4d66de81461050e578063e6ddae4f14610552578063f2fde38b1461062f5761009e565b80630a03e738146100a35780631c3f03c9146101765780632871c8a914610245578063715018a6146103145780638da5cb5b1461031e575b600080fd5b61015c600480360360208110156100b957600080fd5b81019080803590602001906401000000008111156100d657600080fd5b8201836020820111156100e857600080fd5b8035906020019184600183028401116401000000008311171561010a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610673565b604051808215151515815260200191505060405180910390f35b61022f6004803603602081101561018c57600080fd5b81019080803590602001906401000000008111156101a957600080fd5b8201836020820111156101bb57600080fd5b803590602001918460018302840111640100000000831117156101dd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506106a9565b6040518082815260200191505060405180910390f35b6102fe6004803603602081101561025b57600080fd5b810190808035906020019064010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061071c565b6040518082815260200191505060405180910390f35b61031c61074a565b005b6103266108d4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104216004803603602081101561037e57600080fd5b810190808035906020019064010000000081111561039b57600080fd5b8201836020820111156103ad57600080fd5b803590602001918460018302840111640100000000831117156103cf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506108fd565b604051808215151515815260200191505060405180910390f35b6104f46004803603602081101561045157600080fd5b810190808035906020019064010000000081111561046e57600080fd5b82018360208201111561048057600080fd5b803590602001918460018302840111640100000000831117156104a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610c47565b604051808215151515815260200191505060405180910390f35b6105506004803603602081101561052457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f23565b005b6106156004803603604081101561056857600080fd5b810190808035906020019064010000000081111561058557600080fd5b82018360208201111561059757600080fd5b803590602001918460018302840111640100000000831117156105b957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506110d8565b604051808215151515815260200191505060405180910390f35b6106716004803603602081101561064557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111e4565b005b6001818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b60006002826040518082805190602001908083835b602083106106e157805182526020820191506020810190506020830392506106be565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b6107526113f4565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610814576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006109076113f4565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146109c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6001826040518082805190602001908083835b602083106109ff57805182526020820191506020810190506020830392506109dc565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff16610ab2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6964656e7469747920686173207265766f6b656400000000000000000000000081525060200191505060405180910390fd5b60006001836040518082805190602001908083835b60208310610aea5780518252602082019150602081019050602083039250610ac7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff021916908315150217905550426002836040518082805190602001908083835b60208310610b6b5780518252602082019150602081019050602083039250610b48565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055507fdc730e218cfc1f20ca0db2b8f2d6de5a20643197dfce3030ffb6162d979eec56826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610c04578082015181840152602081019050610be9565b50505050905090810190601f168015610c315780820380516001836020036101000a031916815260200191505b509250505060405180910390a160019050919050565b6000610c516113f4565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610d13576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6001826040518082805190602001908083835b60208310610d495780518252602082019150602081019050602083039250610d26565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff1615610dfd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6964656e7469747920686173206578697374000000000000000000000000000081525060200191505060405180910390fd5b600180836040518082805190602001908083835b60208310610e345780518252602082019150602081019050602083039250610e11565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548160ff0219169083151502179055507fd78c379e1912ab387fe6b610ef43af8b64ea376337da41e5c4776532f6d84fa4826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610ee0578082015181840152602081019050610ec5565b50505050905090810190601f168015610f0d5780820380516001836020036101000a031916815260200191505b509250505060405180910390a160019050919050565b600060019054906101000a900460ff1680610f4a57506000809054906101000a900460ff16155b610f9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611423602e913960400191505060405180910390fd5b60008060019054906101000a900460ff161590508015610fef576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b610ff76113f4565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a381600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156110d45760008060016101000a81548160ff0219169083151502179055505b5050565b60006001836040518082805190602001908083835b6020831061111057805182526020820191506020810190506020830392506110ed565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900460ff161561115f57600190506111de565b816002846040518082805190602001908083835b602083106111965780518252602082019150602081019050602083039250611173565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205411156111d957600190506111de565b600090505b92915050565b6111ec6113f4565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146112ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611334576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806113fd6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60003390509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a264697066735822122093aa6fb138aeb0fc269a36f34afad68dc5386ee9cd7f99ac835447d78dd1783864736f6c634300060a0033"

// DeployIdentity deploys a new contract, binding an instance of Identity to it.
func DeployIdentity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Identity, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

func AsyncDeployIdentity(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(IdentityBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Identity is an auto generated Go binding around a Solidity contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around a Solidity contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around a Solidity contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around a Solidity contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// IsValid is a free data retrieval call binding the contract method 0x0a03e738.
//
// Solidity: function isValid(bytes ) constant returns(bool)
func (_Identity *IdentityCaller) IsValid(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "isValid", arg0)
	return *ret0, err
}

// IsValid is a free data retrieval call binding the contract method 0x0a03e738.
//
// Solidity: function isValid(bytes ) constant returns(bool)
func (_Identity *IdentitySession) IsValid(arg0 []byte) (bool, error) {
	return _Identity.Contract.IsValid(&_Identity.CallOpts, arg0)
}

// IsValid is a free data retrieval call binding the contract method 0x0a03e738.
//
// Solidity: function isValid(bytes ) constant returns(bool)
func (_Identity *IdentityCallerSession) IsValid(arg0 []byte) (bool, error) {
	return _Identity.Contract.IsValid(&_Identity.CallOpts, arg0)
}

// LastEfficacyTime is a free data retrieval call binding the contract method 0x2871c8a9.
//
// Solidity: function lastEfficacyTime(bytes ) constant returns(uint256)
func (_Identity *IdentityCaller) LastEfficacyTime(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "lastEfficacyTime", arg0)
	return *ret0, err
}

// LastEfficacyTime is a free data retrieval call binding the contract method 0x2871c8a9.
//
// Solidity: function lastEfficacyTime(bytes ) constant returns(uint256)
func (_Identity *IdentitySession) LastEfficacyTime(arg0 []byte) (*big.Int, error) {
	return _Identity.Contract.LastEfficacyTime(&_Identity.CallOpts, arg0)
}

// LastEfficacyTime is a free data retrieval call binding the contract method 0x2871c8a9.
//
// Solidity: function lastEfficacyTime(bytes ) constant returns(uint256)
func (_Identity *IdentityCallerSession) LastEfficacyTime(arg0 []byte) (*big.Int, error) {
	return _Identity.Contract.LastEfficacyTime(&_Identity.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Identity *IdentityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Identity *IdentitySession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Identity *IdentityCallerSession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// QueryEfficacyTime is a free data retrieval call binding the contract method 0x1c3f03c9.
//
// Solidity: function queryEfficacyTime(bytes _identity) constant returns(uint256)
func (_Identity *IdentityCaller) QueryEfficacyTime(opts *bind.CallOpts, _identity []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "queryEfficacyTime", _identity)
	return *ret0, err
}

// QueryEfficacyTime is a free data retrieval call binding the contract method 0x1c3f03c9.
//
// Solidity: function queryEfficacyTime(bytes _identity) constant returns(uint256)
func (_Identity *IdentitySession) QueryEfficacyTime(_identity []byte) (*big.Int, error) {
	return _Identity.Contract.QueryEfficacyTime(&_Identity.CallOpts, _identity)
}

// QueryEfficacyTime is a free data retrieval call binding the contract method 0x1c3f03c9.
//
// Solidity: function queryEfficacyTime(bytes _identity) constant returns(uint256)
func (_Identity *IdentityCallerSession) QueryEfficacyTime(_identity []byte) (*big.Int, error) {
	return _Identity.Contract.QueryEfficacyTime(&_Identity.CallOpts, _identity)
}

// QueryIdentity is a free data retrieval call binding the contract method 0xe6ddae4f.
//
// Solidity: function queryIdentity(bytes _identity, uint256 _time) constant returns(bool)
func (_Identity *IdentityCaller) QueryIdentity(opts *bind.CallOpts, _identity []byte, _time *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "queryIdentity", _identity, _time)
	return *ret0, err
}

// QueryIdentity is a free data retrieval call binding the contract method 0xe6ddae4f.
//
// Solidity: function queryIdentity(bytes _identity, uint256 _time) constant returns(bool)
func (_Identity *IdentitySession) QueryIdentity(_identity []byte, _time *big.Int) (bool, error) {
	return _Identity.Contract.QueryIdentity(&_Identity.CallOpts, _identity, _time)
}

// QueryIdentity is a free data retrieval call binding the contract method 0xe6ddae4f.
//
// Solidity: function queryIdentity(bytes _identity, uint256 _time) constant returns(bool)
func (_Identity *IdentityCallerSession) QueryIdentity(_identity []byte, _time *big.Int) (bool, error) {
	return _Identity.Contract.QueryIdentity(&_Identity.CallOpts, _identity, _time)
}

// AddIdentity is a paid mutator transaction binding the contract method 0xa77fe1a8.
//
// Solidity: function addIdentity(bytes _identity) returns(bool)
func (_Identity *IdentityTransactor) AddIdentity(opts *bind.TransactOpts, _identity []byte) (*types.Transaction, *types.Receipt, error) {
	return _Identity.contract.Transact(opts, "addIdentity", _identity)
}

func (_Identity *IdentityTransactor) AsyncAddIdentity(handler func(*types.Receipt, error), opts *bind.TransactOpts, _identity []byte) (*types.Transaction, error) {
	return _Identity.contract.AsyncTransact(opts, handler, "addIdentity", _identity)
}

// AddIdentity is a paid mutator transaction binding the contract method 0xa77fe1a8.
//
// Solidity: function addIdentity(bytes _identity) returns(bool)
func (_Identity *IdentitySession) AddIdentity(_identity []byte) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.AddIdentity(&_Identity.TransactOpts, _identity)
}

func (_Identity *IdentitySession) AsyncAddIdentity(handler func(*types.Receipt, error), _identity []byte) (*types.Transaction, error) {
	return _Identity.Contract.AsyncAddIdentity(handler, &_Identity.TransactOpts, _identity)
}

// AddIdentity is a paid mutator transaction binding the contract method 0xa77fe1a8.
//
// Solidity: function addIdentity(bytes _identity) returns(bool)
func (_Identity *IdentityTransactorSession) AddIdentity(_identity []byte) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.AddIdentity(&_Identity.TransactOpts, _identity)
}

func (_Identity *IdentityTransactorSession) AsyncAddIdentity(handler func(*types.Receipt, error), _identity []byte) (*types.Transaction, error) {
	return _Identity.Contract.AsyncAddIdentity(handler, &_Identity.TransactOpts, _identity)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_Identity *IdentityTransactor) Initialize(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Identity.contract.Transact(opts, "initialize", newOwner)
}

func (_Identity *IdentityTransactor) AsyncInitialize(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identity.contract.AsyncTransact(opts, handler, "initialize", newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_Identity *IdentitySession) Initialize(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.Initialize(&_Identity.TransactOpts, newOwner)
}

func (_Identity *IdentitySession) AsyncInitialize(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AsyncInitialize(handler, &_Identity.TransactOpts, newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_Identity *IdentityTransactorSession) Initialize(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.Initialize(&_Identity.TransactOpts, newOwner)
}

func (_Identity *IdentityTransactorSession) AsyncInitialize(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AsyncInitialize(handler, &_Identity.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentityTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Identity.contract.Transact(opts, "renounceOwnership")
}

func (_Identity *IdentityTransactor) AsyncRenounceOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.AsyncTransact(opts, handler, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentitySession) RenounceOwnership() (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.RenounceOwnership(&_Identity.TransactOpts)
}

func (_Identity *IdentitySession) AsyncRenounceOwnership(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Identity.Contract.AsyncRenounceOwnership(handler, &_Identity.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentityTransactorSession) RenounceOwnership() (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.RenounceOwnership(&_Identity.TransactOpts)
}

func (_Identity *IdentityTransactorSession) AsyncRenounceOwnership(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Identity.Contract.AsyncRenounceOwnership(handler, &_Identity.TransactOpts)
}

// RevokeIdentity is a paid mutator transaction binding the contract method 0xa16ce21d.
//
// Solidity: function revokeIdentity(bytes _identity) returns(bool)
func (_Identity *IdentityTransactor) RevokeIdentity(opts *bind.TransactOpts, _identity []byte) (*types.Transaction, *types.Receipt, error) {
	return _Identity.contract.Transact(opts, "revokeIdentity", _identity)
}

func (_Identity *IdentityTransactor) AsyncRevokeIdentity(handler func(*types.Receipt, error), opts *bind.TransactOpts, _identity []byte) (*types.Transaction, error) {
	return _Identity.contract.AsyncTransact(opts, handler, "revokeIdentity", _identity)
}

// RevokeIdentity is a paid mutator transaction binding the contract method 0xa16ce21d.
//
// Solidity: function revokeIdentity(bytes _identity) returns(bool)
func (_Identity *IdentitySession) RevokeIdentity(_identity []byte) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.RevokeIdentity(&_Identity.TransactOpts, _identity)
}

func (_Identity *IdentitySession) AsyncRevokeIdentity(handler func(*types.Receipt, error), _identity []byte) (*types.Transaction, error) {
	return _Identity.Contract.AsyncRevokeIdentity(handler, &_Identity.TransactOpts, _identity)
}

// RevokeIdentity is a paid mutator transaction binding the contract method 0xa16ce21d.
//
// Solidity: function revokeIdentity(bytes _identity) returns(bool)
func (_Identity *IdentityTransactorSession) RevokeIdentity(_identity []byte) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.RevokeIdentity(&_Identity.TransactOpts, _identity)
}

func (_Identity *IdentityTransactorSession) AsyncRevokeIdentity(handler func(*types.Receipt, error), _identity []byte) (*types.Transaction, error) {
	return _Identity.Contract.AsyncRevokeIdentity(handler, &_Identity.TransactOpts, _identity)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentityTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Identity.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_Identity *IdentityTransactor) AsyncTransferOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identity.contract.AsyncTransact(opts, handler, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentitySession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.TransferOwnership(&_Identity.TransactOpts, newOwner)
}

func (_Identity *IdentitySession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AsyncTransferOwnership(handler, &_Identity.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentityTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Identity.Contract.TransferOwnership(&_Identity.TransactOpts, newOwner)
}

func (_Identity *IdentityTransactorSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AsyncTransferOwnership(handler, &_Identity.TransactOpts, newOwner)
}

// IdentityAddIdentityEvent represents a AddIdentityEvent event raised by the Identity contract.
type IdentityAddIdentityEvent struct {
	Account []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchAddIdentityEvent is a free log subscription operation binding the contract event 0xd78c379e1912ab387fe6b610ef43af8b64ea376337da41e5c4776532f6d84fa4.
//
// Solidity: event AddIdentityEvent(bytes account)
func (_Identity *IdentityFilterer) WatchAddIdentityEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.contract.WatchLogs(fromBlock, handler, "AddIdentityEvent")
}

func (_Identity *IdentityFilterer) WatchAllAddIdentityEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.contract.WatchLogs(fromBlock, handler, "AddIdentityEvent")
}

// ParseAddIdentityEvent is a log parse operation binding the contract event 0xd78c379e1912ab387fe6b610ef43af8b64ea376337da41e5c4776532f6d84fa4.
//
// Solidity: event AddIdentityEvent(bytes account)
func (_Identity *IdentityFilterer) ParseAddIdentityEvent(log types.Log) (*IdentityAddIdentityEvent, error) {
	event := new(IdentityAddIdentityEvent)
	if err := _Identity.contract.UnpackLog(event, "AddIdentityEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchAddIdentityEvent is a free log subscription operation binding the contract event 0xd78c379e1912ab387fe6b610ef43af8b64ea376337da41e5c4776532f6d84fa4.
//
// Solidity: event AddIdentityEvent(bytes account)
func (_Identity *IdentitySession) WatchAddIdentityEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.Contract.WatchAddIdentityEvent(fromBlock, handler)
}

func (_Identity *IdentitySession) WatchAllAddIdentityEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.Contract.WatchAllAddIdentityEvent(fromBlock, handler)
}

// ParseAddIdentityEvent is a log parse operation binding the contract event 0xd78c379e1912ab387fe6b610ef43af8b64ea376337da41e5c4776532f6d84fa4.
//
// Solidity: event AddIdentityEvent(bytes account)
func (_Identity *IdentitySession) ParseAddIdentityEvent(log types.Log) (*IdentityAddIdentityEvent, error) {
	return _Identity.Contract.ParseAddIdentityEvent(log)
}

// IdentityLoseEfficacyEvent represents a LoseEfficacyEvent event raised by the Identity contract.
type IdentityLoseEfficacyEvent struct {
	Account []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchLoseEfficacyEvent is a free log subscription operation binding the contract event 0xdc730e218cfc1f20ca0db2b8f2d6de5a20643197dfce3030ffb6162d979eec56.
//
// Solidity: event LoseEfficacyEvent(bytes account)
func (_Identity *IdentityFilterer) WatchLoseEfficacyEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.contract.WatchLogs(fromBlock, handler, "LoseEfficacyEvent")
}

func (_Identity *IdentityFilterer) WatchAllLoseEfficacyEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.contract.WatchLogs(fromBlock, handler, "LoseEfficacyEvent")
}

// ParseLoseEfficacyEvent is a log parse operation binding the contract event 0xdc730e218cfc1f20ca0db2b8f2d6de5a20643197dfce3030ffb6162d979eec56.
//
// Solidity: event LoseEfficacyEvent(bytes account)
func (_Identity *IdentityFilterer) ParseLoseEfficacyEvent(log types.Log) (*IdentityLoseEfficacyEvent, error) {
	event := new(IdentityLoseEfficacyEvent)
	if err := _Identity.contract.UnpackLog(event, "LoseEfficacyEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchLoseEfficacyEvent is a free log subscription operation binding the contract event 0xdc730e218cfc1f20ca0db2b8f2d6de5a20643197dfce3030ffb6162d979eec56.
//
// Solidity: event LoseEfficacyEvent(bytes account)
func (_Identity *IdentitySession) WatchLoseEfficacyEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.Contract.WatchLoseEfficacyEvent(fromBlock, handler)
}

func (_Identity *IdentitySession) WatchAllLoseEfficacyEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.Contract.WatchAllLoseEfficacyEvent(fromBlock, handler)
}

// ParseLoseEfficacyEvent is a log parse operation binding the contract event 0xdc730e218cfc1f20ca0db2b8f2d6de5a20643197dfce3030ffb6162d979eec56.
//
// Solidity: event LoseEfficacyEvent(bytes account)
func (_Identity *IdentitySession) ParseLoseEfficacyEvent(log types.Log) (*IdentityLoseEfficacyEvent, error) {
	return _Identity.Contract.ParseLoseEfficacyEvent(log)
}

// IdentityOwnershipTransferred represents a OwnershipTransferred event raised by the Identity contract.
type IdentityOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentityFilterer) WatchOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log), previousOwner common.Address, newOwner common.Address) error {
	return _Identity.contract.WatchLogs(fromBlock, handler, "OwnershipTransferred", previousOwner, newOwner)
}

func (_Identity *IdentityFilterer) WatchAllOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.contract.WatchLogs(fromBlock, handler, "OwnershipTransferred")
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentityFilterer) ParseOwnershipTransferred(log types.Log) (*IdentityOwnershipTransferred, error) {
	event := new(IdentityOwnershipTransferred)
	if err := _Identity.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentitySession) WatchOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log), previousOwner common.Address, newOwner common.Address) error {
	return _Identity.Contract.WatchOwnershipTransferred(fromBlock, handler, previousOwner, newOwner)
}

func (_Identity *IdentitySession) WatchAllOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Identity.Contract.WatchAllOwnershipTransferred(fromBlock, handler)
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentitySession) ParseOwnershipTransferred(log types.Log) (*IdentityOwnershipTransferred, error) {
	return _Identity.Contract.ParseOwnershipTransferred(log)
}
