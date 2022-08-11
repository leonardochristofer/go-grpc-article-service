package article_handler

import (
	"context"
	"fmt"

	articlePb "gateway-service/pb/article"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	ctx     context.Context
	log     *logrus.Logger
	article articlePb.ArticleServiceClient
}

type Handler interface {
	//HealthCheck(*gin.Context)
	GetArticle(*gin.Context)
	GetListArticle(*gin.Context)

	InsertArticle(*gin.Context)
}

func NewArticleHandler(ctx context.Context, log *logrus.Logger, conn articlePb.ArticleServiceClient) Handler {
	return &handler{
		ctx:     ctx,
		log:     log,
		article: conn,
	}
}

func TranslateArrObjSort(c *gin.Context) ([]articlePb.Sort, string) {
	key := "sort"
	queryMap := c.Request.URL.Query()
	isEndIndex := false
	index := 0
	sortList := make([]articlePb.Sort, 0)
	for isEndIndex == false {
		keyQueryField := key + "[" + fmt.Sprint(index) + "][field]"
		keyQuerySortBy := key + "[" + fmt.Sprint(index) + "][sortBy]"
		field, isFieldOk := queryMap[keyQueryField]
		sortBy, isSortByOk := queryMap[keyQuerySortBy]
		if isFieldOk && isSortByOk {
			sortList = append(sortList, articlePb.Sort{
				Field:  field[0],
				SortBy: sortBy[0],
			})
			index += 1
		} else {
			isEndIndex = true
		}
	}

	return sortList, ""
}
