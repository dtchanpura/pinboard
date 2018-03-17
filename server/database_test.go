package server

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func TestUpdateBoard(t *testing.T) {
	t.SkipNow()
	boardDAO.Connect()
	board := Board{

		Blocks: []Block{Block{
			ID:         bson.NewObjectId(),
			Paragraphs: []Paragraph{"Hello World 12345678"},
			Title:      "Title 1",
		}},
		Title: "Board 1 Test",
		ID:    bson.NewObjectId(),
	}
	fmt.Println(boardDAO.Insert(board))

	board.Blocks = append(board.Blocks, Block{ID: bson.NewObjectId(), Paragraphs: []Paragraph{"Updated Paragraph"}})
	err := boardDAO.UpdateBoard(board)
	fmt.Println(err)
}
func TestAddDeleteBlock(t *testing.T) {
	t.SkipNow()
	boardID := "5aaab0321def2d41da352ee9"
	// blockID := "5aaab0321def2d41da352eea"
	blockID := bson.NewObjectId()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	paragraph := Paragraph(fmt.Sprintf("Hello World Test %d", r1.Intn(10000)))
	err := boardDAO.AddBlocksToBoard(boardID, Block{
		ID:         blockID,
		Paragraphs: []Paragraph{paragraph},
		Title:      "Title 1",
	})

	if err != nil {
		fmt.Println(err)
	}
	board, err := boardDAO.FindByID(boardID)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(len(board.Blocks))
	if findBlockIDInBoard(blockID.Hex(), board) == -1 {
		// Check for given block (it should exist)
		t.Fail()
	}

	boardDAO.RemoveBlockFromBoard(blockID.Hex(), boardID)
	board, err = boardDAO.FindByID(boardID)
	if err != nil {
		fmt.Println(err)
	}

	if findBlockIDInBoard(blockID.Hex(), board) != -1 {
		// Check for given block (it should not exist)
		t.Fail()
	}

}

func TestInvalidBoardID(t *testing.T) {
	boardID := "5aaab0321def2d41da352eex" // invalid id
	_, err := boardDAO.FindByID(boardID)
	if err.Error() != "invalid id" {
		t.Fail()
	}

	err = boardDAO.AddBlocksToBoard(boardID, Block{
		ID:         bson.NewObjectId(),
		Paragraphs: []Paragraph{"paragraph"},
		Title:      "Title 1",
	})
	if err.Error() != "invalid id" {
		t.Fail()
	}

	err = boardDAO.RemoveBlockFromBoard(bson.NewObjectId().Hex(), boardID)
	if err.Error() != "invalid id" {
		t.Fail()
	}
}

func TestUpdateBlockInBoard(t *testing.T) {
	t.SkipNow()
	boardID := bson.NewObjectId()
	blockID := bson.NewObjectId()
	block := Block{
		ID:         blockID,
		Title:      fmt.Sprintf("Test Title: %d", time.Now().Unix()),
		Paragraphs: []Paragraph{"Hello"},
	}
	board := Board{
		Blocks: []Block{
			block,
		},
		ID:    boardID,
		Title: "Board Title",
	}
	err := boardDAO.Insert(board)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	block.Title = fmt.Sprintf("Test Title: %d", 0)

	boardDAO.UpdateBlockInBoard(block, boardID.Hex())
	boardDAO.Delete(boardID)
}

func TestFindAll(t *testing.T) {
	t.SkipNow()
	boards, err := boardDAO.FindAll()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	for _, board := range boards {
		fmt.Println(board.ID, board.Title)
	}
}
