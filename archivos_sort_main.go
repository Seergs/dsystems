package main

import (
	"fmt"
)

func main() {
	fmt.Print("Cantidad de strings: ")
	n := getIntFromUser()

	s := []string{}

	i := 0

	for int64(i) < n {
		fmt.Print("String ", i + 1 , ": ")
		s = append(s, getStringFromUser())
		i++
	}

	SortAsc(s)
	SaveToFile(s, "ascendente.txt")

	SortDesc(s)
	SaveToFile(s, "descendente.txt")
}

func getStringFromUser() string {
	var line string

	fmt.Scan(&line)

	return line
}

func getIntFromUser() int64 {
	var op int64
	fmt.Scan(&op)

	return op
}
