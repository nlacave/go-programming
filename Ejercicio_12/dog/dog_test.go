package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	v := Years(8)
	if v != 56 {
		t.Error("Got:", v, "Expected:", 56)
	}
}

func ExampleYears() {
	v := Years(8)
	fmt.Println(v)
	//Output: 56
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(Years(8))
	}
}

func TestYearsTwo(t *testing.T) {
	v := YearsTwo(8)
	if v != 56 {
		t.Error("Got:", v, "Expected:", 56)
	}
}

func ExampleYearsTwo() {
	v := YearsTwo(8)
	fmt.Println(v)
	//Output: 56
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(8)
	}
}
