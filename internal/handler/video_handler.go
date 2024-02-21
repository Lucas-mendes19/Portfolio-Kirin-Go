package handler

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type VideoDTO struct {
	PlaylistID  string `json:"playlist_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type VideoHandler struct {
	VideoRepo repository.VideoInterface
}

func NewVideoHandler(db *gorm.DB) *VideoHandler {
	videoRepo := repository.NewVideoRepository(db)

	return &VideoHandler{
		VideoRepo: videoRepo,
	}
}

func (v *VideoHandler) Show(c echo.Context) error {
	id := c.Param("id")

	video, err := v.VideoRepo.Find(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, video)
}

func (v *VideoHandler) Store(c echo.Context) error {
	var videoDTO VideoDTO

	if err := c.Bind(&videoDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	video, err := entity.CreateVideo(
		videoDTO.PlaylistID,
		videoDTO.Title,
		videoDTO.Description,
		videoDTO.URL,
	)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = v.VideoRepo.Create(video)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, video)
}
