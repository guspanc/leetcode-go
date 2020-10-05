package main

func main() {

}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	var sum int
	dfs(root, nil, nil, &sum)
	return sum
}

func dfs(node *TreeNode, parent *TreeNode, grandParent *TreeNode, sum *int) {
	if node == nil {
		return
	}
	if grandParent != nil && grandParent.Val%2 == 0 {
		(*sum) += node.Val
	}
	dfs(node.Left, node, parent, sum)
	dfs(node.Right, node, parent, sum)
}
