package router

import (
	"github.com/gorilla/mux"
	"github.com/pravinba9495/iborg-core/src/controllers"
	"net/http"
)

// RegisterRoutes registers all API routes for the application
func RegisterRoutes() {
	r := mux.NewRouter().StrictSlash(true)

	// Users route
	r.HandleFunc("/users", controllers.UserController)

	http.Handle("/", r)
}
