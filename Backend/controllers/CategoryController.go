package controllers

import (
	"BACKEND/Data"
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	collection := Data.GetCollection("SkillArcade", "Quizzes")
	searchText := c.Query("searchText")

	categories, err := services.FetchCategories(c, collection, searchText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
func CategoryRouter(r *gin.Engine) {
	r.GET("/categories", GetCategories)
}
