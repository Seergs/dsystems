package main

import (
	"fmt"
	"math"
)

// AreaCuadrado returns the area of a square
func AreaCuadrado(l float64) float64  {
	return l * l
}

//AreaTriangulo returns the area of a triangle
func AreaTriangulo(b float64, h float64) float64 {
	return b * h / 2
}

//AreaCirculo return the are of a circle
func AreaCirculo(r float64) float64 {
	return math.Pi * (r * r)
}

//FahrenheitToCelcius returns the conversion between fahrenheit and celcius
func FahrenheitToCelcius(f float64) float64 {
	return  (f - 32) * 5 / 9 
}

// GetDataFromUserForAreaCuadrado returns
func GetDataFromUserForAreaCuadrado() float64 {
	var l float64
	fmt.Print("Lado: ")
	fmt.Scan(&l)

	return l
}

// GetDataFromUserForAreaTriangulo returns
func GetDataFromUserForAreaTriangulo() (float64, float64) {
	var b float64
	var h float64

	fmt.Print("Base: ")
	fmt.Scan(&b)

	fmt.Print("Altura: ")
	fmt.Scan(&h)

	return b, h
}

//GetDataFromUserForAreaCirculo returns
func GetDataFromUserForAreaCirculo() float64 {
	var r float64

	fmt.Print("Radio: ")
	fmt.Scan(&r)

	return r
}

//GetDataFromUserForFahrenheitToCelcius returns
func GetDataFromUserForFahrenheitToCelcius() float64 {
	var f float64

	fmt.Print("F°: ")
	fmt.Scan(&f)

	return f
}