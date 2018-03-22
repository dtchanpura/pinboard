package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
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
	var (
		routerHandler http.Handler
		headersOk     handlers.CORSOption
	)
	if configuration.CORS != nil && configuration.CORS.Enable {
		// Default CORS configuration
		if configuration.CORS.AllowedHeaders == nil {
			configuration.CORS.AllowedHeaders = []string{"Content-Type", "X-Requested-With"}
		}
		if configuration.CORS.AllowedOrigins == nil {
			configuration.CORS.AllowedOrigins = []string{"*"}
		}
		if configuration.CORS.AllowedMethods == nil {
			configuration.CORS.AllowedMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}
		}

		headersOk = handlers.AllowedHeaders(configuration.CORS.AllowedHeaders)
		originsOk := handlers.AllowedOrigins(configuration.CORS.AllowedOrigins)
		methodsOk := handlers.AllowedMethods(configuration.CORS.AllowedMethods)

		routerHandler = handlers.CORS(originsOk, headersOk, methodsOk)(router)
	} else {
		routerHandler = router
	}
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: routerHandler,
	}
	server.ListenAndServe()
}
