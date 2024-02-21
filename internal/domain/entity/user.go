package entity

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/pkg/entity"
	"gorm.io/gorm"
)

type User struct {
	ID       entity.Id       `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password entity.Password `json:"-"`
	gorm.Model
}

func CreateUser(name string, email string, password string) (*User, error) {
	hash, err := entity.Encrypt(password)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewId(),
		Name:     name,
		Email:    email,
		Password: hash,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	return entity.Verify(password, u.Password)
}
