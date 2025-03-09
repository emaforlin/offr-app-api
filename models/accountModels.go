package models

import (
	"time"
)

type SignupAccountDto struct {
	Username  string    `json:"username" validate:"required,min=4,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password,min=8,max=72"`
	Firstname string    `json:"firstname" validate:"required,max=50"`
	Lastname  string    `json:"lastname" validate:"required,max=50"`
	Birthday  time.Time `json:"birthday" validate:"required"`
}

func (d *SignupAccountDto) Validate() error {
	err := validate.Struct(d)
	return err
}
