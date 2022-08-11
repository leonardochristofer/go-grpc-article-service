package server

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewServer(log *logrus.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(gin.ErrorLogger())

	return r
}

func Run(r *gin.Engine, listener net.Listener) error {
	// Start http server
	if err := r.RunListener(listener); err != nil {
		return err
	}

	return nil
}
