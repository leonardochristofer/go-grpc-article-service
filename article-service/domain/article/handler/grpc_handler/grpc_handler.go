package grpc_handler

import (
	article "article-service/domain/article"
	"article-service/lib/pkg/logger"

	"github.com/sirupsen/logrus"
)

type ArticleHandler struct {
	log     *logrus.Logger
	usecase article.ArticleUsecaseInterface
}

func NewHandler(usecase article.ArticleUsecaseInterface) *ArticleHandler {
	return &ArticleHandler{
		log:     logger.GetLogger(),
		usecase: usecase,
	}
}
