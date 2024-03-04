package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go agregar(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func agregar(canal chan int) {
	for i := 0; i < 20; i++ {
		canal <- i
	}
	close(canal)
}

func fanOutIn(canal, canal2 chan int) {
	var wg sync.WaitGroup
	for v := range canal {
		wg.Add(1)
		go func(valor int) {
			canal2 <- trabajoConsumeTiempo(valor)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabajoConsumeTiempo(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
