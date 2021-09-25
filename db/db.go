package db

import (
	"github.com/Colocust/strcture"
	"time"
	"tinydb/object"
)

type DB struct {
	db     *strcture.Dict
	expire *strcture.Dict
}

const (
	LookupNone = iota
	LookupNoTouch
	LookupNoNotify
)

func NewDB() *DB {
	return &DB{
		db:     strcture.NewDict(),
		expire: strcture.NewDict(),
	}
}

func (db *DB) SetExpire(key *object.Object, when *object.Object) {
	db.expire.Set(key.GetValue().(string), when)
}

func (db *DB) getExpire(key *object.Object) int {
	expire := db.expire.Get(key.GetValue().(string))
	if expire == nil {
		return -1
	}
	return expire.(*object.Object).GetValue().(int)
}

func (db *DB) keyIsExpired(key *object.Object) bool {
	when := db.getExpire(key)
	if when < 0 {
		return false
	}
	now := int(time.Now().Unix())
	return now > when
}

func (db *DB) expireKey(key *object.Object) {
	db.db.Remove(key.GetValue().(string))
	db.expire.Remove(key.GetValue().(string))
}

// 删除一个key 当它过期的时候
func (db *DB) expireIfNeeded(key *object.Object) bool {
	if !db.keyIsExpired(key) {
		return false
	}
	db.expireKey(key)
	return true
}

func (db *DB) LookupKeyReadOrReply(key *object.Object) *object.Object {
	return db.lookupKeyRead(key)
}

func (db *DB) lookupKeyRead(key *object.Object) *object.Object {
	return db.lookupKeyReadWithFlags(key, LookupNone)
}

func (db *DB) lookupKeyReadWithFlags(key *object.Object, flag int) *object.Object {
	if db.expireIfNeeded(key) {
		return nil
	}
	obj := db.lookupKey(key, flag)
	return obj
}

func (db *DB) lookupKey(key *object.Object, flag int) *object.Object {
	value := db.db.Get(key.GetValue().(string))
	if value == nil {
		return nil
	}

	obj := value.(*object.Object)
	if flag == 1 {
		// 补充LRU
	}
	return obj
}

func (db *DB) LookUpKeyWrite(key *object.Object) *object.Object {
	return db.lookUpKeyWriteWithFlags(key, LookupNone)
}

func (db *DB) lookUpKeyWriteWithFlags(key *object.Object, flag int) *object.Object {
	db.expireIfNeeded(key)
	return db.lookupKey(key, flag)
}

func (db *DB) SetKey(key *object.Object, value *object.Object, keepTTL bool) {
	if !keepTTL {
		db.expire.Remove(key.GetValue().(string))
	}
	db.db.Set(key.GetValue().(string), value)
}
