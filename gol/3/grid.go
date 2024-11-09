package main

type Grid struct {
	boxes []bool
}

func NewGrid() Grid {
	return Grid{}
}

func (g Grid) String() string {
	return `
·····
·····
·····
·····
·····`
}

func (g *Grid) Init(b ...bool) {
	g.boxes = b
}
