package controllers

import (
	"BACKEND/services"

	"github.com/gin-gonic/gin"
)

// SearchRouter registers the search API route
func SearchRouter(r *gin.Engine) {
	r.GET("/search", SearchQuizzesHandler)
}

// SearchQuizzesHandler handles the search API
func SearchQuizzesHandler(c *gin.Context) {
	searchText := c.Query("searchText")
	page := c.Query("page")

	if len(searchText) < 3 {
		c.JSON(400, gin.H{"error": "Search text must be at least 3 characters"})
		return
	}

	results, err := services.SearchQuizzes(searchText, page)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to search quizzes"})
		return
	}

	c.JSON(200, gin.H{"results": results})
}
