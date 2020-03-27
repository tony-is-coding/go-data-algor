package single_linked_list

type List struct {
	size int
	head *Node
	tail *Node
}
type Node struct {
	next *Node
	data string
}

func NewNode(value string) *Node {
	return &Node{nil, value}
}

func (l *List) Append(value string) {
	node := NewNode(value)
	if l.tail == nil { // head == tail == nil
		l.head = node
		l.tail = node
		l.size += 1
		return
	}
	l.tail.next = node
	l.size += 1
	return
}

func (l *List) Search(value string) *Node {
	currNode := l.head
	for currNode.next != nil {
		if currNode.data == value {
			return currNode
		}
	}
	return nil
}
func (l *List) Insert(index int, value string) {
	if index == 0 {

	} else if index == (l.size - 1) {

	} else {

	}
}
func (l *List) Delete(index int) *Node { return nil }


// 链表反转
// 某个节点开始反转
func (l *List) reverse(n *Node) {
	temp := n.next
	n.next = nil
	temp.reverse(n)
}

func (n *Node) reverse(prev *Node) {
	temp := n.next
	n.next = prev
	temp.reverse(n)
}
