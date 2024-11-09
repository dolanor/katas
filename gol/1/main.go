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
	for i, v := range g.boxes {
		if i%5 == 0 {
			s += "\n"
		}
		if v {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

func (g *Grid) Init(b ...bool) {
	g.boxes = b
}

func main() {
}
