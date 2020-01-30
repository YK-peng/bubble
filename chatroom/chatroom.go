package chatroom

import (
	"chatroom/conf"
	"chatroom/dao"
)

// Auth server
type Server struct {
	c *conf.Config

	dao *dao.Dao
	//rpc client to UnifyAuth //todo  带token到统一登录服务去先验证token有效性

	// online
	totalIPs   int64
	totalConns int64
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
