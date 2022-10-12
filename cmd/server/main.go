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

	"algogrit.com/emp_server/employees/repository"
)

var empRepo = repository.NewInMem()

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	emps, err := empRepo.ListAll(req.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emps)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	savedEmp, err := empRepo.Save(req.Context(), newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(savedEmp)
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
