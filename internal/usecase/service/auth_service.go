package service

import "article_recommender/internal/domain"

type AuthRepository interface {
	GetUserByLoginAndPassword(string, string) (*domain.User, error)
	SaveUser(*domain.User) error
}
type AuthService struct {
	authRepo AuthRepository
}

func NewAuthService(authRepo AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (service *AuthService) Login(email string, password string) (*domain.User, error) {
	user, err := service.authRepo.GetUserByLoginAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
