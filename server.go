package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	// "reflect"
)

type Employee struct { // Struct Tags
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "speciality": "%s", "project": %d}`, e.Name, e.Department, e.ProjectID)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Senthil", "Cloud", 10002},
	{3, "Sonali", "SRE", 10010},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)

	// encoder := json.NewEncoder(w)
	// encoder.Encode(employees)
	// fmt.Fprintln(w, employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	// req.Body
	var newEmp Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1

	employees = append(employees, newEmp)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		next.ServeHTTP(w, req)

		log.Printf("%s %s took %s\n", req.Method, req.URL, time.Since(begin))
	}

	return http.HandlerFunc(fn)
}

func main() {
	// http.DefaultServeMux
	// r := http.NewServeMux()
	r := mux.NewRouter()

	// http.HandleFunc
	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		// w.Write([]byte(msg))
		fmt.Fprintln(w, msg)
	})

	// r.HandleFunc("/employees", EmployeesHandler)
	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	// http.ListenAndServe("localhost:3000", nil)
	// http.ListenAndServe("localhost:3000", r)
	http.ListenAndServe("localhost:3000", LoggingMiddleware(r))
}
