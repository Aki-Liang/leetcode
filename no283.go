package leetcode

import "fmt"

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 示例:

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/move-zeroes
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if 0 != nums[j] {
			nums[i] = nums[j]
			if j > i {
				nums[j] = 0
			}
			i++
		}
	}
	fmt.Println(nums)
}

// func moveZeroes(nums []int) {
// 	i := 0
// 	for j := 0; j < len(nums); j++ {
// 		if 0 != nums[j] {
// 			nums[i] = nums[j]
// 			if j > i {
// 				nums[j] = 0
// 			}
// 			i++
// 		}
// 	}
// 	fmt.Println(nums)
// }

// func moveZeroeV2(nums []int) {
// 	for i := 0; i < len(nums); i++ {
// 		if 0 != nums[i] {
// 			continue
// 		}
// 		for j := i + 1; j < len(nums); j++ {
// 			if 0 != nums[j] {
// 				nums[j], nums[i] = nums[i], nums[j]
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(nums)
// }
