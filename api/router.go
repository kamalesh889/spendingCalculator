package api

import "github.com/gorilla/mux"

func Router(s *server) *mux.Router {

	r := s.router

	r.HandleFunc("/login", s.login).Methods("POST")
	r.HandleFunc("/signin", s.signin).Methods("POST")

	return r

}
