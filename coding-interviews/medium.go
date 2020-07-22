package coding_interviews

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//剑指 Offer 07. 重建二叉树
//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
func buildTree(preorder []int, inorder []int) *TreeNode {
	nodeNum := len(preorder)
	if nodeNum == 0 {
		return nil
	}
	if nodeNum == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}
	for i := 0; i < nodeNum; i++ {
		if preorder[0] == inorder[i] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1], inorder[0:i]),
				Right: buildTree(preorder[i+1:nodeNum], inorder[i+1:nodeNum]),
			}
		}
	}
	return nil
}
