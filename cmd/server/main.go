package main

import (
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/domain/entity"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/internal/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Playlist{},
		&entity.Video{},
	)
	if err != nil {
		return
	}

	http.Init(db)
}
