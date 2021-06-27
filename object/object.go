package object

type object struct {
	t        uint8       //类型
	encoding uint8       //编码
	ptr      interface{} //指针
	lru      uint
}

func newObject(t, e uint8, ptr interface{}) *object {
	return &object{
		t:        t,
		encoding: e,
		ptr:      ptr,
	}
}

func NewStringObject(ptr *string) *object {
	return newObject(objString, encodingRaw, ptr)
}

func NewStringObjectByIntValue(ptr *int64) *object {
	return newObject(objString, encodingInt, ptr)
}
