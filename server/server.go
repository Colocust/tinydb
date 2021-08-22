package server

import (
	"github.com/Colocust/strcture"
	"os"
	"tinydb/config"
	"tinydb/db"
)

type (
	Server struct {
		Pid     int            // 进程号
		Cfg     *config.Config // 配置
		DB      *db.DB         // 全局数据库
		Clients *strcture.List // 当前连接的客户端
	}
)

var serv *Server

func InitServer(cfg *config.Config) {
	serv = &Server{
		Pid:     os.Getpid(),
		Cfg:     cfg,
		DB:      db.NewDB(),
		Clients: strcture.NewList(),
	}
}

func GetServ() *Server {
	return serv
}
