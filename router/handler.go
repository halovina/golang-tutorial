package router

import "github.com/gorilla/mux"

func HealthHandler(r *mux.Router) {
	r.Handle("/ping", Ping())
}
