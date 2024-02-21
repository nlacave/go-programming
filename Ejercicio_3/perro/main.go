package main

import (
	"fmt"
	"github.com/nlacave/go-programming/Ejercicio_3/gato"
)
func main() {
	comer()
	saludar()
	gato.Saludar()
	gato.Comer()
}

func saludar() {
	fmt.Println("Guauuuuuuuuu")
}

func comer() {
	fmt.Println("El perro está comiendo")
}

