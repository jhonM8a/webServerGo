package main

import (
	"fmt"
	"net/http"
)

//Server is one struct for server http
type Server struct {
	port   string
	router *Router
}

//NewServer return new instance of @Server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}

}

//Listen init server in port selected
func (s *Server) Listen() error {
	//Se define el endPoint principal
	http.Handle("/", s.router)
	fmt.Println("Listenting in : ", s.port)
	//Con este el server esta antento alas peticiones
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
