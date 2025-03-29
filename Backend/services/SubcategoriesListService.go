package services

import (
	// "BACKEND/Data"
	"BACKEND/models"
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchSubCategories(c context.Context, categoryName, searchText string, collection *mongo.Collection) ([]map[string]string, error) {
	categoryName = strings.TrimSpace(categoryName)
	searchText = strings.TrimSpace(searchText)

	filter := bson.M{"category": categoryName}

	var category models.Category
	err := collection.FindOne(c, filter).Decode(&category)
	if err != nil {
		return nil, fmt.Errorf("error finding category: %v", err)
	}

	var result []map[string]string
	for _, subCategory := range category.SubCategories {
		if len(searchText) >= 3 {
			// case-insensitive contains match
			if !strings.Contains(strings.ToLower(subCategory.SubCategoryName), strings.ToLower(searchText)) {
				continue
			}
		}
		result = append(result, map[string]string{
			"subCategory": subCategory.SubCategoryName,
			"subImgPath":  subCategory.SubImgPath,
		})
	}

	return result, nil
}
