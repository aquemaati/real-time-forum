package main

import "github.com/aquemaati/real-time-forum/server"

// main function
func main() {

	// handling all routes
	handler := server.Routing()

	// launching server
	serv := server.InitServer("8080", handler)
	server.RunServer(serv)
}
