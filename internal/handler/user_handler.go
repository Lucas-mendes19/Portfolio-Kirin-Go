package handler

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	UserRepo repository.UserInterface
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	userRepo := repository.NewUserRepository(db)

	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (h *UserHandler) Store(c echo.Context) error {
	var userDTO UserDTO

	if err := c.Bind(&userDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := entity.CreateUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.UserRepo.Create(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
