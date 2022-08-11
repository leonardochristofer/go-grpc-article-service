package usecase

import (
	"article-service/domain/article/payload"
	articlePb "article-service/pb/article"
	"context"
	"errors"
	"time"

	"article-service/lib/helper"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

func (s *service) GetArticle(ctx context.Context, req *payload.GetArticleRequest) (*articlePb.GetArticleReturn, error) {
	var articleReturn articlePb.GetArticleReturn
	tracestr := "article.service.GetArticle"
	select {
	case <-ctx.Done():
		return nil, helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}
	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	s.logger.WithFields(logrus.Fields{
		"ID": req.Id,
	}).Info("incoming request payload")

	if req.Id == "" {
		return &articleReturn, errors.New("Invalid ID")
	}

	result, err := s.articleRepo.GetArticle(ctx, req.Id)
	if err != nil {
		return &articleReturn, helper.Error(codes.Internal, tracestr, err)
	}

	articleReturn.Id = result.Id

	if result.Author != nil {
		articleReturn.Author = *result.Author
	}

	if result.Title != nil {
		articleReturn.Title = *result.Title
	}

	if result.Body != nil {
		articleReturn.Body = *result.Body
	}

	if result.CreatedAt != nil {
		t := *result.CreatedAt
		articleReturn.CreatedAt = t.Format(time.RFC3339)
	}

	return &articleReturn, nil
}
