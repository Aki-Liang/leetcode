package leetcode

import "math"

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 假设一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/validate-binary-search-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}

	return checkValidBST(root, math.MinInt64, math.MaxInt64)
}

func checkValidBST(root *TreeNode, min, max int) bool {
	if nil == root {
		return true
	}

	if min >= root.Val || max <= root.Val {
		return false
	}

	return checkValidBST(root.Left, min, root.Val) && checkValidBST(root.Right, root.Val, max)
}
