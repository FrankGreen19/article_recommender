package article

import (
	"article_recommender/internal/domain"
	"database/sql"
)

type PgArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *PgArticleRepository {
	return &PgArticleRepository{db: db}
}

func (articleRepo *PgArticleRepository) GetById(id int64) (domain.Article, error) {
	row, err := articleRepo.db.Query("SELECT id, title, content FROM articles WHERE id = $1", id)
	if err != nil {
		return domain.Article{}, err
	}
	defer row.Close()

	var article domain.Article

	err = row.Scan(&article.Id, &article.Title, &article.Content)
	if err != nil {
		return domain.Article{}, err
	}

	return article, err
}
