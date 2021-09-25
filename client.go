package tinydb

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/evio"
	"strconv"
	"strings"
	"tinydb/command"
	"tinydb/db"
	"tinydb/object"
	"tinydb/tool"
)

type Client struct {
	Id   uint                // 唯一自增ID
	DB   *db.DB              // 当前选择的数据库
	Argc int                 // 当前命令的参数个数
	Argv []*object.Object    // 当前命令的参数
	Cmd  command.HandlerFunc // 执行的命令
}

func NewClient() *Client {
	return &Client{
		DB: GetDB(),
	}
}

func (c *Client) ReadQueryFromClient(in []byte) {
	inString := strings.TrimSpace(string(in))
	inSlice := strings.Split(inString, " ")

	var query []*object.Object
	for _, item := range inSlice {
		if len(item) == 0 {
			continue
		}
		var obj *object.Object
		if i, err := strconv.Atoi(item); err == nil {
			obj = object.NewIntObject(i)
		} else {
			obj = object.NewStringObject(item)
		}
		query = append(query, obj)
	}

	c.Argc = len(query)
	c.Argv = query
}

func HandleClient(conn evio.Conn, in []byte) (out []byte, action evio.Action) {
	client := conn.Context().(*Client)
	client.ReadQueryFromClient(in)

	// 空的指令
	if len(client.Argv) == 0 {
		return
	}

	key, argc, argv := tool.FirstUpper(client.Argv[0].GetValue().(string)), client.Argc-1, client.Argv[1:]

	cmd, err := GetCommand(key, argc)
	if err != nil {
		out = []byte(err.Error())
		return
	}

	if resp, err := cmd.Func(client.DB, argv); err != nil {
		out = []byte(err.Error())
		return
	} else {
		if resp == nil {
			out = []byte("nil")
			return
		}
		if resp.GetEncoding() == object.EncodingRaw {
			out = []byte(resp.GetValue().(string))
			return
		}

		out, _ = json.Marshal(resp.GetValue().(int))
		return
	}
}

func GetCommand(key string, argc int) (cmd *command.Command, err error) {
	cmd = command.LookUpCommand(key)

	if cmd == nil {
		err = errors.New("(error) ERR unknown command " + key)
		return
	}

	if (cmd.Arity > 0 && argc != cmd.Arity) || (cmd.Arity < 0 && -cmd.Arity > argc) {
		err = errors.New("wrong number of arguments for " + key + " command")
		return
	}

	return
}
