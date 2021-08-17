package main

import (
	"fmt"
	"os"
	"sort"
)

func SortAsc(s []string) {
	sort.Strings(s)	
}

func SortDesc(s []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
}

func SaveToFile(strings []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	for _,s := range(strings) {
		file.WriteString(s + "\n")
	}
}

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
