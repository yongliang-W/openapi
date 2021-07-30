package model

type OauthOrganization struct {
	Pk
	OrgName string `orm:"column(org_name);unique"`
	TableChangeInfo
}

func (o *OauthOrganization) TableName() string {
	return "oauth_organization"
}

