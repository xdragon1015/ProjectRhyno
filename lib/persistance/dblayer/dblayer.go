package dblayer

import (
	"ProjectRhyno/lib/persistance"
	"ProjectRhyno/lib/persistance/mongolayer"
)

//DATABASETYPE Type for database names
type DATABASETYPE string

//Constant varaibles for database names like mongo, dynamo, Mysql, sqlite
const (
	MONGODB DATABASETYPE = "mongodb"
)

//NewPersistanceLayer allows us to seemlesly change databases if we need to. Recomeneded database: Dynamo
func NewPersistanceLayer(options DATABASETYPE, connection string) (persistance.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	//Might add more cases later
	return nil, nil
}
