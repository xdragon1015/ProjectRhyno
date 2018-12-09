package configuration

import (
	"ProjectRhyno/lib/persistance/dblayer"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Configuration struct {
	DatabaseType    dblayer.DATABASETYPE
	DBConnection    string
	RestfulEndpoint string
}

var (
	DBTYPEDefault       = dblayer.DATABASETYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1:27017"
	RestfulEPDefault    = "localhost:8080"
)

func ExtraConfig(filename string) (Configuration, error) {
	config := Configuration{DatabaseType: DBTYPEDefault, DBConnection: DBConnectionDefault, RestfulEndpoint: RestfulEPDefault}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("configuratio file not found")
	} else {
		fmt.Println("Configuration file found")
		fmt.Println("Processing Configuration file")
		time.Sleep(time.Millisecond * 5000)
		fmt.Println("File Processed")
	}

	err = json.NewDecoder(file).Decode(&config)
	return config, err

}
