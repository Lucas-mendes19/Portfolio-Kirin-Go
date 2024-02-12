package entity

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/pkg/entity"
	"time"
)

type User struct {
	ID        entity.Id       `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  entity.Password `json:"-"`
	CreatedAt time.Time       `json:"created_at"`
}

func CreateUser(name string, email string, password string) (*User, error) {
	hash, err := entity.Encrypt(password)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:        entity.NewId(),
		Name:      name,
		Email:     email,
		Password:  hash,
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	return entity.Verify(password, u.Password)
}
