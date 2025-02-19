package middlewares

import (
	"BACKEND/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// JWTMiddleware checks for a valid JWT in Authorization header
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate JWT
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store user info in context for further use
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

// CORSMiddleware dynamically sets CORS headers based on allowed origins
func CORSMiddleware() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	allowedOrigins := strings.Split(os.Getenv("CORS_ORIGINS"), ",") // Read allowed origins from .env

	return func(c *gin.Context) {
		requestOrigin := c.Request.Header.Get("Origin") // Get request's origin

		// Check if the request origin is in the allowed list
		isAllowed := false
		for _, origin := range allowedOrigins {
			if strings.TrimSpace(origin) == requestOrigin {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			// Set CORS Headers dynamically
			c.Writer.Header().Set("Access-Control-Allow-Origin", requestOrigin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Handle OPTIONS method separately for preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
