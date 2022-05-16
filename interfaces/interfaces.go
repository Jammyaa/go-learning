/*
Interfaces
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
*/
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	fmt.Printf("%T\n", a)

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	fmt.Println(v.Abs()) // a MyFloat implements Abser
	fmt.Println(f.Abs()) // a *Vertex implements Abser

	a = f
	a = &v

	fmt.Println(a.Abs())
}
