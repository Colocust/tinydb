package structure

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	d := NewDict()

	k := NewSds("key")
	d.Set(k,"value")
	fmt.Println(d.Get(k))
	d.Remove(k)
	fmt.Println(d.Get(k))
}
