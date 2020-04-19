package main

import "fmt"

func main() {
	fmt.Println("Server init")
	server := NewServer(":8080")
	//Se define el endPoint principal
	server.Handle("GET", "/", HandlerRoot)
	server.Handle("POST", "/api", server.AddMiddleware(HandlerHome, CheckAuth(), Logging()))
	server.Handle("POST", "/create", HandlerPostRequest)
	server.Listen()
}
