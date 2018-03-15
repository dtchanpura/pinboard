package server

import (
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

// Insert method inserts a block
func (m *BoardDAO) Insert(board Board) error {
	err := db.C(COLLECTION).Insert(&board)
	return err
}

// FindByID to find a board by ID
func (m *BoardDAO) FindByID(id string) (Board, error) {
	var board Board
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&board)
	return board, err
}
