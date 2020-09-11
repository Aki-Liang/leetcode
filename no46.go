package leetcode

// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

// 示例:

// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutations
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func permute(nums []int) [][]int {
	res := [][]int{}
	backTrackPermute(nums, 0, &res)
	return res
}
func backTrackPermute(nums []int, idx int, res *[][]int) {
	if idx == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	}

	for i := idx; i < len(nums); i++ {
		nums[idx], nums[i] = nums[i], nums[idx]
		backTrackPermute(nums, idx+1, res)
	}

	for i := len(nums) - 1; i > idx; i-- {
		nums[idx], nums[i] = nums[i], nums[idx]
	}
}

// func permute(nums []int) [][]int {
// 	res := [][]int{}
// 	l := len(nums)
// 	used := make([]bool, l)
// 	for i := 0; i < l; i++ {
// 		used[i] = true
// 		backTrackPermute(nums, []int{nums[i]}, used, &res)
// 		used[i] = false
// 	}

// 	return res
// }

// func backTrackPermute(nums []int, sortNums []int, used []bool, res *[][]int) {
// 	if len(nums) == len(sortNums) {
// 		tmp := make([]int, len(sortNums))
// 		copy(tmp, sortNums)
// 		*res = append(*res, tmp)
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if !used[i] {
// 			used[i] = true
// 			sortNums = append(sortNums, nums[i])
// 			backTrackPermute(nums, sortNums, used, res)
// 			sortNums = sortNums[:len(sortNums)-1]
// 			used[i] = false
// 		}
// 	}
// }

// var res [][]int

// func permute(nums []int) [][]int {
// 	res = [][]int{}
// 	pathNums := []int{}
// 	used := make([]bool, len(nums))
// 	backTrack(nums, pathNums, used)
// 	return res
// }

// func backTrack(nums, pathNums []int, used []bool) {
// 	if len(nums) == len(pathNums) {
// 		tmp := make([]int, len(pathNums))
// 		copy(tmp, pathNums)
// 		res = append(res, tmp)
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if !used[i] {
// 			used[i] = true
// 			pathNums = append(pathNums, nums[i])
// 			backTrack(nums, pathNums, used)
// 			pathNums = pathNums[:len(pathNums)-1]
// 			used[i] = false
// 		}
// 	}
// }
