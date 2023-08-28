package engine

import (
	"fmt"
	"log"
	"math/big"
	"seal/blockchain/tool"
	"seal/server/common/utils"
	"strconv"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// 全局session
var contractsSession = tool.InitContracts()

func IsContractExisted(contractHash, signHash []byte, sealAddr string, sealType uint8, signAddr string) bool {
	exist, err := contractsSession.IsExist(contractHash, signHash, common.HexToAddress(sealAddr), sealType, common.HexToAddress(signAddr))
	if err != nil {
		log.Println("contractSession queryExist failed")
		return false
	}
	return exist
}

func QuerySignTimel(contractHash, signHash []byte, sealAddr string, sealType uint8, signAddr string) *big.Int {
	time, err := contractsSession.QuerySignTimel(contractHash, signHash, common.HexToAddress(sealAddr), sealType, common.HexToAddress(signAddr))
	if err != nil {
		fmt.Println("contractSession querySignTiml failed")
		return big.NewInt(0)
	}
	return time
}

func AddSign(cid string, contractHash, signHash []byte, sealAddr string, sealType uint8, signAddr string) (bool, *types.Receipt) {
	_, rec, err := contractsSession.AddSign(cid, contractHash, signHash, common.HexToAddress(sealAddr),
		sealType, common.HexToAddress(signAddr))
	if err != nil {
		log.Println("add sign failed")
		return false, nil
	}

	return rec.Status == 0, rec
}

/* ========================================================== */
func IsExisted(_contractHash string, _signHash string, _sealAddr string, sealType uint8, _signAddr string) bool {
	cHash, _ := hexutil.Decode(_contractHash)
	sHash, _ := hexutil.Decode(_signHash)
	res, err := contractsSession.IsExist(
		cHash,
		sHash,
		common.HexToAddress(_sealAddr),
		sealType,
		common.HexToAddress(_signAddr))
	if err != nil {
		return false
	}
	return res
}

func Query(_contractHash string, _signHash string, _sealAddr string, sealType uint8, _signAddr string) *big.Int {
	cHash, _ := hexutil.Decode(_contractHash)
	sHash, _ := hexutil.Decode(_signHash)
	res, err := contractsSession.Query(
		cHash,
		sHash,
		common.HexToAddress(_sealAddr),
		sealType,
		common.HexToAddress(_signAddr))
	if err != nil {
		log.Println(err)
		return nil
	}
	return res
}

func QuerySignTimel2(_contractHash string, _signHash string, _sealAddr string, sealType uint8, _signAddr string) *big.Int {
	cHash, _ := hexutil.Decode(_contractHash)
	sHash, _ := hexutil.Decode(_signHash)
	res, err := contractsSession.QuerySignTimel(
		cHash,
		sHash,
		common.HexToAddress(_sealAddr),
		sealType,
		common.HexToAddress(_signAddr))
	if err != nil {
		log.Println(err)
		return nil
	}
	return res
}

func QueryDetail(_contractHash string, _index uint64) interface{} {
	hash, _ := hexutil.Decode(_contractHash)
	res, err := contractsSession.QueryDetail(hash, new(big.Int).SetUint64(_index))
	if err != nil {
		log.Println(err)
		return nil
	}
	return res
}

func IsOwner_Contract(_gid string) bool {
	owner, err := contractsSession.Owner()
	if err != nil {
		log.Println(err)
		return false
	}
	ownerGid, _ := utils.AddrToGID(owner.String())
	return strings.EqualFold(ownerGid, _gid)
}

func AddSign2(cid, _contractHash string, _signHash string, _sealAddr string, sealType uint8, _signAddr string, base64Sign []byte) {
	//原版建立一个新的session
	client := tool.InitClient()
	defer client.Close()
	contractAddress := common.HexToAddress(tool.CONTRACTS_CONTRACT_ADDRESS)
	instance, err := tool.NewContracts(contractAddress, client)
	if err != nil {
		log.Println(err)
	}
	contractsSessionTemp := &tool.ContractsSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	_, receipt, err := contractsSessionTemp.AddSign(cid, []byte(_contractHash), []byte(_signHash),
		common.HexToAddress(_sealAddr), sealType, common.HexToAddress(_signAddr))
	if err != nil {
		log.Println(err)
	}
	log.Println("AddSign receipt:", receipt)

	//addSignEvent
	from, _ := strconv.ParseUint(receipt.GetBlockNumber(), 0, 64)
	err = contractsSessionTemp.WatchAllAddSign(&from, func(status int, logs []types.Log) {
		log.Println("status code:", status)
		for _, v := range logs {
			logData, errs := contractsSessionTemp.ParseAddSign(v)
			if errs != nil {
				log.Println("ParseAddSign err:", errs)
			}
			log.Println(logData)
		}
	})
	if err != nil {
		log.Println("WatchAddSign err:", err)
	}
	if receipt.GetStatus() == 0 {
		log.Println("AddSign success")
	}

}

// Solidity: function setAsealsRouterAddr(address _addr) returns()
func SetAsealsRouterAddr(_sealsRouterAddr string) bool {
	transaction, receipt, err := contractsSession.SetAsealsRouterAddr(common.HexToAddress(_sealsRouterAddr))
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("SetAsealsRouterAddr transaction:", transaction)
	log.Println("SetAsealsRouterAddr receipt:", receipt)

	//setAsealsRouterAddrEvent
	from, _ := strconv.ParseUint(receipt.GetBlockNumber(), 0, 64)
	err = contractsSession.WatchSetAddr(&from, func(status int, logs []types.Log) {
		log.Println("status code:", status)
		for _, v := range logs {
			logData, errs := contractsSession.ParseSetAddr(v)
			if errs != nil {
				log.Println("ParseSetAddr Event err:", errs)
			}
			log.Println(logData)
		}
	}, common.HexToAddress(_sealsRouterAddr))
	if err != nil {
		log.Println("WatchSetAddr err:", err)
		return false
	}
	return receipt.GetStatus() == 0
}
