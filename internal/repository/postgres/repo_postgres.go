package postgres

import (
	"fmt"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/jmoiron/sqlx"
)

type RepositoryPostgres struct {
	db *sqlx.DB
}

func NewRepositoryPostgres(db *sqlx.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (r *RepositoryPostgres) SaveNewUrl(urlObject models.UrlObject) error {
	query := fmt.Sprintf("INSERT INTO %s (long_url, short_url) VALUES $1, $2 RETURNING id", urlTable)
	_, err := r.db.Exec(query, urlObject.LongUrl, urlObject.ShortUrl)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPostgres) CheckLongUrl(longUrl string) (string, error) {
	var shortUrl string
	query := fmt.Sprintf("SELECT short_url FROM %s WHERE long_url = $1", urlTable)

	if err := r.db.Select(&shortUrl, query, longUrl); err != nil {
		return shortUrl, err
	}

	return shortUrl, nil
}
