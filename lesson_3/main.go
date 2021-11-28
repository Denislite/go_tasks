package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Grade string `json:"grade"`
}

var Students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getStudents")
	json.NewEncoder(w).Encode(Students)
}

func getSingleStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, student := range Students {
		if student.Id == key {
			json.NewEncoder(w).Encode(student)
		}
	}
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliction/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)
	file, _ := json.MarshalIndent(Students, "", " ")
	_ = ioutil.WriteFile("mocks/students.json", file, 0644)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, student := range Students {
		if student.Id == id {
			Students = append(Students[:index], Students[index+1:]...)
		}
	}

	file, _ := json.MarshalIndent(Students, "", " ")
	_ = ioutil.WriteFile("mocks/students.json", file, 0644)
}

func getStudentByGrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	for _, student := range Students {
		if student.Grade == key {
			json.NewEncoder(w).Encode(student)
		}
	}
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/students", getStudents)
	r.HandleFunc("/", createStudent).Methods("POST")
	r.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
	r.HandleFunc("/student/{id}", getSingleStudent)
	r.HandleFunc("/student/grade/{key}", getStudentByGrade)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	JsonFile, err := os.Open("mocks/students.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(JsonFile)

	json.Unmarshal(byteValue, &Students)

	JsonFile.Close()
	handleRequests()
}
