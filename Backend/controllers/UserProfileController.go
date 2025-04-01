package controllers

import (
	"BACKEND/Data"
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	collection := Data.GetCollection("SkillArcade", "UserDetails")
	// Call service to get the user profile
	profile, err := services.GetUserProfileService(c.Request.Context(), collection, userID)
	if err != nil {
		if err.Error() == "invalid user ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		} else if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, profile)
}

func UserProfileRouter(router *gin.Engine) {
	router.GET("/userprofile", GetUserProfile)
}
