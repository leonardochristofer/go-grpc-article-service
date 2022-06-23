package article

type StringMessageResponse struct {
	Id      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type GetArticleResponse struct {
	Id        string `json:"id,omitempty"`
	Author    string `json:"author,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type GetListArticleResponse struct {
	Articles []GetArticleResponse `json:"Articles"`
}

// //---------------------------------FUNCTION---------------------------------

// func NewHealthCheck(in *datapb.StringMessageReturn) *StringMessageResponse {
// 	return &StringMessageResponse{
// 		Id:      in.Id,
// 		Message: in.Message,
// 	}
// }
