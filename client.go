package tinydb

import (
	"strconv"
	"strings"
	"tinydb/db"
	"tinydb/object"
)

type Client struct {
	Id   uint             // 唯一自增ID
	DB   *db.DB           // 当前选择的数据库
	Argc int              // 当前命令的参数个数
	Argv []*object.Object // 当前命令的参数
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
