package server

import (
	"time"

	bson "gopkg.in/mgo.v2/bson"
)

// APIResponse type for responding to API Requests
type APIResponse struct {
	Ok           bool        `json:"ok"`
	Data         interface{} `json:"data,omitempty"`
	Reload       bool        `json:"reload,omitempty"`
	ErrorMessage string      `json:"err,omitempty"`
}

// APIBoardRequest type for structuring board related requests
type APIBoardRequest struct {
	Method string `json:"method"`
	Board  Board  `json:"board,omitempty"`
}

// APIBlockRequest type for structuring block related requests
type APIBlockRequest struct {
	Method string  `json:"method"`
	Block  Block   `json:"block,omitempty"`
	Blocks []Block `json:"blocks,omitempty"`
}

// Board type for structuring multiple blocks
type Board struct {
	ID       bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Title    string        `bson:"title" json:"title"`
	Owner    string        `bson:"owner" json:"owner"`
	Blocks   []Block       `bson:"blocks" json:"blocks"`
	Created  time.Time     `bson:"created,omitempty" json:"created,omitempty"`
	Modified time.Time     `bson:"modified,omitempty" json:"modified,omitempty"`
}

// Block type for showing details about a block
type Block struct {
	ID         bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Title      string        `bson:"title" json:"title"`
	Paragraphs []Paragraph   `bson:"paragraphs" json:"paragraphs"`
	Side       string        `bson:"side" json:"side"`
	Created    time.Time     `bson:"created,omitempty" json:"created,omitempty"`
	Modified   time.Time     `bson:"modified,omitempty" json:"modified,omitempty"`
}

// Paragraph type as string
type Paragraph string
