package medium

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil && root.Next.Left != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
