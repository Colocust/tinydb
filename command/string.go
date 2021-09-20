package command

import (
	"time"
	"tinydb/db"
	"tinydb/errors"
	"tinydb/object"
)

const (
	ObjSetNx = iota
	ObjSetXx
	ObjSetEx
	ObjSetPx
	ObjKeepTTL
)

func Get(db *db.DB, param []*object.Object) (result *object.Object, err error) {
	return get(db, param[0])
}

//
//func Set(db *db.DB, param []*object.Object) (result *object.Object, err error) {
//
//}

func get(d *db.DB, key *object.Object) (result *object.Object, err error) {
	result = d.LookupKeyReadOrReply(key)

	if result == nil {
		return
	}
	if result.GetType() == object.ObjString {
		return
	}

	err = errors.NewTypeError("it`s not a string type value")
	return
}

func set(db *db.DB, flag int, key *object.Object, value *object.Object, expire *object.Object) (err error) {
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
		db.SetExpire(key, object.NewIntObject(ttl))
	}
	return
}
