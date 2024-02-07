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

// modle for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper -file

func (c *Course) IsEmpty() bool {

	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("Welcome to Build API")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: "299", Author: &Author{Fullname: "Harsha", Website: "lco.dev"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "MernStack", CoursePrice: "999", Author: &Author{Fullname: "Harsha", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("Get")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by Golang course</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Tyoe", "application/json")

	// grab id from request
	params := mux.Vars(r)

	fmt.Println(params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		} else {

			json.NewEncoder(w).Encode("No course found with given id")
		}

	}

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about data is  -{}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// TODO : check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// Read the request body and decode it into a Course struct
	// var incomingCourse Course

	// if err := json.NewDecoder(r.Body).Decode(&incomingCourse); err != nil {
	// 	// Handle decoding error (e.g., log it or return an error response)
	// 	http.Error(w, "Error decoding JSON", http.StatusBadRequest)
	// 	return
	// }

	for _, excourse := range courses {
		if &excourse.CourseName == &course.CourseName {
			// json.NewEncoder(w).Encode("Same course name exists")
			fmt.Println("Same course name exists")
			break

			// for _, existingCourse := range courses {
			// 	if incomingCourse.CourseName == existingCourse.CourseName {
			// 		json.NewEncoder(w).Encode("Same course name exists")
			// 		return
		} else {
			rand.Seed(time.Now().UnixNano())
			course.CourseId = strconv.Itoa(rand.Int())
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}

	// generate unique id, string
	// append course into courses

	// rand.Seed(time.Now().UnixNano())
	// course.CourseId = strconv.Itoa(rand.Int())
	// courses = append(courses, course)
	// json.NewEncoder(w).Encode(course)
	// return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		} else {
			json.NewEncoder(w).Encode("No course found with given id")
			//fmt.Println("Didn't found course related to id")
		}

	}
	//  TODO send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}
	}

}
