package main

import (
	"fmt"
	"github.com/aerosystems/checkmail-service/internal/handlers"
	"github.com/aerosystems/checkmail-service/internal/models"
	"github.com/aerosystems/checkmail-service/internal/repository"
	GormPostgres "github.com/aerosystems/checkmail-service/pkg/gorm_postgres"
	"github.com/sirupsen/logrus"
	"net/rpc"
	"os"

	"github.com/aerosystems/checkmail-service/pkg/logger"
	"net/http"
)

const webPort = 80

// @title Checkmail Service
// @version 1.0.6
// @description A part of microservice infrastructure, who responsible for store and check email domains in black/whitelists

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey X-API-KEY
// @in header
// @name X-API-KEY
// @description Should contain Token, digits and letters, 64 symbols length

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Should contain Access JWT Token, with the Bearer started

// @host gw.verifire.com/checkmail
// @schemes https
// @BasePath /
func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))

	clientGORM := GormPostgres.NewClient(logrus.NewEntry(log.Logger))
	if err := clientGORM.AutoMigrate(&models.Domain{}, &models.RootDomain{}); err != nil {
		log.Panic(err)
	}
	domainRepo := repository.NewDomainRepo(clientGORM)
	rootDomainRepo := repository.NewRootDomainRepo(clientGORM)

	lookupClientRPC, err := rpc.Dial("tcp", "lookup-service:5001")
	if err != nil {
		log.Fatal(err)
	}

	app := Config{
		BaseHandler: handlers.NewBaseHandler(domainRepo, rootDomainRepo, lookupClientRPC),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(log.Logger),
	}

	log.Info("starting checkmail-service WEB server on port %d\n", webPort)
	err = srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
