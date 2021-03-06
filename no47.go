package leetcode

import "sort"

// 给定一个可包含重复数字的序列，返回所有不重复的全排列。

// 示例:

// 输入: [1,1,2]
// 输出:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutations-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	backTrackUnique(nums, 0, &res)
	return res
}

func backTrackUnique(nums []int, i int, res *[][]int) {
	if len(nums)-1 == i {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	}

	for j := i; j < len(nums); j++ {
		if i != j && nums[j] == nums[i] {
			continue
		}
		nums[j], nums[i] = nums[i], nums[j]
		backTrackUnique(nums, i+1, res)
	}

	for j := len(nums) - 1; j > i; j-- {
		nums[j], nums[i] = nums[i], nums[j]
	}

}

// func permuteUnique(nums []int) [][]int {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	sort.Ints(nums)
// 	res := [][]int{}
// 	backTrackUnique(nums, 0, &res)
// 	return res
// }

// // 回溯函数实现
// // i表示本次函数需要放置的元素位置
// func backTrackUnique(nums []int, i int, res *[][]int) {
// 	l := len(nums)
// 	if l-1 == i {
// 		tmp := make([]int, len(nums))
// 		copy(tmp, nums)
// 		*res = append(*res, tmp)
// 		return
// 	}

// 	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
// 	for j := i; j < l; j++ {
// 		//跳过相同的值
// 		if i != j && nums[i] == nums[j] {
// 			continue
// 		}

// 		nums[i], nums[j] = nums[j], nums[i]
// 		backTrackUnique(nums, i+1, res)
// 	}

// 	//还原调整过的数组
// 	for j := l - 1; j > i; j-- {
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// }
