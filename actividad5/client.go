package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

type StudentGrade struct {
	StudentName string
	Subject string
	Grade float64
}

func main() {
	c, err := rpc.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println("Algo salio mal")
		return
	}
	for {
		displayMenu()
		option := getIntFromUser()
		
		switch option {
		case 1:
			fmt.Print("Nombre del alumno: ")
			studentName := getStringFromUser()
			fmt.Print("Materia: ")
			subject := getStringFromUser()
			fmt.Print("Calificacion: ")
			if grade, err := strconv.ParseFloat(getStringFromUser(), 64); err == nil {
				var result bool
				err := c.Call("Server.SetStudentGrade", StudentGrade {studentName, subject, grade}, &result)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Print("\n\n\tCalificacion agregada con exito\n\n")
				}
			}
		case 2:
			fmt.Print("Nombre del alumno: ")
			studentName := getStringFromUser()
			var result float64
			err := c.Call("Server.GetStudentGPA", studentName, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\n\tPromedio: %.2f\n\n", result)
			}
		case 3:
			var result float64
			err := c.Call("Server.GetStudentsGPA", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\n\tPromedio: %.2f\n\n", result)
			}
		case 4:
			fmt.Print("Nombre de la materia: ")
			subject := getStringFromUser()
			var result float64
			err := c.Call("Server.GetSubjectGPA", subject, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\n\tPromedio: %.2f\n\n", result)
			}
		
		case 5:
			return
		}
	}
}

func displayMenu() {
	fmt.Println("1. Agregar calificacion de una materia")
	fmt.Println("2. Mostrar el promedio de un alumno")
	fmt.Println("3. Mostrar el promedio general")
	fmt.Println("4. Mostrar el promedio de una materia")
	fmt.Println("5. Salir")
	fmt.Print("Opcion: ")
}

const inputDelimiter = '\n'

func getStringFromUser() string {
	s := ""
	for s == "" {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
	}
	return s
}

func getIntFromUser() int64 {
	var op int64
	fmt.Scan(&op)

	return op
}