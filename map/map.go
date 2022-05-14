// A map maps keys to values.

// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // Keys are string, value are Vertex

func main() {
	// The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]Vertex)
	m["Notional Labs"] = Vertex{
		100, 200,
	}
	m["Chinh"] = Vertex{
		200, 300,
	}
	fmt.Println(m)
	fmt.Println(m["Notional Labs"])

	// map literals
	var m2 = map[string]Vertex{
		"Chinh": Vertex{
			200, 300,
		},
		"Vinh": Vertex{
			400, 500,
		},
		"Son": Vertex{
			500, 600,
		},
	}

	fmt.Println(m2)

	// map literals continued
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m3 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	// Mutating Maps
	m4 := make(map[string]int)
	m4["answer"] = 48

	// Test that a key is present with a two-value assignment: elem, ok = m[key]
	v, ok := m4["answer"]
	fmt.Println(v, ok)

	// Delete an element: delete(m, key)
	delete(m4, "answer")
	v, ok = m4["answer"]
	fmt.Println(v, ok)
}
