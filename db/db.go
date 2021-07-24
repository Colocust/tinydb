package db

import (
	"tinydb/object"
	"tinydb/structure"
)

type DB struct {
	d *structure.Dict
}

func NewDB() *DB {
	return &DB{
		d: structure.NewDict(),
	}
}

func (d *DB) LookupKeyRead(key string) *object.Object {
	o := d.d.Get(key).(*object.Object)
	return o
}
