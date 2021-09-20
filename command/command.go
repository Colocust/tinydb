package command

import (
	"github.com/Colocust/strcture"
	"log"
	"tinydb/db"
	"tinydb/object"
)

type (
	Command struct {
		Func HandlerFunc
	}
	HandlerFunc func(db *db.DB, param []*object.Object) (result *object.Object, err error)
)

var Commands *strcture.Dict

func init() {
	Commands = strcture.NewDict()

	Commands.Set("Get", &Command{
		Func: Get,
	})
}

func LookUpCommand(argv string) *Command {
	cmd := Commands.Get(argv)

	if cmd == nil {
		log.Println("(error) ERR unknown command " + argv)
		return nil
	}

	return cmd.(*Command)
}
