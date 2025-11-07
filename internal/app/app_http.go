package app

import (
	"article_recommender/internal/infrastructure/render"
	"article_recommender/internal/infrastructure/security"
	storage "article_recommender/internal/infrastructure/storage/article"
	"article_recommender/internal/infrastructure/storage/user"
	http2 "article_recommender/internal/interface/http/handler"
	service2 "article_recommender/internal/usecase/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	server *http.Server
}

func NewAppHttp() *App {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/test/article/json", getTestArticleHandler().GetByID)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// GORM group
	mux.HandleFunc("/article/json", geHttpArticleHandler(db).GetByID)
	mux.HandleFunc("/user/login", getHttpAuthHandler(db).Login)
	mux.HandleFunc("/user/register", getHttpAuthHandler(db).Register)

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

func getTestArticleHandler() *http2.ArticleHandler {
	repo := storage.NewTestArticleRepository()
	service := service2.NewArticleService(repo)
	renderer := render.JSONRenderer{}

	return http2.NewArticleHandler(service, renderer)
}

func geHttpArticleHandler(db *gorm.DB) *http2.ArticleHandler {
	repo := storage.NewGormArticleRepository(db)
	service := service2.NewArticleService(repo)
	renderer := render.JSONRenderer{}

	return http2.NewArticleHandler(service, renderer)
}

func getHttpAuthHandler(db *gorm.DB) *http2.AuthHandler {
	repo := user.NewGormUserRepository(db)
	hasher := security.NewBcryptPasswordHasher()
	service := service2.NewAuthService(repo, hasher)
	renderer := render.JSONRenderer{}

	return http2.NewAuthHandler(service, renderer)
}
