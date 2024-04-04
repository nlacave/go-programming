package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	str := "Hola Mundo Hola"
	got := UseCount(str)
	expected := map[string]int{
		"Hola":  2,
		"Mundo": 1,
	}

	val := true
	for k, v := range expected {
		if v2, ok := got[k]; !ok || v2 != v {
			t.Error("Expected:", expected, "Got:", got)
			val = false
		}
	}
	if val {
		fmt.Println("Prueba satisfactoria!")
	}

	type prueba struct {
		datos     string
		resultado map[string]int
	}
	p := []prueba{
		{"Hola Mundo Hola", map[string]int{
			"Hola":  2,
			"Mundo": 1,
		}},
		{"Chau Mundo Mundo", map[string]int{
			"Chau":  1,
			"Mundo": 2,
		}},
	}

	val = true
	for _, v := range p {
		m1 := UseCount(v.datos)
		m2 := v.resultado
		if !mapsEqual(m1, m2) {
			t.Error("Expected:", v.resultado, "Got:", UseCount(v.datos))
			val = false
		}
	}
	if val {
		fmt.Println("El test pasó la prueba!")
	}
}

func TestCount(t *testing.T) {
	str := "Hola Mundo Hola"
	got := Count(str)
	expected := 3
	if got != expected {
		t.Error("Expected:", expected, "Got:", got)
	}
}

func BenchmarkUseCount(b *testing.B) {
	str := "Hola Mundo Hola"
	for i := 0; i < b.N; i++ {
		UseCount(str)
	}
}
func BenchmarkCount(b *testing.B) {
	str := "Hola Mundo Hola"
	for i := 0; i < b.N; i++ {
		Count(str)
	}
}

func ExampleCount() {
	str := "Hola Mundo Hola"
	fmt.Println(Count(str))
	//Output: 3
}

func mapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v2 != v {
			return false
		}
	}
	return true
}
