package main

import "fmt"

func main() {
	//buffered channel (canal con bufer)
	ca := make(chan int)

	go func() {
		ca <- 42
	}()

	fmt.Println(<-ca)
}
