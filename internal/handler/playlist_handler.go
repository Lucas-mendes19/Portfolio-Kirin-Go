package handler

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PlaylistDTO struct {
	Title string `json:"title"`
}

type PlaylistHandler struct {
	PlaylistRepo repository.PlaylistInterface
}

func NewPlaylistHandler(db *gorm.DB) *PlaylistHandler {
	playlistRepo := repository.NewPlaylistRepository(db)

	return &PlaylistHandler{
		PlaylistRepo: playlistRepo,
	}
}

func (h *PlaylistHandler) Index(c echo.Context) error {
	page := c.QueryParam("page")
	pageInt := convertToInt(page)

	limit := c.QueryParam("limit")
	limitInt := convertToInt(limit)

	asc := c.QueryParam("asc")

	playlists, _ := h.PlaylistRepo.FindAll(pageInt, limitInt, asc)

	return c.JSON(http.StatusOK, playlists)
}

func (h *PlaylistHandler) Store(c echo.Context) error {
	var playlistDTO PlaylistDTO

	if err := c.Bind(&playlistDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	playlist, err := entity.CreatePlaylist(playlistDTO.Title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.PlaylistRepo.Create(playlist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, playlist)
}

func convertToInt(str string) int {
	pageInt, err := strconv.Atoi(str)
	if err != nil {
		pageInt = 0
	}

	return pageInt
}
