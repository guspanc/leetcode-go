package main

import "sort"

func main() {

}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := []int{}
	inOrder(root1, &res)
	inOrder(root2, &res)
	sort.Ints(res)
	return res
}

func inOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, res)
	*res = append(*res, node.Val)
	inOrder(node.Right, res)
}
