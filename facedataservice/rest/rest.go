package rest

import (
	"github.com/gorilla/mux"
)

func ServerAPI() {
	r := mux.NewRouter()
	prefixRouter := r.PathPrefix("/data").Subrouter()
	prefixRouter.Methods("GET").Path("/{dataPhoto}/{photo}")
	prefixRouter.Methods("GET").Path("")
}
