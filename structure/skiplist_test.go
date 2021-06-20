package structure

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(NewSds("1"), 1)
	sl.Insert(NewSds("2"), 2)
	sl.Insert(NewSds("2"), 3)

	n := sl.header
	for i := uint(0); i < sl.length; i++ {
		n = n.level[0].forward
		fmt.Println(n)
	}
}

func TestSkipList_Delete(t *testing.T) {
	sl := NewSkipList()

	for i := 10; i > 0; i-- {
		sl.Insert(NewSds(strconv.Itoa(i)), float32(i))
	}
	n := sl.header
	for i := uint(0); i < sl.length; i++ {
		n = n.level[0].forward
		fmt.Println(n)
	}

	fmt.Println("sss")
	sl.Delete(NewSds("10"), 10)

	n = sl.header
	for i := uint(0); i < sl.length; i++ {
		n = n.level[0].forward
		fmt.Println(n)
	}
	fmt.Println(sl.tail)
}
