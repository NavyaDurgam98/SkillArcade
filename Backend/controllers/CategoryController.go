package controllers

import (
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetCategories(c *gin.Context) {
	categories, err := services.FetchCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
func CategoryRouter(r *gin.Engine) {
	r.GET("/categories", GetCategories)
}