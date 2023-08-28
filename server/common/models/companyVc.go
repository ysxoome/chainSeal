package models

type CompanyVcEntity struct {
	//法人did
	LegalDid string `json:"legalDid"`

	//企业/公司名称
	Name string `json:"name"`

	//企业编号/统一社会信用代码
	CreditCode string `json:"creditCode"`

	//法人姓名
	LegalName string `json:"legalName"`

	//企业类型
	CompanyType int `json:"companyType"`
}

