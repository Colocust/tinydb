package command

import (
	"tinydb/db"
	"tinydb/object"
)

func GetCommand(d *db.DB, key *object.Object) *object.Object {
	return getGenericCommand(d, key)
}

func getGenericCommand(d *db.DB, key *object.Object) *object.Object {
	obj := d.LookupKeyReadOrReply(key)
	if obj == nil {
		return nil
	}
	if obj.GetType() != object.ObjString {
		return nil
	}
	return obj
}

func setGenericCommand(d *db.DB, flag int, key string, value string, expire int) {

}
