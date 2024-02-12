package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var playlistData = map[string]string{
	"title": "My Playlist",
}

func TestCreatePlaylist(t *testing.T) {
	playlist, err := CreatePlaylist(playlistData["title"])

	assert.Nil(t, err)
	assert.NotNil(t, playlist)
	assert.Equal(t, playlistData["title"], playlist.Title)
	assert.NotEmpty(t, playlist.CreatedAt)
}

func TestPlaylistValidateTitleRequired(t *testing.T) {
	playlist := &Playlist{}

	err := playlist.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, ErrTitleRequired, err)
}
