package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(2, 3))
}

func Sum(xi ...int) int {
	acum := 0
	for _, v := range xi {
		acum += v
	}
	return acum
}
