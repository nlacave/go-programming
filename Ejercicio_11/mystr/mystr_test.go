package mystr

import (
	"fmt"
	"testing"
)

func TestConcatenar(t *testing.T) {
	xs := []string{"Aprendiendo", "a", "programar", "en", "Go"}
	f := Concatenar(xs)
	if f != "Aprendiendo a programar en Go " {
		t.Error("Got", f, "want", "Aprendiendo a programar en Go ")
	}
}

func ExampleConcatenar() {
	xs := []string{"Aprendiendo", "a", "programar", "en", "Go"}
	f := Concatenar(xs)
	fmt.Println(f)
	//Output:
	//Aprendiendo a programar en Go
}

func BenchmarkConcatenar(b *testing.B) {
	xs := []string{"Aprendiendo", "a", "programar", "en", "Go"}
	for i := 0; i < b.N; i++ {
		Concatenar(xs)
	}
}
