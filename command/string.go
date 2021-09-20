package command

import (
	"time"
	"tinydb/db"
	"tinydb/errors"
	"tinydb/object"
)

const (
	ObjSetWithNoFlag = iota
	ObjSetNx
	ObjSetXx
	ObjSetEx
	ObjSetPx
	ObjKeepTTL
)

func Get(db *db.DB, param []*object.Object) (result *object.Object, err error) {
	return get(db, param[0])
}

func Set(db *db.DB, param []*object.Object) (result *object.Object, err error) {
	var flag int

	//for i := 2; i < len(param); i++ {
	//	if i == len(param)-1 {
	//		next := nil
	//	} else {
	//		next := param[i+1]
	//	}
	//}

	if err = set(db, flag, param[0], param[1], nil); err != nil {
		return
	}
	result = object.NewStringObject("OK")
	return
}

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
