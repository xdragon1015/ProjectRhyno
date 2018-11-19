package rest

import (
	"github.com/gorilla/mux"
)

func ServerAPI() {
	handler := NewDataHandler()
	r := mux.NewRouter()
	prefixRouter := r.PathPrefix("/data").Subrouter()
	prefixRouter.Methods("GET").Path("/{dataPhoto}/{photo}").HandlerFunc(handler.findPhotoData)
	prefixRouter.Methods("GET").Path("").HandlerFunc(handler.findAllPhotoData)

}
