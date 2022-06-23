package article

import (
	"article-service/domain/article/payload"
	articlePb "article-service/pb/article"
	"context"
)

type ArticleUsecaseInterface interface {
	GetArticle(ctx context.Context, request *payload.GetArticleRequest) (*articlePb.GetArticleReturn, error)
	GetListArticle(ctx context.Context, request *payload.GetListArticleRequest) (*articlePb.GetListArticleReturn, error)

	InsertArticle(ctx context.Context, request *articlePb.InsertArticleRequest) (string, error)
}
