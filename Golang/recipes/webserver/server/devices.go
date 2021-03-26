package server

import (
	"net/http"
)

func(s *Server) handleDevicesCreate() http.HandlerFunc {
	
	type device struct {
		SerialID        string `json:"serial_id"`
	}

	return func(w http.ResponseWriter, r *http.Request){
		s.respond(w, r, nil, 201)
		// var d device
		// err := s.decode(w, r, &d)
		// if err == nil {
		// 	err = s.db.InsertDevice(d.SerialID); 
		// 	if err != nil {
		// 		s.respond(w, r, nil, 500)
		// 		return 
		// 	}
		// 	s.respond(w, r, nil, 201)
		// 	return
		// }
	}
}

func(s *Server) handleDevicesRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		s.respond(w, r, nil, 200)
	}
}