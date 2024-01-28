package main

import (
	_ "github.com/aerosystems/checkmail-service/docs" // docs are generated by Swag CLI, you have to import it.
	"github.com/aerosystems/checkmail-service/internal/models"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()

	// Private routes Basic Auth
	docsGroup := e.Group("/docs")
	docsGroup.Use(app.basicAuthMiddleware.BasicAuthMiddleware)
	docsGroup.GET("/*", echoSwagger.WrapHandler)

	// Auth X-Api-Key and reCAPTCHA implemented on API Gateway
	e.POST("/v1/data/inspect", app.baseHandler.Inspect)

	// Protected with reCAPTCHA on API Gateway
	e.POST("/v1/domains/count", app.baseHandler.Count)
	e.POST("/v1/domains/review", app.baseHandler.CreateDomainReview)

	// Private routes OAuth 2.0: check roles [customer, staff]. Auth implemented on API Gateway
	e.GET("/v1/filters", app.baseHandler.GetFilterList, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))
	e.POST("/v1/filters", app.baseHandler.CreateFilter, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))
	e.PUT("/v1/filters/:filterId", app.baseHandler.UpdateFilter, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))
	e.DELETE("/v1/filters/:filterId", app.baseHandler.DeleteFilter, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))

	// Private routes OAuth 2.0: check roles [staff]. Auth implemented on API Gateway
	e.GET("/v1/domains/:domainName", app.baseHandler.GetDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))
	e.POST("/v1/domains", app.baseHandler.CreateDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))
	e.PATCH("/v1/domains/:domainName", app.baseHandler.UpdateDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))
	e.DELETE("/v1/domains/:domainName", app.baseHandler.DeleteDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))

	return e
}
