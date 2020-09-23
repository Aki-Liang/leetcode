package leetcode

import "sort"

// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

// 注意：

// 答案中不可以包含重复的四元组。

// 示例：

// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

// 满足要求的四元组集合为：
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/4sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//回溯法 984 ms 不好使
func fourSumTB(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	trackBack4Sum(nums, []int{}, 0, target, &res)
	return res
}

func trackBack4Sum(nums, ans []int, n, target int, res *[][]int) {
	if len(ans) == 4 {
		if target != 0 {
			return
		}
		tmp := make([]int, 4)
		copy(tmp, ans)
		*res = append(*res, tmp)
		return
	}

	for i := n; i < len(nums); i++ {
		rawI := i
		for ; i+1 < len(nums); i++ {
			if nums[i] != nums[i+1] {
				break
			}
		}

		ans = append(ans, nums[i])
		trackBack4Sum(nums, ans, rawI+1, target-nums[i], res)
		ans = ans[:len(ans)-1]
	}
}

//双指针法
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			m := j + 1
			n := len(nums) - 1
			for m < n {
				sum := nums[i] + nums[j] + nums[m] + nums[n]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
					for m < n {
						n--
						if nums[n] != nums[n+1] {
							break
						}
					}
					for m < n {
						m++
						if nums[m] != nums[m-1] {
							break
						}
					}
				} else if sum > target {
					for m < n {
						n--
						if nums[n] != nums[n+1] {
							break
						}
					}
				} else if sum < target {
					for m < n {
						m++
						if nums[m] != nums[m-1] {
							break
						}
					}
				}
			}

			for j < len(nums)-2 {
				j++
				if nums[j] != nums[j-1] {
					break
				}
			}
		}

		for i < len(nums)-3 {
			i++
			if nums[i] != nums[i-1] {
				break
			}
		}
	}

	return res
}
