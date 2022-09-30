package app

import (
	"fmt"
	"net/http"
)

func (srv *Server) NewRouter() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/create_event", logEvent(http.HandlerFunc(srv.createEvent)))

	return router
}

func (srv *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
