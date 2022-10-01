package app

import (
	"dev11/config"
	"net/http"
)

type Server struct {
	server *http.Server
	cfg config.Config
}

// NewServer func -
func NewServer(config *config.Config) *Server {
	return &Server{cfg: *config}
}

// Run func - runs http.server based on received config
func (s *Server) Run() error {

	// adjusting server settings via provided config
	s.server = &http.Server{
		Addr:         s.cfg.Host + ":" + s.cfg.Port,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	// run http.server
	return s.server.ListenAndServe()
}
