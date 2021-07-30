package model

type OauthClientInfo struct {
	Pk
	ClientId string `orm:"column(client_id);unique;null"`
	ClientSecret    string `orm:"column(client_secret);unique;null"`
	Domain    string `orm:"column(domain);unique;null"`
	TableChangeInfo
}

func (oci *OauthClientInfo) TableName() string {
	return "oauth_client_info"
}

