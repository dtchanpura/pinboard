package server

import (
	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BoardDAO is DAO struct containing Server and Database
type BoardDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION as constant collection name
	COLLECTION = "boards"
)

// Connect method establishes a connection to database
func (m *BoardDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Insert method inserts a board
func (m *BoardDAO) Insert(board Board) error {
	err := db.C(COLLECTION).Insert(&board)
	return err
}

// Delete method deletes a board
func (m *BoardDAO) Delete(boardID bson.ObjectId) error {
	err := db.C(COLLECTION).Remove(bson.M{"_id": boardID})
	return err
}

// FindAll method returns all boards
func (m *BoardDAO) FindAll() ([]Board, error) {
	var boards []Board
	err := db.C(COLLECTION).Find(nil).All(&boards)
	return boards, err
}

// FindByID to find a board by ID
func (m *BoardDAO) FindByID(id string) (Board, error) {
	if bson.IsObjectIdHex(id) {
		return m.FindByObjectID(bson.ObjectIdHex(id))
	}
	return Board{}, errors.New("invalid id")
}

// FindByObjectID to find a board by ID
func (m *BoardDAO) FindByObjectID(id bson.ObjectId) (Board, error) {
	var board Board

	err := db.C(COLLECTION).FindId(id).One(&board)
	return board, err

}

// UpdateBoard updates the given board
func (m *BoardDAO) UpdateBoard(board Board) error {
	err := db.C(COLLECTION).Update(bson.M{"_id": board.ID}, board)
	return err
}

// AddBlocksToBoard adds a new block with provided Block object to provided boardID
func (m *BoardDAO) AddBlocksToBoard(boardID string, block ...Block) error {
	board, err := m.FindByID(boardID)
	if err != nil {
		return err
	}
	board.Blocks = append(board.Blocks, block...)
	return m.UpdateBoard(board)
}

// RemoveBlockFromBoard removes a block with provided blockID from board with provided boardID
func (m *BoardDAO) RemoveBlockFromBoard(blockID string, boardID string) error {
	board, err := m.FindByID(boardID)
	if err != nil {
		return err
	}
	blockIndex := findBlockIDInBoard(blockID, board)
	if blockIndex == -1 {
		return errors.New("block not found")
	}
	board.Blocks = append(board.Blocks[:blockIndex], board.Blocks[blockIndex+1:]...)
	return m.UpdateBoard(board)
}

// UpdateBlockInBoard updates a block in board
func (m *BoardDAO) UpdateBlockInBoard(block Block, boardID string) error {
	board, err := m.FindByID(boardID)
	if err != nil {
		return err
	}
	blockIndex := findBlockIDInBoard(block.ID.Hex(), board)
	if blockIndex == -1 {
		return errors.New("block not found")
	}
	board.Blocks[blockIndex] = block
	return m.UpdateBoard(board)
}

func findBlockIDInBoard(blockID string, board Board) int {
	for i, b := range board.Blocks {
		if b.ID.Hex() == blockID {
			return i
		}
	}
	return -1
}
