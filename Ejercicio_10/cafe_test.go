package cafe

import (
	"testing"
)

func TestPedirCafe(t *testing.T) {
	type prueba struct {
		datos     string
		resultado int
	}

	pruebas := []prueba{
		{"Expresso Macchiato", 36},
		{"Capuccino con Crema", 38},
		{"Latte", 10},
	}

	for _, x := range pruebas {
		v := PedirCafe(x.datos)
		if v != x.resultado {
			t.Error("Expected:", x.resultado, "Got:", v)
		}
	}
}
