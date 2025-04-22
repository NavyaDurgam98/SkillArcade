package controllers

import (
	"BACKEND/Data"
	"BACKEND/models"
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {

	var loginData models.UserLogin
	// Extract username & password from JSON request body and bind
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	collection := Data.GetCollection("SkillArcade", "UserDetails")
	// Authenticate user and generate JWT
	token, userID, err := services.UserLoginService(c.Request.Context(), collection, &loginData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return JWT token in response
	c.JSON(http.StatusOK, gin.H{"token": token, "user_id": userID})

	//connect to service to validate in db
	// _, err := services.UserLoginService(c.Request.Context(), &loginData)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}

func UserLoginRouter(router *gin.Engine) {
	router.POST("/signin", UserLogin)
}
