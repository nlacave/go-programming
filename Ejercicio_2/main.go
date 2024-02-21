package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Arquitectura: %v\nSistema Operativo: %v\n", runtime.GOARCH, runtime.GOOS)
}
