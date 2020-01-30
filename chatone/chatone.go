package chatone

import (
	"chatone/conf"
	"chatone/dao"
)

// chatone server
type Server struct {
	c *conf.Config

	dao *dao.Dao
}

func NewServer(c *conf.Config) (s *Server) {
	s = &Server{
		c:   c,
		dao: dao.New(c),
	}

	return s
}

func (s *Server) Close() {
	s.dao.Close()
}
