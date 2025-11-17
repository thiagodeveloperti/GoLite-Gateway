package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func CreateReverseProxy(target string) http.Handler {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if target == "self" {
			w.Write([]byte("OK"))
		} else {
			proxy.ServeHTTP(w, r)
		}
	})
}
