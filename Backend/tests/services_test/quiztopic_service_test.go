package services_test

import (
	"BACKEND/models"
	"BACKEND/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFetchQuizTopics(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success_with_results", func(mt *mtest.T) {
		// Correctly create a cursor response with multiple quiz topics
		mockCategory := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "category", Value: "Computer Science"},
				{Key: "sub_categories", Value: bson.A{
					bson.D{
						{Key: "sub_category", Value: "Programming Languages"},
						{Key: "quiz_topics", Value: bson.A{
							bson.D{{Key: "quiz_topic_id", Value: "1"}, {Key: "quiz_topic_name", Value: "Python"}},
							bson.D{{Key: "quiz_topic_id", Value: "2"}, {Key: "quiz_topic_name", Value: "Java"}},
						}},
					},
				}},
			},
		)

		mt.AddMockResponses(mockCategory)

		expectedTopics := []models.QuizTopic{
			{QuizTopicID: "1", QuizTopicName: "Python"},
			{QuizTopicID: "2", QuizTopicName: "Java"},
		}

		result, err := services.FetchQuizTopics(context.Background(), "Computer Science", "Programming Languages", "", mt.Coll)
		assert.NoError(t, err)
		assert.Equal(t, expectedTopics, result)
	})

	mt.Run("success_no_results", func(mt *mtest.T) {
		// Mock empty cursor response
		emptyCursor := mtest.CreateCursorResponse(0, "dbname.collection", mtest.FirstBatch)
		mt.AddMockResponses(emptyCursor)

		expectedTopics := []models.QuizTopic(nil)

		result, err := services.FetchQuizTopics(context.Background(), "Math", "Algebra", "", mt.Coll)
		assert.Error(t, err)
		assert.Equal(t, expectedTopics, result)
	})

	mt.Run("search_with_filtered_results", func(mt *mtest.T) {
		// Mock filtered topics response
		mockCategory := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "category", Value: "Computer Science"},
				{Key: "sub_categories", Value: bson.A{
					bson.D{
						{Key: "sub_category", Value: "Programming Languages"},
						{Key: "quiz_topics", Value: bson.A{
							bson.D{{Key: "quiz_topic_id", Value: "1"}, {Key: "quiz_topic_name", Value: "Python"}},
							bson.D{{Key: "quiz_topic_id", Value: "2"}, {Key: "quiz_topic_name", Value: "Java"}},
						}},
					},
				}},
			},
		)

		mt.AddMockResponses(mockCategory)

		expectedTopics := []models.QuizTopic{
			{QuizTopicID: "1", QuizTopicName: "Python"},
		}

		result, err := services.FetchQuizTopics(context.Background(), "Computer Science", "Programming Languages", "Pyt", mt.Coll)
		assert.NoError(t, err)
		assert.Equal(t, expectedTopics, result)
	})

	mt.Run("subcategory_not_found", func(mt *mtest.T) {
		// Mock valid category but missing subcategory
		mockCategory := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "category", Value: "Computer Science"},
				{Key: "sub_categories", Value: bson.A{
					bson.D{
						{Key: "sub_category", Value: "Data Structures"},
						{Key: "quiz_topics", Value: bson.A{
							bson.D{{Key: "quiz_topic_id", Value: "3"}, {Key: "quiz_topic_name", Value: "Arrays"}},
						}},
					},
				}},
			},
		)

		mt.AddMockResponses(mockCategory)

		result, err := services.FetchQuizTopics(context.Background(), "Computer Science", "Programming Languages", "", mt.Coll)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "subcategory 'Programming Languages' not found under category 'Computer Science'")
		assert.Nil(t, result)
	})

	mt.Run("database_error", func(mt *mtest.T) {
		// Simulate a database error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "database error",
		}))

		result, err := services.FetchQuizTopics(context.Background(), "Computer Science", "Programming Languages", "", mt.Coll)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error finding category or subcategory")
	})
}
