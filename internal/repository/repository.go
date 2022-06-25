package repository

import (
	"github.com/Thunderbirrd/ShortLinksService/internal/repository/postgres"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/jmoiron/sqlx"
)

type UrlRepository interface {
	SaveNewUrl(urlObject models.UrlObject) error
	CheckLongUrl(longUrl string) (string, error)
	GetLongUrlByShortUrl(shortUrl string) (string, error)
}

type Repository struct {
	UrlRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UrlRepository: postgres.NewRepositoryPostgres(db),
	}
}
