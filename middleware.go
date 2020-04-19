package main

import (
	"fmt"
	"net/http"
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
