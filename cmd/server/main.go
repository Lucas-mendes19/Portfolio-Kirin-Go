package main

import (
	"encoding/json"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Playlist{})
	if err != nil {
		return
	}

	playlistRepo := repository.NewPlaylistRepository(db)
	playlistHandler := NewPlaylistHandler(playlistRepo)

	http.HandleFunc("/playlists", playlistHandler.CreatePlaylist)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

type PlaylistDTO struct {
	Title string `json:"title"`
}

type PlaylistHandler struct {
	PlaylistDB repository.PlaylistInterface
}

func NewPlaylistHandler(db repository.PlaylistInterface) *PlaylistHandler {
	return &PlaylistHandler{
		PlaylistDB: db,
	}
}

func (handler *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	var dto PlaylistDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playlist, err := entity.CreatePlaylist(dto.Title)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.PlaylistDB.Create(playlist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
