package main

import (
	"fmt"
)

func Burbuja(s [] int64) {
	n := len(s)
	
	for i:=0; i<(n-1); i++ {
		for j:=0; j<(n-i-1); j++ {
			if s[j+1] < s[j] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
}

func Fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n - 2)
}

func Greater(args ...int) int {
	greatest := args[0]

	for _,value := range args {
		if value > greatest {
			greatest = value
		}
	}

	return greatest
}

func GenerarImpar() func() uint  {
	i := uint(1)

	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func Intercambia(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 5, 3
	fmt.Print("Antes del cambio\n")
	fmt.Printf("a = %d, b = %d", a, b)


	Intercambia(&a, &b)

	fmt.Print("\nDespues del cambio\n")
	fmt.Printf("a = %d, b = %d", a, b)
}
