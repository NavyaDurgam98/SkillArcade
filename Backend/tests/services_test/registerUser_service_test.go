package services_test

import (
	"BACKEND/models"
	"BACKEND/services"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang.org/x/crypto/bcrypt"
)

// A variable to hold the bcrypt hash function (we'll mock this in tests)
var bcryptGenerateFromPassword = bcrypt.GenerateFromPassword

// Test function to mock bcrypt password hashing
func mockBcryptGenerateFromPassword(password []byte, cost int) ([]byte, error) {
	// This is a mock, return a predefined value
	return []byte("mockedhashedpassword"), nil
}

func TestUserRegistrationService(t *testing.T) {
	// Initialize the mock MongoDB test environment
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		// Mock user registration data
		mockUser := &models.UserRegister{
			FirstName: "John",
			LastName:  "Doe",
			Username:  "john_doe",
			Password:  "password123",
			Email:     "john.doe@example.com",
			DOB:       "1990-01-01",
			Gender:    "Male",
		}

		// Simulate that the email doesn't exist in DB
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "SkillArcade.UserDetails", mtest.FirstBatch))

		// Mock bcrypt to simulate password hashing by overriding the global variable
		bcryptGenerateFromPassword = mockBcryptGenerateFromPassword
		defer func() {
			// Restore the original bcrypt function after the test
			bcryptGenerateFromPassword = bcrypt.GenerateFromPassword
		}()

		// Mock a successful insertOne operation
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		// Call the service with the mocked bcrypt and the mocked collection
		firstName, err := services.UserRegistrationService(context.Background(), mt.Coll, mockUser)

		// Assert the results
		assert.NoError(t, err)
		assert.Equal(t, mockUser.FirstName, firstName)
	})

	mt.Run("email_already_exists", func(mt *mtest.T) {
		// Mock user data
		mockUser := &models.UserRegister{
			FirstName: "Jane",
			LastName:  "Doe",
			Username:  "jane_doe",
			Password:  "password123",
			Email:     "jane.doe@example.com",
			DOB:       "1992-02-02",
			Gender:    "Female",
		}

		// Simulate the case where the email already exists in the DB
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "SkillArcade.UserDetails", mtest.FirstBatch,
			bson.D{
				{Key: "email", Value: "jane.doe@example.com"},
			},
		))

		// Mock bcrypt to simulate password hashing
		bcryptGenerateFromPassword = mockBcryptGenerateFromPassword
		defer func() {
			// Restore the original bcrypt function after the test
			bcryptGenerateFromPassword = bcrypt.GenerateFromPassword
		}()

		// Call the service
		firstName, err := services.UserRegistrationService(context.Background(), mt.Coll, mockUser)

		// Assert that it returns an error
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email already exists")
		assert.Equal(t, mockUser.FirstName, firstName)
	})

	mt.Run("failed_to_hash_password", func(mt *mtest.T) {
		// Mock user data
		mockUser := &models.UserRegister{
			FirstName: "Mark",
			LastName:  "Smith",
			Username:  "mark_smith",
			Password:  "password123",
			Email:     "mark.smith@example.com",
			DOB:       "1995-03-03",
			Gender:    "Male",
		}

		// Simulate an error in bcrypt password hashing
		bcryptGenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
			return nil, errors.New("bcrypt error")
		}
		defer func() {
			// Restore the original bcrypt function after the test
			bcryptGenerateFromPassword = bcrypt.GenerateFromPassword
		}()

		// Call the service
		firstName, err := services.UserRegistrationService(context.Background(), mt.Coll, mockUser)

		// Assert that the error is as expected
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error creating user")
		assert.Equal(t, mockUser.FirstName, firstName)
	})

	mt.Run("insert_error", func(mt *mtest.T) {
		// Mock user data
		mockUser := &models.UserRegister{
			FirstName: "Alice",
			LastName:  "Johnson",
			Username:  "alice_johnson",
			Password:  "password123",
			Email:     "alice.johnson@example.com",
			DOB:       "1998-04-04",
			Gender:    "Female",
		}

		// Simulate an error in inserting the user into the DB
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "insert error",
		}))

		// Mock bcrypt to simulate password hashing
		bcryptGenerateFromPassword = mockBcryptGenerateFromPassword
		defer func() {
			// Restore the original bcrypt function after the test
			bcryptGenerateFromPassword = bcrypt.GenerateFromPassword
		}()

		// Call the service
		firstName, err := services.UserRegistrationService(context.Background(), mt.Coll, mockUser)

		// Assert that the error is as expected
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error creating user")
		assert.Equal(t, mockUser.FirstName, firstName)
	})
}
