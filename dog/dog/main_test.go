// Package dog allows us to more fully understand dogs.
package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(7)
	if y != 49 {
		t.Error("Got", y, "Expected", 49)
	}

}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7 * 200)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	//Output:
	//49
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(7)
	if y != 49 {
		t.Error("Got", y, "Expected", 49)
	}

}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7 * 200)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	//Output:
	//49
}
