package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  基于单链表实现的 skiplist
*/

const MaxLevel = 16

var Flag = false

type Node struct {
	NextNodes map[int8]*Node
	data      int
	MaxLevel  int8
}

type SkipList struct {
	MaxLevel int8
	Head     *Node
}

func RandomLevel() int8 {
	if Flag == false {
		rand.Seed(time.Now().UnixNano())
		Flag = true
	}
	return int8(rand.Intn(MaxLevel) + 1)
}

// 符合正态分布的随机算法
func BinaryRandom() int8 {
	k := 1
	if Flag == false {
		rand.Seed(time.Now().UnixNano())
		Flag = true
	}
	for rand.Intn(2) == 1 {
		if k <= MaxLevel {
			k += 1
		} else {
			break
		}
	}
	return int8(k)
}

func NewSkipList() *SkipList {

	return &SkipList{
		MaxLevel: 0,
		Head:     &Node{NextNodes: map[int8]*Node{}, data: 0, MaxLevel: 0},
	}
}
func (s *SkipList) Insert(data int) {
	level := RandomLevel()
	fmt.Printf("insert data : %d , the level is: %d \n", data, level)
	if s.MaxLevel < level {
		s.MaxLevel = level
	}
	newNode := &Node{NextNodes: map[int8]*Node{}, data: data, MaxLevel: level}
	currentPrevNode := s.Head
	for l := level; l > 0; l-- { // (当前层l)
		tempNode := currentPrevNode
		for tempNode.data < newNode.data { // 不断向后遍历查找插入区间
			tempNextNode := tempNode.NextNodes[l]
			if tempNextNode == nil || tempNextNode.data > newNode.data { // 找到了插入的区间
				tempNode.NextNodes[l] = newNode
				newNode.NextNodes[l] = tempNextNode
				break
			}
			tempNode = tempNextNode
		}
		currentPrevNode = tempNode //逐个节点后递进
	}
}
func (s *SkipList) Display() {
	currNode := s.Head.NextNodes[1]
	for currNode.NextNodes[1] != nil {
		fmt.Println(currNode.data)
		currNode = currNode.NextNodes[1]
	}
	fmt.Println(currNode.data)
}

func main() {
	sl := NewSkipList()
	sl.Insert(4)
	sl.Insert(23)
	sl.Insert(44)
	sl.Insert(43)
	sl.Insert(276)
	sl.Insert(78)
	sl.Insert(3)
	sl.Insert(2)
	sl.Insert(8)
	sl.Insert(21)
	sl.Insert(1)
	sl.Display()
}
