package main

import (
	"fmt"
	"time"
)

func tomarCafe() {
	fmt.Println("Bienvenido a Starbucks Coffee! Que café quieres tomar?")
	var cafe string
	_, err := fmt.Scanln(&cafe)
	if err != nil {
		fmt.Println("Ocurrio un error inesperado", err)
	}
	fmt.Println()
	precio := len(cafe) * 2
	fmt.Println("Son " + "$" + fmt.Sprint(precio) + " por el cafe.")
	fmt.Println("Esperando dinero...")
	var pago int
	_, err = fmt.Scanln(&pago)
	for pago != precio {
		fmt.Println("Por favor, corrobora tu pago...")
		_, err = fmt.Scanln(&pago)
		if err != nil {
			fmt.Println("Ocurrio un error inesperado", err)
		}
		if pago == precio {
			break
		}
	}
	if err != nil {
		fmt.Println("Ocurrio un error inesperado", err)
	}
	fmt.Println("Genial! Ya comenzamos a preparar tu cafe. Por favor espera unos momentos.")
	timer := time.NewTimer(10 * time.Second)
	<-timer.C
	fmt.Println("Listo! Aqui tienes tu café", cafe, "! Que lo disfrutes!")
}

func main() {
	tomarCafe()
}
