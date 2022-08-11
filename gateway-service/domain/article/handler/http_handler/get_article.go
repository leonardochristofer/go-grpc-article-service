package article_handler

import (
	"encoding/json"
	"fmt"

	article_response "gateway-service/domain/article/response"
	"gateway-service/lib/helper/http_response"
	"gateway-service/lib/helper/timestamp"
	"net/http"

	"github.com/gin-gonic/gin"

	articlePb "gateway-service/pb/article"

	"google.golang.org/protobuf/encoding/protojson"
)

func (h *handler) GetArticle(c *gin.Context) {
	select {
	case <-c.Done():
		h.log.Error(http_response.SendErrorAborted(c))
		return
	default:
	}

	//REQUEST
	in := articlePb.GetArticleRequest{
		Id: c.Query("id"),
	}

	phreturn, err := h.article.GetArticle(h.ctx, &in)
	if err != nil {
		h.log.Error(http_response.SendError(err, c))
		return
	}

	//RETURN
	var out article_response.GetArticleResponse
	// Convert JSON to Proto with CamelCase
	byt, err := protojson.Marshal(phreturn)
	if err != nil {
		fmt.Println("Error marshal Proto to JSON", err)
		h.log.Error(http_response.SendError(err, c))
		return
	}

	err = json.Unmarshal(byt, &out)
	if err != nil {
		fmt.Println("Error unmarshal Proto to JSON", err)
		h.log.Error(http_response.SendError(err, c))
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     http_response.STANDARD_200_STATUS,
		Timestamp:  timestamp.GetNow(),
		Data:       out,
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	c.JSON(http.StatusOK, response)
}
