package main

import "testing"

func TestGridEmpty(t *testing.T) {
	grid := NewGrid()
	if len(grid.boxes) != 25 {
		t.Fatal("grid size is not 25")
	}
}

func TestGrid_String(t *testing.T) {
	grid := NewGrid()
	exp := `
00000
00000
00000
00000
00000`
	got := grid.String()

	if got != exp {
		t.Fatalf("grid misrepresented, expected: \n%s\n\ngot:\n%s", exp, got)
	}

}

func TestGridRule1(t *testing.T) {
	grid := NewGrid()

	grid.Init(true, true)
	exp := `
11000
00000
00000
00000
00000`
	got := grid.String()

	if got != exp {
		t.Fatalf("2 live cells didn't survive, expected: \n%s\n\ngot:\n%s", exp, got)
	}

}
