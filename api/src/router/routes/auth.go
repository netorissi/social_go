package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoutes = []Routes{
	{
		URI:             "/login",
		Method:          http.MethodPost,
		Func:            controllers.Login,
		RequireUserAuth: false,
	},
}
