package main

import (
	"fmt"
	"go-programming/Ejercicio_3/gato"
)

func main() {
	Comer()
	Saludar()
	gato.Saludar()
	gato.Comer()
}

func Saludar() {
	fmt.Println("Guauuuuuuuuu")
}

func Comer() {
	fmt.Println("El perro está comiendo")
}
