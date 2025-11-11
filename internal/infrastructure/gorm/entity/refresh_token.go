package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	REVOKED_FALSE = "false"
	REVOKED_TRUE  = "true"
)

type RefreshToken struct {
	gorm.Model
	UserId    int64
	Jti       string
	TokenHash string
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}

//func NewRefreshTokenFromJwtToken(token *jwt.Token) *RefreshToken {
//	rt := RefreshToken{}
//	rt.Jti = token.Claims.(jwt.MapClaims)["jti"].(string)
//
//}
