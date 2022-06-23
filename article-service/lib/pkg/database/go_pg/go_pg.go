package go_pg

import (
	"fmt"

	"article-service/config"

	"github.com/go-pg/pg/v10"
)

func NewPostgresORM(cfg *config.Config) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DBWrite.Host, cfg.DBWrite.Port),
		User:     cfg.DBWrite.User,
		Password: cfg.DBWrite.Password,
		Database: cfg.DBWrite.Database,
	})
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	return db, err
}
