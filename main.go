package main

import (
	"fmt"
	"os"
	"strconv"
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

// Representation of JSON expected to be used with POST request to the /api/v1/urlkeys endpont {"URL": "www.example.com"}
type redirectURL struct {
	URL string `json:"URL"`
}

// Read the URL to redirect to from a given key
func readURLKeys(c *gin.Context) {
	urlID := c.Param("id")

	urlIDI32, err := strconv.ParseInt(urlID, 10, 32)
	if err != nil {
		panic(err)
	}

	var urlInstance = savedURL{ID: int32(urlIDI32)}

	DB.First(&urlInstance)
	c.JSON(200, gin.H{
		"redirect": urlInstance.URL,
	})
}

// Create a key to url mapping at the database level return the key and URL
func createURLKey(c *gin.Context) {
	var newRedirectURL redirectURL

	if err := c.BindJSON(&newRedirectURL); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	newSavedURL := savedURL{
		URL:       newRedirectURL.URL,
		CreatedAt: time.Now(),
	}

	result := DB.Create(&newSavedURL)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(201, gin.H{
		"URL": newSavedURL.URL,
		"KEY": newSavedURL.ID,
	})
}

func main() {
	// Get main db connection paramaters from host OS
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Build connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	// Connect to the database based on the connection string
	connectDB(dsn)

	// Run automatic datbaase migrations
	err := DB.AutoMigrate(&savedURL{})

	if err != nil {
		return
	}

	// Create an instance of the gin engine
	r := gin.Default()

	// Read route for urlkeys
	r.GET("/api/v1/urlkeys/:id", readURLKeys)

	// Create route for url keys
	r.POST("/api/v1/urlkeys", createURLKey)

	// Run the application
	err = r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
