package entities

type Employee struct { // Struct Tags
	ID         int    `json:"-"`
	Name       string `json:"name" validate:"required"`
	Department string `json:"speciality" validate:"required"`
	ProjectID  int    `json:"project" validate:"required"`
}
