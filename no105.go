package leetcode

// 根据一棵树的前序遍历与中序遍历构造二叉树。

// 注意:
// 你可以假设树中没有重复的元素。

// 例如，给出

// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//找到中序的根节点
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	left := buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	right := buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
	return root
}

// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	if len(preorder) != len(inorder) || len(preorder) == 0 {
// 		return nil
// 	}

// 	root := &TreeNode{
// 		Val: preorder[0],
// 	}
// 	//找到中序遍历的跟节点
// 	i := 0
// 	for ; i < len(inorder); i++ {
// 		if inorder[i] == preorder[0] {
// 			break
// 		}
// 	}
// 	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
// 	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])

// 	return root
// }
