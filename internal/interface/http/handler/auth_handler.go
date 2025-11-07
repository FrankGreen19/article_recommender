package handler

import (
	"article_recommender/internal/domain"
	"article_recommender/internal/infrastructure/render"
	"article_recommender/internal/interface/http/dto/input"
	"article_recommender/internal/interface/http/dto/output"
	"article_recommender/internal/usecase/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service  *service.AuthService
	renderer render.Renderer
}

func NewAuthHandler(service *service.AuthService, renderer render.Renderer) *AuthHandler {
	return &AuthHandler{service: service, renderer: renderer}
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginDto input.UserLoginDto

	err := json.NewDecoder(request.Body).Decode(&loginDto)
	if err != nil {
		http.Error(writer, "invalid JSON", http.StatusBadRequest)

		return
	}

	err = loginDto.Valid()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	user, err := handler.service.Login(loginDto.Email, loginDto.Password)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)

	handler.renderer.Render(writer, output.NewUserOutputDto(user))
}

func (handler *AuthHandler) Register(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registerDto input.UserRegisterDto

	err := json.NewDecoder(request.Body).Decode(&registerDto)
	if err != nil {
		http.Error(writer, "invalid JSON", http.StatusBadRequest)

		return
	}

	if err := registerDto.Valid(); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	user := domain.NewUser()
	user.Email = registerDto.Email
	user.Password = registerDto.Password
	user.Lastname = registerDto.Lastname
	user.Firstname = registerDto.Firstname

	_, err = handler.service.Register(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(http.StatusCreated)

	handler.renderer.Render(writer, output.NewUserOutputDto(user))
}
