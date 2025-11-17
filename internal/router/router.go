package router

import (
	"log"
	"net/http"

	"golite-gateway/internal/auth"
	"golite-gateway/internal/config"
	"golite-gateway/internal/store"
)

func New(cfg config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	store.SetRoutes(cfg.Routes)

	for _, r := range cfg.Routes {
		log.Println("Registering route:", r.Path, "->", r.Upstream)

		handler := store.ProxyHandler(r.Upstream)

		// Apply JWT auth if enabled
		if r.Auth {
			handler = auth.JWTMiddleware(handler)
		}

		mux.Handle(r.Path, handler)
	}

	return mux
}
