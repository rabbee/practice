package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return
	}
	//找到左节点的叶子节点
	if root.Left == nil {
		return
	}
	iter := root.Left
	for iter.Right != nil {
		iter = iter.Right
	}
	iter.Right, root.Right, root.Left = root.Right, root.Left, nil
}
