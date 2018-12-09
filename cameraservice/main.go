package main

import (
	"ProjectRhyno/cameraservice/rest"
	"ProjectRhyno/lib/configuration"
	"ProjectRhyno/lib/persistance/dblayer"
	"flag"
	"fmt"
)

func main() {
	configPath := flag.String("conf", "../lib/configuration/config.json", "flag to get the path to the configuration file")
	flag.Parse()
	config, _ := configuration.ExtraConfig(*configPath)
	fmt.Println("Connecting to database")
	dbHandler, err := dblayer.NewPersistanceLayer(config.DatabaseType, config.DBConnection)
	if err != nil {
		fmt.Println(err)
	}
	rest.ServerAPI(config.RestfulEndpoint, dbHandler)
}
