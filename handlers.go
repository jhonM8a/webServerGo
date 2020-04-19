package main

import (
	"encoding/json"
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

func HandlerPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}
