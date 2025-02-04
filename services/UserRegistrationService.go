package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"SKILLARCADE/Data"
	"SKILLARCADE/models"
	"errors"
	"context"
)


func UserRegistrationService(c context.Context, user *models.UserRegister) (string,error){
	// Access MongoDB collection
	userDetailsCollection := Data.GetCollection("SkillArcade","UserDetails")

	//Check if email already exists in DB
	var existingUser bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		//c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return user.FirstName, errors.New("Email already exists")
	}

	//create user object to insert in DB
	userData := bson.M{
		"firstname": user.FirstName,
		"lastname": user.LastName,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"dob": user.DOB,
		"gender": user.Gender,
	}

	// Insert user into the MongoDB collection
	_, err = userDetailsCollection.InsertOne(c, userData)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return user.FirstName, errors.New("Error creating user")
	}
    return user.FirstName, nil
}