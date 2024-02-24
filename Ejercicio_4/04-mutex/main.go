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
	mu := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			in := incremento
			in++
			incremento = in
			fmt.Println(incremento)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Numero final de Goruoutines:", runtime.NumGoroutine())
}
