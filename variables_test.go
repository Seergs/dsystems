package main

import (
	"testing"
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

	if (IsAlmostEqual(expected, result, 1e-4)) {
    	t.Fatalf("Got %v, expected %v", result, expected)
	}

}

func TestFahrenheitToCelcius(t *testing.T) {
	var f float64 = 65

	result := FahrenheitToCelcius(f)

	if (IsAlmostEqual(18.3333, result, 1e-4)) {
		t.Fatalf("%f is not 18.3333", result)
	}
}