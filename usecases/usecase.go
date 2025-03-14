package usecases

import (
	"context"

	"github.com/emaforlin/offr-app-api/domain/entities"
	"github.com/emaforlin/offr-app-api/models"
)

type AccountUsecase interface {
	SignupAccount(ctx context.Context, account *models.SignupAccountDto) error
	GetAccountByID(ctx context.Context, id uint) (*entities.Account, error)
	BindRole(ctx context.Context, roleBinding *models.RoleBindDto) error
}
