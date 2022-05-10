package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	v3.Y = 1e9
	fmt.Println(v1, v2, v3, p)
	fmt.Printf("%T %T %T %T\n", v1, v2, v3, p)
}
