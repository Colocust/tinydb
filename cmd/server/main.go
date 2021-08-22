package main

import (
	"fmt"
	"log"
	"net"
	"os/signal"
	"strconv"
	"syscall"
	"tinydb/config"
	"tinydb/enum"
	"tinydb/server"
)

func init() {
	fmt.Println(

		" _____    ____   \n" +
			"|  __ \\  |  _ \\ \n" +
			"| |  | | | |_) |  \n" +
			"| |  | | |  _ <   \n" +
			"| |__| | | |_) |  \n" +
			"|_____/  |____/ ")
	log.SetFlags(log.LstdFlags)
}

func main() {
	var (
		cfg  *config.Config
		ok   int
		serv *server.Server
	)

	// 加载配置
	if cfg, ok = config.Load(); ok == enum.ERR {
		return
	}

	server.InitServer(cfg)
	serv = server.GetServ()
	log.Println("Info: Pid = " + strconv.Itoa(serv.Pid))
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)

	handleTcp()
}

func handleTcp() {
	ln, err := net.Listen("tcp", server.GetServ().Cfg.Addr)
	if err != nil {
		log.Println("Error: server boot error，The cause of the error is " + err.Error())
		return
	}

	for {
		var conn net.Conn
		if conn, err = ln.Accept(); err != nil {
			log.Println("Warning: accept error")
			continue
		}

		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer func() {
		c.Close()
		log.Println(c.RemoteAddr().String() + "断开连接")
	}()

	log.Println(c.RemoteAddr().String() + "连接了")

	for {
		var (
			n    int
			err  error
			data string
		)

		buf := make([]byte, 2048)
		if n, err = c.Read(buf); err != nil {
			break
		}

		data = string(buf[:n])
		fmt.Println(data)
		c.Write([]byte("s \n"))
	}
}
