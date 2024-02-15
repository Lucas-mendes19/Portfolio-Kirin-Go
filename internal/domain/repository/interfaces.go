package repository

import "github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type PlaylistInterface interface {
	Find(id string) (*entity.Playlist, error)
	FindAll(page int, limit int, sort string) ([]entity.Playlist, error)
	Create(playlist *entity.Playlist) error
	Update(playlist *entity.Playlist) error
	Delete(id string) error
}
