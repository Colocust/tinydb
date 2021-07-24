package command

import (
	"tinydb/db"
	"tinydb/object"
)

func GetCommand(d *db.DB, key string) *object.Object {
	return getGenericCommand(d, key)
}

func getGenericCommand(d *db.DB, key string) *object.Object {
	obj := d.LookupKeyReadOrReply(key)
	if obj == nil {
		return nil
	}
	if obj.GetType() != object.ObjString {
		return nil
	}
	return obj
}
