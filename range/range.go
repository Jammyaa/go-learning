package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 25, 512, 1024}

	// When ranging over a slice, two values are returned for each iteration. The first "i" is the index, and the second "v" is a copy of the element at that index.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 11)

	// If you only want the index, you can omit the second variable.
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}

	// You can skip the index or value by assigning to _
	for _, v := range pow2 {
		fmt.Printf("%d\n", v)
	}
}
