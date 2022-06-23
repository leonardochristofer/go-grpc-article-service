package article_handler

import (
	article_payload "gateway-service/domain/article/payload"
	"gateway-service/lib/helper/http_response"
	"gateway-service/lib/helper/timestamp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) InsertArticle(c *gin.Context) {
	select {
	case <-c.Done():
		h.log.Error(http_response.SendErrorAborted(c))
		return
	default:
	}

	reqBody := new(article_payload.ArticlePayload)
	c.BindJSON(reqBody)

	in := article_payload.ArticlePayload{
		Id:        reqBody.Id,
		Author:    reqBody.Author,
		Title:     reqBody.Title,
		Body:      reqBody.Body,
		CreatedAt: reqBody.CreatedAt,
	}

	result, err := h.article.InsertArticle(h.ctx, in.ToPB())
	if err != nil {
		h.log.Error(http_response.SendError(err, c))
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     http_response.STANDARD_200_STATUS,
		Timestamp:  timestamp.GetNow(),
		Data:       result,
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	c.JSON(http.StatusOK, response)
}
