package main

import (
	"fmt"
	"./version"
	"./chessPieces"
	"./chessBoard"
)

func main() {
	fmt.Printf("Welcome to Element Version " + version.CurrentVersion())

	// initialize chess board
	chessBoard.InitializeBoard()

	DummyMove := chessPieces.Move{
		Delta_x: 1,
		Delta_y: 4,
		Piece: chessPieces.Piece{
			IsPinned: false,
		},
	}

	fmt.Print(chessPieces.IsLegalMove(DummyMove))
}