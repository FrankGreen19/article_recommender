package user

import (
	"article_recommender/internal/domain"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) GetUserByLoginAndPassword(login string, password string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email = ? AND password = ?", login, password).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) SaveUser(user *domain.User) error {
	err := r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}
