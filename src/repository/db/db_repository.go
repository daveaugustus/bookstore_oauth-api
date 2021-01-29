package db

import (
	"github.com/davetweetlive/bookstore_oauth-api/src/client/cassandra"
	"github.com/davetweetlive/bookstore_oauth-api/src/domain/access_token"
	"github.com/davetweetlive/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	// TODO: implement get access token from CassandraDB
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("Data base connection not implemented yet!")
}
