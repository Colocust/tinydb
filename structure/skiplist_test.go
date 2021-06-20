package structure

import (
	"fmt"
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
