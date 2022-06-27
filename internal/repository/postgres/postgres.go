package postgres

import (
	"fmt"
	"github.com/Thunderbirrd/ShortLinksService/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const urlTable = "urls"

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if err = MigrateUp("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)); err != nil {
		return nil, err
	}

	return db, nil
}
