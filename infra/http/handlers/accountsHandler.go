package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaforlin/offr-app-api/usecases"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	usecase usecases.AccountUsecase
}

func NewAccountHandler(usecase usecases.AccountUsecase) *AccountHandler {
	return &AccountHandler{usecase: usecase}
}

func (h *AccountHandler) HandleGetUserByEmail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	account, err := h.usecase.GetAccountByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}
