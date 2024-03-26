// Package mymath provides math solutions
package mymath

//Sum suma un numero ilimitado de valores de tipo int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
