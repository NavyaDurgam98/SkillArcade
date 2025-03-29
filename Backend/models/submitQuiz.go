package models

type QuizSubmitRequest struct {
	UserID        string `json:"user_id" binding:"required"`
	QuizTopicID   string `json:"quiz_topic_id" binding:"required"`
	QuizTopicName string `json:"quiz_topic_name" binding:"required"`
	Score         int    `json:"score" binding:"required"`
}
