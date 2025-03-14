package models

import (
	"time"
)

type SignupAccountDto struct {
	Username  string    `json:"username" validate:"required,min=4,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"alphanum,min=8,max=72"`
	Firstname string    `json:"firstname" validate:"required,alpha,max=50"`
	Lastname  string    `json:"lastname" validate:"required,alpha,max=50"`
	Birthday  time.Time `json:"birthday" validate:"required"`
}

type RoleBindDto struct {
	AccountID uint   `json:"account_id;omitempty" validate:"required"`
	RoleIDs   []uint `json:"roles_id" validate:"required"`
}
