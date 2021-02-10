package gorilla

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Freeit struct {
	ID         string `json:"id"`
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
	Position   string `json:"Position"`
}
type FreeitV2 struct {
	ID         string `json:"id"`
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
}

var Students []Freeit
var StudentsV2 []FreeitV2

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Students)
}

func GetStudentV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Freeit{})
}
func GetStudentV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range StudentsV2 {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&FreeitV2{})
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Freeit
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(100))
	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Students {
		if item.ID == params["id"] {
			Students = append(Students[:index], Students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Students)
}
