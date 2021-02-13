package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes types for API
type Routes struct {
	URI             string
	Method          string
	Func            func(http.ResponseWriter, *http.Request)
	RequireUserAuth bool
}

func Config(r *mux.Router) *mux.Router {
	var routes []Routes

	// users
	routes = append(routes, userRoutes...)
	// auth
	routes = append(routes, authRoutes...)

	for _, route := range routes {
		if route.RequireUserAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.UserAuthentication(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
