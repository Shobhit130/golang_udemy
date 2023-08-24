package word

import (
	"fmt"
	"mymodule/254-hands-on-exercise-2/quote"
	"testing"
)

func TestCount(t *testing.T) {
	n := Count("one two three")

	if n != 3 {
		t.Error("got", n, "wanted", 3)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")

	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}
