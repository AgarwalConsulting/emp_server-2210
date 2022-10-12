package main

import (
	// "flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	empHTTP "algogrit.com/emp_server/employees/http"
	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/employees/service"
)

func GetEnvOrDefault(key, dflt string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return dflt
	}

	return val
}

func main() {
	port := GetEnvOrDefault("PORT", "3000")

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	var empRepo = repository.NewInMem()
	var empSvc = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvc)

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port:", port, "...")
	log.Fatalln(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r)))
}
