package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter wires handlers and middleware.
func NewRouter(productHandler *ProductHandler, highlightHandler *HighlightHandler, categoryHandler *CategoryHandler) *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(corsMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// Register routes
	productHandler.RegisterRoutes(r)
	highlightHandler.RegisterRoutes(r)
	categoryHandler.RegisterRoutes(r)

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
