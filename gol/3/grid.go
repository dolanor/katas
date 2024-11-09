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
