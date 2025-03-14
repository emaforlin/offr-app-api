package repositories

import (
	"context"

	"github.com/emaforlin/offr-app-api/domain/entities"
)

type AccountRepository interface {
	Create(ctx context.Context, account *entities.Account) error
	BindRoles(ctx context.Context, accountID uint, rolesIDs []uint) error
	GetByID(ctx context.Context, id uint) (*entities.Account, error)
	GetByEmail(ctx context.Context, email uint) (*entities.Account, error)
}
