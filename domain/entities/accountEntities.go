package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email    string `gorm:"not null;unique;"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Roles    []Role `gorm:"many2many:account_roles"`
	Profile  Profile
}

func (a *Account) AfterFind(tx *gorm.DB) error {
	a.Password = ""
	return nil
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hash)
	return nil
}
