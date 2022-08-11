package application

import (
	"article-service/logger"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type Application struct {
	Context     context.Context
	Logger      *logger.CustomLogger
	Tracer      opentracing.Tracer
	DbClients   map[string]*DbClient
	GrpcClients map[string]*grpc.ClientConn
	GrpcServer  *grpc.Server
	HttpServer  *gin.Engine
	ServiceMode string
	TraceCloser func()
}

type DbConnectionType uint32

const (
	ReadConnection DbConnectionType = iota
	WriteConnection
)

type DbClient struct {
	Type       DbConnectionType
	SqlAdapter *sql.DB
	PgAdapter  *pg.DB
}
