package repo

import (
	article "article-service/domain/article"
	"article-service/lib/pkg/logger"
	"database/sql"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type articleRepo struct {
	log  *logrus.Logger
	db   *sql.DB
	gopg *pg.DB
}

func NewPostgresRepo(db *sql.DB, gopg *pg.DB) article.ArticleRepoInterface {
	return &articleRepo{
		log:  logger.GetLogger(),
		db:   db,
		gopg: gopg,
	}
}
