package router

import "github.com/gorilla/mux"

// Create - return new router with configs
func Create() *mux.Router {
	return mux.NewRouter()
}
