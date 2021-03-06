package main

func main() {

}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= (*min) {
		return false
	}
	if max != nil && node.Val >= (*max) {
		return false
	}
	return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
}
