package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"algogrit.com/emp_server/entities"
)

func (h EmployeeHandler) indexV1(w http.ResponseWriter, req *http.Request) {
	emps, err := h.svcV1.Index(req.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emps)
}

func (h EmployeeHandler) createV1(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	savedEmp, err := h.svcV1.Create(req.Context(), newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(savedEmp)
}
