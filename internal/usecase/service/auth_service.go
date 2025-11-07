package service

import (
	"article_recommender/internal/domain"
	"errors"
	"fmt"
)

type AuthRepository interface {
	GetUserByLoginAndPassword(string, string) (*domain.User, error)
	SaveUser(*domain.User) error
	GetUserByEmail(string) (*domain.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}

type AuthService struct {
	authRepo AuthRepository
	hasher   PasswordHasher
}

func NewAuthService(authRepo AuthRepository, hasher PasswordHasher) *AuthService {
	return &AuthService{authRepo: authRepo, hasher: hasher}
}

func (service *AuthService) Login(email string, password string) (*domain.User, error) {
	user, err := service.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !service.hasher.Compare(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (service *AuthService) Register(user *domain.User) (*domain.User, error) {
	if _, err := service.authRepo.GetUserByEmail(user.Email); err == nil {
		return nil, errors.New("user with this email already exists")
	}

	fmt.Println(user.Password)
	hashedPassword, err := service.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	if err := service.authRepo.SaveUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
