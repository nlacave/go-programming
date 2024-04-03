package main

import (
	"Ejercicio_11/mystr"
	"fmt"
	"strings"
)

func main() {
	const str = "Programando en el lenguaje Golang"
	xs := strings.Split(str, " ")
	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Println(mystr.Concatenar(xs))
}
