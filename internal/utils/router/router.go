package router

import (
	"fmt"
	"net/http"
)

func NewRouter(prefix string) Router {
	r :=
		Router{
			prefix: prefix,
			Mux:    http.NewServeMux(),
		}
	return r
}

type Router struct {
	prefix string
	Mux    *http.ServeMux
}

// prefixes resource with provided PREFIX method
func methodRoute(method string, resource string) string {
	return fmt.Sprintf("%s %s", method, resource)
}

func (r *Router) GET(url string, handler http.HandlerFunc) {

	route := methodRoute(http.MethodGet, url)

	r.Mux.HandleFunc(route, handler)
}

func (r *Router) POST(url string, handler http.HandlerFunc) {
	route := methodRoute(http.MethodPost, url)

	r.Mux.HandleFunc(route, handler)
}

func (r *Router) DELETE(url string, handler http.HandlerFunc) {
	route := methodRoute(http.MethodDelete, url)

	r.Mux.HandleFunc(route, handler)
}

func (r *Router) PUT(url string, handler http.HandlerFunc) {
	route := methodRoute(http.MethodPut, url)

	r.Mux.HandleFunc(route, handler)
}

// allows the Mux to serve for an entire "route"
func (r *Router) Serve() http.Handler {
	return http.StripPrefix(r.prefix, r.Mux)
}
