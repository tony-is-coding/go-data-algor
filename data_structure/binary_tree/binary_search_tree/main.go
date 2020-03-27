package main

import "fmt"

/* 二叉查找树满足以下条件
	1. 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
	2. 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
	3. 任意节点的左、右子树也分别为二叉查找树；
	4. 没有键值相等的节点 // ps
*/
type Node struct {
	Label  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) IsEmpty() bool {
	return t.Root == nil
}
func (n *Node) IsEmpty() bool {
	return n == nil
}

func (t *Tree) Insert(label int) {
	newNode := &Node{Label: label}
	if t.IsEmpty() {
		t.Root = newNode
		return
	}
	t.Root.Insert(newNode)
}
func (n *Node) Insert(node *Node) {
	if n.Label >= node.Label { //如果需要插入的值比当前节点值要小, 往左边写入
		if n.Left.IsEmpty() {
			n.Left = node
			node.Parent = n
			return
		}
		n.Left.Insert(node)
	} else { //如果需要插入的值比当前节点值要大, 往右边写入
		if n.Right.IsEmpty() {
			n.Right = node
			node.Parent = n
			return
		}
		n.Right.Insert(node)
	}
}

func (t *Tree) Delete(label int) (hasDelete bool) {
	has, node := t.Search(label)
	if ! has { // 没找到
		return false
	}
	if node.Label == t.Root.Label { // 删除的为root根节点
		//node.Label
	}
	if node.Left == nil && node.Right != nil {
		node.SwapNull(node.Right, t) // 将当前节点的右子节点与当前节点替换
	} else if node.Right == nil && node.Left != nil {
		node.SwapNull(node.Left, t) // 将当前节点的左子节点与当前节点替换
	} else if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			t.Root = nil
		}
	} else {
		return
	}
	return true
}

// 节点替换
func (n *Node) SwapNull(swapNode *Node, tree *Tree) {
	if n.Parent == nil {
		tree.Root = swapNode
	}
	if n.Label == n.Parent.Left.Label { // 是左子树
		swapNode.Parent = n.Parent
		n.Parent.Left = swapNode

	} else {
		swapNode.Parent = n.Parent
		n.Parent.Right = swapNode
	}
}

func (n *Node) Delete() {

}

func (t *Tree) Search(label int) (has bool, node *Node) {
	if t.IsEmpty() {
		return false, nil
	}
	return t.Root.Search(label)
}
func (n *Node) Search(label int) (has bool, node *Node) {
	if n.Label == label { // 如果当前值为需要查找值，则返回
		return true, n
	} else if n.Label > label && !n.Left.IsEmpty() {
		return n.Left.Search(label)
	} else if n.Label < label && !n.Right.IsEmpty() {
		return n.Right.Search(label)
	} else {
		return false, nil
	}
}
func (t *Tree) Traversal() {
	if t.IsEmpty() {
		return
	}
	t.Root.Traversal()
}
func (n *Node) Traversal() {
	if !n.Left.IsEmpty() {
		n.Left.Traversal()
	}
	fmt.Println(n.Label)
	if !n.Right.IsEmpty() {
		n.Right.Traversal()
	}
}

func main() {
	tree := Tree{}
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(5)
	tree.Traversal()
}
