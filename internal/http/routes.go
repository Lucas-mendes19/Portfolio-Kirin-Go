package http

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func routes(e *echo.Echo, db *gorm.DB) {
	userHandler := handler.NewUserHandler(db)
	playlistHandler := handler.NewPlaylistHandler(db)
	videoHandler := handler.NewVideoHandler(db)

	api := e.Group("/api")
	{
		api.POST("/user", userHandler.Store)

		api.GET("/playlist", playlistHandler.Index)
		api.POST("/playlist", playlistHandler.Store)

		api.POST("/video", videoHandler.Store)
		api.GET("/video/:id", videoHandler.Show)
	}
}
