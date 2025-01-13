package models

type LoginRequest struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password"`
}

type LoginResponse struct {
	AccessToken string `bson:"access_token"`
}
