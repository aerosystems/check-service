package main

import (
	"github.com/aerosystems/checkmail-service/internal/common/config"
	GRPCServer "github.com/aerosystems/checkmail-service/internal/presenters/grpc"
	HttpServer "github.com/aerosystems/checkmail-service/internal/presenters/http"
	"github.com/sirupsen/logrus"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	httpServer *HttpServer.Server
	grpcServer *GRPCServer.Server
}

func NewApp(
	log *logrus.Logger,
	cfg *config.Config,
	httpServer *HttpServer.Server,
	grpcServer *GRPCServer.Server,
) *App {
	return &App{
		log:        log,
		cfg:        cfg,
		httpServer: httpServer,
		grpcServer: grpcServer,
	}
}
