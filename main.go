package main

import (
	"forum/backend/database"
	"forum/backend/server"
)

func main() {
	database.InitMainDB()
	server.LoadServer()
}
