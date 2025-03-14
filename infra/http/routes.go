package routes

import (
	"github.com/emaforlin/offr-app-api/infra/http/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, accountHandlers *handlers.AccountHandler) *gin.Engine {
	main := router.Group("/api/v1")

	authGroup := main.Group("/auth")

	authGroup.POST("/signup", accountHandlers.HandleSignupAccount)

	accountGroup := main.Group("/accounts")
	{
		// account group
		accountGroup.GET("/:id", accountHandlers.HandleGetAccountByID)

		// roles group
		accountGroup.POST("/:id/roles", accountHandlers.HandleBindRoles)
	}
	return router
}
