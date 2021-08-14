package object

import (
	"tinydb/errors"
)

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

func (obj *Object) SetType(t uint8) {
	obj.t = t
}

func (obj *Object) SetEncoding(encoding uint8) {
	obj.encoding = encoding
}

func (obj *Object) SetPtr(ptr interface{}) {
	obj.ptr = ptr
}

func NewStringObject(p *string) *Object {
	return newObject(ObjString, EncodingRaw, p)
}

func NewIntObject(p *int) *Object {
	return newObject(ObjString, EncodingInt, p)
}

func (obj *Object) GetIntValue() (value int, err error) {
	if obj.GetType() != ObjString {
		err = errors.NewTypeError("it`s not a string type value")
		return
	}
	if obj.GetEncoding() != EncodingInt {
		err = errors.NewEncodingError("it`s not an int encoding value")
		return
	}

	value = *obj.ptr.(*int)
	return
}
