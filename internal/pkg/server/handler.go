package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *ApiServer) configureRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.helloHandler)
	return r
}

func (s *ApiServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to site")
}
