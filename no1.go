package leetcode

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

//

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func twoSum(nums []int, target int) []int {

	resMap := make(map[int]int)
	for index, num := range nums {
		idx, ok := resMap[target-num]
		if ok {
			return []int{idx, index}
		}
		resMap[num] = index
	}

	return nil
}

// func twoSum(nums []int, target int) []int {

// 	resMap := make(map[int]int)
// 	for index, num := range nums {
// 		indexList, ok := resMap[num]
// 		if ok {
// 			resMap[num] = append(indexList, index)
// 		} else {
// 			resMap[num] = []int{index}
// 		}
// 	}

// 	for index, num := range nums {
// 		num2 := target - num

// 		i, ok := resMap[num2]
// 		if ok {
// 			for _, idx := range i {
// 				if idx != index {
// 					return []int{index, idx}
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }
