package mongolayer

import (
	"ProjectRhyno/lib/persistance"

	"gopkg.in/mgo.v2/bson"

	"github.com/globalsign/mgo"
)

const (
	DB        = "myPhotos"
	PHOTOS    = "photos"
	PHOTODATA = "data"
)

//MongoDBLayer struct that holds a session field from the database
type MongoDBLayer struct {
	session *mgo.Session
}

//NewMongoDBLayer stablishes new connection to mongoDB
func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	session, err := mgo.Dial(connection)
	return &MongoDBLayer{
		session: session,
	}, err

}

//AddPhoto adds photo object into the collection PHOTOS
func (mongoDB *MongoDBLayer) AddPhoto(photo persistance.Photo) ([]byte, error) {
	s := refreshSession(mongoDB)
	defer s.Close()

	if !photo.ID.Valid() {
		photo.ID = bson.NewObjectId()
	}

	return []byte(photo.ID), s.DB(DB).C(PHOTOS).Insert(photo)
}

//FindPhoto finds object from collection return PHOTO object back to the calling function
//Tak 1 => Implement function FindPhoto from interface and return PHOTO object back to the calling function
func (mongoDB MongoDBLayer) FindPhoto(id []byte) (persistance.Photo, error) {
	return persistance.Photo{}, nil
}

func refreshSession(mongoDB *MongoDBLayer) *mgo.Session {
	s := mongoDB.session.Copy()
	return s
}
