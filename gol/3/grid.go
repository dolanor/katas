package main

type Grid struct {
	boxes []bool
}

func NewGrid() Grid {
	return Grid{
		boxes: make([]bool, 25),
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

}
