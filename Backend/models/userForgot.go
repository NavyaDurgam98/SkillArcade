package models

type UserForgot struct {
	Email string `json:"email" binding:"required"`
}
