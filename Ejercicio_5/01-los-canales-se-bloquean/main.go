package main

import "fmt"

func main() {
	//unbuffered channel (canal sin bufer)
	ca := make(chan int)

	ca <- 42

	fmt.Println(<-ca)
}
