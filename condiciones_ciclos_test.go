package main

import (
	"testing"
)

func TestIsValidDay(t *testing.T) {
	isValid := IsValidDay(0)
	if isValid {
		t.Fatal("0 should not be a valid day")
	}
	isValid = IsValidDay(32)
	if isValid {
		t.Fatal("32 should not be a valid day")
	}
	isValid = IsValidDay(17)
		if !isValid {
			t.Fatal("17 should be a valid day")
		}
}

func TestGetZodiacSign(t *testing.T) {
	sign := GetZodiacSign(1, 19)
	if sign != "Capricornio" {
		t.Fatal(sign + " is not Capricornio")
	}

	sign = GetZodiacSign(1, 20)
	if sign != "Acuario" {
		t.Fatal(sign + " is not Acuario")
	}

	sign = GetZodiacSign(2, 18)
	if sign != "Acuario" {
		t.Fatal(sign + " is not Acuario")
	}

	sign = GetZodiacSign(2, 19)
	if sign != "Piscis" {
		t.Fatal(sign + " is not Piscis")
	}

	sign = GetZodiacSign(3, 20)
	if sign != "Piscis" {
		t.Fatal(sign + " is not Piscis")
	}

	sign = GetZodiacSign(3, 21)
	if sign != "Aries" {
		t.Fatal(sign + " is not Aries")
	}

	sign = GetZodiacSign(4, 19)
	if sign != "Aries" {
		t.Fatal(sign + " is not Aries")
	}

	sign = GetZodiacSign(4, 20)
	if sign != "Tauro" {
		t.Fatal(sign + " is not Tauro")
	}

	sign = GetZodiacSign(5, 20)
	if sign != "Tauro" {
		t.Fatal(sign + " is not Tauro")
	}

	sign = GetZodiacSign(5, 21)
	if sign != "Geminis" {
		t.Fatal(sign + " is not Geminis")
	}

	sign = GetZodiacSign(6, 20)
	if sign != "Geminis" {
		t.Fatal(sign + " is not Geminis")
	}

	sign = GetZodiacSign(6, 21)
	if sign != "Cancer" {
		t.Fatal(sign + " is not Cancer")
	}

	sign = GetZodiacSign(7, 22)
	if sign != "Cancer" {
		t.Fatal(sign + " is not Cancer")
	}

	sign = GetZodiacSign(7, 23)
	if sign != "Leo" {
		t.Fatal(sign + " is not Leo")
	}

	sign = GetZodiacSign(8, 22)
	if sign != "Leo" {
		t.Fatal(sign + " is not Leo")
	}

	sign = GetZodiacSign(8, 23) 
	if sign != "Virgo" {
		t.Fatal(sign + " is not Virgo")
	}

	sign = GetZodiacSign(9, 22) 
	if  sign != "Virgo" {
		t.Fatal(sign + " is not Virgo")
	}

	sign = GetZodiacSign(9, 23)
	if sign != "Libra" {
		t.Fatal(sign + " is not Libra")
	}

	sign = GetZodiacSign(10, 22) 
	if sign != "Libra" {
		t.Fatal(sign + " is not Libra")
	}

	sign = GetZodiacSign(10, 23)
	if sign != "Escorpio" {
		t.Fatal(sign + " is not Escorpio")
	}

	sign = GetZodiacSign(11, 21)
	if sign != "Escorpio" {
		t.Fatal(sign + " is not Escorpio")
	}

	sign = GetZodiacSign(11, 22)
	if sign != "Sagitario" {
		t.Fatal(sign + " is not Sagitario")
	}

	sign = GetZodiacSign(12, 21) 
	if sign != "Sagitario" {
		t.Fatal(sign + " is not Sagitario")
	}
	
	sign = GetZodiacSign(12, 22)
	if sign != "Capricornio" {
		t.Fatal(sign + " is not Capricornio")
	}
}

func TestFactorial(t *testing.T) {
	factorial := Factorial(5)
	if factorial != 120 {
		t.Fatalf("%c is not 120", factorial)
	}

	factorial = Factorial(2)
	if factorial != 2 {
		t.Fatalf("%c is not 2", factorial)
	}

	factorial = Factorial(0)
	if factorial != 1 {
		t.Fatalf("%c is not 1", factorial)
	}
}

func TestGetE(t *testing.T) {
	expected := 2.718
	e := GetE()
	if !IsAlmostEqual(expected, e, 1e-4) {
		t.Fatalf("%f is not %f", e, expected)
	}
}