package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func respondWithError(err error, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&APIResponse{Ok: false, ErrorMessage: err.Error()})
}

func respondWithJSON(payload interface{}, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&APIResponse{Ok: true, Data: payload})
}

// FrontendHandler handles the UI requests
var FrontendHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	h := http.StripPrefix("/", http.FileServer(http.Dir(guiPath)))
	h.ServeHTTP(w, r)
})

// GetBoardHandler handles the API requests
var GetBoardHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if boardID := params["boardID"]; boardID != "" {
		board, err := boardDAO.FindByID(boardID)
		if err != nil {
			respondWithError(err, http.StatusNotFound, w)
			return
		}
		respondWithJSON(board, http.StatusOK, w)
	} else {
		boards, err := boardDAO.FindAll()
		if err != nil {
			respondWithError(err, http.StatusInternalServerError, w)
			return
		}
		respondWithJSON(boards, http.StatusOK, w)
	}
	// fmt.Fprintf(w, `{"ok":true,"data":{"":""}}`)

})

// AddBoardHandler for adding a new Board
var AddBoardHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var jsonData APIBoardRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	fmt.Println(jsonData.Board)
	updateBoard := Board{ID: bson.NewObjectId(), Title: jsonData.Board.Title, Blocks: []Block{}}
	if len(jsonData.Board.Blocks) > 0 {
		for _, block := range jsonData.Board.Blocks {
			block.ID = bson.NewObjectId()
			updateBoard.Blocks = append(updateBoard.Blocks, block)
		}
	}
	err := boardDAO.Insert(updateBoard)
	if err != nil {
		respondWithError(err, http.StatusInternalServerError, w)
	} else {
		respondWithJSON(updateBoard, http.StatusOK, w)
	}
})

// AddBlockHandler for adding a new Block in Board
// Only POST allowed
var AddBlockHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var jsonData APIBlockRequest
	params := mux.Vars(r)
	if boardID, ok := params["boardID"]; ok {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
			respondWithError(err, http.StatusBadRequest, w)
		}
		if jsonData.Block.Title != "" {
			jsonData.Block.ID = bson.NewObjectId()
			boardDAO.AddBlocksToBoard(boardID, jsonData.Block)
			respondWithJSON(jsonData.Block, http.StatusOK, w)
		}
		if len(jsonData.Blocks) > 0 {
			for i := range jsonData.Blocks {
				jsonData.Blocks[i].ID = bson.NewObjectId()
			}
			boardDAO.AddBlocksToBoard(boardID, jsonData.Blocks...)
			respondWithJSON(jsonData.Blocks, http.StatusOK, w)
		}
	} else {
		respondWithError(errors.New("invalid board id"), http.StatusBadRequest, w)
	}
})

// UpdateBoardHandler for updating Board
// Only PUT allowed
// Mostly for changing title, owner
var UpdateBoardHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var (
		boardID  string
		ok       bool
		jsonData APIBoardRequest
	)
	if boardID, ok = params["boardID"]; !ok {
		respondWithError(errors.New("boardID invalid"), http.StatusBadRequest, w)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	jsonData.Board.ID = bson.ObjectIdHex(boardID)
	err := boardDAO.UpdateBoardDetails(jsonData.Board)
	if err != nil {
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	respondWithJSON(jsonData.Board, http.StatusOK, w)
	// jsonData.Board
})

// DeleteBoardHandler for updating Board
// Only PUT allowed
// Mostly for changing title, owner
var DeleteBoardHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var (
		boardID string
		ok      bool
	)
	if boardID, ok = params["boardID"]; !ok {
		respondWithError(errors.New("boardID invalid"), http.StatusBadRequest, w)
		return
	}
	if !bson.IsObjectIdHex(boardID) {
		respondWithError(errors.New("invalid board id"), http.StatusBadRequest, w)
		return
	}
	err := boardDAO.Delete(bson.ObjectIdHex(boardID))
	if err != nil {
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	respondWithJSON(APIResponse{Ok: true, Data: "deleted"}, http.StatusOK, w)
	// jsonData.Board
})

// UpdateBlockHandler for updating Block in Board
// Only PUT allowed
var UpdateBlockHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var (
		boardID  string
		blockID  string
		ok       bool
		jsonData APIBlockRequest
	)

	if boardID, ok = params["boardID"]; !ok {
		respondWithError(errors.New("boardID invalid"), http.StatusBadRequest, w)
		return
	}
	if blockID, ok = params["blockID"]; !ok {
		respondWithError(errors.New("blockID invalid"), http.StatusBadRequest, w)
		return
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	jsonData.Block.ID = bson.ObjectIdHex(blockID)
	err := boardDAO.UpdateBlockInBoard(jsonData.Block, boardID)
	if err != nil {
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	respondWithJSON(jsonData.Block, http.StatusOK, w)
})

// DeleteBlockHandler for updating Block in Board
// Only DELETE allowed
var DeleteBlockHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var (
		boardID  string
		blockID  string
		ok       bool
		jsonData APIBlockRequest
	)

	if boardID, ok = params["boardID"]; !ok {
		respondWithError(errors.New("boardID invalid"), http.StatusBadRequest, w)
		return
	}
	if blockID, ok = params["blockID"]; !ok {
		respondWithError(errors.New("blockID invalid"), http.StatusBadRequest, w)
		return
	}
	defer r.Body.Close()
	err := boardDAO.RemoveBlockFromBoard(blockID, boardID)
	if err != nil {
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	respondWithJSON(jsonData.Block, http.StatusOK, w)
})
