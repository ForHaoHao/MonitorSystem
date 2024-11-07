package routers

import (
	"MonitorSystem/controllers"

	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/api/SayHello", controllers.SayHello, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Handler(route.Handler)
		if route.Middleware != nil {
			router.Use(route.Middleware)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
