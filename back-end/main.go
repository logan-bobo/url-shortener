package main

import (
	"github.com/logan-bobo/url_shortener/db"
	"github.com/logan-bobo/url_shortener/models"
	"github.com/logan-bobo/url_shortener/server"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	db.Init()

	DB = db.GetDB()

	// Run automatic datbaase migrations
	err := DB.AutoMigrate(&models.SavedURL{})
	if err != nil {
		panic(err)
	}

	// Run the application
	server.Init()
}
