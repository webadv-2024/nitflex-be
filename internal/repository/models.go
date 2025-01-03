package repository

import (
	"time"
)

type User struct {
	Id                    string     `bson:"_id,omitempty"`
	Username              string     `bson:"username"`
	Email                 string     `bson:"email"`
	Password              string     `bson:"password"`
	RefreshToken          string     `bson:"refresh_token,omitempty"`
	RefreshTokenExpiresAt *time.Time `bson:"refresh_token_expires_at,omitempty"`
	CreatedAt             time.Time  `bson:"created_at"`
	UpdatedAt             time.Time  `bson:"updated_at"`
}

func (User) TableName() string { return "users" }
