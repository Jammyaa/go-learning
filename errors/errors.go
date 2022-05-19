/*
Errors
Go programs express error state with error values.
The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
A nil error denotes success; a non-nil error denotes failure.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := "10000a"

	v, err := strconv.Atoi(i)

	if err != nil {
		fmt.Printf("could not convert %v to int because this error: %v\n", i, err)
		return
	}

	fmt.Println(v)
}
