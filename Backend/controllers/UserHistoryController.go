package controllers

import (
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHistory(c *gin.Context) {
	// Get user_id from query parameters
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Call service function to get user quiz history
	userHistory, err := services.GetUserHistoryService(c.Request.Context(), userID)
	if err != nil {
		if err.Error() == "invalid user ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		} else if err.Error() == "user not found or no quiz history available" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found or no quiz history available"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Return quiz history as JSON
	c.JSON(http.StatusOK, userHistory)
}

func UserHistoryRouter(router *gin.Engine) {
	router.GET("/userhistory", GetUserHistory)
}
