package tool

import (
	"log"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func InitClient() *client.Client {
	configs, err := conf.ParseConfigFile("/Users/ysx/code/go/src/chainSeal/blockchain/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func InitIdentity() *IdentitySession {
	client := InitClient()
	contractAddress := common.HexToAddress(IDENTITY_CONTRACT_ADDRESS)
	instance, err := NewIdentity(contractAddress, client)
	if err != nil {
		errors.New("identity contract init failed")
	}

	identitySession := &IdentitySession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	return identitySession
}

func InitContracts() *ContractsSession {
	client := InitClient()
	contractAddress := common.HexToAddress(CONTRACTS_CONTRACT_ADDRESS)
	instance, err := NewContracts(contractAddress, client)
	if err != nil {
		errors.New("contracts contract init failed")
	}

	contractsSession := &ContractsSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	return contractsSession
}

func InitSealsRouter() *SealsrouterSession {
	client := InitClient()
	contractAddress := common.HexToAddress(SEALSROUTER_CONTRACT_ADDRESS)
	instance, err := NewSealsrouter(contractAddress, client)
	if err != nil {
		errors.New("sealrouter contract init failed")
	}

	sealsRouterSession := &SealsrouterSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	return sealsRouterSession
}
