package repository

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"gorm.io/gorm"
)

type Video struct {
	DB *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *Video {
	return &Video{
		DB: db,
	}
}

func (v *Video) Find(id string) (*entity.Video, error) {
	var video entity.Video
	err := v.DB.Preload("Playlist").First(&video, "id = ?", id).Error
	return &video, err
}

func (v *Video) Create(video *entity.Video) error {
	return v.DB.Create(video).Error
}
