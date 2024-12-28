package business

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"nitflex/constant"
	"nitflex/internal/models"
	"nitflex/internal/repository"
	"nitflex/util"
)

func (b *business) Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
	var (
		user *repository.User
		err  error
	)

	user, err = b.repo.GetUserByUsername(ctx, request.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, util.NewError(constant.ErrorMessage_InvalidUsernameOrPassword)
	}

	// generate access token
	accessToken, err := util.GenerateToken(
		user.Id,
		user.Username,
		time.Now().Add(constant.AccessTokenExpriesIn))
	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshToken, err := util.GenerateToken(
		user.Id,
		user.Username,
		time.Now().Add(constant.RefreshTokenExpriesIn))

	if err != nil {
		return nil, err
	}

	// store refresh token
	err = b.repo.UpdateRefreshToken(ctx, &repository.UpdateRefreshTokenParams{
		UserId:                user.Id,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: time.Now().Add(constant.RefreshTokenExpriesIn),
	})
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		AccessToken: accessToken,
	}, nil
}
