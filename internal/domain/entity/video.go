package entity

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/pkg/entity"
	"gorm.io/gorm"
)

type Video struct {
	ID          entity.Id `json:"id"`
	PlaylistID  string    `json:"playlist_id"`
	Playlist    *Playlist `json:"playlist"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	gorm.Model
}

func CreateVideo(PlaylistId string, title string, description string, url string) (*Video, error) {
	return &Video{
		ID:          entity.NewId(),
		PlaylistID:  PlaylistId,
		Title:       title,
		Description: description,
		Url:         url,
	}, nil
}
