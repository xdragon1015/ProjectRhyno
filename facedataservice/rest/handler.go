package rest

import (
	"ProjectRhyno/lib/persistance"
	"encoding/json"
	"fmt"
	"net/http"
)

type dataFaceHandler struct {
	dbHandler persistance.DatabaseHandler
}

func NewDataHandler(databaseHandler persistance.DatabaseHandler) *dataFaceHandler {
	return &dataFaceHandler{
		dbHandler: databaseHandler,
	}
}

func (eh *dataFaceHandler) findPhotoData(w http.ResponseWriter, r *http.Request) {

}
func (eh *dataFaceHandler) findAllPhotoData(w http.ResponseWriter, r *http.Request) {

}

func (eh *dataFaceHandler) newFaceHandler(w http.ResponseWriter, r *http.Request) {
	photoFaceImage := persistance.Photo{}
	err := json.NewDecoder(r.Body).Decode(photoFaceImage)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while persisting photo %s}", err)
		return
	}

	id, err := eh.dbHandler.AddPhoto(photoFaceImage)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occured while persisting event %d %s}", id, err)
		return
	}
}
