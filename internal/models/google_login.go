package models

type GoogleLoginRequest struct {
	AccessToken string `bson:"access_token" binding:"required"`
}

type GoogleUserInfo struct {
	Email    string `bson:"email"`
	Name     string `bson:"name"`
	Picture  string `bson:"picture"`
	Verified bool   `bson:"verified_email"`
}
