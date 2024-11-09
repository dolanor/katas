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

	if got != exp {
		t.Fatalf(`got:
%s

exp:
%s`, got, exp)
	}
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

	got := grid
	if got != exp {
		t.Fatalf(`got:
%s

exp:
%s`, got, exp)
	}
}
