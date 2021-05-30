package structure

import (
	"fmt"
	"testing"
)

func TestSDS(t *testing.T) {
	s1 := NewSds("s1")
	//s2 := NewSds("s2")
	//s3 := NewSds("s3")

	//fmt.Println(s1.GetLen(), s2.GetLen(), s3.GetLen())
	//fmt.Println(string(s1.GetBuf()), string(s2.GetBuf()), string(s3.GetBuf()))
	//
	//s1.CatSds(s2)
	//fmt.Println(s1.GetLen(), s2.GetLen())
	//fmt.Println(string(s1.GetBuf()), string(s2.GetBuf()))
	//
	//s1.CatString("s4")
	//fmt.Println(s1.GetLen())
	//fmt.Println(string(s1.GetBuf()))
	//
	//s1.Cpy("s5")
	//fmt.Println(s1.GetLen())
	//fmt.Println(string(s1.GetBuf()))

	s1.Cpy("s1s2s3")
	//fmt.Println(s1.GetLen())
	fmt.Println(cap(s1.buf))

	s1.buf = s1.buf[4:4]
	//fmt.Println(s1.GetLen())
	fmt.Println(cap(s1.buf))
	fmt.Println(len(s1.buf))

}
