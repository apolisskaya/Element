package chessPieces

// definition of each chess piece
// value, color, candidate moves, is pinned, etc

type Piece struct {
	value int
	color, name string
	IsPinned bool
}

type Move struct {

	Delta_x, Delta_y int
	Piece Piece
}

func IsLegalMove(move Move) bool {
	if !move.Piece.IsPinned {
		return true
	}

	return false
}