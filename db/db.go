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

func (db *DB) setExpire(key string, when int) {
	db.expire.Set(key, when)
}

func (db *DB) getExpire(key string) int {
	expire := db.expire.Get(key)
	if expire == nil {
		return -1
	}
	return expire.(int)
}

func (db *DB) keyIsExpired(key string) bool {
	when := db.getExpire(key)
	if when < 0 {
		return false
	}
	now := time.Now().Second()
	return now > when
}

func (db *DB) expireKey(key string) {
	db.db.Remove(key)
	db.expire.Remove(key)
}

// 删除一个key 当它过期的时候
func (db *DB) expireIfNeeded(key string) bool {
	if !db.keyIsExpired(key) {
		return false
	}
	db.expireKey(key)
	return true
}

func (db *DB) LookupKeyReadOrReply(key string) *object.Object {
	return db.lookupKeyRead(key)
}

func (db *DB) lookupKeyRead(key string) *object.Object {
	return db.lookupKeyReadWithFlags(key, server.LookupNone)
}

func (db *DB) lookupKeyReadWithFlags(key string, flag int) *object.Object {
	if db.expireIfNeeded(key) {
		return nil
	}
	obj := db.lookupKey(key, flag)
	return obj
}

func (db *DB) lookupKey(key string, flag int) *object.Object {
	obj := db.db.Get(key).(*object.Object)
	if obj == nil {
		return nil
	}

	if flag == 1 {
		// 补充LRU
	}
	return obj
}
