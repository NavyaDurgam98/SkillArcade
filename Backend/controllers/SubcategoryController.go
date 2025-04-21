package controllers

import (
	"BACKEND/Data"
	"BACKEND/services"

	// "BACKEND/Data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSubCategories(c *gin.Context) {
	collection := Data.GetCollection("SkillArcade", "Quizzes")
	categoryName := c.Param("category")
	searchText := c.Query("searchText") // optional

	subCategories, err := services.FetchSubCategories(c, categoryName, searchText, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subCategories)
}
func SubCategoryRouter(r *gin.RouterGroup) {
	r.GET("/categories/:category", GetSubCategories)
}
