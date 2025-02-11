package services

import (
	"BACKEND/Data"
	"BACKEND/models"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/bson"
)

func ForgotPasswordService(c context.Context, requestData *models.UserForgot) (string, error) {

	//Access MongoDB collection
	userDetailsCollection := Data.GetCollection("SkillArcade", "UserDetails")

	// check if user exists in DB
	var userExists bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"email": requestData.Email}).Decode(&userExists)
	if err != nil {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return "", errors.New("email not found")
	}

	token, err := PasswordResetToken(requestData)
	if err != nil {
		return "", err
	}

	err = SendResetEmail(requestData, token)

	if err != nil {
		return "", errors.New("failed to send email")
	} else {
		return "Reset email sent successfully!", err
	}

}

func SendResetEmail(requestData *models.UserForgot, token string) error {

	from := mail.NewEmail("Skillarcade", "spabbathi@ufl.edu") // Sender email
	to := mail.NewEmail("User", requestData.Email)            // Recipient email

	subject := "Password Reset Request"
	resetLink := fmt.Sprintf("https://skillarcade.com/reset-password?token=%s", token)

	// Email content
	plainTextContent := fmt.Sprintf("Click the link below to reset your password:\n\n%s\n\nThis link expires in 15 minutes.", resetLink)
	htmlContent := fmt.Sprintf("<p>Click the link below to reset your password:</p><p><a href='%s'>Reset Password</a></p><p>This link expires in 15 minutes.</p>", resetLink)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// Replace with your SendGrid API Key
	apiKey := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(apiKey)

	// Send the email
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
