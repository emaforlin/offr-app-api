package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/emaforlin/offr-app-api/domain/entities"
	"github.com/emaforlin/offr-app-api/domain/repositories"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BindRoles implements repositories.AccountRepository.
func (r *mysqlRepositoryImpl) BindRoles(ctx context.Context, accountID uint, roleIDs []uint) error {
	var account entities.Account

	if err := r.db.Preload("Roles").First(&account, accountID).Error; err != nil {
		return errors.New("account not found")
	}

	var roles []entities.Role
	if err := r.db.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		return errors.New("failed to retrieve roles")
	}

	if err := r.db.Model(&account).Association("Roles").Replace(roles); err != nil {
		fmt.Printf("error: %v", err)
		return errors.New("role binding failed")
	}

	return nil

}

// Create implements AccountRepository.
func (r *mysqlRepositoryImpl) Create(ctx context.Context, account *entities.Account) error {
	err := r.db.Create(account).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062: // mysql code for DuplicatedEntry
			msg := "email or username already used"
			return errors.New(msg)
		default:
			return errors.New("failed to save account on db")
		}
	}
	return nil
}

// GetByEmail implements AccountRepository.
func (r *mysqlRepositoryImpl) GetByEmail(ctx context.Context, email uint) (*entities.Account, error) {
	var accountFound = &entities.Account{}

	if err := r.db.Preload(clause.Associations).First(accountFound, "email = ?", email).Error; err != nil {
		return nil, errors.New("failed to find account")
	}
	return accountFound, nil
}

// GetByID implements AccountRepository.
func (r *mysqlRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Account, error) {
	var accountFound = &entities.Account{}

	if err := r.db.Preload(clause.Associations).First(accountFound, id).Error; err != nil {
		return nil, errors.New("failed to find account")
	}
	return accountFound, nil
}

func NewAccountRepository(dbConn *gorm.DB) repositories.AccountRepository {
	return &mysqlRepositoryImpl{db: dbConn}
}
