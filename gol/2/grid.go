package main

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
		}
	}
	return g
}

func itocell(i int) rune {
	if i == 0 {
		return '\u00b7'
	} else {

		return '\u25a1'
	}
}

type matrix [][]int

type game [][]rune

func (g game) String() string {
	var b []rune
	for _, line := range g {
		b = append(b, rune('\n'))
		for _, cell := range line {
			b = append(b, rune(cell))
		}
	}
	return string(b)
}
