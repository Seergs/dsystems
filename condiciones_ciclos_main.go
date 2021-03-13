package main

import "fmt"

func main() {
	// Primer programa
	fmt.Print("\n\nSigno zodiacal\n")
	month, day := GetInputDataFromUserForZodiacSign()
	zodiacSign := GetZodiacSign(month, day)
	fmt.Print("Eres " + zodiacSign + "\n\n")


	//Segundo programa
	fmt.Print("\n\nCalculando numero de Euler, espere...\n")
	e := GetE()
	fmt.Printf("Euler = %.3f\n", e)
}