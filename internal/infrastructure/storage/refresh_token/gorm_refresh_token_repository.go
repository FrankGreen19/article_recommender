package refresh_token

import (
	"article_recommender/internal/infrastructure/gorm/entity"
	"context"

	"gorm.io/gorm"
)

type GormRefreshTokenRepository struct {
	db *gorm.DB
}

func NewGormRefreshTokenRepository(db *gorm.DB) *GormRefreshTokenRepository {
	return &GormRefreshTokenRepository{db: db}
}

func (repo *GormRefreshTokenRepository) GetByJti(ctx context.Context, jti string) (*entity.RefreshToken, error) {
	token, err := gorm.G[entity.RefreshToken](repo.db).Where("id = ?", 1).First(ctx)
}
