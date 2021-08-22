package server

import (
	"github.com/Colocust/strcture"
	"os"
	"tinydb/config"
	"tinydb/db"
)

type (
	Server struct {
		pid     int            // 进程号
		cfg     *config.Config // 配置
		db      *db.DB         // 全局数据库
		clients *strcture.List // 当前连接的客户端
	}
)

var Serv *Server

func InitServer() {
	Serv = &Server{
		pid:     os.Getpid(),
		db:      db.NewDB(),
		clients: strcture.NewList(),
	}
}

func (serv *Server) GetPid() int {
	return serv.pid
}

func (serv *Server) GetCfg() *config.Config {
	return serv.cfg
}

func (serv *Server) SetCfg(cfg *config.Config) {
	serv.cfg = cfg
}

func (serv *Server) GetClients() *strcture.List {
	return serv.clients
}
