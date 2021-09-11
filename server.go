package tinydb

import (
	"github.com/Colocust/strcture"
	"os"
	"tinydb/config"
	"tinydb/db"
)

type (
	Server struct {
		pid           int            // 进程号
		cfg           *config.Config // 配置
		db            *db.DB         // 全局数据库
		clients       *strcture.List // 当前连接的客户端
		currentClient *Client        // 当前处理的客户端
	}
)

var Serv *Server

func init() {
	Serv = &Server{
		pid:     os.Getpid(),
		clients: strcture.NewList(),
		db:      db.NewDB(),
		cfg:     config.Load(),
	}
}

func SetCurrentClient(c *Client) {
	Serv.currentClient = c
}

func GetPid() int {
	return Serv.pid
}

func GetCfg() *config.Config {
	return Serv.cfg
}

func GetDB() *db.DB {
	return Serv.db
}
