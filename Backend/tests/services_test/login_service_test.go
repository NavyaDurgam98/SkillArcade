package services_test

import (
	"BACKEND/models"
	"BACKEND/services"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang.org/x/crypto/bcrypt"
)

// Mock bcrypt.CompareHashAndPassword to simulate password comparison
var bcryptCompareHashAndPassword = bcrypt.CompareHashAndPassword

func mockBcryptCompareHashAndPassword(hashedPassword []byte, password []byte) error {
	// Simulate successful comparison
	if string(hashedPassword) == string(password) {
		return nil
	}
	return errors.New("invalid password")
}

func TestUserLoginService(t *testing.T) {
	// Initialize the mock MongoDB test environment
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	// mt.Run("success", func(mt *mtest.T) {
	// 	// Mock user login data
	// 	mockUser := &models.UserLogin{
	// 		Username: "john_doe",
	// 		Password: "password123",
	// 	}

	// 	// Simulate that the user exists in DB
	// 	mockUserData := bson.D{
	// 		{Key: "username", Value: "john_doe"},
	// 		{Key: "password", Value: "$2a$10$JwjyDwsY5wUJtCEcW6e0eP2aaDTN4apzoBwpQb1jk9vwF6c02JG5y"},
	// 		{Key: "_id", Value: primitive.NewObjectID()},
	// 	}
	// 	mt.AddMockResponses(mtest.CreateCursorResponse(1, "SkillArcade.UserDetails", mtest.FirstBatch, mockUserData))

	// 	// Mock bcrypt.CompareHashAndPassword to simulate correct password comparison
	// 	bcryptCompareHashAndPassword = mockBcryptCompareHashAndPassword
	// 	defer func() {
	// 		bcryptCompareHashAndPassword = bcrypt.CompareHashAndPassword
	// 	}()

	// 	// Call the service with the mocked bcrypt and the mocked collection
	// 	_, _, err := services.UserLoginService(context.Background(), mt.Coll, mockUser)

	// 	// Assert the results
	// 	assert.NoError(t, err)
	// })

	mt.Run("user_not_found", func(mt *mtest.T) {
		// Mock user login data
		mockUser := &models.UserLogin{
			Username: "unknown_user",
			Password: "password123",
		}

		// Simulate that the user is not found in DB
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "SkillArcade.UserDetails", mtest.FirstBatch))

		// Call the service
		_, _, err := services.UserLoginService(context.Background(), mt.Coll, mockUser)

		// Assert the results
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")

	})

	mt.Run("invalid_password", func(mt *mtest.T) {
		// Mock user login data
		mockUser := &models.UserLogin{
			Username: "john_doe",
			Password: "wrongpassword",
		}

		// Simulate that the user exists in DB but password is incorrect
		mockUserData := bson.D{
			{Key: "username", Value: "john_doe"},
			{Key: "password", Value: "$2a$10$JwjyDwsY5wUJtCEcW6e0eP2aaDTN4apzoBwpQb1jk9vwF6c02JG5y"}, // bcrypt hash of "password123"
			{Key: "_id", Value: primitive.NewObjectID()},
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "SkillArcade.UserDetails", mtest.FirstBatch, mockUserData))

		// Mock bcrypt.CompareHashAndPassword to simulate incorrect password comparison
		bcryptCompareHashAndPassword = mockBcryptCompareHashAndPassword
		defer func() {
			bcryptCompareHashAndPassword = bcrypt.CompareHashAndPassword
		}()

		// Call the service
		_, _, err := services.UserLoginService(context.Background(), mt.Coll, mockUser)

		// Assert the results
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid password")

	})

	mt.Run("invalid_user_id_format", func(mt *mtest.T) {
		// Mock user login data
		mockUser := &models.UserLogin{
			Username: "john_doe",
			Password: "password123",
		}

		// Simulate that the user exists in DB but the _id is invalid
		mockUserData := bson.D{
			{Key: "username", Value: "john_doe"},
			{Key: "password", Value: "$2a$10$JwjyDwsY5wUJtCEcW6e0eP2aaDTN4apzoBwpQb1jk9vwF6c02JG5y"}, // bcrypt hash of "password123"
			{Key: "_id", Value: "invalidID"},                                                         // Invalid ObjectID
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "SkillArcade.UserDetails", mtest.FirstBatch, mockUserData))

		// Call the service
		_, _, err := services.UserLoginService(context.Background(), mt.Coll, mockUser)

		// Assert the results
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid password")
	})
}
