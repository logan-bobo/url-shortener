package server

import (
	"github.com/gin-gonic/gin"
	"github.com/logan-bobo/url_shortener/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Read route for url keys
	router.GET("/api/v1/urlkeys/:id", controllers.ReadURLKey)

	// Create route for url keys
	router.POST("/api/v1/urlkeys", controllers.CreateURLKey)

	// Delete route for url keys
	router.DELETE("/api/v1/urlkeys/:id", controllers.DeleteURLkey)

	// Update route for url keys
	router.PUT("/api/v1/urlkeys/:id", controllers.UpdateURLKey)

	return router
}
