package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	incremento := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			in := incremento
			runtime.Gosched()
			in++
			incremento = in
			fmt.Println(incremento)
			wg.Done()
		}()
		fmt.Println("Numero de Goruoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Numero final de Goruoutines:", runtime.NumGoroutine())
}
