package leetcode

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

// 示例：
// 二叉树：[3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：
// [
//   [3],
//   [9,20],
//   [15,7]
// ]
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if nil == root {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		layer := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			layer = append(layer, queue[i].Val)
			if nil != queue[i].Left {
				queue = append(queue, queue[i].Left)
			}
			if nil != queue[i].Right {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, layer)
		queue = queue[length:]
	}
	return res
}
