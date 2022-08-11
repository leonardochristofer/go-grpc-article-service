package db

import (
	"database/sql"
	"fmt"

	"article-service/config"
	"article-service/lib/pkg/logger"

	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBRead.Host,
		cfg.DBRead.Port,
		cfg.DBRead.User,
		cfg.DBRead.Password,
		cfg.DBRead.Database,
	)

	logger.GetLogger().Info(dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = PingDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PingDB(db *sql.DB) error {
	_, err := db.Exec("SELECT 1")
	return err
}
