package main

import (
	"net/http"
)

//Router direge un endpoit a su handler correspondiente, es decir
//realiza mapeos de la peticiones
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

//NewRouter devulve una instance re Router con las reglas
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//FindHandler dada una ruta y su tipo de metodo retorna su handler correspondiente y si esta existe
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExists := r.rules[path][method]
	return handler, methodExists, exist
}

//ServerHTTP redirigir las rutas
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExists, exist := r.FindHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)

}
