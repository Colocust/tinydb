package object

import "tinydb/structure"

type Object struct {
	T        uint8
	Encoding uint8
	P        interface{}
}

func newObject(t, e uint8, p interface{}) *Object {
	return &Object{
		T:        t,
		Encoding: e,
		P:        p,
	}
}

func createSdsObject(p *structure.Sds) *Object {
	return newObject(ObjString, EncodingRaw, p)
}
