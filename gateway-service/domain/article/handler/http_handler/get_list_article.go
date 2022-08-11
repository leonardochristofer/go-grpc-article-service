package article_handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	article_response "gateway-service/domain/article/response"
	"gateway-service/lib/helper/http_response"
	"gateway-service/lib/helper/timestamp"
	"net/http"

	"github.com/gin-gonic/gin"

	articlePb "gateway-service/pb/article"

	"google.golang.org/protobuf/encoding/protojson"
)

func (h *handler) GetListArticle(c *gin.Context) {
	select {
	case <-c.Done():
		h.log.Error(http_response.SendErrorAborted(c))
		return
	default:
	}

	//Parse Value
	limitVal, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
	limit := int32(limitVal)

	offsetVal, _ := strconv.ParseInt(c.Query("offset"), 10, 32)
	offset := int32(offsetVal)

	sort := c.QueryArray("sort")
	pbSort := make([]*articlePb.Sort, 0)
	if len(sort) > 0 {
		for _, v := range sort {
			sortSplit := strings.Split(v, " ")
			pbSort = append(pbSort, &articlePb.Sort{
				Field:  sortSplit[0],
				SortBy: sortSplit[1],
			})
		}
	} else {
		sortList, errSort := TranslateArrObjSort(c)
		if errSort != "" {
			responseError := http_response.Response{
				StatusCode: http.StatusForbidden,
				Message:    "error",
				Status:     http_response.STANDARD_403_STATUS,
				Timestamp:  timestamp.GetNow(),
				Data:       "Please Check Your Data",
			}

			c.JSON(200, responseError)
			return
		}

		for index := range sortList {
			pbSort = append(pbSort, &articlePb.Sort{
				Field:  sortList[index].Field,
				SortBy: sortList[index].SortBy,
			})
		}
	}

	//REQUEST
	in := articlePb.GetListArticleRequest{
		Offset:    offset,
		Sort:      pbSort,
		Limit:     limit,
		Search:    c.Query("search"),
		CreatedAt: c.Query("createdAt"),
		Author:    c.Query("author"),
	}

	phreturn, err := h.article.GetListArticle(h.ctx, &in)
	if err != nil {
		h.log.Error(http_response.SendError(err, c))
		return
	}

	//RETURN
	var out article_response.GetListArticleResponse
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
