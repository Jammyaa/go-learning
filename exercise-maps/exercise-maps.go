/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var str = strings.Fields(s)
	m := make(map[string]int)
	for i := 0; i < len(str); i++ {
		m[str[i]] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
