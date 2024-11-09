package main

import "testing"

func TestGrid(t *testing.T) {
	grid := NewGrid()
	exp := `
·····
·····
·····
·····
·····`
	got := grid.String()

	if got != exp {
		t.Fatalf(`got:
%s

exp:
%s`, got, exp)
	}
}
