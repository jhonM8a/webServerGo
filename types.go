package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

//Metadata interface usada para poder hacer decoder de JSON
type Metadata interface{}

//User stuct de HandlerPostRequest
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//ToJSON convierte el struct a un slice de bytes el struct User
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
