package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"SKILLARCADE/services"
	"SKILLARCADE/models"
)

func UserLogin(c *gin.Context) {

	var loginData models.UserLogin
	// Extract username & password from JSON request body and bind
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//connect to service to validate in db
	_, err := services.UserLoginService(c.Request.Context(),&loginData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}

func UserLoginRouter(router *gin.Engine) {
	router.POST("/signin", UserLogin)
}