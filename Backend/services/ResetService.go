package services

import (
	"BACKEND/Data"
	"BACKEND/models"
	"context"
	"errors"

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
