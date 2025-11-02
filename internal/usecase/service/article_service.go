package service

import "article_recommender/internal/domain"

type ArticleRepository interface {
	GetByID(id int64) (*domain.Article, error)
}

type ArticleService struct {
	repo ArticleRepository
}

func NewArticleService(repo ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

//func (service *ArticleService) GetArticles() ([]domain.Article, error) {
//	return service.repo.GetAll()
//}

func (service *ArticleService) GetArticle() (*domain.Article, error) {
	return service.repo.GetByID(1)
}
