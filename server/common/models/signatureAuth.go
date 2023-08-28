package models

/**
 * @bref 签名授权结构体
 * @date 2023.7.25
 */

type SignatureAuth struct {
	SealAddress    string `json:"sealAddress"`    // 签章地址
	TypeAuth       uint8  `json:"typeAuth"`       // 授权类型,0取消授权，1授权
	TargetId       int64  `json:"targetId"`       // 授权用户id 公司授权给管理员
	TargetDid      string `json:"targetDid"`      // 授权用户DID
	AuthId         int64  `json:"authId"`         // 授权人id 公司
	AuthDid        string `json:"authDid"`        // 授权人DID
	AuthSign       string `json:"authSign"`       // 授权签名
	TxId           string `json:"txId"`           // 交易id
	ContractStatus uint8  `json:"contractStatus"` // 上链状态(0待上链，1进行中，2成功，3失败)
	CreateTime     string `json:"createTime"`     // 创建时间
	UpdateTime     string `json:"updateTime"`     // 更新时间
	CreateBy       int64  `json:"createBy"`       // 创建人
	UpdateBy       int64  `json:"updateBy"`       // 更新人
	AuthTime       int64  `json:"authTime"`       // 授权时间
}
