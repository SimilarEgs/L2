package app

import (
	"dev11/config"
	"dev11/internal/repository"
	"net/http"
)

// inject event methods (prolly a bad example of clean architecture)
type Server struct {
	Events repository.EventStorage
}

// Run func - runs http.server based on received config
func (srv *Server) Run(cfg *config.Config) error {

	// representation of http.Server
	server := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      srv.NewRouter(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	// run http server
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
