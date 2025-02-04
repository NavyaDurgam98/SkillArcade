package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"SKILLARCADE/services"
	"SKILLARCADE/models"
)

// ResetPassword handles updating the user's password
func ResetPassword(c *gin.Context) {
	var resetData models.UserReset

	// Bind the reset token and new password from the request body
	if err := c.ShouldBindJSON(&resetData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Simulate validating the reset token (in a real app, you'd store and check the token)
	if resetData.ResetToken != "reset_token_123456" { // Example token validation
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reset token"})
		return
	}

	//connect to service to validate in db
	_, err := services.ResetPasswordService(c.Request.Context(),&resetData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully!"})
}

func ResetRouter(router *gin.Engine) {
	router.POST("/resetpassword", ResetPassword) 
}  