package main

import (
	"fmt"
)

type Persona struct {
	nombre string
}

type Humano interface {
	hablar()
}

func (p *Persona) hablar() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	p := Persona{
		nombre: "Nicolas",
	}

	p2 := &Persona{
		nombre: "Esteban",
	}
	p.hablar()
	diAlgo(&p)
	diAlgo(p2)

	fmt.Println("PUEDES pasar un valor de tipo *persona a diAlgo")
	fmt.Println("NO puedes pasar un valor de tipo persona a diAlgo")
}

func diAlgo(h Humano) {
	h.hablar()
}
