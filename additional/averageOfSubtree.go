package additional

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfSubtree(root *TreeNode) int {
	_, _, ans := helper(root)
	return ans
}
     
func helper(node *TreeNode) (int, int, int) {
	num, sum, ans := 1, node.Val, 0
	if node.Left != nil {
		lnum, lsum, lans := helper(node.Left)
		num += lnum
		sum += lsum
		ans += lans
	}
	if node.Right != nil {
		rnum, rsum, rans := helper(node.Right)
		num += rnum
		sum += rsum
		ans += rans
	}
	if sum/num == node.Val {
		ans += 1
	}
	return num, sum, ans
}
