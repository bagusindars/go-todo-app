package middleware

import (
	"log"
	"net/http"
)

type middleware func(http.Handler) http.Handler

func HandlerWithMiddleware(mux *http.ServeMux, path string, handler http.HandlerFunc, middlewares ...middleware) {
	final := chain(handler, middlewares...)
	mux.Handle(path, final)
}

func chain(h http.Handler, middlewares ...middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
