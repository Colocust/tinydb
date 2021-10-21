package command

import (
	"github.com/Colocust/strcture"
	"log"
	"tinydb/internal/app/db"
	"tinydb/internal/app/object"
)

type (
	Command struct {
		Arity int // 正数代表参数的个数，负数代表参数个数>-Arity
		Func  HandlerFunc
	}
	HandlerFunc func(db *db.DB, param []*object.Object) (result *object.Object, err error)
)

var Commands *strcture.Dict

func init() {
	Commands = strcture.NewDict()

	Commands.Set("Get", &Command{
		Arity: 1,
		Func:  Get,
	})
	Commands.Set("Set", &Command{
		Arity: 2,
		Func:  Set,
	})
	Commands.Set("Setex", &Command{
		Arity: 3,
		Func:  Setex,
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
