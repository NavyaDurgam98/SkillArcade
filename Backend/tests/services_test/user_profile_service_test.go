package services_test

import (
	"BACKEND/models"
	"BACKEND/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type UserProfileServiceTestSuite struct {
	suite.Suite
	mt *mtest.T
}

func TestUserProfileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserProfileServiceTestSuite))
}

func (s *UserProfileServiceTestSuite) SetupTest() {
	s.mt = mtest.New(s.T(), mtest.NewOptions().ClientType(mtest.Mock))
}

func (s *UserProfileServiceTestSuite) TestGetUserProfileService() {
	// Generate a valid ObjectID for testing
	validUserID := primitive.NewObjectID().Hex()
	nonExistentUserID := primitive.NewObjectID().Hex() // Valid but not present in DB

	s.mt.Run("success", func(mt *mtest.T) {
		// Create a mock user profile
		expectedUserProfile := &models.UserProfile{
			Username:  "john_doe",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		}

		// Add mock response to return the user profile
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "dbname.users", mtest.FirstBatch, bson.D{
			{Key: "username", Value: expectedUserProfile.Username},
			{Key: "firstname", Value: expectedUserProfile.FirstName},
			{Key: "lastname", Value: expectedUserProfile.LastName},
			{Key: "email", Value: expectedUserProfile.Email},
		}))

		// Call the service function
		result, err := services.GetUserProfileService(context.Background(), mt.Coll, validUserID)

		// Assertions
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), *expectedUserProfile, result)
	})

	s.mt.Run("user_not_found", func(mt *mtest.T) {
		// Simulate no documents found for a valid ObjectID
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "dbname.users", mtest.FirstBatch))

		// Call the service function with a valid but non-existent user ID
		result, err := services.GetUserProfileService(context.Background(), mt.Coll, nonExistentUserID)

		// Assertions
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "user not found", err.Error())
		assert.Equal(s.T(), models.UserProfile{}, result) // Ensure the result is an empty profile
	})

	s.mt.Run("invalid_user_id_format", func(mt *mtest.T) {
		// Simulate invalid user ID format error
		invalidUserID := "invalid_object_id_format" // Invalid ObjectID
		result, err := services.GetUserProfileService(context.Background(), mt.Coll, invalidUserID)

		// Assertions
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "invalid user ID format", err.Error())
		assert.Equal(s.T(), models.UserProfile{}, result)
	})
}
