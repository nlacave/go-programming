package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type prueba struct {
		datos     []int
		resultado int
	}
	p := []prueba{
		{[]int{1, 2, 3}, 6},
		{[]int{5, 2, 1}, 8},
		{[]int{9, 1, 9}, 19},
	}

	for _, v := range p {
		if Sum(v.datos...) == v.resultado {
			fmt.Println("Operacion pasó el test:", v.datos, ":", v.resultado)
		} else {
			t.Error("Expected", v.resultado, "Got", Sum(v.datos...))
		}
	}
	fmt.Println("Corrió correctamente el programa.")
}
