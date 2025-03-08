package main

import (
	"github.com/emaforlin/offr-app-api/config"
	"github.com/emaforlin/offr-app-api/infra/database"
	routes "github.com/emaforlin/offr-app-api/infra/http"
	"github.com/emaforlin/offr-app-api/infra/http/handlers"
	"github.com/emaforlin/offr-app-api/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init("config")
	conf := config.Load()

	db, err := database.NewDBConn(conf)
	if err != nil {
		panic("can't connect to database")
	}

	// accounts
	accountRepo := database.NewAccountRepository(db)
	accountUsecase := usecases.NewAccountUsecase(accountRepo)
	accountHandlers := handlers.NewAccountHandler(accountUsecase)

	var addr = ":8080"
	if !conf.App.Debugmode {
		addr = ":443"
	}

	routes.InitRoutes(gin.Default(), accountHandlers).Run(addr)
}
