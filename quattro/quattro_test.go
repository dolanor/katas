package quattro_test

import (
	"testing"

	. "github.com/dolanor/katas/quattro"
)

func TestEmptyBoard(t *testing.T) {
	emptyBoard := NewBoard()

	if emptyBoard.Win() {
		t.Errorf("empty board shouldn't win")
	}
}

func TestBoardWith4SmallPiecesOnRow1Wins(t *testing.T) {
	board := NewBoard()

	p1 := Piece{Size: Small, Color: White, Shape: Circle, Filled: Plain}
	p2 := Piece{Size: Small, Color: Black, Shape: Circle, Filled: Plain}
	p3 := Piece{Size: Small, Color: White, Shape: Square, Filled: Plain}
	p4 := Piece{Size: Small, Color: White, Shape: Circle, Filled: Hollow}
	board.Put(p1, 0, 0)
	board.Put(p2, 0, 1)
	board.Put(p3, 0, 2)
	board.Put(p4, 0, 3)

	if !board.Win() {
		t.Errorf("board with 4 small pieces should win")
	}

}

func TestBoardWithOnly3PiecesOnRow1DoesNotWin(t *testing.T) {
	board := NewBoard()

	p1 := Piece{Size: Small, Color: White, Shape: Circle, Filled: Plain}
	p2 := Piece{Size: Small, Color: Black, Shape: Circle, Filled: Plain}
	p3 := Piece{Size: Small, Color: White, Shape: Square, Filled: Plain}
	board.Put(p1, 0, 0)
	board.Put(p2, 0, 1)
	board.Put(p3, 0, 2)

	if board.Win() {
		t.Errorf("board with only 3 pieces should not win")
	}
}

func TestBoardWith4SmallPiecesNotOnSameRowCantWin(t *testing.T) {
	board := NewBoard()

	p1 := Piece{Size: Small, Color: White, Shape: Circle, Filled: Plain}
	p2 := Piece{Size: Small, Color: Black, Shape: Circle, Filled: Plain}
	p3 := Piece{Size: Small, Color: White, Shape: Square, Filled: Plain}
	p4 := Piece{Size: Small, Color: White, Shape: Circle, Filled: Hollow}
	board.Put(p1, 0, 0)
	board.Put(p2, 0, 1)
	board.Put(p3, 0, 2)
	board.Put(p4, 1, 3)

	if board.Win() {
		t.Errorf("board with 4 small pieces not on same row should not win")
	}
}

func TestProperties(t *testing.T) {
	p1 := Piece{Size: Small, Color: White, Shape: Square, Filled: Plain}
	p2 := Piece{Size: Large, Color: Black, Shape: Square, Filled: Hollow}
	p3 := Piece{Size: Large, Color: Black, Shape: Square, Filled: Hollow}

	place := func(p Piece) PlacedPiece {
		return PlacedPiece{
			Piece: p,
		}
	}

	pp := PlacedPieces{place(p1), place(p2), place(p3)}

	if !pp.Similarities() {
		t.Errorf("no similarities found")
	}
}
