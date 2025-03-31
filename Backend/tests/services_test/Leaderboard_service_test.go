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

type LeaderboardServiceTestSuite struct {
	suite.Suite
	mt *mtest.T
}

func TestLeaderboardServiceTestSuite(t *testing.T) {
	suite.Run(t, new(LeaderboardServiceTestSuite))
}

func (s *LeaderboardServiceTestSuite) SetupTest() {
	s.mt = mtest.New(s.T(), mtest.NewOptions().ClientType(mtest.Mock))
}

func (s *LeaderboardServiceTestSuite) TestGetLeaderboardService() {
	validUserID := primitive.NewObjectID().Hex()
	nonExistentUserID := primitive.NewObjectID().Hex()
	invalidUserID := "invalid_object_id_format"

	s.mt.Run("success_top10", func(mt *mtest.T) {
		// Mock top 10 leaderboard entries
		mockResponses := mtest.CreateCursorResponse(1, "dbname.userscores", mtest.FirstBatch,
			bson.D{
				{Key: "username", Value: "john_doe"},
				{Key: "total_score", Value: 100},
				{Key: "quizzes_taken", Value: 5},
			},
			bson.D{
				{Key: "username", Value: "jane_doe"},
				{Key: "total_score", Value: 90},
				{Key: "quizzes_taken", Value: 4},
			},
		)

		mt.AddMockResponses(mockResponses)

		// Call the service function without user ID to get the top 10
		results, userRank, err := services.GetLeaderboardService(context.Background(), mt.Coll, "")

		// Assertions
		assert.NoError(s.T(), err)
		assert.NotNil(s.T(), results)
		assert.Nil(s.T(), userRank)
		assert.Equal(s.T(), 2, len(results))

		assert.Equal(s.T(), "john_doe", results[0].Username)
		assert.Equal(s.T(), 100, results[0].TotalScore)
		assert.Equal(s.T(), 5, results[0].QuizzesTaken)
		assert.Equal(s.T(), 1, results[0].Rank)

		assert.Equal(s.T(), "jane_doe", results[1].Username)
		assert.Equal(s.T(), 90, results[1].TotalScore)
		assert.Equal(s.T(), 4, results[1].QuizzesTaken)
		assert.Equal(s.T(), 2, results[1].Rank)
	})

	s.mt.Run("success_user_rank", func(mt *mtest.T) {
		userObjectID, err := primitive.ObjectIDFromHex(validUserID)
		assert.NoError(s.T(), err)
		// Mock user rank entry
		mockResponses := mtest.CreateCursorResponse(1, "dbname.userscores", mtest.FirstBatch,
			bson.D{
				{Key: "_id", Value: userObjectID},
				{Key: "username", Value: "john_doe"},
				{Key: "total_score", Value: 100},
				{Key: "quizzes_taken", Value: 5},
			},
		)

		mt.AddMockResponses(mockResponses)

		// Call the service function with a valid user ID
		_, userRank, err := services.GetLeaderboardService(context.Background(), mt.Coll, validUserID)

		// Assertions
		assert.NoError(s.T(), err)
		assert.NotNil(s.T(), userRank)
		assert.Equal(s.T(), "john_doe", userRank.Username)
		assert.Equal(s.T(), 100, userRank.TotalScore)
		assert.Equal(s.T(), 5, userRank.QuizzesTaken)
		assert.Equal(s.T(), 1, userRank.Rank)
	})

	s.mt.Run("user_not_found", func(mt *mtest.T) {
		// Simulate no documents found for a valid ObjectID
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "dbname.userscores", mtest.FirstBatch))

		// Call the service function with a valid but non-existent user ID
		_, userRank, err := services.GetLeaderboardService(context.Background(), mt.Coll, nonExistentUserID)

		// Assertions
		assert.Error(s.T(), err)
		assert.Equal(s.T(), "user not found", err.Error())

		// Check if the returned userRank is an empty LeaderboardEntry instead of nil
		expectedEmptyEntry := &models.LeaderboardEntry{}
		assert.Equal(s.T(), expectedEmptyEntry, userRank)
	})

	s.mt.Run("invalid_user_id_format", func(mt *mtest.T) {
		// Call the service function with an invalid user ID format
		_, userRank, err := services.GetLeaderboardService(context.Background(), mt.Coll, invalidUserID)

		// Assertions
		assert.Error(s.T(), err)
		assert.Nil(s.T(), userRank)
		assert.Equal(s.T(), "invalid user ID format", err.Error())
	})

	s.mt.Run("aggregate_error", func(mt *mtest.T) {
		// Mock an aggregate error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "aggregation failed",
		}))

		// Call the service function
		results, userRank, err := services.GetLeaderboardService(context.Background(), mt.Coll, "")

		// Assertions
		assert.Error(s.T(), err)
		assert.Nil(s.T(), results)
		assert.Nil(s.T(), userRank)
		assert.Equal(s.T(), "failed to aggregate leaderboard", err.Error())
	})
}
