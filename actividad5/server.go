package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Server struct {
	subjectGrades map[string]map[string]float64
	studentGrades map[string]map[string]float64
}

type StudentGrade struct {
	StudentName string
	Subject string
	Grade float64
}

func newServer() *Server {
	var server Server
	server.subjectGrades = make(map[string]map[string]float64)
	server.studentGrades = make(map[string]map[string]float64)
	return &server
}

func server() {
	rpc.Register(newServer())
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go rpc.ServeConn(c)
	}
}

func (s *Server) SetStudentGrade(sg StudentGrade, reply *bool) error {
	log.Printf("Attempting to set grade %f on subject %s for student %s", sg.Grade, sg.Subject, sg.StudentName)
	studentData, exists := s.studentGrades[sg.StudentName]
	if !exists {
		log.Println("Student data does not exists yet, creating new entry")
		studentData = make(map[string]float64)
		studentData[sg.Subject] = sg.Grade
		s.studentGrades[sg.StudentName] = studentData
	} else {
		log.Println("Student data exists already, checking if subject exists")
		_, exits := s.studentGrades[sg.StudentName][sg.Subject]
		if !exits {
			log.Printf("Subject %s does not exists for student %s yet, creating new entry", sg.Subject, sg.StudentName)
			s.studentGrades[sg.StudentName][sg.Subject] = sg.Grade
		} else {
			log.Println("Grade already exists for student", sg.StudentName, "and subject", sg.Subject)
			return errors.New("\nYa existe calificacion para este alumno en esta materia, por favor, verifique\n")
		}
	}
	subjectData, exists := s.subjectGrades[sg.Subject]
	if !exists {
		subjectData = make(map[string]float64)
		subjectData[sg.StudentName] = sg.Grade
		s.subjectGrades[sg.Subject] = subjectData
	} else {
		s.subjectGrades[sg.Subject][sg.StudentName] = sg.Grade
	}
	log.Println("Grade set successfully")
	*reply = true
	return nil
}

func (s *Server) GetStudentGPA(studentName string, reply *float64) error {
	*reply = s.GetGPAByName(studentName)
	return nil
}

func (s *Server) GetStudentsGPA(a string, reply *float64) error {
	var grades []float64
	for student, _ := range s.studentGrades {
		grades = append(grades, s.GetGPAByName(student))
	}
	*reply = calculateGPA(grades)
	return nil
}

func (s *Server) GetSubjectGPA(subject string, reply *float64) error {
	subjectGrades := s.subjectGrades[subject]
	var grades []float64
	for _, grade := range subjectGrades {
		grades = append(grades, grade)
	}
	*reply = calculateGPA(grades)
	return nil
}

func (s *Server) GetGPAByName(studentName string) float64 {
	studentGrades := s.studentGrades[studentName]
	var grades []float64
	for _, grade := range studentGrades {
		grades = append(grades, grade)
	}
	return calculateGPA(grades)	
}

func (s *Server) print() {
	fmt.Println(s.studentGrades)
	fmt.Println(s.subjectGrades)
}

func calculateGPA(grades []float64) float64 {
	var gpa float64
	for _, g := range grades {
		gpa += g
	}
	return gpa / float64(len(grades))
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}