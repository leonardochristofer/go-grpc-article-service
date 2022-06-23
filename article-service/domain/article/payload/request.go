package payload

type GetArticleRequest struct {
	Id string `json:"id"`
}

type GetListArticleRequest struct {
	Sort      []Sort `json:"sort"`
	Offset    int32  `json:"offset"`
	Limit     int32  `json:"limit" validate:"required,limit"`
	Search    string `json:"search"`
	CreatedAt string `json:"created_at"`
	Author    string `json:"author"`
}

type Sort struct {
	Field  string `json:"field"`
	SortBy string `json:"sort_by"`
}
