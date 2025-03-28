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

	subCategories, err := services.FetchSubCategories(c, categoryName, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subCategories)
}

func SubCategoryRouter(r *gin.Engine) {
	r.GET("/categories/:category", GetSubCategories)  
}
