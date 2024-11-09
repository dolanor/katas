package main

import (
	"fmt"
	"log"
)

func NewGrid() game {
	return MatrixToGame(NewMatrix())
}

func NewMatrix() matrix {
	return matrix{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
}

func MatrixToGame(m matrix) game {
	var g game = make(game, 5)
	for i, line := range m {
		g[i] = make([]rune, 5)
		for c, cell := range line {
			g[i][c] = itocell(cell)
			fmt.Println(g[i][c])
		}
	}
	return g
}

func itocell(i int) rune {
	if i == 0 {
		log.Println("0")
		return '·'
	} else {

		log.Println("1")
		return '\u25a1'
	}
}

type matrix [][]int

type game [][]rune

func (g game) String() string {
	var b []byte
	for _, line := range g {
		for _, cell := range line {
			b = append(b, byte(cell))
		}
		b = append(b, byte('\n'))
	}
	return string(b)
}
