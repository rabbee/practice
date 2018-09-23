package c103

import (
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func SmallestRangeI(A []int, K int) int {
	sort.Sort(IntSlice(A))
	res := A[len(A)-1] - A[0]
	left := A[0] + K
	right := A[len(A)-1] - K
	for i := 0; i < len(A)-1; i++ {
		maxi := Max(A[i]+K, right)
		mini := Min(A[i+1]-K, left)
		res = Min(res, maxi-mini)
	}
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
