package router

import (
	"net/http"
	"time"

	"github.com/bryanljx/go-rest-api/internal/config"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	developmentOrigins = []string{"*"}
	productionOrigins  = []string{"*"} // FIXME: Replace with appropriate origins
)

func setUpMiddleware(mux *http.ServeMux, config *config.Config) *http.Handler {
	// var middlewares []func(http.HandlerFunc) http.HandlerFunc
	middlewares := []func(http.Handler) http.Handler{
		// Injects a request ID in the context of each request
		middleware.RequestID,
		// Sets a http.Request's RemoteAddr to that of either the X-Forwarded-For or X-Real-IP header
		middleware.RealIP,
		// Recovers from panics and return a 500 Internal Service Error
		middleware.Recoverer,
		// Returns a 504 Gateway Timeout after 1 min
		middleware.Timeout(time.Minute),
		// CORS
		corsMiddleware(config),
		// Security headers
		securityHeadersMiddleware,
		// Handles preflight requests with 200 OK and specific headers, which allows some browsers to verify
		// whether cross-domain requests, e.g. from the volunteer app to the backend, are allowed.
		preflightMiddleware,
	}

	res := http.Handler(mux)
	for i := len(middlewares) - 1; i > 0; i-- {
		m := middlewares[i]
		res = m(res)
	}

	return &res
}

func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Disable caching of responses
		w.Header().Set("Cache-Control", "no-store")
		// Protect against drag-and-drop style clickjacking attacks
		w.Header().Set("Content-Security-Policy", "frame-ancestors 'none'")
		// Protect against drag-and-drop style clickjacking attacks in older browsers
		w.Header().Set("X-Frame-Options", "DENY")
		// Inform browsers to connect directly over HTTPS. Browsers should remember this for one year.
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// Prevent browsers from performing MIME sniffing and inappropriately interpreting responses as HTML
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(config *config.Config) func(http.Handler) http.Handler {
	options := cors.Options{
		AllowedOrigins:   productionOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}

	if config.Env == "dev" {
		options.AllowedOrigins = developmentOrigins
	} else if config.Env == "production" {
		options.AllowedOrigins = productionOrigins
	}

	return cors.Handler(options)
}

func preflightMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Preflight requests utilise the OPTIONS method
		if r.Method == "OPTIONS" {
			// Allows "application/json" to be provided via the "Content-Type" header for cross-origin requests.
			// See https://developer.mozilla.org/en-US/docs/Glossary/CORS-safelisted_request_header for more info.
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			// By default, only "simple" methods are allowed, which does not include PUT, DELETE nor PATCH. Specifying
			// this allows cross-origin requests to use all the following methods.
			// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods for more info.
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
