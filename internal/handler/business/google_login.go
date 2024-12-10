package business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"nitflex/constant"
	"nitflex/internal/handler/models"
	"nitflex/internal/repository"
	"nitflex/util"
	"time"
)

func (b *business) GoogleLogin(ctx context.Context, request *models.GoogleLoginRequest) (*models.LoginResponse, error) {
	// Get user info from Google
	userInfo, err := b.getGoogleUserInfo(request.AccessToken)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InvalidGoogleToken)
	}

	// Check if user exists
	user, err := b.repo.GetUserByEmail(ctx, userInfo.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	// Create new user if not exists
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &repository.User{
			Email:    userInfo.Email,
			Username: userInfo.Name,
		}
		err = b.repo.CreateUser(ctx, user)
		if err != nil {
			return nil, util.NewError(constant.ErrorMessage_InternalServerError)
		}
	}

	// Generate access token
	accessToken, err := util.GenerateToken(
		user.Id,
		user.Username,
		time.Now().Add(constant.AccessTokenExpriesIn),
	)
	if err != nil {
		return nil, util.NewError(constant.ErrorMessage_InternalServerError)
	}

	return &models.LoginResponse{
		AccessToken: accessToken,
	}, nil
}

func (b *business) getGoogleUserInfo(accessToken string) (*models.GoogleUserInfo, error) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info")
	}

	var userInfo models.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
