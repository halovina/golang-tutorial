package router

import (
	"github.com/gorilla/mux"
)

func RegisterRouter() *mux.Router {
	router := mux.NewRouter()
	HealthHandler(router)
	return router
}
