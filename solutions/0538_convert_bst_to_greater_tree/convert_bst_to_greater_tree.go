package convertbsttogreatertree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	prev := 0

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		prev += node.Val
		node.Val = prev
		dfs(node.Left)
	}

	dfs(root)

	return root
}
