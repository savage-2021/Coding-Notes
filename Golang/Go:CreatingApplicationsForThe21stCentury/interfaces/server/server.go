package server

import "fmt"

type Database interface { 
	Insert()
}

type Server struct {
	server string
	db	   Database
}

func Start(db Database) *Server {
	fmt.Println("Created Server")
	s := &Server{"exampleserver", db}
	return s
}

func(s *Server) HandleInsert() {
	s.db.Insert()
}