package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var idx, maxValue = 0, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			idx = i
		}
	}
	return &TreeNode{
		Val:   nums[idx],
		Left:  constructMaximumBinaryTree(nums[0:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}
}

func build(nums []int, lIdx, rIdx int) *TreeNode {
	if lIdx >= rIdx {
		return nil
	}
	var idx, maxValue = 0, -1
	for i := lIdx; i < rIdx; i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			idx = i
		}
	}
	return &TreeNode{
		Val:   nums[idx],
		Left:  build(nums, lIdx, idx),
		Right: build(nums, idx+1, rIdx),
	}
}
