package repository

import (
	"time"
)

type User struct {
	Id                    int64      `gorm:"column:id"`
	Username              string     `gorm:"column:username"`
	Email                 string     `gorm:"column:email"`
	Password              string     `gorm:"column:password"`
	RefreshToken          string     `gorm:"column:refresh_token"`
	RefreshTokenExpiresAt *time.Time `gorm:"column:refresh_token_expires_at"`
	// CreatedAt             *time.Time `gorm:"column:created_at"`
	// UpdatedAt             *time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string { return "users" }
