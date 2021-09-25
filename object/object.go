package object

import (
	"tinydb/errors"
)

type Object struct {
	t        uint8
	encoding uint8
	value    interface{}
}

func newObject(t, e uint8, p interface{}) *Object {
	return &Object{
		t:        t,
		encoding: e,
		value:    p,
	}
}

func (obj *Object) GetType() uint8 {
	return obj.t
}

func (obj *Object) GetEncoding() uint8 {
	return obj.encoding
}

func (obj *Object) GetValue() interface{} {
	return obj.value
}

func (obj *Object) SetType(t uint8) {
	obj.t = t
}

func (obj *Object) SetEncoding(encoding uint8) {
	obj.encoding = encoding
}

func (obj *Object) SetValue(ptr interface{}) {
	obj.value = ptr
}

func NewStringObject(value string) *Object {
	return newObject(ObjString, EncodingRaw, value)
}

func NewIntObject(value int) *Object {
	return newObject(ObjString, EncodingInt, value)
}

func NewLongLongObject(value int64) *Object {
	return newObject(ObjString, EncodingLong, value)
}

func (obj *Object) GetIntValue() (value int, err error) {
	if obj.GetType() != ObjString {
		err = errors.NewTypeError("it`s not a string type value")
		return
	}
	if obj.GetEncoding() != EncodingInt {
		err = errors.NewEncodingError(obj.value.(string) + " is not an int encoding value")
		return
	}

	value = obj.value.(int)
	return
}
