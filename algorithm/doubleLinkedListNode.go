package algorithm

type LinkedListNode struct {
	prev  *LinkedListNode
	next  *LinkedListNode
	value int
}

func (l *LinkedListNode) New(val int) *LinkedListNode {
	node := &LinkedListNode{
		prev:  nil,
		next:  nil,
		value: val,
	}

	return node
}
