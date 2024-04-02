// Package math_test provides math solutions.
package math_test

//Sum suma un numero ilimitado de numeros enteros.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
