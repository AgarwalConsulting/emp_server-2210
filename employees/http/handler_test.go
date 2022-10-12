package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/emp_server/employees/http"
	"algogrit.com/emp_server/employees/service"
	"algogrit.com/emp_server/entities"
)

func TestIndexV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockSvc.EXPECT().Index(gomock.Any()).Return(expectedEmps, nil)

	respRec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/employees", nil)

	// sut.IndexV1(respRec, req)
	// sut.Router.ServeHTTP(respRec, req)
	sut.ServeHTTP(respRec, req)

	res := respRec.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var employees []entities.Employee
	json.NewDecoder(res.Body).Decode(&employees)

	assert.NotNil(t, employees)
	assert.Equal(t, 1, len(employees))
	assert.Equal(t, "Gaurav", employees[0].Name)
	// assert.Equal(t, expectedEmps, employees)
}

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 1001}

	mockSvc.EXPECT().Create(gomock.Any(), expectedEmp).Return(&expectedEmp, nil)

	respRec := httptest.NewRecorder()
	reqBody := `{"name": "Gaurav", "speciality": "LnD", "project": 1001}`
	req := httptest.NewRequest("POST", "/v1/employees", strings.NewReader(reqBody))

	sut.ServeHTTP(respRec, req)

	res := respRec.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var createdEmp entities.Employee
	json.NewDecoder(res.Body).Decode(&createdEmp)

	assert.Equal(t, expectedEmp, createdEmp)
}

func FuzzCreateV1(f *testing.F) {
	f.Add(`{"name": "Gaurav", "speciality": "LnD", "project: 1001}`)

	f.Fuzz(func(t *testing.T, reqBody string) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := service.NewMockEmployeeService(ctrl)

		sut := empHTTP.NewHandler(mockSvc)

		// expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 1001}

		// // mockSvc.EXPECT().Create(gomock.Any(), expectedEmp).Return(&expectedEmp, nil)

		respRec := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/v1/employees", strings.NewReader(reqBody))

		sut.ServeHTTP(respRec, req)

		res := respRec.Result()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
