package store

import (
	"net/http"

	"golite-gateway/internal/config"
	"golite-gateway/internal/proxy"
)

var routes []config.Route

func SetRoutes(r []config.Route) {
	routes = r
}

func ProxyHandler(upstream string) http.Handler {
	return proxy.CreateReverseProxy(upstream)
}
