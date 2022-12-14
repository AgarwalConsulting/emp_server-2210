package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

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
	decoder := json.NewDecoder(req.Body)

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newEmp)

	validate := validator.New()

	errs := validate.Struct(newEmp)

	if err != nil || errs != nil {
		w.WriteHeader(http.StatusBadRequest)

		if err != nil {
			fmt.Fprintln(w, err)
		} else {
			fmt.Fprintln(w, errs)
		}
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
