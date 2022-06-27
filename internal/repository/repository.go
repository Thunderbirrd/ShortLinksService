package repository

import (
	"github.com/Thunderbirrd/ShortLinksService/internal/repository/postgres"
	"github.com/Thunderbirrd/ShortLinksService/internal/repository/ram-storage"
	"github.com/jmoiron/sqlx"
)

type UrlRepository interface {
	SaveNewUrl(longUrl, shortUrl string) error
	CheckLongUrl(longUrl string) (string, error)
	GetLongUrlByShortUrl(shortUrl string) (string, error)
}

type Repository struct {
	UrlRepository
}

func NewRepositoryPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		UrlRepository: postgres.NewRepositoryPostgres(db),
	}
}

func NewRepositoryInternal(storage map[string]string) *Repository {
	return &Repository{UrlRepository: ramstorage.NewRepositoryInternal(storage)}
}
