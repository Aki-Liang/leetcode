package leetcode

import (
	"fmt"
	"testing"
)

func Test_tribonacci(t *testing.T) {
	res := tribonacci(4)
	if 4 != res {
		panic(res)
	}
}

func Test_findUnsortedSubarray(t *testing.T) {
	res := findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
	//res := findUnsortedSubarray([]int{1, 2, 3, 4})
	if 5 != res {
		panic(res)
	}
}

func Test_plusOne(t *testing.T) {
	res := plusOne([]int{9, 9, 9, 9})
	fmt.Println(res)
}

func Test_twoSum(t *testing.T) {
	res := twoSum([]int{3, 3}, 6)
	fmt.Println(res)
}

func Test_threeSum(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
func Test_rangeBitwiseAnd(t *testing.T) {
	res := rangeBitwiseAnd(5, 7)
	fmt.Println(res)
}

func Test_climbStairs(t *testing.T) {
	res := climbStairs(3)
	fmt.Println(res)
}

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

func Test_LRUCache(t *testing.T) {
	// cache := Constructor(2)

	// cache.Put(1, 1)
	// cache.Put(2, 2)
	// cache.Get(1)    // 返回  1
	// cache.Put(3, 3) // 该操作会使得关键字 2 作废
	// cache.Get(2)    // 返回 -1 (未找到)
	// cache.Put(4, 4) // 该操作会使得关键字 1 作废
	// cache.Get(1)    // 返回 -1 (未找到)
	// cache.Get(3)    // 返回  3
	// cache.Get(4)    // 返回  4

}

func Test_removeDuplicates(t *testing.T) {
	res := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(res)
}

func Test_merge(t *testing.T) {
	// nums1 = [1,2,3,0,0,0], m = 3
	// nums2 = [2,5,6],       n = 3
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func Test_maxArea(t *testing.T) {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}

func Test_isValid(t *testing.T) {
	res := isValid("}")
	fmt.Println(res)
}

func Test_swapPairs(t *testing.T) {
	//1->2->3->4
	Node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	res := swapPairs(Node)
	fmt.Println(res)
}

func Test_reverseKGroup(t *testing.T) {
	//1->2->3->4
	Node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	res := reverseKGroup(Node, 2)
	fmt.Println(res)
}

func Test_rotate(t *testing.T) {
	rotate([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3)

}

func Test_isUnique(t *testing.T) {
	fmt.Println(isUnique("abc"))
}

func Test_removeOuterParentheses(t *testing.T) {
	removeOuterParentheses("(()())(())(()(()))")
}
func Test_combine(t *testing.T) {
	fmt.Println(combine(4, 2))
}

func Test_search(t *testing.T) {
	fmt.Println(search([]int{3, 1}, 1))
}

func Test_findLadders(t *testing.T) {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func Test_combinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{2, 2, 2}, 2))
}

func Test_4sum(t *testing.T) {
	fmt.Println(fourSum([]int{5, 5, 3, 5, 1, -5, 1, -2}, 4))
}

func Test_longestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("(()"))
}

func Test_minPathSum(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func Test_numDecodings(t *testing.T) {
	fmt.Println(numDecodings("1210121"))
}

func Test_maxCoins(t *testing.T) {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}

func Test_leastInterval(t *testing.T) {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
