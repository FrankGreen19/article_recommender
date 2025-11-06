package service

import "article_recommender/internal/domain"

type AuthRepository interface {
	GetUserByLoginAndPassword(string, string) (*domain.User, error)
	SaveUser(*domain.User) error
}
type AuthService struct {
	authRepo AuthRepository
}
