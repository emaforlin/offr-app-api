package routes

import (
	"github.com/emaforlin/offr-app-api/infra/http/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, accountHandlers *handlers.AccountHandler) *gin.Engine {
	main := router.Group("/api/v1")
	main.POST("/signup", accountHandlers.HandleSignupAccount)

	accountGroup := main.Group("/accounts")
	{
		accountGroup.GET("/:id", accountHandlers.HandleGetAccountByEmail)
	}
	return router
}
