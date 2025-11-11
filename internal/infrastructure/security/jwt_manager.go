package security

import (
	"article_recommender/internal/domain"
	"article_recommender/internal/infrastructure/gorm/entity"
	"article_recommender/internal/infrastructure/security/dto"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenRepository interface {
	GetByUserJti(ctx context.Context, jti string) (*entity.RefreshToken, error)
	Save(ctx context.Context, token entity.RefreshToken) (entity.RefreshToken, error)
}

type JwtManager struct {
	refreshTokenRepo RefreshTokenRepository
}

func NewJwtManager(repo RefreshTokenRepository) *JwtManager {
	return &JwtManager{refreshTokenRepo: repo}
}

func (jwtManager *JwtManager) Generate(user *domain.User) (dto.JwtPair, error) {
	secret := make([]byte, 32)

	atClaims := jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"jti":     generateJTI(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString(secret)
	if err != nil {
		return dto.JwtPair{}, err
	}

	rtJti := generateJTI()
	rtExpiresAt := time.Now().Add(7 * 24 * time.Hour)
	rtClaims := jwt.MapClaims{
		"user_id": user.Id,
		"exp":     rtExpiresAt.Unix(),
		"jti":     rtJti,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString(secret)
	if err != nil {
		return dto.JwtPair{}, err
	}

	refreshTokenEnt := entity.RefreshToken{}
	refreshTokenEnt.Jti = rtJti
	refreshTokenEnt.TokenHash = refreshToken
	refreshTokenEnt.Revoked = false
	refreshTokenEnt.UserId = user.Id
	refreshTokenEnt.ExpiresAt = rtExpiresAt

	if _, err := jwtManager.refreshTokenRepo.Save(refreshTokenEnt); err != nil {
		return dto.JwtPair{}, err
	}

	return dto.JwtPair{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *JwtManager) Refresh(user *domain.User, refreshToken string) (string, error) {
	stored, ok := s.refreshStore[userID]
	if !ok || stored != refreshToken {
		return "", fmt.Errorf("invalid refresh token")
	}

	rt, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil || !rt.Valid {
		return "", fmt.Errorf("refresh token expired or invalid")
	}

	atClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString(s.secret)
}

func generateJTI() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}
