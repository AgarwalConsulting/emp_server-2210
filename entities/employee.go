package entities

type Employee struct { // Struct Tags
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}
