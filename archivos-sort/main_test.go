package main

import (
	"testing"
	"math"
)

func TestSortAsc(t *testing.T) {
	s := []string{"hola", "mundo", "como", "te", "encuentras"}
	expected := []string{"como", "encuentras", "hola", "mundo", "te"}

	SortAsc(s)

	if !isEqualStringSlice(s, expected) {
		t.Fatalf("%s is not %s", s, expected)
	}
}

func TestSortDesc(t *testing.T) {
	s := []string{"hola", "mundo", "como", "te", "encuentras"}
	expected := []string{"te", "mundo", "hola", "encuentras", "como"}

	SortDesc(s)

	if !isEqualStringSlice(s, expected) {
		t.Fatalf("%s is not %s", s, expected)
	}
}


func isAlmostEqual(expected float64, actual float64, _error float64) bool {
	return math.Abs(actual - expected) > _error
}

func isEqualSlice(a, b [] int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i]{
			return false
		}
	}

	return true
}

func isEqualStringSlice(a, b [] string) bool {
	if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}