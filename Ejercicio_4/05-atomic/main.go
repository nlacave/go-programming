package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var incremento int64
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&incremento, 1)
			fmt.Println(atomic.LoadInt64(&incremento))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Numero final de Goruoutines:", runtime.NumGoroutine())
	fmt.Println("El valor final del incremento es:", incremento)
}
