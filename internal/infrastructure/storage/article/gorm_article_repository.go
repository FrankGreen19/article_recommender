package article

import (
	"article_recommender/internal/domain"

	"gorm.io/gorm"
)

type GormArticleRepository struct {
	db *gorm.DB
}

func NewGormArticleRepository(db *gorm.DB) *GormArticleRepository {
	return &GormArticleRepository{db: db}
}

func (r *GormArticleRepository) GetByID(id int64) (*domain.Article, error) {
	var article domain.Article
	err := r.db.First(&article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}
