package tinydb

import (
	"github.com/Colocust/strcture"
	"tinydb/config"
	"tinydb/db"
)

type (
	Server struct {
		Pid           int            // 进程号
		Cfg           *config.Config // 配置
		DB            *db.DB         // 全局数据库
		Clients       *strcture.List // 当前连接的客户端
		CurrentClient *Client        // 当前处理的客户端
	}
)

var Serv = new(Server)
