package main

import (
	"testing"
)

func TestSumOfElementsOfSlice(t *testing.T) {
	s := []int {5, 6, 1, 4}

	sum := SumOfElementsOfSlice(s)
	
	if (sum != 16) {
		t.Fatalf("%v is not 16", sum)
	}

	s2 := []int {4, 3, 4, 1, 0}

	sum = SumOfElementsOfSlice(s2)
	if (sum != 12) {
		t.Fatalf("%v is no 12", sum)
	}
}