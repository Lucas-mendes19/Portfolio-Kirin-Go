package repository

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userData = map[string]string{
	"name":     "Lucas",
	"email":    "test@gmail.com",
	"password": "123456",
}

func init() {
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func TestUserCreate(t *testing.T) {
	user, _ := entity.CreateUser(userData["name"], userData["email"], userData["password"])
	userRepo := NewUserRepository(db)

	err = userRepo.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err := db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

func TestUserFindByEmail(t *testing.T) {
	userRepo := NewUserRepository(db)

	user, _ := entity.CreateUser(userData["name"], userData["email"], userData["password"])
	err := userRepo.Create(user)
	assert.Nil(t, err)

	user, err = userRepo.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
	assert.Equal(t, user.Email, userData["email"])
}
