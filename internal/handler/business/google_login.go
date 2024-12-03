package business

import (
	"context"
	"errors"
	"google.golang.org/api/oauth2/v2"
	"gorm.io/gorm"
	"nitflex/constant"
	"nitflex/internal/handler/models"
	repository2 "nitflex/internal/repository"
	"nitflex/util"
	"time"
)

func (b *business) GoogleLogin(ctx context.Context, userInfo *oauth2.Userinfo) (*models.LoginResponse, error) {
	// check user existed, if not -> create new one
	_, err := b.repo.GetUserByEmail(ctx, userInfo.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = b.repo.CreateUser(ctx, &repository2.User{
			Email: userInfo.Email,
		})
	}

	user, err := b.repo.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
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
	err = b.repo.UpdateRefreshToken(ctx, &repository2.UpdateRefreshTokenParams{
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
