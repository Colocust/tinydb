package structure

type (
	ListNode struct {
		value interface{}
		prev  *ListNode
		next  *ListNode
	}

	List struct {
		head *ListNode
		tail *ListNode
		len  uint32
	}
)

func NewList() *List {
	l := new(List)
	l.len = 0
	return l
}

func (l *List) AddNodeHead(value interface{}) {
	node := new(ListNode)
	node.value = value

	if l.len == 0 {
		l.head, l.tail = node, node
	} else {
		node.next, l.head.prev, l.head = l.head, node, node
	}

	l.len++
}

func (l *List) AddNodeTail(value interface{}) {
	node := new(ListNode)
	node.value = value

	if l.len == 0 {
		l.head, l.tail = node, node
	} else {
		node.prev, l.tail.next, l.tail = l.tail, node, node
	}

	l.len++
}

func (l *List) Rotate() {

}

//指定节点前后插入一个新节点 after决定前后
func (l *List) InsertNode(oldNode *ListNode, value interface{}, after bool) {
	node := new(ListNode)
	node.value = value

	if after {
		//先更新新节点前后指针
		node.prev, node.next = oldNode, oldNode.next
		if l.tail == oldNode {
			l.tail = node
		}
	} else {
		node.next, node.prev = oldNode, oldNode.prev
		if l.head == oldNode {
			l.head = node
		}
	}

	//更新oldNode以及oldNode.next
	if node.prev != nil {
		node.prev.next = node
	}

	if node.next != nil {
		node.next.prev = node
	}

	l.len++
}

//此方法有问题吧 如果删除的节点不在指定的list 需要开发者自行注意
func (l *List) DelNode(n *ListNode) {
	if n.prev == nil {
		l.head = n.next
	} else {
		n.prev.next = n.next
	}

	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
	}

	l.len--
}

func (l *List) GetLen() uint32 {
	return l.len
}

func (l *List) Head() *ListNode {
	return l.head
}

func (l *List) Tail() *ListNode {
	return l.tail
}

func (n *ListNode) GetValue() interface{} {
	return n.value
}

func (n *ListNode) Prev() *ListNode {
	return n.prev
}

func (n *ListNode) Next() *ListNode {
	return n.next
}
