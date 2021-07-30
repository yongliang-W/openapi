package service

import (
	"github.com/lishimeng/openapi/internal/db/model"
	"github.com/lishimeng/openapi/internal/db/repo"
)

func GetClientInfo(id string) (ci model.OauthClientInfo, err error) {
	return repo.GetClientInfo(id)
}
