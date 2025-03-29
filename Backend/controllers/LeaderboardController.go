package controllers

import (
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLeaderboard(c *gin.Context) {
	userID := c.Query("user_id")

	toprankers, userRank, err := services.GetLeaderboardService(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the appropriate response based on the presence of userRank
	if userRank != nil {
		c.JSON(http.StatusOK, userRank)
	} else {
		c.JSON(http.StatusOK, toprankers)
	}
}

func LeaderboardRouter(router *gin.Engine) {
	router.GET("/leaderboard", GetLeaderboard)
}
