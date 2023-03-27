package api

import (
	"net/http"
)

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("To be implemented"))

}

func (s *server) signin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("To be implemented"))
}
