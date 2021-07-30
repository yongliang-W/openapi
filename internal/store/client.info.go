package store

type PostgresClientInfo struct {
	Id string
	Secret string
	Domain string
	UserID string
}

func (ci *PostgresClientInfo) GetID() string {
	return ci.Id
}

func (ci *PostgresClientInfo) GetSecret() string {
	return ci.Secret
}

func (ci *PostgresClientInfo) GetDomain() string {
	return ci.Domain
}

func (ci *PostgresClientInfo) GetUserID() string {
	return ci.UserID
}
