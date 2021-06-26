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

func TestSkipList_UpdateScore(t *testing.T) {
	sl := CreateSkipList()
	n := sl.header
	for i := uint(0); i < sl.length; i++ {
		n = n.level[0].forward
		fmt.Println(n)
	}

	fmt.Println("s")
	sl.UpdateScore(NewSds("5"), 6, 6)
	n = sl.header
	for i := uint(0); i < sl.length; i++ {
		n = n.level[0].forward
		fmt.Println(n)
	}

}

func TestSkipList_Range(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(NewSds("1"), 1)
	sl.Insert(NewSds("2"), 2)
	sl.Insert(NewSds("11"), 11)

	zrs := NewZRangeSpec(7, 9, true, false)
	fmt.Println(sl.FirstInRange(zrs))
	fmt.Println(sl.LastInRange(zrs))
}

func TestSkipList_DeleteByRange(t *testing.T) {
	sl := CreateSkipList()
	Println(sl)
	r := NewZRangeSpec(6, 9, false, true)
	fmt.Println(sl.DeleteByRange(r))
	fmt.Println("sssss")
	Println(sl)
}

func CreateSkipList() *SkipList {
	sl := NewSkipList()

	for i := 10; i > 0; i-- {
		sl.Insert(NewSds(strconv.Itoa(i)), float32(i))
	}
	return sl
}

func Println(sl *SkipList) {
	n := sl.header
	for i := uint(0); i < sl.length; i++ {
		n = n.level[0].forward
		fmt.Println(n)
	}
}
