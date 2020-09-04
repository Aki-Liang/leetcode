package leetcode

// 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。

// 例如，给定一个 3叉树 :
// 返回其层序遍历:

// [
//      [1],
//      [3,2,4],
//      [5,6]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if nil == root {
		return nil
	}
	res := [][]int{}
	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		level := []int{}
		for i := 0; i < l; i++ {

			node := queue[i]
			level = append(level, node.Val)
			for _, n := range node.Children {
				queue = append(queue, n)
			}
		}
		queue = queue[l:]
		res = append(res, level)
	}

	return res
}

// func levelOrder(root *Node) [][]int {
// 	if nil == root {
// 		return nil
// 	}

// 	res := [][]int{}
// 	queue := list.New()
// 	queue.PushBack(root)
// 	for queue.Len() > 0 {
// 		items := []int{}
// 		l := queue.Len()
// 		for i := 0; i < l; i++ {
// 			elem := queue.Front()
// 			node := elem.Value.(*Node)
// 			items = append(items, node.Val)
// 			queue.Remove(elem)
// 			for _, item := range node.Children {
// 				queue.PushBack(item)
// 			}
// 		}
// 		res = append(res, items)
// 	}

// 	return res
// }

// func levelOrder(root *Node) [][]int {
// 	if nil == root {
// 		return nil
// 	}

// 	res := [][]int{}
// 	queue := list.New()
// 	queue.PushBack(root)
// 	for queue.Len() > 0 {
// 		items := []int{}
// 		for i := 0; i < queue.Len(); i++ {
// 			elem := queue.Front()
// 			node := elem.Value.(*Node)
// 			items = append(items, node.Val)
// 			queue.Remove(elem)
// 			for _, n := range node.Children {
// 				queue.PushBack(n)
// 			}
// 		}
// 		res = append(res, items)
// 	}
// 	return res
// }
