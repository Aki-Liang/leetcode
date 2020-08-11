package leetcode

import "container/list"

// 给定一个 N 叉树，返回其节点值的前序遍历。

// 例如，给定一个 3叉树

// 返回其前序遍历: [1,3,5,6,2,4]。

//

// 说明: 递归法很简单，你可以使用迭代法完成此题吗?

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

//暴力递归
func preorder(root *Node) []int {
	if nil == root {
		return nil
	}
	res := []int{root.Val}
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}
	return res
}

//栈
func preorder_stack(root *Node) []int {
	if nil == root {
		return nil
	}
	res := []int{}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		elem := stack.Back()
		node := elem.Value.(*Node)
		res = append(res, node.Val)
		stack.Remove(elem)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
	}
	return res
}
