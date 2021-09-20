package main

import (
	"fmt"
	"github.com/tidwall/evio"
	"log"
	"os/signal"
	"strconv"
	"syscall"
	"tinydb"
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
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	log.Println("Info: Pid = " + strconv.Itoa(tinydb.GetPid()))

	// 引入evio网络库
	var events evio.Events
	events.Data = func(conn evio.Conn, in []byte) (out []byte, action evio.Action) {
		if conn.Context() == nil {
			conn.SetContext(tinydb.NewClient())
		}

		out, action = tinydb.HandleClient(conn, in)
		return
	}

	if err := evio.Serve(events, "tcp://"+tinydb.GetCfg().Addr); err != nil {
		panic(err.Error())
	}
}
