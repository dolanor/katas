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
