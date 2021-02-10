package routes

import "net/http"

// Routes types for API
type Routes struct {
	URI             string
	Method          string
	Func            func(http.ResponseWriter, *http.Request)
	RequireUserAuth bool
}
