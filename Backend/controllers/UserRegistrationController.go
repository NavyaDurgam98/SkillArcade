package controllers

import (
	"BACKEND/Data"
	"BACKEND/models"
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userData models.UserRegister
	// Bind JSON request to userData struct
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	collection := Data.GetCollection("SkillArcade", "UserDetails")
	//connect to service to validate in db
	_, err := services.UserRegistrationService(c.Request.Context(), collection, &userData)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func UserRegisterRouter(router *gin.Engine) {
	router.POST("/signup", UserRegister)
}
