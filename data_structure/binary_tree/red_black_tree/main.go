package main

/*
	红黑树： 查找，插入，删除的最差时间复杂度为O(logN)，红黑树始终保持logN的高度
		1）每个结点要么是红的，要么是黑的。
		2）根结点是黑的。
		3）每个叶结点（叶结点即指树尾端NIL指针或NULL结点）是黑的。
		4）如果一个结点是红的，那么它的俩个儿子都是黑的。
		5）对于任一结点而言，其到叶结点树尾端NIL指针的每一条路径都包含相同数目的黑结点。
*/

type Color int

const (
	Red   Color = 1
	Black Color = 0
)

var NilNode = &Node{
	Data:   0,
	Color:  Black,
	Left:   nil,
	Right:  nil,
	Parent: nil,
}

type Node struct {
	Data int
	Color
	Left   *Node
	Right  *Node
	Parent *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Search(item int) (node *Node) {
	if t.Root == nil {
		return nil
	}
	return t.Root.Search(item)
}
func (n *Node) Search(item int) (node *Node) {
	if n.Data == item {
		return n
	} else if n.Data > item { // Less then
		if n.Left != nil {
			return n.Left.Search(item)
		}
	} else {
		if n.Right != nil {
			return n.Right.Search(item)
		}
	}
	return nil
}

// 节点插入
func (t *Tree) Insert(data int) {
	newNode := &Node{Data: data, Left: NilNode, Right: NilNode}
	if t.Root == nil {
		t.Root = newNode
		return
	}
	t.Root.Insert(newNode)
}

func (n *Node) Insert(node *Node) {
	if n.Data >= node.Data { //如果需要插入的值比当前节点值要小, 往左边写入
		if n.Left == nil {
			n.Left = node
			node.Parent = n
			return
		}
		n.Left.Insert(node)
	} else { //如果需要插入的值比当前节点值要大, 往右边写入
		if n.Right == nil {
			n.Right = node
			node.Parent = n
			return
		}
		n.Right.Insert(node)
	}
}

// 红黑树的调整，将不平衡的树调整为平衡的树
func (t *Tree) Adjust() {
	//  TODO: To be continue
}
