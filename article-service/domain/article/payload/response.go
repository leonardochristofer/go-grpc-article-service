package payload

import (
	"time"
)

type GetArticleResponse struct {
	Id        string     `json:"id"`
	Author    *string    `json:"author"`
	Title     *string    `json:"title"`
	Body      *string    `json:"body"`
	CreatedAt *time.Time `json:"created_at"`
}

type GetListArticleResponse struct {
	ArticleList []GetArticleResponse
}
