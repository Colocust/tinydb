package main

import (
	"fmt"
	"github.com/Colocust/strcture"
	"github.com/tidwall/evio"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"tinydb"
	"tinydb/config"
	"tinydb/db"
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
	tinydb.Serv.Cfg = config.Load()
	initServer()

	var events evio.Events
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		out = []byte("s")
		return
	}

	if err := evio.Serve(events, "tcp://localhost:5000"); err != nil {
		panic(err.Error())
	}
}

func initServer() {
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)

	tinydb.Serv.Pid = os.Getpid()
	tinydb.Serv.Clients = strcture.NewList()
	tinydb.Serv.DB = db.NewDB()

	log.Println("Info: Pid = " + strconv.Itoa(tinydb.Serv.Pid))
}
