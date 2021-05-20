package app

import (
	"github.com/davetweetlive/bookstore_oauth-api/src/http"
	"github.com/davetweetlive/bookstore_oauth-api/src/repository/db"
	"github.com/davetweetlive/bookstore_oauth-api/src/repository/rest"
	serviceAT "github.com/davetweetlive/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(serviceAT.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
