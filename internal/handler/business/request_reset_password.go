package business

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"nitflex/internal/models"
)

func (b *business) RequestResetPassword(ctx context.Context, request *models.RequestResetPasswordRequest) error {
	// check email existed
	user, err := b.repo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	// generate reset password token
	resetPasswordToken := b.generateResetPasswordToken()

	// store reset password token into db
	err = b.repo.UpdateResetPasswordToken(ctx, user.Username, resetPasswordToken, time.Now().UTC().Add(15*time.Minute))
	if err != nil {
		return err
	}

	// check email activated
	if !user.IsActivated {
		return fmt.Errorf("email is invalid or not activated")
	}

	// send email
	err = b.sendResetPasswordEmail(user.Email, resetPasswordToken)
	if err != nil {
		return err
	}

	return nil
}

func (b *business) generateResetPasswordToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (b *business) sendResetPasswordEmail(toEmail, resetPasswordToken string) error {
	resetPasswordLink := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("FE_URL"), resetPasswordToken)

	// Email headers
	subject := "Subject: Reset Your Password\n"
	to := fmt.Sprintf("To: %s\n", toEmail)
	from := "From: giabach9102@gmail.com\n"
	contentType := "Content-Type: text/plain; charset=\"utf-8\"\n\n"

	// Email body
	body := fmt.Sprintf("Reset your password by clicking this link:\n%s", resetPasswordLink)

	// Full email message
	message := subject + to + from + contentType + body

	// Send email
	auth := smtp.PlainAuth("", "giabach9102@gmail.com", "mtongunughwbzjmt", "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "giabach9102@gmail.com", []string{toEmail}, []byte(message))
	if err != nil {
		return fmt.Errorf("unable to send email: %w", err)
	}
	return nil
}
