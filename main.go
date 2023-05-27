package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type savedURL struct {
	ID        string `gorm:"primaryKey"`
	URL       string
	CreatedAt time.Time
}

type redirectURL struct {
	URL string `json:"URL"`
}

func readURLKeys(c *gin.Context) {
	c.JSON(200, gin.H{
		"redirect": "https//www.google.com",
	})
}

func createURLKey(c *gin.Context) {
	var newRedirectURL redirectURL

	if err := c.BindJSON(&newRedirectURL); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	c.JSON(201, newRedirectURL)

	// Generate a key based on a hash of the url and wrte to DB this means that each URl is unique in the DB and a
	// URL will always output the same hash

}

func main() {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_PORT := os.Getenv("DB_PORT")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&savedURL{})

	if err != nil {
		return
	}

	// create an instance of the gin engine
	r := gin.Default()

	// Read route for urlkeys
	r.GET("/api/v1/urlkeys", readURLKeys)

	// Create route for url keys
	r.POST("/api/v1/urlkeys", createURLKey)

	r.Run() // listen and serve on 0.0.0.0:8080

}
