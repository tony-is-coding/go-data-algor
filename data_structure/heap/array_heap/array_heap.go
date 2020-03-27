package main

import (
	"fmt"
)

/*
	基于纯数组实现的heap, 默认为当前index + 1 为左子树
		父子节点的下标关系
			1. 未使用index = 0元素的数组的父子节点的下标关系，即图一：
			k=父节点的index -> 左子节点的index = 2k， 右子节点的index = 2k + 1
			j = 子节点的index -> 父节点的index = j / 2

			2. 使用index = 0元素的数组的父子节点的下标关系，即图二：
			k=父节点的index -> 左子节点的index = 2k + 1， 右子节点的index = （k + 1） x 2
			j = 子节点的index -> 父节点的index = （j -1） / 2
*/

type Node int
type Heap struct {
	size int
	data []Node // 使用 0 索引元素
}

// 目前假设最多容纳100个元素
func heapNew() *Heap {
	return &Heap{0, make([]Node, 0, 100)}
}

// 返回对应左孩子的index
func (h *Heap) LeftChild(i int) int {
	i = 2*i + 1
	return i
}

// 返回对应右孩子的index
func (h *Heap) RightChild(i int) int {
	i = 2 * (i + 1)
	return i
}

// 可能不存在子节点
func (h *Heap) Parent(i int) int {
	i = (i - 1) / 2
	return i
}

// 堆元素上浮
// 每次写入元素, 都默认写入到最后一个节点, 然后再做上浮调整
func (h *Heap) heapUp(curr int) {
	if curr <= 0 {
		return
	}
	parent := h.Parent(curr)
	if h.data[curr] < h.data[parent] {
		h.data[parent], h.data[curr] = h.data[curr], h.data[parent]
	}
	h.heapUp(parent)
}

// 堆元素下沉
// 每次删除元素, 将
func (h *Heap) heapDown(curr int) {
	left := h.LeftChild(curr)
	right := h.RightChild(curr)
	if left >= h.size {
		return
	}
	minChild := left
	if right < h.size && h.data[left] > h.data[right] {
		minChild = right
	}
	if h.data[curr] > h.data[minChild] { // 满足当前节点大于最小的子节点，进行交换
		h.data[curr], h.data[minChild] = h.data[minChild], h.data[curr]
		h.heapDown(minChild)
	} else {
		return
	}
}

func (h *Heap) Insert(data int) {
	var n Node = Node(data)
	h.data = append(h.data, n)
	h.size += 1
	h.heapUp(h.size - 1)
}

func (h *Heap) Delete() Node {
	popData := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data = h.data[:h.size-1]
	h.size = h.size - 1
	h.heapDown(0)
	return popData
}

func main() {
	heap := heapNew()
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(1)
	heap.Insert(9)
	heap.Insert(16)
	fmt.Println(heap.data)
	size := heap.size
	for i := 0; i < size; i++ {
		fmt.Println(heap.Delete())
	}
}
