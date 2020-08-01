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

//剑指 Offer 12. 矩阵中的路径
//https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
//func exist(board [][]byte, word string) bool {
//
//}

//剑指 Offer 16. 数值的整数次方
//https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	result := myPow(x, n>>1)
	if n&1 > 0 {
		return result * result * x
	}
	return result * result
}
