package main

import (
	"fmt"
)

// GetNumberOfElementsForSlices returns
func GetNumberOfElementsForSlice() int {
	var n int
	fmt.Print("Numero de elementos: ")
	fmt.Scan(&n)

	return n
}

// GetElementsForSlices returns
func GetElementsForSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var num int 
		fmt.Printf("Numero %v: ", i + 1)
		fmt.Scan(&num)
		s = append(s, num)
	}

	return s
}

// SumOfElementsOfSlices returns
func SumOfElementsOfSlice(s []int) int {
	sum := 0
	for _,v := range s {
		sum += v
	}
	
	return sum
}