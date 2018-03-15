package server

import (
	bson "gopkg.in/mgo.v2/bson"
)

// APIResponse type for responding to API Requests
type APIResponse struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data,omitempty"`
}

// Board type for structuring multiple blocks
type Board struct {
	Blocks []Block `json:"blocks"`
}

// Block type for showing details about a block
type Block struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Owner      string        `bson:"owner" json:"owner"`
	Title      string        `bson:"title" json:"title"`
	Paragraphs []Paragraph   `bson:"paragraphs" json:"paragraphs"`
	Side       string        `bson:"side" json:"side"`
}

// Paragraph type as string
type Paragraph string
