package persistance

import (
	"gopkg.in/mgo.v2/bson"
)

//Photo photo structure for the camera microservice
type Photo struct {
	ID    bson.ObjectId `bson:"id"`
	Photo []byte
}

// type PhotoData struct {

// }
