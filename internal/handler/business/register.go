package business

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"nitflex/constant"
	"nitflex/internal/models"
	"nitflex/internal/repository"
	"nitflex/util"
)

func (b *business) Register(ctx context.Context, request *models.RegisterRequest) error {
	// check username existed
	_, err := b.repo.GetUserByUsername(ctx, request.Username)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return util.NewError(constant.ErrorMessage_InternalServerError)
	}
	if err == nil {
		return util.NewError(constant.ErrorMessage_UsernameExisted)
	}

	// check email existed
	_, err = b.repo.GetUserByEmail(ctx, request.Email)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return util.NewError(constant.ErrorMessage_InternalServerError)
	}
	if err == nil {
		return util.NewError(constant.ErrorMessage_EmailExisted)
	}

	// hash password before store
	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		return util.NewError(constant.ErrorMessage_InternalServerError)
	}

	// generate activation token
	activationToken := b.generateActivationToken()

	// store user into db
	err = b.repo.CreateUser(ctx, &repository.User{
		Username:        request.Username,
		Email:           request.Email,
		Password:        hashedPassword,
		ActivationToken: activationToken,
		IsActivated:     false,

		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	// send activation email
	err = b.sendActivationEmail(request.Email, activationToken)
	if err != nil {
		fmt.Print(err)
		return util.NewError(constant.ErrorMessage_InternalServerError)
	}

	return nil
}

func (b *business) generateActivationToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (b *business) sendActivationEmail(toEmail, activationToken string) error {
	var (
		activationLink = fmt.Sprintf("%s/activate?token=%s", os.Getenv("FE_URL"), activationToken)
		body           = fmt.Sprintf("Welcome to Your App!\n\nPlease activate your account by clicking this link:\n%s", activationLink)
		auth           = smtp.PlainAuth("", "giabach9102@gmail.com", "mtongunughwbzjmt", "smtp.gmail.com")
		err            = smtp.SendMail("smtp.gmail.com:587", auth, "giabach9102@gmail.com", []string{toEmail}, []byte(body))
	)
	if err != nil {
		return fmt.Errorf("unable to send email: %w", err)
	}
	return nil
}
