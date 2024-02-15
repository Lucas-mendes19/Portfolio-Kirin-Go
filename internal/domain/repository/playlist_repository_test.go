package repository

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var playlistData = map[string]string{
	"title": "My Playlist",
}

func init() {
	err = db.AutoMigrate(&entity.Playlist{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func TestPlaylistCreate(t *testing.T) {
	playlist, _ := entity.CreatePlaylist(playlistData["title"])
	playlistRepo := NewPlaylistRepository(db)

	err = playlistRepo.Create(playlist)
	assert.Nil(t, err)

	var playlistFound entity.Playlist
	err := db.First(&playlistFound, "id = ?", playlist.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, playlist.ID, playlistFound.ID)
	assert.Equal(t, playlist.Title, playlistFound.Title)
}

func TestPlaylistUpdate(t *testing.T) {
	playlist, _ := entity.CreatePlaylist(playlistData["title"])
	playlistRepo := NewPlaylistRepository(db)

	err := playlistRepo.Create(playlist)
	assert.Nil(t, err)

	playlist.Title = "Updated Playlist"
	err = playlistRepo.Update(playlist)
	assert.Nil(t, err)

	playlistFound, err := playlistRepo.Find(playlist.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, playlist.ID, playlistFound.ID)
	assert.Equal(t, playlist.Title, playlistFound.Title)
}

func TestPlaylistDelete(t *testing.T) {
	playlist, _ := entity.CreatePlaylist(playlistData["title"])
	playlistRepo := NewPlaylistRepository(db)

	err := playlistRepo.Create(playlist)
	assert.Nil(t, err)

	err = playlistRepo.Delete(playlist.ID.String())
	assert.Nil(t, err)
}
