package server

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)


func(s *Server) routes(){
	// CRUD Implementation of Devices that can make calls 
	s.router.HandleFunc("/devices", s.handleDevicesCreate()).Methods("POST")
	s.router.HandleFunc("/devices", s.handleDevicesRead()).Methods("GET")
}

func (s *Server) logPrinter(ctx context.Context, msg string){
	type requestIDKey string
	id := ctx.Value(requestIDKey("id"))
	if id != nil {
		log.Printf("[%d] %s", id, msg)
	} else {
		log.Println("No Request ID Attached")
	}
}

func (s *Server) decorate(f http.HandlerFunc) http.HandlerFunc {
	type requestIDKey string
	return func(w http.ResponseWriter, r *http.Request){
		ctx := r.Context()
		key := requestIDKey("id")
		value := rand.Int63()
		ctx = context.WithValue(ctx, key, value)
		f(w, r.WithContext(ctx))
	}
}

func(s *Server) middlewareLogging(f http.Handler) http.Handler {
	type requestIDKey string
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Printf("[54352521] %s", r.RequestURI)
		f.ServeHTTP(w, r)
	})
}