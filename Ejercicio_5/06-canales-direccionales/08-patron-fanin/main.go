package main

import (
	"fmt"
	"sync"
)

type ValorConOrigen struct {
	Valor  int
	Origen string
}

func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan ValorConOrigen)

	go enviar(par, impar)
	go recibir(par, impar, fanin)

	for obj := range fanin {
		fmt.Printf("Valor: %d, Origen: %s\n", obj.Valor, obj.Origen)
	}

}

func enviar(par, impar chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

func recibir(par, impar <-chan int, fanin chan<- ValorConOrigen) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for v := range par {
			fanin <- ValorConOrigen{
				Valor:  v,
				Origen: "Par",
			}
		}
		wg.Done()
	}()

	go func() {
		for v := range impar {
			fanin <- ValorConOrigen{
				Valor:  v,
				Origen: "Impar",
			}
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
