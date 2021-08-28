package command

import (
	"errors"
	"tinydb"
)

type HandlerFunc func(c *tinydb.Client) (result interface{}, err error)

var Commands = make(map[string]HandlerFunc)

func init() {
	Commands["get"] = Get
}

func GetCommand(command string) (f HandlerFunc, err error) {
	if f = Commands[command]; f == nil {
		err = errors.New("wrong command" + command)
	}
	return
}
