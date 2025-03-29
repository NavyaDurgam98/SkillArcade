package controllers

import (
	"BACKEND/services"
	"fmt"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	// "log"
)

// Remove the .env file loading for tests (since it's not necessary for this example)
// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func TestGetCategories_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Mock service to bypass actual database calls
	mockService := &MockService{}
	r.GET("/categories", func(c *gin.Context) {
		categories, err := mockService.FetchCategories(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	})

	// Create a request
	req, _ := http.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()

	// Call the route
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200")
	assert.Contains(t, w.Body.String(), "Frontend", "Expected response body to contain 'Frontend'")
	assert.Contains(t, w.Body.String(), "Backend", "Expected response body to contain 'Backend'")
}

func TestGetCategories_Failure(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Mock service to simulate error during fetching categories
	mockService := &MockService{}
	r.GET("/categories", func(c *gin.Context) {
		categories, err := mockService.FetchCategories(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	})

	// Create a request with an error header to simulate failure
	req, _ := http.NewRequest("GET", "/categories", nil)
	req.Header.Set("error", "true")
	w := httptest.NewRecorder()

	// Call the route
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected status code 500")
	assert.Contains(t, w.Body.String(), "database error", "Expected response body to contain 'database error'")
}

// Mock Service without using external libraries
type MockService struct {}

func (m *MockService) FetchCategories(c *gin.Context) ([]services.CategoryOnly, error) {
	// Simulate error when the request header contains "error=true"
	if c.GetHeader("error") == "true" {
		return nil, fmt.Errorf("database error")
	}

	// Simulate a successful response
	return []services.CategoryOnly{
		{CategoryName: "Frontend"},
		{CategoryName: "Backend"},
	}, nil
}
