package repo

import (
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/openapi/internal/db/model"
)

func GetClientInfo(id string) (ci model.OauthClientInfo, err error) {
	if len(id) == 0 {
		err = fmt.Errorf("not found")
		return
	}
	err = app.GetOrm().Context.
		QueryTable(new(model.OauthClientInfo)).
		Filter("ClientId", id).
		One(&ci)
	return
}
