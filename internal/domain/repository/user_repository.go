package repository

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}
func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
