package main

import (
	"fmt"
	"github.com/tidwall/evio"
	"log"
	"os/signal"
	"strconv"
	"syscall"
	"tinydb/internal/app"
	_ "tinydb/internal/app/command"
)

func init() {
	fmt.Println(

		" _____    ____   \n" +
			"|  __ \\  |  _ \\ \n" +
			"| |  | | | |_) |  \n" +
			"| |  | | |  _ <   \n" +
			"| |__| | | |_) |  \n" +
			"|_____/  |____/ ")
	fmt.Println()
	log.SetFlags(log.LstdFlags)
}

func main() {
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	log.Println(fmt.Sprintf("Info: Pid = %s", strconv.Itoa(app.Serv.GetPid())))

	var events evio.Events
	events.Data = func(conn evio.Conn, in []byte) (out []byte, action evio.Action) {
		if conn.Context() == nil {
			conn.SetContext(app.NewClient())
		}

		out, action = app.HandleClient(conn, in)
		return
	}

	if err := evio.Serve(events, "tcp://"+app.Serv.GetCfg().Addr); err != nil {
		panic(err.Error())
	}
}
