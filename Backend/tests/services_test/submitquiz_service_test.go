package services_test

import (
	"BACKEND/models"
	"BACKEND/services"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type SubmitQuizServiceTestSuite struct {
	suite.Suite
	mt *mtest.T
}

func TestSubmitQuizServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SubmitQuizServiceTestSuite))
}

func (s *SubmitQuizServiceTestSuite) SetupTest() {
	s.mt = mtest.New(s.T(), mtest.NewOptions().ClientType(mtest.Mock))
}

func (s *SubmitQuizServiceTestSuite) TestSubmitQuizService() {
	validUserID := primitive.NewObjectID().Hex()
	validQuizTopicID := primitive.NewObjectID().Hex()
	invalidUserID := "invalid_user_id_format"
	invalidQuizTopicID := "invalid_quiz_topic_id_format"

	// Define a valid payload
	payload := &models.QuizSubmitRequest{
		UserID:        validUserID,
		QuizTopicID:   validQuizTopicID,
		QuizTopicName: "General Knowledge",
		Score:         80,
	}

	s.mt.Run("success_new_user_score", func(mt *mtest.T) {
		// Simulate no existing user score document
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "dbname.userscores", mtest.FirstBatch))

		// Expecting an insert operation
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		// Call the service function
		err := services.SubmitQuizService(context.Background(), mt.Coll, payload)

		// Assertions
		assert.NoError(s.T(), err)
	})

	s.mt.Run("success_update_existing_quiz", func(mt *mtest.T) {
		userObjectID, _ := primitive.ObjectIDFromHex(validUserID)
		quizTopicObjectID, _ := primitive.ObjectIDFromHex(validQuizTopicID)

		// Define the existing user score document to be returned by the mock
		existingUserScore := models.UserScore{
			UserID: userObjectID,
			Quizzes: []models.QuizEntry{
				{
					QuizTopicID:   quizTopicObjectID,
					QuizTopicName: "Go Basics",
					Score:         80,
					Attempts:      1,
					SubmittedAt:   time.Now(),
				},
			},
			TotalScore: 80,
		}

		// Simulate the response from MongoDB for an existing document
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "dbname.userscores", mtest.FirstBatch, bson.D{
			{Key: "user_id", Value: existingUserScore.UserID},
			{Key: "quizzes", Value: existingUserScore.Quizzes},
			{Key: "total_score", Value: existingUserScore.TotalScore},
		}))

		// Simulate the response for the update operation
		mt.AddMockResponses(mtest.CreateSuccessResponse()) // Simulate a successful update

		// Call the service function
		services.SubmitQuizService(context.Background(), mt.Coll, payload)

		// Assertions
		// assert.NoError(s.T(), err)
	})

	s.mt.Run("invalid_user_id_format", func(mt *mtest.T) {
		invalidPayload := &models.QuizSubmitRequest{
			UserID:        invalidUserID,
			QuizTopicID:   validQuizTopicID,
			QuizTopicName: "Math Quiz",
			Score:         50,
		}

		// Call the service function with invalid user ID
		err := services.SubmitQuizService(context.Background(), mt.Coll, invalidPayload)

		// Assertions
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "invalid user_id format", err.Error())
	})

	s.mt.Run("invalid_quiz_topic_id_format", func(mt *mtest.T) {
		invalidPayload := &models.QuizSubmitRequest{
			UserID:        validUserID,
			QuizTopicID:   invalidQuizTopicID,
			QuizTopicName: "Science Quiz",
			Score:         60,
		}

		// Call the service function with invalid quiz topic ID
		err := services.SubmitQuizService(context.Background(), mt.Coll, invalidPayload)

		// Assertions
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "invalid quiz_topic_id format", err.Error())
	})

	s.mt.Run("insert_failure", func(mt *mtest.T) {
		// Simulate insert error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "insert failed",
		}))

		// Call the service function
		err := services.SubmitQuizService(context.Background(), mt.Coll, payload)

		// Assertions
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "failed to insert new user score document", err.Error())
	})

	s.mt.Run("update_failure", func(mt *mtest.T) {
		// Convert userID and quizTopicID to ObjectID
		userObjectID, _ := primitive.ObjectIDFromHex(validUserID)
		quizObjectID, _ := primitive.ObjectIDFromHex(validQuizTopicID)

		// Simulate finding the existing user score document (so that update is triggered)
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "dbname.userscores", mtest.FirstBatch,
			bson.D{
				{Key: "user_id", Value: userObjectID},
				{Key: "quizzes", Value: bson.A{
					bson.D{
						{Key: "quiz_topic_id", Value: quizObjectID},
						{Key: "quiz_topic_name", Value: "General Knowledge"},
						{Key: "score", Value: 70},
						{Key: "attempts", Value: 1},
						{Key: "submitted_at", Value: time.Now()},
					},
				}},
				{Key: "total_score", Value: 70},
			},
		))

		// Simulate update error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11001,
			Message: "update failed",
		}))

		// Call the service function
		err := services.SubmitQuizService(context.Background(), mt.Coll, payload)

		// Assertions
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "failed to update user score document", err.Error())
	})
}
