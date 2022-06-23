package usecase

import (
	"article-service/domain/article/payload"
	articlePb "article-service/pb/article"
	"context"
	"time"

	"article-service/lib/helper"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *service) GetListArticle(ctx context.Context, req *payload.GetListArticleRequest) (*articlePb.GetListArticleReturn, error) {
	var articleListReturn articlePb.GetListArticleReturn

	tracestr := "article.service.GetListArticle"
	select {
	case <-ctx.Done():
		return nil, helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}
	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	s.logger.WithFields(logrus.Fields{
		"CreatedAt": req.CreatedAt,
		"Sort":      req.Sort,
		"Offset":    req.Offset,
		"Limit":     req.Limit,
		"Search":    req.Search,
		"Author":    req.Author,
	}).Info("incoming request payload")

	result, err := s.articleRepo.GetListArticle(ctx, req)

	if err != nil {
		return &articleListReturn, helper.Error(codes.Internal, tracestr, err)
	}

	for _, article := range result.ArticleList {
		var ArticleReturn articlePb.GetArticleReturn

		ArticleReturn.Id = article.Id

		if article.Author != nil {
			ArticleReturn.Author = *article.Author
		}

		if article.Title != nil {
			ArticleReturn.Title = *article.Title
		}

		if article.Body != nil {
			ArticleReturn.Body = *article.Body
		}

		if article.CreatedAt != nil {
			t := *article.CreatedAt
			ArticleReturn.CreatedAt = t.Format(time.RFC3339)
		}

		articleListReturn.Articles = append(articleListReturn.Articles, &ArticleReturn)
	}

	return &articleListReturn, nil
}
