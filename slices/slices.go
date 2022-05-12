package main

import "fmt"

func main() {

	// names is a array of strings
	names := [4]string{
		"Chinh",
		"Linh",
		"Vinh",
		"Minh",
	}

	fmt.Println(names)

	a := names[0:2] // This is a slice
	b := names[1:3] // This is a slice

	fmt.Println(a, b)

	b[0] = "Son"

	fmt.Println(a, b)
	fmt.Println(names)

	// A slice literal is like an array literal without the length.
	q := []int{2, 4, 5, 6, 7, 8, 9}
	fmt.Println(q)
	r := []bool{true, false, true, false, true, false, false}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
	}
	fmt.Println(s)
}
