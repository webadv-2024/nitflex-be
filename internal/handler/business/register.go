package business

import (
	"context"
	"errors"
	"nitflex/internal/handler/models"
	"nitflex/internal/repository"

	"gorm.io/gorm"

	"nitflex/constant"
	"nitflex/util"
)

func (b *business) Register(ctx context.Context, request *models.RegisterRequest) error {
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
