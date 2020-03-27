package main

import "fmt"

/*
	基本二叉树实现:
		每个节点最多有2 个子节点
*/

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}
type Tree struct {
	Root *Node
}

func (t *Tree) IsEmpty() bool {
	return t.Root == nil
}
func (t *Tree) Display() {
	if t.IsEmpty() {
		return
	}
	t.Root.Display()

}
func (n *Node) IsEmpty() bool {
	return n == nil
}

func (n *Node) Display() {
	fmt.Println(n.Data)
	if !n.Left.IsEmpty() {
		n.Left.Display()
	}
	if !n.Right.IsEmpty() {
		n.Right.Display()
	}
}

func main() {
	t := Tree{Root: &Node{Data: "hello", Left: &Node{Data: "world"}, Right: &Node{Data: "right"}}}
	t.Display()
}
