package structure

type sds struct {
	buf []byte
}

func NewSds(string string) *sds {
	s := new(sds)
	s.buf = []byte(string)

	return s
}

//获取sds长度
func (sds *sds) GetLen() int {
	return len(sds.buf)
}

func (sds *sds) cat(b []byte) {
	sds.buf = append(sds.buf, b...)
}

//将一个字符串追加到指定sds后面
func (sds *sds) CatString(s string) {
	sds.cat([]byte(s))
}

//将一个sds追加到指定sds后面
func (sds *sds) CatSds(s *sds) {
	sds.cat(s.buf)
}

//获取sds内容
func (sds *sds) GetBuf() []byte {
	return sds.buf
}

//设置sds内容
func (sds *sds) Cpy(s string) {
	sds.buf = []byte(s)
}

//保留sds指定区间内的数据
func (sds *sds) Range(start, end int) {
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
