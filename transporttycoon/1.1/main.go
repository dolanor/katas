package main

type Location struct {
	Name   string
	Parent *Location
	Next   map[string]int
}

var (
	Port    = "Port"
	Factory = "Factory"
	A       = "Warehouse A"
	B       = "Warehouse B"

	LFactory = &Location{
		Name: Factory,
	}

	LPort = &Location{
		Name:   Port,
		Parent: LFactory,
	}

	LA = &Location{
		Name:   A,
		Parent: LPort,
	}

	LB = &Location{
		Name:   B,
		Parent: LFactory,
	}

	locations = map[string]*Location{
		Port:    LPort,
		Factory: LFactory,
		A:       LA,
		B:       LB,
	}

	paths = Location{
		Name: "Factory",
		Next: map[string]int{
			Port: 1,
			B:    5,
		},
	}
)

func main() {

	// AABABBAB
	//      A B         B
	// C    +<+++++<<<<<+++++<<<<<
	//      A A B         A B
	// C    +<+<+++++<<<<<+<+++++<<<<<
	//
	// Bat   ++++<<<<++++<<<<++++<<<<++++<<<<
}
