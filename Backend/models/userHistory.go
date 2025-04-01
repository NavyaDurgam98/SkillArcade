package models

import "time"

type UserHistory struct {
	QuizTopicName string    `bson:"quiz_topic_name" json:"quiz_topic_name"`
	Score         int       `bson:"score" json:"score"`
	Attempts      int       `bson:"attempts" json:"attempts"`
	SubmittedAt   time.Time `bson:"submitted_at" json:"submitted_at"`
}
