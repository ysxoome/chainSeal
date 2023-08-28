package models

import "github.com/ethereum/go-ethereum/accounts/keystore"

/*
	Verifiable Credential Struct 可信凭证
*/

type PersonVC struct {

	// 钱包 主要包含钱包id、钱包地址(、版本号)
	KeyStore *keystore.Key `json:"keyStore"`

	// 姓名
	Name string `json:"name"`

	// 身份证号
	IdNumber string `json:"idNumber"`
}