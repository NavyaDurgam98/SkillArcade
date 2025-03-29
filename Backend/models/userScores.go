package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizEntry struct {
	QuizTopicID   primitive.ObjectID `bson:"quiz_topic_id" json:"quiz_topic_id"`
	QuizTopicName string             `bson:"quiz_topic_name" json:"quiz_topic_name"`
	Score         int                `bson:"score" json:"score"`
	Attempts      int                `bson:"attempts" json:"attempts"`
	SubmittedAt   time.Time          `bson:"submitted_at" json:"submitted_at"`
}

type UserScore struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Quizzes    []QuizEntry        `bson:"quizzes" json:"quizzes"`
	TotalScore int                `bson:"total_score" json:"total_score"`
}
