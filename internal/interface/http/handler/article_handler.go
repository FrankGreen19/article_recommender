package handler

import (
	"article_recommender/internal/domain"
	"article_recommender/internal/infrastructure/render"
	"article_recommender/internal/usecase/service"
	"net/http"
)

type ArticleHandler struct {
	service  *service.ArticleService
	renderer render.Renderer
}

func NewArticleHandler(service *service.ArticleService, renderer render.Renderer) *ArticleHandler {
	return &ArticleHandler{service: service, renderer: renderer}
}

func (h *ArticleHandler) List(w http.ResponseWriter, r *http.Request) {
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

func (h *ArticleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	article, err := h.service.GetArticle()
	if err != nil {
		panic(err)
	}

	h.renderer.Render(w, article)
}
