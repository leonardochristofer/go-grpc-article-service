package repo

import (
	"article-service/domain/article/payload"
	"article-service/lib/helper"
	"context"
	"errors"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
)

func (r *articleRepo) GetListArticle(ctx context.Context, request *payload.GetListArticleRequest) (*payload.GetListArticleResponse, error) {
	datas := make([]payload.GetArticleResponse, 0)
	tracestr := "article.repo.GetListArticle"
	select {
	case <-ctx.Done():
		return nil, helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}
	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	sort := ``
	offset := ``
	limit := ``

	paramQueries := []interface{}{}

	query := `SELECT * FROM article`

	where := ` WHERE created_at IS NOT NULL `

	if request.CreatedAt != "" {
		where += `AND created_at <= TO_TIMESTAMP('` + request.CreatedAt + `', 'YYYY-MM-DD HH24:MI:SS') `
	}

	if request.Search != "" {
		where += `AND (title ILIKE '%` + request.Search + `%' OR body ILIKE '%` + request.Search + `%') `
	}

	if request.Author != "" {
		where += `AND author = '` + request.Author + `' `
	}

	if len(request.Sort) > 0 {
		isIdExist := false
		sortObj := request.Sort
		sort = `ORDER BY `
		for index := range sortObj {
			sort += sortObj[index].Field + ` ` + sortObj[index].SortBy
			if len(sortObj)-1 != index {
				sort += `,`
			}

			if sortObj[index].Field == "id" {
				isIdExist = true
			}
		}

		if !isIdExist {
			sort += ` title asc`
		}

		sort += ` `
	}

	if request.Offset > 0 {
		offset = fmt.Sprintf(`OFFSET %d `, request.Offset) + ` `
	}

	if request.Limit > 0 {
		limit = fmt.Sprintf(`LIMIT %d `, request.Limit)
	}

	query = query + where + sort + offset + limit

	fmt.Println(query)

	rows, err := r.db.QueryContext(ctx, query, paramQueries...)
	if err != nil {
		r.log.Error(tracestr + ":" + err.Error())
		return nil, helper.Error(codes.InvalidArgument, tracestr, err)
	}
	defer rows.Close()

	for rows.Next() {
		var rowResult payload.GetArticleResponse
		err := rows.Scan(
			&rowResult.Id,
			&rowResult.Author,
			&rowResult.Title,
			&rowResult.Body,
			&rowResult.CreatedAt,
		)
		if err != nil {
			r.log.Error(tracestr + ":" + err.Error())
			return nil, err
		}

		datas = append(datas, rowResult)
	}

	result := payload.GetListArticleResponse{}
	result.ArticleList = append(result.ArticleList, datas...)

	return &result, nil
}

func (r *articleRepo) GetArticle(ctx context.Context, id string) (*payload.GetArticleResponse, error) {
	var result payload.GetArticleResponse
	tracestr := "article.repo.GetArticle"
	select {
	case <-ctx.Done():
		return &result, helper.Error(codes.Aborted, tracestr+": Request Aborted", nil)
	default:
	}
	span, ctx := opentracing.StartSpanFromContext(ctx, tracestr)
	defer span.Finish()

	if id == "" {
		return &result, helper.Error(codes.InvalidArgument, tracestr+" : ", errors.New("ID cannot empty"))
	}
	paramQueries := []interface{}{}
	paramQueries = append(paramQueries, id)

	query := `SELECT * FROM article where id = $1`

	err := r.db.QueryRowContext(ctx, query, paramQueries...).Scan(
		&result.Id,
		&result.Author,
		&result.Title,
		&result.Body,
		&result.CreatedAt,
	)
	if err != nil {
		r.log.Error(tracestr + ":" + err.Error())
		return &result, helper.Error(codes.InvalidArgument, tracestr, err)
	}

	return &result, nil
}
