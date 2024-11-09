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
	grid.Next()

	got := grid.String()
	assert(t, got, exp)
}

func TestGridRule1_2(t *testing.T) {
	exp := `
·····
·11··
·····
·····
·····`

	grid := NewGrid()
	grid.Init(false, false, false, false, false, false, true, true)

	got := grid.String()
	assert(t, got, exp)
}

func TestGridRule2(t *testing.T) {
	exp := `
1·1··
1·1··
·····
·····
·····`
	grid := NewGrid()
	grid.Init(true, true, true, false, false,
		true, true)

	got := grid.String()
	assert(t, got, exp)
}

func assert(t *testing.T, got, exp string) {
	t.Helper()
	if got != exp {
		t.Fatalf(`got:
%s

exp:
%s`, got, exp)
	}

}
