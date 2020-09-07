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

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	pathNums := []int{}
	used := make([]bool, len(nums))
	backTrack(nums, pathNums, used)
	return res
}

func backTrack(nums, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(pathNums))
		copy(tmp, pathNums)
		res = append(res, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backTrack(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}
