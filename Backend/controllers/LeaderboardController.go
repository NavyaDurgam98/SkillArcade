package controllers

import (
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLeaderboard(c *gin.Context) {
	// Check if user_id is present as a query parameter
	userID := c.Query("user_id")

	// Call the service function to get the leaderboard data
	top10, userRank, err := services.GetLeaderboardService(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the appropriate response based on the presence of userRank
	if userRank != nil {
		// Return the user rank directly as a single JSON object
		c.JSON(http.StatusOK, userRank)
	} else {
		// Return leaderboard data directly as an array
		c.JSON(http.StatusOK, top10)
	}
}

func LeaderboardRouter(router *gin.Engine) {
	router.GET("/leaderboard", GetLeaderboard)
}
