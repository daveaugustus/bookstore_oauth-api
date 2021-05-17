package app

import (
	"github.com/davetweetlive/bookstore_oauth-api/src/domain/access_token"
	"github.com/davetweetlive/bookstore_oauth-api/src/repository/db"
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
}
