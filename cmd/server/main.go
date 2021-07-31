package main

import (
	"fmt"
	"tinydb/db"
)

type (
	Server struct {
		db *db.DB
	}
)

func main() {
	s := new(Server)
	fmt.Println(s)
}
