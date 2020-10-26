package leetcode

// 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。

// 你需要返回给定数组中的重要翻转对的数量。

// 示例 1:

// 输入: [1,3,2,3,1]
// 输出: 2
// 示例 2:

// 输入: [2,4,3,5,1]
// 输出: 3
// 注意:

// 给定数组的长度不会超过50000。
// 输入数组中的所有数字都在32位整数的表示范围内。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-pairs
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// func reversePairs(nums []int) int {
// 	return mergeSortAndCount(nums, 0, len(nums)-1)
// }

// func mergeSortAndCount(nums []int, start, end int) {
// 	if start < end {
// 		mid := (start + end) / 2
// 		count := mergeSortAndCount(nums, start, mid) + mergeSortAndCount(nums, mid+1, end)
// 		j := mid + 1
// 		for i := start; i <= mid; i++ {

// 		}
// 	}
// 	return 0
// }
