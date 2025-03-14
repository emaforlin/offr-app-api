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

func (u *accountUsecaseImpl) BindRole(ctx context.Context, roleBinding *models.RoleBindDto) error {
	return u.repo.BindRoles(ctx, roleBinding.AccountID, roleBinding.RoleIDs)
}

// GetAccountByID implements AccountUsecase.
func (u *accountUsecaseImpl) GetAccountByID(ctx context.Context, id uint) (*entities.Account, error) {
	return u.repo.GetByID(ctx, id)
}

// SignupAccount implements AccountUsecase.
func (u *accountUsecaseImpl) SignupAccount(ctx context.Context, account *models.SignupAccountDto) error {
	var accountData = entities.Account{
		Email:    account.Email,
		Username: account.Username,
		Password: account.Password,
		Profile: entities.Profile{
			Firstname: account.Firstname,
			Lastname:  account.Lastname,
			Birthday:  account.Birthday,
		},
	}

	if err := u.repo.Create(ctx, &accountData); err != nil {
		return err
	}
	return nil
}

func NewAccountUsecase(repo repositories.AccountRepository) AccountUsecase {
	return &accountUsecaseImpl{repo: repo}
}
