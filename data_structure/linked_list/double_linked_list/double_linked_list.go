package double_linked_list

type Node struct {
	data string
	next *Node
	prev *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}

