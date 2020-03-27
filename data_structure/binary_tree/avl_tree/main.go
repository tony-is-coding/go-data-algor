package main

/* AVL tree 以人名取名的的一种 自平衡二叉查找树, 符合以下约束：  # TODO: To be continue
	1. 左右子树的高度差小于等于1
	2. 其每一个子树均为平衡二叉树
  # 监督机制： 为了保证二叉树的平衡， AVL 树引入了所谓监督机制，就是在树的某一部分的不平衡度超过一个阈值后触发相应的平衡操作。
			  保证树的平衡度在可以接受的范围内 (基于平衡因子 balance factor)
	平衡因子： 某个结点的左子树的高度减去右子树的高度得到的差值。 |left - right|
	AVL 树： 所有结点的平衡因子的绝对值都不超过 1 的二叉树
*/

type Node struct {
	data   int
	height int
	left   *Node
	right  *Node
	parent *Node
}
type Tree struct {
	root *Node
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func (n *Node) nodeHeight() int {
	if n == nil {
		return 0
	}
	var left, right int
	if n.left != nil {
		left = n.left.nodeHeight()
	}
	if n.right != nil {
		right = n.right.nodeHeight()
	}
	return max(left, right) + 1
}
func (n *Node) nodeBalanceFactor() int {
	if n == nil {
		return 0
	}
	var leftHeight, rightHeight int
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}
	return abs(leftHeight - rightHeight)
}

func (n *Node) leftRotation() {

}
