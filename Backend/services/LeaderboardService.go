package services

import (
	"BACKEND/Data"
	"BACKEND/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLeaderboardService(ctx context.Context, userID string) ([]models.LeaderboardEntry, *models.LeaderboardEntry, error) {
	userScoreCollection := Data.GetCollection("SkillArcade", "UserScores")

	// Aggregation pipeline to sort users by total_score and join with UserDetails
	pipeline := mongo.Pipeline{
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "UserDetails"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user"},
			},
		}},
		{{Key: "$unwind", Value: "$user"}},
		{{
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: "$user._id"},
				{Key: "username", Value: "$user.username"},
				{Key: "total_score", Value: 1},
				{Key: "quizzes_taken", Value: bson.D{{Key: "$size", Value: "$quizzes"}}},
			},
		}},
		{{Key: "$sort", Value: bson.D{{Key: "total_score", Value: -1}}}}, // Sort by total score descending
	}

	cursor, err := userScoreCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, nil, errors.New("failed to aggregate leaderboard")
	}
	defer cursor.Close(ctx)

	if userID == "" {
		top10, err := getTop10Users(ctx, cursor)
		return top10, nil, err
	} else {
		userRank, err := getUserRank(ctx, cursor, userID)
		return nil, &userRank, err
	}
}

// Helper function to get top 10 users
func getTop10Users(ctx context.Context, cursor *mongo.Cursor) ([]models.LeaderboardEntry, error) {
	var results []models.LeaderboardEntry
	rank := 1

	for cursor.Next(ctx) {
		var entry struct {
			Username     string `bson:"username"`
			TotalScore   int    `bson:"total_score"`
			QuizzesTaken int    `bson:"quizzes_taken"`
		}

		if err := cursor.Decode(&entry); err != nil {
			continue
		}

		results = append(results, models.LeaderboardEntry{
			Username:     entry.Username,
			TotalScore:   entry.TotalScore,
			QuizzesTaken: entry.QuizzesTaken,
			Rank:         rank,
		})

		if rank == 10 {
			break
		}
		rank++
	}

	return results, nil
}

// Helper function  to get a specific user's rank
func getUserRank(ctx context.Context, cursor *mongo.Cursor, userID string) (models.LeaderboardEntry, error) {
	rank := 1

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.LeaderboardEntry{}, errors.New("invalid user ID format")
	}
	// fmt.Println("Received User ID from URL:", userID)
	// fmt.Println("Converted User Object ID:", userObjectID.Hex())

	for cursor.Next(ctx) {
		var entry struct {
			Username     string             `bson:"username"`
			TotalScore   int                `bson:"total_score"`
			QuizzesTaken int                `bson:"quizzes_taken"`
			ID           primitive.ObjectID `bson:"_id"`
		}

		if err := cursor.Decode(&entry); err != nil {
			continue
		}

		// Check if this is the user we are looking for
		if entry.ID == userObjectID {
			// fmt.Println("User found!")
			return models.LeaderboardEntry{
				Username:     entry.Username,
				TotalScore:   entry.TotalScore,
				QuizzesTaken: entry.QuizzesTaken,
				Rank:         rank,
			}, nil
		}
		rank++
	}

	return models.LeaderboardEntry{}, errors.New("user not found")
}
