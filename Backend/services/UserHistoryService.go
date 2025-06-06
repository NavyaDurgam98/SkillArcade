package services

import (
	"BACKEND/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserHistoryService(ctx context.Context, collection *mongo.Collection, userID string) ([]models.UserHistory, error) {
	userScoreCollection := collection

	// Convert userID to ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Find the user document by user_id
	var result struct {
		Quizzes []models.UserHistory `bson:"quizzes"`
	}
	filter := bson.M{"user_id": userObjectID}
	err = userScoreCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, errors.New("user not found or no quiz history available")
	}

	return result.Quizzes, nil
}
