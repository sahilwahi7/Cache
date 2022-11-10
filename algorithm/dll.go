package algorithm

type Dll interface {
	buildConnection()
	DetachNode(node *LinkedListNode)
	AddNodetoLast(node *LinkedListNode)
	AddElementAtLast(val int) *LinkedListNode
	getLastNode() *LinkedListNode
	GetFirstNode() *LinkedListNode
}
type dll struct {
	dummyhead *LinkedListNode
	dummytail *LinkedListNode
	Node      LinkedListNode
}

func (d *dll) buildConnection() {
	d.dummytail.prev = d.dummyhead
	d.dummyhead.next = d.dummytail
}

func (d *dll) DetachNode(node *LinkedListNode) {
	if node == nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}
func (d *dll) AddNodetoLast(node *LinkedListNode) {
	tailprev := d.dummytail.prev
	tailprev.next = node
	node.next = d.dummytail
	d.dummytail.prev = node
	node.prev = tailprev
}

func (d *dll) AddElementAtLast(val int) *LinkedListNode {
	newNode := d.Node.New(val)
	d.AddNodetoLast(newNode)
	return newNode
}

func (d *dll) isItemPresent() bool {
	return d.dummyhead.next != d.dummytail
}

func (d *dll) GetFirstNode() *LinkedListNode {
	if d.isItemPresent() == false {
		return nil
	}
	return d.dummyhead.next
}

func (d *dll) getLastNode() *LinkedListNode {
	if d.isItemPresent() == false {
		return nil
	}
	return d.dummytail.prev
}
