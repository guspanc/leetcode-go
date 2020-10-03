package main

func main() {

}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := -1
	counter := 0
	dfs(root, &counter, &res, k)
	return res
}

func dfs(node *TreeNode, counter *int, res *int, k int) {
	if node == nil {
		return
	}
	dfs(node.Left, counter, res, k)
	(*counter)++
	if (*counter) == k {
		(*res) = node.Val
		return
	}
	dfs(node.Right, counter, res, k)
}
