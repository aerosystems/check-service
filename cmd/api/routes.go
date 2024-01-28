package main

import (
	_ "github.com/aerosystems/checkmail-service/docs" // docs are generated by Swag CLI, you have to import it.
	"github.com/aerosystems/checkmail-service/internal/models"
	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
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

	// Private routes OAuth 2.0: check roles [staff]. Auth implemented on API Gateway

}

func (app *Config) routes(log *logrus.Logger) http.Handler {
	mux := chi.NewRouter()
	mux.Use(logger.Logger("router", log))

	// Public routes
	mux.Use(middleware.Heartbeat("/ping"))

	// Must be protected with reCAPTCHA on API Gateway
	mux.Post("/v1/domains/count", app.BaseHandler.Count)
	mux.Post("/v1/domains/review", app.BaseHandler.CreateDomainReview)

	// Auth X-Api-Key and reCAPTCHA implemented on API Gateway
	mux.Post("/v1/data/inspect", app.BaseHandler.Inspect)

	// Private routes Basic Auth
	mux.Group(func(mux chi.Router) {
		mux.Use(middleware.BasicAuth("Swagger Docs", map[string]string{
			os.Getenv("BASIC_AUTH_DOCS_USERNAME"): os.Getenv("BASIC_AUTH_DOCS_PASSWORD"),
		}))
		mux.Get("/docs/*", httpSwagger.Handler(
			httpSwagger.URL("doc.json"), // The url pointing to API definition
		))
	})

	// Private routes OAuth 2.0: check roles [staff, business]. Auth implemented on API Gateway
	mux.Group(func(mux chi.Router) {
		mux.Use(func(next http.Handler) http.Handler {
			return app.TokenAuthMiddleware(next, "business", "staff")
		})

		mux.Get("/v1/filters", app.BaseHandler.GetFilterList)
		mux.Post("/v1/filters", app.BaseHandler.CreateFilter)
		mux.Put("/v1/filters/{filterId}", app.BaseHandler.UpdateFilter)
		mux.Delete("/v1/filters/{filterId}", app.BaseHandler.DeleteFilter)
	})

	// Private routes OAuth 2.0: check roles [staff]. Auth implemented on API Gateway
	mux.Group(func(mux chi.Router) {
		mux.Use(func(next http.Handler) http.Handler {
			return app.TokenAuthMiddleware(next, "staff")
		})

		mux.Get("/v1/domains/{domainName}", app.BaseHandler.GetDomain)
		mux.Post("/v1/domains", app.BaseHandler.CreateDomain)
		mux.Patch("/v1/domains/{domainName}", app.BaseHandler.UpdateDomain)
		mux.Delete("/v1/domains/{domainName}", app.BaseHandler.DeleteDomain)
	})

	return mux
}
