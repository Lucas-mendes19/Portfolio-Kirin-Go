package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

func init() {
	if err != nil {
		panic("failed to connect database")
	}
}
