package rest

import (
	"ProjectRhyno/lib/persistance"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerAPI(endPoint string, dbHandler persistance.DatabaseHandler) {
	handler := NewCameraServiceHandler(dbHandler)
	r := mux.NewRouter()
	prefixRouter := r.PathPrefix("/data").Subrouter()
	prefixRouter.Methods("POST").Path("").HandlerFunc(handler.addPictureHandler)
	http.ListenAndServe(endPoint, r)
}
