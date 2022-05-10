package main

import "fmt"

func main() {
	var a [2]string // An array's length is part of its type, so arrays cannot be resized
	a[0] = "Hello"
	a[1] = "Chinh"
	fmt.Println(a[0], a[1])
	a[1] = "Chinhnotional"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

}
