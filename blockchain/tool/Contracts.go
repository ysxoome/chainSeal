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

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"contractHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sealaddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"sealType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signAddr\",\"type\":\"address\"}],\"name\":\"AddSign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RevokeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sealsRouterAddr\",\"type\":\"address\"}],\"name\":\"SetAddr\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_contractHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sealaddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_sealType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_signAddr\",\"type\":\"address\"}],\"name\":\"addSign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_signHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sealaddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_sealType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_signAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_signTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_contractHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sealaddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_sealType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_signAddr\",\"type\":\"address\"}],\"name\":\"isExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_contractHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sealaddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_sealType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_signAddr\",\"type\":\"address\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_contractHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"queryDetail\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_signHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sealaddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_sealType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_signAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_signTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_contractHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sealaddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_sealType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_signAddr\",\"type\":\"address\"}],\"name\":\"querySignTimel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_signTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAsealsRouterAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"signatureMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x608060405234801561001057600080fd5b50612be4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806384c284801161007157806384c28480146107fb5780638da5cb5b14610a1f5780639d2a449414610a69578063bdbafcd414610c1c578063c4d66de814610e6a578063f2fde38b14610eae576100b4565b80630b03ea30146100b95780630e445722146100fd578063382dab96146102b05780633d92b594146104165780636417c8b51461063a578063715018a6146107f1575b600080fd5b6100fb600480360360208110156100cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ef2565b005b61029a600480360360a081101561011357600080fd5b810190808035906020019064010000000081111561013057600080fd5b82018360208201111561014257600080fd5b8035906020019184600183028401116401000000008311171561016457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101c757600080fd5b8201836020820111156101d957600080fd5b803590602001918460018302840111640100000000831117156101fb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110e6565b6040518082815260200191505060405180910390f35b610400600480360360408110156102c657600080fd5b81019080803590602001906401000000008111156102e357600080fd5b8201836020820111156102f557600080fd5b8035906020019184600183028401116401000000008311171561031757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561037a57600080fd5b82018360208201111561038c57600080fd5b803590602001918460018302840111640100000000831117156103ae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061143b565b6040518082815260200191505060405180910390f35b6104d96004803603604081101561042c57600080fd5b810190808035906020019064010000000081111561044957600080fd5b82018360208201111561045b57600080fd5b8035906020019184600183028401116401000000008311171561047d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919050505061148c565b6040518080602001806020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018660ff1660ff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838103835289818151815260200191508051906020019080838360005b83811015610593578082015181840152602081019050610578565b50505050905090810190601f1680156105c05780820380516001836020036101000a031916815260200191505b50838103825288818151815260200191508051906020019080838360005b838110156105f95780820151818401526020810190506105de565b50505050905090810190601f1680156106265780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6107d7600480360360a081101561065057600080fd5b810190808035906020019064010000000081111561066d57600080fd5b82018360208201111561067f57600080fd5b803590602001918460018302840111640100000000831117156106a157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561070457600080fd5b82018360208201111561071657600080fd5b8035906020019184600183028401116401000000008311171561073857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061185a565b604051808215151515815260200191505060405180910390f35b6107f9611887565b005b6108be6004803603604081101561081157600080fd5b810190808035906020019064010000000081111561082e57600080fd5b82018360208201111561084057600080fd5b8035906020019184600183028401116401000000008311171561086257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611a11565b6040518080602001806020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018660ff1660ff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838103835289818151815260200191508051906020019080838360005b8381101561097857808201518184015260208101905061095d565b50505050905090810190601f1680156109a55780820380516001836020036101000a031916815260200191505b50838103825288818151815260200191508051906020019080838360005b838110156109de5780820151818401526020810190506109c3565b50505050905090810190601f168015610a0b5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b610a27611bfa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610c06600480360360a0811015610a7f57600080fd5b8101908080359060200190640100000000811115610a9c57600080fd5b820183602082011115610aae57600080fd5b80359060200191846001830284011164010000000083111715610ad057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190640100000000811115610b3357600080fd5b820183602082011115610b4557600080fd5b80359060200191846001830284011164010000000083111715610b6757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611c23565b6040518082815260200191505060405180910390f35b610e50600480360360c0811015610c3257600080fd5b8101908080359060200190640100000000811115610c4f57600080fd5b820183602082011115610c6157600080fd5b80359060200191846001830284011164010000000083111715610c8357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190640100000000811115610ce657600080fd5b820183602082011115610cf857600080fd5b80359060200191846001830284011164010000000083111715610d1a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190640100000000811115610d7d57600080fd5b820183602082011115610d8f57600080fd5b80359060200191846001830284011164010000000083111715610db157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611dea565b604051808215151515815260200191505060405180910390f35b610eac60048036036020811015610e8057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612603565b005b610ef060048036036020811015610ec457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506127b8565b005b610efa6129c8565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610fbc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561105f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f5f6164647220697320300000000000000000000000000000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f2f75ca815bd84b784dc14bf0e44146ca33760aa65a4c5eb4f041ab85ac1ea58160405160405180910390a250565b60006110f5868686868661185a565b611167576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f5f7369676e206973206e6f74206578697374000000000000000000000000000081525060200191505060405180910390fd5b60006111768787878787611c23565b90506111806129d0565b6002886040518082805190602001908083835b602083106111b65780518252602082019150602081019050602083039250611193565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060018303815481106111f757fe5b90600052602060002090600502016040518060c0016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112a95780601f1061127e576101008083540402835291602001916112a9565b820191906000526020600020905b81548152906001019060200180831161128c57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561134b5780601f106113205761010080835404028352916020019161134b565b820191906000526020600020905b81548152906001019060200180831161132e57829003601f168201915b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff1660ff1660ff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152505090508060a001519250505095945050505050565b60038280516020810182018051848252602083016020850120818352809550505050505081805160208101820180518482526020830160208501208183528095505050505050600091509150505481565b606080600080600080866002896040518082805190602001908083835b602083106114cc57805182526020820191506020810190506020830392506114a9565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020805490501015611577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f5f696e6465782067726561746572207468656e206c656e67746800000000000081525060200191505060405180910390fd5b61157f6129d0565b6002896040518082805190602001908083835b602083106115b55780518252602082019150602081019050602083039250611592565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902088815481106115f357fe5b90600052602060002090600502016040518060c0016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116a55780601f1061167a576101008083540402835291602001916116a5565b820191906000526020600020905b81548152906001019060200180831161168857829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117475780601f1061171c57610100808354040283529160200191611747565b820191906000526020600020905b81548152906001019060200180831161172a57829003601f168201915b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff1660ff1660ff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815250509050806000015181602001518260400151836060015184608001518560a00151965096509650965096509650509295509295509295565b60008061186a8787878787611c23565b1115611879576001905061187e565b600090505b95945050505050565b61188f6129c8565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611951576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6002828051602081018201805184825260208301602085012081835280955050505050508181548110611a4057fe5b906000526020600020906005020160009150915050806000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611aed5780601f10611ac257610100808354040283529160200191611aed565b820191906000526020600020905b815481529060010190602001808311611ad057829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611b8b5780601f10611b6057610100808354040283529160200191611b8b565b820191906000526020600020905b815481529060010190602001808311611b6e57829003601f168201915b5050505050908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900460ff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905086565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006003866040518082805190602001908083835b60208310611c5b5780518252602082019150602081019050602083039250611c38565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020858585856040516020018085805190602001908083835b60208310611cc85780518252602082019150602081019050602083039250611ca5565b6001836020036101000a0380198251168184511680821785525050505050509050018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018360ff1660ff1660f81b81526001018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014019450505050506040516020818303038152906040526040518082805190602001908083835b60208310611dab5780518252602082019150602081019050602083039250611d88565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054905095945050505050565b6000611df9868686868661185a565b15611e6c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f5f7369676e20697320657869737400000000000000000000000000000000000081525060200191505060405180910390fd5b60008360ff161415611f1f57600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614611f1a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f5f7365616c61646472206572726f72000000000000000000000000000000000081525060200191505060405180910390fd5b61205e565b60008473ffffffffffffffffffffffffffffffffffffffff16630e800b9f85856040518363ffffffff1660e01b8152600401808360ff1660ff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b158015611fac57600080fd5b505afa158015611fc0573d6000803e3d6000fd5b505050506040513d6020811015611fd657600080fd5b810190808051906020019092919050505090508061205c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f5f7369676e4164647220686173206e6f74207065726d697373696f6e0000000081525060200191505060405180910390fd5b505b60006002876040518082805190602001908083835b602083106120965780518252602082019150602081019050602083039250612073565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390209050806040518060c001604052808a81526020018881526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018660ff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001428152509080600181540180825580915050600190039060005260206000209060050201600090919091909150600082015181600001908051906020019061216d929190612a35565b50602082015181600101908051906020019061218a929190612ab5565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff021916908360ff16021790555060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040155505080805490506003886040518082805190602001908083835b60208310612281578051825260208201915060208101905060208303925061225e565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020878787876040516020018085805190602001908083835b602083106122ee57805182526020820191506020810190506020830392506122cb565b6001836020036101000a0380198251168184511680821785525050505050509050018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018360ff1660ff1660f81b81526001018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014019450505050506040516020818303038152906040526040518082805190602001908083835b602083106123d157805182526020820191506020810190506020830392506123ae565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055507f60ce163fbdbd2df8979e1fdf4b7d557de670c2964294198fcd7be653084b2105888888888888604051808060200180602001806020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018660ff1660ff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184810384528a818151815260200191508051906020019080838360005b838110156124e75780820151818401526020810190506124cc565b50505050905090810190601f1680156125145780820380516001836020036101000a031916815260200191505b50848103835289818151815260200191508051906020019080838360005b8381101561254d578082015181840152602081019050612532565b50505050905090810190601f16801561257a5780820380516001836020036101000a031916815260200191505b50848103825288818151815260200191508051906020019080838360005b838110156125b3578082015181840152602081019050612598565b50505050905090810190601f1680156125e05780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a160019150509695505050505050565b600060019054906101000a900460ff168061262a57506000809054906101000a900460ff16155b61267f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180612b81602e913960400191505060405180910390fd5b60008060019054906101000a900460ff1615905080156126cf576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6126d76129c8565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a381600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156127b45760008060016101000a81548160ff0219169083151502179055505b5050565b6127c06129c8565b73ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612882576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612908576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612b5b6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b6040518060c001604052806060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612a7657805160ff1916838001178555612aa4565b82800160010185558215612aa4579182015b82811115612aa3578251825591602001919060010190612a88565b5b509050612ab19190612b35565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612af657805160ff1916838001178555612b24565b82800160010185558215612b24579182015b82811115612b23578251825591602001919060010190612b08565b5b509050612b319190612b35565b5090565b612b5791905b80821115612b53576000816000905550600101612b3b565b5090565b9056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220a265f56befaa0a717f1d135505e072eae5d156239eb2261b84e08157d724bc0164736f6c634300060a0033"

// DeployContracts deploys a new contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

func AsyncDeployContracts(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Contracts is an auto generated Go binding around a Solidity contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around a Solidity contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around a Solidity contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around a Solidity contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// ContractMap is a free data retrieval call binding the contract method 0x84c28480.
//
// Solidity: function contractMap(bytes , uint256 ) constant returns(string _cid, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr, uint256 _signTime)
func (_Contracts *ContractsCaller) ContractMap(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (struct {
	Cid      string
	SignHash []byte
	Sealaddr common.Address
	SealType uint8
	SignAddr common.Address
	SignTime *big.Int
}, error) {
	ret := new(struct {
		Cid      string
		SignHash []byte
		Sealaddr common.Address
		SealType uint8
		SignAddr common.Address
		SignTime *big.Int
	})
	out := ret
	err := _Contracts.contract.Call(opts, out, "contractMap", arg0, arg1)
	return *ret, err
}

// ContractMap is a free data retrieval call binding the contract method 0x84c28480.
//
// Solidity: function contractMap(bytes , uint256 ) constant returns(string _cid, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr, uint256 _signTime)
func (_Contracts *ContractsSession) ContractMap(arg0 []byte, arg1 *big.Int) (struct {
	Cid      string
	SignHash []byte
	Sealaddr common.Address
	SealType uint8
	SignAddr common.Address
	SignTime *big.Int
}, error) {
	return _Contracts.Contract.ContractMap(&_Contracts.CallOpts, arg0, arg1)
}

// ContractMap is a free data retrieval call binding the contract method 0x84c28480.
//
// Solidity: function contractMap(bytes , uint256 ) constant returns(string _cid, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr, uint256 _signTime)
func (_Contracts *ContractsCallerSession) ContractMap(arg0 []byte, arg1 *big.Int) (struct {
	Cid      string
	SignHash []byte
	Sealaddr common.Address
	SealType uint8
	SignAddr common.Address
	SignTime *big.Int
}, error) {
	return _Contracts.Contract.ContractMap(&_Contracts.CallOpts, arg0, arg1)
}

// IsExist is a free data retrieval call binding the contract method 0x6417c8b5.
//
// Solidity: function isExist(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(bool)
func (_Contracts *ContractsCaller) IsExist(opts *bind.CallOpts, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "isExist", _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
	return *ret0, err
}

// IsExist is a free data retrieval call binding the contract method 0x6417c8b5.
//
// Solidity: function isExist(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(bool)
func (_Contracts *ContractsSession) IsExist(_contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (bool, error) {
	return _Contracts.Contract.IsExist(&_Contracts.CallOpts, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// IsExist is a free data retrieval call binding the contract method 0x6417c8b5.
//
// Solidity: function isExist(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(bool)
func (_Contracts *ContractsCallerSession) IsExist(_contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (bool, error) {
	return _Contracts.Contract.IsExist(&_Contracts.CallOpts, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Query is a free data retrieval call binding the contract method 0x9d2a4494.
//
// Solidity: function query(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(uint256)
func (_Contracts *ContractsCaller) Query(opts *bind.CallOpts, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "query", _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
	return *ret0, err
}

// Query is a free data retrieval call binding the contract method 0x9d2a4494.
//
// Solidity: function query(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(uint256)
func (_Contracts *ContractsSession) Query(_contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*big.Int, error) {
	return _Contracts.Contract.Query(&_Contracts.CallOpts, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// Query is a free data retrieval call binding the contract method 0x9d2a4494.
//
// Solidity: function query(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(uint256)
func (_Contracts *ContractsCallerSession) Query(_contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*big.Int, error) {
	return _Contracts.Contract.Query(&_Contracts.CallOpts, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// QueryDetail is a free data retrieval call binding the contract method 0x3d92b594.
//
// Solidity: function queryDetail(bytes _contractHash, uint256 _index) constant returns(string _cid, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr, uint256 _signTime)
func (_Contracts *ContractsCaller) QueryDetail(opts *bind.CallOpts, _contractHash []byte, _index *big.Int) (struct {
	Cid      string
	SignHash []byte
	Sealaddr common.Address
	SealType uint8
	SignAddr common.Address
	SignTime *big.Int
}, error) {
	ret := new(struct {
		Cid      string
		SignHash []byte
		Sealaddr common.Address
		SealType uint8
		SignAddr common.Address
		SignTime *big.Int
	})
	out := ret
	err := _Contracts.contract.Call(opts, out, "queryDetail", _contractHash, _index)
	return *ret, err
}

// QueryDetail is a free data retrieval call binding the contract method 0x3d92b594.
//
// Solidity: function queryDetail(bytes _contractHash, uint256 _index) constant returns(string _cid, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr, uint256 _signTime)
func (_Contracts *ContractsSession) QueryDetail(_contractHash []byte, _index *big.Int) (struct {
	Cid      string
	SignHash []byte
	Sealaddr common.Address
	SealType uint8
	SignAddr common.Address
	SignTime *big.Int
}, error) {
	return _Contracts.Contract.QueryDetail(&_Contracts.CallOpts, _contractHash, _index)
}

// QueryDetail is a free data retrieval call binding the contract method 0x3d92b594.
//
// Solidity: function queryDetail(bytes _contractHash, uint256 _index) constant returns(string _cid, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr, uint256 _signTime)
func (_Contracts *ContractsCallerSession) QueryDetail(_contractHash []byte, _index *big.Int) (struct {
	Cid      string
	SignHash []byte
	Sealaddr common.Address
	SealType uint8
	SignAddr common.Address
	SignTime *big.Int
}, error) {
	return _Contracts.Contract.QueryDetail(&_Contracts.CallOpts, _contractHash, _index)
}

// QuerySignTimel is a free data retrieval call binding the contract method 0x0e445722.
//
// Solidity: function querySignTimel(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(uint256 _signTime)
func (_Contracts *ContractsCaller) QuerySignTimel(opts *bind.CallOpts, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "querySignTimel", _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
	return *ret0, err
}

// QuerySignTimel is a free data retrieval call binding the contract method 0x0e445722.
//
// Solidity: function querySignTimel(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(uint256 _signTime)
func (_Contracts *ContractsSession) QuerySignTimel(_contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*big.Int, error) {
	return _Contracts.Contract.QuerySignTimel(&_Contracts.CallOpts, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// QuerySignTimel is a free data retrieval call binding the contract method 0x0e445722.
//
// Solidity: function querySignTimel(bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) constant returns(uint256 _signTime)
func (_Contracts *ContractsCallerSession) QuerySignTimel(_contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*big.Int, error) {
	return _Contracts.Contract.QuerySignTimel(&_Contracts.CallOpts, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// SignatureMap is a free data retrieval call binding the contract method 0x382dab96.
//
// Solidity: function signatureMap(bytes , bytes ) constant returns(uint256)
func (_Contracts *ContractsCaller) SignatureMap(opts *bind.CallOpts, arg0 []byte, arg1 []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "signatureMap", arg0, arg1)
	return *ret0, err
}

// SignatureMap is a free data retrieval call binding the contract method 0x382dab96.
//
// Solidity: function signatureMap(bytes , bytes ) constant returns(uint256)
func (_Contracts *ContractsSession) SignatureMap(arg0 []byte, arg1 []byte) (*big.Int, error) {
	return _Contracts.Contract.SignatureMap(&_Contracts.CallOpts, arg0, arg1)
}

// SignatureMap is a free data retrieval call binding the contract method 0x382dab96.
//
// Solidity: function signatureMap(bytes , bytes ) constant returns(uint256)
func (_Contracts *ContractsCallerSession) SignatureMap(arg0 []byte, arg1 []byte) (*big.Int, error) {
	return _Contracts.Contract.SignatureMap(&_Contracts.CallOpts, arg0, arg1)
}

// AddSign is a paid mutator transaction binding the contract method 0xbdbafcd4.
//
// Solidity: function addSign(string _cid, bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) returns(bool)
func (_Contracts *ContractsTransactor) AddSign(opts *bind.TransactOpts, _cid string, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.contract.Transact(opts, "addSign", _cid, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

func (_Contracts *ContractsTransactor) AsyncAddSign(handler func(*types.Receipt, error), opts *bind.TransactOpts, _cid string, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*types.Transaction, error) {
	return _Contracts.contract.AsyncTransact(opts, handler, "addSign", _cid, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// AddSign is a paid mutator transaction binding the contract method 0xbdbafcd4.
//
// Solidity: function addSign(string _cid, bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) returns(bool)
func (_Contracts *ContractsSession) AddSign(_cid string, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.AddSign(&_Contracts.TransactOpts, _cid, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

func (_Contracts *ContractsSession) AsyncAddSign(handler func(*types.Receipt, error), _cid string, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncAddSign(handler, &_Contracts.TransactOpts, _cid, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// AddSign is a paid mutator transaction binding the contract method 0xbdbafcd4.
//
// Solidity: function addSign(string _cid, bytes _contractHash, bytes _signHash, address _sealaddr, uint8 _sealType, address _signAddr) returns(bool)
func (_Contracts *ContractsTransactorSession) AddSign(_cid string, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.AddSign(&_Contracts.TransactOpts, _cid, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

func (_Contracts *ContractsTransactorSession) AsyncAddSign(handler func(*types.Receipt, error), _cid string, _contractHash []byte, _signHash []byte, _sealaddr common.Address, _sealType uint8, _signAddr common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncAddSign(handler, &_Contracts.TransactOpts, _cid, _contractHash, _signHash, _sealaddr, _sealType, _signAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.contract.Transact(opts, "initialize", newOwner)
}

func (_Contracts *ContractsTransactor) AsyncInitialize(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.AsyncTransact(opts, handler, "initialize", newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_Contracts *ContractsSession) Initialize(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, newOwner)
}

func (_Contracts *ContractsSession) AsyncInitialize(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncInitialize(handler, &_Contracts.TransactOpts, newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) Initialize(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, newOwner)
}

func (_Contracts *ContractsTransactorSession) AsyncInitialize(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncInitialize(handler, &_Contracts.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

func (_Contracts *ContractsTransactor) AsyncRenounceOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.AsyncTransact(opts, handler, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

func (_Contracts *ContractsSession) AsyncRenounceOwnership(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncRenounceOwnership(handler, &_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

func (_Contracts *ContractsTransactorSession) AsyncRenounceOwnership(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncRenounceOwnership(handler, &_Contracts.TransactOpts)
}

// SetAsealsRouterAddr is a paid mutator transaction binding the contract method 0x0b03ea30.
//
// Solidity: function setAsealsRouterAddr(address _addr) returns()
func (_Contracts *ContractsTransactor) SetAsealsRouterAddr(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.contract.Transact(opts, "setAsealsRouterAddr", _addr)
}

func (_Contracts *ContractsTransactor) AsyncSetAsealsRouterAddr(handler func(*types.Receipt, error), opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Contracts.contract.AsyncTransact(opts, handler, "setAsealsRouterAddr", _addr)
}

// SetAsealsRouterAddr is a paid mutator transaction binding the contract method 0x0b03ea30.
//
// Solidity: function setAsealsRouterAddr(address _addr) returns()
func (_Contracts *ContractsSession) SetAsealsRouterAddr(_addr common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.SetAsealsRouterAddr(&_Contracts.TransactOpts, _addr)
}

func (_Contracts *ContractsSession) AsyncSetAsealsRouterAddr(handler func(*types.Receipt, error), _addr common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncSetAsealsRouterAddr(handler, &_Contracts.TransactOpts, _addr)
}

// SetAsealsRouterAddr is a paid mutator transaction binding the contract method 0x0b03ea30.
//
// Solidity: function setAsealsRouterAddr(address _addr) returns()
func (_Contracts *ContractsTransactorSession) SetAsealsRouterAddr(_addr common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.SetAsealsRouterAddr(&_Contracts.TransactOpts, _addr)
}

func (_Contracts *ContractsTransactorSession) AsyncSetAsealsRouterAddr(handler func(*types.Receipt, error), _addr common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncSetAsealsRouterAddr(handler, &_Contracts.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_Contracts *ContractsTransactor) AsyncTransferOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.AsyncTransact(opts, handler, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

func (_Contracts *ContractsSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncTransferOwnership(handler, &_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

func (_Contracts *ContractsTransactorSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AsyncTransferOwnership(handler, &_Contracts.TransactOpts, newOwner)
}

// ContractsAddSign represents a AddSign event raised by the Contracts contract.
type ContractsAddSign struct {
	Cid          string
	ContractHash []byte
	SignHash     []byte
	Sealaddr     common.Address
	SealType     uint8
	SignAddr     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// WatchAddSign is a free log subscription operation binding the contract event 0x60ce163fbdbd2df8979e1fdf4b7d557de670c2964294198fcd7be653084b2105.
//
// Solidity: event AddSign(string cid, bytes contractHash, bytes signHash, address sealaddr, uint8 sealType, address signAddr)
func (_Contracts *ContractsFilterer) WatchAddSign(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "AddSign")
}

func (_Contracts *ContractsFilterer) WatchAllAddSign(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "AddSign")
}

// ParseAddSign is a log parse operation binding the contract event 0x60ce163fbdbd2df8979e1fdf4b7d557de670c2964294198fcd7be653084b2105.
//
// Solidity: event AddSign(string cid, bytes contractHash, bytes signHash, address sealaddr, uint8 sealType, address signAddr)
func (_Contracts *ContractsFilterer) ParseAddSign(log types.Log) (*ContractsAddSign, error) {
	event := new(ContractsAddSign)
	if err := _Contracts.contract.UnpackLog(event, "AddSign", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchAddSign is a free log subscription operation binding the contract event 0x60ce163fbdbd2df8979e1fdf4b7d557de670c2964294198fcd7be653084b2105.
//
// Solidity: event AddSign(string cid, bytes contractHash, bytes signHash, address sealaddr, uint8 sealType, address signAddr)
func (_Contracts *ContractsSession) WatchAddSign(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.Contract.WatchAddSign(fromBlock, handler)
}

func (_Contracts *ContractsSession) WatchAllAddSign(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.Contract.WatchAllAddSign(fromBlock, handler)
}

// ParseAddSign is a log parse operation binding the contract event 0x60ce163fbdbd2df8979e1fdf4b7d557de670c2964294198fcd7be653084b2105.
//
// Solidity: event AddSign(string cid, bytes contractHash, bytes signHash, address sealaddr, uint8 sealType, address signAddr)
func (_Contracts *ContractsSession) ParseAddSign(log types.Log) (*ContractsAddSign, error) {
	return _Contracts.Contract.ParseAddSign(log)
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log), previousOwner common.Address, newOwner common.Address) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "OwnershipTransferred", previousOwner, newOwner)
}

func (_Contracts *ContractsFilterer) WatchAllOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "OwnershipTransferred")
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsSession) WatchOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log), previousOwner common.Address, newOwner common.Address) error {
	return _Contracts.Contract.WatchOwnershipTransferred(fromBlock, handler, previousOwner, newOwner)
}

func (_Contracts *ContractsSession) WatchAllOwnershipTransferred(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.Contract.WatchAllOwnershipTransferred(fromBlock, handler)
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsSession) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	return _Contracts.Contract.ParseOwnershipTransferred(log)
}

// ContractsRevokeEvent represents a RevokeEvent event raised by the Contracts contract.
type ContractsRevokeEvent struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchRevokeEvent is a free log subscription operation binding the contract event 0x0d845ae82024b310a3b75b2ae201bda3809831b599b661519ac4dc43fe80802e.
//
// Solidity: event RevokeEvent(address indexed account, bool indexed status)
func (_Contracts *ContractsFilterer) WatchRevokeEvent(fromBlock *uint64, handler func(int, []types.Log), account common.Address, status bool) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "RevokeEvent", account, status)
}

func (_Contracts *ContractsFilterer) WatchAllRevokeEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "RevokeEvent")
}

// ParseRevokeEvent is a log parse operation binding the contract event 0x0d845ae82024b310a3b75b2ae201bda3809831b599b661519ac4dc43fe80802e.
//
// Solidity: event RevokeEvent(address indexed account, bool indexed status)
func (_Contracts *ContractsFilterer) ParseRevokeEvent(log types.Log) (*ContractsRevokeEvent, error) {
	event := new(ContractsRevokeEvent)
	if err := _Contracts.contract.UnpackLog(event, "RevokeEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRevokeEvent is a free log subscription operation binding the contract event 0x0d845ae82024b310a3b75b2ae201bda3809831b599b661519ac4dc43fe80802e.
//
// Solidity: event RevokeEvent(address indexed account, bool indexed status)
func (_Contracts *ContractsSession) WatchRevokeEvent(fromBlock *uint64, handler func(int, []types.Log), account common.Address, status bool) error {
	return _Contracts.Contract.WatchRevokeEvent(fromBlock, handler, account, status)
}

func (_Contracts *ContractsSession) WatchAllRevokeEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.Contract.WatchAllRevokeEvent(fromBlock, handler)
}

// ParseRevokeEvent is a log parse operation binding the contract event 0x0d845ae82024b310a3b75b2ae201bda3809831b599b661519ac4dc43fe80802e.
//
// Solidity: event RevokeEvent(address indexed account, bool indexed status)
func (_Contracts *ContractsSession) ParseRevokeEvent(log types.Log) (*ContractsRevokeEvent, error) {
	return _Contracts.Contract.ParseRevokeEvent(log)
}

// ContractsSetAddr represents a SetAddr event raised by the Contracts contract.
type ContractsSetAddr struct {
	SealsRouterAddr common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// WatchSetAddr is a free log subscription operation binding the contract event 0x2f75ca815bd84b784dc14bf0e44146ca33760aa65a4c5eb4f041ab85ac1ea581.
//
// Solidity: event SetAddr(address indexed sealsRouterAddr)
func (_Contracts *ContractsFilterer) WatchSetAddr(fromBlock *uint64, handler func(int, []types.Log), sealsRouterAddr common.Address) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "SetAddr", sealsRouterAddr)
}

func (_Contracts *ContractsFilterer) WatchAllSetAddr(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.contract.WatchLogs(fromBlock, handler, "SetAddr")
}

// ParseSetAddr is a log parse operation binding the contract event 0x2f75ca815bd84b784dc14bf0e44146ca33760aa65a4c5eb4f041ab85ac1ea581.
//
// Solidity: event SetAddr(address indexed sealsRouterAddr)
func (_Contracts *ContractsFilterer) ParseSetAddr(log types.Log) (*ContractsSetAddr, error) {
	event := new(ContractsSetAddr)
	if err := _Contracts.contract.UnpackLog(event, "SetAddr", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchSetAddr is a free log subscription operation binding the contract event 0x2f75ca815bd84b784dc14bf0e44146ca33760aa65a4c5eb4f041ab85ac1ea581.
//
// Solidity: event SetAddr(address indexed sealsRouterAddr)
func (_Contracts *ContractsSession) WatchSetAddr(fromBlock *uint64, handler func(int, []types.Log), sealsRouterAddr common.Address) error {
	return _Contracts.Contract.WatchSetAddr(fromBlock, handler, sealsRouterAddr)
}

func (_Contracts *ContractsSession) WatchAllSetAddr(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Contracts.Contract.WatchAllSetAddr(fromBlock, handler)
}

// ParseSetAddr is a log parse operation binding the contract event 0x2f75ca815bd84b784dc14bf0e44146ca33760aa65a4c5eb4f041ab85ac1ea581.
//
// Solidity: event SetAddr(address indexed sealsRouterAddr)
func (_Contracts *ContractsSession) ParseSetAddr(log types.Log) (*ContractsSetAddr, error) {
	return _Contracts.Contract.ParseSetAddr(log)
}
