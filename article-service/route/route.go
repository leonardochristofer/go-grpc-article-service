package route

import (
	"article-service/application"
	"article-service/config"
	"article-service/domain/article/handler/grpc_handler"
	"article-service/domain/article/repo"
	"article-service/domain/article/usecase"
	article "article-service/pb/article"
)

func SetupRouter(cfg *config.Config, app *application.Application) {
	switch app.ServiceMode {
	case "grpc":
		SetupGrpcRouter(cfg, app)
	}
}

func SetupGrpcRouter(cfg *config.Config, app *application.Application) {
	articleRepo := repo.NewPostgresRepo(
		app.DbClients["read"].SqlAdapter,
		app.DbClients["write"].PgAdapter,
	)

	articleUsecase := usecase.NewService(usecase.Dependencies{
		Env:         cfg.Application.Env,
		Gopg:        app.DbClients["write"].PgAdapter,
		Config:      cfg,
		ArticleRepo: articleRepo,
	})

	articleHandler := grpc_handler.NewHandler(
		articleUsecase,
	)

	article.RegisterArticleServiceServer(app.GrpcServer, articleHandler)
}
