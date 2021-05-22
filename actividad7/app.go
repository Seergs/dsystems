package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type AdminGrades struct {
	studentGrades map[string]map[string]float64
}
type StudentGrade struct {
	StudentName string
	Subject string
	Grade float64
}

func newAdmin() *AdminGrades {
	return &AdminGrades{
		studentGrades: make(map[string]map[string]float64),
	}
}

func (a *AdminGrades) grades(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		grades, _ := a.getAllGrades()
		res.Write(grades)
	case "POST":
		var studentGrade StudentGrade
		err := json.NewDecoder(req.Body).Decode(&studentGrade)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		if !a.addNew(studentGrade) {
			http.Error(res, "Ya existe calificación para este alumno en esta materia, intenta con PUT", http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func (a *AdminGrades) grade(res http.ResponseWriter, req *http.Request) {
	studentId := strings.TrimPrefix(req.URL.Path, "/grades/")
	switch req.Method {
	case "GET":
		grades, _ := a.getStudentGrades(studentId)
		res.Write(grades)
	case "PUT":
		var studentGrade StudentGrade
		err := json.NewDecoder(req.Body).Decode(&studentGrade)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		studentGrade.StudentName = studentId
		if !a.updateStudentGrade(studentGrade) {
			http.Error(res, "User does not exists", http.StatusNotFound)
		} else {
			res.WriteHeader(http.StatusOK)
		}
	case "DELETE":
		if !a.deleteStudent(studentId) {
			http.Error(res, "User does not exists", http.StatusNotFound)
		} else {
			res.WriteHeader(http.StatusOK)
		}
	}
}

func (a *AdminGrades) getAllGrades() ([]byte, error) {
	log.Println("Getting all students grades")
	gradesJson, err := json.MarshalIndent(a.studentGrades, "", "	")
	if err != nil {
		return gradesJson, nil
	}
	return gradesJson, err
}

func (a *AdminGrades) addNew(studentGrade StudentGrade) bool {
	log.Printf("Attempting to set grade %f on subject %s for student %s", studentGrade.Grade, studentGrade.Subject, studentGrade.StudentName)
	studentData, exists := a.studentGrades[studentGrade.StudentName]
	if !exists {
		log.Println("Student data does not exists yet, creating new entry")
		studentData = make(map[string]float64)
		studentData[studentGrade.Subject] = studentGrade.Grade
		a.studentGrades[studentGrade.StudentName] = studentData
	} else {
		log.Println("Student data exists already, checking if subject exists")
		_, exits := a.studentGrades[studentGrade.StudentName][studentGrade.Subject]
		if !exits {
			log.Printf("Subject %s does not exists for student %s yet, creating new entry", studentGrade.Subject, studentGrade.StudentName)
			a.studentGrades[studentGrade.StudentName][studentGrade.Subject] = studentGrade.Grade
		} else {
			log.Println("Grade already exists for student", studentGrade.StudentName, "and subject", studentGrade.Subject)
			return false
		}
	}
	log.Println("Grade set successfully")
	return true
}

func (a *AdminGrades) getStudentGrades(studentId string) ([]byte, error) {
	log.Println("Getting student grades")
	gradesJson, err := json.MarshalIndent(a.studentGrades[studentId], "", "		")
	if err != nil {
		return gradesJson, nil
	}
	return gradesJson, err
}

func (a *AdminGrades) updateStudentGrade(newStudentGrade StudentGrade) bool {
	log.Println("Updating student grade")
	_, exists := a.studentGrades[newStudentGrade.StudentName]
	if !exists {
		log.Printf("User %s does not exists", newStudentGrade.StudentName)
		return false
	}
	a.studentGrades[newStudentGrade.StudentName][newStudentGrade.Subject] = newStudentGrade.Grade
	return true
}

func (a *AdminGrades) deleteStudent(studentId string) bool {
	_, exists := a.studentGrades[studentId]
	if !exists {
		log.Printf("User %s does not exists", studentId)
		return false
	}
	log.Println("Deleting student")
	delete(a.studentGrades, studentId)
	return true
}

func main() {
	admin := newAdmin()
	http.HandleFunc("/grades", admin.grades)
	http.HandleFunc("/grades/", admin.grade)
	log.Println("Server running on http://localhost:5000")
	http.ListenAndServe(":5000", logRequest(http.DefaultServeMux))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
