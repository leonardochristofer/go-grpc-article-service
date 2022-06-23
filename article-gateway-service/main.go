package main

import (
	"context"
	"gateway-service/config"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"

	"gateway-service/lib/pkg/logger"
	"gateway-service/lib/pkg/server"
	"gateway-service/route"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

	// loading config xyz
	_, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	// Logrus instance
	log := logger.GetLogger()

	t := opentracer.New(
		tracer.WithServiceName(os.Getenv("SERVICE_NAME")),
		tracer.WithAnalytics(true),
		tracer.WithAgentAddr(os.Getenv("DD_AGENT_HOST")),
	)
	opentracing.SetGlobalTracer(t)
	defer tracer.Stop()

	ctx := context.Background()
	srv := server.NewServer(log)
	route.SetupRouter(srv)

	// Add your service route here
	// Please sort it alphabetically to reduced conflicts
	route.SetupRouterArticle(ctx, log, srv)

	log.Info("Starting server " + os.Getenv("PORT"))
	srv.Run()
}
