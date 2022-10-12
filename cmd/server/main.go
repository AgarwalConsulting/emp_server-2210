package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"algogrit.com/emp_server/entities"
)

var employees = []entities.Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Senthil", "Cloud", 10002},
	{3, "Sonali", "SRE", 10010},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	// req.Body
	var newEmp entities.Employee
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

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	log.Println("Starting server on port: 3000...")
	http.ListenAndServe("localhost:3000", handlers.LoggingHandler(os.Stdout, r))
}
