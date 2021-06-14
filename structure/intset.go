package structure

import "math/rand"

type IntSet struct {
	contents []int
}

func NewIntSet() *IntSet {
	return &IntSet{
		contents: make([]int, 0),
	}
}

//添加新元素时先查找对应的下标，如果元素已存在，那么添加失败
func (is *IntSet) Add(value int) bool {
	pos, exist := is.search(value)
	if exist {
		return false
	}
	is.contents = append(is.contents, 0)
	copy(is.contents[pos+1:], is.contents[pos:])
	is.Set(pos, value)
	return true
}

func (is *IntSet) Remove(value int) bool {
	pos, exist := is.search(value)
	if !exist {
		return false
	}
	is.contents = append(is.contents[:pos], is.contents[pos+1:]...)
	return true
}

func (is *IntSet) Find(value int) (exist bool) {
	_, exist = is.search(value)
	return
}

func (is *IntSet) Random() int {
	return is.contents[rand.Intn(len(is.contents))]
}

func (is *IntSet) Get(pos int) int {
	return is.contents[pos]
}

func (is *IntSet) Set(pos int, value int) {
	is.contents[pos] = value
}

// 根据value找到所对应的下标
// 如果IntSet为空或者value小于IntSet最小值，那么返回0
// 如果value大于IntSet最大值 返回IntSet长度
// 采用二分搜索，查找value
func (is *IntSet) search(value int) (pos int, exist bool) {
	min, max, mid, cur := 0, len(is.contents)-1, -1, -1

	if len(is.contents) == 0 {
		pos, exist = min, false
		return
	} else {
		if value > is.Get(max) {
			pos, exist = len(is.contents), false
		} else if value < is.Get(min) {
			pos, exist = min, false
		}
	}

	for max >= min {
		mid = (min + max) >> 1
		cur = is.Get(mid)

		if value > cur {
			min = mid + 1
		} else if value < cur {
			max = mid - 1
		} else {
			break
		}
	}

	if cur == value {
		pos, exist = mid, true
	} else {
		pos, exist = min, false
	}
	return
}
