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