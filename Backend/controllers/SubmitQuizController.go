package controllers

import (
	"BACKEND/models"
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitQuiz(c *gin.Context) {
	var requestData models.QuizSubmitRequest

	// Bind the JSON payload into the request
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	// Call the service layer to handle the business logic
	err := services.SubmitQuizService(c.Request.Context(), &requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Quiz submitted successfully",
	})
}

func SubmitQuizRouter(router *gin.Engine) {
	router.POST("/submitquiz", SubmitQuiz)
}
