package quattro

import "fmt"

type PlacedPiece struct {
	Piece
	Row int
	Col int
}

type PlacedPieces []PlacedPiece

func (pp *PlacedPieces) Similarities() bool {
	if len(*pp) <= 2 {
		return false
	}

	fmt.Println("=====")
	x := (*pp)[0].Properties()
	fmt.Printf("%04b\n", (*pp)[0].Properties())
	for _, p := range (*pp)[1:] {
		fmt.Printf("%04b\n", p.Properties())
		x = x ^ p.Properties()
	}
	fmt.Println("-----")
	fmt.Printf("%04b | %04b\n", x, (15 ^ x))
	return (15 ^ x) > 0
}

type Board struct {
	pieces PlacedPieces
}

func NewBoard() Board {
	return Board{}
}

func (b *Board) Put(p Piece, x, y int) {
	b.pieces = append(b.pieces, PlacedPiece{Piece: p, Row: x, Col: y})
}

func (b *Board) Win() bool {
	if len(b.pieces) < 4 {
		return false
	}

	lastRow := b.pieces[0].Row
	for _, p := range b.pieces[1:] {
		if lastRow != p.Row {
			return false
		}
	}

	return true
}

type Color int

const (
	White Color = 0x0
	Black       = 0x1
)

type Size int

const (
	Small Size = 0x0
	Large      = 0x2
)

type Shape int

const (
	Circle Shape = 0x0
	Square       = 0x4
)

type Filled int

const (
	Plain  Filled = 0x0
	Hollow        = 0x8
)

type Piece struct {
	Color  Color
	Shape  Shape
	Size   Size
	Filled Filled
}

func (p *Piece) Properties() int {
	return int(p.Color) + int(p.Shape) + int(p.Size) + int(p.Filled)
}
