package main

import (
	"interfaces/database"
	"interfaces/server"
)


func main() {
	db := database.Start()
	s := server.Start(db)
	s.HandleInsert()
}