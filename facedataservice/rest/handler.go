package rest

import (
	"net/http"
)

type dataFaceHandler struct {
}

func NewDataHandler() *dataFaceHandler {
	return &dataFaceHandler{}
}

func (*dataFaceHandler) findPhotoData(w http.ResponseWriter, r *http.Request) {

}

func (*dataFaceHandler) findAllPhotoData(w http.ResponseWriter, r *http.Request) {

}
