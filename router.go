package main

import (
	"fmt"
	"net/http"
)

//Router direge un endpoit a su handler correspondiente, es decir
//realiza mapeos de la peticiones
type Router struct {
	rules map[string]http.HandlerFunc
}

//NewRouter devulve una instance re Router con las reglas
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

//ServerHTTP permite responder todos los request
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	//Usa un Writer
	fmt.Fprintf(w, "Hello World!")
}
