package main

import "fmt"

func main() {
	fmt.Println("Server init")
	server := NewServer(":8080")
	//Se define el endPoint principal
	server.Handle("/", HandlerRoot)
	server.Handle("/api", server.AddMiddleware(HandlerHome, CheckAuth()))
	server.Listen()
}
