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

func (repo *GormRefreshTokenRepository) GetByJti(ctx context.Context, jti string) (entity.RefreshToken, error) {
	return gorm.G[entity.RefreshToken](repo.db).Where("jti = ?", jti).First(ctx)
}

func (repo *GormRefreshTokenRepository) Save(ctx context.Context, refreshToken entity.RefreshToken) (entity.RefreshToken, error) {
	var err error

	if refreshToken.ID == 0 {
		err = gorm.G[entity.RefreshToken](repo.db).Create(ctx, &refreshToken)
	} else {
		_, err = gorm.G[entity.RefreshToken](repo.db).Updates(ctx, refreshToken)
	}

	return refreshToken, err
}

func (repo *GormRefreshTokenRepository) Remove(ctx context.Context, token entity.RefreshToken) (entity.RefreshToken, error) {
	affected, err := gorm.G[entity.RefreshToken](repo.db).Where("jti = ?", token.Jti).Delete(ctx)
	if err != nil {
		return entity.RefreshToken{}, err
	}

	if affected == 0 {
		return entity.RefreshToken{}, gorm.ErrRecordNotFound
	}

	return token, nil
}
