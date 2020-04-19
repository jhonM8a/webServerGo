package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//CheckAuth verifica que se este autenticado
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Check auth")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

//Logging crea un Log al llamar un endpoint
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			//ejecuta la función al final
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}() //llama a la función anonima
			f(w, r)
		}
	}
}
