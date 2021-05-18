package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type AdminGrades struct {
	subjectGrades map[string]map[string]float64
	studentGrades map[string]map[string]float64
}
type StudentGrade struct {
	studentName string
	subject string
	grade float64
}

func NewAdmin() *AdminGrades {
	return &AdminGrades{
		subjectGrades: make(map[string]map[string]float64),
		studentGrades: make(map[string]map[string]float64),
	}
}

var grades = NewAdmin()

func (g *AdminGrades) setStudentGrade(studentGrade StudentGrade) bool {
	log.Printf("Attempting to set grade %f on subject %s for student %s", studentGrade.grade, studentGrade.subject, studentGrade.studentName)
	studentData, exists := g.studentGrades[studentGrade.studentName]
	if !exists {
		log.Println("Student data does not exists yet, creating new entry")
		studentData = make(map[string]float64)
		studentData[studentGrade.subject] = studentGrade.grade
		g.studentGrades[studentGrade.studentName] = studentData
	} else {
		log.Println("Student data exists already, checking if subject exists")
		_, exits := g.studentGrades[studentGrade.studentName][studentGrade.subject]
		if !exits {
			log.Printf("Subject %s does not exists for student %s yet, creating new entry", studentGrade.subject, studentGrade.studentName)
			g.studentGrades[studentGrade.studentName][studentGrade.subject] = studentGrade.grade
		} else {
			log.Println("Grade already exists for student", studentGrade.studentName, "and subject", studentGrade.subject)
			return false
		}
	}
	log.Println("Adding to second map (subjectGrades)")
	subjectData, exists := g.subjectGrades[studentGrade.subject]
	if !exists {
		log.Println("Grades does not exists, creating new entry")
		subjectData = make(map[string]float64)
		subjectData[studentGrade.studentName] = studentGrade.grade
		g.subjectGrades[studentGrade.subject] = subjectData
	} else {
		log.Println("Subject exists, adding student entry")
		g.subjectGrades[studentGrade.subject][studentGrade.studentName] = studentGrade.grade
	}
	log.Println("Grade set successfully")
	return true
}

func (g *AdminGrades) getStudentGPA(studentName string) float64 {
	log.Println("Getting student GAP")
	gpa := g.getGPAByName(studentName)
	log.Println("GPA:", gpa)
	return gpa
}

func (g *AdminGrades) getStudentsGPA() float64 {
	log.Println("Getting students GPA")
	if len(g.studentGrades) == 0 {
		return 0
	}
	var grades []float64
	for student := range g.studentGrades {
		grades = append(grades, g.getGPAByName(student))
	}
	gpa := calculateGPA(grades)
	log.Println("GPA:", gpa)
	return gpa
}

func (g *AdminGrades) getSubjectGPA(subject string) float64 {
	log.Println("Getting subject GPA")
	subjectGrades := g.subjectGrades[subject]
	var grades []float64
	for _, grade := range subjectGrades {
		grades = append(grades, grade)
	}
	gpa := calculateGPA(grades)
	log.Println("GPA:", gpa)
	return gpa
}

func (g *AdminGrades) getGPAByName(studentName string) float64 {
	studentGrades := g.studentGrades[studentName]
	var grades []float64
	for _, grade := range studentGrades {
		grades = append(grades, grade)
	}
	return calculateGPA(grades)	
}

func (g *AdminGrades) getSubjectNames() []string {
	var subjects []string
	for subject := range g.subjectGrades {
		subjects = append(subjects, subject)
	}
	return subjects
}

func (g *AdminGrades) getStudentNames() []string {
	var students []string
	for student := range g.studentGrades {
		students = append(students, student)
	}
	return students
}

func (g *AdminGrades) studentsToHtml() string {
	students := g.getStudentNames()
	var html string
	for _,student := range students {
		html += "<option value='" + student + "'>" +
			student + "</option>"
	}
	return html
}

func (g *AdminGrades) subjectsToHtml() string {
	subjects := g.getSubjectNames()
	log.Println(subjects)
	var html string
	for _,subject := range subjects {
		html += "<option value='" + subject + "'>" +
			subject + "</option>"
	}
	return html
}

func calculateGPA(grades []float64) float64 {
	var gpa float64
	for _, g := range grades {
		gpa += g
	}
	return gpa / float64(len(grades))
}

func main() {
	http.HandleFunc("/new", newGradeCrontoller)
	http.HandleFunc("/gpa", gpaController)
	log.Println("Server started on port 5000")
	http.ListenAndServe(":5000", nil)
}

func newGradeCrontoller(res http.ResponseWriter, req *http.Request) {
	log.Printf("New request received to /new with method, %s", req.Method )
	switch(req.Method) {
	case "GET":
		fmt.Fprintf(res, render("new.html"))
		break
	case "POST":
		if err := req.ParseForm(); err != nil {
			log.Println(err.Error())
			fmt.Fprintf(res, render("error.html"), "No se pudo procesar la solicitud, intenta de nuevo más tarde")
			return
		}
		grade, err := strconv.ParseFloat(req.FormValue("grade"), 64)
		if (err != nil) {
			fmt.Fprintf(res, render("error.html"), "Algo salió mal al convertir la calificación a número")
			return
		}
		if !grades.setStudentGrade(StudentGrade { studentName: req.FormValue("studentName"), subject: req.FormValue("subject"), grade: grade }) {
			fmt.Fprintf(res, render("error.html"), "Ya existe calificación para este alumno en esta materia, verifica")
			return
		}
		fmt.Fprintf(res, render("added.html"))
		break
	default:
		fmt.Fprintf(res, render("error.html"), "No se pudo procesar la solicitud, intenta de nuevo más tarde")
	}
}

func gpaController(res http.ResponseWriter, req *http.Request) {
	log.Printf("New request received to /gpa with method, %s", req.Method )
	gpaType := req.URL.Query().Get("type")

	switch gpaType {
	case "":
		gpa := grades.getStudentsGPA()
		if gpa == 0 {
			fmt.Fprintf(res, render("error.html"), "Aún no existen calificaciones de alumnos")
		}else{ 
			fmt.Fprintf(res, render("gpa.html"), "general", grades.getStudentsGPA())
		}
		break
	case "student":
		if req.Method == "GET"{
			fmt.Fprintf(res, render("list.html"), grades.studentsToHtml())
		}
			
		break
	case "subject":
		if req.Method == "GET" {
			fmt.Fprintf(res, render("list.html"), grades.subjectsToHtml())
		}
	}
}


func render(path string) string {
	html, _ := ioutil.ReadFile(path)
	return string(html)
}
