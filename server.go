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

//Handle asigna una ruta a un hadler creado
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exists := s.router.rules[path]
	if !exists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

//AddMiddleware añade el handler a ejecutar despues de ejecutar uno o varios middleware
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		//Encadena los llamados de los middleware
		f = m(f)
	}
	return f
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
