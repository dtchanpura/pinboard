package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// StartListener for starting the listener
func StartListener() {
	// router := http.NewServeMux()
	router := mux.NewRouter()

	router.Handle("/board/{id}", GetBoardHandler).Methods(http.MethodGet)
	router.Handle("/board", AddBoardHandler).Methods(http.MethodPost)
	router.PathPrefix("/").Handler(FrontendHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: router,
	}

	server.ListenAndServe()
}
