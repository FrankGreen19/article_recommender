package output

import "article_recommender/internal/domain"

type UserOutputDto struct {
	Lastname  string
	Firstname string
	Email     string
}

func NewUserOutputDto(model *domain.User) UserOutputDto {
	return UserOutputDto{
		Lastname:  model.Lastname,
		Firstname: model.Firstname,
		Email:     model.Email,
	}
}
