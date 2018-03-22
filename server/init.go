package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	configuration config
	boardDAO      BoardDAO
	host          string
	port          int
	guiPath       string
)

func init() {
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(raw, &configuration)
	if err != nil {
		log.Fatal(err)
	}
	boardDAO.Database = configuration.Database
	boardDAO.Server = configuration.Server
	host = configuration.Host
	port = configuration.Port
	guiPath = configuration.GUIPath

	boardDAO.Connect()
}
