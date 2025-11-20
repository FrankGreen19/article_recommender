package security

import (
	"article_recommender/internal/infrastructure/gorm/entity"
	"article_recommender/internal/infrastructure/security/dto"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenRepository interface {
	GetByJti(ctx context.Context, jti string) (entity.RefreshToken, error)
	Save(ctx context.Context, token entity.RefreshToken) (entity.RefreshToken, error)
	Remove(ctx context.Context, token entity.RefreshToken) (entity.RefreshToken, error)
}

type JwtManager struct {
	secret           []byte
	refreshTokenRepo RefreshTokenRepository
}

func NewJwtManager(repo RefreshTokenRepository) *JwtManager {
	secret := []byte(os.Getenv("JWT_SECRET"))

	return &JwtManager{secret: secret, refreshTokenRepo: repo}
}

func (jwtManager *JwtManager) Generate(ctx context.Context, userId int64) (dto.JwtPair, error) {
	atClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"jti":     generateJTI(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString(jwtManager.secret)
	if err != nil {
		return dto.JwtPair{}, err
	}

	rtJti := generateJTI()
	rtExpiresAt := time.Now().Add(7 * 24 * time.Hour)
	rtClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     rtExpiresAt.Unix(),
		"jti":     rtJti,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString(jwtManager.secret)
	if err != nil {
		return dto.JwtPair{}, err
	}

	refreshTokenEnt := entity.RefreshToken{}
	refreshTokenEnt.Jti = rtJti
	refreshTokenEnt.TokenHash = refreshToken
	refreshTokenEnt.Revoked = false
	refreshTokenEnt.UserId = userId
	refreshTokenEnt.ExpiresAt = rtExpiresAt

	if _, err := jwtManager.refreshTokenRepo.Save(ctx, refreshTokenEnt); err != nil {
		return dto.JwtPair{}, err
	}

	return dto.JwtPair{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (jwtManager *JwtManager) Refresh(ctx context.Context, refreshToken string) (dto.JwtPair, error) {
	rt, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return jwtManager.secret, nil
	})
	if err != nil || !rt.Valid {
		return dto.JwtPair{}, errors.New("refresh token expired or invalid")
	}

	claims, ok := rt.Claims.(jwt.MapClaims)
	if !ok {
		return dto.JwtPair{}, errors.New("invalid claims format")
	}

	jti := claims["jti"].(string)

	refreshTokenEnt, err := jwtManager.refreshTokenRepo.GetByJti(ctx, jti)
	if err != nil {
		return dto.JwtPair{}, err
	}

	if _, err := jwtManager.refreshTokenRepo.Remove(ctx, refreshTokenEnt); err != nil {
		return dto.JwtPair{}, err
	}

	tokens, err := jwtManager.Generate(ctx, int64(claims["user_id"].(float64)))
	if err != nil {
		return dto.JwtPair{}, err
	}

	return tokens, nil
}

func generateJTI() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}
