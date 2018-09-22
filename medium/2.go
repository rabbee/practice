/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 * 给定两个非空链表来表示两个非负整数。
 * 位数按照逆序方式存储，它们的每个节点只存储单个数字。
 * 将两数相加返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 */

package medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	res := root
	current1 := l1
	current2 := l2
	var up = false
	for {
		sum, bit1, bit2 := 0, 0, 0
		current := &ListNode{}
		if current1 == nil && current2 == nil {
			if up {
				current.Val = 1
				root.Next = current
				root = root.Next
			}
			break
		}
		if current1 != nil {
			bit1 = current1.Val
			current1 = current1.Next
		}
		if current2 != nil {
			bit2 = current2.Val
			current2 = current2.Next
		}
		fmt.Printf("%d %d\n", bit1, bit2)
		if up {
			sum = bit1 + bit2 + 1
		} else {
			sum = bit1 + bit2
		}
		fmt.Println(sum)
		if sum >= 10 {
			up = true
			current.Val = sum % 10
		} else {
			up = false
			current.Val = sum
		}
		root.Next = current
		root = root.Next
		fmt.Println(current.Val)
	}
	return res.Next
}
