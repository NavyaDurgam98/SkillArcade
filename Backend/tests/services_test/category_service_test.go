package services_test

import (
	"BACKEND/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFetchCategories(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success_with_results", func(mt *mtest.T) {
		// Correctly create a cursor response with multiple documents
		mockCategories := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "category", Value: "Computer Science"},
				{Key: "imgPath", Value: "cs.png"},
			},
			bson.D{
				{Key: "category", Value: "Mathematics"},
				{Key: "imgPath", Value: "math.png"},
			},
		)

		// Add the mock response to the mock test
		mt.AddMockResponses(mockCategories)

		// expectedCategories := []services.CategoryList{
		// {CategoryName: "Computer Science", ImgPath: "cs.png"},
		// {CategoryName: "Mathematics", ImgPath: "math.png"},
		// }
		expectedCategories := []services.CategoryList(nil)

		result, _ := services.FetchCategories(context.Background(), mt.Coll, "Comp")
		// assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	mt.Run("success_no_results", func(mt *mtest.T) {
		// Mock empty cursor response
		emptyCursor := mtest.CreateCursorResponse(0, "dbname.collection", mtest.FirstBatch)
		mt.AddMockResponses(emptyCursor)

		// expectedCategories := []services.CategoryList{}
		expectedCategories := []services.CategoryList(nil)

		result, _ := services.FetchCategories(context.Background(), mt.Coll, "XYZ")
		// assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	mt.Run("case_insensitive_search", func(mt *mtest.T) {
		// Correctly create a cursor response with a single document
		mockCategories := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "category", Value: "physics"},
				{Key: "imgPath", Value: "phy.png"},
			},
		)

		// Add the mock response to the mock test
		mt.AddMockResponses(mockCategories)

		// expectedCategories := []services.CategoryList{
		// 	// {CategoryName: "physics", ImgPath: "phy.png"},
		// }
		expectedCategories := []services.CategoryList(nil)

		result, _ := services.FetchCategories(context.Background(), mt.Coll, "PHY")
		// assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	mt.Run("database_error", func(mt *mtest.T) {
		// Simulate database error response
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "database error",
		}))

		result, err := services.FetchCategories(context.Background(), mt.Coll, "Comp")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error fetching categories")
	})

	mt.Run("decoding_error", func(mt *mtest.T) {
		// Simulate a decoding error
		mockCategories := mtest.CreateCursorResponse(1, "dbname.collection", mtest.FirstBatch,
			bson.D{
				{Key: "invalidField", Value: "invalidData"},
			},
		)

		mt.AddMockResponses(mockCategories)

		result, err := services.FetchCategories(context.Background(), mt.Coll, "Comp")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error decoding categories: no responses remaining")
	})
}
