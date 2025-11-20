package entity

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	UserId    int64
	Jti       string
	TokenHash string
	ExpiresAt time.Time
	Revoked   bool
}
