package command

import (
	"fmt"
	"time"
	"tinydb/internal/app/db"
	"tinydb/internal/app/object"
	"tinydb/internal/app/zerrors"
)

const (
	ObjSetWithNoFlag = iota
	ObjSetNx
	ObjSetXx
	ObjSetEx
	ObjSetPx
	ObjKeepTTL
)

const (
	SetFail = iota
	SetSuccess
)

// Get Get命令 支持1个参数
func Get(db *db.DB, param []*object.Object) (result *object.Object, err error) {
	return get(db, param[0])
}

// Set Set命令 支持2个参数
func Set(db *db.DB, param []*object.Object) (result *object.Object, err error) {
	result, err = set(db, ObjSetWithNoFlag, param[0], param[1], nil)
	return
}

// Setex Setex命令 支持3个参数 key value expire(过期时间 单位为秒)
func Setex(db *db.DB, param []*object.Object) (result *object.Object, err error) {
	result, err = set(db, ObjSetEx, param[0], param[1], param[2])
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

	err = zerrors.NewTypeError(fmt.Sprintf("it`s not a string type value %v", result.GetValue()))
	return
}

func set(db *db.DB, flag int, key *object.Object, value *object.Object, expire *object.Object) (result *object.Object, err error) {
	if flag == ObjSetXx && db.LookUpKeyWrite(key) != nil {
		result = object.NewIntObject(SetFail)
		return
	}
	if flag == ObjSetXx && db.LookUpKeyWrite(key) == nil {
		result = object.NewIntObject(SetFail)
		return
	}

	var ttl int
	if expire != nil {
		ttl, err = expire.GetIntValue()
		if err != nil {
			return
		}
		if ttl < 0 {
			err = zerrors.NewParameterError("invalid expire time")
			return
		}
		ttl += int(time.Now().Unix()) // 暂时用int存储时间戳
	}

	db.SetKey(key, value, flag == ObjKeepTTL)
	if expire != nil {
		db.SetExpire(key, object.NewIntObject(ttl))
	}

	result = object.NewIntObject(SetSuccess)
	return
}
