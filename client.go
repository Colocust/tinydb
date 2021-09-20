package tinydb

import (
	"github.com/tidwall/evio"
	"strconv"
	"strings"
	"tinydb/command"
	"tinydb/db"
	"tinydb/object"
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
			obj = object.NewIntObject(&i)
		} else {
			obj = object.NewStringObject(&item)
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

	cmd := *(client.Argv[0].GetPtr().(*string))

	f := command.LookUpCommand(cmd)
	if f == nil {
		out = []byte("(error) ERR unknown command " + cmd + "\n")
		return
	}

	_, _ = f.Func(client.Argv[1:])

	return out, action
}
