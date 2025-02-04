package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"SKILLARCADE/services"
	"SKILLARCADE/models"
)

func UserRegister(c *gin.Context) {
	var userData models.UserRegister
	// Bind JSON request to userData struct
	if err := c.ShouldBindJSON(&userData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

	//connect to service to validate in db
	_, err := services.UserRegistrationService(c.Request.Context(),&userData)
    if err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusCreated, gin.H{"message":  "User registered successfully!"})
}

func UserRegisterRouter(router *gin.Engine) {
	router.POST("/signup",UserRegister)
}