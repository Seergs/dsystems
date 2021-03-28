package main

import (
	"testing"
)

func TestSortAsc(t *testing.T) {
	s := []string{"hola", "mundo", "como", "te", "encuentras"}
	expected := []string{"como", "encuentras", "hola", "mundo", "te"}

	SortAsc(s)

	if !IsEqualStringSlice(s, expected) {
		t.Fatalf("%s is not %s", s, expected)
	}
}

func TestSortDesc(t *testing.T) {
	s := []string{"hola", "mundo", "como", "te", "encuentras"}
	expected := []string{"te", "mundo", "hola", "encuentras", "como"}

	SortDesc(s)

	if !IsEqualStringSlice(s, expected) {
		t.Fatalf("%s is not %s", s, expected)
	}
}