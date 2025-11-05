package article

import (
	"article_recommender/internal/domain"
)

type TestArticleRepository struct {
}

func NewTestArticleRepository() *TestArticleRepository {
	return &TestArticleRepository{}
}

func (articleRepo *TestArticleRepository) GetByID(id int64) (*domain.Article, error) {
	article := domain.Article{Id: 1, Title: "Test 1", Content: "Content 1"}

	return &article, nil
}
