package http

import (
	"algogrit.com/emp_server/employees/service"
	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	*mux.Router
	svcV1 service.EmployeeService
}

func (h EmployeeHandler) SetupRoutes(r *mux.Router) {
	h.Router = r

	r.HandleFunc("/v1/employees", h.indexV1).Methods("GET")
	r.HandleFunc("/v1/employees", h.createV1).Methods("POST")
}

func NewHandler(svcV1 service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{svcV1: svcV1}

	h.SetupRoutes(mux.NewRouter())

	return h
}
