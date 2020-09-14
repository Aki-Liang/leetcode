package leetcode

// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

// k 是一个正整数，它的值小于或等于链表的长度。

// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

//
// 示例：

// 给你这个链表：1->2->3->4->5

// 当 k = 2 时，应当返回: 2->1->4->3->5

// 当 k = 3 时，应当返回: 3->2->1->4->5

// 说明：

// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0
	//找到k+1个节点
	for curr != nil && count != k {
		curr = curr.Next
		count++
		if count == k {
			break
		}
	}
	if count == k { //有k+1个节点
		curr = reverseKGroup(curr, k) //递归换后面的
		for ; count > 0; count-- {
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
		}
		head = curr
	}

	return head

}

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	curr := head
// 	count := 0
// 	for curr != nil && count != k { // find the k+1 node
// 		curr = curr.Next
// 		count++
// 	}
// 	if count == k { // if k+1 node is found
// 		curr = reverseKGroup(curr, k) // reverse list with k+1 node as head
// 		// head - head-pointer to direct part,
// 		// curr - head-pointer to reversed part;
// 		for ; count > 0; count-- { // reverse current k-group:
// 			tmp := head.Next // tmp - next head in direct part
// 			head.Next = curr // preappcurring "direct" head to the reversed list
// 			curr = head      // move head of reversed part to a new node
// 			head = tmp       // move "direct" head to the next node in direct part
// 		}
// 		head = curr
// 	}
// 	return head
// }
