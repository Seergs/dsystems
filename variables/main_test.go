package main

import (
	"testing"
	"math"
)

func TestAreaCuadrado(t *testing.T)	 {
	var  l float64 = 4

	result := AreaCuadrado(l)

	if (result != 16.0) {
		t.Fatalf("%f is not 16", result)
	}
}

func TestAreaTriangulo(t *testing.T) {
	var b float64 = 4
	var h float64 = 6

	result := AreaTriangulo(b,h)

	if (result != 12.0) {
		t.Fatalf("%f is not 12", result)
	}
}

func TestAreaCirculo(t *testing.T) {
	var r float64 = 5.5
	expected := 95.0331777711

	result := AreaCirculo(r)

	if (isAlmostEqual(expected, result, 1e-4)) {
    	t.Fatalf("Got %v, expected %v", result, expected)
	}

}

func TestFahrenheitToCelcius(t *testing.T) {
	var f float64 = 65

	result := FahrenheitToCelcius(f)

	if (isAlmostEqual(18.3333, result, 1e-4)) {
		t.Fatalf("%f is not 18.3333", result)
	}
}

func isAlmostEqual(expected float64, actual float64, _error float64) bool {
	return math.Abs(actual - expected) > _error
}