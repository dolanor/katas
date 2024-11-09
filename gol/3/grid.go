package main

type Grid struct {
	boxes []bool
	width int
}

func NewGrid() Grid {
	return Grid{
		boxes: make([]bool, 25),
		width: 5,
	}
}

func (g Grid) String() string {
	var s string
	for i, b := range g.boxes {
		if i%5 == 0 {
			s += "\n"
		}

		if b {
			s += "1"
		} else {
			s += "·"
		}
	}

	return s
}

func (g *Grid) Init(b ...bool) {
	for i, v := range b {
		g.boxes[i] = v
	}
}

func (g *Grid) Next() {
	for i, b := range g.boxes {
		_, _ = i, b
	}
}

func numNeighbour(g Grid, i int) {
	nw := i - 1 - g.width
	n := i - g.width
	ne := i + 1 - g.width
	w := i - 1
	e := i + 1
	sw := i - 1 + g.width
	s := i + g.width
	se := i + 1 + g.width

	_, _, _, _, _, _, _, _ = nw, n, ne, w, e, sw, s, se
}
