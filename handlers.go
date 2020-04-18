package main

import (
	"fmt"
	"net/http"
)

//HandlerRoot responde a la ruta principal
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola con handlres")
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello home")
}
