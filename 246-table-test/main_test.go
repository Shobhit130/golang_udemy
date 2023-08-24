package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{4, 2, 4}, 10},
		{[]int{-1, 0, 2}, 1},
		{[]int{2, 3, 1, 2}, 8},
	}

	for _, v := range tests {
		ans := mySum(v.data...)
		if ans != v.answer {
			t.Error("Expected", v.answer, "Got", ans)
		}
	}
}
