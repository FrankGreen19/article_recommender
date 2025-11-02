package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
}

func Load() *Config {
	// Загружаем .env в окружение
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	dbURL := loadDbDsn()

	return &Config{DBURL: dbURL}
}

func loadDbDsn() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		"disable",
	)
}
