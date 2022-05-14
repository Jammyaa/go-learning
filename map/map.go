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
}
