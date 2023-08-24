package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(2, 3)
	want := 6

	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}
