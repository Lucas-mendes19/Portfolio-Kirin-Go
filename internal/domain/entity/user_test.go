package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var userData = map[string]string{
	"name":     "Lucas",
	"email":    "test@gmail.com",
	"password": "123456",
}

func TestCreateUser(t *testing.T) {
	user, err := CreateUser(userData["name"], userData["email"], userData["password"])

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userData["name"], user.Name)
	assert.Equal(t, userData["email"], user.Email)
	assert.NotEqual(t, userData["password"], user.Password)
	assert.NotEmpty(t, user.CreatedAt)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := CreateUser(userData["name"], userData["email"], userData["password"])

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(userData["password"]))
	assert.False(t, user.ValidatePassword("invalid"))
	assert.NotEqual(t, userData["password"], user.Password)
}
