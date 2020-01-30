package dao

import (
	"context"

	"auth/conf"
)

// Dao dao.
type Dao struct {
	c *conf.Config
	//todo
}

// New new a dao and return.
func New(c *conf.Config) *Dao {
	d := &Dao{
		c: c,
	}
	return d
}

func (d *Dao) Close() error {
	return nil
}

// Ping dao ping.
func (d *Dao) Ping(c context.Context) error {
	return nil
}
