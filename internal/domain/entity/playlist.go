package entity

import (
	"errors"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/pkg/entity"
	"time"
)

var (
	ErrTitleRequired = errors.New("title is required")
)

type Playlist struct {
	ID        entity.Id `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePlaylist(title string) (*Playlist, error) {
	playlist := &Playlist{
		ID:        entity.NewId(),
		Title:     title,
		CreatedAt: time.Now(),
	}

	err := playlist.Validate()
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (p *Playlist) Validate() error {
	if p.Title == "" {
		return ErrTitleRequired
	}

	return nil
}
