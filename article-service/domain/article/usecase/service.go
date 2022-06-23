package usecase

import (
	"article-service/config"
	"article-service/domain/article"
	"article-service/lib/pkg/logger"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type service struct {
	env         string
	logger      *logrus.Logger
	gopg        *pg.DB
	config      *config.Config
	articleRepo article.ArticleRepoInterface
}

type Dependencies struct {
	Env         string
	Gopg        *pg.DB
	Config      *config.Config
	ArticleRepo article.ArticleRepoInterface
}

func NewService(deps Dependencies) article.ArticleUsecaseInterface {
	return &service{
		env:         deps.Env,
		logger:      logger.GetLogger(),
		gopg:        deps.Gopg,
		config:      deps.Config,
		articleRepo: deps.ArticleRepo,
	}
}
