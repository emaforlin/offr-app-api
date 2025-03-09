package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emaforlin/offr-app-api/models"
	"github.com/emaforlin/offr-app-api/usecases"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	usecase usecases.AccountUsecase
}

func NewAccountHandler(usecase usecases.AccountUsecase) *AccountHandler {
	return &AccountHandler{usecase: usecase}
}

func (h *AccountHandler) HandleGetAccountByEmail(c *gin.Context) {
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

func (h *AccountHandler) HandleSignupAccount(c *gin.Context) {
	body := &models.SignupAccountDto{} // equal to: body := &entities.Account{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		fmt.Printf("error: %v", err)
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
	}
	if err := h.usecase.SignupAccount(c, body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signing up failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully"})
}
