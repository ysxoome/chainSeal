package engine

import (
	"log"
	"math/big"
	"seal/blockchain/tool"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// 全局session
var identitySession = tool.InitIdentity()

/* ========================================================== */
// Solidity: function addIdentity(bytes _identity) returns(bool)
func AddIdentity(identityByte string) bool {
	hashByte, _ := hexutil.Decode(identityByte)
	_, _, err := identitySession.AddIdentity(hashByte)
	if err != nil {
		log.Println(err)
		return false
	}

	//addIdentityEventEvents, err := identitySession.ParseAddIdentityEvent(receipt)
	return true
}

// Solidity: function queryIdentity(bytes _identity, uint256 _time) constant returns(bool)
func QueryIdentity(identityByte string, time *big.Int) bool {
	hash, _ := hexutil.Decode(identityByte)
	response, err := identitySession.QueryIdentity(hash, time)
	if err != nil {
		log.Println(err)
		return false
	}
	// log.Println("queryIdentity responseDataBool:", responseDataBool)
	return response
}
