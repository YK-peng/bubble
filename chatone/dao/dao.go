package dao

import (
	"chatone/xframe/xlog"
	"context"
	"database/sql"
	"fmt"
	"time"

	"chatone/conf"
)

// Dao dao.
type Dao struct {
	c  *conf.Config
	db *sql.DB
}

// New new a dao and return.
func New(c *conf.Config) (*Dao, error) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", c.Mysql.DbUser, c.Mysql.DbPwd, "tcp", c.Mysql.Host, c.Mysql.Port, c.Mysql.Database)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		xlog.Error("connection to mysql failed:", err)
		return nil, err
	}
	db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
	db.SetMaxOpenConns(100)                  //设置最大连接数

	return &Dao{
		c:  c,
		db: db,
	}, nil
}

func (d *Dao) Close() error {
	err := d.db.Close()
	if err != nil {
		xlog.Error("mysql db close error:", err)
	}
	return err
}

// Ping dao ping.
func (d *Dao) Ping(c context.Context) error {
	err := d.db.Ping()
	if err != nil {
		xlog.Error("mysql db ping error:", err)
	}
	return err
}
