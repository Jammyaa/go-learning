package main

import (
	"fmt"
	"example.com/math"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("Hello"))
	fmt.Println(math.Add(3,4))
	fmt.Println(math.Sub(3,4))
}
