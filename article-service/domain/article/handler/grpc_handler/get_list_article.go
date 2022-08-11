package grpc_handler

import (
	"context"

	"article-service/domain/article/payload"
	"article-service/lib/helper"
	articlePb "article-service/pb/article"

	"google.golang.org/grpc/codes"

	"github.com/opentracing/opentracing-go"
)

func (h *ArticleHandler) GetListArticle(ctx context.Context, msg *articlePb.GetListArticleRequest) (*articlePb.GetListArticleReturn, error) {
	tracestr := "article.handler.GetListArticle"
	select {
	case <-ctx.Done():
		return nil, helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	var result articlePb.GetListArticleReturn

	sort := msg.GetSort()
	sortList := make([]payload.Sort, 0)

	for index := range sort {
		sortList = append(sortList, payload.Sort{
			Field:  sort[index].GetField(),
			SortBy: sort[index].GetSortBy(),
		})
	}

	ret, err := h.usecase.GetListArticle(ctx, &payload.GetListArticleRequest{
		Author:    msg.GetAuthor(),
		CreatedAt: msg.GetCreatedAt(),
		Sort:      sortList,
		Offset:    msg.GetOffset(),
		Limit:     msg.GetLimit(),
		Search:    msg.GetSearch(),
	})
	if err != nil {
		// no need to log error, since its calls usecase
		// error comes from usecase should already been handled on usecase
		return &result, helper.Error(codes.Internal, tracestr, err)
	}

	return ret, nil
}
