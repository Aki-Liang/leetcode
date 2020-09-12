package leetcode

// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

// 示例 1：

// 输入：head = [1,3,2]
// 输出：[2,3,1]

// 限制：

// 0 <= 链表长度 <= 10000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//递归
// func reversePrint(head *ListNode) []int {
// 	if nil == head {
// 		return []int{}
// 	}
// 	return append(reversePrint(head.Next), head.Val)
// }

//循环
func reversePrint(head *ListNode) []int {
	res := []int{}
	if nil == head {
		return res
	}

	for nil != head {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
