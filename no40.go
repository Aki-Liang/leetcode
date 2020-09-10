package leetcode

import "sort"

// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明：
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	trackBackCombinationSum2(candidates, []int{}, 0, 0, target, &res)
	return res
}

func trackBackCombinationSum2(candidates, nums []int, index, curr, target int, res *[][]int) {
	if curr == target {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	l := len(candidates)
	for i := index; i < l; i++ {
		// 去除多余的重复项（调整当前层的选择进度即可，下层需要继续可选）
		rawI := i
		for i < l-1 && candidates[i] == candidates[i+1] {
			i++
		}

		curr += candidates[i]
		if curr > target {
			return
		}
		nums = append(nums, candidates[i])
		trackBackCombinationSum2(candidates, nums, rawI+1, curr, target, res)
		curr -= candidates[i]
		nums = nums[:len(nums)-1]
	}
}
