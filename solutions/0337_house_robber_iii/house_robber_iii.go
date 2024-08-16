package houserobberiii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(cur *TreeNode) []int
	dfs = func(cur *TreeNode) []int {
		if cur == nil {
			return []int{0, 0}
		}

		left := dfs(cur.Left)
		right := dfs(cur.Right)

		robCur := cur.Val + left[0] + right[0]
		notRobCur := max(left[0], left[1]) + max(right[0], right[1])
		return []int{notRobCur, robCur}
	}

	res := dfs(root)
	return max(res[0], res[1])
}
