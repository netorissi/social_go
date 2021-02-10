package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Routes{
	{
		URI:             "/users",
		Method:          http.MethodGet,
		Func:            controllers.GetUser,
		RequireUserAuth: false,
	},
	{
		URI:             "/users",
		Method:          http.MethodPost,
		Func:            controllers.CreateUser,
		RequireUserAuth: false,
	},
	{
		URI:             "/users",
		Method:          http.MethodPut,
		Func:            controllers.UpdateUser,
		RequireUserAuth: false,
	},
	{
		URI:             "/users",
		Method:          http.MethodDelete,
		Func:            controllers.DeleteUser,
		RequireUserAuth: false,
	},
}
