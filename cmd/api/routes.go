package main

import (
	_ "github.com/aerosystems/checkmail-service/docs" // docs are generated by Swag CLI, you have to import it.
	"github.com/aerosystems/checkmail-service/internal/models"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (app *App) NewRouter() *echo.Echo {
	e := echo.New()

	// Private routes Basic Auth
	docsGroup := e.Group("/docs")
	docsGroup.Use(app.basicAuthMiddleware.BasicAuthMiddleware)
	docsGroup.GET("/*", echoSwagger.WrapHandler)

	// Auth X-Api-Key and reCAPTCHA implemented on API Gateway
	e.POST("/v1/data/inspect", app.inspectHandler.Inspect)

	// Protected with reCAPTCHA on API Gateway
	e.POST("/v1/domains/count", app.domainHandler.Count)
	e.POST("/v1/reviews", app.reviewHandler.CreateReview)

	// Private routes OAuth 2.0: check roles [customer, staff]. Auth implemented on API Gateway
	e.GET("/v1/filters", app.filterHandler.GetFilterList, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))
	e.POST("/v1/filters", app.filterHandler.CreateFilter, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))
	e.PUT("/v1/filters/:filterId", app.filterHandler.UpdateFilter, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))
	e.DELETE("/v1/filters/:filterId", app.filterHandler.DeleteFilter, app.oauthMiddleware.AuthTokenMiddleware(models.CustomerRole, models.StaffRole))

	// Private routes OAuth 2.0: check roles [staff]. Auth implemented on API Gateway
	e.GET("/v1/domains/:domainName", app.domainHandler.GetDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))
	e.POST("/v1/domains", app.domainHandler.CreateDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))
	e.PATCH("/v1/domains/:domainName", app.domainHandler.UpdateDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))
	e.DELETE("/v1/domains/:domainName", app.domainHandler.DeleteDomain, app.oauthMiddleware.AuthTokenMiddleware(models.StaffRole))

	return e
}
