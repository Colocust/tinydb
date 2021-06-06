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

	PrintlnListFromHead(l)
}

func TestList_AddNodeTail(t *testing.T) {
	l := NewList()
	l.AddNodeTail("s1")
	l.AddNodeTail("s2")
	l.AddNodeTail("s3")

	PrintlnListFromHead(l)
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

	PrintlnListFromHead(l)

	tail := l.Tail()
	l.DelNode(tail)

	PrintlnListFromHead(l)

	fmt.Println(l.head, l.tail)
}

func TestList_InsertNode(t *testing.T) {
	l := NewList()
	l.AddNodeTail("s1")
	l.AddNodeTail("s2")
	l.AddNodeTail("s3")

	tail := l.Tail()
	l.InsertNode(tail, "s4", true)

	PrintlnListFromHead(l)
	PrintlnListFromTail(l)
}

func TestList_SearchKey(t *testing.T) {
	l := NewList()

	for i := 1; i <= 1000; i++ {
		l.AddNodeTail(i)
	}

	n := l.SearchKey(499)
	fmt.Println(n)

	PrintlnListFromHead(l)
}

func TestList_Index(t *testing.T) {
	l := NewList()
	for i := 1; i <= 1000; i++ {
		l.AddNodeTail(i)
	}

	index := -1000

	fmt.Println(l.Index(index))
}

func TestList_RotateTailToHead(t *testing.T) {
	l := NewList()
	for i := 1; i <= 500; i++ {
		l.AddNodeTail(i)
	}

	l.RotateTailToHead()
	PrintlnListFromTail(l)
}

func TestList_RotateHeadToTail(t *testing.T) {
	l := NewList()
	for i := 1; i <= 500; i++ {
		l.AddNodeTail(i)
	}

	l.RotateHeadToTail()
	PrintlnListFromTail(l)
}


func PrintlnListFromHead(l *List) {
	cur := l.head
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.next
	}
	fmt.Println(l.len)
}

func PrintlnListFromTail(l *List) {
	cur := l.tail
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.prev
	}
	fmt.Println(l.len)
}
