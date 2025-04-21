package services

import (
	"BACKEND/Data"
	"BACKEND/models"
	"BACKEND/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func UserLoginService(c context.Context, user *models.UserLogin) (string, string, error) {

	// Access MongoDB collection
	userDetailsCollection := Data.GetCollection("SkillArcade", "UserDetails")

	// check if user exists in DB
	var userExists bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"username": user.Username}).Decode(&userExists)
	if err != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return user.Username, "", errors.New("user not found")
	}

	// Check if the passwords match
	hashedPassword, ok := userExists["password"].(string)
	if !ok {
		return user.Username, "", errors.New("stored password format is invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return user.Username, "", errors.New("invalid password")
	}

	userID, ok := userExists["_id"].(primitive.ObjectID)

	if !ok {
		return user.Username, "", errors.New("invalid user ID format")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username, user.Username)
	if err != nil {
		return "", "", err
	}

	return token, userID.Hex(), nil
	//return user.Username, nil
}
