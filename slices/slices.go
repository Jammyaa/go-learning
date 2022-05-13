package main

import (
	"fmt"
	"strings"
)

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
	fmt.Printf("%T %T\n", s, names)
	// Slice defaults
	v := [4]int{12, 24, 35, 46}
	fmt.Println(v[0:4], v[:], v[:4], v[0:])

	// Slice length and capacity
	t := []int{2, 3, 5, 7, 11, 13}
	printSlice(t)

	// Slice the slice to give it zero length.
	t = t[:0]
	printSlice(t)

	// Extend its length.
	t = t[:4]
	printSlice(t)

	// Drop its first two values.
	t = t[2:]
	printSlice(t)

	// Nil slices
	var nils []int
	fmt.Println(nils, len(nils), cap(nils))

	// Using make function to create a slice
	u := make([]int, 5) // allocate a zeroed array
	printSlice(u)

	w := make([]int, 0, 5) // specify len and cap of slices
	printSlice(w)

	w = w[:cap(w)]
	printSlice(w)

	w = w[1:]
	printSlice(w)

	// Create slices of slices
	// Slices can contain any type, including other slices.
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println(len(board), cap(board))
	fmt.Println(board[1])
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	var nils2 []int
	printSlice(nils2)

	// append works on nil slices.
	nils2 = append(nils2, 0)
	printSlice(nils2)

	// The slice grows as needed.
	nils2 = append(nils2, 1)
	printSlice(nils2)

	// We can add more than one element at a time.
	nils2 = append(nils2, 2, 3, 4)
	printSlice(nils2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
