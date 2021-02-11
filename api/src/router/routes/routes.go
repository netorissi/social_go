package routes

import (
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

	for _, route := range userRoutes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
