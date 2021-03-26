package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


type Server struct {
	router 	*mux.Router
	// db 		database
}

func buildServer(s *Server){
	s.router = mux.NewRouter()
	// s.db = db
	s.routes()
	s.router.Use(s.middlewareLogging)
}

func New(){
	fmt.Println("Started Server")
	s := &Server{}
	buildServer(s)
		srv := http.Server{
		Handler: 		s.router,
		Addr:			"0.0.0.0:80",
		IdleTimeout: 	120 * time.Second,
		ReadTimeout: 	120 * time.Second,
		WriteTimeout: 	120 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	return
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int){
	_, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		
		if data != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		}
		w.WriteHeader(status)
	}
}

func (s *Server) respondError(w http.ResponseWriter, r *http.Request, err string, status int){
	log.Printf("HTTP %d - %s", status, err)
	http.Error(w, err, http.StatusInternalServerError)
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}