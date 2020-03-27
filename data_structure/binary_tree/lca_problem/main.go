package main

/*
	lastCommentAncestor 树两个节点最近公共祖先问题
*/

// 二叉搜索树的最近公共祖先问题
type BstNode struct {
	Data  int
	Left  *BstNode
	Right *BstNode
}
type Bst struct {
	Root *BstNode
}

func (t *Bst) BstLca(p, q int) (node *BstNode) {
	if t.Root == nil {
		return t.Root
	}
	return t.Root.BstLca(p, q)
}
func (n *BstNode) BstLca(p, q int) (node *BstNode) { //认为是 p < q 的
	if p > q {
		return n.BstLca(q, p)
	}
	if n.Data >= p && n.Data <= q {
		return n
	}
	if n.Data > q {
		return n.Left.BstLca(p, q)
	}
	if n.Data < p {
		return n.Right.BstLca(p, q)
	}
	return nil
}

// 普通二叉树的最近公共祖先问题 # TODO: To be continue

type BtNode struct {
	Data  int
	Left  *BtNode
	Right *BtNode
}
type Bt struct {
	Root *BtNode
}

func (b *Bt) GetDepth(node *BstNode) int {

}
func (b *Bt) BtLca(p, q int) (node *BtNode) {

}
