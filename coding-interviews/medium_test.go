package coding_interviews

import (
	"fmt"
	"testing"
)

func preOrder(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d,", node.Val)
		if node.Left != nil {
			preOrder(node.Left)
		}
		if node.Right != nil {
			preOrder(node.Right)
		}
	}
}

func inoOrder(node *TreeNode) {
	if node != nil {
		if node.Left != nil {
			inoOrder(node.Left)
		}
		fmt.Printf("%d,", node.Val)
		if node.Right != nil {
			inoOrder(node.Right)
		}
	}
}

func Test_buildTree(t *testing.T) {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	preOrder(tree)
	fmt.Println()
	inoOrder(tree)
	fmt.Println()
	tree = buildTree([]int{}, []int{})
	preOrder(tree)
	fmt.Println()
	inoOrder(tree)
	fmt.Println()
	tree = buildTree([]int{3, 20, 15, 7, 9}, []int{15, 20, 7, 3, 9})
	preOrder(tree)
	fmt.Println()
	inoOrder(tree)
	fmt.Println()
}

func Test_myPow(t *testing.T) {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}