package rest

import (
	"ProjectRhyno/lib/persistance"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type CameraServiceHandler struct {
	dbhandler persistance.DatabaseHandler
}

//NewCameraServiceHandler initializises a CameraServiceHandler object with dbHandler for which allows to work directly with the database
func NewCameraServiceHandler(databaseHandler persistance.DatabaseHandler) *CameraServiceHandler {
	return &CameraServiceHandler{
		dbhandler: databaseHandler,
	}
}

func getPictures() []byte {
	var photo []byte
	filepath.Walk("../lib/photos", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, err := ioutil.ReadFile(path)
		photo = file
		return err
	})
	return photo
}

func (eh *CameraServiceHandler) addPictureHandler(w http.ResponseWriter, r *http.Request) {
	pic := persistance.Photo{}
	err := json.NewDecoder(r.Body).Decode(&pic)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occurred while persisting picture %s}", err)
		return
	}

	id, err := eh.dbhandler.AddPhoto(pic)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: error occurred while persisting picture %d %s}", id, err)
		return
	}
}
