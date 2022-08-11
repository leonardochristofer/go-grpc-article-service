package usecase

import (
	"article-service/lib/helper"
	articlePb "article-service/pb/article"
	"context"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
)

func (s *service) InsertArticle(ctx context.Context, in *articlePb.InsertArticleRequest) (string, error) {
	tracestr := "article.service.InsertArticle"
	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	tx, err := s.gopg.Begin()
	if err != nil {
		s.logger.Error(tracestr, err)
		return "", helper.Error(codes.Internal, tracestr, err)
	}

	err = s.articleRepo.InsertArticle(ctx, tx, in)
	if err != nil {
		tx.Rollback()
		return "", helper.Error(codes.Internal, tracestr, err)
	}

	err = tx.Commit()
	if err != nil {
		s.logger.Error(tracestr, err)
		return "", helper.Error(codes.Internal, tracestr, err)
	}

	return in.Id, nil
}
