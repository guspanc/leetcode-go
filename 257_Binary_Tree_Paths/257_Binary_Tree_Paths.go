package main

import "fmt"

func main() {

}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	dfs("", &paths, root)
	return paths
}

func dfs(path string, paths *[]string, node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		path = fmt.Sprintf("%s%d", path, node.Val)
		(*paths) = append((*paths), path)
		return
	}
	path = fmt.Sprintf("%s%d->", path, node.Val)
	dfs(path, paths, node.Left)
	dfs(path, paths, node.Right)
}
