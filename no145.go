package leetcode

// 给定一个二叉树，返回它的 后序 遍历。

// 示例:

// 输入: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// 输出: [3,2,1]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if nil == root {
		return res
	}

	l, r := root.Left, root.Right
	root.Left, root.Right = nil, nil
	stack := []*TreeNode{root}
	if nil != r {
		stack = append(stack, r)
	}
	if nil != l {
		stack = append(stack, l)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if nil == node.Left && nil == node.Right {
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
		}
		l, r = node.Left, node.Right
		node.Left, node.Right = nil, nil
		if nil != r {
			stack = append(stack, r)
		}
		if nil != l {
			stack = append(stack, l)
		}
	}

	return res
}
