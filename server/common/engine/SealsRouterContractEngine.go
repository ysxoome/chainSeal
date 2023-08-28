package engine

import (
	"encoding/base64"
	"log"
	"math/big"
	"seal/blockchain/tool"
	"seal/server/common/utils"
	"strconv"
	"strings"
	"time"

	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
)

// 全局session
var sealsRouterSession = tool.InitSealsRouter()

/* ========================================================== */
// Solidity: function calculateSeal(address _legalAddr, bytes _creditCode) constant returns(address seal)
func CalculateSeal(legalDid string, _creditCode string) string {
	_legalAddr := common.HexToAddress(utils.GidToAddr(legalDid))
	responseData, err := sealsRouterSession.CalculateSeal(_legalAddr, []byte(_creditCode))
	if err != nil {
		log.Println(err)
		return ""
	}
	// log.Println("CalculateSeal responseData:", responseData)
	return responseData.String()
}

// Solidity: function createSeal(address _legalAddr, bytes _creditCode) returns(address sealAddr)
func CreateSeal(legalDid string, _creditCode string) bool {
	_legalAddr := common.HexToAddress(utils.GidToAddr(legalDid))
	_, receipt, err := sealsRouterSession.CreateSeal(_legalAddr, []byte(_creditCode))
	if err != nil {
		log.Println(err)
		return false
	}

	//sealCreate event
	timeout := 10 * time.Second
	queryTicker := time.NewTicker(timeout)
	done := make(chan bool)
	from, _ := strconv.ParseUint(receipt.GetBlockNumber(), 0, 64)
	err = sealsRouterSession.WatchSealsCreated(&from, func(status int, logs []types.Log) {
		for _, v := range logs {
			_, errs := sealsRouterSession.ParseSealsCreated(v)
			if errs != nil {
				log.Println("ParseSealCreated err:", errs)
			}
		}
		queryTicker.Stop()
		queryTicker = time.NewTicker(timeout)
		done <- true
	}, _legalAddr)
	if err != nil {
		log.Println("WatchSealCreated err:", err)
		return false
	}
	//根据交易回执状态码判断是否成功
	return receipt.GetStatus() == 0
}

// Solidity: function encodePacked(address sealAddr, uint8 _sealType, address _addr, uint256 _signTime) constant returns(bytes res)
func EncodePacked(sealAddr string, sealType uint8, addr string, signTime int64) []byte {

	_sealAddr := common.HexToAddress(sealAddr)
	_sealType := sealType
	_addr := common.HexToAddress(utils.GidToAddr(addr))
	_signTime := new(big.Int).SetInt64(signTime)
	encodePacked, err := sealsRouterSession.EncodePacked(_sealAddr, _sealType, _addr, _signTime)
	if err != nil {
		log.Println(err)
	}
	return encodePacked
}

// Solidity: function querySealApprovl(address sealAddr, uint8 _sealType, address _addr) constant returns(bool)
func QuerySealApprovl(sealAddr string, _sealType uint8, addr string) bool {
	_sealAddr := common.HexToAddress(sealAddr)
	_addr := common.HexToAddress(utils.GidToAddr(addr))
	res, err := sealsRouterSession.QuerySealApprovl(_sealAddr, _sealType, _addr)
	if err != nil {
		log.Println(err)
		return false
	}
	return res
}

// Solidity: function querySealStatus(address sealAddr) constant returns(bool)
func QuerySealStatus(sealAddr string) bool {
	_sealAddr := common.HexToAddress(sealAddr)
	status, err := sealsRouterSession.QuerySealStatus(_sealAddr)
	if err != nil {
		log.Println(err)
	}
	return status
}

// Solidity: function approvalDelegate(address sealAddr, uint8 _sealType, address _addr, uint256 _signTime, uint8 v, bytes32 r, bytes32 s) returns(bool)
func ApprovalDelegate(sealAddr string, sealType uint8, addr string, signTime int64, base64Sign string) (bool, *types.Receipt) {
	_sealAddr := common.HexToAddress(sealAddr)
	_sealType := sealType
	_approvalAddr := common.HexToAddress(utils.GidToAddr(addr))
	_signTime := new(big.Int).SetInt64(signTime)

	_base64Sign, _ := base64.StdEncoding.DecodeString(base64Sign)
	var v uint8
	var r, s [32]byte
	v = _base64Sign[64] + 27
	copy(r[:], _base64Sign[:32])
	copy(s[:], _base64Sign[32:64])

	_, receipt, err := sealsRouterSession.ApprovalDelegate(_sealAddr, _sealType, _approvalAddr, _signTime, v, r, s)
	if err != nil {
		log.Println(err)
		return false, nil
	}
	// log.Println("ApprovalDelegateHash:", receipt.TransactionHash)

	//根据交易回执状态码判断是否成功
	return receipt.GetStatus() == 0, receipt
}

// Solidity: function revokeDelegate(address sealAddr, uint8 _sealType, address _addr, uint256 _signTime, uint8 v, bytes32 r, bytes32 s) returns(bool)
func RevokeDelegate(sealAddr string, sealType uint8, delegateAddr string, signTime int64, base64Sign string) (bool, *types.Receipt) {
	_sealAddr := common.HexToAddress(sealAddr)
	_delegateAddr := common.HexToAddress(utils.GidToAddr(delegateAddr))
	_sealType := sealType
	_signTime := new(big.Int).SetInt64(signTime)
	_base64Sign, _ := base64.StdEncoding.DecodeString(base64Sign)
	var v uint8
	var r, s [32]byte
	v = _base64Sign[64] + 27
	copy(r[:], _base64Sign[:32])
	copy(s[:], _base64Sign[32:64])

	_, receipt, err := sealsRouterSession.RevokeDelegate(_sealAddr, _sealType, _delegateAddr, _signTime, v, r, s)
	if err != nil {
		log.Println(err)
		return false, nil
	}
	// log.Println("RevokeDelegateHash:", receipt.TransactionHash)

	//revokeDelegate event

	//根据交易回执状态码判断是否成功
	return receipt.GetStatus() == 0, receipt
}

// Solidity: function querySealOwner(address sealAddr) constant returns(address)
func IsSealOwner(sealAddr string, _did string) bool {
	_sealAddr := common.HexToAddress(sealAddr)
	address, err := sealsRouterSession.QuerySealOwner(_sealAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	//strings.EqualFold() 忽略大小写比较
	addrDid, _ := utils.AddrToGID(address.String())
	if !strings.EqualFold(addrDid, _did) {
		log.Println("address.String() != _did")
		return false
	}
	return true
}

// Solidity: function owner() constant returns(address)
func IsOwner_SealsRouter(_gid string) bool {
	ownr, err := sealsRouterSession.Owner()
	if err != nil {
		log.Println(err)
		return false
	}
	ownrGid, _ := utils.AddrToGID(ownr.String())
	return strings.EqualFold(ownrGid, _gid)
}

// Solidity: function getSeal(address , bytes ) constant returns(address)
func GetSeal(sealAddr string, _creditCode string) string {
	_sealAddr := common.HexToAddress(sealAddr)
	_arg1 := []byte(_creditCode)
	address, err := sealsRouterSession.GetSeal(_sealAddr, _arg1)
	if err != nil {
		log.Println(err)
	}
	log.Println("GetSeal address:", address)
	return address.String()
}

func GetAllSealAddress() {
	num, err := sealsRouterSession.AllSealsLength()
	if err != nil {
		log.Println(err)
	}
	log.Println("All seal number:", num)

	for i := 0; i < int(num.Int64()); i++ {
		address, err := sealsRouterSession.AllSeals(big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		log.Printf("第%d条seal address:::%s", i+1, address)
	}
}
