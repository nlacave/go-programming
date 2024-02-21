package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Número de CPUS al principio:", runtime.NumCPU())
	fmt.Println("Número de Gorutinas al principio:", runtime.NumGoroutine())
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Valor indice:", i)
		}
		wg.Done()
	}()
	go func() {
		fmt.Println("Estoy imprimiendooo!")
		wg.Done()
	}()

	fmt.Println("Número de CPUS en el medio:", runtime.NumCPU())
	fmt.Println("Número de Gorutinas en el medio:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Número de CPUS al final:", runtime.NumCPU())
	fmt.Println("Número de Gorutinas al final:", runtime.NumGoroutine())
}
