package controllers

import (
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(c *gin.Context) {
	var requestData struct {
		Email string `json:"email" bindind:"required"`
	}

	// Bind the email address from the request body
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//connect to service to validate in db
	resetToken, err := services.ForgotPasswordService(c.Request.Context(), requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Simulate sending an email with the reset token (Here we're just sending it in the response)
	// In production, you'd send the token via email to the user.
	c.JSON(http.StatusOK, gin.H{
		"message":     "Password reset link sent.",
		"reset_token": resetToken,
		"email":       requestData.Email,
	})

}

func ForgotRouter(router *gin.Engine) {
	router.POST("/forgotpassword", ForgotPassword)
}
