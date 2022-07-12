package main

import (
	"fmt"
	"sort"
)

type initSlice []int

func (p initSlice) Len() int {
	return len(p)
}

func (p initSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
// 交换
func (p initSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	arr := []int{1, 2, 6, 4, 5}

	// 判断是否是排序的
	flag := sort.IntsAreSorted(arr)
	fmt.Println(flag)

	// 正序排列
	sort.Ints(arr)
	fmt.Println(arr)

	// 降序排列
	dd := make(initSlice, 0)
	dd = append(dd, 1, 3, 4, 5, 6, 1)
	sort.Sort(sort.Reverse(dd))
	fmt.Println(dd)
}
