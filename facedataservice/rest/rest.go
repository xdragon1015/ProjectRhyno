package rest

import (
	"ProjectRhyno/lib/persistance"

	"github.com/gorilla/mux"
)

func ServerAPI(dbhandler persistance.DatabaseHandler) {
	handler := NewDataHandler(dbhandler)
	r := mux.NewRouter()
	prefixRouter := r.PathPrefix("/data").Subrouter()
	prefixRouter.Methods("GET").Path("/{photo}").HandlerFunc(handler.findPhotoData)
	prefixRouter.Methods("GET").Path("").HandlerFunc(handler.findAllPhotoData)

}
