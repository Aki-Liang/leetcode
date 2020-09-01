package leetcode

import "sort"

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

//

// 示例：

// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		m := i + 1
		n := l - 1
		for m < n {
			s := nums[i] + nums[m] + nums[n]
			if s > 0 {
				n--
			} else if s < 0 {
				m++
			} else {
				res = append(res, []int{nums[i], nums[m], nums[n]})
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
			}
		}
	}

	return res
}

// func threeSum(nums []int) [][]int {
// 	//使用夹逼法
// 	length := len(nums)
// 	if 3 > length {
// 		return nil
// 	}
// 	res := [][]int{}
// 	sort.Ints(nums)
// 	for i := 0; i < length; i++ {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			//新元素和上一个元素相等时跳过
// 			continue
// 		}
// 		m := i + 1
// 		n := length - 1
// 		for m < n {
// 			sum := nums[i] + nums[m] + nums[n]

// 			if sum > 0 {
// 				n = n - 1
// 				continue
// 			}

// 			if sum < 0 {
// 				m = m + 1
// 				continue
// 			}
// 			if sum == 0 {
// 				res = append(res, []int{nums[i], nums[m], nums[n]})
// 				for n > m {
// 					n -= 1
// 					//尾部指针移动一步，如果相同则再移动一步
// 					if nums[n] != nums[n+1] {
// 						break
// 					}
// 				}

// 				for n > m {
// 					m += 1
// 					//头部指针移动一步，如果相同则再移动一步
// 					if nums[m] != nums[m-1] {
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return res
// }

// func threeSum(nums []int) [][]int {
// 	length := len(nums)
// 	if length < 3 {
// 		return nil
// 	}

// 	res := [][]int{}
// 	sort.Ints(nums)
// 	for i := 0; i < length; i++ {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			//新元素和上一个元素相等时跳过
// 			continue
// 		}

// 		l, r := i+1, length-1
// 		for r > l {
// 			s := nums[i] + nums[l] + nums[r]
// 			if s < 0 { //有序数组，相加小于0，需要把左指针右移
// 				l += 1
// 			} else if s > 0 { //有序数组，相加大于0，需要把右指针左移
// 				r -= 1
// 			} else {
// 				res = append(res, []int{nums[i], nums[l], nums[r]})
// 				for r > l {
// 					r -= 1
// 					//尾部指针移动一步，如果相同则再移动一步
// 					if nums[r] != nums[r+1] {
// 						break
// 					}
// 				}

// 				for r > l {
// 					l += 1
// 					//尾部指针移动一步，如果相同则再移动一步
// 					if nums[l] != nums[l-1] {
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return res
// }
