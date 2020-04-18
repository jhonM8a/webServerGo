package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}

}

func (s *Server) Listen() error {
	fmt.Println("Listenting in : ", s.port)
	//Con este el server esta antento alas peticiones
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
