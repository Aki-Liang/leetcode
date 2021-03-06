package leetcode

import "container/list"

// 给定一个二叉树，返回它的中序 遍历。

// 示例:

// 输入: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3

// 输出: [1,3,2]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func inorderTraversal(root *TreeNode) []int {
// 	if nil == root {
// 		return nil
// 	}

// 	res := []int{}
// 	res = append(res, inorderTraversal(root.Left)...)
// 	res = append(res, root.Val)
// 	res = append(res, inorderTraversal(root.Right)...)
// 	return res
// }

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := list.New()
	for stack.Len() > 0 || nil != root {
		for nil != root {
			stack.PushBack(root)
			root = root.Left
		}
		elem := stack.Back()
		node := elem.Value.(*TreeNode)
		res = append(res, node.Val)
		stack.Remove(elem)
		root = node.Right
	}

	return res
}
