package main

type userOption int

const (
	cuadrado userOption = iota
	triangulo
	circulo
	grados
)
/*
func main() {
	option := getOptionFromUserInput()

	printSelectedOption(option)

	switch option {
		case int(cuadrado): 
			l := GetDataFromUserForAreaCuadrado()
			result := AreaCuadrado(l)
			fmt.Printf("Area = %f\n", result)
			break	
		case int(triangulo): 
			b, h := GetDataFromUserForAreaTriangulo()
			result := AreaTriangulo(b,h)
			fmt.Printf("Area = %f\n", result)
			break
		case int(circulo):
			r := GetDataFromUserForAreaCirculo()
			result := AreaCirculo(r)
			fmt.Printf("Area = %f\n", result)
			break
		case int(grados):
			f := GetDataFromUserForFahrenheitToCelcius()
			result := FahrenheitToCelcius(f)
			fmt.Printf("Celcius = %f\n", result)
			break
		default:
			fmt.Print("Unsupported operation")
			break
	}
}


func getOptionFromUserInput() int {
	printOptions()
	var option int
	fmt.Print("Opcion: ")


	fmt.Scan(&option)

	fmt.Print("\n")

	return option - 1
}

func printOptions()  {
	fmt.Print("\n\n")
	fmt.Println("1. Area cuadrado")
	fmt.Println("2. Area triangulo")
	fmt.Println("3. Area circulo")
	fmt.Println("4. Fahrenheit a Celcius")
}

func printSelectedOption(option int) {
	switch option {
		case int(cuadrado):
			fmt.Print("AREA DEL CUADRADO\n")
			break
		case int(triangulo):
			fmt.Print("AREA DEL TRIANGULO\n")
			break
		case int(circulo):
			fmt.Print("AREA DEL CIRCULO\n")
			break
		case int(grados):
			fmt.Print("FAHRENHEIT A CELCIUS\n")
			break
		default:
			fmt.Print("Unsupported option")
			break
	}
}
*/