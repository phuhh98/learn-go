package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type sample struct {
		data   []int
		answer float64
	}

	samples := []sample{
		sample{data: []int{1, 4, 6, 8, 100}, answer: 6},
		sample{data: []int{0, 8, 10, 1000}, answer: 9},
		sample{data: []int{9000, 4, 10, 8, 6, 12}, answer: 9},
		sample{data: []int{123, 744, 140, 200}, answer: 170},
	}
	for _, sample := range samples {
		cen := CenteredAvg(sample.data)

		if cen != sample.answer {
			t.Errorf("Got %v expected %v", cen, sample.answer)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{9000, 4, 10, 8, 6, 12})
	}
}

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	xi := [][]int{a, b, c, d}

	for _, v := range xi {
		fmt.Println(CenteredAvg(v))
	}
	//Output:
	//6
	//9
	//9
	//170
}
