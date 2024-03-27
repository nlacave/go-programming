package main

import (
	"Ejercicio_7/mymath"
	"Ejercicio_7/perro"
	"fmt"
)

func main() {
	edadPerruna := perro.Edad(9)
	fmt.Println(edadPerruna)
	sumaDeEdades := mymath.Sum(int(edadPerruna), 10)
	fmt.Println(sumaDeEdades)
}
