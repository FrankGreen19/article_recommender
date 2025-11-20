package handler

import (
	"article_recommender/internal/domain"
	"article_recommender/internal/infrastructure/render"
	"article_recommender/internal/infrastructure/security"
	"article_recommender/internal/interface/http/dto/input"
	"article_recommender/internal/interface/http/dto/output"
	"article_recommender/internal/usecase/service"
	"encoding/json"
	"io"
	"net/http"
)

type AuthHandler struct {
	service    *service.UserService
	renderer   render.Renderer
	jwtManager *security.JwtManager
}

func NewAuthHandler(service *service.UserService, renderer render.Renderer, jwtManager *security.JwtManager) *AuthHandler {
	return &AuthHandler{service: service, renderer: renderer, jwtManager: jwtManager}
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

	user, err := handler.service.GetByEmailAndPassword(loginDto.Email, loginDto.Password)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	ctx := request.Context()
	tokens, err := handler.jwtManager.Generate(ctx, user.Id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(http.StatusOK)

	handler.renderer.Render(writer, output.NewLoginOutputDto(tokens.AccessToken, tokens.RefreshToken))
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

	_, err = handler.service.Create(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(http.StatusCreated)

	handler.renderer.Render(writer, output.NewUserOutputDto(user))
}

func (handler *AuthHandler) Refresh(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
	}

	body, _ := io.ReadAll(request.Body)
	m := make(map[string]string)
	err := json.Unmarshal(body, &m)
	if err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)

		return
	}

	tokens, err := handler.jwtManager.Refresh(request.Context(), m["refresh_token"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	handler.renderer.Render(writer, output.NewLoginOutputDto(tokens.AccessToken, tokens.RefreshToken))
}

func (handler *AuthHandler) Logout(writer http.ResponseWriter, request *http.Request) {}
