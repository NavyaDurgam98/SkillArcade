package services

import (
	"BACKEND/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistrationService(c context.Context, collection *mongo.Collection, user *models.UserRegister) (string, error) {
	// Access MongoDB collection
	userDetailsCollection := collection

	// Check if email already exists in DB
	var existingUser bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		// c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return user.FirstName, errors.New("email already exists")
	}

	//  Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.FirstName, errors.New("failed to hash password")
	}

	// Create user object to insert in DB
	userData := bson.M{
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"username":  user.Username,
		"password":  string(hashedPassword), // 🔐 save the hashed password
		"email":     user.Email,
		"dob":       user.DOB,
		"gender":    user.Gender,
	}

	// Insert user into the MongoDB collection
	_, err = userDetailsCollection.InsertOne(c, userData)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return user.FirstName, errors.New("error creating user")
	}

	return user.FirstName, nil
}
