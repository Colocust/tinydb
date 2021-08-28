package tinydb

import (
	"net"
	"tinydb/db"
	"tinydb/object"
)

type Client struct {
	Id   uint            // 唯一自增ID
	Conn *net.Conn       // 连接对象
	DB   *db.DB          // 当前选择的数据库
	Argc int             // 当前命令的参数个数
	Argv []object.Object // 当前命令的参数
}

func NewClient() *Client {
	return new(Client)
}
