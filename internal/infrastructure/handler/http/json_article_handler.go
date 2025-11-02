package http

import (
	"article_recommender/internal/domain"
	"article_recommender/internal/infrastructure/render"
	"article_recommender/internal/usecase/service"
	"net/http"
)

type JsonArticleHandler struct {
	service  *service.ArticleService
	renderer render.Renderer
}

func NewArticleHandler(service *service.ArticleService, renderer render.Renderer) *JsonArticleHandler {
	return &JsonArticleHandler{service: service, renderer: renderer}
}

func (h *JsonArticleHandler) List(w http.ResponseWriter, r *http.Request) {
	//articles, err := h.service.GetArticles()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	articles := []domain.Article{
		{1, "Test 1", "Content 1"},
		{2, "Test 2", "Content 2"},
	}
	h.renderer.Render(w, articles)
}

func (h *JsonArticleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	article, err := h.service.GetArticle()
	if err != nil {
		panic(err)
	}

	h.renderer.Render(w, article)
}
