package leetcode

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

//

// 示例:

// 给定 1->2->3->4, 你应该返回 2->1->4->3.

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func swapPairs(head *ListNode) *ListNode {
// 	if nil == head || nil == head.Next {
// 		return head
// 	}

// 	first := head
// 	second := head.Next
// 	first.Next = second.Next
// 	second.Next = first
// 	first.Next = swapPairs(first.Next)
// 	return second

// }

//递归解法
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	first := head
// 	second := head.Next

// 	first.Next = swapPairs(second.Next)
// 	second.Next = first
// 	return second
// }

// 循环解法

// func swapPairs(head *ListNode) *ListNode {
// 	pre := &ListNode{
// 		Next: head,
// 	}
// 	tmp := pre
// 	for nil != tmp.Next && nil != tmp.Next.Next {
// 		first := tmp.Next
// 		second := tmp.Next.Next
// 		tmp.Next = second
// 		first.Next = second.Next
// 		second.Next = first
// 		tmp = first
// 	}
// 	return pre.Next
// }
