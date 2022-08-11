package article

import (
	articlePb "gateway-service/pb/article"
)

type ArticlePayload struct {
	Id        string `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

func (o *ArticlePayload) ToPB() *articlePb.InsertArticleRequest {
	pbReturn := articlePb.InsertArticleRequest{
		Id:        o.Id,
		Author:    o.Author,
		Title:     o.Title,
		Body:      o.Body,
		CreatedAt: o.CreatedAt,
	}
	return &pbReturn
}

// //---------------------------------FUNCTION---------------------------------

// func (o *HealthCheckTypesPayload) ToPB() *datapb.StringMessageReturn {
// 	return &datapb.StringMessageReturn{
// 		Id:      o.Id,
// 		Message: o.Message,
// 	}
// }
