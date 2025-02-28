package services

import (
	"BACKEND/Data"
	"BACKEND/models"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ResetPasswordService(c context.Context, resetData *models.UserReset) (string, error) {
	// Access MongoDB collection
	userDetailsCollection := Data.GetCollection("SkillArcade", "UserDetails")

	// Update the password in DB
	updatePassword := bson.M{"$set": bson.M{"password": resetData.Password}}
	_, err := userDetailsCollection.UpdateOne(c, bson.M{"email": resetData.Email}, updatePassword)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password"})
		return resetData.Email, errors.New("error updating password")
	}
	return resetData.Email, nil
}
func GenerateToken() (string, error) {

	bufferValue := make([]byte, 32) // 256-bit token
	_, err := rand.Read(bufferValue)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bufferValue), nil

}
func PasswordResetToken(resetEmail *models.UserForgot) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, err := GenerateToken()

	if err != nil {
		return "", err
	}

	resetToken := models.PasswordResetToken{
		Email:      resetEmail.Email,
		ResetToken: token,
		ExpiresAt:  time.Now().Add(15 * time.Minute), // Token expires in 15 minutes
	}

	passwordResetTokenCollection := Data.GetCollection("SkillArcade", "PasswordResetToken")
	_, err = passwordResetTokenCollection.InsertOne(ctx, resetToken)
	if err != nil {
		return "", errors.New("error inserting reset token")
	}

	return token, nil

}
