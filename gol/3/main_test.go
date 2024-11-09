package main

import "testing"

func TestGridEmpty(t *testing.T) {
	exp := `
·····
·····
·····
·····
·····`

	grid := NewGrid()
	got := grid.String()

	assert(t, got, exp)
}

func TestGridRule1(t *testing.T) {
	exp := `
11···
·····
·····
·····
·····`

	grid := NewGrid()
	grid.Init(true, true)

	got := grid.String()
	assert(t, got, exp)
}

func TestGridRule1_2(t *testing.T) {
	exp := `
111··
·····
·····
·····
·····`

	grid := NewGrid()
	grid.Init(true, true, true)

	got := grid.String()
	assert(t, got, exp)
}

func assert(t *testing.T, got, exp string) {
	if got != exp {
		t.Fatalf(`got:
%s

exp:
%s`, got, exp)
	}

}
