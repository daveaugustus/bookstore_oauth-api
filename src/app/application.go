package app

import (
	"github.com/davetweetlive/bookstore_oauth-api/src/client/cassandra"
	"github.com/davetweetlive/bookstore_oauth-api/src/domain/access_token"
	"github.com/davetweetlive/bookstore_oauth-api/src/http"
	"github.com/davetweetlive/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session_, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session_.Close()

	atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":9090")
}
