package main

import "fmt"

/* # TODO: To be continue
	树状索引树/树状数组(Binary Indexed Tree/Fenwick Tree), 是一种用于高效处理对一个存储数字的列表进行更新及求前缀和的数据结构：
		1. update(idx, delta)：将num加到位置idx的数字上。
		2. prefixSum(idx)：求从数组第一个位置到第idx（含idx）个位置所有数字的和。
		3. rangeSum(from_idx, to_idx)：求从数组第from_idx个位置到第to_idx个位置的所有数字的和
*/

type FenwickTree struct {
	Size int
	FT   []int
}

func (f *FenwickTree) Update(index int, value int) {

}

func LowBit(input uint) (output uint) {
	return input & (^input + 1)
}

func main() {
	var a uint
	a = 34
	fmt.Println(LowBit(a))
}
