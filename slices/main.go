package main

import (
	"fmt"
)

func GetNumberOfElementsForSlice() int {
	var n int
	fmt.Print("Numero de elementos: ")
	fmt.Scan(&n)

	return n
}

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

func SumOfElementsOfSlice(s []int) int {
	sum := 0
	for _,v := range s {
		sum += v
	}
	
	return sum
}

func main() {
	fmt.Print("\n\nSuma de elementos\n")
	n := GetNumberOfElementsForSlice()
	s := GetElementsForSlice(n)

	fmt.Printf("Suma = %v\n\n", SumOfElementsOfSlice(s))
}
