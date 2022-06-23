package grpc_handler

import (
	"context"

	"article-service/domain/article/payload"
	"article-service/lib/helper"
	articlePb "article-service/pb/article"

	"google.golang.org/grpc/codes"

	"github.com/opentracing/opentracing-go"
)

func (h *ArticleHandler) GetArticle(ctx context.Context, msg *articlePb.GetArticleRequest) (*articlePb.GetArticleReturn, error) {
	tracestr := "article.handler.GetArticle"
	select {
	case <-ctx.Done():
		return nil, helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	var result articlePb.GetArticleReturn
	id := msg.GetId()
	ret, err := h.usecase.GetArticle(ctx, &payload.GetArticleRequest{
		Id: id,
	})
	if err != nil {
		// no need to log error, since its calls usecase
		// error comes from usecase should already been handled on usecase
		return &result, helper.Error(codes.Internal, tracestr, err)
	}

	return ret, nil
}
