package models

import "time"

type UserReset struct {
	Email      string `json:"email"`
	ResetToken string `json:"reset_token"`
	Password   string `json:"password"`
}

type PasswordResetToken struct {
	Email      string    `json:"email"`
	ResetToken string    `json:"reset_token"`
	ExpiresAt  time.Time `bson:"expires_at"`
}
