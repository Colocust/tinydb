package db

import (
	"time"
	"tinydb/object"
	"tinydb/server"
	"tinydb/structure"
)

type DB struct {
	db     *structure.Dict
	expire *structure.Dict
}

func NewDB() *DB {
	return &DB{
		db: structure.NewDict(),
	}
}

func (db *DB) setExpire(key *object.Object, when *object.Object) {
	db.expire.Set(*key.GetPtr().(*string), when)
}

func (db *DB) getExpire(key *object.Object) int {
	expire := db.expire.Get(*key.GetPtr().(*string))
	if expire == nil {
		return -1
	}
	return *expire.(*object.Object).GetPtr().(*int)
}

func (db *DB) keyIsExpired(key *object.Object) bool {
	when := db.getExpire(key)
	if when < 0 {
		return false
	}
	now := time.Now().Second()
	return now > when
}

func (db *DB) expireKey(key *object.Object) {
	db.db.Remove(*key.GetPtr().(*string))
	db.expire.Remove(*key.GetPtr().(*string))
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
	return db.lookupKeyReadWithFlags(key, server.LookupNone)
}

func (db *DB) lookupKeyReadWithFlags(key *object.Object, flag int) *object.Object {
	if db.expireIfNeeded(key) {
		return nil
	}
	obj := db.lookupKey(key, flag)
	return obj
}

func (db *DB) lookupKey(key *object.Object, flag int) *object.Object {
	obj := db.db.Get(*key.GetPtr().(*string)).(*object.Object)
	if obj == nil {
		return nil
	}

	if flag == 1 {
		// 补充LRU
	}
	return obj
}
