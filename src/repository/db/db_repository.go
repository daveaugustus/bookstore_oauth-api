package db

import (
	"github.com/davetweetlive/bookstore_oauth-api/src/client/cassandra"
	"github.com/davetweetlive/bookstore_oauth-api/src/domain/access_token"
	"github.com/davetweetlive/bookstore_oauth-api/src/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// TODO: implement get access token from CassandraDB.
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
