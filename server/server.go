package server

import (
	"github.com/Colocust/strcture"
	"net"
	"tinydb/command"
	"tinydb/config"
	"tinydb/db"
	"tinydb/object"
)

type (
	Server struct {
		Pid           int            // 进程号
		Cfg           *config.Config // 配置
		DB            *db.DB         // 全局数据库
		Clients       *strcture.List // 当前连接的客户端
		CurrentClient *Client        // 当前处理的客户端
	}

	Client struct {
		Id      uint             // 唯一自增ID
		Conn    *net.Conn        // 连接对象
		DB      *db.DB           // 当前选择的数据库
		Command *command.Command // 当前命令
		Argc    int              // 当前命令的参数个数
		Argv    object.Object    // 当前命令的参数
	}
)

var Serv = new(Server)

func NewClient() *Client {
	return new(Client)
}
