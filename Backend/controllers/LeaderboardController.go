package controllers

import (
	"BACKEND/Data"
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLeaderboard(c *gin.Context) {
	userID := c.Query("user_id")

	collection := Data.GetCollection("SkillArcade", "UserScores")
	toprankers, userRank, err := services.GetLeaderboardService(c.Request.Context(), collection, userID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusOK, gin.H{"message": "User not found"})
			return
		}
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

func LeaderboardRouter(router *gin.RouterGroup) {
	router.GET("/leaderboard", GetLeaderboard)
}
