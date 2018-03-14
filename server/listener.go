package server

import (
	"fmt"
	"net/http"
)

var (
	guiPath = "./gui"
	host    = "0.0.0.0"
	port    = 8080
)

// StartListener for starting the listener
func StartListener() {
	router := http.NewServeMux()

	router.Handle("/", FrontendHandler)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: router,
	}

	server.ListenAndServe()

}

// FrontendHandler handles the UI requests
var FrontendHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	h := http.StripPrefix("/", http.FileServer(http.Dir(guiPath)))
	h.ServeHTTP(w, r)
})
