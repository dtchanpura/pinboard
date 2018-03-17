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

	// Operations related to given board -> blocks (Update, Delete)
	router.Handle("/board/{boardID}/block/{blockID}", UpdateBlockHandler).Methods(http.MethodPut)
	router.Handle("/board/{boardID}/block/{blockID}", DeleteBlockHandler).Methods(http.MethodDelete)

	// Add new block to board
	router.Handle("/board/{boardID}/block", AddBlockHandler).Methods(http.MethodPost)

	// Useless...
	// router.Handle("/board/{boardID}/block", GetBoardHandler).Methods(http.MethodGet)

	// Operations related to given board ID (Get, Update, Delete)
	router.Handle("/board/{boardID}", GetBoardHandler).Methods(http.MethodGet)
	router.Handle("/board/{boardID}", UpdateBoardHandler).Methods(http.MethodPut)
	router.Handle("/board/{boardID}", DeleteBoardHandler).Methods(http.MethodDelete)

	// Add new board
	router.Handle("/board", AddBoardHandler).Methods(http.MethodPost)
	// Get all boards
	router.Handle("/board", GetBoardHandler).Methods(http.MethodGet)

	// GUI Rendering
	router.PathPrefix("/").Handler(FrontendHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: router,
	}

	server.ListenAndServe()
}
