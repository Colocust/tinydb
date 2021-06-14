package structure

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(randomLevel())
	}
}
