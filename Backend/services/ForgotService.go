package services

import (
	"BACKEND/Data"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type requestData struct {
	Email string `json:"email" bindind:"required"`
}

func ForgotPasswordService(c context.Context, requestData requestData) (string, error) {
	// Access MongoDB collection
	userDetailsCollection := Data.GetCollection("SkillArcade", "UserDetails")

	// check if user exists in DB
	var userExists bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"email": requestData.Email}).Decode(&userExists)
	if err != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return "", errors.New("Email not found")
	}

	//static token for temporary use
	resetToken := "reset_token_123456"

	return resetToken, nil
}
