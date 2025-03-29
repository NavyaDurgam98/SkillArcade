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

func FetchQuizTopics(c context.Context, categoryName, subCategoryName, searchText string, collection *mongo.Collection) ([]models.QuizTopic, error) {
	categoryName = strings.TrimSpace(categoryName)

	subCategoryName = strings.TrimSpace(subCategoryName)
	searchText = strings.TrimSpace(searchText)

	filter := bson.M{
		"category":                    categoryName,
		"sub_categories.sub_category": subCategoryName,
	}

	var category models.Category
	err := collection.FindOne(c, filter).Decode(&category)
	if err != nil {
		return nil, fmt.Errorf("error finding category or subcategory: %v", err)
	}

	for _, subCategory := range category.SubCategories {
		if subCategory.SubCategoryName == subCategoryName {
			if len(searchText) >= 3 {
				var filteredTopics []models.QuizTopic
				for _, topic := range subCategory.QuizTopics {
					if strings.Contains(strings.ToLower(topic.QuizTopicName), strings.ToLower(searchText)) {
						filteredTopics = append(filteredTopics, topic)
					}
				}
				return filteredTopics, nil
			}
			// No filtering
			return subCategory.QuizTopics, nil
		}
	}

	return nil, fmt.Errorf("subcategory '%s' not found under category '%s'", subCategoryName, categoryName)
}
