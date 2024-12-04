package business

import (
	"context"
	"errors"
	"net/mail"
	"nitflex/internal/handler/models"
	"nitflex/internal/repository"
	"unicode/utf8"

	"gorm.io/gorm"

	"nitflex/constant"
	"nitflex/util"
)

func (b *business) Register(ctx context.Context, request *models.RegisterRequest) error {
	// validate username length
	usernameLen := utf8.RuneCountInString(request.Username)
	if usernameLen < constant.MinUsernameLength || usernameLen > constant.MaxUsernameLength {
		return util.NewError(constant.ErrorMessage_InvalidUsernameLength)
	}

	// validate password length
	passwordLen := utf8.RuneCountInString(request.Password)
	if passwordLen < constant.MinPasswordLength || passwordLen > constant.MaxPasswordLength {
		return util.NewError(constant.ErrorMessage_InvalidPasswordLength)
	}

	// validate email length and format
	emailLen := utf8.RuneCountInString(request.Email)
	if emailLen == 0 || emailLen > constant.MaxEmailLength {
		return util.NewError(constant.ErrorMessage_InvalidEmailLength)
	}
	if _, err := mail.ParseAddress(request.Email); err != nil {
		return util.NewError(constant.ErrorMessage_InvalidEmailFormat)
	}

	// check username existed
	_, err := b.repo.GetUserByUsername(ctx, request.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return util.NewError(constant.ErrorMessage_InternalServerError)
	}
	if err == nil {
		return util.NewError(constant.ErrorMessage_UsernameExisted)
	}

	// check email existed
	_, err = b.repo.GetUserByEmail(ctx, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

	// store user into db
	err = b.repo.CreateUser(ctx, &repository.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	})

	return err
}
