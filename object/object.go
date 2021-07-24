package object

import "tinydb/structure"

type Object struct {
	t        uint8
	encoding uint8
	ptr      interface{}
}

func newObject(t, e uint8, p interface{}) *Object {
	return &Object{
		t:        t,
		encoding: e,
		ptr:      p,
	}
}

func (obj *Object) GetType() uint8 {
	return obj.t
}

func (obj *Object) GetEncoding() uint8 {
	return obj.encoding
}

func (obj *Object) GetPtr() interface{} {
	return obj.ptr
}

func createSdsObject(p *structure.Sds) *Object {
	return newObject(ObjString, EncodingRaw, p)
}
