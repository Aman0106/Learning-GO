package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {

	courses = append(courses, []Course{
		{
			CourseId:    "101",
			CourseName:  "Go Programming",
			CoursePrice: 199,
			Author: &Author{
				Fullname: "Hitesh Chowdhary",
				Website:  "https://hiteshchowdhary.com",
			},
		},
		{
			CourseId:    "102",
			CourseName:  "Data Structures & Algorithms",
			CoursePrice: 149,
			Author: &Author{
				Fullname: "Striver",
				Website:  "https://striver.com",
			},
		},
		{
			CourseId:    "103",
			CourseName:  "Native Android Development",
			CoursePrice: 299,
			Author: &Author{
				Fullname: "Philip Lackner",
				Website:  "https://philip.dev",
			},
		},
	}...)

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course", addOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	fmt.Println("Server Started")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Fake DB
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// controllers

// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving Home")
	w.Write([]byte("<h1>Welcome To Home</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by id")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	noCouseMessage := "Now course found with id: " + params["id"]

	json.NewEncoder(w).Encode(noCouseMessage)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found")
}

func addOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No request body")
		return
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Request body is empty")
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	fmt.Println("course to be add:", course)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}
