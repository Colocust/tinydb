package structure

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()
	fmt.Println(sl)
	fmt.Println(sl.header)
	fmt.Println(sl.header.ele)
}
