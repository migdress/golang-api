package model

type Person struct {
	FullName  string `json:"full_name"`
	DNI       string `json:"dni"`
	Birthdate string `json:"birthdate"`
}
