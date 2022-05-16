/*
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
*/

package main

import "fmt"

type I interface {
	M()
	N()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}
func (t T) N() {
	fmt.Printf("%T\n", t.S)
}

func main() {
	var i I = T{"Chinh yeu Linh"}
	i.M()
	i.N()
}
