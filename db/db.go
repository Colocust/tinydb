package db

import (
	"tinydb/object"
	"tinydb/server"
	"tinydb/structure"
)

type DB struct {
	d *structure.Dict
	e *structure.Dict
}

func NewDB() *DB {
	return &DB{
		d: structure.NewDict(),
	}
}

func (db *DB) keyIsExpired(key string) bool {
	return true
}

// 删除一个key 当它过期的时候
func (db *DB) expireIfNeeded(key string) bool {
	if !db.keyIsExpired(key) {
		return false
	}

	server.StatExpiredKey++
	// 同步/异步删除key (根据配置而定)

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
		server.StatMissesKey++
	}
	obj := db.lookupKey(key, flag)
	if obj == nil {
		server.StatMissesKey++
	} else {
		server.StatHitsKey++
	}

	return obj
}

func (db *DB) lookupKey(key string, flag int) *object.Object {
	obj := db.d.Get(key).(*object.Object)
	if obj != nil {
		if flag == 1 {
			// 补充LRU
		}
		return obj
	}
	return nil
}
