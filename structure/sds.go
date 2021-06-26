package structure

import (
	"strings"
)

type Sds struct {
	buf []byte
}

//创建一个包含给定字符串的sds
func NewSds(string string) *Sds {
	return &Sds{buf: []byte(string)}
}

//创建一个空的sds
func EmptySds() *Sds {
	return new(Sds)
}

//获取sds长度
func (sds *Sds) GetLen() int {
	return len(sds.buf)
}

func (sds *Sds) cat(b []byte) {
	sds.buf = append(sds.buf, b...)
}

//将一个字符串追加到指定sds后面
func (sds *Sds) CatString(s string) {
	sds.cat([]byte(s))
}

//将一个sds追加到指定sds后面
func (sds *Sds) CatSds(s *Sds) {
	sds.cat(s.buf)
}

//获取sds内容
func (sds *Sds) GetBuf() []byte {
	return sds.buf
}

func (sds *Sds) GetString() string {
	return string(sds.GetBuf())
}

//设置sds内容
func (sds *Sds) Cpy(s string) {
	sds.buf = []byte(s)
}

//保留sds指定区间内的数据
func (sds *Sds) Range(start, end int) {
	l := len(sds.buf)
	if l == 0 {
		return
	}

	if start < 0 {
		if start = l + start; start < 0 {
			start = 0
		}
	}
	if end < 0 {
		if end = l + end; end < 0 {
			end = 0
		}
	}

	if start >= l {
		start = 0
	}
	if end >= l {
		end = l - 1
	}

	if start > end {
		start, end = 0, -1
	}

	sds.buf = sds.buf[start : end+1]
}

//比较两个sds字符串是否相同
func (sds *Sds) Cmp(s *Sds) int {
	return strings.Compare(string(sds.buf), string(s.buf))
}

//去除在sds出现在s中的字符
func (sds *Sds) Trim(s string) {
	buf := make([]byte, 0)
loop:
	for _, element := range sds.buf {
		for _, j := range s {
			if int32(element) == j {
				continue loop
			}
		}
		buf = append(buf, element)
	}

	sds.buf = buf
}
