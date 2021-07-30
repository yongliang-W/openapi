package store

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/lishimeng/openapi/internal/db/service"
)

type PostgresClientStore struct {

}


func New() oauth2.ClientStore {
	var cs oauth2.ClientStore
	cs = &PostgresClientStore{}
	return cs
}

func (pcs *PostgresClientStore) GetByID(_ context.Context, id string) (ci oauth2.ClientInfo, err error) {
	var pci PostgresClientInfo
	oci, err := service.GetClientInfo(id)
	if err != nil {
		err = errors.ErrInvalidClient
		return
	}
	pci.Id = oci.ClientId
	pci.Secret = oci.ClientSecret
	pci.Domain = oci.Domain
	ci = &pci
	return
}
