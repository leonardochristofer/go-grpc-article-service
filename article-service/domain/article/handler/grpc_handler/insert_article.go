package grpc_handler

import (
	"context"

	"article-service/lib/helper"
	articlePb "article-service/pb/article"

	"google.golang.org/grpc/codes"

	"github.com/opentracing/opentracing-go"
)

func (h *ArticleHandler) InsertArticle(ctx context.Context, req *articlePb.InsertArticleRequest) (*articlePb.StringMessageReturn, error) {
	tracestr := "article.handler.InsertArticle"
	select {
	case <-ctx.Done():
		return nil, helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}
	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()
	id, err := h.usecase.InsertArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return &articlePb.StringMessageReturn{
		Id:      id,
		Message: "Success Insert Article",
	}, nil
}
