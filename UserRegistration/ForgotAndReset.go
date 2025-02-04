package UserRegistration

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"SKILLARCADE/Data"
)


// ForgotPassword handles sending a password reset token
func ForgotPassword(c *gin.Context) {
	var requestData struct {
		Email string `json:"email"`
	}

	// Bind the email address from the request body
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Access MongoDB collection
	userDetailsCollection := DatabaseConnection.GetCollection("SkillArcade","UserDetails")

	// check if user exists in DB
	var userExists bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"email": requestData.Email}).Decode(&userExists)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return
	}

	//static token for temporary use
	resetToken := "reset_token_123456"

	// Simulate sending an email with the reset token (Here we're just sending it in the response)
	// In production, you'd send the token via email to the user.
	c.JSON(http.StatusOK, gin.H{
		"message":      "Password reset link sent.",
		"reset_token":  resetToken,
		"email":        requestData.Email,
	})
}

// ResetPassword handles updating the user's password
func ResetPassword(c *gin.Context) {
	var resetData struct {
		Email string `json:"email"`
		ResetToken string `json:"reset_token"`
		Password   string `json:"password"`
	}

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

	// Access MongoDB collection
	userDetailsCollection := DatabaseConnection.GetCollection("SkillArcade","UserDetails")

	// Update the password in DB
	updatePassword := bson.M{"$set": bson.M{"password": resetData.Password}}
	_, err := userDetailsCollection.UpdateOne(c, bson.M{"email": resetData.Email}, updatePassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully!"})

}

func ForgotAndResetRoutes(router *gin.Engine) {
	router.POST("/forgotpassword", ForgotPassword) 
	router.POST("/resetpassword", ResetPassword) 
}  
