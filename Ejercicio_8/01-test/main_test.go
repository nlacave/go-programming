package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	v := Sum(1, 2)
	if v != 3 {
		t.Error("Expected", 3, "Got", v)
	}
	fmt.Println("Corrió correctamente el programa.")
}
