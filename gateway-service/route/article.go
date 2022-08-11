package route

import (
	"context"
	"os"

	article "gateway-service/domain/article/handler/http_handler"
	articlePb "gateway-service/pb/article"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func SetupRouterArticle(ctx context.Context, log *logrus.Logger, r *gin.Engine) {
	serviceName := "MY_SERVICE"
	grpcAddress := os.Getenv(serviceName)
	grpcConn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %s", err)
	} else {
		log.Info(serviceName + " connected on " + grpcAddress)
	}
	grpcService := articlePb.NewArticleServiceClient(grpcConn)

	// HTTP Handler
	httpHandler := article.NewArticleHandler(ctx, log, grpcService)

	r.GET("/article", httpHandler.GetArticle)
	r.GET("/articles", httpHandler.GetListArticle)

	r.POST("/articles", httpHandler.InsertArticle)
}
