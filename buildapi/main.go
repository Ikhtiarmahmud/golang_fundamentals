package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for a Course -

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice string `json:"courseprice"`
	Author *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"authorname"`
	AuthorWebsite string `json:"authorwebsite"`
}

// fake db
var courses []Course

// middleware, helper - fire

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to Home")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "PHP", CoursePrice: "100", Author: 
	&Author{AuthorName: "Ikhtiar", AuthorWebsite: "ikhtiar.net"}})
	
	courses = append(courses, Course{CourseId: "2", CourseName: "Go", CoursePrice: "300", Author: 
	&Author{AuthorName: "Mitul", AuthorWebsite: "mitul.net"}})

	// router
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/all-courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course", createNewCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateSingleCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteSingleCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// method for course controller

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go API</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) // encode course slice to json
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One courses")
	w.Header().Set("Content-Type", "application/json")

	// grab course id from query params
	params := mux.Vars(r)

	for _, course := range courses {
		if params["id"] == course.CourseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Sorry! No Course Found")
	return
}

func createNewCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One courses")
	w.Header().Set("Content-Type", "application/json")

	// check body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("No Data Found")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No Course Name Found")
	}

	// generate random id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100)) // itoa does convert interger to string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One courses")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...) // remove course first
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Id Found :(")
}

func deleteSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One courses")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("No Course Found :(")
}
