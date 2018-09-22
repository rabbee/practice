package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	l1 := &medium.ListNode{2, &medium.ListNode{4, nil}}
	l2 := &medium.ListNode{5, &medium.ListNode{6, nil}}
	res := medium.AddTwoNumbers(l1, l2)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}
