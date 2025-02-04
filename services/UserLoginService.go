package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"SKILLARCADE/Data"
	"SKILLARCADE/models"
	"errors"
	"context"
)

func UserLoginService(c context.Context,user *models.UserLogin) (string, error) {

	// Access MongoDB collection
	userDetailsCollection := Data.GetCollection("SkillArcade","UserDetails")

	// check if user exists in DB
	var userExists bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"username": user.Username}).Decode(&userExists)
	if err != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return user.Username, errors.New("User not found")
	}

	// Check if the passwords match
	if userExists["password"] != user.Password {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return user.Username, errors.New("Invalid password")
	}
    return user.Username, nil
}