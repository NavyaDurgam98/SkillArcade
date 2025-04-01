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

type UserHistoryServiceTestSuite struct {
	suite.Suite
	mt *mtest.T
}

func TestUserHistoryServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserHistoryServiceTestSuite))
}

func (s *UserHistoryServiceTestSuite) SetupTest() {
	s.mt = mtest.New(s.T(), mtest.NewOptions().ClientType(mtest.Mock))
}

func (s *UserHistoryServiceTestSuite) TestGetUserHistoryService() {
	validUserID := primitive.NewObjectID().Hex()
	invalidUserID := "invalid_object_id_format"

	// Mocked quiz history data
	mockedQuizHistory := []models.UserHistory{
		{
			QuizTopicName: "General Knowledge",
			Score:         80,
			Attempts:      2,
			SubmittedAt:   time.Now(),
		},
		{
			QuizTopicName: "Science Quiz",
			Score:         90,
			Attempts:      1,
			SubmittedAt:   time.Now(),
		},
	}

	// Success: User history found
	s.mt.Run("success_user_history_found", func(mt *mtest.T) {
		userObjectID, _ := primitive.ObjectIDFromHex(validUserID)

		// Mock response simulating a successful FindOne call
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "dbname.userscores", mtest.FirstBatch,
			bson.D{
				{Key: "user_id", Value: userObjectID},
				{Key: "quizzes", Value: bson.A{
					bson.D{
						{Key: "quiz_topic_name", Value: mockedQuizHistory[0].QuizTopicName},
						{Key: "score", Value: mockedQuizHistory[0].Score},
						{Key: "attempts", Value: mockedQuizHistory[0].Attempts},
						{Key: "submitted_at", Value: mockedQuizHistory[0].SubmittedAt},
					},
					bson.D{
						{Key: "quiz_topic_name", Value: mockedQuizHistory[1].QuizTopicName},
						{Key: "score", Value: mockedQuizHistory[1].Score},
						{Key: "attempts", Value: mockedQuizHistory[1].Attempts},
						{Key: "submitted_at", Value: mockedQuizHistory[1].SubmittedAt},
					},
				}},
			},
		))

		// Call the service function
		result, err := services.GetUserHistoryService(context.Background(), mt.Coll, validUserID)

		// Assertions
		assert.NoError(s.T(), err)
		assert.Len(s.T(), result, 2)
		assert.Equal(s.T(), mockedQuizHistory[0].QuizTopicName, result[0].QuizTopicName)
		assert.Equal(s.T(), mockedQuizHistory[1].QuizTopicName, result[1].QuizTopicName)
		assert.Equal(s.T(), mockedQuizHistory[0].Score, result[0].Score)
		assert.Equal(s.T(), mockedQuizHistory[1].Score, result[1].Score)
	})

	// Failure: Invalid user ID format
	s.mt.Run("invalid_user_id_format", func(mt *mtest.T) {
		// Call the service function with an invalid user ID
		result, err := services.GetUserHistoryService(context.Background(), mt.Coll, invalidUserID)

		// Assertions
		assert.Error(s.T(), err)
		assert.Nil(s.T(), result)
		assert.Equal(s.T(), "invalid user ID format", err.Error())
	})

	// Failure: User not found or no quiz history available
	s.mt.Run("user_not_found_or_no_history", func(mt *mtest.T) {
		// Simulate a scenario where the user history is not found
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "dbname.userscores", mtest.FirstBatch))

		// Call the service function with a non-existent user ID
		result, err := services.GetUserHistoryService(context.Background(), mt.Coll, validUserID)

		// Assertions
		assert.Error(s.T(), err)
		assert.Nil(s.T(), result)
		assert.Equal(s.T(), "user not found or no quiz history available", err.Error())
	})

}
