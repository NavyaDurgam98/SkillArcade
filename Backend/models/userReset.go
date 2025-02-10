package models

type UserReset struct {
    Email      string `json:"email"`
    ResetToken string `json:"reset_token"`
    Password   string `json:"password"`
}