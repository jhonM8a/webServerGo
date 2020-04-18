package main

import "fmt"

func main() {
	fmt.Println("Server init")
	server := NewServer(":8080")
	server.Listen()
}
