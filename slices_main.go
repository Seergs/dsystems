package main

import (
	"fmt"
)

func main() {
	fmt.Print("\n\nSuma de elementos\n")
	n := GetNumberOfElementsForSlice()
	s := GetElementsForSlice(n)

	fmt.Printf("Suma = %v\n\n", SumOfElementsOfSlice(s))
}