package repository

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"gorm.io/gorm"
)

type Playlist struct {
	DB *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) *Playlist {
	return &Playlist{
		DB: db,
	}
}

func (p *Playlist) Find(id string) (*entity.Playlist, error) {
	var playlist entity.Playlist
	err := p.DB.First(&playlist, "id = ?", id).Error
	return &playlist, err
}

func (p *Playlist) FindAll(page int, limit int, sort string) ([]entity.Playlist, error) {
	var playlists []entity.Playlist
	var err error

	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Offset((page - 1) * limit).Limit(limit).Order("created_at " + sort).Find(&playlists).Error
		return playlists, err
	}

	err = p.DB.Find(&playlists).Error
	return playlists, err
}

func (p *Playlist) Create(playlist *entity.Playlist) error {
	return p.DB.Create(playlist).Error
}

func (p *Playlist) Update(playlist *entity.Playlist) error {
	return p.DB.Save(playlist).Error
}

func (p *Playlist) Delete(id string) error {
	playlist, err := p.Find(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(playlist).Error
}
