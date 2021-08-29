package network

import (
	"fmt"
	"log"
	"net"
	"tinydb"
)

func Listen() {
	ln, err := net.Listen("tcp", tinydb.Serv.Cfg.Addr)
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

		// 注册事件 设置事件回调
		acceptTcpHandler(conn)
	}
}

func acceptTcpHandler(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Println(conn.RemoteAddr().String() + "断开连接")
	}()

	log.Println(conn.RemoteAddr().String() + "连接了")

	for {
		var (
			n    int
			err  error
			//data string
		)

		buf := make([]byte, 2048)
		if n, err = conn.Read(buf); err != nil {
			break
		}

		readQueryFromClient(buf, n)

		//data = string(buf[:n])
		//fmt.Println(data)
		//conn.Write([]byte("s \n"))
	}
}

func readQueryFromClient(buf []byte, n int) {
	fmt.Println(string(buf[:n]))
}
