package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	healthRes := healthCheckResponse{
		Data: "OK",
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthRes)
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthRes)
	})
}

type healthCheckResponse struct {
	Data string `json:"data"`
}
