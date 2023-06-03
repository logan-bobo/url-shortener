package controllers

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/logan-bobo/url_shortener/db"
	"github.com/logan-bobo/url_shortener/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = db.GetDB()
}

// Representation of JSON expected to be used with POST request to the /api/v1/urlkeys endpont {"URL": "www.example.com"}
type RedirectURL struct {
	URL string `json:"URL" binding:"required"`
}

// Read the URL to redirect to from a given key
func ReadURLKey(c *gin.Context) {
	urlID := c.Param("id")

	urlIDI32, err := strconv.ParseInt(urlID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request please follow the documented format of /api/v1/urlkeys/1",
		})

		return
	}

	var urlInstance = models.SavedURL{ID: int32(urlIDI32)}

	exists := DB.First(&urlInstance)
	if exists.Error != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Given key does not exist %s", urlID),
		})

		return
	}

	c.JSON(200, gin.H{
		"redirect": urlInstance.URL,
	})
}

// Create a key to url mapping at the database level return the key and URL
func CreateURLKey(c *gin.Context) {
	var newRedirectURL RedirectURL

	// Request cant be bound to struct redirectURL
	if err := c.BindJSON(&newRedirectURL); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request, please ensure POST reqeusts to this endpoint match the required JSON structure",
		})

		return
	}

	// validate url is valid before we write to database
	_, err := url.ParseRequestURI(newRedirectURL.URL)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request, URL is invalid",
		})

		return
	}

	newSavedURL := models.SavedURL{
		URL:       newRedirectURL.URL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := DB.Create(&newSavedURL)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, gin.H{
		"KEY": newSavedURL.ID,
		"URL": newSavedURL.URL,
	})
}

// Delete a url key pair
func DeleteURLkey(c *gin.Context) {
	urlID := c.Param("id")

	urlIDI32, err := strconv.ParseInt(urlID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request please follow the docuemnted format of /api/v1/urlkeys/1",
		})

		return
	}

	var urlInstance = models.SavedURL{ID: int32(urlIDI32)}

	exists := DB.First(&urlInstance)
	if exists.Error != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("url does not exist with the given key %s", urlID),
		})

		return
	}

	result := DB.Delete(&urlInstance)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}

// Update a url that exists at a given key
func UpdateURLKey(c *gin.Context) {
	var updateRedirectURL RedirectURL

	// Request cant be bound to struct redirectURL
	if err := c.BindJSON(&updateRedirectURL); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request, please ensure PUT reqeusts to this endpoint match the required JSON structure",
		})

		return
	}

	// validate url is valid before we write to database
	_, err := url.ParseRequestURI(updateRedirectURL.URL)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request, URL is invalid",
		})

		return
	}

	urlID := c.Param("id")

	urlIDI32, err := strconv.ParseInt(urlID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request please follow the documented format of /api/v1/urlkeys/1",
		})

		return
	}

	var urlInstance = models.SavedURL{
		ID:        int32(urlIDI32),
		URL:       updateRedirectURL.URL,
		UpdatedAt: time.Now(),
	}

	exists := DB.First(&urlInstance)
	if exists.Error != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("Given key does not exist %s", urlID),
		})

		return
	}

	result := DB.Save(&urlInstance)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, gin.H{
		"KEY": urlInstance.ID,
		"URL": urlInstance.URL,
	})
}
