package main

func main() {

}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	result := 0
	getSum(root, &result, L, R)
	return result
}

func getSum(n *TreeNode, sum *int, l int, r int) {
	if n == nil {
		return
	}
	if n.Val >= l && n.Val <= r {
		(*sum) += n.Val
	}
	getSum(n.Left, sum, l, r)
	getSum(n.Right, sum, l, r)
}
