package app

import (
	http2 "article_recommender/internal/infrastructure/handler/http"
	"article_recommender/internal/infrastructure/render"
	storage "article_recommender/internal/infrastructure/storage/article"
	service2 "article_recommender/internal/usecase/service"
	"log"
	"net/http"
)

type App struct {
	server *http.Server
}

func NewAppHttp() *App {
	//db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/dbname?sslmode=disable")
	//if err != nil {
	//	log.Fatal(err)
	//}

	repo := storage.NewTestArticleRepository()
	service := service2.NewArticleService(repo)
	renderer := render.JSONRenderer{}
	handler := http2.NewArticleHandler(service, renderer)

	mux := http.NewServeMux()
	mux.HandleFunc("/article", handler.GetByID)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &App{server: srv}
}

func (a *App) Run() {
	log.Println("Server running on :8080")
	if err := a.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
