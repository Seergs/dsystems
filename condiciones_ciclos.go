package main

import (
	"fmt"
)

type month int

const (
	jan month = iota + 1 
	feb
	march
	april
	may
	jun
	jul
	ag
	sep
	oct
	nov
	dic
	end
)

//GetZodiacSign returns
func GetZodiacSign(month int, day int) string {
	switch month {
	case int(jan):
		if day < 20 {
			return "Capricornio"
		}
		return "Acuario"
	case int(feb):
		if day < 19 {
			return "Acuario"
		}
		return "Piscis"
	case int(march):
		if day < 21 {
			return "Piscis"
		}
		return "Aries"
	case int(april):
		if day < 20 {
			return "Aries"
		}
		return "Tauro"
	case int(may):
		if day < 21 {
			return "Tauro"
		}
		return "Geminis"
	case int(jun):
		if day < 21 {
			return "Geminis"
		}
		return "Cancer"
	case int(jul):
		if day < 23 {
			return "Cancer"
		}
		return "Leo"
	case int(ag):
		if day < 23 {
			return "Leo"
		}
		return "Virgo"
	case int(sep):
		if day < 23 {
			return "Virgo"
		}
		return "Libra"
	case int(oct):
		if day < 23 {
			return "Libra"
		}
		return "Escorpio"
	case int(nov):
		if day < 22 {
			return "Escorpio"
		}
		return "Sagitario"
	case int(dic):
		if day < 22 {
			return "Sagitario"
		}
		return "Capricornio"
	}
	return "Algo salio mal"
}

//GetInputDataFromUserForZodiacSign returns
func GetInputDataFromUserForZodiacSign() (int, int) {
	var month int
	var day int

	fmt.Print("Que dia naciste?: ")
	fmt.Scan(&day)
	if !IsValidDay(day) {
		fmt.Print("Invalid day")
	}

	fmt.Print("De que mes?: ")
	fmt.Scan(&month)
	if !IsValidMonth(month) {
		fmt.Print("Invalid month")
	}


	return month, day 
}

func IsValidMonth(month int) bool {
	return month > 0 && month <= 12
}

//IsValidDay returns
func IsValidDay(day int) bool {
	return day > 0 && day <= 31
}


//GetE retursn
func GetE() float64 {
	e := 0.0
	
	for i:=0; i<10; i++ {
		e += 1 / float64(Factorial(i))
	}
	
	return e
}

//Factorial returns
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}