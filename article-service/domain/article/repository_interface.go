package article

import (
	"article-service/domain/article/payload"
	articlePb "article-service/pb/article"
	"context"

	"github.com/go-pg/pg/v10"
)

type ArticleRepoInterface interface {
	GetArticle(ctx context.Context, id string) (*payload.GetArticleResponse, error)
	GetListArticle(ctx context.Context, article *payload.GetListArticleRequest) (*payload.GetListArticleResponse, error)

	InsertArticle(ctx context.Context, tx *pg.Tx, in *articlePb.InsertArticleRequest) error
}
