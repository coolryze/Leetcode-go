package binarytreecameras

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	res := 0

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 2
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		if left == 0 || right == 0 {
			res++
			return 1
		} else if left == 1 || right == 1 {
			return 2
		} else if left == 2 && right == 2 {
			return 0
		}

		return -1
	}

	if dfs(root) == 0 {
		res++
	}

	return res
}
