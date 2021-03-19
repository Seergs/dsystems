package main

//Burbuja orders
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

//Fibonacci returns
func Fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n - 2)
}

//Greater returns
func Greater(args ...int) int {
	greatest := args[0]

	for _,value := range args {
		if value > greatest {
			greatest = value
		}
	}

	return greatest
}

//GenerarImpar returns
func GenerarImpar() func() uint  {
	i := uint(1)

	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

//Intercambia swaps
func Intercambia(a, b *int) {
	*a, *b = *b, *a
}