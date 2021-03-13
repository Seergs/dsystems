package main

import "math"

//IsAlmostEqual returns wheter two floating point values are close based on an error
func IsAlmostEqual(expected float64, actual float64, _error float64) bool {
	return math.Abs(actual - expected) > _error
}