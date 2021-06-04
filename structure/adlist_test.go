package structure

import (
	"fmt"
	"testing"
)

func TestList_AddNodeHead(t *testing.T) {
	l := NewList()
	l.AddNodeHead("s1")
	l.AddNodeHead("s2")
	l.AddNodeHead("s3")

	PrintlnList(l)
}

func TestList_AddNodeTail(t *testing.T) {
	l := NewList()
	l.AddNodeTail("s1")
	l.AddNodeTail("s2")
	l.AddNodeTail("s3")

	PrintlnList(l)
}

func TestList_Head(t *testing.T) {
	l := NewList()
	l.AddNodeTail("s1")
	fmt.Println(l.Head().value)
}

func TestList_DelNode(t *testing.T) {
	l := NewList()
	l.AddNodeTail("s1")
	l.AddNodeTail("s2")
	l.AddNodeTail("s3")

	head := l.Head()
	l.DelNode(head)

	PrintlnList(l)

	tail := l.Tail()
	l.DelNode(tail)

	PrintlnList(l)

	fmt.Println(l.head, l.tail)
}

func PrintlnList(l *List) {
	cur := l.head
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.next
	}
	fmt.Println(l.len)
}
