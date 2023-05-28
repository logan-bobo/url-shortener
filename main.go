package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB connection that can be used across all route functions
var DB *gorm.DB

// Instantiate a DB connection via a connection string
func connectDB(connectionString string) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = db
}

// Model for use with GORM
type savedURL struct {
	ID        int32 `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	URL       string
	CreatedAt time.Time
}

// Representation of JSON expected to be used with POST request to the /api/v1/urlkeys endpont
type redirectURL struct {
	URL string `json:"URL"`
}

// Dummy return to front end for now at /api/v1/urlkeys on GET request
func readURLKeys(c *gin.Context) {
	c.JSON(200, gin.H{
		"redirect": "https//www.google.com",
	})
}

// Cteate a key to url mapping at the database level return the key and URL
func createURLKey(c *gin.Context) {
	var newRedirectURL redirectURL

	if err := c.BindJSON(&newRedirectURL); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	url := savedURL{
		URL:       newRedirectURL.URL,
		CreatedAt: time.Now(),
	}

	result := DB.Create(&url)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(201, gin.H{
		"URL": url.URL,
		"KEY": url.ID,
	})

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

	connectDB(dsn)

	err := DB.AutoMigrate(&savedURL{})

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
