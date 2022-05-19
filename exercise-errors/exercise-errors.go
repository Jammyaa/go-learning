/*
Exercise: Errors
Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type:
type ErrNegativeSqrt float64

and make it an error by giving it a
func (e ErrNegativeSqrt) Error() string

method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
*/

package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
