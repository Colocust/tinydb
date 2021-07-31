package object

import (
	"errors"
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

func (obj *Object) GetIntOrReply(target *int) error {
	var value int
	if err := obj.getInt(&value); err != nil {
		//addReply
		return err
	}
	target = &value
	return nil
}

func (obj *Object) getInt(target *int) error {
	var value int
	if obj.GetType() != ObjString || obj.GetEncoding() != EncodingInt {
		return errors.New("wrong type or wrong encoding")
	}

	value = *obj.ptr.(*int)
	target = &value
	return nil
}
