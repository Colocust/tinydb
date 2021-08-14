package command

import (
	"time"
	"tinydb/db"
	"tinydb/errors"
	"tinydb/object"
)

const (
	ObjSetNx   = 1 << 0
	ObjSetXx   = 1 << 1
	ObjSetEx   = 1 << 2
	ObjSetPx   = 1 << 3
	ObjKeepTTL = 1 << 4
)

func GetCommand(d *db.DB, key *object.Object) (obj *object.Object, err error) {
	return getGenericCommand(d, key)
}

func getGenericCommand(d *db.DB, key *object.Object) (obj *object.Object, err error) {
	obj = d.LookupKeyReadOrReply(key)

	if obj.GetType() == object.ObjString {
		return
	}

	err = errors.NewTypeError("it`s not a string type value")
	return
}

func SetCommand(d *db.DB) {

}

func setGenericCommand(db *db.DB, flag int, key *object.Object, value *object.Object, expire *object.Object) (err error) {
	if (flag == ObjSetNx && db.LookUpKeyWrite(key) != nil) || (flag == ObjSetXx && db.LookUpKeyWrite(key) == nil) {
		return
	}

	var ttl int
	if expire != nil {
		ttl, err = expire.GetIntValue()

		if err != nil {
			return
		}
		if ttl <= 0 {
			err = errors.NewParameterError("invalid expire time")
			return
		}

		ttl += time.Now().Second()
	}

	db.SetKey(key, value, flag == ObjKeepTTL)
	if expire != nil {
		db.SetExpire(key, object.NewIntObject(&ttl))
	}
	return
}
