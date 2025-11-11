package service

import (
	"article_recommender/internal/domain"
	"errors"
	"fmt"
)

type UserRepository interface {
	GetUserByLoginAndPassword(string, string) (*domain.User, error)
	SaveUser(*domain.User) error
	GetUserByEmail(string) (*domain.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}

type UserService struct {
	authRepo UserRepository
	hasher   PasswordHasher
}

func NewUserService(authRepo UserRepository, hasher PasswordHasher) *UserService {
	return &UserService{authRepo: authRepo, hasher: hasher}
}

func (service *UserService) GetByEmailAndPassword(email string, password string) (*domain.User, error) {
	user, err := service.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !service.hasher.Compare(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (service *UserService) Create(user *domain.User) (*domain.User, error) {
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
