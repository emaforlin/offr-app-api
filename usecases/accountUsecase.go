package usecases

import (
	"context"

	"github.com/emaforlin/offr-app-api/domain/entities"
	"github.com/emaforlin/offr-app-api/domain/repositories"
	"github.com/emaforlin/offr-app-api/models"
)

type accountUsecaseImpl struct {
	repo repositories.AccountRepository
}

// GetAccountByID implements AccountUsecase.
func (u *accountUsecaseImpl) GetAccountByID(ctx context.Context, id uint) (*entities.Account, error) {
	return u.repo.GetByID(ctx, id)
}

// SignupAccount implements AccountUsecase.
func (u *accountUsecaseImpl) SignupAccount(ctx context.Context, account *models.SignupAccountDto) error {
	return u.repo.Create(ctx, &entities.Account{
		Email:    account.Email,
		Username: account.Username,
		Password: account.Password,
		Profile: entities.Profile{
			Firstname: account.Firstname,
			Lastname:  account.Lastname,
			Birthday:  account.Birthday,
		},
	})
}

func NewAccountUsecase(repo repositories.AccountRepository) AccountUsecase {
	return &accountUsecaseImpl{repo: repo}
}
