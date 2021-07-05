package config

import (
	"errors"
	"os"
)

type configs struct {
	DBURL   string
	APIURL  string
	APIPort string
}

var Production configs
var Development configs

func Initialize() error {
	if os.Args[1] == "dev" {
		Development.DBURL = "mongodb://localhost:27017"
		Development.APIURL = "http://localhost"
		Development.APIPort = "3000"
		return nil
	} else if os.Args[1] == "prod" {
		Production.DBURL = "mongodb://localhost:27017"
		Production.APIURL = "http://localhost"
		Production.APIPort = "3000"
		return nil
	}
	return errors.New("Check Your Runtime Args")
}
