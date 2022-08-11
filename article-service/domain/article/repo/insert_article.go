package repo

import (
	articlePb "article-service/pb/article"
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"

	"article-service/lib/helper"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
)

func (r *articleRepo) InsertArticle(ctx context.Context, tx *pg.Tx, in *articlePb.InsertArticleRequest) error {
	tracestr := "article.repo.InsertArticle"
	select {
	case <-ctx.Done():
		return helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}
	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	paramQueries := []interface{}{}
	param := ``
	values := ``

	if in.Id == "" {
		in.Id = uuid.New().String()
	}
	if in.CreatedAt == "" {
		in.CreatedAt = time.Now().Format(time.RFC3339)
	}
	if in.Author == "" {
		in.Author = ""
	}
	if in.Id != "" {
		param += `id,`
		values += `?,`
		paramQueries = append(paramQueries, in.Id)
	}
	if in.Author != "" {
		param += `author,`
		values += `?,`
		paramQueries = append(paramQueries, in.Author)
	}
	if in.Title != "" {
		param += `title,`
		values += `?,`
		paramQueries = append(paramQueries, in.Title)
	}
	if in.Body != "" {
		param += `body,`
		values += `?,`
		paramQueries = append(paramQueries, in.Body)
	}
	if in.CreatedAt != "" {
		param += `created_at,`
		values += `?,`
		paramQueries = append(paramQueries, in.CreatedAt)
	}

	cleanParam := param[:len(param)-1]
	cleanValues := values[:len(values)-1]

	query := `INSERT INTO article (` + cleanParam + `) values (` + cleanValues + `)`

	_, err := tx.ExecContext(ctx, query, paramQueries...)
	if err != nil {
		r.log.Error(tracestr + ":" + err.Error())
		return helper.Error(codes.Internal, tracestr, err)
	}

	return nil
}
