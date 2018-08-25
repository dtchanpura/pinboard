package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	// Statik FileSystem import
	_ "github.com/dtchanpura/pinboard/statik"
	"github.com/rakyll/statik/fs"
)

var (
	configuration config
	boardDAO      BoardDAO
	host          string
	port          int
	statikFS      http.FileSystem
	err           error
)

func init() {
	statikFS, err = fs.New()
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
	// guiPath = configuration.GUIPath

	boardDAO.Connect()
}
