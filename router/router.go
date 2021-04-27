package router

import (
	"github.com/alfanherya/jakarsafx/jakarsafx-registrasi/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")

	return router
}
