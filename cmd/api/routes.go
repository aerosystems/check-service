package main

import (
	_ "github.com/aerosystems/checkmail-service/docs" // docs are generated by Swag CLI, you have to import it.
	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
)

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

	// Private routes OAuth 2.0: check roles [admin, support, business]. Auth implemented on API Gateway
	mux.Group(func(mux chi.Router) {
		mux.Use(func(next http.Handler) http.Handler {
			return app.TokenAuthMiddleware(next, "business", "admin", "support")
		})

		mux.Get("/v1/filters", app.BaseHandler.GetFilterList)
		mux.Post("/v1/filters", app.BaseHandler.CreateFilter)
		mux.Put("/v1/filters/{filterId}", app.BaseHandler.UpdateFilter)
		mux.Delete("/v1/filters/{filterId}", app.BaseHandler.DeleteFilter)
	})

	// Private routes OAuth 2.0: check roles [admin, support]. Auth implemented on API Gateway
	mux.Group(func(mux chi.Router) {
		mux.Use(func(next http.Handler) http.Handler {
			return app.TokenAuthMiddleware(next, "admin", "support")
		})

		mux.Get("/v1/domains/{domainName}", app.BaseHandler.GetDomain)
		mux.Post("/v1/domains", app.BaseHandler.CreateDomain)
		mux.Patch("/v1/domains/{domainName}", app.BaseHandler.UpdateDomain)
		mux.Delete("/v1/domains/{domainName}", app.BaseHandler.DeleteDomain)
	})

	return mux
}
