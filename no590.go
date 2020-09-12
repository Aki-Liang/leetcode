package leetcode

// 给定一个 N 叉树，返回其节点值的后序遍历。

// 例如，给定一个 3叉树 :

// 返回其后序遍历: [5,6,3,2,4,1].

// 说明: 递归法很简单，你可以使用迭代法完成此题吗?

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	res := []int{}
	if nil == root {
		return res
	}

	for _, node := range root.Children {
		ret := postorder(node)
		res = append(res, ret...)
	}
	res = append(res, root.Val)
	return res
}
