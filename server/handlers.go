package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// FrontendHandler handles the UI requests
var FrontendHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	h := http.StripPrefix("/", http.FileServer(http.Dir(guiPath)))
	h.ServeHTTP(w, r)
})

// GetBoardHandler handles the API requests
var GetBoardHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	board, err := boardDAO.FindByID(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	apiResponse := APIResponse{
		Ok:   true,
		Data: board,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(&apiResponse)
	// fmt.Fprintf(w, `{"ok":true,"data":{"":""}}`)

})

// AddBoardHandler for adding a new Board
var AddBoardHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})
