package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepositoryPostgres struct {
	db *sqlx.DB
}

func NewRepositoryPostgres(db *sqlx.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (r *RepositoryPostgres) SaveNewUrl(longUrl, shortUrl string) error {
	query := fmt.Sprintf("INSERT INTO %s (long_url, short_url) VALUES ($1, $2) RETURNING id", urlTable)
	_, err := r.db.Exec(query, longUrl, shortUrl)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryPostgres) CheckLongUrl(longUrl string) (string, error) {
	var shortUrl []string
	query := fmt.Sprintf("SELECT short_url FROM %s WHERE long_url = $1", urlTable)

	if err := r.db.Select(&shortUrl, query, longUrl); err != nil {
		return "", err
	}

	if len(shortUrl) != 0 {
		return shortUrl[0], nil
	}
	return "", nil
}

func (r *RepositoryPostgres) GetLongUrlByShortUrl(shortUrl string) (string, error) {
	var longUrl []string
	query := fmt.Sprintf("SELECT long_url FROM %s WHERE short_url = $1", urlTable)

	if err := r.db.Select(&longUrl, query, shortUrl); err != nil {
		return "", err
	}

	if len(longUrl) != 0 {
		return longUrl[0], nil
	}

	return "", nil
}
