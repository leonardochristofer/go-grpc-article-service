package application

import (
	"article-service/config"
	"article-service/lib/pkg/database/db"
	"article-service/lib/pkg/database/go_pg"
	"article-service/logger"
	"context"
	"net"
	"os"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func Setup(cfg *config.Config, c *cli.Context) (*Application, error) {
	app := new(Application)

	mode := c.String("mode")
	baseInit := []func(*Application) error{
		initLogger(cfg.Logger),
		initTracer(cfg, c.String("tracer")),
		initDatabase(cfg),
		initGrpcClient(cfg),
	}

	switch mode {
	case "grpc":
		baseInit = append(baseInit, initGrpcServer(cfg))
	}
	app.ServiceMode = mode

	if err := runInit(baseInit...)(app); err != nil {
		return app, err
	}

	return app, nil
}

func runInit(appFuncs ...func(*Application) error) func(*Application) error {
	return func(app *Application) error {
		app.Context = context.Background()
		for _, fn := range appFuncs {
			if err := fn(app); err != nil {
				return err
			}
		}
		return nil
	}
}

func initLogger(cfg config.LoggerConfig) func(*Application) error {
	return func(app *Application) error {
		config := &logger.LoggerConfig{
			Level:           logger.DebugLevel,
			Fulltimestamp:   cfg.Fulltimestamp,
			TimestampFormat: cfg.TimestampFormat,
		}

		switch strings.ToLower(cfg.Level) {
		case "debug":
			config.Level = logger.DebugLevel
		case "info":
			config.Level = logger.InfoLevel
		case "error":
			config.Level = logger.ErrorLevel
		case "fatal":
			config.Level = logger.FatalLevel
		}
		log := logger.InitLogger(config)

		app.Logger = log
		return nil
	}
}

func initTracer(cfg *config.Config, tracerType string) func(*Application) error {
	return func(app *Application) error {
		var (
			t      opentracing.Tracer
			isMock bool
		)

		env := strings.ToLower(cfg.Application.Env)
		if os.Getenv("ENV") != "" {
			env = os.Getenv("ENV")
		}

		switch strings.ToLower(env) {
		default:
			isMock = true
		}

		if isMock {
			t = mocktracer.New()
			app.Tracer = t
			app.Logger.Info("init tracer done")
			return nil
		}

		app.Tracer = t
		app.Logger.Info("init tracer done")
		return nil
	}
}

func initDatabase(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		readDB, err := db.NewPostgresDB(cfg)
		if err != nil {
			return err
		}

		writeDB, err := go_pg.NewPostgresORM(cfg)
		if err != nil {
			return err
		}

		app.DbClients = map[string]*DbClient{
			"read": {
				Type:       ReadConnection,
				SqlAdapter: readDB,
				PgAdapter:  writeDB,
			},
			"write": {
				Type:       WriteConnection,
				SqlAdapter: readDB,
				PgAdapter:  writeDB,
			},
		}
		app.Logger.Info("init database done")
		return nil
	}
}

func initGrpcClient(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {

		// add services host
		hosts := map[string]string{}

		app.GrpcClients = make(map[string]*grpc.ClientConn, len(hosts))

		// connect and add to app.GrpcClients
		for key, host := range hosts {
			grpcConn, err := grpc.Dial(host, grpc.WithInsecure())
			if err != nil {
				app.Logger.Error(err)
			} else {
				app.Logger.Info(key + " connected on " + host)

				app.GrpcClients[key] = grpcConn
			}
		}

		app.Logger.Info("init GrpcClient done")
		return nil
	}
}

func initGrpcServer(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		g := grpc.NewServer()
		app.GrpcServer = g
		return nil
	}
}

func grpcRun(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		app.Logger.Info("running grpc server done")
		opentracing.SetGlobalTracer(app.Tracer)
		lis, err := net.Listen("tcp", ":"+cfg.Application.ServerPort)
		if err != nil {
			app.Logger.Error(err)
			return err
		}
		if err := app.GrpcServer.Serve(lis); err != nil {
			app.Logger.Error(err)
			return err
		}
		app.GrpcServer.GracefulStop()
		return nil
	}
}
