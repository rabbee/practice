package coding_interviews

import (
	"strings"
)

//findRepeatNumber 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}
		m[nums[i]] = true
	}
	return nums[0]
}

//findNumberIn2DArray 04. 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var i = 0
	var j = len(matrix[i]) - 1
	for i < len(matrix) && j >= 0 {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			i++
			continue
		}
		if target < matrix[i][j] {
			j--
			continue
		}
	}
	return false
}

//replaceSpace 05. 替换空格
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//剑指 Offer 06. 从尾到头打印链表
//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	result := []int{}
	for p := head; p != nil; p = p.Next {
		result = append(result, p.Val)
	}
	l := len(result)
	for i := 0; i < l/2; i++ {
		result[i], result[l-i-1] = result[l-i-1], result[i]
	}
	return result
}

//剑指 Offer 09. 用两个栈实现队列
//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
//type CQueue struct {
//
//}
//
//
//func Constructor() CQueue {
//
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//
//}
//
//
//func (this *CQueue) DeleteHead() int {
//
//}
//end*******剑指 Offer 09. 用两个栈实现队列

//剑指 Offer 10- I. 斐波那契数列
//https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
func powMatrix(n int) [2][2]int {
	if n == 1 {
		return [2][2]int{
			{1, 1},
			{1, 0},
		}
	}
	if n&1 == 1 {
		tmp := powMatrix((n - 1) / 2)
		return multiMatrix(multiMatrix(tmp, tmp), powMatrix(1))
	} else {
		tmp := powMatrix(n / 2)
		return multiMatrix(tmp, tmp)
	}
}

func multiMatrix(a, b [2][2]int) [2][2]int {
	return [2][2]int{
		{(a[0][0]*b[0][0] + a[0][1]*b[1][0]) % (1e9 + 7), (a[0][0]*b[0][1] + a[0][1]*b[1][1]) % (1e9 + 7)},
		{(a[1][0]*b[0][0] + a[1][1]*b[1][0]) % (1e9 + 7), (a[1][0]*b[0][1] + a[1][1]*b[1][1]) % (1e9 + 7)},
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return powMatrix(n - 1)[0][0]
}
//end****************剑指 Offer 10- I. 斐波那契数列

//剑指 Offer 11. 旋转数组的最小数字
//https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
//TODO: 当前平均时间o(n/2)，用二分时间可以降到o(log n)
func minArray(numbers []int) int {
	l := len(numbers)
	if l == 0 {
		return 0
	}
	for i := 0; i < l-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}
