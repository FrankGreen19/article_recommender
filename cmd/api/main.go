package main

import (
	"article_recommender/internal/app"
)

func main() {
	appHttp := app.NewAppHttp()
	appHttp.Run()
}
