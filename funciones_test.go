package main

import (
	"testing"
)

func TestBurbuja(t *testing.T) {
	s := [] int64 { 4, 7, 8, 1, 4, 3, 6 }
	sorted_slice := [] int64 { 1, 3, 4, 4, 6, 7, 8 }

	Burbuja(s)

	if !IsEqualSlice(s, sorted_slice) {
		t.Fatalf("%v is not %v", s, sorted_slice)
	}
	
}

func TestFibonacci(t *testing.T) {
	n := 10

	fib := Fibonacci(int64(n))

	var expected int64 = 55

	if (fib != expected) {
		t.Fatalf("%o is not %o", fib, expected)
	}
}

func TestGreater(t *testing.T) {
	s := []int {5, 9, 3, 4, 6, 1}

	greatest := Greater(s...)

	if (greatest != 9) {
		t.Fatalf("%o is not %o", greatest, 9)
	}
}

func TestGenerarImpar(t *testing.T) {
	nextImpar := GenerarImpar()

	first := nextImpar()
	if first != 1 {
		t.Fatalf("%o is not %o", first, 1)
	}

	second := nextImpar()
	if second != 3 {
		t.Fatalf("%o is not %o", second, 3)
	}
	third := nextImpar()
	if third != 5 {
		t.Fatalf("%o is not %o", third, 5)
	}
}

func TestIntercambia(t *testing.T) {
	a, b := 5, 6

	Intercambia(&a, &b)

	if (a != 6 || b != 5) {
		t.Fatal("Swap did not work")
	}
}