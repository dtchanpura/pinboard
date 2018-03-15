package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Println(err)
	}
	json.Unmarshal(raw, &configuration)
	boardDAO.Database = configuration.Database
	boardDAO.Server = configuration.Server
	host = configuration.Host
	port = configuration.Port
	guiPath = configuration.GUIPath

	boardDAO.Connect()
}
