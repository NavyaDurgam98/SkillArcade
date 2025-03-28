package services

import (
	"BACKEND/Data"
	"BACKEND/models"
	"BACKEND/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserLoginService(c context.Context, user *models.UserLogin) (string, string, error) {

	// Access MongoDB collection
	userDetailsCollection := Data.GetCollection("SkillArcade", "UserDetails")

	// check if user exists in DB
	var userExists bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"username": user.Username}).Decode(&userExists)
	if err != nil {
		return user.Username, "", errors.New("user not found")
	}

	// Check if the passwords match
	if userExists["password"] != user.Password {
		return user.Username, "", errors.New("invalid password")
	}

	// Extract ObjectID
	userID := userExists["_id"].(primitive.ObjectID).Hex()

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username, user.Username)
	if err != nil {
		return "", "", err
	}

	return token, userID, nil
}

// package services

// import (
// 	"BACKEND/Data"
// 	"BACKEND/models"
// 	"BACKEND/utils"
// 	"context"
// 	"errors"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// func UserLoginService(c context.Context, user *models.UserLogin) (string, error) {

// 	// Access MongoDB collection
// 	userDetailsCollection := Data.GetCollection("SkillArcade", "UserDetails")

// 	// check if user exists in DB
// 	var userExists bson.M
// 	err := userDetailsCollection.FindOne(c, bson.M{"username": user.Username}).Decode(&userExists)
// 	if err != nil {
// 		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
// 		return user.Username, errors.New("user not found")
// 	}

// 	// Check if the passwords match
// 	if userExists["password"] != user.Password {
// 		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
// 		return user.Username, errors.New("invalid password")
// 	}
// 	// Generate JWT token
// 	token, err := utils.GenerateJWT(user.Username, user.Username)
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// 	//return user.Username, nil
// }
