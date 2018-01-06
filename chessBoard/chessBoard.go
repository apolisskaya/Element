package chessBoard

import "fmt"

// properties of a single square
type Square struct {
	value_x, value_y int
	color string
	occupied bool
}

type Board struct {
	board [8][8]Square
}

// returns a 2-dimentional 8x8 array
func InitializeBoard() {

	var ChessBoard Board

	ChessBoard.board[0][0].color = "White"

	fmt.Printf(ChessBoard.board[0][0].color)
}