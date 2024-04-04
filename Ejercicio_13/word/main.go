// Package word provides assistance with converting maps and string values.
package word

import "strings"

//UseCount takes a string as input and returns a map where each word is a key associated with its count
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count takes a string as input and returns the number of words contained in that string.
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
