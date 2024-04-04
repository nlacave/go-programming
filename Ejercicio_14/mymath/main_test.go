package mymath

import (
	"fmt"
	"testing"
)

var sliceDeSlices = [][]int{
	{1, 4, 6, 8, 100},
	{0, 8, 10, 1000},
	{9000, 4, 10, 8, 6, 12},
	{123, 744, 140, 200},
}

func TestCenteredAvg(t *testing.T) {
	for i, v := range sliceDeSlices {
		x := CenteredAvg(v)
		switch i {
		case 0:
			if x != 6 {
				t.Error("Expected:", 6, "Got:", x)
			}
		case 1:
			if x != 9 {
				t.Error("Expected:", 9, "Got:", x)
			}
		case 2:
			if x != 9 {
				t.Error("Expected:", 9, "Got:", x)
			}
		case 3:
			if x != 170 {
				t.Error("Expected:", 170, "Got", x)
			}
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range sliceDeSlices {
			CenteredAvg(v)
		}
	}
}

func ExampleCenteredAvg() {
	for _, v := range sliceDeSlices {
		fmt.Println(CenteredAvg(v))
		//Output:
		//6
		//9
		//9
		//170
	}
}
