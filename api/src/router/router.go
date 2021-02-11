package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Create - return new router with configs
func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
