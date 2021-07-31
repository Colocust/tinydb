package command

import (
	"time"
	"tinydb/db"
	"tinydb/object"
)

const (
	ObjSetNx   = 1 << 0
	ObjSetXx   = 1 << 1
	ObjSetEx   = 1 << 2
	ObjSetPx   = 1 << 3
	ObjKeepTTL = 1 << 4
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

func SetCommand(d *db.DB) {

}

func setGenericCommand(d *db.DB, flag int, key *object.Object, value *object.Object, expire *object.Object) {
	millisecond := 0
	if expire != nil {
		if err := expire.GetIntOrReply(&millisecond); err != nil {
			return
		}
		if millisecond <= 0 {
			//addReply
			return
		}
	}
	if (flag == ObjSetNx && d.LookUpKeyWrite(key) != nil) || (flag == ObjSetXx && d.LookUpKeyWrite(key) == nil) {
		return
	}

	d.SetKey(key, value, flag == ObjKeepTTL)
	if expire != nil {
		millisecond += time.Now().Second()
		d.SetExpire(key, object.NewIntObject(&millisecond))
	}
}
