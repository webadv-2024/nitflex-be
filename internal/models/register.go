package models

type RegisterRequest struct {
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}
