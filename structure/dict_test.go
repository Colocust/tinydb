package structure

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	d := make(map[string]int)

	for i := 0; i < 10000000; i++ {
		d[strconv.Itoa(i)] = i
	}
	fmt.Println("s")

	delete(d,"1")
	fmt.Println(len(d))
}
