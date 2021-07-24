package command

import (
	"tinydb/db"
	"tinydb/object"
)

func GetCommand(d *db.DB, key string) *object.Object {
	return getGenericCommand(d, key)
}

func getGenericCommand(d *db.DB, key string) *object.Object {
	o := d.LookupKeyRead(key)
	if o == nil {
		return nil
	}
	if o.T != object.ObjString {
		return nil
	}
	return o
}
