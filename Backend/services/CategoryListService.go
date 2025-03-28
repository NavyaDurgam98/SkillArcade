package services

import (
	// "BACKEND/Data"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryList struct {
	CategoryName string `json:"category" bson:"category"`
	ImgPath      string `json:"imgPath" bson:"imgPath"`
}

func FetchCategories(c context.Context, collection *mongo.Collection, searchText string) ([]CategoryList, error) {
	// Create filter based on searchText
	filter := bson.M{}
	if len(searchText) >= 3 {
		filter = bson.M{
			"category": bson.M{
				"$regex":   searchText,
				"$options": "i", // case-insensitive
			},
		}
	}

	findOptions := options.Find()
	findOptions.SetProjection(bson.M{
		"category": 1,
		"imgPath":  1,
		"_id":      0,
	})

	cursor, err := collection.Find(c, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error fetching categories: %v", err)
	}
	defer cursor.Close(c)

	var categories []CategoryList
	if err = cursor.All(c, &categories); err != nil {
		return nil, fmt.Errorf("error decoding categories: %v", err)
	}

	return categories, nil
}

// func FetchCategories(c context.Context,collection *mongo.Collection) ([]CategoryList, error) {
// 	// collection := Data.GetCollection("SkillArcade", "Quizzes")
// 	filter := bson.M{}
// 	findOptions := options.Find()
// 	findOptions.SetProjection(bson.M{
// 		"category": 1,
// 		"imgPath":  1,
// 		"_id":      0,
// 	})

// 	cursor, err := collection.Find(c, filter, findOptions)
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching categories: %v", err)
// 	}
// 	defer cursor.Close(c)

// 	var categories []CategoryList
// 	if err = cursor.All(c, &categories); err != nil {
// 		return nil, fmt.Errorf("error decoding categories: %v", err)
// 	}

// 	return categories, nil
// }
