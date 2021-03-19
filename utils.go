package main

import "math"

//IsAlmostEqual returns wheter two floating point values are close based on an error
func IsAlmostEqual(expected float64, actual float64, _error float64) bool {
	return math.Abs(actual - expected) > _error
}

//IsEqualSlice retursn
func IsEqualSlice(a, b [] int64) bool {
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