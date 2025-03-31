package controllers

import (
	"BACKEND/Data"
	"BACKEND/services"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetQuizTopics(c *gin.Context) {
	collection := Data.GetCollection("SkillArcade", "Quizzes")

	categoryName := c.Param("category")
	subCategoryName := c.Param("sub_category")
	decodedCategoryName, err := url.QueryUnescape(categoryName)
	if err != nil {
		log.Println("Error decoding category:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode category"})
		return
	}

	decodedSubCategoryName, err := url.QueryUnescape(subCategoryName)
	if err != nil {
		log.Println("Error decoding subcategory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode subcategory"})
		return
	}
	searchText := c.Query("searchText")

	quizTopics, err := services.FetchQuizTopics(c, decodedCategoryName, decodedSubCategoryName, searchText, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"quiz_topics": quizTopics})
}

func QuizTopicRouter(r *gin.Engine) {
	r.GET("/categories/:category/subcategories/:sub_category/quiz_topics", GetQuizTopics)
}
