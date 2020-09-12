package leetcode

// 给定一个链表，判断链表中是否有环。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if nil == head {
		return false
	}
	walker := head
	runner := head
	for nil != runner.Next && nil != runner.Next.Next {
		walker = walker.Next
		runner = runner.Next.Next
		if walker == runner {
			return true
		}
	}
	return false
}

// func hasCycle(head *ListNode) bool {
// 	if nil == head || nil == head.Next {
// 		return false
// 	}
// 	pFast := head.Next
// 	pSlow := head
// 	for {
// 		if pFast == pSlow {
// 			return true
// 		}

// 		if nil == pSlow.Next {
// 			return false
// 		}
// 		pSlow = pSlow.Next

// 		if nil == pFast.Next || nil == pFast.Next.Next {
// 			return false
// 		}
// 		pFast = pFast.Next.Next
// 	}

// }
