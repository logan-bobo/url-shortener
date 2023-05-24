package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedirectURL struct {
	URL string `json:"url"`
}

func readUrlKeys(c *gin.Context) {
	c.JSON(200, gin.H{
		"redirect": "https//www.google.com",
	})
}

func createUrlKey(c *gin.Context) {
	var newRedirectURL RedirectURL

	if err := c.BindJSON(&newRedirectURL); err != nil {
		return
	}

	fmt.Println(newRedirectURL)
	c.IndentedJSON(http.StatusCreated, newRedirectURL)

	// Generate a key based on a hash of the url and wrte to DB this means that each URl is unique in the DB and a
	// URL will always output the same hash
}

func main() {
	r := gin.Default()

	// Read route for urlkeys
	r.GET("/api/v1/urlkeys", readUrlKeys)

	// Create route for url keys
	r.POST("/api/v1/urlkeys", createUrlKey)

	r.Run() // listen and serve on 0.0.0.0:8080

}
