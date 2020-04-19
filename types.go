package main

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

//Metadata interface usada para poder hacer decoder de JSON
type Metadata interface{}

//User stuct de HandlerPostRequest
type User struct {
	Name  string
	Email string
	Phone string
}
