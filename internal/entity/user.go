package entity

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/pkg/entity"
)

type User struct {
	Id       entity.Id       `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password entity.Password `json:"-"`
}

func create(name string, email string, password string) (*User, error) {
	hash, err := entity.Encrypt(password)

	if err != nil {
		return nil, err
	}

	return &User{
		Id:       entity.NewId(),
		Name:     name,
		Email:    email,
		Password: hash,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	return entity.Verify(password, u.Password)
}
