package services_test

import (
	"BACKEND/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFetchSubCategories(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success_with_results", func(mt *mtest.T) {
		// Correctly create a cursor response with multiple subcategories
		mockCategory := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "category", Value: "Computer Science"},
				{Key: "sub_categories", Value: bson.A{
					bson.D{
						{Key: "sub_category", Value: "Programming Languages"},
						{Key: "subImgPath", Value: "prog.png"},
					},
					bson.D{
						{Key: "sub_category", Value: "Data Structures"},
						{Key: "subImgPath", Value: "ds.png"},
					},
				}},
			},
		)

		mt.AddMockResponses(mockCategory)

		expectedSubCategories := []map[string]string{
			{"subCategory": "Programming Languages", "subImgPath": "prog.png"},
			{"subCategory": "Data Structures", "subImgPath": "ds.png"},
		}

		result, err := services.FetchSubCategories(context.Background(), "Computer Science", "", mt.Coll)
		assert.NoError(t, err)
		assert.Equal(t, expectedSubCategories, result)
	})

	mt.Run("success_no_results", func(mt *mtest.T) {
		// Mock the error response when no documents are found
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "dbname.collection", mtest.FirstBatch))

		result, err := services.FetchSubCategories(context.Background(), "Math", "", mt.Coll)
		assert.Error(t, err)                                      // Expect an error because no documents found
		assert.Contains(t, err.Error(), "error finding category") // Check for the specific error message
		assert.Nil(t, result)                                     // The result should be nil
	})

	mt.Run("search_with_filtered_results", func(mt *mtest.T) {
		// Mock filtered subcategories response
		mockCategory := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "category", Value: "Computer Science"},
				{Key: "sub_categories", Value: bson.A{
					bson.D{
						{Key: "sub_category", Value: "Programming Languages"},
						{Key: "subImgPath", Value: "prog.png"},
					},
					bson.D{
						{Key: "sub_category", Value: "Data Structures"},
						{Key: "subImgPath", Value: "ds.png"},
					},
				}},
			},
		)

		mt.AddMockResponses(mockCategory)

		expectedSubCategories := []map[string]string{
			{"subCategory": "Programming Languages", "subImgPath": "prog.png"},
		}

		result, err := services.FetchSubCategories(context.Background(), "Computer Science", "Pro", mt.Coll)
		assert.NoError(t, err)
		assert.Equal(t, expectedSubCategories, result)
	})

	mt.Run("category_not_found", func(mt *mtest.T) {
		// Simulate category not found scenario
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    404,
			Message: "category not found",
		}))

		result, err := services.FetchSubCategories(context.Background(), "Unknown Category", "", mt.Coll)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error finding category")
	})

	mt.Run("database_error", func(mt *mtest.T) {
		// Simulate a database error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "database error",
		}))

		result, err := services.FetchSubCategories(context.Background(), "Computer Science", "", mt.Coll)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error finding category")
	})
}
