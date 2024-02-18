// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/aerosystems/checkmail-service/internal/config"
	"github.com/aerosystems/checkmail-service/internal/http"
	"github.com/aerosystems/checkmail-service/internal/infrastructure/rest"
	"github.com/aerosystems/checkmail-service/internal/infrastructure/rpc"
	"github.com/aerosystems/checkmail-service/internal/models"
	"github.com/aerosystems/checkmail-service/internal/repository/pg"
	"github.com/aerosystems/checkmail-service/internal/repository/rpc"
	"github.com/aerosystems/checkmail-service/internal/usecases"
	"github.com/aerosystems/checkmail-service/pkg/gorm_postgres"
	"github.com/aerosystems/checkmail-service/pkg/logger"
	"github.com/aerosystems/checkmail-service/pkg/oauth"
	"github.com/aerosystems/checkmail-service/pkg/rpc_client"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Injectors from wire.go:

//go:generate wire
func InitApp() *App {
	logger := ProvideLogger()
	logrusLogger := ProvideLogrusLogger(logger)
	config := ProvideConfig()
	baseHandler := ProvideBaseHandler(logrusLogger, config)
	entry := ProvideLogrusEntry(logger)
	db := ProvideGormPostgres(entry, config)
	domainRepo := ProvideDomainRepo(db)
	rootDomainRepo := ProvideRootDomainRepo(db)
	domainUsecase := ProvideDomainUsecase(domainRepo, rootDomainRepo)
	domainHandler := ProvideDomainHandler(baseHandler, domainUsecase)
	filterRepo := ProvideFilterRepo(db)
	projectRepo := ProvideProjectRepo(config)
	filterUsecase := ProvideFilterUsecase(rootDomainRepo, filterRepo, projectRepo)
	filterHandler := ProvideFilterHandler(baseHandler, filterUsecase)
	inspectUsecase := ProvideInspectUsecase(logrusLogger, domainRepo, rootDomainRepo, filterRepo)
	inspectHandler := ProvideInspectHandler(baseHandler, inspectUsecase)
	reviewRepo := ProvideReviewRepo(db)
	reviewUsecase := ProvideReviewUsecase(reviewRepo, rootDomainRepo)
	reviewHandler := ProvideReviewHandler(baseHandler, reviewUsecase)
	accessTokenService := ProvideAccessTokenService(config)
	server := ProvideHttpServer(logrusLogger, config, domainHandler, filterHandler, inspectHandler, reviewHandler, accessTokenService)
	rpcServerServer := ProvideRpcServer(logrusLogger, inspectUsecase)
	app := ProvideApp(logrusLogger, config, server, rpcServerServer)
	return app
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server, rpcServer *RpcServer.Server) *App {
	app := NewApp(log, cfg, httpServer, rpcServer)
	return app
}

func ProvideLogger() *logger.Logger {
	loggerLogger := logger.NewLogger()
	return loggerLogger
}

func ProvideConfig() *config.Config {
	configConfig := config.NewConfig()
	return configConfig
}

func ProvideHttpServer(log *logrus.Logger, cfg *config.Config, domainHandler *rest.DomainHandler, filterHandler *rest.FilterHandler, inspectHandler *rest.InspectHandler, reviewHandler *rest.ReviewHandler, tokenService HttpServer.TokenService) *HttpServer.Server {
	server := HttpServer.NewServer(log, domainHandler, filterHandler, inspectHandler, reviewHandler, tokenService)
	return server
}

func ProvideRpcServer(log *logrus.Logger, inspectUsecase RpcServer.InspectUsecase) *RpcServer.Server {
	server := RpcServer.NewServer(log, inspectUsecase)
	return server
}

func ProvideDomainHandler(baseHandler *rest.BaseHandler, domainUsecase rest.DomainUsecase) *rest.DomainHandler {
	domainHandler := rest.NewDomainHandler(baseHandler, domainUsecase)
	return domainHandler
}

func ProvideFilterHandler(baseHandler *rest.BaseHandler, filterUsecase rest.FilterUsecase) *rest.FilterHandler {
	filterHandler := rest.NewFilterHandler(baseHandler, filterUsecase)
	return filterHandler
}

func ProvideInspectHandler(baseHandler *rest.BaseHandler, inspectUsecase rest.InspectUsecase) *rest.InspectHandler {
	inspectHandler := rest.NewInspectHandler(baseHandler, inspectUsecase)
	return inspectHandler
}

func ProvideReviewHandler(baseHandler *rest.BaseHandler, reviewUsecase rest.ReviewUsecase) *rest.ReviewHandler {
	reviewHandler := rest.NewReviewHandler(baseHandler, reviewUsecase)
	return reviewHandler
}

func ProvideDomainUsecase(domainRepo usecases.DomainRepository, rootDomainRepo usecases.RootDomainRepository) *usecases.DomainUsecase {
	domainUsecase := usecases.NewDomainUsecase(domainRepo, rootDomainRepo)
	return domainUsecase
}

func ProvideFilterUsecase(rootDomainRepo usecases.RootDomainRepository, filterRepo usecases.FilterRepository, projectRepo usecases.ProjectRepository) *usecases.FilterUsecase {
	filterUsecase := usecases.NewFilterUsecase(rootDomainRepo, filterRepo, projectRepo)
	return filterUsecase
}

func ProvideInspectUsecase(log *logrus.Logger, domainRepo usecases.DomainRepository, rootDomainRepo usecases.RootDomainRepository, filterRepo usecases.FilterRepository) *usecases.InspectUsecase {
	inspectUsecase := usecases.NewInspectUsecase(log, domainRepo, rootDomainRepo, filterRepo)
	return inspectUsecase
}

func ProvideReviewUsecase(domainReviewRepo usecases.ReviewRepository, rootDomainRepo usecases.RootDomainRepository) *usecases.ReviewUsecase {
	reviewUsecase := usecases.NewReviewUsecase(domainReviewRepo, rootDomainRepo)
	return reviewUsecase
}

func ProvideDomainRepo(db *gorm.DB) *pg.DomainRepo {
	domainRepo := pg.NewDomainRepo(db)
	return domainRepo
}

func ProvideRootDomainRepo(db *gorm.DB) *pg.RootDomainRepo {
	rootDomainRepo := pg.NewRootDomainRepo(db)
	return rootDomainRepo
}

func ProvideFilterRepo(db *gorm.DB) *pg.FilterRepo {
	filterRepo := pg.NewFilterRepo(db)
	return filterRepo
}

func ProvideReviewRepo(db *gorm.DB) *pg.ReviewRepo {
	reviewRepo := pg.NewReviewRepo(db)
	return reviewRepo
}

// wire.go:

func ProvideLogrusEntry(log *logger.Logger) *logrus.Entry {
	return logrus.NewEntry(log.Logger)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideGormPostgres(e *logrus.Entry, cfg *config.Config) *gorm.DB {
	db := GormPostgres.NewClient(e, cfg.PostgresDSN)
	if err := db.AutoMigrate(&models.Domain{}, &models.RootDomain{}, &models.Filter{}, &models.Review{}); err != nil {
		panic(err)
	}
	return db
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *rest.BaseHandler {
	return rest.NewBaseHandler(log, cfg.Mode)
}

func ProvideProjectRepo(cfg *config.Config) *RpcRepo.ProjectRepo {
	rpcClient := RPCClient.NewClient("tcp", cfg.ProjectServiceRPCAddress)
	return RpcRepo.NewProjectRepo(rpcClient)
}

func ProvideAccessTokenService(cfg *config.Config) *OAuthService.AccessTokenService {
	return OAuthService.NewAccessTokenService(cfg.AccessSecret)
}
