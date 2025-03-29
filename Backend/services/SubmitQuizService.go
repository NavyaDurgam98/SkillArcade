package services

import (
	"BACKEND/Data"
	"BACKEND/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubmitQuizService handles logic to insert/update a user's quiz attempt
func SubmitQuizService(ctx context.Context, payload *models.QuizSubmitRequest) error {
	userScoreCollection := Data.GetCollection("SkillArcade", "UserScores")

	// Convert user_id and quiz_topic_id to ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(payload.UserID)
	if err != nil {
		return errors.New("invalid user_id format")
	}

	quizTopicObjectID, err := primitive.ObjectIDFromHex(payload.QuizTopicID)
	if err != nil {
		return errors.New("invalid quiz_topic_id format")
	}

	// Try to find existing user score document
	var userScore models.UserScore
	err = userScoreCollection.FindOne(ctx, bson.M{"user_id": userObjectID}).Decode(&userScore)

	// CASE 1: No document found → Create new user score document
	if err != nil {
		newUserScore := models.UserScore{
			UserID: userObjectID,
			Quizzes: []models.QuizEntry{
				{
					QuizTopicID:   quizTopicObjectID,
					QuizTopicName: payload.QuizTopicName,
					Score:         payload.Score,
					Attempts:      1,
					SubmittedAt:   time.Now(),
				},
			},
			TotalScore: payload.Score,
		}

		_, insertErr := userScoreCollection.InsertOne(ctx, newUserScore)
		if insertErr != nil {
			return errors.New("failed to insert new user score document")
		}
		return nil
	}

	// CASE 2: Document exists → Update quiz or append new
	found := false
	totalScore := 0

	for i := range userScore.Quizzes {
		quiz := &userScore.Quizzes[i]
		if quiz.QuizTopicID == quizTopicObjectID {
			quiz.Attempts += 1
			quiz.SubmittedAt = time.Now()

			// Only update score if new score is higher
			if payload.Score > quiz.Score {
				quiz.Score = payload.Score
			}
			found = true
		}
		// Calculate totalScore regardless
		totalScore += quiz.Score
	}

	// Quiz not found → add new quiz entry
	if !found {
		newQuiz := models.QuizEntry{
			QuizTopicID:   quizTopicObjectID,
			QuizTopicName: payload.QuizTopicName,
			Score:         payload.Score,
			Attempts:      1,
			SubmittedAt:   time.Now(),
		}
		userScore.Quizzes = append(userScore.Quizzes, newQuiz)
		totalScore += payload.Score
	}

	userScore.TotalScore = totalScore

	// Update existing user document
	_, updateErr := userScoreCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userObjectID},
		bson.M{
			"$set": bson.M{
				"quizzes":     userScore.Quizzes,
				"total_score": userScore.TotalScore,
			},
		},
	)

	if updateErr != nil {
		return errors.New("failed to update user score document")
	}

	return nil
}
