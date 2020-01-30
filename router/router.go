package router

import (
	"router/conf"
)

// router server
type Server struct {
	c *conf.Config
}

func NewServer(c *conf.Config) (s *Server) {
	s = &Server{
		c: c,
	}

	return s
}

func (s *Server) Close() {
}
