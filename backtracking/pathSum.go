package backtracking

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) [][]int {
	var (
		res       [][]int
		backtrack func(pathSum int, path []int, root *TreeNode, flag bool)
	)

	backtrack = func(pathSum int, path []int, root *TreeNode, flag bool) {
		if pathSum == targetSum && root == nil && flag {
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}

		for root != nil {
			pathSum += root.Val
			path = append(path, root.Val)
			flag = true
			if root.Right != nil {
				flag = false
			}
			backtrack(pathSum, path, root.Left, flag)
			root = root.Right
		}

	}

	if root == nil {
		return [][]int{}
	}

	backtrack(0, []int{}, root, true)

	return res
}
