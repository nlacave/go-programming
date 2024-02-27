package main

import (
	"fmt"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	go enviar(par, impar, salir)

	recibir(par, impar, salir)
	fmt.Println("Finalizando...")
}

func enviar(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	close(s)
}

func recibir(p, i, s <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("Valor par recibido:", v)
		case v := <-i:
			fmt.Println("Valor impar recibido", v)
		case v, ok := <-s:
			if !ok {
				fmt.Println("Valor salir recibido con canal cerrado", v, ok)
				return
			} else {
				fmt.Println("Valor salir recibido con canal abierto", v, ok)
			}
		}
	}
}
