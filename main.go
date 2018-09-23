package main

import (
	"fmt"

	"github.com/rabbee/leetcode/contest/c103"
)

func main() {
	var a []int
	var res int

	a = []int{1}
	res = c103.SmallestRangeI(a, 0)
	fmt.Println(res)

	a = []int{0, 10}
	res = c103.SmallestRangeI(a, 2)
	fmt.Println(res)

	a = []int{1, 3, 6}
	res = c103.SmallestRangeI(a, 3)
	fmt.Println(res)

	a = []int{2, 7, 2}
	res = c103.SmallestRangeI(a, 1)
	fmt.Println(res)

	a = []int{4, 5, 6}
	res = c103.SmallestRangeI(a, 2)
	fmt.Println(res)

	a = []int{7, 8, 8}
	res = c103.SmallestRangeI(a, 5)
	fmt.Println(res)

	a = []int{3, 1, 10}
	res = c103.SmallestRangeI(a, 4)
	fmt.Println(res)
}
